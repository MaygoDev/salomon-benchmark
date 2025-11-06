package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Game struct {
	dimensions int
	cells      []Cell
	generation int

	lastPrint           time.Time
	lastGenerationCount int
}

func (g *Game) Populate(count int) {
	for i := 0; i < count; i++ {
		genome := make([]float64, g.dimensions)

		for j := 0; j < g.dimensions; j++ {
			genome[j] = rand.Float64()
		}

		g.cells = append(g.cells, Cell{
			genome: genome,
		})
	}
	g.lastPrint = time.Now()
}

func (g *Game) Generate(percent int8) {
	g.generation++
	for i := range g.cells {
		g.cells[i].Apply()
	}

	// Sort by lowest score
	sort.Sort(ByScore(g.cells))

	survivors := g.cells[:len(g.cells)*int(percent)/100]

	if time.Since(g.lastPrint) > time.Second {
		g.lastPrint = time.Now()

		generationsPerSecond := g.generation - g.lastGenerationCount
		fmt.Printf("Generation/s: %d | Best score: %\f | Genome: %v\n", generationsPerSecond, survivors[0].GetScore(), survivors[0].genome)
		g.lastGenerationCount = g.generation
	}

	for len(survivors) < len(g.cells) {
		parent := survivors[rand.Intn(len(survivors))]
		child := parent.Child()
		survivors = append(survivors, *child)
	}

	g.cells = survivors
}
