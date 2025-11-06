package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Game struct {
	dimensions int
	cells      []Cell
	generation int
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
}

func (g *Game) Generate(percent int8) {
	g.generation++
	for i := range g.cells {
		g.cells[i].Apply()
	}

	// Sort by lowest score
	sort.Sort(ByScore(g.cells))

	survivors := g.cells[:len(g.cells)*int(percent)/100]

	fmt.Printf("Generation %d: Best score = %.15f\n", g.generation, survivors[0].GetScore())

	for len(survivors) < len(g.cells) {
		parent := survivors[rand.Intn(len(survivors))]
		child := parent.Child()
		survivors = append(survivors, *child)
	}

	g.cells = survivors
}
