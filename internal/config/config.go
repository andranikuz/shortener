package config

import "flag"

type AppConfig struct {
	RunAddr     string
	DefaultHost string
}

var Config AppConfig

func Init() {
	runAddr := *flag.String("a", "localhost:8080", "address and port to run server")
	defaultHost := *flag.String("b", "http://localhost:8080", "default url host")
	flag.Parse()
	Config = AppConfig{RunAddr: runAddr, DefaultHost: defaultHost}
}
