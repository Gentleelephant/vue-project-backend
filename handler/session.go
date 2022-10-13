package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func SetSession(ctx context.Context, rdb *redis.Client, key string, value any) error {
	out, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, key, out, time.Second*360).Err()
	if err != nil {
		return err
	}
	return err
}
