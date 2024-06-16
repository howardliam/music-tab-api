package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port int32
}

type PostgresConfig struct {
	Address      string
	Port         int32
	DatabaseName string
	Username     string
	Password     string
}

func defaultConfig() Config {
	server := ServerConfig{
		Port: 3000,
	}

	postgres := PostgresConfig{
		Address:      "localhost",
		Port:         5432,
		DatabaseName: "music-tabs",
		Username:     "postgres",
		Password:     "postgres",
	}

	return Config{
		Server:   server,
		Postgres: postgres,
	}
}

func LoadConfig() Config {
	if _, err := os.Stat("./config.yaml"); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("./config.yaml")
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		defaultConfig := defaultConfig()
		configString, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			log.Fatalf("Failed to parse struct to YAML: %v", err)
		}

		_, err = file.Write(configString)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
		file.Sync()
		return defaultConfig
	}

	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse YAML to struct: %v", err)
	}

	return config
}

func ConfigIsDefault(config *Config) bool {
	def := defaultConfig()

	return def == *config
}
