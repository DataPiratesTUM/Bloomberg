package endpoint

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var trendingQuery = `
SELECT 
	s.id
FROM matches AS m, securities AS s
WHERE m.security = s.id AND s.creation_date + s.ttl_2 * interval '1 second' > now()
GROUP BY s.id
ORDER BY SUM(m.quantity)
LIMIT 5;
`

func Trending(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(trendingQuery)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	valuesJson := make([]string, 0)
	for rows.Next() {
		var security string

		if err := rows.Scan(&security); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		valuesJson = append(valuesJson, security)
	}

	c.JSON(http.StatusOK, valuesJson)
}
