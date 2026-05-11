package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	
	// for an array: 
	fmt.Println("------------- For  index + value ---------- of a map")
	for i , num:= range nums { 
		fmt.Println("index", i)
		fmt.Println(num)
	}

	// declare a map: 
	var person1 = map[string]string{}
	person1["name"] = "Azmir Khan"
	person1["age"] = "Twenty five"
	person1["school"] = "Bramanbaria high school"

	fmt.Println("________ For key and value both -------------- ", )
	// Iterate a map with range key: 
	for key, value :=range person1 { 
		fmt.Println("Key ->", key, "Value ->", value)
	}

	// Give single variable to get only keys: 
	fmt.Println("________ Only for keys -------------- ", )
	for key := range person1 { 
		fmt.Println("Key -> ", key)
	}

	//  ------------ Run a range loop into a string

	str := "mostafizur"

	for i, letter := range str { 
		fmt.Println("Index => ", i)
		fmt.Println(" Unicode Letter => ", letter)
		fmt.Println("Letter => ", string(letter))
		
	}



}