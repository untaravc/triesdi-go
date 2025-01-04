package db_config

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Address of the Redis server
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err) // handle error appropriately
	}
	log.Println("Connected to REDIS")
	return rdb
}

func RedisSet(ctx context.Context, key string, value interface{}, ttl int) error {
	rdb := InitRedisClient()

	data, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}

	strVal := string(data)

	er := rdb.Set(ctx, key, strVal, time.Duration(ttl)*time.Second).Err()

	if er != nil {
		return er
	}
	return nil
}

// // Get retrieves the value of a given key from Redis.
func RedisGet(ctx context.Context, key string) (string, error) {
	rdb := InitRedisClient()
	return rdb.Get(ctx, key).Result()
}

// // Del deletes one or more keys from Redis.
func RedisDel(ctx context.Context, keys ...string) error {
	rdb := InitRedisClient()
	return rdb.Del(ctx, keys...).Err()
}
