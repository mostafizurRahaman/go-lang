// package main

// import (
// 	"fmt"
// 	"time"
// )

// // ! Senario 2:

// func task(numChan chan int, num int) {
// 	numChan <- num
// }

// // ! Senario 3:
// func receiverTask(strChan chan string) {
// 	msg := <-strChan
// 	fmt.Println(msg)
// }

// // ! Senario 4:
// func renderSum(queueChan chan int) {

// 	defer time.Sleep(time.Second)
// 	sum := 0
// 	for num := range queueChan {

// 		sum = sum + num
// 		fmt.Printf("The value of Sum : %v \n", sum)
// 		time.Sleep(time.Second)
// 	}

// }

// // ! Senario 5: go routine synchronizer:
// func completeTask(done chan bool) {

// 	defer func() {
// 		fmt.Println("Defered Called")
// 		done <- true
// 	}()
// 	fmt.Println("Processing .............")
// 	for i := range 30 {
// 		fmt.Println(i)
// 	}
// 	fmt.Println("Processed")

// }

// func main() {

// 	fmt.Println("Channel")
// 	/*
// 		! Senario one:

// 		! go channel created for int type carry
// 		numChan := make(chan int)

// 		? Send data into channel
// 		numChan <- 1

// 		? Receive data into cahnnel
// 		res := <-numChan

// 		! We will receive an Fatal (DEADLOCK ERROR) Because Reciver will wait for response. But because of single go routine (main go routine) go routine is busy with
// 		fmt.Println(res) */

// 	/*
// 		! Senario Two:
// 	*/

// 	/*
// 		? Another go routine for integer carray :
// 		var numChan = make(chan int)

// 		go task(numChan, 243)

// 		 ? Receive data into channel :
// 		msg := <-numChan

// 		fmt.Println(msg)
// 	*/

// 	/*
// 		 ? Senario Three alternative :We will recieve inside go routine and send from main go routine:

// 		var strChan = make(chan string)

// 		go receiverTask(strChan)
// 		strChan <- "Mostafizur Rahaman"

// 		time.Sleep(time.Millisecond * 5)
// 	*/

// 	/*
// 		 ! 4. Senario 4 : Here we can receive multiple data like queue :
// 		var queueChan = make(chan int)

// 		i := 0
// 		go renderSum(queueChan)
// 		for {
// 			i++
// 			queueChan <- i
// 		}
// 	*/
// 	// ! 5 Go Routine synchronizer:

// 	var done = make(chan bool)

// 	go completeTask(done)

// 	<-done

// }
