package main

import "fmt"

// ! 1. For getting multiple:
func logger(logType string, a ...interface{}) {
	fmt.Print(logType)

	for _, v := range a {
		fmt.Printf(" %v ", v)
	}
	fmt.Println()
}

// ! 2. Create antoher one variadic function which will get n number of parameters and return sum of them.

func sumUp(n ...int) int {
	sum := 0

	for _, v := range n {
		sum = sum + v
	}

	return sum
}

func main() {
	fmt.Println("Veriadic functions")

	logger("INFO", 1, 2, 3, 4, 5, "MOstafizur", "Osman goni")
	logger("INFO", 1, 2, 3, 4, 5, "MOstafizur", "Osman goni")
	logger("INFO", 1, 2, 3, 4, 5, "MOstafizur", "Osman goni")
	logger("INFO", 1, 2, 3, 4, 5, "MOstafizur", "Osman goni")
	logger("INFO", 1, 2, 3, 4, 5, "MOstafizur", "Osman goni")

	fmt.Println("Print", sumUp(20, 20, 20, 10))

	nums := []int{2, 3, 5, 5, 6}

	sumUp(nums...)

}
