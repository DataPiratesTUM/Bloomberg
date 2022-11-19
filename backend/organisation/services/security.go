package services

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CreateSecuritySql   string = "INSERT INTO securities (id, name, description, creator, ttl_1, ttl_2, funding_goal) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	GetSecuritySql      string = "SELECT * FROM securities s WHERE s.id = $1"
	GetAllSecuritiesSql string = "SELECT * FROM securities s, open_orders oo WHERE s.id = oo.security AND oo.user = $1"
	DeleteSecuritySql   string = "DELETE FROM securities WHERE id = $1 AND creator = $2"
)

type CreateSecurityReqest struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
	FundingGoal uint64 `json:"FundingGoal" binding:"required"`
	TtlPhase1   uint64 `json:"TtlPhase1" binding:"required"`
	TtlPhase2   uint64 `json:"TtlPhase2" binding:"required"`
}

type Security struct {
	Id           string `json:"Id"`
	Name         string `json:"Name" binding:"required"`
	Description  string `json:"Description" binding:"required"`
	Creator      string `json:"Creator" binding:"required"`
	CreationDate uint64 `json:"CreationDate" binding:"required"`
	TtlPhase1    uint64 `json:"TtlPhase1" binding:"required"`
	TtlPhase2    uint64 `json:"TtlPhase2" binding:"required"`
	FundingGoal  uint64 `json:"FundingGoal" binding:"required"`
	FundingDate  uint64 `json:"FundingDate" binding:"required"`
}

func CreateSecurity(c *gin.Context, db *sql.DB) {
	userId, ok := getHeaderUuid(c)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	//Query the user to see if he is in an organisation
	user, err := getUser(db, userId)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	//Return forbidden if the user is not in an organisation
	if user.OrganisationId == "" {
		c.Status(http.StatusForbidden)
		return
	}

	//Parses the body to obtain the request as a struct
	var req CreateSecurityReqest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//Insert the new entry into the database
	id := uuid.New()
	res, err := db.Exec(CreateSecuritySql, id, req.Name, req.Description, userId, req.TtlPhase1, req.TtlPhase2, req.FundingGoal)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	_ = res

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func GetSecurityAdapter(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "all" {
		getAllSecurities(c, db)
	} else {
		getSecurity(c, db)
	}
}

func parseSecurity(rows *sql.Rows) (Security, error) {
	var creationDate time.Time
	var fundingDate sql.NullInt64
	var security Security
	err := rows.Scan(&security.Id, &security.Name, &security.Description, &security.Creator, &creationDate, &security.TtlPhase1, &security.TtlPhase2, &security.FundingGoal, &fundingDate)
	if err != nil {
		return security, err
	}
	security.CreationDate = uint64(creationDate.Unix())
	if fundingDate.Valid {
		security.FundingDate = uint64(fundingDate.Int64)
	}
	return security, nil
}

func getAllSecurities(c *gin.Context, db *sql.DB) {
	uuid, ok := getHeaderUuid(c)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	//Query the database for the security or all securities
	rows, err := db.Query(GetAllSecuritiesSql, uuid)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	securities := make([]Security, 0)
	for rows.Next() {
		security, err := parseSecurity(rows)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		securities = append(securities, security)
	}

	c.JSON(http.StatusOK, securities)

}

func getSecurity(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	//Query the database for the security or all securities
	rows, err := db.Query(GetSecuritySql, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//Check if the security exists
	if !rows.Next() {
		c.Status(http.StatusNotFound)
		return
	}

	security, err := parseSecurity(rows)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, security)

}

func DeleteSecurity(c *gin.Context, db *sql.DB) {
	uuid, ok := getHeaderUuid(c)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	id := c.Param("id")

	res, err := db.Exec(DeleteSecuritySql, id, uuid)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	//Check if the security has been deleted
	affected, err := res.RowsAffected()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if affected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}