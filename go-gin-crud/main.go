package main

import (
	"database/sql"
	"fmt"
)

func init() {
	postgresqlInit()
}

func postgresqlInit() {
	config := loadConfigFileName(".", ".env")
	fmt.Println(config)
	ConnectInstance = NewInstance(config)
	if ConnectInstance.GetConnectInstance().(*sql.DB) != nil {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connect")
	}

}

func main() {
	//	defer ConnectInstance.Disconnect()

}
