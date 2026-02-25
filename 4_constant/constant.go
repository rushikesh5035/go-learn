package main

import "fmt"

const age = 30 // you can also declare constants outside functions

const name = "Alice" // ✅
// name := "Alice" // this will cause an error, := can only be used inside functions ❌

func main() {
	// constant values - cannot be changed

	const pi float32 = 3.14159
	const language string = "GoLang"
	const year int = 2024

	// pi = 3.53 // This will cause a compile-time error

	fmt.Println("Constant Pi:", pi)
	fmt.Println("Programming Language:", language)
	fmt.Println("Year:", year)

	// accessing constants declared inside function
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// const grouping -> multiple constants declared together
	const (
		port  = 5000
		host  = "localhost"
		debug = true
	)

	fmt.Println(port, host)
}