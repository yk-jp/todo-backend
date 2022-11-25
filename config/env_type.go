package config

type DbConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

type Config struct {
	Db DbConfig `mapstructure:"db"`
}
