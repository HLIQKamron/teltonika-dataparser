package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

var RDB = redis.NewClient(&redis.Options{
	Addr:     "redis-10180.c326.us-east-1-3.ec2.cloud.redislabs.com:10180",
	Password: "egsgroup", // no password set
	DB:       0,          // use default DB
})

var ctx = context.Background()
