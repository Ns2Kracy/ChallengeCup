package config

type System struct {
	Name string `yaml:"name"`
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	LogPath string `yaml:"log_path"`
}
