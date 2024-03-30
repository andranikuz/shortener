package config

import (
	"flag"
	"os"
)

type AppConfig struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
}

var Config AppConfig

func Init() {
	flag.StringVar(&Config.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&Config.BaseURL, "b", "http://localhost:8080", "default url host")
	flag.StringVar(&Config.FileStoragePath, "f", "/tmp/short-url-db.json", "file storage path")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		Config.ServerAddress = envRunAddr
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		Config.BaseURL = envBaseURL
	}
	if envBaseURL := os.Getenv("FILE_STORAGE_PATH"); envBaseURL != "" {
		Config.FileStoragePath = envBaseURL
	}
}
