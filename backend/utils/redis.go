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

// setting up the redis server
func NewRedisClient(ctx context.Context, addr string, password string, db int) (*RedisClient, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{
		client: client,
		ctx:    ctx,
	}, nil
}

// caching the price of a stock into redis client
func (rc *RedisClient) CacheStockPrice(symbol string, price float64) error {
	if rc.client == nil {
		return redis.ErrClosed
	}

	return rc.client.Set(rc.ctx, "stock:"+symbol, price, time.Second).Err()
}

// retrieveing data from the cache in redis
func (rc *RedisClient) GetCacheStockQuote(symbol string) (float64, error) {
	if rc.client == nil {
		return 0, redis.ErrClosed
	}

	data, err := rc.client.Get(rc.ctx, "stock:"+symbol).Float64()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return data, nil
}

// caching the price of a crypto into redis client
func (rc *RedisClient) CacheCryptoPrice(symbol string, price float64) error {
	if rc.client == nil {
		return redis.ErrClosed
	}

	return rc.client.Set(rc.ctx, "crypto:"+symbol, price, time.Second).Err()
}

// retrieveing data from the cache in redis
func (rc *RedisClient) GetCacheCryptoQuote(symbol string) (float64, error) {
	if rc.client == nil {
		return 0, redis.ErrClosed
	}

	data, err := rc.client.Get(rc.ctx, "crypto:"+symbol).Float64()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return data, nil
}

func (rc *RedisClient) Close() error {
	if rc.client != nil {
		return rc.client.Close()
	}
	return nil
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
