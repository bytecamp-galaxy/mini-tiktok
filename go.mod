module github.com/bytecamp-galaxy/mini-tiktok

go 1.19

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.13.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cloudwego/hertz v0.5.1
	github.com/cloudwego/kitex v0.4.4
	github.com/gavv/httpexpect/v2 v2.10.0
	github.com/google/uuid v1.3.0
	github.com/hertz-contrib/gzip v0.0.1
	github.com/hertz-contrib/jwt v1.0.1
	github.com/hertz-contrib/logger/zap v0.0.0-20221227100845-46a8693d7847
	github.com/hertz-contrib/obs-opentelemetry/provider v0.0.0-20221123024949-68d0df9511cf
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.0.0-20221123024949-68d0df9511cf
	github.com/hertz-contrib/pprof v0.1.0
	github.com/hertz-contrib/registry/etcd v0.0.0-20221226122036-3c451682dc72
	github.com/hertz-contrib/websocket v0.0.1
	github.com/kitex-contrib/obs-opentelemetry v0.1.0
	github.com/kitex-contrib/obs-opentelemetry/logging/zap v0.0.0-20221109071748-a433b0b57972
	github.com/kitex-contrib/registry-etcd v0.1.0
	github.com/marmotedu/errors v1.0.2
	github.com/minio/minio-go/v7 v7.0.47
	github.com/novalagung/gubrak v1.0.0
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	github.com/u2takey/ffmpeg-go v0.4.1
	github.com/wagslane/go-password-validator v0.3.0
	go.opentelemetry.io/otel v1.9.0
	go.uber.org/zap v1.23.0
	golang.org/x/crypto v0.5.0
	golang.org/x/tools v0.5.0
	gorm.io/driver/mysql v1.4.5
	gorm.io/gen v0.3.19
	gorm.io/gorm v1.24.3
	gorm.io/plugin/dbresolver v1.4.0
	gorm.io/plugin/opentelemetry v0.1.0
	moul.io/zapgorm2 v1.2.0
)