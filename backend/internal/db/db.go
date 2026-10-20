package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"storycraft-ai/backend/config"
	"storycraft-ai/backend/internal/model"
)

var DB *gorm.DB

func Connect() error {
	cfg := config.Load()
	dsn := cfg.DSN()
	log.Printf("connecting mysql: %s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)

	gormLogger := logger.Default.LogMode(logger.Info)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger, SkipDefaultTransaction: true})
	if err != nil {
		return err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := conn.AutoMigrate(
		&model.Project{},
		&model.Chapter{},
		&model.Scene{},
		&model.Shot{},
		&model.AITask{},
		&model.AIResult{},
	); err != nil {
		return err
	}

	DB = conn
	return nil
}
