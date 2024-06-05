package expr

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// Arg outputs the dialect-specific argument matcher, and wraps the passed value.
func Arg(value any) litsql.Expression {
	return arg{value: value}
}

// ArgNamed outputs the dialect-specific argument matcher, and wraps the passed value as a named argument.
func ArgNamed(name string) litsql.Expression {
	return Arg(internal.NamedArgument{ArgName: name})
}

// ArgDBNamed outputs the dialect-specific argument matcher, and wraps the passed value as a named argument.
func ArgDBNamed(name string) litsql.Expression {
	return Arg(internal.DBNamedArgument{ArgName: name})
}

// ArgFunc returns the argument value in a callback.
func ArgFunc(f func() (any, error)) litsql.Expression {
	return Arg(internal.FuncArgument{FN: f})
}

// Args wraps multiple values in a slice of Arg.
func Args(values []any) []litsql.Expression {
	return MapSlice(values, func(a any) litsql.Expression {
		return Arg(a)
	})
}

// ArgsNamed wraps multiple values in a slice of ArgNamed.
func ArgsNamed(names ...string) []litsql.Expression {
	return MapSlice(names, func(a string) litsql.Expression {
		return ArgNamed(a)
	})
}

// ArgsDBNamed wraps multiple values in a slice of ArgDBNamed.
func ArgsDBNamed(names ...string) []litsql.Expression {
	return MapSlice(names, func(a string) litsql.Expression {
		return ArgDBNamed(a)
	})
}

// ArgsFunc wraps multiple values in a slice of ArgFunc.
func ArgsFunc(fs []func() (any, error)) []litsql.Expression {
	return MapSlice(fs, func(f func() (any, error)) litsql.Expression {
		return ArgFunc(f)
	})
}

// In outputs the list of values as Arg separated by commas.
func In(values []any) litsql.Expression {
	return argList{
		values:    values,
		separator: internal.CommaSpace,
	}
}

// InT outputs the list of values as Arg separated by commas.
func InT[T any](values ...T) litsql.Expression {
	return argList{
		values:    internal.ToAnySlice(values),
		separator: internal.CommaSpace,
	}
}

// InP outputs the list of values as Arg separated by commas, wrapped in parentheses.
func InP(values []any) litsql.Expression {
	return argList{
		values:    values,
		separator: internal.CommaSpace,
		prefix:    internal.OpenPar,
		suffix:    internal.ClosePar,
	}
}

// InPT outputs the list of values as Arg separated by commas, wrapped in parentheses.
func InPT[T any](values ...T) litsql.Expression {
	return argList{
		values:    internal.ToAnySlice(values),
		separator: internal.CommaSpace,
		prefix:    internal.OpenPar,
		suffix:    internal.ClosePar,
	}
}

type arg struct {
	value any
}

func (a arg) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	return clauseWriteArg(w, d, start, a.value)
}

type argList struct {
	values    []any
	separator string
	prefix    string
	suffix    string
}

func (a argList) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	w.Write(a.prefix)

	startAt := start
	for i, v := range a.values {
		if i > 0 {
			w.Write(a.separator)
		}
		newArgs, err := clauseWriteArg(w, d, startAt, v)
		if err != nil {
			return nil, err
		}
		startAt += len(newArgs)
		args = append(args, newArgs...)
	}

	w.Write(a.suffix)

	return args, nil
}
