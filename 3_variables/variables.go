package main

import "fmt"

func main() {

	// Variable with var with explicit type:
	// Declare,initialization, you can do seperately 
	// Reassign value, you can 
	// Redeclare you can't 
	var name string = "Mostafizur Rahaman"
	fmt.Println(name)
	name = "Ratul Hossain"
	fmt.Println(name)


	//  Var for (Infered Typed)
	// variable infered type based on first initialized value: 
	var firstName = "Rohim Miah" // this variable infered it's type from the value. 
	fmt.Println(firstName)
	firstName = "Rahul Miah"
	fmt.Println(firstName)
	var isAdult = true
	var height = 3.5 
	fmt.Println("is Adult", isAdult)
	fmt.Println("Height", height)


	// Sort hand syntax of GO Lang var: 
	// Declare and initialization same line (should be)
	// Don't need to define explicit type 
	// This variable auto infered from first initialization. 
	// But you can Reassign with = sign multiple times: 
	language := "Go Lang"
	fmt.Println("Language", language)
	language = "JavaScript"
	fmt.Println("Re Assigned Shorthand var", language)


	// Declare first but assign seperately: (Have to define explicit type)
	var langName string
	langName = "Python"
	fmt.Println(langName)
	

	var langVersion float32
	langVersion = 3.5
	fmt.Println(langVersion)

	// Multi varibale  declaration:

	// (Explicit Type) Variables type should be same 
	// Cannot provide multiple types for single line with explicit type 
	var a , b, c int = 50,60,70
	fmt.Println(a,b,c)

	// If Implicit you can define multiple typed variable in a single line: 
	var d,e,f= "D is a letter", 0.5, false
	fmt.Println("D", d)
	fmt.Println("E", e)
	fmt.Println("F", f)

	// Multi line shorthand variable, with multiple typed data: 
	g,h, i := "Mostafizur Rahaman", 3, true
	fmt.Println("G", g)
	fmt.Println("H",h)
	fmt.Println("I", i)


	// Block variable declaration 

	var (
		ab string = "Mostafiz"
		bc int = 20
		cd bool = true
	) 

	fmt.Println("AB", ab)
	fmt.Println("BC", bc)
	fmt.Println("CD", cd)

	
}