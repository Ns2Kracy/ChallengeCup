package config

import (
	"os"

	"ChallengeCup/utils/file"

	"gopkg.in/yaml.v3"
)

type Config struct {
	System *System `yaml:"system"`
	Mysql  *Mysql  `yaml:"mysql"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{
		System: &System{
			Name: "Challenge Cup",
			Mode: "debug",
			Host: "localhost",
			Port: "8080",
		},
		Mysql: &Mysql{
			Host:     "localhost",
			Port:     "3306",
			Database: "challenge_cup",
			User:     "username",
			Pwd:      "password",
			Driver:   "mysql",
		},
	}
	if !file.IsExist(path) {
		f, _ := file.NewFile(path)
		encoder := yaml.NewEncoder(f)
		err := encoder.Encode(config)
		if err != nil {
			return nil, err
		}
		return config, nil
	} else {
		config = LoadConfig()
	}

	return config, nil
}

func LoadConfig() *Config {
	f, _ := os.Open("config.yaml")
	decoder := yaml.NewDecoder(f)
	config := &Config{}
	err := decoder.Decode(config)
	if err != nil {
		return nil
	}
	return config
}
