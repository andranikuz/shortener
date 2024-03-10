package config

import "flag"

type AppConfig struct {
	RunAddr     string
	DefaultHost string
}

var Config AppConfig

func Init() {
	flag.StringVar(&Config.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&Config.DefaultHost, "b", "http://localhost:8080", "default url host")
	flag.Parse()
}
