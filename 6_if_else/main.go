package main

import "fmt"

func main() {

	// CASE : If else:

	// age := 16
	// if age >= 18 {
	// 	fmt.Println("Person is an adult!")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// CASE: IF , else if, else
	age := 18
	if age >= 18 {
		fmt.Println("Person is an adult!")
	} else if age >= 12 {
		fmt.Println("Person is teenager!")
	} else {
		fmt.Println("Person is a kid")
	}

	// CASE: IF ELSE WITH LOGICAL OPERATOR: 
	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission{ 
		fmt.Println("1. You are an admin and have permission.")
	} else if role == "admin" && !hasPermission { 
		fmt.Println("2. You are an admin and have no permission.")
	} else if role != "admin" && hasPermission { 
		fmt.Println("3. You are not an admin and have permission.")
	} else { 
		fmt.Println("4. You are not an admin and have no permission.")
	}

	hasAccess := role == "admin" && hasPermission
	fmt.Println("Has Access", hasAccess)


	// CASE : We can declare variable inside the if construct
	// CASE :

	if age:=10; age >= 18{ 
		fmt.Println("Person is an adult", age)
	} else if (age >= 12) { 
		 fmt.Println("Person is a teenager", age)
	} else { 
		fmt.Println("Person is a kid", age)
	}

	//  CASE : go doesn't have no ternary operator: 
	
}