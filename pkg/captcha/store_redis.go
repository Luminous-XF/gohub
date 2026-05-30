package captcha

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.Client
	KeyPrefix   string
}

func (store *RedisStore) Set(key string, value string) {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := store.RedisClient.Set(store.KeyPrefix+key, value, ExpireTime); !ok {
		logger.ErrorString("Captcha", "Set", "Captcha store error")
	}
}

func (store *RedisStore) Get(key string, clear bool) string {
	key = store.KeyPrefix + key
	value := store.RedisClient.Get(key)
	if clear {
		store.RedisClient.Del(key)
	}
	return value
}

func (store *RedisStore) Verify(key, answer string, clear bool) bool {
	value := store.Get(key, clear)
	return value == answer
}
