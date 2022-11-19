package main

import (
	"bloomberg/organisation/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := lib.OpenDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	registerRoutes(r, db)
	r.Run()
}
