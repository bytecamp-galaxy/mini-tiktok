package dal

import (
	"mini-tiktok-v2/pkg/dal/mysql"
	"mini-tiktok-v2/pkg/dal/query"
)

func Init() {
	mysql.Init()
	query.SetDefault(mysql.DB)
}
