package config

import (
	"flag"
	"os"
)

type AppConfig struct {
	ServerAddress string
	BaseURL       string
	LogLevel      string
}

var Config AppConfig

func Init() {
	flag.StringVar(&Config.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&Config.BaseURL, "b", "http://localhost:8080", "default url host")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		Config.ServerAddress = envRunAddr
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		Config.BaseURL = envBaseURL
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		Config.LogLevel = envLogLevel
	}
}
