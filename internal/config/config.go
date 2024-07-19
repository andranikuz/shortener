package config

import (
	"encoding/json"
	"flag"
	"os"
)

// AppConfig структура отвечает за конфигурацию приложения.
type AppConfig struct {
	isInit          bool
	ServerAddress   string `json:"server_address"`
	BaseURL         string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
}

// Config публичная переменная, которой пользуются приложения дл получения параметров конфигурации приложения.
var Config AppConfig

// Init функция инициализирует конфигурацию приложения.
func Init() error {
	if !Config.isInit {
		var configPath string
		flag.StringVar(&configPath, "c", "", "config path")
		flag.StringVar(&configPath, "config", "", "config path")
		if config := os.Getenv("CONFIG"); config != "" {
			configPath = config
		}
		if configPath != "" {
			if err := loadConfigFromFile(configPath); err != nil {
				return err
			}
		}

		flag.StringVar(&Config.ServerAddress, "a", "localhost:8080", "address and port to run server")
		flag.StringVar(&Config.BaseURL, "b", "http://localhost:8080", "default url host")
		flag.StringVar(&Config.FileStoragePath, "f", "/tmp/short-url-db.json", "file storage path")
		flag.StringVar(&Config.DatabaseDSN, "d", "", "database dsn")
		flag.BoolVar(&Config.EnableHTTPS, "s", false, "enable https")
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
		if databaseDSN := os.Getenv("DATABASE_DSN"); databaseDSN != "" {
			Config.DatabaseDSN = databaseDSN
		}
		if enableHTTPS := os.Getenv("ENABLE_HTTPS"); enableHTTPS == "true" {
			Config.EnableHTTPS = true
		}
		Config.isInit = true
	}

	return nil
}

func loadConfigFromFile(filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var config AppConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	Config = config

	return nil
}
