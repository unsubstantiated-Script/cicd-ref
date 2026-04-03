package config

import "os"

type Config struct {
	Port    string
	Message string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	message := os.Getenv("APP_MESSAGE")
	if message == "" {
		message = "secure pipeline demo"
	}

	return Config{
		Port:    port,
		Message: message,
	}
}
