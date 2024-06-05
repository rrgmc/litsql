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

func ValuesE(clauses ...litsql.Expression) mysql.InsertMod {
	return iim.ValuesE[tag.InsertTag](clauses...)
}

func ValuesS(clauses ...string) mysql.InsertMod {
	return iim.ValuesS[tag.InsertTag](clauses...)
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

func ConflictSetAN(column string, argumentName string) InsertConflictUpdateMod {
	return iim.ConflictSetAN[tag.InsertTag](column, argumentName)
}

func ConflictSetE(column string, value litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetE[tag.InsertTag](column, value)
}

func ConflictSetQ(column string, q mysql.SelectQuery) InsertConflictUpdateMod {
	return iim.ConflictSetQ[tag.InsertTag, tag.SelectTag](column, q)
}

func ConflictSetS(column string, right string) InsertConflictUpdateMod {
	return iim.ConflictSetS[tag.InsertTag](column, right)
}

func ConflictSetR(raw string) InsertConflictUpdateMod {
	return iim.ConflictSetR[tag.InsertTag](raw)
}

func ConflictSetRE(assignment litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetRE[tag.InsertTag](assignment)
}

func ConflictWhere(condition string) InsertConflictUpdateMod {
	return iim.Where[tag.InsertTag](condition)
}

func ConflictWhereE(condition litsql.Expression) InsertConflictUpdateMod {
	return iim.WhereE[tag.InsertTag](condition)
}

func ConflictWhereC(query string, args ...any) InsertConflictUpdateMod {
	return iim.WhereC[tag.InsertTag](query, args...)
}

func With(name string, columns ...string) WithChain {
	return iim.With[tag.InsertTag](name, columns...)
}

func WithE(name string, columns ...litsql.Expression) WithChain {
	return iim.WithE[tag.InsertTag](name, columns...)
}
