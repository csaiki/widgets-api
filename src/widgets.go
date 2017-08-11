package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Widget defines the MySQL schema for Widget
type Widget struct {
	gorm.Model
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Inventory int    `json:"inventory"`
	Melts     bool   `json:"melts"`
}

// ReturnedWidget defines the struct to be returned by the API
type ReturnedWidget struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Inventory int    `json:"inventory"`
	Melts     bool   `json:"melts"`
}

// GetWidgets is responsible for returning all the users
func GetWidgets(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	db.LogMode(true)

	// Close connection database
	defer db.Close()

	var widgets []Widget
	var _widgets []ReturnedWidget
	// SELECT * FROM users
	db.Find(&widgets)

	if len(widgets) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No widgets found!"})
		return
	}

	for _, widget := range widgets {
		_widgets = append(_widgets, ReturnedWidget{
			ID:        widget.ID,
			Name:      widget.Name,
			Color:     widget.Color,
			Price:     widget.Price,
			Inventory: widget.Inventory,
			Melts:     widget.Melts})
	}

	// Display JSON result
	c.JSON(200, _widgets)

	// curl -i http://localhost:8080/widgets
}

// GetWidgetByID is responsible for returning the specified user
func GetWidgetByID(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	db.LogMode(true)

	// Close connection database
	defer db.Close()

	var widget Widget
	var _widget ReturnedWidget

	widgetID := c.Param("id")
	// SELECT * FROM users
	db.First(&widget, widgetID)

	if widget.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No widget found!"})
		return
	}

	_widget = ReturnedWidget{
		ID:        widget.ID,
		Name:      widget.Name,
		Color:     widget.Color,
		Price:     widget.Price,
		Inventory: widget.Inventory,
		Melts:     widget.Melts}

	// Display JSON result
	c.JSON(200, _widget)

}

// UpdateWidgetByID is responsible for returning the specified user
func UpdateWidgetByID(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	db.LogMode(true)

	// Close connection database
	defer db.Close()

	var widget Widget

	widgetID := c.Param("id")
	// SELECT * FROM widgets where id = {id}
	db.First(&widget, widgetID)

	if widget.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No widget found!"})
		return
	}

	var updatedWidget Widget
	c.Bind(&updatedWidget)
	db.Model(&widget).Updates(updatedWidget)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Widget updated successfully!"})
}

// AddWidgets is responsible for returning the specified user
func AddWidgets(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	db.LogMode(true)

	// Close connection database
	defer db.Close()

	var widgets []Widget
	c.Bind(&widgets)

	for _, widget := range widgets {
		db.Create(&widget)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Widgets created successfully!"})

}
