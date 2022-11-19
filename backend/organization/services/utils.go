package services

import "github.com/gin-gonic/gin"

const uuidHeaderField = "x-user-id"

func getHeaderUuid(c *gin.Context) (string, bool) {
	uuid, err := c.Request.Header[uuidHeaderField]
	return uuid[0], err
}
