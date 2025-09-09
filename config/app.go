// Package config 站点配置信息
package config

import (
	"gohub/pkg/config"
)

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env[string]("APP_NAME", "GoHub"),

			// 当前环境, 用以区分多环境, 一般为 local, stage, production, test
			"env": config.Env[string]("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env[bool]("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env[string]("APP_PORT", "8080"),

			// 加密会话, JWT 加密
			"key": config.Env[string]("APP_KEY", ""),

			// 用以生成链接
			"url": config.Env[string]("APP_URL", "http://localhost:8080"),

			// 设置时区, JWT 中会使用, 日志记录中也会使用到
			"timezone": config.Env[string]("APP_TIMEZONE", "Asia/Shanghai"),
		}
	})
}
