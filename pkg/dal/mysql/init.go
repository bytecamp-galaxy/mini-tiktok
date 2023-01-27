package mysql

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

func Init() {
	var err error

	v := conf.Init().V
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		v.GetString("mysql.user"),
		v.GetString("mysql.password"),
		v.GetString("mysql.host"),
		v.GetInt("mysql.port"),
		v.GetString("mysql.dbname"),
		v.GetString("mysql.charset"),
		v.GetBool("mysql.parseTime"),
		v.GetString("mysql.loc"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 log.InitDBLogger(),
	})
	if err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.Relation{}); err != nil {
		panic(err)
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
}
