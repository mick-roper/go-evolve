package main

import (
	"math/rand"
	"time"
)

type individual struct {
	genes []int
}

func newIndividual(genes int) *individual {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	g := make([]int, genes)

	for i := range g {
		g[i] = r.Int()
	}

	return &individual{g}
}
