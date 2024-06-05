package um

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func Apply(f func(a psql.UpdateModApply)) psql.UpdateMod {
	return ium.Apply(f)
}

func From(table string) FromChain {
	return ium.From[tag.UpdateTag](table)
}

func FromE(table litsql.Expression) FromChain {
	return ium.FromE[tag.UpdateTag](table)
}

func FromQ(q psql.SelectQuery) FromChain {
	return ium.FromQ[tag.UpdateTag, tag.SelectTag](q)
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

func Only() psql.UpdateMod {
	return ium.Only[tag.UpdateTag](true)
}

func Returning(clauses ...string) psql.UpdateMod {
	return ium.Returning[tag.UpdateTag](clauses...)
}

func Set(column string, arg any) psql.UpdateMod {
	return ium.Set[tag.UpdateTag](column, arg)
}

func SetAN(column string, argumentName string) psql.UpdateMod {
	return ium.SetAN[tag.UpdateTag](column, argumentName)
}

func SetE(column string, value litsql.Expression) psql.UpdateMod {
	return ium.SetE[tag.UpdateTag](column, value)
}

func SetQ(column string, q psql.SelectQuery) psql.UpdateMod {
	return ium.SetQ[tag.UpdateTag, tag.SelectTag](column, q)
}

func SetS(column string, right string) psql.UpdateMod {
	return ium.SetS[tag.UpdateTag](column, right)
}

func SetR(raw string) psql.UpdateMod {
	return ium.SetR[tag.UpdateTag](raw)
}

func SetRE(assignment litsql.Expression) psql.UpdateMod {
	return ium.SetRE[tag.UpdateTag](assignment)
}

func Table(name string) psql.UpdateMod {
	return ium.Table[tag.UpdateTag](name)
}

func Where(condition string) psql.UpdateMod {
	return ium.Where[tag.UpdateTag](condition)
}

func WhereE(condition litsql.Expression) psql.UpdateMod {
	return ium.WhereE[tag.UpdateTag](condition)
}

func WhereC(query string, args ...any) psql.UpdateMod {
	return ium.WhereC[tag.UpdateTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return ium.With[tag.UpdateTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return ium.WithE[tag.UpdateTag](name, columns...)
}
