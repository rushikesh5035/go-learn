package main

import (
	"fmt"
)

// GENERICS -> generics introduced in go 1.18
// Generics allow you to write functions and data structures that can work with any type, without sacrificing type safety.


func printSlice (items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice (items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we need to write separate functions for each type of slice we want to print, which can lead to code duplication


// With generics, we can write a single function that can handle slices of any type:
// The [T any] syntax defines a type parameter T that can be of any type. 
// [T any] or [T interface{}] both indicate that T can be any type.
func printGenericSlice [T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// but any is not good way to define a generic type parameter because it allows any type, which can lead to less type safety.

// this will accept int or string types
func printGenericSlice2[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// [T comparable] -> comparable contains all types

// we can pass multiple type parameters to a generic function, for example:
func printMultipleGenericSlices[T any, U any](items1 []T, items2 []U) {
	fmt.Println("Items 1:")	
	for _, item := range items1 {
		fmt.Println(item)
	}
	fmt.Println("Items 2:")
	for _, item := range items2 {
		fmt.Println(item)
	}
}

// ------------------------------------------------------------------------------
// ------------------------------------------------------------------------------

// GENRICS are also used on structs

type stack struct{
	elements []int // elements are the slice of int type
}

type stack2[T any] struct{
	elements []T // elements are the slice of T type - any type
}

func main(){
	// nums := []int {1, 2, 3, 4}
	// printSlice(nums)

	// names := []string {"golang", "typescript", "rust"}
	// printSlice(names) // This will cause a compile-time error because printSlice expects a slice of int, not string.
	// printStringSlice(names)


	// names2 := []string {"golang", "typescript", "rust"}
	// nums2 := []int {1, 2, 3, 4}

	// printGenericSlice(names2) // This will work because printGenericSlice can handle slices of any type, including string.
	// printGenericSlice(nums2)


	// printGenericSlice2(nums2)
	// printGenericSlice2(names2)

	// vals := []bool {true, false, true}
	// printGenericSlice2(vals) // this will not allow because printGenericSlice2 only accepts int or string types, not bool.

	// printMultipleGenericSlices(
	// 	[]int {1, 2, 3},
	// 	[]string {"golang", "typescript", "rust"},
	// )

	// ----------------------------------------------------------
	myStack := stack{
		elements: []int {1, 2, 3},
	}
	fmt.Println(myStack)

	myStack2 := stack2[string]{ // we define T is string type for this instance of stack2
		elements: []string {"golang", "typescript", "rust"},
	}
	fmt.Println(myStack2)

}