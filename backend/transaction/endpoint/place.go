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

func createMatch(db *sql.DB, buyer string, seller string, buyPrice int, sellPrice int, security string, quantity int) error {
	_, err := db.Exec(
		`INSERT INTO matches (buyer, seller, buy_price, sell_price, security, quantity) VALUES ($1, $2, $3, $4, $5, $6);`,
		buyer,
		seller,
		buyPrice,
		sellPrice,
		security,
		quantity,
	)

	return err
}

func matchBuy(db *sql.DB, user string, security string, price int, quantity int) error {
	rows, err := db.Query(
		`SELECT security, "user", quantity, price FROM "open_orders" WHERE "security" = $1 AND NOT side AND "price" <= $2;`,
		security,
		price,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var other order
		err = rows.Scan(&other.security, &other.user, &other.quantity, &other.price)
		if err != nil {
			return err
		}

		if other.quantity >= quantity {
			return createMatch(db, user, other.user, price, other.price, security, quantity)
		} else {
			quantity -= other.quantity

			err = createMatch(db, user, other.user, price, other.price, security, other.quantity)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func matchSell(db *sql.DB, user string, security string, price int, quantity int) error {
	rows, err := db.Query(
		`SELECT security, "user", quantity, price FROM "open_orders" WHERE "security" = $1 AND side AND "price" >= $2;`,
		security,
		price,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var other order
		err = rows.Scan(&other.security, &other.user, &other.quantity, &other.price)
		if err != nil {
			return err
		}

		if other.quantity >= quantity {
			return createMatch(db, other.user, user, other.price, price, security, quantity)
		} else {
			quantity -= other.quantity

			err = createMatch(db, other.user, user, other.price, price, security, other.quantity)
			if err != nil {
				return err
			}
		}
	}

	return nil
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

	_, err := db.Exec(
		`INSERT INTO "orders" ("security", "quantity", "price", "side", "user") VALUES ($1, $2, $3, $4, $5);`,
		body.Security,
		body.Quantity,
		body.Price,
		orderSide,
		user,
	)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	if orderSide {
		err = matchBuy(db, user, body.Security, body.Price, body.Quantity)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}
	} else {
		err = matchSell(db, user, body.Security, body.Price, body.Quantity)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}
	}

	c.Status(http.StatusOK)
}
