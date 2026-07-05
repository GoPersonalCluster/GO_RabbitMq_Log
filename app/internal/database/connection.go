package database

import (
	"fmt"
	"log"

	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/config"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config := config.NewEnvironmentConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	err = DB.AutoMigrate(&models.ErrorLog{}, &models.PipelineLog{})
	if err != nil {
		log.Fatal(err)
	}
}
