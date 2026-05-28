package logger

import (
	"encoding/json"
	"fmt"
	"gohub/pkg/app"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger(fileName string, maxSize, maxBackups, maxAge int, compress bool, logType string, level string) {
	// 获取日志写入介质
	writeSyncer := getLogWriter(fileName, maxSize, maxBackups, maxAge, compress, logType)

	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("init logger error", zap.String("error", err.Error()))
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化 Logger
	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	zap.ReplaceGlobals(Logger)
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if app.IsLocal() {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 本地使用内置的 Console 解码器
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.DateTime))
}

// getLogWriter 日志记录介质 os.Stdout & file
func getLogWriter(fileName string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		logName := time.Now().Format(time.DateOnly + ".log")
		fileName = strings.ReplaceAll(fileName, "logs.log", logName)
	}

	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	if app.IsLocal() {
		// 本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(lumberJackLogger),
		)
	}

	// 生产环境只记录文件
	return zapcore.AddSync(lumberJackLogger)
}

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息
func Dump(data interface{}, msg ...string) {
	valueString := jsonString(data)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

// LogIf 当 err != nil 时, 记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时, 记录 warning 等级的日志
func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志，详尽的程序日志
// logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

// Info 告知类日志
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// Warn 警告类日志
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error 错误时记录, 不应该中断程序, 查看日志时重点关注
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志
// logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志, 使用 json.Marshal 进行编码
// logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, data interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(data)))
}

func InfoJSON(moduleName, name string, data interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(data)))
}

func WarnJSON(moduleName, name string, data interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(data)))
}

func ErrorJSON(moduleName, name string, data interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(data)))
}

func FatalJSON(moduleName, name string, data interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(data)))
}

func jsonString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
