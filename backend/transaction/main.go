package main

import (
	"bloomberg/transaction/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	db := lib.OpenDatabase()
	r := gin.Default()

	registerRoutes(r, db)
	r.Run()
}