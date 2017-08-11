package src

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// InitDb connects to a MySQL database
func InitDb() *gorm.DB {
	var host string
	// If the application is running inside docker it should connect to a different tcp connection
	if os.Getenv("DOCKER") == "true" {
		host = "tcp(db:3306)"
	} else {
		host = ""
	}
	connString := fmt.Sprintf("%s%s%s", "chris:1234@", host, "/widgetsapi?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(connString)
	//open a db connection
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
