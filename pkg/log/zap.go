package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetZapOptions() []zap.Option {
	return []zap.Option{zap.AddCaller(), zap.AddCallerSkip(3)}
}

func GetLevel() zapcore.Level {
	return zapcore.DebugLevel
}

func GetAtomicLevel() zap.AtomicLevel {
	level := zap.NewAtomicLevel()
	level.SetLevel(GetLevel())
	return level
}

func GetLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func GetLogger() *zap.Logger {
	core := zapcore.NewCore(GetEncoder(), GetLogWriter(), GetLevel())
	return zap.New(core) // TODO(vgalaxy): add caller
}
