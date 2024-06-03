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

func InnerJoinE(table litsql.Expression) JoinChain {
	return idm.InnerJoinE[tag.DeleteTag](table)
}

func LeftJoin(table string) JoinChain {
	return idm.LeftJoin[tag.DeleteTag](table)
}

func LeftJoinE(table litsql.Expression) JoinChain {
	return idm.LeftJoinE[tag.DeleteTag](table)
}

func RightJoin(table string) JoinChain {
	return idm.RightJoin[tag.DeleteTag](table)
}

func RightJoinE(table litsql.Expression) JoinChain {
	return idm.RightJoinE[tag.DeleteTag](table)
}

func FullJoin(table string) JoinChain {
	return idm.FullJoin[tag.DeleteTag](table)
}

func FullJoinE(table litsql.Expression) JoinChain {
	return idm.FullJoinE[tag.DeleteTag](table)
}

func CrossJoin(table string) JoinChain {
	return idm.CrossJoin[tag.DeleteTag](table)
}

func CrossJoinE(table litsql.Expression) JoinChain {
	return idm.CrossJoinE[tag.DeleteTag](table)
}

func StraightJoin(table string) JoinChain {
	return idm.StraightJoin[tag.DeleteTag](table)
}

func StraightJoinE(table litsql.Expression) JoinChain {
	return idm.StraightJoinE[tag.DeleteTag](table)
}

func Only() mysql.DeleteMod {
	return idm.Only[tag.DeleteTag](true)
}

func Returning(clauses ...string) mysql.DeleteMod {
	return idm.Returning[tag.DeleteTag](clauses...)
}

func Using(table string) FromChain {
	return idm.Using[tag.DeleteTag](table)
}

func UsingE(table litsql.Expression) FromChain {
	return idm.UsingE[tag.DeleteTag](table)
}

func UsingQ(q mysql.SelectQuery) FromChain {
	return idm.UsingQ[tag.DeleteTag, tag.SelectTag](q)
}

func Where(condition string) mysql.DeleteMod {
	return idm.Where[tag.DeleteTag](condition)
}

func WhereE(condition litsql.Expression) mysql.DeleteMod {
	return idm.WhereE[tag.DeleteTag](condition)
}

func WhereC(query string, args ...any) mysql.DeleteMod {
	return idm.WhereC[tag.DeleteTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return idm.With[tag.DeleteTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return idm.WithE[tag.DeleteTag](name, columns...)
}
