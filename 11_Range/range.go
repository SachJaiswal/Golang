package main

import "fmt"

//iterating over datastructure using range keyword
func main() {
	// #1. iterating over array
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println(num)
	}
	fmt.Println("Sum of the numbers is ", sum)

	for index, num := range nums {
		fmt.Println(index, " = ", num)
	}

	//#2. iterating over string

	m1 := map[string]string{"name": "John", "age": "30", "city": "New York"}

	for key, value := range m1 {
		fmt.Println(key, " = ", value)
	}

	// #3 . iterating over stringfmt
	// unicode code point
	for i, c := range "golangz" {
		fmt.Println(i, " = ", c) //0  =  104 1  =  101 2  =  108 3  =  108 4  =  111
	}

}
