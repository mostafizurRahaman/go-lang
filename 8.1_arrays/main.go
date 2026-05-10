package main

import "fmt"

func main() {

	fmt.Println("---------Learn Array------------")

	// Declare array without initialing values:
	var names [5]int
	names[2] = 3

	fmt.Println(names)
	fmt.Println(names[2])

	// Another way to declare an fixed  length array:
	var friendNames = [10]int{}
	fmt.Println(friendNames)

	// Another way to define a array () not fixed.
	// ->  inferred length from initialized values:
	var friendNames2 = [...]int{2, 3, 4, 56, 6}
	fmt.Println(friendNames2)

	// Sorthand array :

	// Declare length, type:
	nums := [6]int{}
	fmt.Println(nums)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	fmt.Println(nums)

	// Declare type, not length and shorthand way:

	bools := [...]bool{true, false, true} // based  on initial values infered the length.
	fmt.Println(bools)
	bools[2] = true
	bools[1] = false
	fmt.Println(bools)

	// Check  length of the array:

	fmt.Println(len(bools))

	// Two dimentional array:
	twoDimentionalArr := [2][2]int{{2, 2}, {2, 3}}
	fmt.Println(twoDimentionalArr)

	fmt.Printf("twoDimentionalArr: %v\n", twoDimentionalArr)
	fmt.Println(twoDimentionalArr[1][0])
	fmt.Println(len(twoDimentionalArr))

	// Run a for lop into the first array

	for i := 0; i < len(twoDimentionalArr); i++ {
		for j := 0; j < len(twoDimentionalArr[i]); j++ {
			nums := twoDimentionalArr[i]
			fmt.Println(i, ":", j, "<--->", nums[j])
		}
	}

}
