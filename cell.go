package main

import (
	"math"
	"math/rand"
)

type Cell struct {
	output float64
	genome []float64
}

func (c *Cell) Apply() {
	var translated []float64
	for _, v := range c.genome {
		translated = append(translated, translator(v))
	}
	c.output = salomon(translated)
}

func (c *Cell) Child() *Cell {
	genomeSize := len(c.genome)
	newGenome := make([]float64, genomeSize)

	randomIndex := rand.Intn(genomeSize)

	for i := range newGenome {
		newGenome[i] = c.genome[i]
	}
	// mutation sur 1 gène
	newGenome[randomIndex] = rand.Float64()

	return &Cell{
		genome: newGenome,
		output: 0,
	}
}

func (c *Cell) GetScore() float64 {
	return c.output
}

func (c *Cell) Reset() {
	c.output = 0
}

func translator(x float64) float64 {
	return 8*x - 4
}

func salomon(x []float64) float64 {
	var somme float64

	for _, v := range x {
		somme += v * v
	}

	return 1 - math.Cos(2*math.Pi*math.Sqrt(somme)) + .1*math.Sqrt(somme)
}

type ByScore []Cell

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Less(i, j int) bool { return a[i].GetScore() < a[j].GetScore() }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
