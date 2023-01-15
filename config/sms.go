package config

type SMS struct {
	SignName     string `yaml:"sign_name"`
	AccessKeyID  string `yaml:"access_key_id"`
	AppSecret    string `yaml:"app_secret"`
	TemplateCode string `yaml:"template_code"`
	RegionId     string `yaml:"region_id"`
}
