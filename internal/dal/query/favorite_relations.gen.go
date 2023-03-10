// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
)

func newFavoriteRelation(db *gorm.DB, opts ...gen.DOOption) favoriteRelation {
	_favoriteRelation := favoriteRelation{}

	_favoriteRelation.favoriteRelationDo.UseDB(db, opts...)
	_favoriteRelation.favoriteRelationDo.UseModel(&model.FavoriteRelation{})

	tableName := _favoriteRelation.favoriteRelationDo.TableName()
	_favoriteRelation.ALL = field.NewAsterisk(tableName)
	_favoriteRelation.ID = field.NewInt64(tableName, "id")
	_favoriteRelation.UserID = field.NewInt64(tableName, "user_id")
	_favoriteRelation.VideoID = field.NewInt64(tableName, "video_id")
	_favoriteRelation.User = favoriteRelationBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
	}

	_favoriteRelation.Video = favoriteRelationBelongsToVideo{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Video", "model.Video"),
		Author: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Video.Author", "model.User"),
		},
	}

	_favoriteRelation.fillFieldMap()

	return _favoriteRelation
}

type favoriteRelation struct {
	favoriteRelationDo

	ALL     field.Asterisk
	ID      field.Int64
	UserID  field.Int64
	VideoID field.Int64
	User    favoriteRelationBelongsToUser

	Video favoriteRelationBelongsToVideo

	fieldMap map[string]field.Expr
}

func (f favoriteRelation) Table(newTableName string) *favoriteRelation {
	f.favoriteRelationDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f favoriteRelation) As(alias string) *favoriteRelation {
	f.favoriteRelationDo.DO = *(f.favoriteRelationDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *favoriteRelation) updateTableName(table string) *favoriteRelation {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.UserID = field.NewInt64(table, "user_id")
	f.VideoID = field.NewInt64(table, "video_id")

	f.fillFieldMap()

	return f
}

func (f *favoriteRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *favoriteRelation) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 5)
	f.fieldMap["id"] = f.ID
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["video_id"] = f.VideoID

}

func (f favoriteRelation) clone(db *gorm.DB) favoriteRelation {
	f.favoriteRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f favoriteRelation) replaceDB(db *gorm.DB) favoriteRelation {
	f.favoriteRelationDo.ReplaceDB(db)
	return f
}

type favoriteRelationBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a favoriteRelationBelongsToUser) Where(conds ...field.Expr) *favoriteRelationBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a favoriteRelationBelongsToUser) WithContext(ctx context.Context) *favoriteRelationBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a favoriteRelationBelongsToUser) Model(m *model.FavoriteRelation) *favoriteRelationBelongsToUserTx {
	return &favoriteRelationBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type favoriteRelationBelongsToUserTx struct{ tx *gorm.Association }

func (a favoriteRelationBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a favoriteRelationBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a favoriteRelationBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a favoriteRelationBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a favoriteRelationBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a favoriteRelationBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type favoriteRelationBelongsToVideo struct {
	db *gorm.DB

	field.RelationField

	Author struct {
		field.RelationField
	}
}

func (a favoriteRelationBelongsToVideo) Where(conds ...field.Expr) *favoriteRelationBelongsToVideo {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a favoriteRelationBelongsToVideo) WithContext(ctx context.Context) *favoriteRelationBelongsToVideo {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a favoriteRelationBelongsToVideo) Model(m *model.FavoriteRelation) *favoriteRelationBelongsToVideoTx {
	return &favoriteRelationBelongsToVideoTx{a.db.Model(m).Association(a.Name())}
}

type favoriteRelationBelongsToVideoTx struct{ tx *gorm.Association }

func (a favoriteRelationBelongsToVideoTx) Find() (result *model.Video, err error) {
	return result, a.tx.Find(&result)
}

func (a favoriteRelationBelongsToVideoTx) Append(values ...*model.Video) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a favoriteRelationBelongsToVideoTx) Replace(values ...*model.Video) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a favoriteRelationBelongsToVideoTx) Delete(values ...*model.Video) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a favoriteRelationBelongsToVideoTx) Clear() error {
	return a.tx.Clear()
}

func (a favoriteRelationBelongsToVideoTx) Count() int64 {
	return a.tx.Count()
}

type favoriteRelationDo struct{ gen.DO }

