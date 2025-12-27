package env

import (
	"net"
	"os"

	"github.com/AndreiShkolnyi/go-auth/internal/config"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

func NewGRPConfig() (config.GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		port = "50051"
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
