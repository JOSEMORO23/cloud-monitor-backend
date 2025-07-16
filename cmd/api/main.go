package main

import (
	"log"

	"github.com/JOSEMORO23/cloud-monitor-backend/internal/handlers"
	"github.com/JOSEMORO23/cloud-monitor-backend/internal/middlewares"

	//"github.com/JOSEMORO23/cloud-monitor-backend/pkg/cloud"
	"github.com/JOSEMORO23/cloud-monitor-backend/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Services CRUD (protegido)
	services := r.Group("/services")
	services.Use(middlewares.AuthMiddleware())
	{
		services.GET("", handlers.GetServices)
		services.POST("", handlers.CreateService)
		services.GET("/:id", handlers.GetServiceByID)
		services.PUT("/:id", handlers.UpdateService)
		services.DELETE("/:id", handlers.DeleteService)
	}

	// Metrics CRUD
	metrics := r.Group("/metrics")
	metrics.Use(middlewares.AuthMiddleware())
	{
		metrics.GET("", handlers.GetMetrics)
		metrics.POST("", handlers.CreateMetric)
		metrics.GET("/:id", handlers.GetMetricByID)
		metrics.PUT("/:id", handlers.UpdateMetric)
		metrics.DELETE("/:id", handlers.DeleteMetric)
	}
	// Users CRUD
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// ✔️ Alerts CRUD protegido
	alerts := r.Group("/alerts")
	alerts.Use(middlewares.AuthMiddleware())
	{
		alerts.GET("", handlers.GetAlerts)
		alerts.POST("", handlers.CreateAlert)
		alerts.GET("/:id", handlers.GetAlertByID)
		alerts.PUT("/:id", handlers.UpdateAlert)
		alerts.DELETE("/:id", handlers.DeleteAlert)
	}

	// ✔️ Logs CRUD protegido
	logs := r.Group("/logs")
	logs.Use(middlewares.AuthMiddleware())
	{
		logs.GET("", handlers.GetLogs)
		logs.POST("", handlers.CreateLog)
		logs.GET("/:id", handlers.GetLogByID)
		logs.PUT("/:id", handlers.UpdateLog)
		logs.DELETE("/:id", handlers.DeleteLog)
	}

	// Solo para prueba, comenta si no lo quieres al iniciar siempre
	//err := cloud.ListInstancesAWS()
	//if err != nil {
	//	log.Println("AWS Error:", err)
	//}

	//err = cloud.ListInstancesGCP("your-project-id", "us-central1-a", "path/to/credentials.json")
	//if err != nil {
	//	log.Println("GCP Error:", err)
	//}

	// Cloud Automation
	cloudRoutes := r.Group("/cloud")
	cloudRoutes.Use(middlewares.AuthMiddleware())
	{
		cloudRoutes.GET("/aws/instances", handlers.GetAWSInstances)
		cloudRoutes.GET("/gcp/instances", handlers.GetGCPInstances)
	}

	// Auth
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	log.Println("Servidor escuchando en http://localhost:8080")
	r.Run(":8080")
}
