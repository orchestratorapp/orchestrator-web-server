package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// The server configuration structure. This struct maps the config.yaml
// structure, so that it can easily be parsed and read when necessary.
// It is recommended to change this struct according to the changes that
// are made to the config.yaml, to keep it consistent.
type ServerConfig struct {
	Orchestrator struct {
		Server struct {
			Banner  string `yaml:"banner"`
			Version string `yaml:"version"`
			AppName string `yaml:"application-name"`
			Port    string `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"orchestrator"`
}

// Loads the configuration from the config.yaml file
// and makes it available for all the application.
func LoadConfig() ServerConfig {
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalf("\033[41m FATAL \033[0m %v", err)
	}
	defer f.Close()

	var cfg ServerConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("\033[41m FATAL \033[0m %v", err)
	}
	cfg.Orchestrator.Server.Port = ":" + cfg.Orchestrator.Server.Port
	printBanner(cfg)
	return cfg
}

func printBanner(cfg ServerConfig) {
	banner, err := os.ReadFile(cfg.Orchestrator.Server.Banner)
	if err != nil {
		log.Fatalf("\033[41m FATAL \033[0m %v", err)
	}
	fmt.Println(string(banner))
	fmt.Printf("%s v. \033[35m%s\033[0m\n",
		cfg.Orchestrator.Server.AppName, cfg.Orchestrator.Server.Version,
	)
}
