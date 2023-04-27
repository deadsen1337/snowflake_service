package handlers

import (
	"context"
	snowflake_api "github.com/deadsen1337/snowflake_api"
)

type snowflakeService interface {
	NewID(ctx context.Context) (uint64, error)
}

type Implementation struct {
	snowflake_api.UnimplementedSnowflakeServiceServer
	snowflakeService snowflakeService
}

func NewSnowflakeApiService(snowflakeService snowflakeService) *Implementation {
	return &Implementation{snowflakeService: snowflakeService}
}
