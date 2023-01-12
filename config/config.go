package config

import (
	"os"

	"ChallengeCup/utils/file"

	"gopkg.in/yaml.v3"
)

type Config struct {
	System *System `yaml:"system"`
	Mysql  *Mysql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
	SMS    *SMS    `yaml:"sms"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{
		System: &System{},
		Mysql:  &Mysql{},
		Redis:  &Redis{},
		SMS:    &SMS{},
	}
	if !file.IsExist(path) {
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		encoder := yaml.NewEncoder(file)
		err = encoder.Encode(config)
		if err != nil {
			return nil, err
		}
	}
	config = LoadConfig()
	return config, nil
}

func LoadConfig() *Config {
	config := &Config{}
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil
	}
	return config
}