type IFavoriteRelationDo interface {
	gen.SubQuery
	Debug() IFavoriteRelationDo
	WithContext(ctx context.Context) IFavoriteRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFavoriteRelationDo
	WriteDB() IFavoriteRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFavoriteRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFavoriteRelationDo
	Not(conds ...gen.Condition) IFavoriteRelationDo
	Or(conds ...gen.Condition) IFavoriteRelationDo
	Select(conds ...field.Expr) IFavoriteRelationDo
	Where(conds ...gen.Condition) IFavoriteRelationDo
	Order(conds ...field.Expr) IFavoriteRelationDo
	Distinct(cols ...field.Expr) IFavoriteRelationDo
	Omit(cols ...field.Expr) IFavoriteRelationDo
	Join(table schema.Tabler, on ...field.Expr) IFavoriteRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFavoriteRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFavoriteRelationDo
	Group(cols ...field.Expr) IFavoriteRelationDo
	Having(conds ...gen.Condition) IFavoriteRelationDo
	Limit(limit int) IFavoriteRelationDo
	Offset(offset int) IFavoriteRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFavoriteRelationDo
	Unscoped() IFavoriteRelationDo
	Create(values ...*model.FavoriteRelation) error
	CreateInBatches(values []*model.FavoriteRelation, batchSize int) error
	Save(values ...*model.FavoriteRelation) error
	First() (*model.FavoriteRelation, error)
	Take() (*model.FavoriteRelation, error)
	Last() (*model.FavoriteRelation, error)
	Find() ([]*model.FavoriteRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FavoriteRelation, err error)
	FindInBatches(result *[]*model.FavoriteRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FavoriteRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFavoriteRelationDo
	Assign(attrs ...field.AssignExpr) IFavoriteRelationDo
	Joins(fields ...field.RelationField) IFavoriteRelationDo
	Preload(fields ...field.RelationField) IFavoriteRelationDo
	FirstOrInit() (*model.FavoriteRelation, error)
	FirstOrCreate() (*model.FavoriteRelation, error)
	FindByPage(offset int, limit int) (result []*model.FavoriteRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFavoriteRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f favoriteRelationDo) Debug() IFavoriteRelationDo {
	return f.withDO(f.DO.Debug())
}

func (f favoriteRelationDo) WithContext(ctx context.Context) IFavoriteRelationDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f favoriteRelationDo) ReadDB() IFavoriteRelationDo {
	return f.Clauses(dbresolver.Read)
}

func (f favoriteRelationDo) WriteDB() IFavoriteRelationDo {
	return f.Clauses(dbresolver.Write)
}

func (f favoriteRelationDo) Session(config *gorm.Session) IFavoriteRelationDo {
	return f.withDO(f.DO.Session(config))
}

func (f favoriteRelationDo) Clauses(conds ...clause.Expression) IFavoriteRelationDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f favoriteRelationDo) Returning(value interface{}, columns ...string) IFavoriteRelationDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f favoriteRelationDo) Not(conds ...gen.Condition) IFavoriteRelationDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f favoriteRelationDo) Or(conds ...gen.Condition) IFavoriteRelationDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f favoriteRelationDo) Select(conds ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f favoriteRelationDo) Where(conds ...gen.Condition) IFavoriteRelationDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f favoriteRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFavoriteRelationDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f favoriteRelationDo) Order(conds ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f favoriteRelationDo) Distinct(cols ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f favoriteRelationDo) Omit(cols ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f favoriteRelationDo) Join(table schema.Tabler, on ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f favoriteRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f favoriteRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f favoriteRelationDo) Group(cols ...field.Expr) IFavoriteRelationDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f favoriteRelationDo) Having(conds ...gen.Condition) IFavoriteRelationDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f favoriteRelationDo) Limit(limit int) IFavoriteRelationDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f favoriteRelationDo) Offset(offset int) IFavoriteRelationDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f favoriteRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFavoriteRelationDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f favoriteRelationDo) Unscoped() IFavoriteRelationDo {
	return f.withDO(f.DO.Unscoped())
}

func (f favoriteRelationDo) Create(values ...*model.FavoriteRelation) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f favoriteRelationDo) CreateInBatches(values []*model.FavoriteRelation, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f favoriteRelationDo) Save(values ...*model.FavoriteRelation) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f favoriteRelationDo) First() (*model.FavoriteRelation, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteRelation), nil
	}
}

func (f favoriteRelationDo) Take() (*model.FavoriteRelation, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteRelation), nil
	}
}

func (f favoriteRelationDo) Last() (*model.FavoriteRelation, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteRelation), nil
	}
}

func (f favoriteRelationDo) Find() ([]*model.FavoriteRelation, error) {
	result, err := f.DO.Find()
	return result.([]*model.FavoriteRelation), err
}

func (f favoriteRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FavoriteRelation, err error) {
	buf := make([]*model.FavoriteRelation, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f favoriteRelationDo) FindInBatches(result *[]*model.FavoriteRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f favoriteRelationDo) Attrs(attrs ...field.AssignExpr) IFavoriteRelationDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f favoriteRelationDo) Assign(attrs ...field.AssignExpr) IFavoriteRelationDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f favoriteRelationDo) Joins(fields ...field.RelationField) IFavoriteRelationDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f favoriteRelationDo) Preload(fields ...field.RelationField) IFavoriteRelationDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f favoriteRelationDo) FirstOrInit() (*model.FavoriteRelation, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteRelation), nil
	}
}

func (f favoriteRelationDo) FirstOrCreate() (*model.FavoriteRelation, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FavoriteRelation), nil
	}
}

func (f favoriteRelationDo) FindByPage(offset int, limit int) (result []*model.FavoriteRelation, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f favoriteRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f favoriteRelationDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f favoriteRelationDo) Delete(models ...*model.FavoriteRelation) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *favoriteRelationDo) withDO(do gen.Dao) *favoriteRelationDo {
	f.DO = *do.(*gen.DO)
	return f
}
