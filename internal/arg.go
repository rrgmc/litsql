package internal

import "github.com/rrgmc/litsql"

type NamedArgument struct {
	litsql.ArgumentBase
	ArgName string
}

func (a NamedArgument) Name() string {
	return a.ArgName
}

type NamedArgumentWithDefault struct {
	litsql.ArgumentBase
	ArgName      string
	DefaultValue any
}

func (a NamedArgumentWithDefault) Name() string {
	return a.ArgName
}

func (a NamedArgumentWithDefault) Value() (any, error) {
	return a.DefaultValue, nil
}

type DBNamedArgument struct {
	litsql.ArgumentBase
	ArgName string
}

func (a DBNamedArgument) DBName() string {
	return a.ArgName
}

type DBNamedArgumentWithDefault struct {
	litsql.ArgumentBase
	ArgName      string
	DefaultValue any
}

func (a DBNamedArgumentWithDefault) DBName() string {
	return a.ArgName
}

func (a DBNamedArgumentWithDefault) Value() (any, error) {
	return a.DefaultValue, nil
}

type FuncArgument struct {
	litsql.ArgumentBase
	FN func() (any, error)
}

func (f FuncArgument) Value() (any, error) {
	return f.FN()
}
