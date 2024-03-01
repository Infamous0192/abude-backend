package app

import (
	"abude-backend/internal/config"
	"abude-backend/internal/pkg/auth"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInstance struct {
	*gorm.DB
	Config *config.DatabaseConfig
}

func (d *DatabaseInstance) Setup(config *config.DatabaseConfig) {
	d.Config = config

	var err error
	connectionString := ""
	if d.DB != nil {
		return
	}

	switch d.Config.Driver {
	case "postgres":
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", d.Config.Host, d.Config.Port, d.Config.Username, d.Config.Name, d.Config.Password)
		d.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(d.debugMode()),
		})
	default:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", d.Config.Username, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.Name)
		d.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(d.debugMode()),
		})
	}
	if err != nil {
		panic(err)
	}

	d.DB.Callback().Create().Before("gorm:create").Register("editor:before_create", auth.AssignEditor)
	d.DB.Callback().Update().Before("gorm:update").Register("editor:before_update", auth.AssignEditor)
}

func (d *DatabaseInstance) debugMode() logger.LogLevel {
	switch d.Config.Debug {
	case "info":
		return logger.Info
	case "error":
		return logger.Error
	default:
		return logger.Silent
	}
}

// Custom Query

func (d *DatabaseInstance) Filter(queries interface{}) *DatabaseInstance {
	return d
}
