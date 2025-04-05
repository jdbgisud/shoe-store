package routes

import (
	"github.com/gin-gonic/gin"
	"shoe-store/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/shoes", controllers.GetShoes)
		api.GET("/shoes/:id", controllers.GetShoe)
		api.POST("/shoes", controllers.CreateShoe)
		api.PUT("/shoes/:id", controllers.UpdateShoe)
		api.DELETE("/shoes/:id", controllers.DeleteShoe)
	}
}
