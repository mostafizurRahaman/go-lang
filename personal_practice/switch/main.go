package main

import (
	"fmt"
	"time"
)

// ! Single switch:

// Check grade based on result:

func main() {

	marks := 33

	switch marks {

	case 100:
		fmt.Println("How brilliant student you are!")
	case 80:
		fmt.Println("You got A+. Congratulations!")
	case 60:
		fmt.Println("Not Bad. You achieve B!")
	case 33:
		fmt.Println("You passed!")
	case 0:
		fmt.Println("You  are fail!")
	default:
		fmt.Println("Default value  is printing")

	}

	// ! Multiple switch  condition in single case:
	today := time.Now().Weekday()
	fmt.Println(today)

	switch today {
	case time.Friday, time.Saturday:
		fmt.Println("Weekdays")
	case time.Sunday:
		fmt.Println("First office day!")
	case time.Monday:
		fmt.Println("Practice day!")
	case time.Tuesday, time.Thursday, time.Wednesday:
		fmt.Println("Being serious to back to track")
	default:
		fmt.Println("Nothing")
	}

	//! Typed switch

	whoAmI := func(i interface{}) {

		switch i.(type) {

		case int:
			fmt.Println("INTEger", i)
		case float32, float64:
			fmt.Println("Floating", i)
		case string:
			fmt.Println("String", i)
		case bool:
			fmt.Println("String", i)
		default:
			fmt.Println("OTHERS")
		}
	}

	whoAmI(2)
	whoAmI("mostafiz")
	whoAmI(2.3)
	whoAmI(true)

	v := [3]int{1, 2, 3} // array
	whoAmI(v)

	v1 := []string{"Mostafiz"} // slice
	whoAmI(v1)

	v3 := map[string]string{ // map
		"name": "Mostafiz",
	}
	whoAmI(v3)
}
