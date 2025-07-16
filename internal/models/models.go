package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Password string
	Services []Service
}

type Service struct {
	gorm.Model
	Name    string
	Cloud   string
	UserID  uint
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

type Alert struct {
	gorm.Model
	ServiceID uint
	Message   string
	Active    bool
}
