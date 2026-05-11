package main

import "fmt"

func main() {

	nums := func(a ...int) int {
		fmt.Println(a)
		acc := 0
		for _, elem := range a {
			acc += elem
		}
		return acc
	}

	fmt.Println(nums(1, 2, 3, 4, 5))
	fmt.Println(nums(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(nums(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12))
	fmt.Println(nums(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15))
	fmt.Println(nums(30, 20, 20))

	numerics := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20}

	fmt.Println(numerics)
	fmt.Println(nums(numerics...))
	// fmt.Println()

}
