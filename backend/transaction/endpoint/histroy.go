package endpoint

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type match struct {
	buyer        string
	buyPrice     int
	seller       string
	sellPrice    int
	quantity     int
	creationDate time.Time
}

func securityHistory(db *sql.DB, security string) ([]*match, error) {
	matches := make([]*match, 0)

	rows, err := db.Query(
		`SELECT "buyer", "buy_price", "seller", "sell_price", "quantity", "creation_date" FROM "matches" WHERE "security" = $1 ORDER BY "creation_date"`,
		security,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		match := &match{}
		if err := rows.Scan(&match.buyer, &match.buyPrice, &match.seller, &match.sellPrice, &match.quantity, &match.creationDate); err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func SecurityHistory(c *gin.Context, db *sql.DB) {
	security := c.Param("id")

	matches, err := securityHistory(db, security)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	var matchesJson []gin.H = make([]gin.H, len(matches))
	for i, m := range matches {
		matchesJson[i] = gin.H{
			"quantity": m.quantity,
			"price":    m.sellPrice,
			"created":  m.creationDate.Unix(),
		}
	}

	c.JSON(http.StatusOK, matchesJson)
}
