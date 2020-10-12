package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

const (
	MODE_DEV = "dev"
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
	once sync.Once
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