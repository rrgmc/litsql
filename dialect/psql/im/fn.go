package im

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/iim"
)

func Apply(f func(a psql.InsertModApply)) psql.InsertMod {
	return iim.Apply(f)
}

func Into(name string, columns ...string) psql.InsertMod {
	return iim.Into[tag.InsertTag](name, columns...)
}

func OverridingSystem() psql.InsertMod {
	return iim.OverridingSystem[tag.InsertTag]()
}

func OverridingUser() psql.InsertMod {
	return iim.OverridingUser[tag.InsertTag]()
}

// Insert from a query
func Query(q psql.SelectQuery) psql.InsertMod {
	return iim.Query[tag.InsertTag, tag.SelectTag](q)
}

func Returning(clauses ...string) psql.InsertMod {
	return iim.Returning[tag.InsertTag](clauses...)
}

func Values(values ...any) psql.InsertMod {
	return iim.Values[tag.InsertTag](values...)
}

func ValuesAN(argumentNames ...string) psql.InsertMod {
	return iim.ValuesAN[tag.InsertTag](argumentNames...)
}

func ValuesE(clauses ...litsql.Expression) psql.InsertMod {
	return iim.ValuesE[tag.InsertTag](clauses...)
}

func ValuesS(clauses ...string) psql.InsertMod {
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

func ConflictSetE(column string, value litsql.Expression) InsertConflictUpdateMod {
	return iim.ConflictSetE[tag.InsertTag](column, value)
}

func ConflictSetQ(column string, q psql.SelectQuery) InsertConflictUpdateMod {
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
