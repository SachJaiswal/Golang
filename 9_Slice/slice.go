package main

import (
	"fmt"
	"slices"
)

// slice-> dynamic array
// most used data structure in go
// useful methods
func main() {

	// uninitialize slice is nil
	var nums []int
	fmt.Println(nums == nil) //true
	fmt.Println(len(nums))   //0

	var nums1 = make([]int, 0, 5)

	//capacity => maximun numbers of element can fit in the slice
	fmt.Println(cap(nums1)) //2
	fmt.Println(nums1)      //[0 0]
	nums1 = append(nums1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(nums1)
	fmt.Println("Length of the slice == ", len(nums1)) //12
	fmt.Println(cap(nums1))                            //10

	var num1 = make([]int, 0, 5)
	num1 = append(num1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	nums2 := make([]int, len(num1))
	//copy function
	copy(nums2, nums1)
	fmt.Println(nums2) // [1 2 3 4 5 6 7 8 9 10 11 12]

	var nums3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums3[0:5])

	//Slices Equal function
	var n1 = []int{1, 2}
	var n2 = []int{1, 2}

	fmt.Println(slices.Equal(n1, n2)) // true

	//2d slice
	var nums4 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(nums4) // 1
}
