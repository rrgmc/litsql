package im

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/sqlite"
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/iim"
)

func Apply(f func(a sqlite.InsertModApply)) sqlite.InsertMod {
	return iim.Apply(f)
}

func ConflictSet(column string, arg any) InsertConflictUpdateMod {
	return iim.ConflictSet[tag.InsertTag](column, arg)
}

func ConflictSetArgNamed(column string, argumentName string) InsertConflictUpdateMod {
	return iim.ConflictSetArgNamed[tag.InsertTag](column, argumentName)
}

func ConflictSetClause(query string, args ...any) InsertConflictUpdateMod {
	return iim.ConflictSetClause[tag.InsertTag](query, args...)
}

func ConflictSetExpr(column string, value litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetExpr[tag.InsertTag](column, value)
}

func ConflictSetExprClause(assignment litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetExprClause[tag.InsertTag](assignment)
}

func ConflictSetQuery(column string, q sqlite.SelectQuery) InsertConflictUpdateMod {
	return iim.ConflictSetQuery[tag.InsertTag, tag.SelectTag](column, q)
}

func ConflictSetString(column string, right string) InsertConflictUpdateMod {
	return iim.ConflictSetString[tag.InsertTag](column, right)
}

func ConflictWhere(condition string) InsertConflictUpdateMod {
	return iim.ConflictWhere[tag.InsertTag](condition)
}

func ConflictWhereClause(query string, args ...any) InsertConflictUpdateMod {
	return iim.ConflictWhereClause[tag.InsertTag](query, args...)
}

func ConflictWhereExpr(condition litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictWhereExpr[tag.InsertTag](condition)
}

func Into(name string, columns ...string) sqlite.InsertMod {
	return iim.Into[tag.InsertTag](name, columns...)
}

func OnConflict(columns ...string) InsertConflictChain {
	return iim.OnConflict[tag.InsertTag](columns...)
}

func OnConflictOnConstraint(constraint string) InsertConflictChain {
	return iim.OnConflictOnConstraint[tag.InsertTag](constraint)
}

func OverridingSystem() sqlite.InsertMod {
	return iim.OverridingSystem[tag.InsertTag]()
}

func OverridingUser() sqlite.InsertMod {
	return iim.OverridingUser[tag.InsertTag]()
}

func Query(q sqlite.SelectQuery) sqlite.InsertMod {
	return iim.Query[tag.InsertTag, tag.SelectTag](q)
}

func Returning(clauses ...string) sqlite.InsertMod {
	return iim.Returning[tag.InsertTag](clauses...)
}

func Values(values ...any) sqlite.InsertMod {
	return iim.Values[tag.InsertTag](values...)
}

func ValuesExpr(clauses ...litsql.Expression) sqlite.InsertMod {
	return iim.ValuesExpr[tag.InsertTag](clauses...)
}

func ValuesString(clauses ...string) sqlite.InsertMod {
	return iim.ValuesString[tag.InsertTag](clauses...)
}

func With(name string, columns ...string) WithChain {
	return iim.With[tag.InsertTag](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return iim.WithExpr[tag.InsertTag](name, columns...)
}
