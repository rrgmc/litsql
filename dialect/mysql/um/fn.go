package um

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func Apply(f func(a mysql.UpdateModApply)) mysql.UpdateMod {
	return ium.Apply(f)
}

func From(table string) FromChain {
	return ium.From[tag.UpdateTag](table)
}

func FromE(table litsql.Expression) FromChain {
	return ium.FromExpr[tag.UpdateTag](table)
}

func FromQ(q mysql.SelectQuery) FromChain {
	return ium.FromQuery[tag.UpdateTag, tag.SelectTag](q)
}

func InnerJoin(table string) JoinChain {
	return ium.InnerJoin[tag.UpdateTag](table)
}

func InnerJoinE(table litsql.Expression) JoinChain {
	return ium.InnerJoinExpr[tag.UpdateTag](table)
}

func LeftJoin(table string) JoinChain {
	return ium.LeftJoin[tag.UpdateTag](table)
}

func LeftJoinE(table litsql.Expression) JoinChain {
	return ium.LeftJoinExpr[tag.UpdateTag](table)
}

func RightJoin(table string) JoinChain {
	return ium.RightJoin[tag.UpdateTag](table)
}

func RightJoinE(table litsql.Expression) JoinChain {
	return ium.RightJoinExpr[tag.UpdateTag](table)
}

func FullJoin(table string) JoinChain {
	return ium.FullJoin[tag.UpdateTag](table)
}

func FullJoinE(table litsql.Expression) JoinChain {
	return ium.FullJoinExpr[tag.UpdateTag](table)
}

func CrossJoin(table string) JoinChain {
	return ium.CrossJoin[tag.UpdateTag](table)
}

func CrossJoinE(table litsql.Expression) JoinChain {
	return ium.CrossJoinExpr[tag.UpdateTag](table)
}

func StraightJoin(table string) JoinChain {
	return ium.StraightJoin[tag.UpdateTag](table)
}

func StraightJoinE(table litsql.Expression) JoinChain {
	return ium.StraightJoinExpr[tag.UpdateTag](table)
}

func Only() mysql.UpdateMod {
	return ium.Only[tag.UpdateTag](true)
}

func Returning(clauses ...string) mysql.UpdateMod {
	return ium.Returning[tag.UpdateTag](clauses...)
}

func Set(column string, arg any) mysql.UpdateMod {
	return ium.Set[tag.UpdateTag](column, arg)
}

func SetAN(column string, argumentName string) mysql.UpdateMod {
	return ium.SetArgNamed[tag.UpdateTag](column, argumentName)
}

func SetE(column string, value litsql.Expression) mysql.UpdateMod {
	return ium.SetExpr[tag.UpdateTag](column, value)
}

func SetQ(column string, q mysql.SelectQuery) mysql.UpdateMod {
	return ium.SetQuery[tag.UpdateTag, tag.SelectTag](column, q)
}

func SetS(column string, right string) mysql.UpdateMod {
	return ium.SetString[tag.UpdateTag](column, right)
}

func SetC(query string, args ...any) mysql.UpdateMod {
	return ium.SetClause[tag.UpdateTag](query, args...)
}

func SetRE(assignment litsql.Expression) mysql.UpdateMod {
	return ium.SetExprClause[tag.UpdateTag](assignment)
}

func Table(name string) mysql.UpdateMod {
	return ium.Table[tag.UpdateTag](name)
}

func Where(condition string) mysql.UpdateMod {
	return ium.Where[tag.UpdateTag](condition)
}

func WhereE(condition litsql.Expression) mysql.UpdateMod {
	return ium.WhereExpr[tag.UpdateTag](condition)
}

func WhereC(query string, args ...any) mysql.UpdateMod {
	return ium.WhereClause[tag.UpdateTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return ium.With[tag.UpdateTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return ium.WithExpr[tag.UpdateTag](name, columns...)
}
