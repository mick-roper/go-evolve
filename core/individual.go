package core

import (
	"math/rand"
)

type chromosome []int

type individual struct {
	chromosome chromosome
	genes      int
	fitness    int
}

func newIndividual(genes int, r *rand.Rand) *individual {
	var g chromosome = make([]int, genes)

	for i := range g {
		g[i] = r.Int() % 2
	}

	return &individual{g, genes, 0}
}

func (i *individual) calculateFitness() {
	i.fitness = 0

	for _, n := range i.chromosome {
		i.fitness += n
	}
}
