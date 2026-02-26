package main

import "sync"

// in the previous example, we uses the Sleep to wait for the goroutines to finish, but this is not a good way to do it, because we don't know how long the goroutines will take to finish, and if we set the Sleep time too short, the program will exit before the goroutines finish, and if we set the Sleep time too long, the program will wait unnecessarily long.
// to solve this problem, we can use a WaitGroup from the sync package, which allows us to wait for a collection of goroutines to finish. A WaitGroup is a counter that we can increment and decrement, and we can wait for the counter to reach zero before proceeding.

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // signal that the task is done after the execution of the function
	println("Doing Task", id)
}

func main() {

	// create wait group
	var wg sync.WaitGroup // 

	for i:= 0; i < 5; i++ {
		wg.Add(1) // increment the wait group counter by 1 for each goroutine we are going to start
		go task(i, &wg)
	}

	wg.Wait() // wait for all the goroutines to finish, this will block the main goroutine until the counter is zero
}


// 1st we create a wait group variable wg of type sync.WaitGroup, which will be used to keep track of the number of goroutines that are running.
// 2nd we start a loop that will run 5 times, and in each iteration, we call wg.Add(1) to increment the wait group counter by 1, indicating that we are starting a new goroutine.
// 3rd we start a new goroutine by calling go task(i, &wg), which will execute the task function with the current value of i and a pointer to the wait group variable.
// 4th inside the task function, we call defer w.Done() to signal that the task is done after the execution of the function. This will decrement the wait group counter by 1 when the function returns.
// 5th after starting all the goroutines, we call wg.Wait() to wait for all the goroutines to finish. This will block the main goroutine until the counter is zero, which means that all the tasks have completed.