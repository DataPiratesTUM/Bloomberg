package main

import (
	"bloomberg/organization/services"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	/*
		Routes to manage users
	*/

	r.GET("/user/:id", services.GetUser)
	r.POST("/user/:id/balance") //actions deposit, withdraw
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
