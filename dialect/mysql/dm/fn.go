package dm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/idm"
)

func Apply(f func(a mysql.DeleteModApply)) mysql.DeleteMod {
	return idm.Apply(f)
}

func From(table string) mysql.DeleteMod {
	return idm.From[tag.DeleteTag](table)
}

func InnerJoin(table string) JoinChain {
	return idm.InnerJoin[tag.DeleteTag](table)
}

func InnerJoinExpr(table litsql.Expression) JoinChain {
	return idm.InnerJoinExpr[tag.DeleteTag](table)
}

func LeftJoin(table string) JoinChain {
	return idm.LeftJoin[tag.DeleteTag](table)
}

func LeftJoinExpr(table litsql.Expression) JoinChain {
	return idm.LeftJoinExpr[tag.DeleteTag](table)
}

func RightJoin(table string) JoinChain {
	return idm.RightJoin[tag.DeleteTag](table)
}

func RightJoinExpr(table litsql.Expression) JoinChain {
	return idm.RightJoinExpr[tag.DeleteTag](table)
}

func FullJoin(table string) JoinChain {
	return idm.FullJoin[tag.DeleteTag](table)
}

func FullJoinExpr(table litsql.Expression) JoinChain {
	return idm.FullJoinExpr[tag.DeleteTag](table)
}

func CrossJoin(table string) JoinChain {
	return idm.CrossJoin[tag.DeleteTag](table)
}

func CrossJoinExpr(table litsql.Expression) JoinChain {
	return idm.CrossJoinExpr[tag.DeleteTag](table)
}

func StraightJoin(table string) JoinChain {
	return idm.StraightJoin[tag.DeleteTag](table)
}

func StraightJoinExpr(table litsql.Expression) JoinChain {
	return idm.StraightJoinExpr[tag.DeleteTag](table)
}

func Only() mysql.DeleteMod {
	return idm.Only[tag.DeleteTag]()
}

func Returning(clauses ...string) mysql.DeleteMod {
	return idm.Returning[tag.DeleteTag](clauses...)
}

func Using(table string) FromChain {
	return idm.Using[tag.DeleteTag](table)
}

func UsingExpr(table litsql.Expression) FromChain {
	return idm.UsingExpr[tag.DeleteTag](table)
}

func UsingQuery(q mysql.SelectQuery) FromChain {
	return idm.UsingQuery[tag.DeleteTag, tag.SelectTag](q)
}

func Where(condition string) mysql.DeleteMod {
	return idm.Where[tag.DeleteTag](condition)
}

func WhereExpr(condition litsql.Expression) mysql.DeleteMod {
	return idm.WhereExpr[tag.DeleteTag](condition)
}

func WhereClause(query string, args ...any) mysql.DeleteMod {
	return idm.WhereClause[tag.DeleteTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return idm.With[tag.DeleteTag](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return idm.WithExpr[tag.DeleteTag](name, columns...)
}
