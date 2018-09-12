package core

import (
	"math"
	"math/rand"
	"time"
)

// Population that we are going to evolve
type Population struct {
	r               *rand.Rand
	individuals     []*individual
	previousFitness int
	Fitness         int
	Generation      int
}

// NewPopulation creates a new population
func NewPopulation(size int) *Population {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	i := make([]*individual, size)
	genes := 5

	for x := range i {
		i[x] = newIndividual(genes, r)
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
		in := p.individuals[i]
		in.calculateFitness()
		p.Fitness += in.fitness
	}
}

// HasConverged checks if the population has converged
func (p *Population) HasConverged() bool {
	return p.previousFitness == p.Fitness
}

// Evolve the population
func (p *Population) Evolve() {
	// select
	i1 := p.getClosestToFitness(math.MaxInt32)                 // strongest
	i2 := p.getClosestToFitness(p.individuals[i1].fitness - 1) // second strongest
	i3 := p.getClosestToFitness(-1)                            // weakest

	// combine
	offspring := p.combine(p.individuals[i1], p.individuals[i2])

	// mutate
	if p.r.Int31n(10) < 3 {
		p.mutate(offspring)
	}

	p.individuals[i3] = offspring
}

func (p *Population) getClosestToFitness(f int) int {
	return -1
}

func (p *Population) combine(i1, i2 *individual) *individual {
	genes := i1.genes
	xp := p.r.Intn(genes)

	chrom := make(chromosome, genes)

	for n := 0; n < genes; n++ {
		var t int
		if n < xp {
			t = i1.chromosome[n]
		} else {
			t = i2.chromosome[n]
		}

		chrom[n] = t
	}

	newI := &individual{
		chromosome: chrom,
		genes:      genes,
		fitness:    0,
	}

	newI.calculateFitness()

	return newI
}

func (p *Population) mutate(i *individual) {
	mPoint := p.r.Intn(i.genes)

	for n := 0; n < mPoint; n++ {
		if i.chromosome[n] == 1 {
			i.chromosome[n] = 0
		} else {
			i.chromosome[n] = 1
		}
	}
}
