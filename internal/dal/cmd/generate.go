package main

import (
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/dal/query",                                             // relative to project root
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// Generate struct `User` based on table `users`
	// mysql.Init(true)
	// g.UseDB(mysql.DB)
	// g.GenerateModel("users")

	g.ApplyBasic(model.User{}, model.Video{}, model.Comment{}, model.FollowRelation{}, model.FavoriteRelation{})
	g.Execute()
}
