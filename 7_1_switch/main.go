package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("SWITCH PRACTICING")


	//  NOTE: Single switch:

	letter := "b"

	switch (letter){ 

		case "a": 
				fmt.Print("A is a vowel")
		case "e": 
				fmt.Print("E is a vowel")
		case "i": 
				fmt.Print("I is a vowel")
		case "o": 
				fmt.Print("O is a vowel")
		case "u": 
				fmt.Print("U is a vowel")
		default :
			fmt.Print("This is a consonant")

	}


	// NOTE: Multiple switch: 

	switch (time.Now().Weekday()){ 
		case time.Friday, time.Saturday: 
			fmt.Println("Today is weekend!")
		case time.Sunday, time.Monday: 
			fmt.Println("Today is holiday!")
		default :
			fmt.Print("Today is weekday.")
	}

	// NOTE: TYPE Switch 

	whoAmI := func(i interface{}) {

		switch i.(type) { 
			case int: 
				fmt.Println("This is an integer")
			case string: 
				fmt.Println("This is a string")
			case bool: 
				fmt.Println("This is a bool")
			case float32, float64: 
				fmt.Println("This is a float.")
			default: 	
				fmt.Println("Other types")
		}

	}

	whoAmI("s")
	whoAmI(25)
	whoAmI(25.26)
	whoAmI(true)
	whoAmI([5]int{1,2,6})

	
}