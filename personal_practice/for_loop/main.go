package main

import "fmt"

// ! Only construct for looping in go:

func main() {

	// !  1. for loop as while loop (finite):

	fmt.Println("-----------For as  while (finite) ------------")
	i := 0
	for i <= 5 {

		fmt.Printf("The value is : %v \n", i)

		// increase the value:
		i++
	}

	// ! 2. for loop as while loop (infinite):

	// fmt.Println("-----------For as while infinite ------------")
	// j := 0
	// for {
	// 	fmt.Println(j)
	// 	j++
	// }

	// ! 3. For as clasic  for loop:

	fmt.Println("======= CLASIC FOR LOOP =========")
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	// ! 4. For loop with continue key:
	fmt.Println("======= CLASIC FOR LOOP With continue key =========")
	for i := 0; i <= 5; i++ {

		if i == 2 {
			continue //! skip the if matched
		}

		fmt.Println(i)
	}

	// ! 5. For loop with break key:
	fmt.Println("======= CLASIC FOR LOOP With break key =========")
	for i := 0; i <= 5; i++ {

		if i == 4 {
			break // ! stop the loop here if matched.
		}

		fmt.Println(i)
	}

	// ! 6. for loop with range keyword:
	fmt.Println("======= For loop  with range keyword =========")
	for i := range 5 {
		fmt.Println(i)
	}

	// ! 7. for loop with string and range :
	fmt.Println("======= For loop  with range keyword for iterating string =========")
	str := "Mostafizur"
	for i, v := range str {
		fmt.Println("Index : ", i, "Unicode value :  ", v, "Character :", string(v))
	}

}
