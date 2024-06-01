package um

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/sqlite"
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func From(table string) FromChain {
	return ium.From[tag.UpdateTag](table)
}

func FromE(table litsql.Expression) FromChain {
	return ium.FromE[tag.UpdateTag](table)
}

func FromQ(q sqlite.SelectQuery) FromChain {
	return ium.FromQ[tag.UpdateTag, tag.SelectTag](q)
}

func Func(f func(a sqlite.UpdateModApply)) sqlite.UpdateMod {
	return ium.Func(f)
}

func InnerJoin(table string) JoinChain {
	return ium.InnerJoin[tag.UpdateTag](table)
}

func InnerJoinE(table litsql.Expression) JoinChain {
	return ium.InnerJoinE[tag.UpdateTag](table)
}

func LeftJoin(table string) JoinChain {
	return ium.LeftJoin[tag.UpdateTag](table)
}

func LeftJoinE(table litsql.Expression) JoinChain {
	return ium.LeftJoinE[tag.UpdateTag](table)
}

func RightJoin(table string) JoinChain {
	return ium.RightJoin[tag.UpdateTag](table)
}

func RightJoinE(table litsql.Expression) JoinChain {
	return ium.RightJoinE[tag.UpdateTag](table)
}

func FullJoin(table string) JoinChain {
	return ium.FullJoin[tag.UpdateTag](table)
}

func FullJoinE(table litsql.Expression) JoinChain {
	return ium.FullJoinE[tag.UpdateTag](table)
}

func CrossJoin(table string) JoinChain {
	return ium.CrossJoin[tag.UpdateTag](table)
}

func CrossJoinE(table litsql.Expression) JoinChain {
	return ium.CrossJoinE[tag.UpdateTag](table)
}

func StraightJoin(table string) JoinChain {
	return ium.StraightJoin[tag.UpdateTag](table)
}

func StraightJoinE(table litsql.Expression) JoinChain {
	return ium.StraightJoinE[tag.UpdateTag](table)
}

func Only() sqlite.UpdateMod {
	return ium.Only[tag.UpdateTag](true)
}

func Returning(clauses ...string) sqlite.UpdateMod {
	return ium.Returning[tag.UpdateTag](clauses...)
}

func Set(column string, arg any) sqlite.UpdateMod {
	return ium.Set[tag.UpdateTag](column, arg)
}

func SetE(column string, value litsql.Expression) sqlite.UpdateMod {
	return ium.SetE[tag.UpdateTag](column, value)
}

func SetQ(column string, q sqlite.SelectQuery) sqlite.UpdateMod {
	return ium.SetQ[tag.UpdateTag, tag.SelectTag](column, q)
}

func SetS(column string, right string) sqlite.UpdateMod {
	return ium.SetS[tag.UpdateTag](column, right)
}

func SetR(raw string) sqlite.UpdateMod {
	return ium.SetR[tag.UpdateTag](raw)
}

func SetRE(assignment litsql.Expression) sqlite.UpdateMod {
	return ium.SetRE[tag.UpdateTag](assignment)
}

func Table(name string) sqlite.UpdateMod {
	return ium.Table[tag.UpdateTag](name)
}

func Where(condition string) sqlite.UpdateMod {
	return ium.Where[tag.UpdateTag](condition)
}

func WhereE(condition litsql.Expression) sqlite.UpdateMod {
	return ium.WhereE[tag.UpdateTag](condition)
}

func WhereC(query string, args ...any) sqlite.UpdateMod {
	return ium.WhereC[tag.UpdateTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return ium.With[tag.UpdateTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return ium.WithE[tag.UpdateTag](name, columns...)
}
