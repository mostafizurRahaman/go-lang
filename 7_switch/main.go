package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch

	i := 2


	// CASE : SINGLE SWITCH 
	switch i {

		case 1:
			fmt.Println("Print  One", i)
		case 2:
			fmt.Println("Print Two", i)
		case 3:
			fmt.Println("Print Three", i)
		default:  // NOTE: optional 
			fmt.Println("Other values", i)
	}

	// CASE : MULTI SWITCH 
	switch time.Now().Weekday() {

		case time.Saturday, time.Sunday:
			fmt.Println("It's weekend")
		default: 
			fmt.Println("It's weekday")
	}

	// CASE : TYPE Switch: 

	whoAmI := func(i interface{}){

		switch i.(type){
			case int:	
				fmt.Println("its an integer")
			case string:	
				fmt.Println("its a string")
			case float32:	
				fmt.Println("its a float32")
			case float64:	
				fmt.Println("its a float32")
			case bool:	
				fmt.Println("its a bool")
			default: 
				fmt.Println("Other")
		}

	}

	whoAmI(true)
	

}