package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"snowflake_service/internal/app/handlers"
	"snowflake_service/internal/snowflake"
	snowflake_service "snowflake_service/internal/snowflake-service"
	snowflake_api_service "snowflake_service/pkg/snowflake-api"
)

func main() {
	sf := snowflake.New()
	snowflakeService := snowflake_service.NewService(sf)
	service := handlers.NewSnowflakeApiService(snowflakeService)
	listener, err := net.Listen("tcp", ":7002")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	snowflake_api_service.RegisterSnowflakeServiceServer(grpcServer, service)

	log.Fatal(grpcServer.Serve(listener))
}
