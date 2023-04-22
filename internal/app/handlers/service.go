package handlers

import (
	"context"
	snowflake_service "snowflake_service/pkg/snowflake-api"
)

type snowflakeService interface {
	NewID(ctx context.Context) (uint64, error)
}

type Implementation struct {
	snowflake_service.UnimplementedSnowflakeServiceServer
	snowflakeService snowflakeService
}

func NewSnowflakeApiService(snowflakeService snowflakeService) *Implementation {
	return &Implementation{snowflakeService: snowflakeService}
}
