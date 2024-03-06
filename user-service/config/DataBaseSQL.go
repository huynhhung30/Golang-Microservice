package config

import (
	"fmt"
	"os"
)

type infoDatabaseSQL struct {
	Hostname   string
	Name       string
	Username   string
	Password   string
	Port       string
	DriverConn string
}

func getDiverConn() (infoDB infoDatabaseSQL) {
	infoDB.Hostname = os.Getenv("DB_HOST")
	infoDB.Name = os.Getenv("DB_NAME")
	infoDB.Username = os.Getenv("DB_USER")
	infoDB.Password = os.Getenv("DB_PASS")
	infoDB.Port = os.Getenv("DB_PORT")
	infoDB.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", infoDB.Username, infoDB.Password, infoDB.Hostname, infoDB.Port, infoDB.Name)
	return infoDB
}
