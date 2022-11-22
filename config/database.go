package config

import (
	"fmt"
	"strings"
)

type DatabaseConfig struct {
	DbName          string `env:"DB_NAME"`
	Host            string `env:"DB_HOST" env-default:"localhost:5432"`
	Username        string `env:"DB_USERNAME"`
	Password        string `env:"DB_PASSWORD"`
	MaxIddleConn    int    `env:"DB_MAX_IDDLE_CONN"`
	MaxOpenConn     int    `env:"DB_MAX_OPEN_CONN"`
	ConnMaxLifetime int    `env:"DB_CONN_MAX_LIFETIME"`
}

func (config DatabaseConfig) GetDatabaseConnectionString() string {
	var (
		host = "localhost"
		port = "5432"
	)
	if splittedHost := strings.Split(config.Host, ":"); len(splittedHost) == 2 {
		host, port = splittedHost[0], splittedHost[1]
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, config.Username, config.Password, config.DbName, port)
}
