package snowflake

import "github.com/sony/sonyflake"

func New() *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(sonyflake.Settings{})
}
