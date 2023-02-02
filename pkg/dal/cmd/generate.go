package main

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dal/query",                                                  // relative to project root
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	mysql.Init(true)
	g.UseDB(mysql.DB)

	// Generate struct `User` based on table `users`
	// g.GenerateModel("users")

	g.ApplyBasic(model.User{}, model.Video{}, model.Comment{}, model.Relation{})
	g.Execute()
}
