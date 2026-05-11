package main

import "fmt"

// variadic function :

var numsTotal = func(a ...int) {

	total := 0

	for _, i := range a {
		total = total + i
	}

	fmt.Println(total)

}

func main() {

	// 1. Wait youtube interfaces
	nums2 := []int{1, 2, 3, 4, 5}
	numsTotal(nums2...)

	numsTotal(2, 3)
	numsTotal(2, 3, 6)
	numsTotal(2, 3, 6, 7, 2)
	numsTotal(2, 3, 6, 7, 2, 5)
}
