package main

import "fmt"

func main() {

	// ! 1.  Pointers first example:
	// ! &variableName for dereferencing. which will return address
	// !  & ampersand
	// ! &variableName = variable address

	// *Pointer * always come first for type and variable *typeName *varName

	num := 10

	// * Store number pointer into ptr :
	ptr := &num

	fmt.Println(ptr)

	*ptr = 20

	fmt.Println(num)

	a := 20

	pointer := &a

	*pointer = 500

	fmt.Println(a)

	var ptr2 *int = pointer
	fmt.Println(ptr2)

	*ptr2 = 400
	fmt.Println(a)
}
