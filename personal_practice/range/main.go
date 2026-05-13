package main

import "fmt"

func main() {
	fmt.Println("Range")

	// ! 1. Classic range:
	for i := range 3 {
		fmt.Println(i)
	}

	// ! 2. Run range with string :
	name := "Mostafizur"
	for i, v := range name {
		fmt.Println("I : ", i, "Unicode value :", v, "value", string(v))
	}

	// ! 3. Range run into slice :
	friends := []string{"Mostafiz", "Ratul", "sakib"}

	for i, v := range friends {
		fmt.Println("I : ", i, "value : ", v)
	}

	// ! 4. Run range into map :
	person := map[string]string{
		"name":      "Mostafizur Rahaman",
		"age":       "Tweenty four",
		"isStudent": "Yes",
	}

	fmt.Println(person)
	for k, v := range person {
		fmt.Println("Property Name ", k, "Value ", v)
	}

	// ! 5. To get only we can access key, we can use access only the fisrt value:

	for k := range person {
		fmt.Println("Property Name ", k)
	}

}
