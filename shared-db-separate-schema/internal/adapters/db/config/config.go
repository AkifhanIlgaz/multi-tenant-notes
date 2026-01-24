package config

import "fmt"

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabaseConfig() *DatabaseConfig {
	// env dosyasindan degerler alinabilir
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "myuser",
		Password: "mypassword",
		DBName:   "mydatabase",
		SSLMode:  "disable",
	}
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}
