package main

import (
	"bloomberg/transaction/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := lib.OpenDatabase()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
	}))

	registerRoutes(r, db)
	r.Run()
}
