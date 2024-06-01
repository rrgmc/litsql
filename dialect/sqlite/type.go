package sqlite

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/sq"
)

type SelectQuery = sq.Query[tag.SelectTag]
type InsertQuery = sq.Query[tag.InsertTag]
type UpdateQuery = sq.Query[tag.UpdateTag]
type DeleteQuery = sq.Query[tag.DeleteTag]

type SelectMod = sq.QueryMod[tag.SelectTag]
type InsertMod = sq.QueryMod[tag.InsertTag]
type UpdateMod = sq.QueryMod[tag.UpdateTag]
type DeleteMod = sq.QueryMod[tag.DeleteTag]

type SelectModApply = sq.QueryModApply[tag.SelectTag]
type InsertModApply = sq.QueryModApply[tag.InsertTag]
type UpdateModApply = sq.QueryModApply[tag.UpdateTag]
type DeleteModApply = sq.QueryModApply[tag.DeleteTag]
