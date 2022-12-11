package config

type Mysql struct {
	Driver   string `yaml:"driver"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pwd      string `yaml:"pwd"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}
