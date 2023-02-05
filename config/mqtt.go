package config

type Mqtt struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	ClientId string `mapstructure:"client_id" yaml:"client_id"`
	Broker   string `mapstructure:"broker" yaml:"broker"`
}
