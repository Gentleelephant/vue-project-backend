package handler

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
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
