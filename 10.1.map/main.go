package main

import (
	"fmt"
	"maps"
)

func main() {

	//  Decalaration method one :
	m1 := map[string]int{}
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 2
	m1["four"] = 4
	fmt.Println("M1", m1)

	// Declaration method two:
	var m2 = map[string]int{}
	m2["one"] = 1
	m2["two"] = 2
	m2["three"] = 2
	m2["four"] = 4
	fmt.Println("M2", m2)

	// Declaration method three (in this case value will be nil)
	// value not assigned yet: 
	var m3 map[string]int

	m3 = map[string]int{
		"one": 1, 
		"two": 2, 
		"three": 3, 
	}

	// update value 
	m3["one"] = 2
	m3["three"] = 5

	// get value 
	fmt.Println(m3["three"])

	//  delete value : 
	delete(m3, "three")
	fmt.Println(m3)

	// after accessing missing key: Return falsy value
	fmt.Println(m3["three"])  // 0, false or empty


	// After access in go you can check the property like this: 
	// okay will return true, false
	// if value exists "i" will value, if not exists i will be Falsy value Like (0, "", false)
	i, ok := m3["two"]

	if ok { 
		fmt.Println("All ok", i, ok)
	} else { 
		fmt.Println("All not ok", i, ok)
	}

	// Check multiple array is equal or not: 

	m5 := map[string]int{"one": 2, "two": 1, "three": 3}
	m6 := map[string]int{"one": 2, "two": 1, "three": 3}

	fmt.Println(maps.Equal(m5, m6))

}