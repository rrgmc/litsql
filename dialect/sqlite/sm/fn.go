package sm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/sqlite"
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/ism"
)

func Apply(f func(a sqlite.SelectModApply)) sqlite.SelectMod {
	return ism.Apply(f)
}

func Columns(names ...string) sqlite.SelectMod {
	return ism.Columns[tag.SelectTag](names...)
}

func ColumnsClause(query string, args ...any) sqlite.SelectMod {
	return ism.ColumnsClause[tag.SelectTag](query, args...)
}

func ColumnsExpr(names ...litsql.Expression) sqlite.SelectMod {
	return ism.ColumnsExpr[tag.SelectTag](names...)
}

func CrossJoin(table string) JoinChain {
	return ism.CrossJoin[tag.SelectTag](table)
}

func CrossJoinExpr(table litsql.Expression) JoinChain {
	return ism.CrossJoinExpr[tag.SelectTag](table)
}

func Distinct(on ...string) sqlite.SelectMod {
	return ism.Distinct[tag.SelectTag](on...)
}

func DistinctExpr(on ...litsql.Expression) sqlite.SelectMod {
	return ism.DistinctExpr[tag.SelectTag](on...)
}

func Except(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Except[tag.SelectTag](q)
}

func ExceptAll(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.ExceptAll[tag.SelectTag](q)
}

func From(table string) FromChain {
	return &fromChainAdapter{chain: ism.From[tag.SelectTag](table)}
}

func FromExpr(table litsql.Expression) FromChain {
	return &fromChainAdapter{chain: ism.FromExpr[tag.SelectTag](table)}
}

func FromQuery(q sqlite.SelectQuery) FromChain {
	return &fromChainAdapter{chain: ism.FromQuery[tag.SelectTag, tag.SelectTag](q)}
}

func FullJoin(table string) JoinChain {
	return ism.FullJoin[tag.SelectTag](table)
}

func FullJoinExpr(table litsql.Expression) JoinChain {
	return ism.FullJoinExpr[tag.SelectTag](table)
}

func GroupBy(columns ...string) GroupByChain {
	return ism.GroupBy[tag.SelectTag](columns...)
}

func GroupByExpr(columns ...litsql.Expression) GroupByChain {
	return ism.GroupByExpr[tag.SelectTag](columns...)
}

func Having(condition string) sqlite.SelectMod {
	return ism.Having[tag.SelectTag](condition)
}

func HavingClause(query string, args ...any) sqlite.SelectMod {
	return ism.HavingClause[tag.SelectTag](query, args...)
}

func HavingExpr(condition litsql.Expression) sqlite.SelectMod {
	return ism.HavingExpr[tag.SelectTag](condition)
}

func InnerJoin(table string) JoinChain {
	return ism.InnerJoin[tag.SelectTag](table)
}

func InnerJoinExpr(table litsql.Expression) JoinChain {
	return ism.InnerJoinExpr[tag.SelectTag](table)
}

func Intersect(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func IntersectAll(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func LeftJoin(table string) JoinChain {
	return ism.LeftJoin[tag.SelectTag](table)
}

func LeftJoinExpr(table litsql.Expression) JoinChain {
	return ism.LeftJoinExpr[tag.SelectTag](table)
}

func Limit(count int) sqlite.SelectMod {
	return ism.Limit[tag.SelectTag](count)
}

func LimitArg(arg any) sqlite.SelectMod {
	return ism.LimitArg[tag.SelectTag](arg)
}

func LimitArgNamed(argumentName string) sqlite.SelectMod {
	return ism.LimitArgNamed[tag.SelectTag](argumentName)
}

func LimitExpr(count litsql.Expression) sqlite.SelectMod {
	return ism.LimitExpr[tag.SelectTag](count)
}

func Offset(count int) sqlite.SelectMod {
	return ism.Offset[tag.SelectTag](count)
}

func OffsetArg(arg any) sqlite.SelectMod {
	return ism.OffsetArg[tag.SelectTag](arg)
}

func OffsetArgNamed(argumentName string) sqlite.SelectMod {
	return ism.OffsetArgNamed[tag.SelectTag](argumentName)
}

func OffsetExpr(count litsql.Expression) sqlite.SelectMod {
	return ism.OffsetExpr[tag.SelectTag](count)
}

func OrderBy(names ...string) sqlite.SelectMod {
	return ism.OrderBy[tag.SelectTag](names...)
}

func OrderByExpr(names ...litsql.Expression) sqlite.SelectMod {
	return ism.OrderByExpr[tag.SelectTag](names...)
}

func RightJoin(table string) JoinChain {
	return ism.RightJoin[tag.SelectTag](table)
}

func RightJoinExpr(table litsql.Expression) JoinChain {
	return ism.RightJoinExpr[tag.SelectTag](table)
}

func StraightJoin(table string) JoinChain {
	return ism.StraightJoin[tag.SelectTag](table)
}

func StraightJoinExpr(table litsql.Expression) JoinChain {
	return ism.StraightJoinExpr[tag.SelectTag](table)
}

func Union(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Union[tag.SelectTag](q)
}

func UnionAll(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.UnionAll[tag.SelectTag](q)
}

func Where(condition string) sqlite.SelectMod {
	return ism.Where[tag.SelectTag](condition)
}

func WhereClause(query string, args ...any) sqlite.SelectMod {
	return ism.WhereClause[tag.SelectTag](query, args...)
}

func WhereExpr(condition litsql.Expression) sqlite.SelectMod {
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
