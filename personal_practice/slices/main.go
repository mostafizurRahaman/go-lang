package main

import (
	"fmt"
	"slices"
)

// ! There are some way we can declare a slice:
// ! 1. []int{} by using this format
// ! 2. shorthand or var
// ! 3. make() by using make function
// ! 4. By using slice operator

func main() {

	// ! 1. If we declare a slice, don't initialize value in same line that slice will be nil
	var slice1 []int
	fmt.Println(slice1)
	fmt.Println(slice1 == nil)

	// ! Append new element into the array:
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 4)

	// ! Check capacity of the slice :
	fmt.Println(cap(slice1))

	// ! Check the length of the slice:
	fmt.Println(len(slice1))

	// ! ** If need capacity increase like this: newCapacity = previousCapity * 2

	// ! 2. Declare a slice with make() :-> for var/ :=

	slice2 := make([]int, 2, 3) // ! parameters: type, length, capacity
	// ! If capacity not provided, the default capacity will be equal to length.
	fmt.Println(slice2)

	var slice3 = make([]string, 2, 3)
	fmt.Println(slice3)

	slice3[0] = "Khadakar efaz"
	slice3[1] = "Rafiq Vai"

	slice3 = append(slice3, "Mostafiz")
	slice3 = append(slice3, "Osman goni")
	slice3 = append(slice3, "Masum sakar", "Ligger Masum", "Nm sujon", "Azmir azam khan", "fahim")
	fmt.Println(slice3)

	// ! 3. Create another one slice from another array or slice:
	friendList := []string{"Mostafiz", "Ratul Hossain", "Ajay", "Azmir", "Rayhan"}

	// ! [startIndex:endIndex]
	smallerFriendList := friendList[1:4]
	fmt.Println(smallerFriendList)

	// ! [startIndex:]
	smallerFriendList1 := friendList[1:]
	fmt.Println("smallerFriendList1", smallerFriendList1)

	// ! [:endIndex]
	smallerFriendList2 := friendList[:4]
	fmt.Println("smallerFriendList2", smallerFriendList2)

	// ! Copy one slice to another slice (dest, source)
	copy(smallerFriendList1, smallerFriendList2)
	fmt.Println(smallerFriendList1)

	// ! Slices builtin functions:
	fmt.Println(slices.Equal(smallerFriendList1, smallerFriendList2))

	// ! Check is value available inside slice:
	fmt.Println(slices.Index(smallerFriendList1, "Ratul Hossain2")) // ! if available will return 1 or -1

	// ! Check is slices shorted?:
	fmt.Println(slices.IsSorted(smallerFriendList))
}
