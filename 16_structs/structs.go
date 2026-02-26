package main

import (
	"fmt"
	"time"
)

// Structs are a way to group together related data.
// structs is similar to classes in other programming languages
// but they do not have methods or inheritance.

// order struct
type Order struct{
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
}

// constructor function to create a new order struct
// *Order - type of  newOrder function, it returns a pointer to an Order struct, able to change struct value
func newOrder (id string, amount float32, status string) *Order {
	// initial setup of the struct fields
	myOrder := Order{
	id: id,
	amount: amount,
	status: status,
	createdAt: time.Now(),
	}

	return &myOrder // returning a pointer to the struct.
}


// method to change the status of an order 
// (o *Order) is a pointer receiver, which means that the method can modify the original struct 
// if we use (o Order) instead of (o *Order), the method will receive a copy of the struct and any changes made to it will not affect the original struct.
func (o * Order)changeStatus(status string){
	o.status = status 
}

// method to get the amount of an order
// * is only used when we want to modify the struct value, 
// if we just want to read the value we can use (o Order) instead of (o *Order)
func (o Order) getAmount() float32{
	return o.amount
}

func main() {
	

	// instanceating a struct 
	order:= Order{
		id: "1",
		amount: 5435.0,
		status: "pending",
	}

	// setting values in a struct
	// order.createdAt = time.Now()
	// order.updatedAt = time.Now()

	// fmt.Println("order structs", order)
	// fmt.Println(order.id)

	// if you don't set any field, default value is zero value
	// int -> 0, float -> 0, string -> "", bool -> false
	myOrder2 := Order{
		id: "2",
		// amount:1200.0, // if value is not set, default value will be 0 for float32
		status:"delivered",
		createdAt: time.Now(),
	}

	order.status = "paid" // not affecting to myOrder 2 struct
	// fmt.Println("Order", order)
	
	fmt.Println("myOrder2", myOrder2)

	
	// order.changeStatus("shipped") // calling the method on the order struct
	// fmt.Println("Order after changeStatus method", order)


	// order.getAmount() // calling method on the order struct
	// fmt.Println("Order amount: ", order.getAmount())



	myOrder := newOrder("1", 5435.0, "pending") // using constructor function to create a new order struct
	fmt.Println("myOrder", myOrder)


	// inline struct - struct without a name
	language := struct {
		name string
		isGood bool
	} {"golang",true}

	fmt.Println(language)

	
}