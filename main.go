package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	game := Game{
		dimensions: 10,
	}
	game.Populate(200)
	for i := 0; i < 20000; i++ {
		game.Generate(35)
	}
	elapsed := time.Since(start)
	fmt.Printf("It took %5fseconds\n", elapsed.Seconds())
}
