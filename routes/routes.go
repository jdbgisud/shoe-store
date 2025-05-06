package routes

import (
	"github.com/gin-gonic/gin"
	"shoe-store/controllers"
	"shoe-store/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Аутентификация
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Защищенные маршруты
		shoes := api.Group("/shoes")
		shoes.Use(middleware.AuthMiddleware())
		{
			shoes.GET("", controllers.GetShoes)
			shoes.GET("/:id", controllers.GetShoe)
			shoes.POST("", middleware.RoleMiddleware("ADMIN"), controllers.CreateShoe)
			shoes.PUT("/:id", middleware.RoleMiddleware("ADMIN"), controllers.UpdateShoe)
			shoes.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), controllers.DeleteShoe)
		}

	}
}
