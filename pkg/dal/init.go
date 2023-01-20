package dal

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
)

func Init() {
	mysql.Init()
	query.SetDefault(mysql.DB)
}
