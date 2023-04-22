package snowflake_service

import (
	"context"
	"github.com/sony/sonyflake"
)

type SnowflakeService struct {
	Sf *sonyflake.Sonyflake
}

func (s *SnowflakeService) NewID(ctx context.Context) (uint64, error) {
	return s.Sf.NextID()
}

func NewService(sf *sonyflake.Sonyflake) *SnowflakeService {
	return &SnowflakeService{Sf: sf}
}
