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

func NewClauseErrorInvalidMerge(clauseName string) error {
	return NewClauseError("%s: invalid merge", clauseName)
}

func NewClauseErrorInvalidMergeHasChanges(clauseName string) error {
	return NewClauseError("%s: invalid merge: new instance has changes", clauseName)
}

func NewClauseErrorInvalidMergeCannotHaveMultiple(clauseName string) error {
	return NewClauseError("%s: invalid merge: clause cannot have multiple instances", clauseName)
}
