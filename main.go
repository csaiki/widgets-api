package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	src "github.com/csaiki/widgets-api/src"
)

func main() {
	router := gin.Default()

	routes := router.Group("/", gin.BasicAuth(gin.Accounts{
		"rv": "R3D-V3ntur35"}))

	{
		routes.GET("users", src.GetUsers)
		routes.GET("users/:id", src.GetUserByID)
		routes.GET("widgets", src.GetWidgets)
		routes.POST("widgets", src.AddWidgets)
		routes.GET("/widgets/:id", src.GetWidgetByID)
		routes.PUT("/widgets/:id", src.UpdateWidgetByID)
	}

	router.Run(":8080")

}
