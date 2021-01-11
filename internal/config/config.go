package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/moonliightz/telegram-webhookinfo-exporter/internal/model"
	"gopkg.in/yaml.v2"
)

// Config holds configuration details
type Config struct {
	App        model.App        `yaml:"app"`
	Telegram   model.Telegram   `yaml:"telegram"`
	Prometheus model.Prometheus `yaml:"prometheus"`
	HTTP       model.HTTP       `yaml:"http"`
}

// Load loads the config.yml file
func Load(configPath string) (*Config, error) {
	config := Config{
		App: model.App{
			Interval: 10,
		},
		Prometheus: model.Prometheus{
			Name: "pending_update_count",
		},
		HTTP: model.HTTP{
			Addr: "0.0.0.0",
			Port: 2112,
		},
	}

	filename, _ := filepath.Abs(configPath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// var config model.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	if config.Telegram.Token == "" {
		return nil, errors.New("Telegram Token is missing in config.yml")
	}

	return &config, nil
}
