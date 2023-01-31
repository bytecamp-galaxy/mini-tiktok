package log

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func InitHLogger() {
	hlog.SetLogger(hertzzap.NewLogger(hertzzap.WithCoreEnc(GetEncoder()),
		hertzzap.WithCoreWs(GetLogWriter()),
		hertzzap.WithCoreLevel(GetAtomicLevel()),
		hertzzap.WithZapOptions(GetZapOptions()...)))
}

func InitKLogger() {
	klog.SetLogger(kitexzap.NewLogger(kitexzap.WithCoreEnc(GetEncoder()),
		kitexzap.WithCoreWs(GetLogWriter()),
		kitexzap.WithCoreLevel(GetAtomicLevel()),
		kitexzap.WithZapOptions(GetZapOptions()...)))
}

func InitDBLogger() gormlogger.Interface {
	logger := zapgorm2.New(GetLogger())
	return logger.LogMode(gormlogger.Info)
}
