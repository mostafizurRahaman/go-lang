package main

import (
	"fmt"
	"slices"
)

// slices => slices
// most used construct in go
// + useful methods

// What we learnt from here:
// 1. Slice declaration types (3 way to declare)
// 2. make function
// 3. access slice element by using index
// 4. slice len()
// 5. slice cap()
// 6. Slice appending value
// 7. Slice copying
// 8. Slice operator [startIndex:endIndex]
// 	--		8.1. [startIndex:]
// 	--		8.2. [:endIndex]
// 9. slice

func main() {

	// 1. First declaration way:
	var nums []int

	// uninitialized slice is nil
	fmt.Println(nums == nil)
	fmt.Println(nums)
	// Append  element on this slice:
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)
	nums = append(nums, 8)
	nums = append(nums, 9)
	fmt.Println(nums[len(nums)-1])

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// 2. If you can an slice with make(), it will not be equal to  nil
	var friendList = make([]string, 0)
	fmt.Println(friendList)
	fmt.Println(friendList == nil)
	fmt.Println(len(friendList))
	fmt.Println("capacity", cap(friendList))

	friendList = append(friendList, "Mostafiz")
	friendList = append(friendList, "Mostafiz")
	friendList = append(friendList, "Mostafiz")
	friendList = append(friendList, "Mostafiz")
	friendList = append(friendList, "Mostafiz")
	fmt.Println(friendList)
	fmt.Println(friendList[3])
	fmt.Println("length", len(friendList))
	fmt.Println("capacity", cap(friendList))

	// 3.Third way to delcare a slice:
	coligueList := []string{}
	fmt.Println(coligueList)
	fmt.Println("Length : ", len(coligueList))
	fmt.Println("cap : ", cap(coligueList))
	fmt.Println("Is Nil", coligueList == nil)

	coligueList = append(coligueList, "NM Sujon")
	coligueList = append(coligueList, "Osman goni")
	coligueList = append(coligueList, "Osman goni")
	coligueList = append(coligueList, "Osman goni")
	fmt.Println(coligueList)
	fmt.Println("Length : ", len(coligueList))
	fmt.Println("cap : ", cap(coligueList))

	// 4:  Make an slice from an array:
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("ARR", arr)

	// syntax: [startIndex:endIndex]
	// endIndex is not counted
	arrSlice := arr[2:5]
	fmt.Println("arrSlice", arrSlice)

	// startIndex:
	// return full startIndex to last index  element
	arrSlice2 := arr[2:]
	fmt.Println("arrSlice", arrSlice2)

	// :endIndex
	// return from 0 index  to before endIndex
	arrSlice3 := arr[:8]
	fmt.Println(arrSlice3)

	// Copy slices:
	nums2 := []int{}

	nums2 = append(nums2, 2)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 2)
	nums3 := make([]int, len(nums2))
	nums2 = append(nums2, 8)
	fmt.Println(nums2)
	fmt.Println(nums3)

	copy(nums3, nums2)

	fmt.Println(nums2)
	fmt.Println(nums3)

	// slice operator:

	newSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(newSlice[1:4])
	fmt.Println(newSlice[:4])
	fmt.Println(newSlice[2:])

	// Slices packages:

	// compare slices:
	var numbers1 = []int{1, 2, 3}
	var numbers2 = []int{1, 2, 3}
	var numbers3 = []int{1, 2, 3, 4, 6, 5}
	// var numbers3 = []int{1, 2, 3, 4}

	// check is any values is not slice or not:
	fmt.Println(slices.Index(numbers1, 4))

	// compare two slices: Equal()
	fmt.Println(slices.Equal(numbers1, numbers2))
	fmt.Println(slices.Equal(numbers1, numbers3))

	// to check is sorted?:
	fmt.Println(slices.IsSorted(numbers3))

	// 2D  array :

	var newNums2 = [][]string{{"A", "B"}, {"C", "D"}}
	fmt.Println(newNums2)

	for i := 0; i < len(newNums2); i++ {
		for j := 0; j < len(newNums2[i]); j++ {

			fmt.Println(i, "x", j, " : ", newNums2[i][j])
		}
	}

	// print like matrics
	for i := 0; i < len(newNums2); i++ {

		for j := 0; j < len(newNums2[i]); j++ {

			fmt.Print(newNums2[i][j], " ")
		}

		fmt.Println()
	}

}
