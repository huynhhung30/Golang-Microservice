package config

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	SslMode  string
}

func GormOpen() (gormDB *gorm.DB, err error) {
	infoDatabase := getDiverConn()
	// functions.ShowLog("infoDatabase", infoDatabase)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", infoDatabase.Hostname, infoDatabase.Port, infoDatabase.Username, infoDatabase.Password, infoDatabase.Name, infoDatabase.SslMode)
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = gormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.Open(dsn)},
	}))
	return
}
