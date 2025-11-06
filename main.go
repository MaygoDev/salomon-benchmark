package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	game := Game{
		dimensions: 5,
	}
	game.Populate(1000)
	for {
		game.Generate(20)
	}
	elapsed := time.Since(start)
	fmt.Printf("It took %5fseconds\n", elapsed.Seconds())
}
