// Package config 负责配置信息
package config

import (
	"os"
	"reflect"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

// CfgFunc 动态加载配置信X息
type CfgFunc func() map[string]interface{}

// CfgValue 是可以从配置中获取的类型
type CfgValue interface {
	~string | ~int | ~int64 | ~float64 | ~bool | ~map[string]string
}

// CfgFuncs 先加载到此数组, loadConfig 再动态生成配置信息
var CfgFuncs map[string]CfgFunc

func init() {
	viper = viperlib.New()

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()

	CfgFuncs = make(map[string]CfgFunc)
}

// InitConfig 初始化配置信息, 完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)

	// 2. 注册配置信息
	loadConfig()
}

func loadConfig() {
	for name, fn := range CfgFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件, 如果有传参 --env=name, 则加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env" + envSuffix
		if _, err := os.Stat(filePath); err == nil {
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

	// 监控 .env 文件, 配置文件发生变更时重新加载
	viper.WatchConfig()
}

// Env 读取环境变量, 支持默认值
func Env[T CfgValue](envName string, defaultValue T) T {
	return GetWithDefault(envName, defaultValue)
}

// Add 新增配置项
func Add(name string, configFn CfgFunc) {
	CfgFuncs[name] = configFn
}

func Get[T CfgValue](path string) (T, bool) {
	// 定义一个泛型默认值, 如果没有取到值则返回零值和 false
	var fallback T

	value := viper.Get(path)
	if value == nil {
		return fallback, false
	}

	// 获取目标类型
	targetType := reflect.TypeOf(fallback)
	valueType := reflect.ValueOf(value)

	if valueType.Kind() == targetType.Kind() {
		return value.(T), true
	}

	// 尝试类型转换
	switch any(fallback).(type) {
	case string:
		strValue := cast.ToString(value)
		return any(strValue).(T), true
	case int:
		intValue := cast.ToInt(value)
		return any(intValue).(T), true
	case int64:
		intValue := cast.ToInt64(value)
		return any(intValue).(T), true
	case float64:
		floatValue := cast.ToFloat64(value)
		return any(floatValue).(T), true
	case map[string]string:
		stringMapValue := cast.ToStringMapString(value)
		return any(stringMapValue).(T), true
	default:
		return fallback, false
	}
}

func GetWithDefault[T CfgValue](path string, defaultValue T) T {
	value, ok := Get[T](path)
	if !ok {
		return defaultValue
	}
	return value
}
