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

func ColumnsE(names ...litsql.Expression) sqlite.SelectMod {
	return ism.ColumnsE[tag.SelectTag](names...)
}

func Distinct(on ...string) sqlite.SelectMod {
	return ism.Distinct[tag.SelectTag](on...)
}

func DistinctE(on ...litsql.Expression) sqlite.SelectMod {
	return ism.DistinctE[tag.SelectTag](on...)
}

func From(table string) FromChain {
	return &fromChainAdapter{chain: ism.From[tag.SelectTag](table)}
}

func FromE(table litsql.Expression) FromChain {
	return &fromChainAdapter{chain: ism.FromE[tag.SelectTag](table)}
}

func FromQ(q sqlite.SelectQuery) FromChain {
	return &fromChainAdapter{chain: ism.FromQ[tag.SelectTag, tag.SelectTag](q)}
}

func GroupBy(columns ...string) GroupByChain {
	return ism.GroupBy[tag.SelectTag](columns...)
}

func GroupByE(columns ...litsql.Expression) GroupByChain {
	return ism.GroupByE[tag.SelectTag](columns...)
}

func Having(condition string) sqlite.SelectMod {
	return ism.Having[tag.SelectTag](condition)
}

func HavingE(condition litsql.Expression) sqlite.SelectMod {
	return ism.HavingE[tag.SelectTag](condition)
}

func HavingC(query string, args ...any) sqlite.SelectMod {
	return ism.HavingC[tag.SelectTag](query, args...)
}

func InnerJoin(table string) JoinChain {
	return ism.InnerJoin[tag.SelectTag](table)
}

func InnerJoinE(table litsql.Expression) JoinChain {
	return ism.InnerJoinE[tag.SelectTag](table)
}

func LeftJoin(table string) JoinChain {
	return ism.LeftJoin[tag.SelectTag](table)
}

func LeftJoinE(table litsql.Expression) JoinChain {
	return ism.LeftJoinE[tag.SelectTag](table)
}

func RightJoin(table string) JoinChain {
	return ism.RightJoin[tag.SelectTag](table)
}

func RightJoinE(table litsql.Expression) JoinChain {
	return ism.RightJoinE[tag.SelectTag](table)
}

func FullJoin(table string) JoinChain {
	return ism.FullJoin[tag.SelectTag](table)
}

func FullJoinE(table litsql.Expression) JoinChain {
	return ism.FullJoinE[tag.SelectTag](table)
}

func CrossJoin(table string) JoinChain {
	return ism.CrossJoin[tag.SelectTag](table)
}

func CrossJoinE(table litsql.Expression) JoinChain {
	return ism.CrossJoinE[tag.SelectTag](table)
}

func StraightJoin(table string) JoinChain {
	return ism.StraightJoin[tag.SelectTag](table)
}

func StraightJoinE(table litsql.Expression) JoinChain {
	return ism.StraightJoinE[tag.SelectTag](table)
}

func Limit(count int) sqlite.SelectMod {
	return ism.Limit[tag.SelectTag](count)
}

func LimitE(count litsql.Expression) sqlite.SelectMod {
	return ism.LimitE[tag.SelectTag](count)
}

func LimitA(arg any) sqlite.SelectMod {
	return ism.LimitA[tag.SelectTag](arg)
}

func Offset(count int) sqlite.SelectMod {
	return ism.Offset[tag.SelectTag](count)
}

func OffsetE(count litsql.Expression) sqlite.SelectMod {
	return ism.OffsetE[tag.SelectTag](count)
}

func OffsetA(arg any) sqlite.SelectMod {
	return ism.OffsetA[tag.SelectTag](arg)
}

func OrderBy(names ...string) sqlite.SelectMod {
	return ism.OrderBy[tag.SelectTag](names...)
}

func OrderByE(names ...litsql.Expression) sqlite.SelectMod {
	return ism.OrderByE[tag.SelectTag](names...)
}

func Union(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Union[tag.SelectTag](q)
}

func UnionAll(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.UnionAll[tag.SelectTag](q)
}

func Intersect(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func IntersectAll(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Intersect[tag.SelectTag](q)
}

func Except(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.Except[tag.SelectTag](q)
}

func ExceptAll(q sqlite.SelectQuery) sqlite.SelectMod {
	return ism.ExceptAll[tag.SelectTag](q)
}

func Where(condition string) sqlite.SelectMod {
	return ism.Where[tag.SelectTag](condition)
}

func WhereE(condition litsql.Expression) sqlite.SelectMod {
	return ism.WhereE[tag.SelectTag](condition)
}

func WhereC(query string, args ...any) sqlite.SelectMod {
	return ism.WhereC[tag.SelectTag](query, args...)
}

func Window(name string) WindowChain {
	return ism.Window[tag.SelectTag](name)
}

func With(name string, columns ...string) WithChain {
	return ism.With[tag.SelectTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return ism.WithE[tag.SelectTag](name, columns...)
}
