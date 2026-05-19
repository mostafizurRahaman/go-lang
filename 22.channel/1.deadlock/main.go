package main

import "fmt"

func main() {

	// ! Create a channel :
	numChan := make(chan int)

	numChan <- 20 // Send  (Blocking)

	// ! Sending is block, So The  programing couldn't execute the receiving. And Deadlock will happen.

	res := <-numChan // Receive (Blocking)

	fmt.Println(res)

}
