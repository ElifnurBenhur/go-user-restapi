package config

import (
	"github.com/ElifnurBenhur/go-gin-gorm-try/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// server db connect: postgres://elif:hurKUS.18072022@93.115.79.25:5432/hurkus
//local connect :     postgres://postgres:1234@localhost:5432/postgres
func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:1234@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
