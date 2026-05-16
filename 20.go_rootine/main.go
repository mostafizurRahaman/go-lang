package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	task := func(id int) {
		defer func() {
			elapsed := time.Since(start)

			fmt.Println("Total execution  time", elapsed.Milliseconds())

		}()

		fmt.Println("Doing task", id)
	}

	for i := range 10 {
		go task(i + 1)
	}

	// !self calling function for  go routine:
	for i := range 10 {
		go func() {
			fmt.Println("Doing self task", i+1)
		}()
	}

	time.Sleep(time.Second * 2)

}
