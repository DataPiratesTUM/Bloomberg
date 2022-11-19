package services

import (
	"github.com/gin-gonic/gin"
)

const uuidHeaderField = "X-User-Id"

func getHeaderUuid(c *gin.Context) (string, bool) {
	uuid, err := c.Request.Header[uuidHeaderField]
	if !err || len(uuid) == 0 {
		return "", false
	}
	return uuid[0], true
}
