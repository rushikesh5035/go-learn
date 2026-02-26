package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They allow you to run functions concurrently without blocking the main thread.


func task(id int){
	fmt.Println("doing task", id)
}

func main(){
	for i:=0; i<=10; i++{
		// task(i) // block the main thread until the task is done, so the tasks will be executed sequentially.

		// go task(i) // run the task in a separate goroutine, so the tasks will be executed concurrently. non-blocking the main thread. 11 go routines will be created to run the tasks concurrently. The main thread will continue to execute the next lines of code without waiting for the tasks to finish. 
		// Output -> 10,2, 3, 0, 4, 5, 6, 7, 1, 8, 9 (the order may vary each time you run the program because the goroutines are executed concurrently and the scheduling is non-deterministic.)

		// inlinne anonymous function
		go func(i int){
			fmt.Println(i) 
		}(i) // 

	}
	
	time.Sleep(time.Second * 2) // sleep for 2 seconds to allow the goroutines to finish before the main thread exits.
}