package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GetUserSql    string = "SELECT * FROM users WHERE uuid = ?"
	AddBalanceSQL        = "UPDATE users SET balance = (balance + ?)"
)

const (
	Withdraw uint8 = 0
	Deposit        = 1
)

type BalanceRequest struct {
	action uint8
	amount uint64
}

func GetUser(c *gin.Context) {
	uuid := c.Param("id")
}

func ChangeBalance(c *gin.Context) {
	//Gets the uuid of the requesting user (Header: "x-user-id")
	uuid, exists := getHeaderUuid(c)
	if !exists {
		c.Status(http.StatusUnauthorized)
		return
	}

	//Parses the body to obtain the request as a struct
	var req BalanceRequest
	if err := c.ShouldBindJSON(&req); err == nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
