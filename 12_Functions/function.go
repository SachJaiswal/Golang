package main

import (
	"fmt"
	"time"
)

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguage() (string, string, string) {
// 	return "Go", "Python", "Java"
// }

func sayHello() {
	fmt.Println("Hello World")
}

func main() {

	// res := add(1, 2)
	// fmt.Println(res)

	// lang1, lang2, lang3 := getLanguage()
	// fmt.Println(lang1, lang2, lang3)

	// // surpass any value if you dont need it
	// _, lang2, _ = getLanguage()
	// fmt.Println(lang2)

	go sayHello()
	fmt.Println("Main Func finished")
	time.Sleep(1 * time.Second)

}
