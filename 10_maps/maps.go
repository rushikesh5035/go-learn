package main

import (
	"fmt"
	"maps"
)

// maps -> maps are a collection of key value pairs like hash, objects in js, dict in python

func main() {
	// creating a map
	// var mapName make(map[keyType]valueType) OR
	// mapName := make(map[keyType]valueType)

	m := make(map[string]string)

	// adding key value pairs to the map
	m["name"] = "goland"
	m["domain"] = "golang.org"
	m["city"] = "New York"

	// get an element from the map
	println(m["name"]) // goland
	println(m["domain"])  // golang.org
	println(m["city"]) // New York
	println(m["c"]) // empty string (if key does not exist, it returns the zero value of the value type)

	m1 := make(map[string]int)
	m1["age"] = 30
	// fmt.Println(m1["phone"]) // 0 -> key doesn't exiist

	// map Length of a map
	m1["price"] = 100
	fmt.Println(len(m1))

	// delete element from a map
	delete(m1, "age") // delete a key value pair from the map
	fmt.Println(m1)

	// clear a map
	clear(m1)
	fmt.Println(m1) // -> map[]

	// define and initialize a map in one line
	m2 := map[string] int{"price":100, "age":30} // map literal
	fmt.Println(m2)

	// check if a key exists in the map
	value, ok := m2["age"]
	fmt.Println(value) // value of the key "age"
	if ok {
		// fmt.Println("all ok")
		fmt.Println("all ok", value) // if key exists, ok will be true and value will be the value of the key
	} else {
		// fmt.Println("not ok")
		fmt.Println("not ok", value) // if key doesn't exist, ok will be false and value will be the zero value of the value type
	}


	// camparing two maps
	m3 := map[string] int{"price":100, "age":30}
	m4 := map[string] int{"price":100, "age":30}

	// fmt.Println(m3 == m4) // false -> maps cannot be compared using == operator

	fmt.Println(maps.Equal(m3, m4)) // true
}