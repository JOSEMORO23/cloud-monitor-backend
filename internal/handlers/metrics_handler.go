package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
)

// Obtener todas las métricas
func GetMetrics(c *gin.Context) {
	var metrics []models.Metric
	if err := db.DB.Find(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metrics)
}

// Crear nueva métrica
func CreateMetric(c *gin.Context) {
	var metric models.Metric
	if err := c.ShouldBindJSON(&metric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&metric).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, metric)
}

// Obtener métrica por ID
func GetMetricByID(c *gin.Context) {
	id := c.Param("id")
	var metric models.Metric
	if err := db.DB.First(&metric, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metric not found"})
		return
	}
	c.JSON(http.StatusOK, metric)
}

// Actualizar métrica
func UpdateMetric(c *gin.Context) {
	id := c.Param("id")
	var metric models.Metric
	if err := db.DB.First(&metric, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metric not found"})
		return
	}

	if err := c.ShouldBindJSON(&metric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&metric)
	c.JSON(http.StatusOK, metric)
}

// Eliminar métrica
func DeleteMetric(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DB.Delete(&models.Metric{}, idNum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Metric deleted"})
}