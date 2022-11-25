package config

import (
	"log"

	"github.com/spf13/viper"
)

var vp *viper.Viper

func LoadEnvVariables() Config {

	var config Config
	vp = viper.New()

	vp.SetConfigName("env")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = vp.Unmarshal(&config)

	if err != nil {
		log.Fatalf("Error while allocating values %s", err)
	}

	return config
}
