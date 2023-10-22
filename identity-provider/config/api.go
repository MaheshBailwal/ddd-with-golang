package config

type ApiConfig struct {
	Host string `validate:"required"`
	Port string `validate:"required"`
}
