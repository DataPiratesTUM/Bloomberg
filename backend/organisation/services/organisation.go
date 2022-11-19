package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GetOrganisationSql string = "SELECT o.name FROM organisation o WHERE o.id = $1"
)

type Organisation struct {
	Name string `json:"Name"`
}

func GetOrganization(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	//Query the database for the user
	rows, err := db.Query(GetOrganisationSql, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//Check if the user has been found
	if !rows.Next() {
		c.Status(http.StatusNotFound)
		return
	}

	//Try to parse the user data
	var user Organisation
	err = rows.Scan(&user.Name)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
