package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Пока простая проверка по хардкоду
		token := c.GetHeader("Authorization")
		if token != "Bearer mysecrettoken" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизован"})
			c.Abort()
			return
		}
		c.Next()
	}
}
