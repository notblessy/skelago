package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/notblessy/skelago/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB :nodoc:
func InitDB() *gorm.DB {
	logLevel := logger.Info

	if config.ENV() == "PRODUCTION" {
		logLevel = logger.Error
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, config.DBUser(), config.DBPassword(), config.DBHost(), config.DBPort(), config.DBName())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to connect: %s", err))
	}

	return db
}

// CloseDB :nodoc:
func CloseDB(db *gorm.DB) {
	postgres, err := db.DB()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to disconnect: %s", err))
	}

	err = postgres.Close()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to close: %s", err))
	}
}
