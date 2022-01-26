package interceptor

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

func CodeLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.Internal:
		return zapcore.ErrorLevel
	case codes.OutOfRange:
		return zapcore.WarnLevel
	case codes.DataLoss:
		return zapcore.WarnLevel
	default:
		return zapcore.InfoLevel
	}
}

func ZapInterceptor(logger *zap.Logger) *zap.Logger {
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return logger
}
