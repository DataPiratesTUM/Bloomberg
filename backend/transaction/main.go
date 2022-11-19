package main

import (
	"net/http"

	"bloomberg/transaction/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Open()
	_ = db

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
