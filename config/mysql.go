package config

type Mysql struct {
	Driver   string `yaml:"driver"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pwd      string `yaml:"pwd"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

func (m *Mysql) GetDialect() string {
	return m.Driver
}

func (m *Mysql) GetDsn() string {
	return m.User + ":" + m.Pwd + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}
