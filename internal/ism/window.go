package ism

import (
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
)

func Window[T, CHAIN any](name string) *ichain.WindowChain[T, CHAIN] {
	w := &iclause.Windows{
		Windows: []*iclause.NamedWindow{{Name: name}},
	}

	return ichain.NewWindowChain[T, CHAIN](&ichain.WindowChain[T, CHAIN]{
		Windows:     w,
		NamedWindow: w.Windows[0],
	})
}
