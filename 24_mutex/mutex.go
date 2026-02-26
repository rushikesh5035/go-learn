package main

import (
	"fmt"
	"sync"
)

// Mutex is a mutual exclusion lock.
// Mutex is used to prevent race conditions in concurrent programming.
// when the multiple process try to access the same resource at the same time, then it can cause a race condition. To prevent this, we can use a mutex to lock the resource while it is being accessed by one process, and unlock it when the process is done.

// social media post struct
type posts struct{
	views int

	mu sync.Mutex // mutex for posts struct
}

// 
func (p *posts) inc(wg *sync.WaitGroup){
	// defer wg.Done() // goroutine is done

	defer func ()  {
		p.mu.Unlock()
		wg.Done()
	}()


	p.mu.Lock() // lock the mutex before accessing the views

	p.views++

	// p.mu.Unlock() // unlock the mutex after accessing the views
	// good way to unlock the mutex and mark the goroutine as done is to use defer statement, so we can ensure that the mutex will be unlocked and the goroutine will be marked as done even if there is an error in function. 
}


func main(){

	var wg sync.WaitGroup

	// instance of posts struct
	myPost := posts{views: 0}

	// myPost.inc() // increment views by 1
	// myPost.inc() // increment views by 1
	
	for i :=0; i<100; i++{
		// myPost.inc() // synchronous incr views by 1

		// to make it asynchronous, we can use goroutine
		wg.Add(1)

		go myPost.inc(&wg) //concurrently/asynchronous incr views by 1
	}

	wg.Wait() // wait for all goroutines to finish
	// fmt.Println(myPost.views) // output will be 100, 95, 93 or any number less than 100 because of race condition

	// to prevent race condition, we can use mutex -> we create a mutex for posts struct so the good way is to create mutex inside the struct 

	fmt.Println(myPost.views) // output will be 100 because of mutex

}