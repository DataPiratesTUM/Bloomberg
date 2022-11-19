package services

import "github.com/gin-gonic/gin"

func GetOrganization(c *gin.Context) {
	uuid := c.Param("id")
	_ = uuid
}
