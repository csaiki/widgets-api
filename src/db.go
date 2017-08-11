package src

import "github.com/jinzhu/gorm"

// InitDb connects to a MySQL database
func InitDb() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "chris:1234@tcp(db:3306)/widgetsapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
