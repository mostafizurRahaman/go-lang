package main

import "fmt"

// slice -> slice length is dynamically increased we don't have to define the length while initialize the slice.
// Most used construct in go.
// + usefull methods for slices

func main() {

	// Uninitialized sliced is nil:
	var nums []int

	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))
	fmt.Println("----------------")
	// Create slice with make(), it won't be nill
	var nums2 = make([]int, 2, 5)
	fmt.Println(nums2)
	// fmt.Println(nums2 == nil)
	fmt.Println(len(nums2))
	// capacity -> maximum number of element fit
	fmt.Println(cap(nums2))

	//
	fmt.Println("----------------")
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))
}
