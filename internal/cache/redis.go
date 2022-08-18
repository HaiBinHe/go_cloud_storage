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
	return rdb.HSet(ctx, key, value...).Err()
}
func HGet(ctx context.Context, key, field string) (string, error) {
	return rdb.HGet(ctx, key, field).Result()
}
func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return rdb.HGetAll(ctx, key).Result()
}

//返回ID
func XAdd(ctx context.Context, x *redis.XAddArgs) (string, error) {
	return rdb.XAdd(ctx, x).Result()
}
func XGroupCreate(ctx context.Context, key, group, start string) error {
	//如果stream不存在则创建一个长度为0的stream
	return rdb.XGroupCreateMkStream(ctx, key, group, start).Err()
}
func XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) ([]redis.XStream, error) {
	return rdb.XReadGroup(ctx, a).Result()
}

//查看消费者群组 key:消息队列
func XInfo(ctx context.Context, key string) []redis.XInfoGroup {
	return rdb.XInfoGroups(ctx, key).Val()
}
func XAck(ctx context.Context, key, group string, ID ...string) error {
	return rdb.XAck(ctx, key, group, ID...).Err()
}

//修剪消息队列的长度
//当队列长度超过上限后，旧消息会被删除，只保留固定长度的新消息。
func XTrim(ctx context.Context, key string, length int64) error {
	return rdb.XTrimMaxLen(ctx, key, length).Err()
}
