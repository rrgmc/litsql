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

// ArgValues is the supplier of values for named arguments.
type ArgValues interface {
	Get(string) (any, bool, error)
}

// ArgValuesFunc is a functional implementation of ArgValues.
type ArgValuesFunc func(string) (any, bool, error)

func (f ArgValuesFunc) Get(s string) (any, bool, error) {
	return f(s)
}

// helpers

type ArgumentBase struct{}

func (a ArgumentBase) isArgument() {}
