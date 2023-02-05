package mysql

import (
	"fmt"
	model2 "github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

func Init(migrated bool) {
	var err error

	v := conf.Init()
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
		Logger:                 log.GetDBLogger(),
	})

	if err != nil {
		panic(err)
	}

	if migrated {
		// NOTE: concurrent `AutoMigrate` is not supported
		// only `AutoMigrate` when api server setup
		// drop table for the convenience of test
		if err := DB.Migrator().DropTable(&model2.User{}, &model2.Video{}, &model2.Comment{}, &model2.FollowRelation{}, &model2.FavoriteRelation{}); err != nil {
			panic(err)
		}
		if err := DB.AutoMigrate(&model2.User{}, &model2.Video{}, &model2.Comment{}, &model2.FollowRelation{}, &model2.FavoriteRelation{}); err != nil {
			panic(err)
		}
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
