// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	Collect    *collect
	Lang       *lang
	Repository *repository
	User       *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Collect = &Q.Collect
	Lang = &Q.Lang
	Repository = &Q.Repository
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Collect:    newCollect(db, opts...),
		Lang:       newLang(db, opts...),
		Repository: newRepository(db, opts...),
		User:       newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Collect    collect
	Lang       lang
	Repository repository
	User       user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Collect:    q.Collect.clone(db),
		Lang:       q.Lang.clone(db),
		Repository: q.Repository.clone(db),
		User:       q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Collect:    q.Collect.replaceDB(db),
		Lang:       q.Lang.replaceDB(db),
		Repository: q.Repository.replaceDB(db),
		User:       q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Collect    ICollectDo
	Lang       ILangDo
	Repository IRepositoryDo
	User       IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Collect:    q.Collect.WithContext(ctx),
		Lang:       q.Lang.WithContext(ctx),
		Repository: q.Repository.WithContext(ctx),
		User:       q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
