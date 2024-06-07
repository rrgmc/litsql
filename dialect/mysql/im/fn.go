package im

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/iim"
)

func Apply(f func(a mysql.InsertModApply)) mysql.InsertMod {
	return iim.Apply(f)
}

func Into(name string, columns ...string) mysql.InsertMod {
	return iim.Into[tag.InsertTag](name, columns...)
}

func OverridingSystem() mysql.InsertMod {
	return iim.OverridingSystem[tag.InsertTag]()
}

func OverridingUser() mysql.InsertMod {
	return iim.OverridingUser[tag.InsertTag]()
}

// Insert from a query
func Query(q mysql.SelectQuery) mysql.InsertMod {
	return iim.Query[tag.InsertTag, tag.SelectTag](q)
}

func Returning(clauses ...string) mysql.InsertMod {
	return iim.Returning[tag.InsertTag](clauses...)
}

func Values(values ...any) mysql.InsertMod {
	return iim.Values[tag.InsertTag](values...)
}

func ValuesExpr(clauses ...litsql.Expression) mysql.InsertMod {
	return iim.ValuesExpr[tag.InsertTag](clauses...)
}

func ValuesString(clauses ...string) mysql.InsertMod {
	return iim.ValuesString[tag.InsertTag](clauses...)
}

func OnConflict(columns ...string) InsertConflictChain {
	return iim.OnConflict[tag.InsertTag](columns...)
}

func OnConflictOnConstraint(constraint string) InsertConflictChain {
	return iim.OnConflictOnConstraint[tag.InsertTag](constraint)
}

func ConflictSet(column string, arg any) InsertConflictUpdateMod {
	return iim.ConflictSet[tag.InsertTag](column, arg)
}

func ConflictSetArgNamed(column string, argumentName string) InsertConflictUpdateMod {
	return iim.ConflictSetArgNamed[tag.InsertTag](column, argumentName)
}

func ConflictSetExpr(column string, value litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetExpr[tag.InsertTag](column, value)
}

func ConflictSetQuery(column string, q mysql.SelectQuery) InsertConflictUpdateMod {
	return iim.ConflictSetQuery[tag.InsertTag, tag.SelectTag](column, q)
}

func ConflictSetString(column string, right string) InsertConflictUpdateMod {
	return iim.ConflictSetString[tag.InsertTag](column, right)
}

func ConflictSetClause(query string, args ...any) InsertConflictUpdateMod {
	return iim.ConflictSetClause[tag.InsertTag](query, args...)
}

func ConflictSetExprClause(assignment litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetExprClause[tag.InsertTag](assignment)
}

func ConflictWhere(condition string) InsertConflictUpdateMod {
	return iim.Where[tag.InsertTag](condition)
}

func ConflictWhereExpr(condition litsql.Expression) InsertConflictUpdateMod {
	return iim.WhereExpr[tag.InsertTag](condition)
}

func ConflictWhereClause(query string, args ...any) InsertConflictUpdateMod {
	return iim.WhereClause[tag.InsertTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return iim.With[tag.InsertTag](name, columns...)
}

func WithExpr(name string, columns ...litsql.Expression) WithChain {
	return iim.WithExpr[tag.InsertTag](name, columns...)
}
