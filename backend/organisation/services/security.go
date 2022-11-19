package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CreateSecuritySql   string = "INSERT INTO securities (id, name, description, creator, ttl_1, ttl_2, funding_goal) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	GetSecuritySql      string = "SELECT  s.*, p.sell_price, 0 AS qty FROM  securities s LEFT JOIN ( SELECT m.security, m.sell_price FROM matches m WHERE m.id = (SELECT m2.id FROM matches m2 WHERE m2.security = m.security ORDER BY creation_date DESC LIMIT 1) ) AS p ON p.security = s.id WHERE s.id = $1"
	GetAllSecuritiesSql string = "SELECT  s.*,  p.sell_price, tq.qty FROM  securities s,  (   SELECT    tmp.security,    SUM(tmp.qty) AS qty   FROM (    SELECT      m.security,      (CASE       WHEN buyer = $1 THEN m.quantity      ELSE (-1) * m.quantity     END) AS qty    FROM matches m WHERE buyer = $1 OR seller = $1   ) AS tmp   GROUP BY tmp.security  ) AS tq,  (   SELECT m.security, m.sell_price FROM matches m WHERE m.id = (SELECT m2.id FROM matches m2 WHERE m2.security = m.security ORDER BY creation_date DESC LIMIT 1)  ) AS p WHERE tq.security = s.id AND p.security = s.id"
	DeleteSecuritySql   string = "DELETE FROM securities WHERE id = $1 AND creator = $2"
	SearchSecuritySql   string = "SELECT s.id, s.name FROM securities s WHERE s.name LIKE ('%' || $1 || '%') LIMIT 10"
)

type CreateSecurityReqest struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
	FundingGoal uint64 `json:"FundingGoal" binding:"required"`
	TtlPhase1   uint64 `json:"TtlPhase1" binding:"required"`
	TtlPhase2   uint64 `json:"TtlPhase2" binding:"required"`
}

type Security struct {
	SecurityId    string `json:"security_id"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Creator       string `json:"creator" binding:"required"`
	CreationDate  uint64 `json:"creationDate" binding:"required"`
	TtlPhase1     uint64 `json:"ttl_phase_one" binding:"required"`
	TtlPhase2     uint64 `json:"ttl_phase_two" binding:"required"`
	FundingAmount uint64 `json:"fundingAmount" binding:"required"`
	FundingDate   uint64 `json:"fundingDate" binding:"required"`
	Price         uint64 `json:"price" binding:"required"`
	Quantity      uint64 `json:"quantity"`
}

type SecuritySearchRequest struct {
	Query string `json:"Query" binding:"required"`
}

type RawSecurity struct {
	Id   string `json:"Id" binding:"required"`
	Name string `json:"Name" binding:"required"`
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
	var price sql.NullInt64
	var security Security
	err := rows.Scan(&security.Id, &security.Name, &security.Description, &security.Creator, &creationDate, &security.TtlPhase1, &security.TtlPhase2, &security.FundingGoal, &fundingDate, &price, &security.Quantity)
	if err != nil {
		fmt.Println(err)
		return security, err
	}
	security.CreationDate = uint64(creationDate.Unix())
	if fundingDate.Valid {
		security.FundingDate = uint64(fundingDate.Int64)
	}
	if price.Valid {
		security.Price = uint64(price.Int64)
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

func SearchSecurity(c *gin.Context, db *sql.DB) {

	var query = c.Query("query")

	//Query the database for the security or all securities
	rows, err := db.Query(SearchSecuritySql, query)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//Parse results
	securities := make([]RawSecurity, 0)
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		securities = append(securities, RawSecurity{id, name})
	}

	c.JSON(http.StatusOK, securities)
}
