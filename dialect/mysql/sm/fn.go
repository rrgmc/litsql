// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package sm

import (
	litsql "github.com/rrgmc/litsql"
	mysql "github.com/rrgmc/litsql/dialect/mysql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	ism "github.com/rrgmc/litsql/internal/ism"
)

func Apply(f func(a mysql.SelectModApply)) mysql.SelectMod {
	return ism.Apply[tag.SelectTag](f)
}

// Columns adds column names to SELECT queries.
func Columns(names ...string) mysql.SelectMod {
	return ism.Columns[tag.SelectTag](names...)
}

// ColumnsClause adds column names to SELECT queries.
func ColumnsClause(query string, args ...any) mysql.SelectMod {
	return ism.ColumnsClause[tag.SelectTag](query, args...)
}

// ColumnsExpr adds column names to SELECT queries.
func ColumnsExpr(names ...litsql.Expression) mysql.SelectMod {
	return ism.ColumnsExpr[tag.SelectTag](names...)
}

func CrossJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ism.CrossJoin[tag.SelectTag](table),
	}
}

func CrossJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ism.CrossJoinExpr[tag.SelectTag](table),
	}
}

func Distinct(on ...string) mysql.SelectMod {
	return ism.Distinct[tag.SelectTag](on...)
}

func DistinctExpr(on ...litsql.Expression) mysql.SelectMod {
	return ism.DistinctExpr[tag.SelectTag](on...)
}

func Except(q mysql.SelectQuery) mysql.SelectMod {
	return ism.Except[tag.SelectTag](q)
}

func ExceptAll(q mysql.SelectQuery) mysql.SelectMod {
	return ism.ExceptAll[tag.SelectTag](q)
}

func From(table string) FromChain {
	return &fromChainAdapter{
		chain: ism.From[tag.SelectTag](table),
	}
}

func FromExpr(table litsql.Expression) FromChain {
	return &fromChainAdapter{
		chain: ism.FromExpr[tag.SelectTag](table),
	}
}

func FromQuery(q mysql.SelectQuery) FromChain {
	return &fromChainAdapter{
		chain: ism.FromQuery[tag.SelectTag, tag.SelectTag](q),
	}
}

func FullJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ism.FullJoin[tag.SelectTag](table),
	}
}

func FullJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ism.FullJoinExpr[tag.SelectTag](table),
	}
}

func GroupBy(columns ...string) GroupByChain {
	return &groupByChainAdapter{
		chain: ism.GroupBy[tag.SelectTag](columns...),
	}
}

func GroupByExpr(columns ...litsql.Expression) GroupByChain {
	return &groupByChainAdapter{
		chain: ism.GroupByExpr[tag.SelectTag](columns...),
	}
}

func Having(condition string) mysql.SelectMod {
	return ism.Having[tag.SelectTag](condition)
}

func HavingClause(query string, args ...any) mysql.SelectMod {
	return ism.HavingClause[tag.SelectTag](query, args...)
}

func HavingExpr(condition litsql.Expression) mysql.SelectMod {
	return ism.HavingExpr[tag.SelectTag](condition)
}

func InnerJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ism.InnerJoin[tag.SelectTag](table),
	}
}

func InnerJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ism.InnerJoinExpr[tag.SelectTag](table),
	}
}

func Intersect(q mysql.SelectQuery) mysql.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func IntersectAll(q mysql.SelectQuery) mysql.SelectMod {
	return ism.IntersectAll[tag.SelectTag](q)
}

func LeftJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ism.LeftJoin[tag.SelectTag](table),
	}
}

func LeftJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ism.LeftJoinExpr[tag.SelectTag](table),
	}
}

func Limit(count int) mysql.SelectMod {
	return ism.Limit[tag.SelectTag](count)
}

func LimitArg(arg any) mysql.SelectMod {
	return ism.LimitArg[tag.SelectTag](arg)
}

func LimitArgNamed(argumentName string) mysql.SelectMod {
	return ism.LimitArgNamed[tag.SelectTag](argumentName)
}

func LimitExpr(count litsql.Expression) mysql.SelectMod {
	return ism.LimitExpr[tag.SelectTag](count)
}

func Offset(count int) mysql.SelectMod {
	return ism.Offset[tag.SelectTag](count)
}

func OffsetArg(arg any) mysql.SelectMod {
	return ism.OffsetArg[tag.SelectTag](arg)
}

func OffsetArgNamed(argumentName string) mysql.SelectMod {
	return ism.OffsetArgNamed[tag.SelectTag](argumentName)
}

func OffsetExpr(count litsql.Expression) mysql.SelectMod {
	return ism.OffsetExpr[tag.SelectTag](count)
}

func OrderBy(names ...string) mysql.SelectMod {
	return ism.OrderBy[tag.SelectTag](names...)
}

func OrderByExpr(names ...litsql.Expression) mysql.SelectMod {
	return ism.OrderByExpr[tag.SelectTag](names...)
}

func RightJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ism.RightJoin[tag.SelectTag](table),
	}
}

func RightJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ism.RightJoinExpr[tag.SelectTag](table),
	}
}

func StraightJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ism.StraightJoin[tag.SelectTag](table),
	}
}

func StraightJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ism.StraightJoinExpr[tag.SelectTag](table),
	}
}

func Union(q mysql.SelectQuery) mysql.SelectMod {
	return ism.Union[tag.SelectTag](q)
}

func UnionAll(q mysql.SelectQuery) mysql.SelectMod {
	return ism.UnionAll[tag.SelectTag](q)
}

func Where(condition string) mysql.SelectMod {
	return ism.Where[tag.SelectTag](condition)
}

func WhereClause(query string, args ...any) mysql.SelectMod {
	return ism.WhereClause[tag.SelectTag](query, args...)
}

func WhereExpr(condition litsql.Expression) mysql.SelectMod {
	return ism.WhereExpr[tag.SelectTag](condition)
}

func Window(name string) WindowChain {
	return ism.Window[tag.SelectTag](name)
}

func With(name string, columns ...string) WithChain {
	return ism.With[tag.SelectTag](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return ism.WithExpr[tag.SelectTag](name, columns...)
}
