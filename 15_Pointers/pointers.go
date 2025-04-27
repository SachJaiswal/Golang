package main

import "fmt"

//by value
// func changeNumber(n int) {
// 	n = 10
// 	fmt.Println("Inside changeNumber:", n) // 10
// }

//by reference
func changeNumber(n *int) {
	*n = 10
	fmt.Println("Inside changeNumber:", *n) // 10
}

func main() {
	n := 5
	changeNumber(&n)
	//fmt.Println("Memory address of n:", &n) // Memory address of n
	fmt.Println("After Change Inside main:", n) // 5
}
