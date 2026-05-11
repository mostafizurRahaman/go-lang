package main

import "fmt"

// 1. Create a function which will took two parameter :
var add = func(a int, b int) int {
	return a + b
}

// 2. Create a function with same type parameter:
// if  parameters types are same we can right like this,
// a, b  int
var add1 = func(a, b int, firstName string) int {
	fmt.Println(firstName)
	return a + b
}

// 3. A function can return another receive another function as parameter:
var processIt = func(fn func(a int) int) int {
	return fn(2)
}

func main() {

	fmt.Println("Print languages")

	fmt.Println("-----------------------------")
	sum := add(1, 2)
	fmt.Println(sum)
	fmt.Println("-----------------------------")

	// fmt.Println("-----------------------------")
	// sum := add(1, 2)
	// fmt.Println(sum)
	// fmt.Println("-----------------------------")

	sum1 := add1(2, 3, "Mostafizur Rahaman")
	fmt.Println(sum1)

	// Recive function as an parameter:
	fn := func(a int) int {
		return a * a
	}

	var processIt = func(fn func(a int) int) int {
		return fn(10)
	}

	fmt.Println(processIt(fn))

	// Return a function as parameter:
	processIt2 := func() func(a int) int {

		accumulator := 0

		return func(a int) int {
			accumulator += a
			return accumulator
		}
	}

	var accFun = processIt2()
	fmt.Println("ACC FUN 1 : ->", accFun(5))
	fmt.Println("ACC FUN 1 : ->", accFun(5))
	fmt.Println("ACC FUN 1 : ->", accFun(30))
	fmt.Println("ACC FUN 1 : ->", accFun(40))

	var accFun2 = processIt2()
	fmt.Println("ACC FUN 2 : ->", accFun2(5))
	fmt.Println("ACC FUN 2 : ->", accFun2(5))
	fmt.Println("ACC FUN 2 : ->", accFun2(30))
	fmt.Println("ACC FUN 2 : ->", accFun2(80))
}
