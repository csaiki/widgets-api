package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database connects to a MySQL database
func InitDb() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "chris:1234@/widgetsapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {
	db := InitDb()
	db.AutoMigrate(&User{}, &Widget{})
	router := gin.Default()

	usersRoutes := router.Group("/users")
	{
		usersRoutes.GET("", GetUsers)
		// users_routes.GET("/:id", GetUserById)
	}

	// widgets_routes := router.Group("/widgets")
	// {
	// widgets_routes.GET("/", GetWidgets)
	// widgets_routes.POST("/", AddWigets)
	// widgets_routes.GET("/:id", GetWidgetById)
	// widgets_routes.PUT(":/id", UpdateWidget)
	// }

	router.Run()

}

// User defines the struct of the user data
type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

type ReturnedUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

// Widget defines the MySQL schema for Widget
type Widget struct {
	gorm.Model
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Color     bool   `json:"color"`
	Price     string `json:"price"`
	Inventory int    `json:"inventory"`
	Melts     bool   `json:"melts"`
}

func GetUsers(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	db.LogMode(true)

	// Close connection database
	defer db.Close()

	var users []User
	var _users []ReturnedUser
	// SELECT * FROM users
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	for _, user := range users {
		_users = append(_users, ReturnedUser{ID: user.ID, Name: user.Name, Gravatar: user.Gravatar})
	}

	// Display JSON result
	c.JSON(200, _users)

	// curl -i http://localhost:8080/api/v1/users
}
