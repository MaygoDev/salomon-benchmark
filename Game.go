package main

import (
	"math/rand"
	"sort"
	"sync"
)

type Game struct {
	dimensions int
	cells      []Cell
}

func (g *Game) Populate(count int) {
	for i := 0; i < count; i++ {
		genome := make([]float64, g.dimensions)
		for j := 0; j < g.dimensions; j++ {
			genome[j] = rand.Float64()
		}
		g.cells = append(g.cells, Cell{genome: genome})
	}
}

func (g *Game) Generate(percent int, generations *int, mutex *sync.Mutex) {
	mutex.Lock()
	*generations++
	mutex.Unlock()

	for i := range g.cells {
		g.cells[i].Apply()
	}

	// tri selon le score
	sort.Sort(ByScore(g.cells))

	survivors := g.cells[:len(g.cells)*percent/100]

	for len(survivors) < len(g.cells) {
		parent := survivors[rand.Intn(len(survivors))]
		child := parent.Child()
		survivors = append(survivors, *child)
	}

	g.cells = survivors
}
