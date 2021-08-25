package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	mysql         = "mysql"
	postgresql    = "postgresql"
	mssql         = "mssql"
	oraclesql     = "oraclesql"
	mysqlUrl      = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	postgresqlUrl = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
)

type DatabaseSQL struct {
	db           *sql.DB
	dbName       string
	url          string
	instanceName string
}

var ConnectInstance *DatabaseSQL

func NewInstance(config Config) *DatabaseSQL {
	c := &DatabaseSQL{}
	c.instance(config)
	return c
}

func (c *DatabaseSQL) GetConnectInstance() interface{} {
	return c.db
}

func (c *DatabaseSQL) instance(config Config) {
	switch config.DbDriver {
	case "mysql":
		c.instanceName = config.DbDriver
		c.url = config.DbDriver
		c.connect(config)
	case "postgres":
		c.instanceName = config.DbDriver
		c.url = postgresqlUrl
		c.connect(config)
	default:
		log.Fatal(fmt.Sprintf("not handler this database %s", config.DbDriver))
	}
}

func (c *DatabaseSQL) connect(config Config) {
	address := fmt.Sprintf(
		c.url,
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	db, err := sql.Open(config.DbDriver, address)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	c.db = db
}

func (c *DatabaseSQL) Disconnect() {
	err := c.db.Close()
	if err != nil {
		log.Fatal("Failed to connect to  database", err)
	}
	log.Println("Disconnected from database")
}
