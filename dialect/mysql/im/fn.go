// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package im

import (
	litsql "github.com/rrgmc/litsql"
	mysql "github.com/rrgmc/litsql/dialect/mysql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	iim "github.com/rrgmc/litsql/internal/iim"
)

func Apply(f func(a mysql.InsertModApply)) mysql.InsertMod {
	return iim.Apply[tag.InsertTag](f)
}

func Into(name string, columns ...string) mysql.InsertMod {
	return iim.Into[tag.InsertTag](name, columns...)
}

func IntoAs(name string, alias string, columns ...string) mysql.InsertMod {
	return iim.IntoAs[tag.InsertTag](name, alias, columns...)
}

func OnDuplicateKeySet(column string, arg any) mysql.InsertMod {
	return iim.OnDuplicateKeySet[tag.InsertTag](column, arg)
}

func OnDuplicateKeySetArgNamed(column string, argumentName string) mysql.InsertMod {
	return iim.OnDuplicateKeySetArgNamed[tag.InsertTag](column, argumentName)
}

func OnDuplicateKeySetClause(query string, args ...any) mysql.InsertMod {
	return iim.OnDuplicateKeySetClause[tag.InsertTag](query, args...)
}

func OnDuplicateKeySetExpr(column string, value litsql.Expression) mysql.InsertMod {
	return iim.OnDuplicateKeySetExpr[tag.InsertTag](column, value)
}

func OnDuplicateKeySetExprClause(assignment litsql.Expression) mysql.InsertMod {
	return iim.OnDuplicateKeySetExprClause[tag.InsertTag](assignment)
}

func OnDuplicateKeySetQuery(column string, q mysql.SelectQuery) mysql.InsertMod {
	return iim.OnDuplicateKeySetQuery[tag.InsertTag, tag.SelectTag](column, q)
}

func OnDuplicateKeySetString(column string, right string) mysql.InsertMod {
	return iim.OnDuplicateKeySetString[tag.InsertTag](column, right)
}

func Query(q mysql.SelectQuery) mysql.InsertMod {
	return iim.Query[tag.InsertTag, tag.SelectTag](q)
}

func Returning(clauses ...string) mysql.InsertMod {
	return iim.Returning[tag.InsertTag](clauses...)
}

func Values(values ...any) mysql.InsertMod {
	return iim.Values[tag.InsertTag](values...)
}

func ValuesArgNamed(argumentNames ...string) mysql.InsertMod {
	return iim.ValuesArgNamed[tag.InsertTag](argumentNames...)
}

func ValuesExpr(clauses ...litsql.Expression) mysql.InsertMod {
	return iim.ValuesExpr[tag.InsertTag](clauses...)
}

func ValuesString(clauses ...string) mysql.InsertMod {
	return iim.ValuesString[tag.InsertTag](clauses...)
}

func With(name string, columns ...string) WithChain {
	return iim.With[tag.InsertTag, WithChain](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return iim.WithExpr[tag.InsertTag, WithChain](name, columns...)
}
