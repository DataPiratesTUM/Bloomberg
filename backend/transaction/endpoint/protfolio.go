package endpoint

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var portfolioQuery = `
WITH security_value AS (
    SELECT
        security,
        creation_date AS t,
        sell_price AS value
    FROM matches
    GROUP BY security, creation_date, sell_price
), security_amount AS (
    SELECT
        security,
        creation_date AS t,
        SUM(SUM(
            CASE WHEN buyer = $1 THEN
                quantity
            WHEN seller = $1 THEN
                -quantity
            ELSE 
                0
            END
        )) OVER (PARTITION BY security ORDER BY creation_date) AS amount
	FROM matches
    GROUP BY security, creation_date
)
SELECT 
    v.security,
    v.t,
    v.value * a.amount AS v
FROM security_value AS v, security_amount AS a
WHERE v.security = a.security AND v.t = a.t ORDER BY v.t ASC ;
`

func PortfolioValue(c *gin.Context, db *sql.DB) {
	user, ok := getUser(c)
	if !ok {
		return
	}

	rows, err := db.Query(portfolioQuery, user)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	values := make(map[string]int64, 0)
	valuesJson := make([]gin.H, 0)

	for rows.Next() {
		var security string
		var time time.Time
		var value int64

		if err := rows.Scan(&security, &time, &value); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		values[security] = value
		fmt.Println(values)

		var valueSum int64 = 0
		for _, v := range values {
			valueSum += v
		}

		valuesJson = append(valuesJson, gin.H{
			"time":  time.Unix(),
			"value": valueSum,
		})
	}

	c.JSON(http.StatusOK, valuesJson)
}
