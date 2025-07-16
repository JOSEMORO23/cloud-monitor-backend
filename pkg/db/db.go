package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"

	"github.com/JOSEMORO23/cloud-monitor-backend/internal/models"
)

var DB *gorm.DB

func Connect() {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	database := "cloud_monitor"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, host, port, database)

	databaseConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error conectando a MySQL:", err)
	}

	err = databaseConn.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Metric{},
		&models.Log{},
		&models.Alert{},
	)
	if err != nil {
		logrus.Fatal("❌ Error en migración:", err)
	}

	DB = databaseConn
	fmt.Println("✅ Conexión a MySQL exitosa y migración realizada")
}
