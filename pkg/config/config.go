package config

import (
	"gohub/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

// viper 库实例
var viper *viperlib.Viper

// LoadConfigFunc 动态加载配置信息
type LoadConfigFunc func() map[string]interface{}

// LoadConfigFuncs 先加载到此数组, loadConfig 再动态生成配置信息
var LoadConfigFuncs map[string]LoadConfigFunc

func init() {
	// 1. 初始化 Viper 库
	viper = viperlib.New()

	// 2. 配置信息
	// 支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")

	// 3. 环境变量配置文件查找的路径, 相对于 main.go
	viper.AddConfigPath(".")

	// 4.设置环境变量前缀, 用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")

	// 5. 读取环境变量(支持 flags)
	viper.AutomaticEnv()

	LoadConfigFuncs = make(map[string]LoadConfigFunc)
}

// InitConfig 初始化配置信息, 完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)

	// 2. 注册配置信息
	loadConfig()
}

func loadConfig() {
	for name, fn := range LoadConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件, 如果有传参 --env=name 则加载 .env.name
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env." + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filePath
		} else {
			panic(err)
		}
	}

	// 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 .env 文件, 变更时重新加载
	viper.WatchConfig()
}

// Env 读取环境变量, 支持默认值
func Env(name string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(name, defaultValue[0])
	}
	return internalGet(name)
}

// Add 新增配置项
func Add(name string, loadConfigFunc LoadConfigFunc) {
	LoadConfigFuncs[name] = loadConfigFunc
}

// Get 获取配置项
// path: 允许使用点式获取, 如 app.name
// defaultValue: 允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

// internalGet
func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或 Env 不存在或为零值时, 尝试使用默认值, 若无默认值则返回零值
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 String 类型配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string, defaultValue ...interface{}) map[string]string {
	return viper.GetStringMapString(path)
}
