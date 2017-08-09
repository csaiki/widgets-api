package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	src "github.com/csaiki/widgets-api/src"
)

func main() {
	router := gin.Default()

	usersRoutes := router.Group("/users")
	{
		usersRoutes.GET("", src.GetUsers)
		usersRoutes.GET("/:id", src.GetUserByID)
	}

	widgetsRoutes := router.Group("/widgets")
	{
		widgetsRoutes.GET("", src.GetWidgets)
		widgetsRoutes.POST("", src.AddWidgets)
		widgetsRoutes.GET(":id", src.GetWidgetByID)
		widgetsRoutes.PUT(":id", src.UpdateWidgetByID)
	}

	router.Run(":8080")

}
