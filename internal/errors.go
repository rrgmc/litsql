package internal

import (
	"fmt"

	"github.com/rrgmc/litsql"
)

func NewClauseError(f string, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: %s", litsql.ErrClause, f)
	}
	return fmt.Errorf("%w: %s", litsql.ErrClause, fmt.Sprintf(f, args...))
}

func NewClauseErrorWrap(err error) error {
	return fmt.Errorf("%w: %w", litsql.ErrClause, err)
}
