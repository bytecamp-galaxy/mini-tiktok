package main

import (
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
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

	g.ApplyBasic(model.User{}, model.Video{}, model.Comment{}, model.FollowRelation{}, model.FavoriteRelation{})
	g.Execute()
}
