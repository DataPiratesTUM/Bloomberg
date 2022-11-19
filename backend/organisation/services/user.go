package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GetUserSql string = "SELECT u.name, u.organisation FROM users u WHERE u.id = $1"
)

type User struct {
	Name           string `json:"Name" binding:"required"`
	OrganisationId string `json:"OrganisationId"`
}

func getUser(db *sql.DB, uuid string) (User, error) {
	//Query the database for the user
	rows, err := db.Query(GetUserSql, uuid)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	//Check if the user has been found
	if !rows.Next() {
		return User{}, fmt.Errorf("User not found")
	}

	//Try to parse the user data
	var orgIdNull sql.NullString
	var user User
	err = rows.Scan(&user.Name, &orgIdNull)
	if err != nil {
		return User{}, err
	}

	if orgIdNull.Valid {
		user.OrganisationId = orgIdNull.String
	}

	return user, nil
}

func GetUserAdapter(c *gin.Context, db *sql.DB) {
	uuid := c.Param("id")

	user, err := getUser(db, uuid)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}
