package oauth2gorm

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBType int8

// Config gorm configuration
type Config struct {
	TableName   string
	MaxLifetime time.Duration
	Dialector   gorm.Dialector
}

var defaultConfig = &gorm.Config{
	Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // slow SQL
			LogLevel:      logger.Info, // log level
			Colorful:      true,        // color
		},
	),
}

func NewConfig(dsn string, tableName string) *Config {
	var d gorm.Dialector

	d = postgres.New(postgres.Config{
		DSN: dsn,
	})

	config := &Config{
		TableName:   tableName,
		MaxLifetime: time.Hour * 2,
		Dialector:   d,
	}

	return config
}
