package main

import (
	"local/wall-crawl/core"
	"log"
)

func main() {
	pop := core.NewPopulation(10)

	for !pop.HasConverged() {
		pop.Evolve()

		pop.CalculateFitness()

		log.Printf("")
	}

	log.Println()
	log.Println("Simulation complete!")
	log.Printf("Fitness %v attained after %v generations\n", pop.Fitness, pop.Generation)
}
