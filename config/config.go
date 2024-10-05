package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	configPath string
	DBPath     string `yaml:"dbPath"`
	ListName   string `yaml:"listName"`
}

func NewConfig(configPath string) *Config {
	return &Config{
		configPath: configPath,
	}
}

func (c *Config) Load() error {
	rawData, err := os.ReadFile(c.configPath)
	if err != nil {
		return err
	}
	var fileConfig Config
	err = yaml.Unmarshal(rawData, &fileConfig)
	if err != nil {
		return err
	}
	c.DBPath = fileConfig.DBPath
	c.ListName = fileConfig.ListName
	return nil
}
