package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var userHeader = "X-User-Id"

func getUser(c *gin.Context) (string, bool) {
	if user, ok := c.Request.Header[userHeader]; ok && len(user) == 1 {
		return user[0], true
	} else {
		c.Status(http.StatusUnauthorized)
		return "", false
	}
}

func sendError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"msg": err.Error()})
}
