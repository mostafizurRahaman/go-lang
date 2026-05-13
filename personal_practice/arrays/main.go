package main

import "fmt"

// ! We can declare array in mutliple way in go lang:

func main() {

	// ! 1. Declare an array:
	var friends [5]int

	// set value of index:
	friends[0] = 1

	// get value from index:
	fmt.Println(friends[0])

	// check length of the array by using len()
	fmt.Println(len(friends))

	// we can override the value also :
	friends[0] = 3

	fmt.Println(friends)

	// ! 2. Declare and initialize in the same line:
	var myFriends = [5]string{}
	fmt.Println(myFriends)

	// ! 3. Infer length based on the initial value of an array:
	var nums = [...]int{5, 5}
	fmt.Println(len(nums))

	// override the array value:
	nums = [...]int{3, 5}
	fmt.Println(len(nums))

	// ! 4. Declare shorthand array with fixed length:
	arr1 := [7]int{}
	fmt.Println(arr1)

	// ! 5. Declare shorthand array with infered length :
	arr2 := [...]int{}
	fmt.Println(arr2)

	// ! 6 Two dymentional array:
	arr2d := [3][3]int{{1, 2, 3}, {2, 5, 6}, {1, 6, 8}}

	fmt.Println(arr2d)

}

// FALSY: 0, "", nil,
