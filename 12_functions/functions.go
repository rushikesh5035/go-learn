package main

import "fmt"

// This is a simple function that adds two integers and returns the result as an integer.
// func add(a int, b int) int {
// 	return a + b
// }

// if both parameters are of the same type, we can shorten the syntax by only specifying the type once:
func add(a,b int) int {
	return a + b
}

func getLanguages() (string, string, string) { // func return three string values, so we specify three string types in the func
	return "golang", "python", "javascript" 
}

// func return three different types of values, so we specify the type of each value in the func
func getLanguages2() (int, string, bool) { 
	return 2024, "golang", true 
}

// return a func from a func
func processIt(fn func(a int, b string)(int, string)){
	fn(1, "hello") // we can call the
}

func processIt2() func(a int)int {
	return func(a int)int {
		return a + 1
	}
}

// this is also the func
func main() {
	// result := add(3, 5)
	// fmt.Println(result)
	// fmt.Println(getLanguages())
	
	// fmt.Println(getLanguages2())
	// lang1, lang2, _ := getLanguages() // if we only want to use the first two values returned by the func, we can use the blank identifier _ to ignore the third value
	// fmt.Println(lang1, lang2)


	fn := func (a int, b string)(int, string)  {
		return a + 1, b + " processed"
	}
	processIt(fn)

	fn2 := processIt2() // fn2 is now a func that takes an int and returns an int
	fmt.Println(fn2(5)) // 6
}
