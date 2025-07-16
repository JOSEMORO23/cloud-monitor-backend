package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
)

// Obtener todas las alertas
func GetAlerts(c *gin.Context) {
	var alerts []models.Alert
	if err := db.DB.Find(&alerts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}

// Crear nueva alerta
func CreateAlert(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&alert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, alert)
}

// Obtener alerta por ID
func GetAlertByID(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	if err := db.DB.First(&alert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}
	c.JSON(http.StatusOK, alert)
}

// Actualizar alerta
func UpdateAlert(c *gin.Context) {
	id := c.Param("id")
	var alert models.Alert
	if err := db.DB.First(&alert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&alert)
	c.JSON(http.StatusOK, alert)
}

// Eliminar alerta
func DeleteAlert(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DB.Delete(&models.Alert{}, idNum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted"})
}
