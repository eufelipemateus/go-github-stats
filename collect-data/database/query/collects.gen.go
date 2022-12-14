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

	"github.com/eufelipemateus/go-github-stats/collect-data/models"
)

func newCollect(db *gorm.DB, opts ...gen.DOOption) collect {
	_collect := collect{}

	_collect.collectDo.UseDB(db, opts...)
	_collect.collectDo.UseModel(&models.Collect{})

	tableName := _collect.collectDo.TableName()
	_collect.ALL = field.NewAsterisk(tableName)
	_collect.ID = field.NewUint(tableName, "id")
	_collect.CountLanguages = field.NewInt(tableName, "count_languages")
	_collect.CountFollowers = field.NewInt(tableName, "count_followers")
	_collect.CountRepos = field.NewInt(tableName, "count_repos")
	_collect.CountStars = field.NewInt(tableName, "count_stars")
	_collect.UserID = field.NewUint(tableName, "user_id")
	_collect.CollectStartedAt = field.NewTime(tableName, "collect_started_at")
	_collect.FinishedAt = field.NewTime(tableName, "finished_at")
	_collect.User = collectBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "models.User"),
		Following: struct {
			field.RelationField
			User struct {
				field.RelationField
			}
			Followers struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("User.Following", "models.Collect"),
			User: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Following.User", "models.User"),
			},
			Followers: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Following.Followers", "models.User"),
			},
		},
	}

	_collect.Followers = collectManyToManyFollowers{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Followers", "models.User"),
	}

	_collect.fillFieldMap()

	return _collect
}

type collect struct {
	collectDo

	ALL              field.Asterisk
	ID               field.Uint
	CountLanguages   field.Int
	CountFollowers   field.Int
	CountRepos       field.Int
	CountStars       field.Int
	UserID           field.Uint
	CollectStartedAt field.Time
	FinishedAt       field.Time
	User             collectBelongsToUser

	Followers collectManyToManyFollowers

	fieldMap map[string]field.Expr
}

func (c collect) Table(newTableName string) *collect {
	c.collectDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c collect) As(alias string) *collect {
	c.collectDo.DO = *(c.collectDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *collect) updateTableName(table string) *collect {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CountLanguages = field.NewInt(table, "count_languages")
	c.CountFollowers = field.NewInt(table, "count_followers")
	c.CountRepos = field.NewInt(table, "count_repos")
	c.CountStars = field.NewInt(table, "count_stars")
	c.UserID = field.NewUint(table, "user_id")
	c.CollectStartedAt = field.NewTime(table, "collect_started_at")
	c.FinishedAt = field.NewTime(table, "finished_at")

	c.fillFieldMap()

	return c
}

func (c *collect) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *collect) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["count_languages"] = c.CountLanguages
	c.fieldMap["count_followers"] = c.CountFollowers
	c.fieldMap["count_repos"] = c.CountRepos
	c.fieldMap["count_stars"] = c.CountStars
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["collect_started_at"] = c.CollectStartedAt
	c.fieldMap["finished_at"] = c.FinishedAt

}

func (c collect) clone(db *gorm.DB) collect {
	c.collectDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c collect) replaceDB(db *gorm.DB) collect {
	c.collectDo.ReplaceDB(db)
	return c
}

type collectBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	Following struct {
		field.RelationField
		User struct {
			field.RelationField
		}
		Followers struct {
			field.RelationField
		}
	}
}

