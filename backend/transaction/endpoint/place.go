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
		"INSERT INTO \"orders\" (\"security\", \"quantity\", \"price\", \"side\", \"user\") VALUES ($1, $2, $3, $4, $5);",
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

	c.Status(http.StatusOK)
}
