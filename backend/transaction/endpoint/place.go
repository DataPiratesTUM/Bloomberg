package endpoint

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type body struct {
	Security string `json:"security" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Side     string `json:"side" binding:"required"`
}

type order struct {
	security string
	user     string
	quantity int64
	price    int64
}

type transaction struct {
	buyer     string
	seller    string
	buyPrice  int64
	sellPrice int64
	quantity  int64
}

func createMatch(tx *sql.Tx, security string, tr *transaction) error {
	_, err := tx.Exec(
		`INSERT INTO matches ("buyer", "seller", "buy_price", "sell_price", "security", "quantity") VALUES ($1, $2, $3, $4, $5, $6);`,
		tr.buyer,
		tr.seller,
		tr.buyPrice,
		tr.sellPrice,
		security,
		tr.quantity,
	)

	return err
}

func findMatches(tx *sql.Tx, user string, security string, side bool, price int64, quantity int64) ([]*transaction, error) {
	transactions := make([]*transaction, 0)

	var rows *sql.Rows

	if side {
		r, err := tx.Query(
			`SELECT "security", "user", "quantity", "price" FROM "open_orders" WHERE "security" = $1 AND NOT side AND "price" <= $2 AND "quantity" > 0;`,
			security,
			price,
		)
		if err != nil {
			return nil, err
		}

		rows = r
	} else {
		r, err := tx.Query(
			`SELECT "security", "user", "quantity", "price" FROM "open_orders" WHERE "security" = $1 AND side AND "price" >= $2 AND "quantity" > 0;`,
			security,
			price,
		)
		if err != nil {
			return nil, err
		}

		rows = r
	}
	defer rows.Close()

	for rows.Next() && quantity > 0 {
		var other order
		err := rows.Scan(&other.security, &other.user, &other.quantity, &other.price)
		if err != nil {
			return nil, err
		}

		var transactionQuantity int64
		if other.quantity >= quantity {
			transactionQuantity = quantity
		} else {
			transactionQuantity = other.quantity
		}

		quantity -= transactionQuantity

		if side {
			transactions = append(transactions, &transaction{
				buyer:     user,
				seller:    other.user,
				buyPrice:  price,
				sellPrice: other.price,
				quantity:  transactionQuantity,
			})
		} else {
			transactions = append(transactions, &transaction{
				buyer:     other.user,
				seller:    user,
				buyPrice:  other.price,
				sellPrice: price,
				quantity:  transactionQuantity,
			})
		}
	}

	for _, tr := range transactions {
		if err := createMatch(tx, security, tr); err != nil {
			return nil, err
		}
	}

	return transactions, nil
}

func PlaceOrder(c *gin.Context, db *sql.DB) {
	user, ok := getUser(c)
	if !ok {
		return
	}

	var body body
	if err := c.BindJSON(&body); err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	orderSide := body.Side == "buy"

	//Check if the user has enourgh shares to sell
	if !orderSide {
		sn, snErr := db.Query(
			`SELECT 
				SUM(
					CASE
					WHEN buyer = $1 THEN quantity
					ELSE -quantity
					END
				) AS count 
			FROM matches m
			WHERE security = $2 AND (buyer = $1 OR seller = $1)
			GROUP BY security;`,
			user,
			body.Security,
		)
		if snErr != nil {
			sendError(c, http.StatusInternalServerError, snErr)
			return
		}
		defer sn.Close()

		if !sn.Next() {
			c.Status(http.StatusForbidden)
			return
		}
		var amount int64
		e := sn.Scan(&amount)
		if e != nil {
			sendError(c, http.StatusInternalServerError, e)
			return
		}

		if amount < body.Quantity {
			c.Status(http.StatusForbidden)
			return
		}
	}

	//Check if the security is in phase 1 and funding_date is not set. THen calculate current price based on a linear function. Check if funding has been reached
	if orderSide {
		rows, err := db.Query(
			`SELECT s.creation_date, s.ttl_1, s.funding_date, funding_goal, funding_remaining FROM securities s WHERE s.id = $1`,
			body.Security,
		)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()
		if !rows.Next() {
			c.Status(http.StatusNotFound)
			return
		}
		var creationDate time.Time
		var ttl1 int64
		var funding_date sql.NullTime
		var funding_goal int64
		var funding_remaining int64

		if err = rows.Scan(&creationDate, &ttl1, &funding_date, &funding_goal, &funding_remaining); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		//Security failed if phase 1 is over and funding_date is not set
		if !funding_date.Valid {
			if time.Now().After(creationDate.Add(time.Duration(ttl1) * time.Second)) {
				c.Status(http.StatusBadRequest)
				return
			} else {
				diff := time.Since(creationDate).Seconds()

				m := (-funding_goal / 1000) / (ttl1 / 86400)
				currentPrice := m*((ttl1-int64(diff))/86400) + (funding_goal / 1000)
				if currentPrice < 0 {
					currentPrice = 0
				}

				//If the price is to high dont buy
				if currentPrice < body.Price {

					var tmp int64
					fullyFunded := false
					if funding_remaining-body.Quantity*currentPrice < currentPrice {
						tmp = body.Quantity - funding_remaining/currentPrice
						body.Quantity = funding_remaining / currentPrice
						fullyFunded = true
					}
					fmt.Println(body.Quantity)

					//Insert a match without a seller to indicate initial shares
					_, err = tx.Exec(
						`INSERT INTO "matches" ("security", "buyer", "buy_price", "sell_price", "quantity") VALUES ($1, $2, $3, $4, $5);`,
						body.Security,
						user,
						body.Price,
						currentPrice,
						body.Quantity,
					)

					if err != nil {
						sendError(c, http.StatusInternalServerError, err)
						tx.Rollback()
						return
					}

					//Check if phase 1 should be ended
					if fullyFunded {
						_, err = tx.Exec(
							`UPDATE securities SET funding_date = $1 WHERE id = $2`,
							time.Now(),
							body.Security,
						)
						if err != nil {
							sendError(c, http.StatusInternalServerError, err)
						}
					}

					if tmp > 0 {
						body.Quantity = tmp
					}
				}

			}

		}
	}

	if body.Quantity > 0 {
		_, err = tx.Exec(
			`INSERT INTO "orders" ("security", "quantity", "price", "side", "user") VALUES ($1, $2, $3, $4, $5);`,
			body.Security,
			body.Quantity,
			body.Price,
			orderSide,
			user,
		)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			tx.Rollback()
			return
		}

		transactions, err := findMatches(tx, user, body.Security, orderSide, body.Price, body.Quantity)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			tx.Rollback()
			return
		}

		var transactionJson []gin.H = make([]gin.H, len(transactions))
		for i, tr := range transactions {
			transactionJson[i] = gin.H{"quantity": tr.quantity, "price": tr.sellPrice}
		}

		c.JSON(http.StatusOK, transactionJson)
	} else {
		err = tx.Commit()
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			tx.Rollback()
			return
		}

		c.Status(http.StatusOK)
	}
}
