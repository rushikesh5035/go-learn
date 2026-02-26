package main

import "fmt"

// ENUMS -> enumerated types
// ENUMS are a way to define a set of named constants. They can be used to represent a fixed number of values, such as the status of an order. "pending", "shipped", "delivered" etc
// In Go, we can use the iota keyword to create an enum-like type.

// We can define an enum for order status using integer values
type OrderStatus int

// 
const (
	Received OrderStatus= iota // iota is a special keyword that is used to create a sequence of constants. It starts at 0 and increments by 1 for each constant. So, Received will be 0, Pending will be 1, Shipped will be 2, and Delivered will be 3.
	Pending
	Shipped
	Delivered
)

// We can also define an enum for order status using string values
type OrderStatus2 string

const (
	Received2 OrderStatus2 = "received"
	Confirmed2 			   = "confirmed"
	Shipped2 			   = "shipped"
	Delivered2 			   = "delivered"
)


func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}
	
func changeOrderStatus2(status OrderStatus2) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Received) // o/p: Changing order status to 0
	changeOrderStatus(Pending) // 1
	changeOrderStatus(Shipped) // 2
	changeOrderStatus(Delivered) // 3

	changeOrderStatus2(Received2) // o/p: received
	changeOrderStatus2(Confirmed2) // o/p: confirmed
	changeOrderStatus2(Shipped2) // o/p: shipped
	changeOrderStatus2(Delivered2) // o/p: delivered
}