package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type SConfig struct {
	Service       *SService       `validate:"required"`
	Logger        *SLogger        `validate:"required"`
	Tracer        *STracer        `validate:"required"`
	MsSql         *SMSSql         `validate:"required"`
	RemoteService *SRemoteService `validate:"required"`
}

func NewConfig() *SConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error in read config ", err)
	}

	config := new(SConfig)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("error in unmarshal config ", err)
	}

	if err := validator.New().Struct(config); err != nil {
		log.Fatalln("error in validate config ", err)
	}

	return config
}
