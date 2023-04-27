package handlers

import (
	"context"
	snowflake_api "github.com/deadsen1337/snowflake_api"
)

func (i *Implementation) NewID(ctx context.Context, _ *snowflake_api.IdRequest) (*snowflake_api.IdResponse, error) {
	id, err := i.snowflakeService.NewID(ctx)
	if err != nil {
		return nil, err
	}

	return &snowflake_api.IdResponse{Id: id}, nil
}
