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

	/*
		r.GET("order/placed", func(c *gin.Context) {
			endpoint.DeleteOrder(c, db)
		})*/
}
