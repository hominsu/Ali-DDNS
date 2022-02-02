package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewProductionLogger new a zap logger
func NewProductionLogger() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "/logs/ads.log",
		MaxSize:   64,
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		w,
		zap.NewAtomicLevel(),
	)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
