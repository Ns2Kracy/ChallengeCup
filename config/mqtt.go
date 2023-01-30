package config

type Mqtt struct {
	Qos          int    `mapstructure:"qos" yaml:"qos"`
	Username     string `mapstructure:"username" yaml:"username"`
	Password     string `mapstructure:"password" yaml:"password"`
	ClientId     string `mapstructure:"client_id" yaml:"client_id"`
	Topic        string `mapstructure:"topic" yaml:"topic"`
	Channel      int    `mapstructure:"channel" yaml:"channel"`
	Host         string `mapstructure:"host" yaml:"host"`
	Port         int    `mapstructure:"port" yaml:"port"`
	CleanSession bool   `mapstructure:"clean_session" yaml:"clean_session"`
}
