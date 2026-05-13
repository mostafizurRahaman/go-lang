package main

import "fmt"

func parent() func() int {

	var counter = 0
	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	fmt.Println("Elementor")

	a := parent()

	fmt.Println("A _> ", a())
	fmt.Println("A _> ", a())
	fmt.Println("A _> ", a())
	fmt.Println("A _> ", a())
	fmt.Println("A _> ", a())
	fmt.Println("A _> ", a()) // 6

	b := parent()

	fmt.Println("B _> ", b())
	fmt.Println("B _> ", b())
	fmt.Println("B _> ", b())
	fmt.Println("B _> ", b()) // 4
}
