package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	client *redis.Client
}

func NewRedisPubSub(addr string, password string, db int) *RedisPubSub {
	return &RedisPubSub{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// Publish publishes a message to the specified channel using the Redis client.
// It takes a context, channel name, and location data as inputs.
// It returns an error if the publish operation fails.
func (s *RedisPubSub) Publish(ctx context.Context, channel string, location []byte) error {
	err := s.client.Publish(ctx, channel, location).Err()
	fmt.Printf("channel %s location : %+v\n", channel, location)
	return err
}

// Subscribe subscribes to the specified channel using the Redis client.
// It takes a context and channel name as inputs.
// It returns a PubSub object that can be used to receive messages from the subscribed channel.
func (s *RedisPubSub) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	pubsub := s.client.Subscribe(ctx, channel)
	return pubsub
}
