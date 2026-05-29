package redis

import (
	"context"
	"errors"
	"gohub/pkg/logger"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client Redis 服务
type Client struct {
	Client  *redis.Client
	Context context.Context
}

var once sync.Once

var Redis *Client

func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

func NewClient(address string, username string, password string, db int) *Client {
	rds := &Client{}
	rds.Context = context.Background()

	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

func (rds Client) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

func (rds Client) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

func (rds Client) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

func (rds Client) Has(key string) bool {
	_, err := rds.Client.Exists(rds.Context, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

func (rds Client) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

func (rds Client) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

func (rds Client) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "IncrBy", "parameters len error")
		return false
	}

	return true
}

func (rds Client) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "parameters len error")
		return false
	}

	return true
}
