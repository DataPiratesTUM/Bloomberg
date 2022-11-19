package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GetUserSql       string = "SELECT u.name, u.balance FROM users u WHERE u.id = $1"
	AddBalanceSQL           = "UPDATE users SET balance = (balance + $1) WHERE id = $2"
	RemoveBalanceSQL        = "UPDATE users SET balance = (balance - $1) WHERE id = $2"
)

type BalanceRequest struct {
	IsWithdraw bool   `json:"IsWithdraw"`
	Amount     uint64 `json:"Amount"`
}

type User struct {
	Name    string `json:"Name"`
	Balance uint64 `json:"Balance"`
}

func GetUser(c *gin.Context, db *sql.DB) {
	uuid := c.Param("id")

	//Query the database for the user
	rows, err := db.Query(GetUserSql, uuid)
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
	var user User
	err = rows.Scan(&user.Name, &user.Balance)
	if err != nil {
		c.Status(http.StatusInternalServerError)
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
