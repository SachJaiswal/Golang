package main

func main() {

	// switch statement
	// Example 1: Basic switch statement
	//we have not to right the break statement in switch case

	// var day int = 3
	// switch day {
	// case 1:
	// 	println("Monday")
	// case 2:
	// 	println("Tuesday")
	// case 3:
	// 	println("Wednesday")
	// case 4:
	// 	println("Thursday")
	// case 5:
	// 	println("Friday")
	// default:
	// 	println("Weekend")
	// }

	//multiple case in switch statement
	// switch time.Now().Weekday() {
	// 	case time.Saturday , time.Sunday:
	// 		fmt.Println("It's the weekend!")
	// 	default:
	// 		fmt.Println("It's a weekday!")
	// }

	//type switch statement
	whoamI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am an int")
		case string:
			println("I am a string")
		case bool:
			println("I am a bool")
		default:
			println("I don't know what I am")
		}
	}

	whoamI(42)      // I am an int
	whoamI("hello") // I am a string
}