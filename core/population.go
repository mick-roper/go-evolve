package core

import (
	"log"
	"math/rand"
	"time"
)

// Population that we are going to evolve
type Population struct {
	r               *rand.Rand
	individuals     []individual
	previousFitness int
	Fitness         int
	Generation      int
}

// NewPopulation creates a new population
func NewPopulation(size int) *Population {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	i := make([]individual, size)
	genes := 5

	for x := range i {
		i[x] = *newIndividual(genes, r)
	}

	return &Population{
		r:               r,
		individuals:     i,
		previousFitness: -1,
		Fitness:         -1,
		Generation:      0,
	}
}

// CalculateFitness of the population
func (p *Population) CalculateFitness() {
	p.previousFitness = p.Fitness
	p.Fitness = 0

	for i := range p.individuals {
		p.Fitness += p.individuals[i].calculateFitness()
	}
}

// HasConverged checks if the population has converged
func (p *Population) HasConverged() bool {
	return p.previousFitness == p.Fitness
}

// Evolve the population
func (p *Population) Evolve() {
	log.Panic("not implemented")
}
