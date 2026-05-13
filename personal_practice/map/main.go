package main

import (
	"fmt"
	"maps"
)

func main() {

	// ! 1. Declare without assigning value:
	// ! If you declare a map without assigning a value, it is equal nil
	var persons map[string]string
	fmt.Println(persons)
	fmt.Println(persons == nil)

	// ! Assign value
	persons = map[string]string{
		"name": "Goni",
	}

	// ! Add new field: map[fieldName] = value
	persons["age"] = "Tweenty four"
	persons["isStudent"] = "yes"
	persons["status"] = "pending"
	fmt.Println(persons)

	// ! Access value : mapName["key"]
	fmt.Println(persons["age"])

	// ! Access a field value which does not exist: it will return falsy value:
	fmt.Println(persons["age1"])

	// ! Override a field value: map[fieldName] = newValue
	persons["name"] = "Fahim hossain"
	fmt.Println(persons)

	// ! delete field from map: delete(person, "fieldName")
	delete(persons, "isStudent")
	fmt.Println(persons)

	// ! 2. Declare a map with make() function:
	var person2 = make(map[string]string)
	fmt.Println(person2 == nil)
	person2["name"] = "Azmir Azam Khan"
	person2["age"] = "Tweenty five"
	person2["profession"] = "Flutter App Developer"
	person2["phoneNumber"] = "unknown"
	fmt.Println(person2)

	// ! 3. Create map with var and := shorthand var:
	var person3 = map[string]string{}
	fmt.Println(person3 == nil)
	person3["name"] = "Mostafizur Rahaman"
	fmt.Println(person3)

	// ! 4. Assign a map with := shorthand variable:
	person4 := map[string]string{}
	fmt.Println(person4 == nil)
	person4["name"] = "Osman goni"
	person4["age"] = "Tweenty six"
	fmt.Println(person4)

	// ! Compare two map: with maps and it's utilities

	m1 := map[string]int{"unit": 3, "quantity": 3}
	m2 := map[string]int{"unit": 3, "quantity": 3}
	m3 := map[string]int{"unit": 1, "quantity": 4}

	fmt.Println("Is m1 is equal m2 ?", maps.Equal(m1, m2))
	fmt.Println("Is m2 is equal m3", maps.Equal(m2, m3))

	// ! Go lang idoms to check map value is exists or not:

	i, ok := person4["name"] // ! i return value, ok return true or false (if value exists) if value doesn't exists, i return falsy value ('', false or 0)

	if ok {
		fmt.Println("All is ok", i)
	} else {
		fmt.Println("All is not ok", i)
	}

	// ! 

}
