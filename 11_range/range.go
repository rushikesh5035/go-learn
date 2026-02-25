package main

import (
	"fmt"
)

// range -> iterating over data structures (arrays, slices, maps, strings)

func main(){
	//******************* range with slices *********************
	// nums := []int{6, 7, 8}

	// for i:= 0; i< len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	// range returns two values, the index and the value
	// for _, num := range nums{ 
	// 	fmt.Println(num)
	// }

	// perform sum operation on the slice
	// sum := 0
	// for _, num := range nums{
	// 	sum += num
	// }
	// fmt.Println("sum is", sum)

	// element with index
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	//********************* range with maps *********************
	// m := map[string] string {
	// 	"fname"	: "John",
	// 	"lname" : "Doe",
	// }

	// for k, v := range m{
	// 	fmt.Println(k, v) // k - key, v - value
	// }

	//********************* range with strings ******************
	for i, c := range "golang" {
		// fmt.Println(i, c) // i - index, c - character (unicode code point)

		fmt.Println(i, string(c)) // convert unicode code point to string
	}
}