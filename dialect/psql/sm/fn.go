package sm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/ism"
)

func Apply(f func(a psql.SelectModApply)) psql.SelectMod {
	return ism.Apply(f)
}

func Columns(names ...string) psql.SelectMod {
	return ism.Columns[tag.SelectTag](names...)
}

func ColumnsE(names ...litsql.Expression) psql.SelectMod {
	return ism.ColumnsExpr[tag.SelectTag](names...)
}

func ColumnsC(query string, args ...any) psql.SelectMod {
	return ism.ColumnsClause[tag.SelectTag](query, args...)
}

func Distinct(on ...string) psql.SelectMod {
	return ism.Distinct[tag.SelectTag](on...)
}

func DistinctE(on ...litsql.Expression) psql.SelectMod {
	return ism.DistinctExpr[tag.SelectTag](on...)
}

func From(table string) FromChain {
	return ism.From[tag.SelectTag](table)
}

func FromE(table litsql.Expression) FromChain {
	return ism.FromExpr[tag.SelectTag](table)
}

func FromQ(q psql.SelectQuery) FromChain {
	return ism.FromQuery[tag.SelectTag, tag.SelectTag](q)
}

func GroupBy(columns ...string) GroupByChain {
	return ism.GroupBy[tag.SelectTag](columns...)
}

func GroupByE(columns ...litsql.Expression) GroupByChain {
	return ism.GroupByExpr[tag.SelectTag](columns...)
}

func Having(condition string) psql.SelectMod {
	return ism.Having[tag.SelectTag](condition)
}

func HavingE(condition litsql.Expression) psql.SelectMod {
	return ism.HavingExpr[tag.SelectTag](condition)
}

func HavingC(query string, args ...any) psql.SelectMod {
	return ism.HavingClause[tag.SelectTag](query, args...)
}

func InnerJoin(table string) JoinChain {
	return ism.InnerJoin[tag.SelectTag](table)
}

func InnerJoinE(table litsql.Expression) JoinChain {
	return ism.InnerJoinExpr[tag.SelectTag](table)
}

func LeftJoin(table string) JoinChain {
	return ism.LeftJoin[tag.SelectTag](table)
}

func LeftJoinE(table litsql.Expression) JoinChain {
	return ism.LeftJoinExpr[tag.SelectTag](table)
}

func RightJoin(table string) JoinChain {
	return ism.RightJoin[tag.SelectTag](table)
}

func RightJoinE(table litsql.Expression) JoinChain {
	return ism.RightJoinExpr[tag.SelectTag](table)
}

func FullJoin(table string) JoinChain {
	return ism.FullJoin[tag.SelectTag](table)
}

func FullJoinE(table litsql.Expression) JoinChain {
	return ism.FullJoinExpr[tag.SelectTag](table)
}

func CrossJoin(table string) JoinChain {
	return ism.CrossJoin[tag.SelectTag](table)
}

func CrossJoinE(table litsql.Expression) JoinChain {
	return ism.CrossJoinExpr[tag.SelectTag](table)
}

func StraightJoin(table string) JoinChain {
	return ism.StraightJoin[tag.SelectTag](table)
}

func StraightJoinE(table litsql.Expression) JoinChain {
	return ism.StraightJoinExpr[tag.SelectTag](table)
}

func Limit(count int) psql.SelectMod {
	return ism.Limit[tag.SelectTag](count)
}

func LimitE(count litsql.Expression) psql.SelectMod {
	return ism.LimitExpr[tag.SelectTag](count)
}

func LimitA(arg any) psql.SelectMod {
	return ism.LimitArg[tag.SelectTag](arg)
}

func LimitAN(argumentName string) psql.SelectMod {
	return ism.LimitArgNamed[tag.SelectTag](argumentName)
}

func Offset(count int) psql.SelectMod {
	return ism.Offset[tag.SelectTag](count)
}

func OffsetE(count litsql.Expression) psql.SelectMod {
	return ism.OffsetExpr[tag.SelectTag](count)
}

func OffsetA(arg any) psql.SelectMod {
	return ism.OffsetArg[tag.SelectTag](arg)
}

func OffsetAN(argumentName string) psql.SelectMod {
	return ism.OffsetArgNamed[tag.SelectTag](argumentName)
}

func OrderBy(names ...string) psql.SelectMod {
	return ism.OrderBy[tag.SelectTag](names...)
}

func OrderByE(names ...litsql.Expression) psql.SelectMod {
	return ism.OrderByExpr[tag.SelectTag](names...)
}

func Union(q psql.SelectQuery) psql.SelectMod {
	return ism.Union[tag.SelectTag](q)
}

func UnionAll(q psql.SelectQuery) psql.SelectMod {
	return ism.UnionAll[tag.SelectTag](q)
}

func Intersect(q psql.SelectQuery) psql.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func IntersectAll(q psql.SelectQuery) psql.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func Except(q psql.SelectQuery) psql.SelectMod {
	return ism.Except[tag.SelectTag](q)
}

func ExceptAll(q psql.SelectQuery) psql.SelectMod {
	return ism.ExceptAll[tag.SelectTag](q)
}

func Where(condition string) psql.SelectMod {
	return ism.Where[tag.SelectTag](condition)
}

func WhereE(condition litsql.Expression) psql.SelectMod {
	return ism.WhereExpr[tag.SelectTag](condition)
}

func WhereC(query string, args ...any) psql.SelectMod {
	return ism.WhereClause[tag.SelectTag](query, args...)
}

func Window(name string) WindowChain {
	return ism.Window[tag.SelectTag](name)
}

func With(name string, columns ...string) WithChain {
	return ism.With[tag.SelectTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return ism.WithExpr[tag.SelectTag](name, columns...)
}
