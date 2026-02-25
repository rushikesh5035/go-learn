package main

import "fmt"

// variadic function -> a function that can take number of arguments

// we can use ... before the type of the parameter to indicate that it is a variadic parameter
func sum(nums ...int)int{
	total := 0
	for _, num:= range nums{
		total += num
	}
	return total
}

// we can also use ...interface{} to accept any type of arguments
func printAll(args ...interface{}){
	for _, arg := range args{
		fmt.Println(arg)
	}
}

func main() {
	// fmt.Println(1, 2, 3, 4) 
	// fmt.Println("Hello", "World", "!")
	// fmt.Println(12, "golang", true, 2.43)
	// these are all variadic functions, they can take any number of arguments of any type

	// fmt.Println(sum(1, 2, 3, 4, 5)) // 15 (we can pass no of int)
	// fmt.Println(sum(10, 20))
	printAll(1, 2, 3, 4, 5)
	printAll("Hello", "World")

	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...) // we can also pass a slice to a variadic function by using ... after the slice variable
	fmt.Println(result) // 15
}