package log

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	hertzzap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	gormlogger "gorm.io/gorm/logger"
	"io"
	"moul.io/zapgorm2"
	"os"
)

var (
	writer io.Writer = os.Stdout
)

func SetOutput(path string) {
	switch path {
	case "stdout":
		writer = os.Stdout
	case "stderr":
		writer = os.Stderr
	default:
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		writer = f
	}
}

func InitHLogger() {
	hlog.SetLogger(hertzzap.NewLogger(hertzzap.WithCoreEnc(getEncoder()),
		hertzzap.WithCoreWs(getLogWriter()),
		hertzzap.WithCoreLevel(getAtomicLevel()),
		hertzzap.WithZapOptions(getZapOptions()...)))
	hlog.SetOutput(writer)
}

func InitKLogger() {
	klog.SetLogger(kitexzap.NewLogger(kitexzap.WithCoreEnc(getEncoder()),
		kitexzap.WithCoreWs(getLogWriter()),
		kitexzap.WithCoreLevel(getAtomicLevel()),
		kitexzap.WithZapOptions(getZapOptions()...)))
	klog.SetOutput(writer)
}

func GetDBLogger() gormlogger.Interface {
	logger := zapgorm2.New(getDBLogger())
	return logger.LogMode(gormlogger.Info)
}
