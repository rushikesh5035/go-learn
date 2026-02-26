package main

import (
	"fmt"
	"time"
)

// Struct embedding is a way to include one struct within another struct.
// It allows us to reuse the fields and methods of the embedded struct in the outer struct.
// The embedded struct is also known as the "anonymous field" because it does not have a name.

type customer struct {
	name  string
	phone string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {

	newCustomer := customer{
		name:  "John",
		phone: "1234567890",
	}

	newOrder := Order{
		id:		"1",
		amount: 5435.0,
		status: "receive",
		// customer: newCustomer, or
		customer: customer{ // inline struct literal
			name:"John",
			phone: "1234567890",
		},
	}

	fmt.Println(newOrder) // {1 5435 receive {0 0 <nil>} { }} -> {} this is customer struct
	fmt.Println(newOrder.customer) // { } 


	fmt.Println(newCustomer.name)

	newCustomer.name = "Doe" // changing the name of the customer
	fmt.Println(newCustomer.name) // Doe
}