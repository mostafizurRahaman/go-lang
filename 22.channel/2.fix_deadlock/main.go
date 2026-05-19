package main

import "fmt"

// ! To fix  deadlock we have receive or send  data from another one go  routine.
// ! Channel  is a  communication between  go routines for memory sharing.

func main() {

	// ! Buffered  channel  (which will blocking while sending and receiving)
	numChan := make(chan int)

	// ! Create a go  routine which  will receive the value
	// ! Self  calling function
	go func() {
		res := <-numChan
		fmt.Println(res)
	}()

	// ! Blocking Send
	numChan <- 300

	fmt.Println("Executation done")

	type order struct {
		name string
	}

	structChain := make(chan order)

	go func() {
		structChain <- order{
			name: "Mostafiz",
		}
	}()

	res := <-structChain

	fmt.Println(res)

}
