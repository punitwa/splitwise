package config

import (
	"encoding/json"
	"os"
)

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	SSLMode  string `json:"sslmode"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
}

var AppConfig Config

func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}

	return nil
}
