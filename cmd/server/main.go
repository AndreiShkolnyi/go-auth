package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/AndreiShkolnyi/go-auth/internal/config"
	"github.com/AndreiShkolnyi/go-auth/internal/config/env"
	"github.com/AndreiShkolnyi/go-auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

var configPath string

func init() {
	flag.StringVar(&configPath, "config=path", ".env", "path to config file")
}

type server struct {
	auth_v1.UnimplementedAuthV1Server
	pool *pgxpool.Pool
}

func (s *server) Get(_ context.Context, req *auth_v1.GetRequest) (*auth_v1.GetResponse, error) {
	log.Printf("Received: %v", req.GetId())

	return &auth_v1.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      auth_v1.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	grpcConfig, err := env.NewGRPConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load pg config %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcConfig.Address()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	auth_v1.RegisterAuthV1Server(s, &server{pool: pool})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
