package core

import (
	"math"
	"math/rand"
	"sort"
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
	// select fittest
	sortedIndividuals := p.individuals[:]
	sort.Slice(sortedIndividuals, func(i, j int) bool {
		return sortedIndividuals[i].fitness > sortedIndividuals[j].fitness
	})

	i1 := sortedIndividuals[0]
	i2 := sortedIndividuals[1]

	// combine
	offspring := p.combine(i1, i2)

	// mutate
	if p.r.Int31n(10) < 3 {
		p.mutate(offspring)
	}

	// replace weakest
	weakestIx := p.getIndexOfWeakest()
	p.individuals[weakestIx] = offspring

	// increment the generation
	p.Generation++
}

func (p *Population) getIndexOfWeakest() int {
	x := math.MaxInt32
	ix := -1

	for a, b := range p.individuals {
		if b.fitness < x {
			ix = a
		}
	}

	return ix
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
