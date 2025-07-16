package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/cloud"
)

// Handler para AWS
func GetAWSInstances(c *gin.Context) {
	err := cloud.ListInstancesAWS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Instancias AWS listadas en servidor (ver logs consola)"})
}

// Handler para GCP
func GetGCPInstances(c *gin.Context) {
	// Cambia estos valores por tus reales si quieres probar en vivo
	projectID := "YOUR_GCP_PROJECT_ID"
	zone := "YOUR_GCP_ZONE"
	credentialsPath := "path/to/your/credentials.json"

	err := cloud.ListInstancesGCP(projectID, zone, credentialsPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Instancias GCP listadas en servidor (ver logs consola)"})
}
