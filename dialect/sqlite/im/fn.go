package im

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/sqlite"
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/iim"
)

func Func(f func(a sqlite.InsertModApply)) sqlite.InsertMod {
	return iim.Func(f)
}

func Into(name string, columns ...string) sqlite.InsertMod {
	return iim.Into[tag.InsertTag](name, columns...)
}

func OverridingSystem() sqlite.InsertMod {
	return iim.OverridingSystem[tag.InsertTag]()
}

func OverridingUser() sqlite.InsertMod {
	return iim.OverridingUser[tag.InsertTag]()
}

// Insert from a query
func Query(q sqlite.SelectQuery) sqlite.InsertMod {
	return iim.Query[tag.InsertTag, tag.SelectTag](q)
}

func Returning(clauses ...string) sqlite.InsertMod {
	return iim.Returning[tag.InsertTag](clauses...)
}

func Values(values ...any) sqlite.InsertMod {
	return iim.Values[tag.InsertTag](values...)
}

func ValuesE(clauses ...litsql.Expression) sqlite.InsertMod {
	return iim.ValuesE[tag.InsertTag](clauses...)
}

func ValuesS(clauses ...string) sqlite.InsertMod {
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

func ConflictSetQ(column string, q sqlite.SelectQuery) InsertConflictUpdateMod {
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
