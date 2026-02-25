package main

// Closures -> A closure is a function that captures and retains access to variables from its surrounding scope, even after that scope has finished executing. This allows the closure to maintain state and access those variables whenever it is called.

func counter () func() int {
	count := 0

	// The anonymous function returned by counter is a closure because it captures the count variable from its surrounding scope. Each time the returned function is called, it increments and returns the value of count, allowing it to maintain state across multiple calls.
	return  func() int {
		count++
		return count
	}
}

func main() {

	increment := counter()
	println(increment()) // 1
	println(increment()) // 2
	println(increment()) // 3


}