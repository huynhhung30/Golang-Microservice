package config

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type infoDatabaseSQL struct {
	Hostname   string
	Name       string
	Username   string
	Password   string
	Port       string
	DriverConn string
	SslMode    string
}

func getDiverConn() (infoDB infoDatabaseSQL) {
	infoDB.Hostname = os.Getenv("DB_HOST")
	infoDB.Name = os.Getenv("DB_NAME")
	infoDB.Username = os.Getenv("DB_USER")
	infoDB.Password = os.Getenv("DB_PASS")
	infoDB.Port = os.Getenv("DB_PORT")
	infoDB.SslMode = os.Getenv("DB_SSLMODE")
	// infoDB.DriverConn = fmt.Sprintf("%s://%s@%s:%s/%s", infoDB.Username, infoDB.Password, infoDB.Hostname, infoDB.Port, infoDB.Name, infoDB.SslMode)
	return infoDB
}
