package main

import (
	"fmt"
	"slices"
)

// Slices in Go are built on top of arrays and provide a more flexible way to work with sequences of data.
// They are dynamically sized, allowing for easy addition and removal of elements.
// Most used construct in Go.
// + usefull methods: append, copy, len, cap

func main() {

	// decleartion of slices
	// var nums []int  // nil slice -> means not initialized slice (nil is like null)

	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums))   // 0 - length


	// ****************************************************************
	// var nums2 = make([]int, 2) // initialize slice with length 2 


	// fmt.Println(nums2 == nil) // false
	// fmt.Println(len(nums2)) // 2 - length
	// fmt.Println(cap(nums2)) // 2 - capacity 

	// capacity -> Maximum number of elements can fit 

	// initial capacity is equal to length
	// var nums3 = make([]int, 2, 5) // length 2, capacity 5
	// fmt.Println(len(nums3)) // 2 - length
	// fmt.Println(cap(nums3)) // 5 - capacity

	// nums3 = append(nums3, 1)
	// fmt.Println(len(nums3)) // 3 - length 
	// fmt.Println(cap(nums3)) // 5 - capacity

	//********************* Copy function *********************
	// var num4 = make([]int, 0, 5);
	// num4 = append(num4, 54) // add element to num4

	// var num5 = make([]int, len(num4)) // create num5 with the same length as num4
	
	// copy(num5, num4) // copy elements from num4 to num5
	// fmt.Println(num4, num5) // [54] [54] - copied slice


	//********************* slice operator *********************
	// var nums = []int{1, 2, 3, 4, 5}

	// fmt.Println(nums[0:2]) // [1 2] - slice from index 0 to 2 (exclusive)
	// fmt.Println(nums[:3]) // [1, 2, 3] - slice from the beginning to index 3 (exclusive)
	// fmt.Println(nums[2:]) // [3, 4, 5] - slice from index 2 to the end

	//********************* Comparing slices *********************
	var num1 = []int {1, 2}
	var num2 = []int {1, 2}
	fmt.Println(slices.Equal(num1, num2)) // true - compares the contents of the slices

	var num3 = []int {1, 2, 3}
	var num4 = []int {1, 2, 4}
	fmt.Println(slices.Equal(num3, num4)) // false - it compares like num3 - 1 and num4 - 1 , num3- 2 and num4- 2, num3- 3 and num4- 4 - not equal


	//********************* 2D slices *********************
	var nums = [][]int {{1, 2, 3},{4, 5, 6}}
	fmt.Println(nums)

}