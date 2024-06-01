package litsql

// Argument is the base interface for query arguments.
type Argument interface {
	isArgument()
}

// NamedArgument represents an argument were its value will be provided by name.
type NamedArgument interface {
	Argument
	Name() string
}

// ValuedArgument represents an argument were its value will be provided by this instance.
type ValuedArgument interface {
	Argument
	Value() (any, error)
}

// DBNamedArgument is like NamedArgument but its value will be wrapped using [sql.Named].
type DBNamedArgument interface {
	Argument
	DBName() string
}

// helpers

type ArgumentBase struct{}

func (a ArgumentBase) isArgument() {}
