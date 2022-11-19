package endpoint

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func OrderHistory(c *gin.Context, db *sql.DB) {
	_, ok := getUser(c)
	if !ok {
		return
	}

}
