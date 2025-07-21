package handlers

import (
	"net/http"
	"strconv"

	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
	"github.com/gin-gonic/gin"
)

// 📌 Obtener todos los servicios
func GetServices(c *gin.Context) {
	var services []models.Service
	if err := db.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// 📌 Crear nuevo servicio
func CreateService(c *gin.Context) {
	var service models.Service

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que user_id es mayor que cero
	if service.UserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id inválido"})
		return
	}

	// Verificar que el usuario existe en la base de datos
	var existingUser models.User
	if err := db.DB.First(&existingUser, service.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := db.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, service)
}

// 📌 Obtener un servicio por ID
func GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

// 📌 Actualizar servicio por ID
func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&service)
	c.JSON(http.StatusOK, service)
}

// 📌 Eliminar servicio por ID
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DB.Delete(&models.Service{}, idNum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}
