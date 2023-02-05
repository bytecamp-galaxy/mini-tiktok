package main

import (
	model2 "github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/dal/query",                                             // relative to project root
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	mysql.Init(true)
	g.UseDB(mysql.DB)

	// Generate struct `User` based on table `users`
	// g.GenerateModel("users")

	g.ApplyBasic(model2.User{}, model2.Video{}, model2.Comment{}, model2.FollowRelation{}, model2.FavoriteRelation{})
	g.Execute()
}
