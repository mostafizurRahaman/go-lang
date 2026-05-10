package main

import "fmt"

// Array is  a list of same data types:

func main() {

	// Array  declaration :
	var names [4]string
	names[0] = "Mostafizur"
	names[1] = "Ratul Hossain"
	fmt.Println(names)

	// Array for int:
	var nums [10]int
	nums[4] = 5
	nums[5] = 10
	fmt.Println(nums)

	// Array for bool:
	var bools [5]bool

	bools[1] = true
	bools[4] = true
	fmt.Println(bools)

	// Another way of array declaration:
	var myFriends = [10]string{"Mostafizur", "Ratul Hossain", "Rakib Khan"}
	fmt.Println(myFriends)
	myFriends[5] = "Rahul"
	fmt.Println(myFriends)

	//Assume variable length from data:
	var myColigues = [...]string{"NM SUJON", "GONI", "RAFIQ", "EFAZ"}
	fmt.Println(myColigues)

	// Single line array:
	friendList := [10]string{}
	friendList[0] = "AMI"
	friendList[1] = "TUME"
	friendList[2] = "SHE"
	friendList[3] = "THEY"
	fmt.Println(friendList)

	// Single line array: Assugme length from data :
	pronouns := [...]string{"They"}
	fmt.Println(pronouns)

	// Two dimentional array:
	arr2d := [2][2]int{{2, 3}, {3, 4}}
	fmt.Println(arr2d)

}
