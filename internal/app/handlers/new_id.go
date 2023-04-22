package handlers

import (
	"context"
	snowflake_service "snowflake_service/pkg/snowflake-api"
)

func (i *Implementation) NewID(ctx context.Context, _ *snowflake_service.IdRequest) (*snowflake_service.IdResponse, error) {
	id, err := i.snowflakeService.NewID(ctx)
	if err != nil {
		return nil, err
	}
	
	return &snowflake_service.IdResponse{Id: id}, nil
}
