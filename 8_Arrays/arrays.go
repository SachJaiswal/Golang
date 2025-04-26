package main

import "fmt"

//arrays in go ia nu mbered sequence of elements of the same type
// arrays are fixed in size and cannot be resized
func main() {

	// int -> 0, string->"", bool->false

	// var nums [10]int

	// fmt.Println(len(nums))

	// var nums[10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	// fmt.Println(nums[9])

	var names [5]string = [5]string{"Sachin", "Jaiswal"}
	fmt.Println(names)

	// PS D:\GO_Tutorial> go run "d:\GO_Tutorial\8_Arrays\arrays.go"
	// // [Sachin Jaiswal   ]

	// declare in single line
	nums := [3]int{1, 2, 3}
	fmt.Println(nums[0])

	//2d array
	var num [3][2]int = [3][2]int{
		{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(num) //6

	//Application of arrays
	// -fixed size , that is predictale
	//-memory optimization
	//-costant time access to elements

}
