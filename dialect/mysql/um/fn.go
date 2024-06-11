// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package um

import (
	litsql "github.com/rrgmc/litsql"
	mysql "github.com/rrgmc/litsql/dialect/mysql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	ium "github.com/rrgmc/litsql/internal/ium"
)

func Apply(f func(a mysql.UpdateModApply)) mysql.UpdateMod {
	return ium.Apply[tag.UpdateTag](f)
}

func CrossJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ium.CrossJoin[tag.UpdateTag](table),
	}
}

func CrossJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ium.CrossJoinExpr[tag.UpdateTag](table),
	}
}

func From(table string) FromChain {
	return &fromChainAdapter{
		chain: ium.From[tag.UpdateTag](table),
	}
}

func FromExpr(table litsql.Expression) FromChain {
	return &fromChainAdapter{
		chain: ium.FromExpr[tag.UpdateTag](table),
	}
}

func FromQuery(q mysql.SelectQuery) FromChain {
	return &fromChainAdapter{
		chain: ium.FromQuery[tag.UpdateTag, tag.SelectTag](q),
	}
}

func FullJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ium.FullJoin[tag.UpdateTag](table),
	}
}

func FullJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ium.FullJoinExpr[tag.UpdateTag](table),
	}
}

func InnerJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ium.InnerJoin[tag.UpdateTag](table),
	}
}

func InnerJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ium.InnerJoinExpr[tag.UpdateTag](table),
	}
}

func LeftJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ium.LeftJoin[tag.UpdateTag](table),
	}
}

func LeftJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ium.LeftJoinExpr[tag.UpdateTag](table),
	}
}

func Only() mysql.UpdateMod {
	return ium.Only[tag.UpdateTag]()
}

func Returning(clauses ...string) mysql.UpdateMod {
	return ium.Returning[tag.UpdateTag](clauses...)
}

func RightJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ium.RightJoin[tag.UpdateTag](table),
	}
}

func RightJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ium.RightJoinExpr[tag.UpdateTag](table),
	}
}

func Set(column string, arg any) mysql.UpdateMod {
	return ium.Set[tag.UpdateTag](column, arg)
}

func SetArgNamed(column string, argumentName string) mysql.UpdateMod {
	return ium.SetArgNamed[tag.UpdateTag](column, argumentName)
}

func SetClause(query string, args ...any) mysql.UpdateMod {
	return ium.SetClause[tag.UpdateTag](query, args...)
}

func SetExpr(column string, value litsql.Expression) mysql.UpdateMod {
	return ium.SetExpr[tag.UpdateTag](column, value)
}

func SetExprClause(assignment litsql.Expression) mysql.UpdateMod {
	return ium.SetExprClause[tag.UpdateTag](assignment)
}

func SetQuery(column string, q mysql.SelectQuery) mysql.UpdateMod {
	return ium.SetQuery[tag.UpdateTag, tag.SelectTag](column, q)
}

func SetString(column string, right string) mysql.UpdateMod {
	return ium.SetString[tag.UpdateTag](column, right)
}

func StraightJoin(table string) JoinChain {
	return &joinChainAdapter{
		chain: ium.StraightJoin[tag.UpdateTag](table),
	}
}

func StraightJoinExpr(table litsql.Expression) JoinChain {
	return &joinChainAdapter{
		chain: ium.StraightJoinExpr[tag.UpdateTag](table),
	}
}

func Table(name string) mysql.UpdateMod {
	return ium.Table[tag.UpdateTag](name)
}

func TableAs(name string, alias string) mysql.UpdateMod {
	return ium.TableAs[tag.UpdateTag](name, alias)
}

func Where(condition string) mysql.UpdateMod {
	return ium.Where[tag.UpdateTag](condition)
}

func WhereClause(query string, args ...any) mysql.UpdateMod {
	return ium.WhereClause[tag.UpdateTag](query, args...)
}

func WhereExpr(condition litsql.Expression) mysql.UpdateMod {
	return ium.WhereExpr[tag.UpdateTag](condition)
}

func With(name string, columns ...string) WithChain {
	return ium.With[tag.UpdateTag](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return ium.WithExpr[tag.UpdateTag](name, columns...)
}
