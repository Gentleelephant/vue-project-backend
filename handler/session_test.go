package handler

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestSetSession(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "139.198.155.59:6389",
		DB:   0,
	})
	//ping := client.Ping(context.Background())
	//t.Log(ping.String())
	val := client.Get(context.Background(), "zzp").Val()
	t.Log(val)
}
