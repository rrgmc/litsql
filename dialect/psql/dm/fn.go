package dm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/idm"
)

func Apply(f func(a psql.DeleteModApply)) psql.DeleteMod {
	return idm.Apply(f)
}

func From(table string) psql.DeleteMod {
	return idm.From[tag.DeleteTag](table)
}

func InnerJoin(table string) JoinChain {
	return idm.InnerJoin[tag.DeleteTag](table)
}

func InnerJoinE(table litsql.Expression) JoinChain {
	return idm.InnerJoinExpr[tag.DeleteTag](table)
}

func LeftJoin(table string) JoinChain {
	return idm.LeftJoin[tag.DeleteTag](table)
}

func LeftJoinE(table litsql.Expression) JoinChain {
	return idm.LeftJoinExpr[tag.DeleteTag](table)
}

func RightJoin(table string) JoinChain {
	return idm.RightJoin[tag.DeleteTag](table)
}

func RightJoinE(table litsql.Expression) JoinChain {
	return idm.RightJoinExpr[tag.DeleteTag](table)
}

func FullJoin(table string) JoinChain {
	return idm.FullJoin[tag.DeleteTag](table)
}

func FullJoinE(table litsql.Expression) JoinChain {
	return idm.FullJoinExpr[tag.DeleteTag](table)
}

func CrossJoin(table string) JoinChain {
	return idm.CrossJoin[tag.DeleteTag](table)
}

func CrossJoinE(table litsql.Expression) JoinChain {
	return idm.CrossJoinExpr[tag.DeleteTag](table)
}

func StraightJoin(table string) JoinChain {
	return idm.StraightJoin[tag.DeleteTag](table)
}

func StraightJoinE(table litsql.Expression) JoinChain {
	return idm.StraightJoinExpr[tag.DeleteTag](table)
}

func Only() psql.DeleteMod {
	return idm.Only[tag.DeleteTag](true)
}

func Returning(clauses ...string) psql.DeleteMod {
	return idm.Returning[tag.DeleteTag](clauses...)
}

func Using(table string) FromChain {
	return idm.Using[tag.DeleteTag](table)
}

func UsingE(table litsql.Expression) FromChain {
	return idm.UsingExpr[tag.DeleteTag](table)
}

func UsingQ(q psql.SelectQuery) FromChain {
	return idm.UsingQuery[tag.DeleteTag, tag.SelectTag](q)
}

func Where(condition string) psql.DeleteMod {
	return idm.Where[tag.DeleteTag](condition)
}

func WhereE(condition litsql.Expression) psql.DeleteMod {
	return idm.WhereExpr[tag.DeleteTag](condition)
}

func WhereC(query string, args ...any) psql.DeleteMod {
	return idm.WhereClause[tag.DeleteTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return idm.With[tag.DeleteTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return idm.WithExpr[tag.DeleteTag](name, columns...)
}
