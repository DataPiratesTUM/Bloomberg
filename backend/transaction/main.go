package main

import (
	"net/http"

	"bloomberg/transaction/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	db := lib.OpenDatabase()
	_ = db

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
