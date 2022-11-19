package endpoint

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body struct {
	Security string `json:"security" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Side     string `json:"side" binding:"required"`
}

type order struct {
	security string
	user     string
	quantity int
	price    int
}

type transaction struct {
	buyer     string
	seller    string
	buyPrice  int
	sellPrice int
	quantity  int
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

func match(tx *sql.Tx, user string, security string, side bool, price int, quantity int) ([]*transaction, error) {
	transactions := make([]*transaction, 0)

	var rows *sql.Rows

	if side {
		r, err := tx.Query(
			`SELECT security, "user", quantity, price FROM "open_orders" WHERE "security" = $1 AND NOT side AND "price" <= $2;`,
			security,
			price,
		)
		if err != nil {
			return nil, err
		}

		rows = r
	} else {
		r, err := tx.Query(
			`SELECT security, "user", quantity, price FROM "open_orders" WHERE "security" = $1 AND side AND "price" >= $2;`,
			security,
			price,
		)
		if err != nil {
			return nil, err
		}

		rows = r
	}

	for rows.Next() && quantity > 0 {
		var other order
		err := rows.Scan(&other.security, &other.user, &other.quantity, &other.price)
		if err != nil {
			return nil, err
		}

		var transactionQuantity int
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

	if err := rows.Close(); err != nil {
		return nil, err
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

	orderSide := body.Side == "buy"

	tx, err := db.Begin()
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

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

	if body.Quantity > 0 {
		transactions, err := match(tx, user, body.Security, orderSide, body.Price, body.Quantity)
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
