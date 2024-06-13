package litsql

import (
	"errors"
)

var (
	ErrClause    = errors.New("clause error")
	ErrParseArgs = errors.New("parse args error")
)
