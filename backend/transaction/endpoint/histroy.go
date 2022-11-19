package endpoint

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type match struct {
	sellPrice    int64
	quantity     int64
	security     string
	creationDate time.Time
}

var historySelect = `SELECT "sell_price", "quantity", "security", "creation_date" FROM "matches"`

func history(query func() (*sql.Rows, error)) ([]*match, error) {
	matches := make([]*match, 0)

	rows, err := query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		match := &match{}
		if err := rows.Scan(&match.sellPrice, &match.quantity, &match.security, &match.creationDate); err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func SecurityHistory(c *gin.Context, db *sql.DB) {
	security := c.Param("id")

	matches, err := history(func() (*sql.Rows, error) {
		return db.Query(
			historySelect+`WHERE "security" = $1 ORDER BY "creation_date"`,
			security,
		)
	})
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	matchesJson := make([]gin.H, len(matches))
	for i, m := range matches {
		matchesJson[i] = gin.H{
			"quantity": m.quantity,
			"price":    m.sellPrice,
			"created":  uint64(m.creationDate.Unix()),
		}
	}

	c.JSON(http.StatusOK, matchesJson)
}

func OrderHistory(c *gin.Context, db *sql.DB) {
	user, ok := getUser(c)
	if !ok {
		return
	}

	matches, err := history(func() (*sql.Rows, error) {
		return db.Query(
			historySelect+`WHERE "buyer" = $1 OR "seller" = $1 ORDER BY "creation_date"`,
			user,
		)
	})
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	matchesJson := make([]gin.H, len(matches))
	for i, m := range matches {
		matchesJson[i] = gin.H{
			"quantity": m.quantity,
			"price":    m.sellPrice,
			"created":  m.creationDate.Unix(),
			"security": m.security,
		}
	}

	c.JSON(http.StatusOK, matchesJson)
}

func AllHistory(c *gin.Context, db *sql.DB) {
	matches, err := history(func() (*sql.Rows, error) {
		return db.Query(historySelect + `ORDER BY "creation_date"`)
	})
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	matchesJson := make([]gin.H, len(matches))
	for i, m := range matches {
		matchesJson[i] = gin.H{
			"quantity": m.quantity,
			"price":    m.sellPrice,
			"created":  m.creationDate.Unix(),
			"security": m.security,
		}
	}

	c.JSON(http.StatusOK, matchesJson)
}
