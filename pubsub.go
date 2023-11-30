package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type PubSub interface {
	Publish(ctx context.Context, channel string, location []byte) error
	Subscribe(ctx context.Context, channel string) *redis.PubSub
}
