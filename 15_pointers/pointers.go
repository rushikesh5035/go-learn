package main

import "fmt"

// Pointers are variables that store the memory address of another variable.
// They allow you to indirectly access and manipulate the value stored at that memory address.
// In Go, you can declare a pointer using the * operator, and you can get the address of a variable using the & operator.
// For example, if you have a variable x, you can declare a pointer to it as follows:
// var p *int = &x

// ------------------------------------------------------------------------------------------------------------------------------------------------------

// In the example below, we have a function changeNum that takes an integer as an argument. When we call this function with num, it creates a copy of num and changes the value to 5. However, this does not affect the original num variable in the main function, which remains 1.
// by value, the function receives a copy of the variable, and any changes made to it do not affect the original variable outside the function. To modify the original variable, we would need to use pointers.
// func changeNum(num int){
// 	num = 5 //
// 	fmt.Println("In change num", num)
// }

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

// by reference - In the example below, we have a function changeNum2 that takes a pointer to an integer as an argument. When we call this function with the memory address of num, it allows us to change the value at that memory address. As a result, the original num variable in the main function is modified to 5.
func changeNum2(num *int){
	*num = 5 // dereference the pointer to change the value at the memory address
	fmt.Println("In change num2 =>", *num)
}

func main() {
	num := 1

	// changeNum(num) 

	// fmt.Println("Memory address", &num) // print memory address of num variable. & is used to get memory address
	// fmt.Println("after change num in main", num)

	changeNum2(&num) // pass the memory address of num variable to changeNum2 function

	fmt.Println("after change num2 in main =>", num)
}