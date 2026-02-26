package main

import "fmt"


func main() {
	// variables in go

	// string type ---> (in go if you declare a variable, you have to use it otherwise it will give an error)
	// var name string = "golang"
	// fmt.Println(name)

	// infer type ---> (you don't have to declare the type (string etc), GO will infer it)
	// (infer means = compiler will automatically determine the type of variable based on the value assigned to it)
	// var name = "GOLANG"
	// fmt.Println(name)

	// var isAdult = true
	// fmt.Println(isAdult)

	// int type
	// var age int = 25
	// var age = 25  // infer type
	// fmt.Println(age)

	// shothand declaration
	city := "New York" // when we want to declare and assign a variable in one line
	fmt.Println(city)

	marks:=90.65
	fmt.Println(marks)

	///////////////////////
	var name string // declaring variable (default value is empty string)
	name="Rushi"    // assigning value

	fmt.Println(name)

	var isAdult bool // default value is false
	isAdult=true
	fmt.Println(isAdult)

	// var price float32 = 102.50
	// var price = 102.50 // infer type
	price:= 150.75       // shorthand declaration
	fmt.Println(price)
}