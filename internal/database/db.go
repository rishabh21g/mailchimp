package database

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rishabh21g/mailchimp/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB() (*gorm.DB, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		slog.Warn("Error while fetching DB credentials")
		return nil, err

	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		slog.Error("Database Url not set in environments", "error", err)
		return nil, err

	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Error while connecting with Database", "error", err)
		return nil, err

	}
	err = db.AutoMigrate(&types.User{})
	if err != nil {
		slog.Error("Error while migrating  model", "error", err)
		return nil, err
	}
	DB = db
	slog.Info("Database connected successfully")
	return db, nil

}
