package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
)

// 📌 Obtener todos los usuarios
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := db.DB.Preload("Services").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// 📌 Crear nuevo usuario
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// 📌 Obtener un usuario por ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.DB.Preload("Services").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// 📌 Actualizar usuario por ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// 📌 Eliminar usuario por ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DB.Delete(&models.User{}, idNum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
