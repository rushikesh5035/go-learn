package main

import "fmt"

// for loop -> in go there is only one looping construct i.e. for loop (not while, do-while etc.)

// syntax
// for initialization; condition; post {
//     // code to be executed
// }

func main(){

	// while loop implement using for loop 
	// i:=1
	// for i<= 4 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	// for i := 0; i < 5; i++ {
	// 	// break // (to exit from loop)
	// 	if i == 2 {  // skip when i is 2
	// 		continue  // to skip current iteration and move to next iteration
	// 	} 
	// 	fmt.Println(i)
	// }

	// range -> it literate from 0 to n-1 )
	for i := range 3 { // (it excludes 3 and print from 0 to 2)
		fmt.Println(i)
	}

}