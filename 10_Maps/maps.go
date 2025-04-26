package main

import "fmt"

//maps-> key-value pair
// unordered collection of key-value pairs
// key-> unique identifier for each value
// value-> data associated with the key

func main() {

	//creating map

	m := make(map[string]string)
	m["name"] = "John"
	m["age"] = "30"
	m["city"] = "New York"
	m["country"] = "USA"

	fmt.Println(m) //map[age:30 city:New York country:USA name:John]
	//unordered collection of key-value pairs

	fmt.Println((m["state"])) //empty string

	// m1 := make(map[string]int)
	// m1["a"] = 1
	// fmt.Println((m1["b"]))   //0

	//delete key-value pair from map
	// delete(m, "age")  //map[city:New York country:USA name:John]
	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m) //map[]

	//without make crating map
	m2 := map[string]int{"price": 100, "price2": 200}
	fmt.Println(m2) //map[price:100 price2:200

	//iterating over map
	for key, value := range m {
		fmt.Println(key, value) //name John age 30 city New York country USA
	}

	//checking if key exists in map
	if value, ok := m2["price"]; ok {
		fmt.Println("Key exists in map", value) //Key exists in map 100
	} else {
		fmt.Println("Key does not exist in map")
	}
}
