package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// If parameters are some we can write a time : 

func add1 ( a, b int) int { 
	return a + b 
}

func getLanguages()(string, string, bool) { 
	return "golang", "javascript", true
}

func processIt (fn func(a int) int)  {

	fn(2)
}


func processIt2 ( ) func(a int) int { 
	 acc := 0
	return func (a int) int { 
		// fmt.Println(a)
		return acc + a
	}
}





func main() {

	result := add(1, 3)
	fmt.Println(result)

	result1 := add1(1, 5)
	fmt.Println(result1)

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	fn:= func(a int) int{ 
		return a
	}

	processIt(fn)

	fn1 := processIt2()
	 b := fn1(5)
	 c := fn1(5)
	 d := fn1(5)

	
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	


}