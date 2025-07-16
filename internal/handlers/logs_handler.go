package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
)

// Obtener todos los logs
func GetLogs(c *gin.Context) {
	var logs []models.Log
	if err := db.DB.Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// Crear nuevo log
func CreateLog(c *gin.Context) {
	var logEntry models.Log
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&logEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, logEntry)
}

// Obtener log por ID
func GetLogByID(c *gin.Context) {
	id := c.Param("id")
	var logEntry models.Log
	if err := db.DB.First(&logEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}
	c.JSON(http.StatusOK, logEntry)
}

// Actualizar log
func UpdateLog(c *gin.Context) {
	id := c.Param("id")
	var logEntry models.Log
	if err := db.DB.First(&logEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&logEntry)
	c.JSON(http.StatusOK, logEntry)
}

// Eliminar log
func DeleteLog(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DB.Delete(&models.Log{}, idNum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Log deleted"})
}
