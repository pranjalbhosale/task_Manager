package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	DBConfig  DBConfig `yaml:"db"`
	Port      string   `yaml:"port"`
	JWTSecret string   `yaml:"jwt_secret"`
}

func LoadConfig(filePath string) Config {
	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Failed to read config file", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatal("failed to parse config file:", err)
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg

}
