package main

import "fmt"

//A variadic function is a function that can accept any number of arguments of a specific type.

func sum(numbers ...int) int {

	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}
func main() {

	fmt.Println("✅ Instead of fixing the number of parameters,\nYou use ... (three dots) before the type.\nIt collects all passed arguments into a slice.")

	num := []int{1, 2, 3, 4, 5, 10}
	res := sum(num...)
	fmt.Println("sum==", res)

}
