package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Loads the configuration from the config.yaml file
// and makes it available for all the application.
func LoadConfig() (*ServerConfig, *ProfileConfig) {
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
	profileCfg, err := loadProfile(cfg.Orchestrator.ActiveProfile)
	if err != nil {
		log.Fatalf("\033[41m FATAL \033[0m %v", err)
	}
	return &cfg, profileCfg
}

func loadProfile(profile string) (*ProfileConfig, error) {
	if len(profile) > 0 {
		f, err := os.Open("./config-" + profile + ".yaml")
		if err != nil {
			log.Fatalf("\033[41m FATAL \033[0m %v", err)
		}
		defer f.Close()
		var cfg ProfileConfig
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&cfg)
		if err != nil {
			log.Fatalf("\033[41m FATAL \033[0m %v", err)
		}

		log.Printf("Active profile: \033[32m%v\033[0m", profile)
		return &cfg, nil
	}
	return nil, errors.New("no profiles found")
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
