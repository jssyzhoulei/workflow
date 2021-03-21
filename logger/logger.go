package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
//func NewLogger(filePath string, level zapcore.Level, maxSize int,
//	maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
//	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
//	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
//}

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)
}

//var Logger *zap.Logger
//var GatewayLogger *zap.Logger

//func init() {
//
//	Logger = NewLogger("./logs/org.log", zapcore.InfoLevel, 128, 30, 7, true, "Main")
//	GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
//}

//
//func Logger() *zap.Logger {
//	once.Do(func() {
//		logger, _ = NewLogger(MODE_DEV)
//	})
//	return logger
//}

const (
	MODE_DEV  = "dev"
	MODE_PROD = "prod"
)

func NewLogger(mode string) (*zap.Logger, error) {
	var (
		logConfig zap.Config
	)
	switch mode {
	case "", MODE_DEV:
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeTime = timeEncoder
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logConfig.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	case MODE_PROD:
		logConfig = zap.NewProductionConfig()
	default:
		panic("unknown run mode it mast dev or prod")
	}
	return logConfig.Build()
}

var (
	logger *zap.Logger
	once   sync.Once
)

func Logger() *zap.Logger {
	once.Do(func() {
		logger, _ = NewLogger(MODE_DEV)
	})
	return logger
}

func timeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	s := fmt.Sprintf("\x1b[0;33m%s\x1b[0m", time.Format("[2006-01-02 15:04:05]"))
	encoder.AppendString(s)
}
