package main

import "fmt"

func rePrint  (num *int)  int { 
	fmt.Println("Print before", *num)
	*num = 20
	fmt.Println("print value after change")
	return *num
}

func main() {
	fmt.Println("Welcome to a class on pointers")
	// i, j := 21, 3232

	// p := &i
	// fmt.Println(p)

	var num = 5 
	fmt.Println(num)

	rePrint(&num)
	fmt.Println(num)
	// num = 4 
	// fmt.Println(num)


	parent := func() func() int {

		counter:= 0

		return func () int {  // increment
			counter = counter  +1
			
			return counter
		}
	}
	a := parent()
	b := parent()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("B", b())
	fmt.Println("B", b())
	fmt.Println("B", b())


	fmt.Println("----------Address sharing")
	abc := 10
	fmt.Println(abc)
	ptr  := &abc
	fmt.Println(ptr)
	*ptr = 100
	fmt.Println(abc)
}