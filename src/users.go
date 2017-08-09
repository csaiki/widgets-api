package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// User defines the struct of the user data
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

// ReturnedUser defines the struct to be returned by the API
type ReturnedUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
}

// GetUsers is responsible for returning all the users
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
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No users found!"})
		return
	}

	for _, user := range users {
		_users = append(_users, ReturnedUser{ID: user.ID, Name: user.Name, Gravatar: user.Gravatar})
	}

	// Display JSON result
	c.JSON(200, _users)

	// curl -i http://localhost:8080/users
}

// GetUserByID is responsible for returning the specified user
func GetUserByID(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	db.LogMode(true)

	// Close connection database
	defer db.Close()

	var user User
	var _user ReturnedUser

	userID := c.Param("id")
	// SELECT * FROM users
	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	_user = ReturnedUser{ID: user.ID, Name: user.Name, Gravatar: user.Gravatar}

	// Display JSON result
	c.JSON(200, _user)

	// curl -i http://localhost:8080/users
}
