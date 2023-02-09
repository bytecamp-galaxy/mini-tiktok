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
	Writer io.Writer = os.Stdout
)

func SetOutput(path string) {
	switch path {
	case "stdout":
		Writer = os.Stdout
	case "stderr":
		Writer = os.Stderr
	default:
		path = "./" + path
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		Writer = f
	}
}

func InitHLogger() {
	hlog.SetLogger(hertzzap.NewLogger(hertzzap.WithCoreEnc(getEncoder()),
		hertzzap.WithCoreWs(getLogWriter()),
		hertzzap.WithCoreLevel(getAtomicLevel()),
		hertzzap.WithZapOptions(getZapOptions()...)))

	hlog.SetOutput(Writer)

}

func InitKLogger() {
	klog.SetLogger(kitexzap.NewLogger(kitexzap.WithCoreEnc(getEncoder()),
		kitexzap.WithCoreWs(getLogWriter()),
		kitexzap.WithCoreLevel(getAtomicLevel()),
		kitexzap.WithZapOptions(getZapOptions()...)))

	klog.SetOutput(Writer)
}

func GetDBLogger() gormlogger.Interface {
	logger := zapgorm2.New(getDBLogger())
	return logger.LogMode(gormlogger.Info)
}
