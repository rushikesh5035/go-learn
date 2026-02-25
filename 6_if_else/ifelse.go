package main

import "fmt"

func main() {
	// if else condition
	// age := 10

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// else if condition
	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	// 	fmt.Println("Person is a child")
	// }

	// logical operators in if condition
	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("Access granted")
	}

	// we can declare variable inside if statement
	if  age:= 11; age >= 18 {
		fmt.Println("Person is an adult with age:", age)
	} else if age >= 12 {
		fmt.Println("Person is a teenager with age:", age)
	} else {
		fmt.Println("Person is a child with age:", age)
	}

	// go does not have ternary operator,so you will have to use if else
	
}