package main

import (
	"local/wall-crawl/core"
	"log"
)

func main() {
	pop := core.NewPopulation(10)

	pop.CalculateFitness()

	for !pop.HasConverged() {
		pop.Evolve()

		pop.CalculateFitness()

		log.Printf("Generation %v - Fitness %v\n", pop.Generation, pop.Fitness)
	}

	log.Println()
	log.Println("Simulation complete!")
	log.Printf("Fitness %v attained after %v generations\n", pop.Fitness, pop.Generation)
}
