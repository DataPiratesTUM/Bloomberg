package main

import (
	"bloomberg/transaction/endpoint"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("order/place", func(c *gin.Context) {
		endpoint.PlaceOrder(c, db)
	})

	r.GET("order/history/security/:id", func(c *gin.Context) {
		endpoint.SecurityHistory(c, db)
	})

	r.GET("order/history", func(c *gin.Context) {
		endpoint.OrderHistory(c, db)
	})

	r.GET("match/history", func(c *gin.Context) {
		endpoint.AllHistory(c, db)
	})

	r.GET("order/value", func(c *gin.Context) {
		endpoint.PortfolioValue(c, db)
	})

	r.GET("trending", func(c *gin.Context) {
		endpoint.Trending(c, db)
	})

	r.GET("open_orders/:id", func(c *gin.Context) {
		endpoint.OpenOrders(c, db)
	})
}