func (a collectBelongsToUser) Where(conds ...field.Expr) *collectBelongsToUser {
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

func (a collectBelongsToUser) WithContext(ctx context.Context) *collectBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a collectBelongsToUser) Model(m *models.Collect) *collectBelongsToUserTx {
	return &collectBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type collectBelongsToUserTx struct{ tx *gorm.Association }

func (a collectBelongsToUserTx) Find() (result *models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a collectBelongsToUserTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a collectBelongsToUserTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a collectBelongsToUserTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a collectBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a collectBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type collectManyToManyFollowers struct {
	db *gorm.DB

	field.RelationField
}

func (a collectManyToManyFollowers) Where(conds ...field.Expr) *collectManyToManyFollowers {
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

func (a collectManyToManyFollowers) WithContext(ctx context.Context) *collectManyToManyFollowers {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a collectManyToManyFollowers) Model(m *models.Collect) *collectManyToManyFollowersTx {
	return &collectManyToManyFollowersTx{a.db.Model(m).Association(a.Name())}
}

type collectManyToManyFollowersTx struct{ tx *gorm.Association }

func (a collectManyToManyFollowersTx) Find() (result []*models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a collectManyToManyFollowersTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a collectManyToManyFollowersTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a collectManyToManyFollowersTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a collectManyToManyFollowersTx) Clear() error {
	return a.tx.Clear()
}

func (a collectManyToManyFollowersTx) Count() int64 {
	return a.tx.Count()
}

type collectDo struct{ gen.DO }

type ICollectDo interface {
	gen.SubQuery
	Debug() ICollectDo
	WithContext(ctx context.Context) ICollectDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICollectDo
	WriteDB() ICollectDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICollectDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICollectDo
	Not(conds ...gen.Condition) ICollectDo
	Or(conds ...gen.Condition) ICollectDo
	Select(conds ...field.Expr) ICollectDo
	Where(conds ...gen.Condition) ICollectDo
	Order(conds ...field.Expr) ICollectDo
	Distinct(cols ...field.Expr) ICollectDo
	Omit(cols ...field.Expr) ICollectDo
	Join(table schema.Tabler, on ...field.Expr) ICollectDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICollectDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICollectDo
	Group(cols ...field.Expr) ICollectDo
	Having(conds ...gen.Condition) ICollectDo
	Limit(limit int) ICollectDo
	Offset(offset int) ICollectDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectDo
	Unscoped() ICollectDo
	Create(values ...*models.Collect) error
	CreateInBatches(values []*models.Collect, batchSize int) error
	Save(values ...*models.Collect) error
	First() (*models.Collect, error)
	Take() (*models.Collect, error)
	Last() (*models.Collect, error)
	Find() ([]*models.Collect, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Collect, err error)
	FindInBatches(result *[]*models.Collect, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Collect) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICollectDo
	Assign(attrs ...field.AssignExpr) ICollectDo
	Joins(fields ...field.RelationField) ICollectDo
	Preload(fields ...field.RelationField) ICollectDo
	FirstOrInit() (*models.Collect, error)
	FirstOrCreate() (*models.Collect, error)
	FindByPage(offset int, limit int) (result []*models.Collect, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICollectDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c collectDo) Debug() ICollectDo {
	return c.withDO(c.DO.Debug())
}

func (c collectDo) WithContext(ctx context.Context) ICollectDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c collectDo) ReadDB() ICollectDo {
	return c.Clauses(dbresolver.Read)
}

func (c collectDo) WriteDB() ICollectDo {
	return c.Clauses(dbresolver.Write)
}

func (c collectDo) Session(config *gorm.Session) ICollectDo {
	return c.withDO(c.DO.Session(config))
}

func (c collectDo) Clauses(conds ...clause.Expression) ICollectDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c collectDo) Returning(value interface{}, columns ...string) ICollectDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c collectDo) Not(conds ...gen.Condition) ICollectDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c collectDo) Or(conds ...gen.Condition) ICollectDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c collectDo) Select(conds ...field.Expr) ICollectDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c collectDo) Where(conds ...gen.Condition) ICollectDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c collectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICollectDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c collectDo) Order(conds ...field.Expr) ICollectDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c collectDo) Distinct(cols ...field.Expr) ICollectDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c collectDo) Omit(cols ...field.Expr) ICollectDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c collectDo) Join(table schema.Tabler, on ...field.Expr) ICollectDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c collectDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICollectDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c collectDo) RightJoin(table schema.Tabler, on ...field.Expr) ICollectDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c collectDo) Group(cols ...field.Expr) ICollectDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c collectDo) Having(conds ...gen.Condition) ICollectDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c collectDo) Limit(limit int) ICollectDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c collectDo) Offset(offset int) ICollectDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c collectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c collectDo) Unscoped() ICollectDo {
	return c.withDO(c.DO.Unscoped())
}

func (c collectDo) Create(values ...*models.Collect) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c collectDo) CreateInBatches(values []*models.Collect, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c collectDo) Save(values ...*models.Collect) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c collectDo) First() (*models.Collect, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Collect), nil
	}
}

func (c collectDo) Take() (*models.Collect, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Collect), nil
	}
}

func (c collectDo) Last() (*models.Collect, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Collect), nil
	}
}

func (c collectDo) Find() ([]*models.Collect, error) {
	result, err := c.DO.Find()
	return result.([]*models.Collect), err
}

func (c collectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Collect, err error) {
	buf := make([]*models.Collect, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c collectDo) FindInBatches(result *[]*models.Collect, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c collectDo) Attrs(attrs ...field.AssignExpr) ICollectDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c collectDo) Assign(attrs ...field.AssignExpr) ICollectDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c collectDo) Joins(fields ...field.RelationField) ICollectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c collectDo) Preload(fields ...field.RelationField) ICollectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c collectDo) FirstOrInit() (*models.Collect, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Collect), nil
	}
}

func (c collectDo) FirstOrCreate() (*models.Collect, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Collect), nil
	}
}

func (c collectDo) FindByPage(offset int, limit int) (result []*models.Collect, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c collectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c collectDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c collectDo) Delete(models ...*models.Collect) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *collectDo) withDO(do gen.Dao) *collectDo {
	c.DO = *do.(*gen.DO)
	return c
}
