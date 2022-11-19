package endpoint

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body struct {
	Request  string `json:"request" binding:"required"`
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
	fmt.Println(body)

	orderType := body.Request == "add"
	orderSide := body.Side == "buy"

	rows, err := db.Query(
		"INSERT INTO \"order\" (\"type\", \"security\", \"quantity\", \"price\", \"side\", \"user\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING \"id\";",
		orderType,
		body.Security,
		body.Quantity,
		body.Price,
		orderSide,
		user,
	)
	if err != nil || !rows.Next() {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	var id string
	if err := rows.Scan(&id); err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
