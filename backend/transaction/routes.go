package main

import "github.com/gin-gonic/gin"

func registerRoutes(r *gin.Engine) {
	r.POST("order/buy")
	r.POST("order/sell")

	r.DELETE("order/:id")
}
