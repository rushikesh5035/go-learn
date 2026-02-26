package main

import "fmt"

// Arrays -> arrays is a fixed-size collection of elements of the same type.

// Syntax: var arrayName [size]type

func main(){

	// declare an array

	// ************** Integer array **************
	// var nums [10]int  // default value is 0

	// length of the array
	// fmt.Println(len(nums)) 

	// assign value in array
	// nums[0] = 12

	// fmt.Println(nums[0]) // access first element
	// fmt.Println(nums) // print entire array

	// ************** Boolean array **************
	// var vals[4]bool  // default value is false

	// assign values
	// vals[0] = true
	// vals[1] = true

	// fmt.Println(vals)


	// ************** String array **************
	// var str[3]string  // default value is ""

	// str[0] = "Hello"
	// str[1] = "World"
	// str[2] = "World"

	// fmt.Println(str)

	// declare and assign value to array in single line
	// arrayName := [size]type{value1, value2, ...}

	nums := [3]int{1,2,3}
	fmt.Println(nums)


	// ************** 2D Array **************
	// arrayName := [rows][columns]type{ {row1_values}, {row2_values}, ...}
	
	nums2D := [2][2] int {{1, 2},{3, 4}}
	fmt.Println(nums2D)

	// Advantages of arrays:
	// -> fixed size , that is predictable
	// - Memory optimization
	// - Constant time access to elements using index
}