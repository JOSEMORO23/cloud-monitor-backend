package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Services []Service
}

type Service struct {
	gorm.Model
	Name    string `json:"name"`
	Cloud   string `json:"cloud"`
	UserID  uint   `json:"user_id"`
	Url     string `json:"url"`
	Metrics []Metric
	Logs    []Log
	Alerts  []Alert
}

type Metric struct {
	gorm.Model
	ServiceID uint
	Name      string
	Value     float64
	Timestamp string
}

type Log struct {
	gorm.Model
	ServiceID uint
	Message   string
	Level     string
	Timestamp string
}

//	type Alert struct {
//		gorm.Model
//		ServiceID uint
//		Message   string
//		Active    bool
//	}
type Alert struct {
	gorm.Model
	ServiceID uint      `json:"service_id"`
	Message   string    `json:"message"`
	Active    bool      `json:"active"`
	Tipo      string    `json:"tipo"`      // Ej: "caida", "cpu", "ram", etc.
	Severidad string    `json:"severidad"` // Ej: "baja", "media", "alta"
	Fecha     time.Time `json:"fecha"`
}
