package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
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
	return zapcore.AddSync(writer)
}

func getDBLogger() *zap.Logger {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), getLevel())
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(-1)) // TODO(vgalaxy): inconsistent caller skip
}
