package log

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	hertzzap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func InitHLogger() {
	hlog.SetLogger(hertzzap.NewLogger(hertzzap.WithCoreEnc(getEncoder()),
		hertzzap.WithCoreWs(getLogWriter()),
		hertzzap.WithCoreLevel(getAtomicLevel()),
		hertzzap.WithZapOptions(getZapOptions()...)))
}

func InitKLogger() {
	klog.SetLogger(kitexzap.NewLogger(kitexzap.WithCoreEnc(getEncoder()),
		kitexzap.WithCoreWs(getLogWriter()),
		kitexzap.WithCoreLevel(getAtomicLevel()),
		kitexzap.WithZapOptions(getZapOptions()...)))
}

func GetDBLogger() gormlogger.Interface {
	logger := zapgorm2.New(getDBLogger())
	return logger.LogMode(gormlogger.Info)
}
