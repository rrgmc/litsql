package ism

import (
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

func Func[T any](f func(a sq.QueryModApply[T])) sq.QueryMod[T] {
	return isq.Func[T](f)
}
