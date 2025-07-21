package handlers

import (
	"net/http"
	"strconv"

	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
	"github.com/gin-gonic/gin"
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
		// 🔍 Este print mostrará qué campo está fallando en el bind
		println("Error al hacer el binding del JSON:", err.Error())
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
// Actualizar alerta
func UpdateAlert(c *gin.Context) {
	id := c.Param("id")

	var existingAlert models.Alert
	if err := db.DB.First(&existingAlert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	var input models.Alert
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizamos únicamente los campos relevantes
	existingAlert.Message = input.Message
	existingAlert.Active = input.Active
	existingAlert.Tipo = input.Tipo
	existingAlert.Severidad = input.Severidad
	existingAlert.Fecha = input.Fecha

	if err := db.DB.Save(&existingAlert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingAlert)
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
