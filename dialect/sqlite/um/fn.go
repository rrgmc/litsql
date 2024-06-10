package um

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/sqlite"
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func Apply(f func(a sqlite.UpdateModApply)) sqlite.UpdateMod {
	return ium.Apply(f)
}

func CrossJoin(table string) JoinChain {
	return ium.CrossJoin[tag.UpdateTag](table)
}

func CrossJoinExpr(table litsql.Expression) JoinChain {
	return ium.CrossJoinExpr[tag.UpdateTag](table)
}

func From(table string) FromChain {
	return ium.From[tag.UpdateTag](table)
}

func FromExpr(table litsql.Expression) FromChain {
	return ium.FromExpr[tag.UpdateTag](table)
}

func FromQuery(q sqlite.SelectQuery) FromChain {
	return ium.FromQuery[tag.UpdateTag, tag.SelectTag](q)
}

func FullJoin(table string) JoinChain {
	return ium.FullJoin[tag.UpdateTag](table)
}

func FullJoinExpr(table litsql.Expression) JoinChain {
	return ium.FullJoinExpr[tag.UpdateTag](table)
}

func InnerJoin(table string) JoinChain {
	return ium.InnerJoin[tag.UpdateTag](table)
}

func InnerJoinExpr(table litsql.Expression) JoinChain {
	return ium.InnerJoinExpr[tag.UpdateTag](table)
}

func LeftJoin(table string) JoinChain {
	return ium.LeftJoin[tag.UpdateTag](table)
}

func LeftJoinExpr(table litsql.Expression) JoinChain {
	return ium.LeftJoinExpr[tag.UpdateTag](table)
}

func Only() sqlite.UpdateMod {
	return ium.Only[tag.UpdateTag](true)
}

func Returning(clauses ...string) sqlite.UpdateMod {
	return ium.Returning[tag.UpdateTag](clauses...)
}

func RightJoin(table string) JoinChain {
	return ium.RightJoin[tag.UpdateTag](table)
}

func RightJoinExpr(table litsql.Expression) JoinChain {
	return ium.RightJoinExpr[tag.UpdateTag](table)
}

func StraightJoin(table string) JoinChain {
	return ium.StraightJoin[tag.UpdateTag](table)
}

func StraightJoinExpr(table litsql.Expression) JoinChain {
	return ium.StraightJoinExpr[tag.UpdateTag](table)
}

func Set(column string, arg any) sqlite.UpdateMod {
	return ium.Set[tag.UpdateTag](column, arg)
}

func SetArgNamed(column string, argumentName string) sqlite.UpdateMod {
	return ium.SetArgNamed[tag.UpdateTag](column, argumentName)
}

func SetClause(query string, args ...any) sqlite.UpdateMod {
	return ium.SetClause[tag.UpdateTag](query, args...)
}

func SetExpr(column string, value litsql.Expression) sqlite.UpdateMod {
	return ium.SetExpr[tag.UpdateTag](column, value)
}

func SetExprClause(assignment litsql.Expression) sqlite.UpdateMod {
	return ium.SetExprClause[tag.UpdateTag](assignment)
}

func SetQuery(column string, q sqlite.SelectQuery) sqlite.UpdateMod {
	return ium.SetQuery[tag.UpdateTag, tag.SelectTag](column, q)
}

func SetString(column string, right string) sqlite.UpdateMod {
	return ium.SetString[tag.UpdateTag](column, right)
}

func Table(name string) sqlite.UpdateMod {
	return ium.Table[tag.UpdateTag](name)
}

func Where(condition string) sqlite.UpdateMod {
	return ium.Where[tag.UpdateTag](condition)
}

func WhereClause(query string, args ...any) sqlite.UpdateMod {
	return ium.WhereClause[tag.UpdateTag](query, args...)
}

func WhereExpr(condition litsql.Expression) sqlite.UpdateMod {
	return ium.WhereExpr[tag.UpdateTag](condition)
}

func With(name string, columns ...string) WithChain {
	return ium.With[tag.UpdateTag](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return ium.WithExpr[tag.UpdateTag](name, columns...)
}
