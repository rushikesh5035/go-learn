package main

import "fmt"

func main() {
	// simple switch statement
	// i := 7

	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// default:
	// 	fmt.Println("Number not between 1 and 5")
	// }

	// multiple conditions switch statement

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a workday.", time.Now().Weekday())
	// }

	// Type switch statement
	whoAmI := func (i interface{}){  // empty interface means it can accept any type

		switch i.(type){
		case int:
			fmt.Println(("Its an interger"))
		
		case string:
			fmt.Println("Its a string")
		
		case bool:
			fmt.Println("Its a boolean")
		
		default:
			fmt.Println("Other")
		}
	}  

	whoAmI(23)
	whoAmI("Hello")
	whoAmI(true)
	whoAmI(2.54)
}
