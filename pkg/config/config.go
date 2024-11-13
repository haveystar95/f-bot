package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
	Monobank MonobankConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type LoggingConfig struct {
	Level string
}

type MonobankConfig struct {
	Token string
}

func LoadConfig() Config {
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./")
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("Error reading config file, %s", err)
	//}
	//
	var config Config
	//if err := viper.Unmarshal(&config); err != nil {
	//	log.Fatalf("Unable to decode into struct, %v", err)
	//}

	// Override with environment variables
	config.Server.Port = getEnvAsInt("APP_PORT", config.Server.Port)
	config.Database.Host = getEnv("POSTGRES_HOST", config.Database.Host)
	config.Database.Port = getEnvAsInt("POSTGRES_PORT", config.Database.Port)
	config.Database.User = getEnv("POSTGRES_USER", config.Database.User)
	config.Database.Password = getEnv("POSTGRES_PASSWORD", config.Database.Password)
	config.Database.DBName = getEnv("POSTGRES_DB", config.Database.DBName)
	config.Monobank.Token = getEnv("MONOBANK_TOKEN", config.Monobank.Token)

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	log.Printf("Invalid value for %s: %s, using default %d", name, valueStr, defaultValue)
	return defaultValue
}
