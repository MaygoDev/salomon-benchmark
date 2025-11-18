package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	game := Game{
		dimensions: 1,
	}
	game.Populate(200)

	var (
		generations         int
		mutex               sync.Mutex
		lastPrint           = time.Now()
		lastGenerationCount = 0
	)

	for {
		go game.Generate(80, &generations, &mutex)

		if time.Since(lastPrint) > time.Second {
			lastPrint = time.Now()

			mutex.Lock()
			generationsPerSecond := generations - lastGenerationCount
			mutex.Unlock()

			fmt.Printf("Generation/s: %d\n", generationsPerSecond)
			lastGenerationCount = generations
		}
	}

	fmt.Printf("Generations : %d", generations)
}
