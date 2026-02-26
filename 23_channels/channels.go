package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// Channels are a way to communicate between goroutines. T
// channels are like pipes, where we send data from one side and receive it on the other side.
// We can create a channel using the make function, and we can specify the type of data that the channel will carry.

// if we have single goroutine then we can use channels
// if we have multiple goroutines then we can use waitgroups to wait for all goroutines to finish before exiting the program.

// receving data to the channel
func processNum (numChannel chan int){ 
	// fmt.Println("Processing number: ", <- numChannel) // we are trying to receive data from the channel before we send it, so we will get a deadlock error.

	// for loop to receive data from the channel multiple times.
	for num := range numChannel { //range keyword to receive data from the channel until it is closed.
		fmt.Println("Processing number: ", num) // trying to receive data from the channel before we send it, so we will get a deadlock error.
		time.Sleep(time.Second * 1)
	}
}


// sending data from the channel 
func sum(result chan int, num1 int, num2 int){
	numResult:= num1 + num2

	result <- numResult // 
}

// goroutine synchronization using channels
func task(done chan bool){
	defer func (){
		done <- true // we are sending true to the channel to indicate that the task is done.
	}()

	fmt.Println("Processing....")
}


// queue system for email sending using channels
func emailSender(emailChan chan string, doneChan chan bool){
	defer func() { doneChan <- true}() // we are sending true to the channel to indicate that the email sending is done.

	for email:= range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second) // sleep for 1 sec to simulate the time taken to send an email.
	}
}	


func main(){
	// messageChannel := make(chan string) 

	// We can send data to the channel using the `<-`` operator.
	// messageChannel <- "Hello, World!" 

	// we can receive data from the channel using the same operator
	// message := <- messageChannel

	// fmt.Println(message) // we got deadlock error because we are trying to send data to the channel before we receive it. 


	// numChannel := make(chan int)

	// go processNum(numChannel)

	// for {
	// 	numChannel <- rand.Intn(100) // we are sending random data to the channel after we receive it, so we will not get a deadlock error.
	// }

	//-----------------------------------------------------------
	// result := make(chan int)

	// go sum(result, 5, 10)

	// fmt.Println("Result: ", <-result) // receiving the result from the sum function and printing it.

	//---------------------------------------------------------------
	//---------------------------------------------------------------
	// done := make(chan bool)

	// go task(done) // 
	
	// <-done //  blocking the main goroutine until we receive a signal from the task goroutine, which will be sent when the task is done. This is a way to synchronize the goroutines, ensuring that the main goroutine waits for the task to complete before exiting. 
	// receiving from the channel to indicate that the task is done.

	//------------------------------------------------------------
	
	// Queue system for email sending using channels
	// emailChan := make(chan string, 100) // buffered channel with a capacity of 100 and this is non-blocking channel, so we can send data to the channel without waiting for it to be received.

	// doneChan := make(chan bool) // unbuffered channel for synchronization

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// go emailSender(emailChan, doneChan)

	// for i := 0; i < 10; i++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i) // it generate 100 email like -> 1@gmail.com, 2@gmail.com
	// }

	// fmt.Println("Done sending emails to the channel")

	// close(emailChan) // closing the channel after we are done sending data to it. if not close, it gives deadlock error
	
	// <-doneChan

	//-------------------------------------------------------------

	// what if we need to listen to multiple channels at the same time? we can use select statement to listen to multiple channels and perform actions based on which channel receives data first.

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1 <- 10
	}()

	go func ()  {
		chan2 <- "Hello from chan2"
	}()

	for i:= 0; i < 2; i++ {
		select{
		case chan1Data := <- chan1:
			fmt.Println("Received from chan1: ", chan1Data)
		
		case chan2Data := <- chan2:
			fmt.Println("Received from chan2: ", chan2Data)
		
		default:
			fmt.Println("No data received from either channel")
			time.Sleep(time.Second) // sleep for 1 sec before checking the channels again.
		}
	}
}


// 1. Channels are a way to communicate between goroutines.
// 2. We can create a channel using the make function, and we can specify the type of data that the channel will carry.
// 3. We can send data to the channel using the `<-`` operator.
// 4. We can receive data from the channel using the same operator.
// 5. We can use the range keyword to receive data from the channel until it is closed.	
// 6. We can use the defer keyword to ensure that we send a signal to the channel when a task is done.	
// 7. We can use the close function to close a channel when we are done sending data to it.