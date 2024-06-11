package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

//litsql:dialects mysql
func OnDuplicateKeySet[T any](column string, arg any) sq.QueryMod[T] {
	return OnDuplicateKeySetExprClause[T](expr.JoinSep(" = ", expr.String(column), expr.Arg(arg)))
}

//litsql:dialects mysql
func OnDuplicateKeySetArgNamed[T any](column string, argumentName string) sq.QueryMod[T] {
	return OnDuplicateKeySetExprClause[T](expr.JoinSep(" = ", expr.String(column), expr.ArgNamed(argumentName)))
}

//litsql:dialects mysql
func OnDuplicateKeySetExpr[T any](column string, value litsql.Expression) sq.QueryMod[T] {
	return OnDuplicateKeySetExprClause[T](expr.JoinSep(" = ", expr.String(column), value))
}

//litsql:dialects mysql
func OnDuplicateKeySetQuery[T, A any](column string, q isq.Query[A]) sq.QueryMod[T] {
	return OnDuplicateKeySetExpr[T](column, q)
}

//litsql:dialects mysql
func OnDuplicateKeySetString[T any](column string, right string) sq.QueryMod[T] {
	return OnDuplicateKeySetExpr[T](column, expr.String(right))
}

//litsql:dialects mysql
func OnDuplicateKeySetClause[T any](query string, args ...any) sq.QueryMod[T] {
	return OnDuplicateKeySetExprClause[T](expr.Clause(query, args...))
}

//litsql:dialects mysql
func OnDuplicateKeySetExprClause[T any](assignment litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.InsertDuplicateKey{
			Set: iclause.Set{
				Set: []litsql.Expression{assignment},
			},
		})
	})
}
