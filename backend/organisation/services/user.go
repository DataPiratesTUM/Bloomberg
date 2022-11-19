package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GetUserSql       string = "SELECT u.name, u.balance, u.organisation FROM users u WHERE u.id = $1"
	AddBalanceSQL           = "UPDATE users SET balance = (balance + $1) WHERE id = $2"
	RemoveBalanceSQL        = "UPDATE users SET balance = (balance - $1) WHERE id = $2"
)

type BalanceRequest struct {
	IsWithdraw bool   `json:"IsWithdraw" binding:"required"`
	Amount     uint64 `json:"Amount" binding:"required"`
}

type User struct {
	Name           string `json:"Name" binding:"required"`
	Balance        uint64 `json:"Balance" binding:"required"`
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
	err = rows.Scan(&user.Name, &user.Balance, &orgIdNull)
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

func ChangeBalance(c *gin.Context, db *sql.DB) {
	//Gets the uuid of the requesting user (Header: "x-user-id")
	uuid := c.Param("id")

	//Parses the body to obtain the request as a struct
	var req BalanceRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	query := AddBalanceSQL
	if req.IsWithdraw {
		query = RemoveBalanceSQL
	}
	res, err := db.Exec(query, req.Amount, uuid)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	_ = res
	c.Status(http.StatusOK)
}
