package main

import (
	"bloomberg/organisation/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, db *sql.DB) {
	/*
		Routes to manage users
	*/

	r.GET("/user/:id", func(c *gin.Context) {
		services.GetUser(c, db)
	})
	r.POST("/user/:id/balance", func(c *gin.Context) {
		services.ChangeBalance(c, db)
	})
	/*
		Routes to manage organisations
	*/
	r.GET("/organization/:id")

	/*
		Routes to manage securities
	*/
	r.POST("/security/create")
	r.GET("/security/:id")    // id is security id or "all"
	r.DELETE("/security/:id") //only possible in the first phase
}
