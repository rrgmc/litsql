package expr

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/rrgmc/litsql"
)

// C parses the query, replacing any "?" with the dialect-specific argument matching, and wraps the passed arguments.
func C(query string, args ...any) litsql.Expression {
	return &clause{
		query: query,
		args:  args,
	}
}

type clause struct {
	query string // The clause with ? used for placeholders
	args  []any  // The replacements for the placeholders in order
}

func (r clause) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	// replace the args with positional args appropriately
	total, args, err := r.convertQuestionMarks(w, d, start)
	if err != nil {
		return nil, err
	}

	if len(r.args) != total {
		return r.args, &clauseError{args: len(r.args), placeholders: total, clause: r.query}
	}

	return args, nil
}

// convertQuestionMarks converts each occurrence of ? with $<number>
// where <number> is an incrementing digit starting at startAt.
// If question-mark (?) is escaped using back-slash (\), it will be ignored.
func (r clause) convertQuestionMarks(w litsql.Writer, d litsql.Dialect, startAt int) (int, []any, error) {
	if startAt == 0 {
		panic("Not a valid start number.")
	}

	paramIndex := 0
	total := 0
	var args []any

	clause := r.query
	for {
		if paramIndex >= len(clause) {
			break
		}

		clause = clause[paramIndex:]
		paramIndex = strings.IndexByte(clause, '?')

		if paramIndex == -1 {
			w.Write(clause)
			break
		}

		escapeIndex := strings.Index(clause, `\?`)
		if escapeIndex != -1 && paramIndex > escapeIndex {
			w.Write(clause[:escapeIndex] + "?")
			paramIndex++
			continue
		}

		w.Write(clause[:paramIndex])

		var arg any
		if total < len(r.args) {
			arg = r.args[total]
		}

		newArgs, err := clauseWriteArg(w, d, startAt, arg)
		if err != nil {
			return total, nil, err
		}

		startAt += len(newArgs)
		args = append(args, newArgs...)

		total++
		paramIndex++
	}

	return total, args, nil
}

func clauseWriteArg(w litsql.Writer, d litsql.Dialect, startAt int, arg any) (args []any, err error) {
	if ex, ok := arg.(litsql.Expression); ok {
		// inner [litsql.Expression]
		eargs, err := ex.WriteSQL(w, d, startAt)
		if err != nil {
			return nil, err
		}
		args = append(args, eargs...)
	} else if nv, ok := arg.(litsql.DBNamedArgument); ok {
		// DB-specific named argument
		dn, ok := d.(litsql.DialectWithNamed)
		if !ok {
			return nil, ErrNoNamedArgs
		}
		dn.WriteNamedArg(w, nv.DBName())
		args = append(args, arg)
	} else if nv, ok := arg.(sql.NamedArg); ok {
		// sql.NamedArg
		dn, ok := d.(litsql.DialectWithNamed)
		if !ok {
			return nil, ErrNoNamedArgs
		}
		dn.WriteNamedArg(w, nv.Name)
		args = append(args, arg)
	} else {
		// dialect argument
		d.WriteArg(w, startAt)
		args = append(args, arg)
	}
	return
}

type clauseError struct {
	args         int
	placeholders int
	clause       string
}

func (s *clauseError) Error() string {
	return fmt.Sprintf(
		"Bad Statement: has %d placeholders but %d args: %s",
		s.placeholders, s.args, s.clause,
	)
}

func (s *clauseError) Equal(I error) bool {
	var s2 *clauseError
	if errors.As(I, &s2) {
		return s2.args == s.args && s2.placeholders == s.placeholders
	}

	return false
}
