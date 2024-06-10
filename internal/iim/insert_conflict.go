package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq/chain"
	"github.com/rrgmc/litsql/sq/mod"
)

func OnConflict[T any](columns ...string) chain.InsertConflictUpdate[T, imod.InsertConflictUpdateModTag] {
	return &ichain.InsertConflictChain[T]{
		InsertConflictUpdate: &iclause.InsertConflictUpdate{
			Target: iclause.InsertConflictTarget{
				Columns: columns,
			},
			Set: iclause.Set{
				Starter: true,
			},
		},
	}
}

func OnConflictOnConstraint[T any](constraint string) chain.InsertConflictUpdate[T, imod.InsertConflictUpdateModTag] {
	return &ichain.InsertConflictChain[T]{
		InsertConflictUpdate: &iclause.InsertConflictUpdate{
			Target: iclause.InsertConflictTarget{
				Constraint: constraint,
			},
			Set: iclause.Set{
				Starter: true,
			},
		},
	}
}

func ConflictSet[T any](column string, arg any) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T](expr.JoinSep(" = ", expr.String(column), expr.Arg(arg)))
}

func ConflictSetArgNamed[T any](column string, argumentName string) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T](expr.JoinSep(" = ", expr.String(column), expr.ArgNamed(argumentName)))
}

func ConflictSetExpr[T any](column string, value litsql.Expression) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T](expr.JoinSep(" = ", expr.String(column), value))
}

func ConflictSetQuery[T, A any](column string, q isq.Query[A]) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExpr[T](column, q)
}

func ConflictSetString[T any](column string, right string) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExpr[T](column, expr.String(right))
}

func ConflictSetClause[T any](query string, args ...any) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictSetExprClause[T](expr.Clause(query, args...))
}

func ConflictSetExprClause[T any](assignment litsql.Expression) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return mod.InsertConflictUpdateModFunc[T, imod.InsertConflictUpdateModTag](func(a *iclause.InsertConflictUpdate) {
		a.Set.Set = append(a.Set.Set, assignment)
	})
}

func ConflictWhere[T any](condition string) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictWhereExpr[T](expr.String(condition))
}

func ConflictWhereExpr[T any](condition litsql.Expression) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return mod.InsertConflictUpdateModFunc[T, imod.InsertConflictUpdateModTag](func(a *iclause.InsertConflictUpdate) {
		a.Where.Conditions = append(a.Where.Conditions, condition)
	})
}

func ConflictWhereClause[T any](query string, args ...any) mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag] {
	return ConflictWhereExpr[T](expr.Clause(query, args...))
}
