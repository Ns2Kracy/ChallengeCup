package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	System *System `yaml:"system"`
	Mysql  *Mysql  `yaml:"mysql"`
}

func InitConfig() *Config {
	config := &Config{
		System: &System{
			Name: "Challenge Cup",
			Mode: "debug",
			Host: "localhost",
			Port: "8080",
		},
		Mysql: &Mysql{},
	}
	// 查看配置文件是否存在
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		// 不存在则创建
		f, _ := os.Create("config.yaml")
		// 并且写入默认配置
		encoder := yaml.NewEncoder(f)
		err := encoder.Encode(config)
		if err != nil {
			return config
		}
		defer f.Close()
	}
	// 存在则读取
	f, _ := os.Open("config.yaml")
	decoder := yaml.NewDecoder(f)
	err := decoder.Decode(config)
	if err != nil {
		return config
	}
	return config
}
