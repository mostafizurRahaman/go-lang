package main

import (
	"fmt"
	"time"
)

// ! Reciever from chain to num chain
// func processNum(numChan chan int) {

// 	result := <-numChan // ! blocking

// 	fmt.Println(result)
// }

// ! Sender to numchian:
func sendNewValueToNumChan(numChan chan int, num int) {
	numChan <- num // blocking
}

// ! Reciever queue chain
func receiveQueueValue(queueChan chan int) {

	for i := range queueChan {

		res := i

		fmt.Println(res)
		time.Sleep(time.Second)
	}

}

func main() {

	// ! go routine :

	// ! declare a chan :
	// numChan := make(chan int) // unbuffered channel (sending and receiving is blocking)

	// go processNum(numChan)
	// numChan <- 20 // ! send to channel (blocking)

	// time.Sleep(time.Second)

	// ! Declare a num chan: (alternative of first one)
	// numChan := make(chan int)

	// go sendNewValueToNumChan(numChan, 450)

	// result := <-numChan // blocking

	// fmt.Println(result)

	// ! Send data as queue:
	queueChain := make(chan int)

	go receiveQueueValue(queueChain)

	for i := range 20 {
		queueChain <- i
	}

}
