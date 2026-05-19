package main

import "fmt"

func processTask(done chan bool) {

	defer func() { done <- true }()

	fmt.Println("------------Processing ----------------")
	for i := range 400 {
		fmt.Println("Doing task :->", i)
	}

}

func main() {

	fmt.Println("---------Done channel ----------")

	done := make(chan bool)
	go processTask(done)

	// ! This will work as blocking like wait
	<-done

	// ! It will work like  wait group:
	// ! Recommend :when you have single go routine to render use this way
	// ! When you have multiple go routine use waitgroups

}
