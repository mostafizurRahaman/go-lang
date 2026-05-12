package main

import "fmt"

// ! Import we can declare const inside our outside of the main function:
// from any  where we can use it
// ! const are not redeclarable and reassignable.

//! Typed  const
const siteName string = "GO LANG EXPLORER"

// ! Untyped const : (without mention type we can create  untyped constant)
const platformFee = 16.70

func main() {

	// ! console the constant
	fmt.Println(siteName, platformFee)

	// ! Const (block) declaration

	const (
		name              = "Mostafizur" //	! untyped const
		role      int     = 02           // ! typed const
		height    float64 = 3.56         // ! typed  const
		isStudent                        // ! inherite the immidiate previous variable value and type  while block
	)

	fmt.Println(name, role, height, isStudent)

	// ! Multiple const declaration (Untyped)
	const a, b = 2, "Mostafiz"

	// ! Same type multiple const declaration  (typed)
	const c, d float64 = 3.34, 4.6
	fmt.Println(c, d)

}
