package dal

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
)

func Init(migrated bool) {
	mysql.Init(migrated)
	query.SetDefault(mysql.DB)
}
