package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	increment := counter()

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
