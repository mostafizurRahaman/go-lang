package main

import "fmt"

// ! 1. Declare a func with parameter type and return type:
func add(num1 int, num2 int) int {
	return num1 + num2
}

// ! 2. If one or more parameter has same type you can define type only once:
func multiply(num1, num2, num3 int) int {
	return num1 * num2 * num3
}

func main() {

	// ! Call the first function for addition:
	fmt.Println(add(1, 4))
	fmt.Println(add(1, 6))
	fmt.Println(add(6.0, 4))

	// ! Call the second function for multiplication :
	fmt.Println(multiply(1, 2, 3))
	fmt.Println(multiply(3, 3, 3))

	// ! 3. A function which can receive another function as parameter:
	getSqaure := func(a int) int {
		return a * a
	}

	getCubicSqaure := func(a int) int {
		return a * a * a
	}

	processFn := func(fn func(a int) int, a int) int {
		return fn(a)
	}

	sqr1 := processFn(getSqaure, 20)
	sqr2 := processFn(getCubicSqaure, 10)
	fmt.Println(sqr1, sqr2)

	// ! Create a function which will return a function :
	processFn2 := func() func(a int) int {
		return func(a int) int {

			return a
		}
	}

	a := processFn2()
	fmt.Println(a(50))

}
