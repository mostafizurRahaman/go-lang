package main

import "fmt"

// const: declare and assign value in same line, cannot redeclare and reassign.
// you can write it outside of function globally :

// Typed const
const lang string = "PROGRAMMING"

// UnTyped const
const langName = "GO LANG"

// You can define var also oustide but you can able to delcare shorthand variable outside of the function: 


func main() { 
// Access the global const 
fmt.Println("Lang", lang)
fmt.Println("Lang", langName)


	// Grouped( Block) const declaration: 

	const (
		a = 20 // UnTyped, 
		b string = "Rajib"
		c float32 = 6.8
		isSameVariable = false
	)

}