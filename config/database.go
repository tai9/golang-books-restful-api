package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tai9/golang_jwt/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Fail to load env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to create a connection to database")
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPostgre, err := db.DB()
	if err != nil {
		panic("Fail to close connection from database")
	}
	dbPostgre.Close()
}
