package main

import (
	"github.com/gin-gonic/gin"

	controllers "shorturl/controller"
	"shorturl/models"
)

func main() {
	r := gin.Default()

	r.GET("/:alias", controllers.GetAlias)

	r.GET("/url", controllers.FindUrl)
	r.POST("/url", controllers.CreateUrl)
	r.PATCH("/url/:id", controllers.UpdateUrl)
	r.DELETE("/url/:id", controllers.DeleteUrl)

	models.ConnectDataBase()

	r.Run(":8090")
}
