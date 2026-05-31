package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.Client
	KeyPrefix   string
}

func (store *RedisStore) Set(id string, code string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return store.RedisClient.Set(store.KeyPrefix+id, code, ExpireTime)
}

func (store *RedisStore) Get(id string, clear bool) string {
	key := store.KeyPrefix + id
	value := store.RedisClient.Get(key)
	if clear {
		store.RedisClient.Del(key)
	}
	return value
}

func (store *RedisStore) Verify(id string, answer string, clear bool) bool {
	value := store.Get(id, clear)
	return value == answer
}
