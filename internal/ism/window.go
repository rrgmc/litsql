package ism

import (
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq/chain"
)

func Window[T any](name string) chain.Window[T] {
	w := &iclause.Windows{
		Windows: []*iclause.NamedWindow{{Name: name}},
	}
	return &ichain.WindowChain[T]{
		Windows:     w,
		NamedWindow: w.Windows[0],
	}
}
