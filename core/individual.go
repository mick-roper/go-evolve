package core

import (
	"math/rand"
	"time"
)

type gene int

type chromosome []gene

type individual struct {
	chromosome chromosome
}

func newIndividual(genes int) *individual {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	var g chromosome = make([]gene, genes)

	for i := range g {
		g[i] = gene(r.Int())
	}

	return &individual{g}
}
