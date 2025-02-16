package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MetaService MetaService `yaml:"meta_service"`
}

type MetaService struct {
	DB Database `yaml:"db"`
}

type Database struct {
	Address    string `yaml:"address"`
	Migrations string `yaml:"migrations"`
}

// GetConfig - get config from config file
func GetConfig() (*Config, error) {
	cfg := &Config{}
	f, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	if err = yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, err
}
