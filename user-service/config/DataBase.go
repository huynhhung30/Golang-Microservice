package config

import (
	"user_service/utils/functions"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func GormOpen() (gormDB *gorm.DB, err error) {
	infoDatabase := getDiverConn()
	functions.ShowLog("infoDatabase", infoDatabase)

	gormDB, err = gorm.Open(mysql.Open(infoDatabase.DriverConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return
	}

	err = gormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(infoDatabase.DriverConn)},
	}))
	functions.ShowLog("err", err)

	return
}
