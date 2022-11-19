package endpoint

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var openOrderQuery = `
SELECT quantity, price, side, security
FROM open_orders
WHERE security = $1
`

func OpenOrders(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(openOrderQuery, c.Param("id"))
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	valuesJson := make([]gin.H, 0)
	for rows.Next() {
		var quantity int64
		var price int64
		var side bool
		var security string

		if err := rows.Scan(&quantity, &price, &side, &security); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		var sideJson string
		if side {
			sideJson = "buy"
		} else {
			sideJson = "sell"
		}

		valuesJson = append(valuesJson, gin.H{
			"quantity": quantity,
			"price":    price,
			"side":     sideJson,
			"security": security,
		})
	}

	c.JSON(http.StatusOK, valuesJson)
}
