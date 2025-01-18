package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(ctx context.Context, addr string, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{client: client}, nil
}

func (rc *RedisClient) CacheStockPrice(symbol string, price float64) error {

	return rc.client.Set(rc.ctx, "stock:"+symbol, price, time.Minute).Err()
}

func (rc *RedisClient) GetCacheStockQuote(symbol string) (float64, error) {
	data, err := rc.client.Get(rc.ctx, "stock:"+symbol).Float64()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return data, nil
}

func (rc *RedisClient) CacheCryptoPrice(symbol string, price float64) error {

	return rc.client.Set(rc.ctx, "crypto:"+symbol, price, time.Minute).Err()
}

func (rc *RedisClient) GetCacheCryptoQuote(symbol string) (float64, error) {
	data, err := rc.client.Get(rc.ctx, "crypto:"+symbol).Float64()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return data, nil
}

/* extra functions which is not necessary right now
func (rc *RedisClient) CachePortfolio(username, portfolioType string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	key := formatPortFolioKey(username, portfolioType)
	return rc.client.Set(rc.ctx, key, jsonData, 5*time.Minute).Err()
}

func (rc *RedisClient) GetCachePortfolio(username, portfolioType string, dest interface{}) error {
	key := formatPortFolioKey(username, portfolioType)
	data, err := rc.client.Get(rc.ctx, key).Result()
	if err == redis.Nil{
		return nil
	}
	if err != nil{
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}
*/
