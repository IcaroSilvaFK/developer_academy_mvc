package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	cache *redis.Client
}

var ctx = context.Background()

type CacheServiceInterface interface {
	Get(string, interface{}) error
	Set(string, interface{}) error
	Delete(string) error
}

func NewCacheService() CacheServiceInterface {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(utils.REDIS_SERVER),
		Password: os.Getenv(utils.REDIS_PASS),
		DB:       0,
	})

	return &CacheService{
		rdb,
	}
}

func (c CacheService) Get(key string, dest interface{}) error {
	r, err := c.cache.Get(ctx, key).Result()

	if err == redis.Nil {
		return fmt.Errorf("THE KEY %s NOT EXISTS IN REDIS", key)
	}

	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(r), dest); err != nil {
		return err
	}

	return nil
}

func (c CacheService) Set(key string, data interface{}) error {

	bt, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return c.cache.Set(ctx, key, bt, 5*time.Minute).Err()
}

func (c CacheService) Delete(key string) error {
	return c.cache.Del(ctx, key).Err()
}
