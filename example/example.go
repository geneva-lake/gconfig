package main

import (
	"github.com/geneva-lake/gconfig"
)

// Some configuration for some service
type Config struct {
	Name               string `yaml:"name"`
	Port               string `yaml:"port"`
	DBConnectionString string `yaml:"db_connection_string"`
}

var serviceConfig *Config
var err error

func main() {
	// parsing yaml file into Config struct
	serviceConfig, err = gconfig.NewConfig[Config]().FromFile("config.yaml").Yaml()
}
