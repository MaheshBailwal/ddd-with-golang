package config

type SRemoteService struct {
	DeparmentServiceIp   string `validate:"required"`
	DeparmentServicePort int    `validate:"required"`
}
