package main

import (
	"fmt"
	"maps"
)

// map = dict (python), object(javascript), hash (ruby)

func main() {

	// Declare a map with var or := shorthand operator:

	var person1 = map[string]string{
		"name":       "Rabby Hosaain",
		"education":  "honours",
		"fatherName": "Sekantor Bhuiyan",
	}

	// set values :
	person1["motherName"] = "Rohima  Khatun"
	person1["schoolName"] = "Mohammod Ullah High school"
	fmt.Println(person1)

	// get values :
	// IMP:If key does not exists in map, its return ZERO value:
	fmt.Println("Father Name -> ", person1["fatherName"])
	fmt.Println(person1["new"])

	// Declare a map with make function:
	var person2 = make(map[string]int)

	// set values:
	person2["age"] = 20
	person2["height"] = 5
	person2["weight"] = 48
	person2["bmi"] = 2600

	fmt.Println(person2)

	// if not exists will return ZERO values / Falsy values:
	fmt.Println(person2["nei"])

	// delete price from map:

	delete(person2, "bmi")

	fmt.Println(person2)

	// clear to make map empty:
	// clear(person2)
	fmt.Println(person2)

	// check is the values  exists:

	k, OK := person2["height"]

	if OK {
		fmt.Println("All  Ok", k)
	} else {
		fmt.Println("Not ok", k)
	}

	// is equal both map?:
	m1 := map[string]int{"price": 20, "phones": 3}

	m2 := map[string]int{"price": 20, "phones": 3}

	fmt.Println(maps.Equal(m1, m2))

	// Declare  an empty map:
	var m3 map[string]int // this map will  be nil  because the map  declared but not initialized .

	fmt.Println(m3)
	fmt.Println(m3 == nil)
}
