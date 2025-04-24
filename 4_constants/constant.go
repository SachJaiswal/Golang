package main

import "fmt"

const (
	port=8080 
	host="Loclhost"
)

func main() {
	const name string = "golang" // constant declaration
	const age int = 20           // constant declaration
	age1 := 20                   // constant declaration

	fmt.Println(age1) // prints the value of name

	fmt.Println(port)
}