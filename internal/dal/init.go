package dal

import (
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
)

func Init(migrated bool) {
	mysql.Init(migrated)
	query.SetDefault(mysql.DB)
}
