package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapOptions() []zap.Option {
	return []zap.Option{zap.AddCaller(), zap.AddCallerSkip(3)}
}

func getLevel() zapcore.Level {
	return zapcore.DebugLevel
}

func getAtomicLevel() zap.AtomicLevel {
	level := zap.NewAtomicLevel()
	level.SetLevel(getLevel())
	return level
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getDBLogger() *zap.Logger {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), getLevel())
	return zap.New(core) // TODO(vgalaxy): add caller
}
