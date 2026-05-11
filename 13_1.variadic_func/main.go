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

var log = func(prefix string, params ...interface{}) {

	fmt.Print(prefix, " : ")
	for _, value := range params {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")

}

func main() {

	// 1. Wait youtube interfaces
	nums2 := []int{1, 2, 3, 4, 5}
	numsTotal(nums2...)

	numsTotal(2, 3)
	numsTotal(2, 3, 6)
	numsTotal(2, 3, 6, 7, 2)
	numsTotal(2, 3, 6, 7, 2, 5)

	// Call logger:
	log("INFO", 1, 2, "Mostafizur")
	log("Error", 1, 2, "Mostafizur")

}
