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
		services.GetUserAdapter(c, db)
	})

	/*
		Routes to manage organisations
	*/
	r.GET("/organisation/:id", func(c *gin.Context) {
		services.GetOrganization(c, db)
	})

	/*
		Routes to manage securities
	*/
	r.POST("/security/create", func(c *gin.Context) {
		services.CreateSecurity(c, db)
	})
	r.GET("/security/:id", func(c *gin.Context) {
		services.GetSecurityAdapter(c, db)
	}) // id is security id or "all"
	r.DELETE("/security/:id", func(c *gin.Context) {
		services.DeleteSecurity(c, db)
	}) //only possible in the first phase
	r.GET("/security/search/title", func(c *gin.Context) {
		services.SearchSecurity(c, db)
	})
}
