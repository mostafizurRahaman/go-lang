package main

import "fmt"

func main() {

	// ! 1. normal if statement for condition check:
	i := 50
	if i > 20 {
		fmt.Println("I is greater then 20")
	}

	// ! 2. If condition not true we have use else block :

	age := 17
	if age >= 18 {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("You are not an adult!")
	}

	// ! 3. else if :
	childAge := 12
	if childAge <= 12 {
		fmt.Println("You are a kid")
	} else if childAge < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You  are adult")
	}

	// ! 4. You can  declare a variable into if statement and they you can use it: and those are accessible between the if block, if else and else blocks.

	if a, b := 1, 3; a > b {
		fmt.Println("A is greater then b", a, b)
	} else if a == b {
		fmt.Println("A is equal B.", a, b)
	} else {
		fmt.Println("A is smaller then B.", a, b)
	}

	// !  5. We can use logical OR (|| ) or Logical  And (&&) operator inside if condition for multiple condition checking :

	isStudent := true
	hasNidCard := false

	if isStudent && hasNidCard {
		fmt.Println("You have both access!")
	} else if !isStudent && hasNidCard {
		fmt.Println("You have Nid card")
	} else if isStudent && !hasNidCard {
		fmt.Println("You have student card!")
	} else {
		fmt.Println("You have nothing!")
	}

	// ! IMR: Go  lang don't have any ternary operator.

}
