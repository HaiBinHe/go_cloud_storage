package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-cloud/conf"
	"log"
	"time"
)

var (
	rdb *redis.Client
)

func InitRedisConn() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisSetting.Host, conf.RedisSetting.Port),
		Password: conf.RedisSetting.Password,
		PoolSize: conf.RedisSetting.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("redis ping  err:", err)
		return err
	}
	log.Println("======redis server start======")

	return nil
}

func Set(ctx context.Context, key, value string, time time.Duration) error {
	return rdb.Set(ctx, key, value, time).Err()
}
func Get(ctx context.Context, key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// key是否存在
func IsExists(ctx context.Context, key string) bool {
	_, err := rdb.Exists(ctx, key).Result()
	if err == redis.Nil {
		return true
	}
	return false
}
func HSet(ctx context.Context, key string, value ...interface{}) error {
	return rdb.HSet(ctx, key, value).Err()
}
func HGet(ctx context.Context, key, field string) (string, error) {
	return rdb.HGet(ctx, key, field).Result()
}
func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return rdb.HGetAll(ctx, key).Result()
}
