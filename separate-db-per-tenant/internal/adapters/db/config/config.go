package config

import (
	"fmt"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabaseConfig(dbName string, port int) *DatabaseConfig {
	// env dosyasindan degerler alinabilir
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     port,
		User:     "myuser",
		Password: "mypassword",
		DBName:   dbName,
		SSLMode:  "disable",
	}
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}
