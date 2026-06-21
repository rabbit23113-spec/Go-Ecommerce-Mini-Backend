package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Port     string   `yaml:"port"`
	DBConfig DBConfig `yaml:"db"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Url      string `yaml:"url"`
}

func (c *Config) ReadConfig() (*Config, error) {
	var config *Config
	os.Chdir("../internal/config")
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		return &Config{}, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return &Config{}, err
	}
	return config, nil
}
