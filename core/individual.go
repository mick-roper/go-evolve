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

	for n := range i.chromosome {
		if i.chromosome[n] == 1 {
			i.fitness++
		}
	}
}
