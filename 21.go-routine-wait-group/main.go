package main

import (
	"fmt"
	"sync"
)

// ! A function with waiting group:

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Doing Task", i+1)
}

func main() {

	// ! Create waiting time:
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

}
