package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq/mod"
)

//litsql:dialects psql,sqlite
func OnConflict[T, CHAIN any](columns ...string) *ichain.InsertConflictUpdateChain[T, CHAIN] {
	return ichain.NewInsertConflictUpdateChain[T, CHAIN](&ichain.InsertConflictUpdateChain[T, CHAIN]{
		InsertConflictUpdate: &iclause.InsertConflictUpdate{
			Target: iclause.InsertConflictTarget{
				Columns: columns,
			},
			Set: iclause.Set{
				Starter: true,
			},
		},
	})
}

//litsql:dialects psql,sqlite
func OnConflictOnConstraint[T, CHAIN any](constraint string) *ichain.InsertConflictUpdateChain[T, CHAIN] {
	return ichain.NewInsertConflictUpdateChain[T, CHAIN](&ichain.InsertConflictUpdateChain[T, CHAIN]{
		InsertConflictUpdate: &iclause.InsertConflictUpdate{
			Target: iclause.InsertConflictTarget{
				Constraint: constraint,
			},
			Set: iclause.Set{
				Starter: true,
			},
		},
	})
}

//litsql:dialects psql,sqlite
func ConflictSet[T, CHAIN any](column string, arg any) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T, CHAIN](expr.JoinSep(" = ", expr.String(column), expr.Arg(arg)))
}

//litsql:dialects psql,sqlite
func ConflictSetArgNamed[T, CHAIN any](column string, argumentName string) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T, CHAIN](expr.JoinSep(" = ", expr.String(column), expr.ArgNamed(argumentName)))
}

//litsql:dialects psql,sqlite
func ConflictSetExpr[T, CHAIN any](column string, value litsql.Expression) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T, CHAIN](expr.JoinSep(" = ", expr.String(column), value))
}

//litsql:dialects psql,sqlite
func ConflictSetQuery[T, CHAIN, A any](column string, q isq.Query[A]) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExpr[T, CHAIN](column, q)
}

//litsql:dialects psql,sqlite
func ConflictSetString[T, CHAIN any](column string, right string) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExpr[T, CHAIN](column, expr.String(right))
}

//litsql:dialects psql,sqlite
func ConflictSetClause[T, CHAIN any](query string, args ...any) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T, CHAIN](expr.Clause(query, args...))
}

//litsql:dialects psql,sqlite
func ConflictSetExprClause[T, CHAIN any](assignment litsql.Expression) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return mod.InsertConflictUpdateModFunc[T, imod.InsertConflictUpdateModTag](func(a *iclause.InsertConflictUpdate) {
		a.SetSet(assignment)
	})
}

//litsql:dialects psql,sqlite
func ConflictWhere[T, CHAIN any](condition string) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictWhereExpr[T, CHAIN](expr.String(condition))
}

//litsql:dialects psql,sqlite
func ConflictWhereExpr[T, CHAIN any](condition litsql.Expression) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return mod.InsertConflictUpdateModFunc[T, imod.InsertConflictUpdateModTag](func(a *iclause.InsertConflictUpdate) {
		a.Where.Conditions = append(a.Where.Conditions, condition)
	})
}

//litsql:dialects psql,sqlite
func ConflictWhereClause[T, CHAIN any](query string, args ...any) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictWhereExpr[T, CHAIN](expr.Clause(query, args...))
}
