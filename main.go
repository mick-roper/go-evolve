package main

import (
	"flag"
	"local/wall-crawl/core"
	"log"
)

func main() {
	popSize := flag.Int("popsize", 10, "The size of the population")
	genes := flag.Int("genes", 10, "The number of genes in each chromosome")

	flag.Parse()

	log.Printf("Creating population of %v individuals with %v genes in each chromosome\n", *popSize, *genes)

	pop := core.NewPopulation(*popSize, *genes)

	pop.CalculateFitness()

	for !pop.HasConverged() {
		pop.Evolve()

		pop.CalculateFitness()

		log.Printf("Generation %v - Fitness %v\n", pop.Generation, pop.Fitness)
	}

	log.Println()
	log.Println("Simulation complete!")
	log.Printf("Population converged after %v generations\n", pop.Generation)
}
