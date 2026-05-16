package main

import "fmt"

// ! A generics which will receive int and float32 value
func sum[T int | float32](nums []T) T {

	var counter T = 0

	for _, i := range nums {
		counter = counter + i
	}

	return counter

}

// ! A  function with any type : any and inferface{} doing same work:
func renderAnyData[T any](v []T) {
	for _, v := range v {
		fmt.Println("Value of ", v)
	}
}

// ! A function with comparable type:
func renderComparableData[T comparable](v []T, k T) int {

	for i, val := range v {
		if k == val {
			return i
		}

	}

	return -1

}

// ! Create a function which will identify duplicate value in slice :

func removeDuplicateValFromSlice[T comparable](slice []T) []T {
	seen := map[T]bool{}
	result := []T{}

	for _, v := range slice {

		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}

	}

	return result

}

// Create an function which identify struct matched with another struct:
// ! But the struct values should only be comparable supported
func FindIndex[T comparable](list []T, target T) int {

	for i, v := range list {
		if v == target {

			return i

		}
	}
	return -1

}

// # Custom type with interface
type customType interface {
	int | float32 | float64
}

func identifyLargeNumber[T customType](a, b T) T {

	if a > b {
		return a
	} else {
		return b
	}

}

func main() {

	// ! Generics type:
	nums := []int{1, 2, 3, 4}
	nums2 := []int{34, 343, 30}
	decimalNums := []float32{3.4, 5.6, 60.3}
	names := []string{"Mostafiz", "Rakib", "Fahim"}
	fmt.Println(sum(nums))
	fmt.Println(sum(decimalNums))
	fmt.Println(sum(nums2))

	// ! In Go lang, To recieve multiple all values, go has provided two types
	// ?  any, comparable
	// ? any can receive any type of data like string, float32, float62, int, int8, int16, int32, int64, bool

	renderAnyData(nums)
	renderAnyData(nums2)
	renderAnyData(names)

	fmt.Println("--- Compareable ------", renderComparableData(nums, 3))
	fmt.Println("--- Compareable ------", renderComparableData(nums2, 3))
	fmt.Println("--- Compareable ------", renderComparableData(names, "Mostafiz"))

	// ! Filter duplicate value from different types of slice:

	// String
	fmt.Println(removeDuplicateValFromSlice([]string{"Mostafiz", "Ratul", "Hridoy", "Mostafiz", "Hridoy"}))
	fmt.Println(removeDuplicateValFromSlice([]int{20, 20, 40, 50, 50, 60, 50}))

	// ! A comparable type value diye jodi struct make kora hoy tahole sekhane of == kaj kore:
	type order struct {
		id          int
		productName string
		price       float32
	}

	// ! Custome type for slice of struct: 
	type TOrderList []order

	orderList := TOrderList{
		{
			id:          1,
			productName: "Mostafiz",
			price:       3.5,
		},
		{
			id:          2,
			productName: "Ratul",
			price:       30.3,
		},
		{
			id:          3,
			productName: "Fhaim",
			price:       350,
		},
	}

	fmt.Println(FindIndex(orderList, order{
		id:          2,
		productName: "Ratul",
		price:       30.3,
	}))

	fmt.Println(FindIndex(orderList, order{
		id:          3,
		productName: "Ratul",
		price:       30.3,
	}))

	fmt.Println(identifyLargeNumber(20, 30))
	fmt.Println(identifyLargeNumber(50, 30))

}
