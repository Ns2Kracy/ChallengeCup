package config

import (
	"os"

	"ChallengeCup/utils/file"

	"gopkg.in/yaml.v3"
)

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	SMS    SMS    `yaml:"sms"`
	Mail   Mail   `yaml:"mail"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{
		System: System{},
		Mysql:  Mysql{},
		Redis:  Redis{},
		SMS:    SMS{},
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	err = encoder.Encode(config)
	if err != nil {
		return nil, err
	}
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

func InitConfig(path string) *Config {
	// 判断配置文件是否存在
	if !file.IsExist(path) {
		// 不存在则创建
		config, err := NewConfig(path)
		if err != nil {
			panic(err)
		}
		return config
	} else {
		// 存在则加载
		return LoadConfig()
	}
}
