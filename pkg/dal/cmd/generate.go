package main

import (
	"gorm.io/gen"
	"mini-tiktok-v2/pkg/dal/model"
	"mini-tiktok-v2/pkg/dal/mysql"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dal/query",                                                  // relative to project root
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	mysql.Init()
	g.UseDB(mysql.DB)

	// Generate struct `User` based on table `users`
	// g.GenerateModel("users")

	g.ApplyBasic(model.User{})
	g.Execute()
}
