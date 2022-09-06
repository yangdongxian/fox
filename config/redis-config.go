package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	ctx = context.Background()
	Rdb *redis.Client
)

type IRedis interface {
	Set(key string, value string, expiration time.Duration) (result string, err error)
	Get(key string) (result string, err error)
	GetRdb() *redis.Client
}
type RedisConnInfo struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}
type RedisClient struct {
	Client *redis.Client
}

func SetupRedisConnection(r *RedisConnInfo) *redis.Client {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env files")
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(redisDb)
	if err != nil {
		log.Println("string to int failed", err.Error())
	}

	if len(r.Addr) < 1 || (len(r.Password) < 1) || (r.DB < 0) {
		r.Addr = redisAddr
		r.Password = redisPassword
		r.DB = db
		r.PoolSize = 20
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})
	pingErr := rdb.Ping(ctx).Err()
	if pingErr != nil {
		fmt.Println("redis.client connection is failed,err:", pingErr.Error())
	}
	err = rdb.Set(ctx, "test", "test", 0).Err()
	if err != nil {
		fmt.Println("redis.client connection is failed,err:", err.Error())
	}

	return rdb
}
func CloseRedisClient(rdb *redis.Client) {
	rdb.Close()
}

func NewRedisClient() IRedis {
	var rc = RedisConnInfo{Addr: "", Password: "", DB: -1, PoolSize: 20}
	Rdb = SetupRedisConnection(&rc)
	fmt.Println("... NewRedisClient ...")
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("redis.client connection is failed error:%v", err.Error())
	}
	fmt.Println("NewRidsClient poolStatus:", Rdb.PoolStats())

	return &RedisClient{
		Client: Rdb,
	}
}
func (r *RedisClient) Set(key string, value string, expiration time.Duration) (result string, err error) {
	return r.Client.Set(ctx, key, value, expiration).Result()
}

func (r *RedisClient) Get(key string) (result string, err error) {
	return r.Client.Get(ctx, key).Result()
}
func (r *RedisClient) GetRdb() *redis.Client {
	return r.Client
}
