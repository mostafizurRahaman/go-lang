package main

import (
	"fmt"
	"time"
)

func queueRender(queueChan chan int) {

	for i := range queueChan {
		res := i

		fmt.Println(res)

		time.Sleep(time.Second)
	}

}

func main() {

	fmt.Println("Something done...............")

	queueChan := make(chan int)

	go queueRender(queueChan)

	for i := range 20 {
		queueChan <- i
	}

}
