package tag

type AllTag interface {
	SelectTag | InsertTag | UpdateTag | DeleteTag
}

type SelectTag struct{}

type InsertTag struct{}

type UpdateTag struct{}

type DeleteTag struct{}
