package config

import "gohub/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			"connection": config.Env("DB_CONNECTION", "mysql"),
			"mysql": map[string]interface{}{
				"host": config.Env("DB_HOST", "127.0.0.1"),
			},
		}
	})
}
