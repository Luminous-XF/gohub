// Package captcha 处理图片验证码逻辑
package captcha

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}

		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}

		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.max_skew"),
			config.GetInt("captcha.dot_count"),
		)

		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

func (captcha *Captcha) GenerateCaptcha() (id, b64s string, err error) {
	return captcha.Base64Captcha.Generate()
}

func (captcha *Captcha) VerifyCaptcha(id, answer string) (match bool) {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}

	return captcha.Base64Captcha.Verify(id, answer, false)
}
