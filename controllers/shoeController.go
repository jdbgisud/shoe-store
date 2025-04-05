package controllers

import (
	"net/http"
	"shoe-store/database"
	"shoe-store/models"

	"github.com/gin-gonic/gin"
)

func GetShoes(c *gin.Context) {
	var shoes []models.Shoe
	database.DB.Find(&shoes)
	c.JSON(http.StatusOK, shoes)
}

func GetShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := database.DB.First(&shoe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	c.JSON(http.StatusOK, shoe)
}

func CreateShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := c.ShouldBindJSON(&shoe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&shoe)
	c.JSON(http.StatusCreated, shoe)
}

func UpdateShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := database.DB.First(&shoe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	if err := c.ShouldBindJSON(&shoe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&shoe)
	c.JSON(http.StatusOK, shoe)
}

func DeleteShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := database.DB.First(&shoe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	database.DB.Delete(&shoe)
	c.JSON(http.StatusOK, gin.H{"message": "Товар удалён"})
}
