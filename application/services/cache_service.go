package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

type CacheService struct {
	cache *bigcache.BigCache
}

type CacheServiceInterface interface {
	Get(string, interface{}) error
	Set(string, interface{}) error
	Delete(string) error
}

func NewCacheService() CacheServiceInterface {

	c, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	return &CacheService{
		c,
	}
}

func (c CacheService) Get(key string, dest interface{}) error {
	r, err := c.cache.Get(key)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(r, dest); err != nil {
		return err
	}

	return nil
}

func (c CacheService) Set(key string, data interface{}) error {

	bt, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return c.cache.Set(key, bt)
}

func (c CacheService) Delete(key string) error {
	return c.cache.Delete(key)
}
