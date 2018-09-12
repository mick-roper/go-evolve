package core

import (
	"log"
	"math/rand"
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
	return nil
}

// CalculateFitness of the population
func (p *Population) CalculateFitness() {

}

// HasConverged checks if the population has converged
func (p *Population) HasConverged() bool {
	return p.previousFitness == p.Fitness
}

// Evolve the population
func (p *Population) Evolve() {
	log.Panic("not implemented")
}
