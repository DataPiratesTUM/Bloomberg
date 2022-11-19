package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	registerRoutes(r)
	r.Run()
}
