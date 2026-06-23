package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("worker %d done\n", id)
		}(i)
	}

	wg.Wait() // blocks until all 3 call Done()
	fmt.Println("all workers finished")
}
