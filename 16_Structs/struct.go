package main

// order struct
// type Order struct {
// 	id        string
// 	amount    float64
// 	status    string
// 	createdAt time.Time //nanosecond precision
// }

// // * Order is a pointer to the Order struct usd when we need to change the value of the struct
// func (O *Order) changeStatus(status string) {
// 	O.status = status
// 	fmt.Println("Inside changeStatus:", status) // 10
// }

// func (O Order) getStatus() string {
// 	return O.status
// }

// func (O Order) getAmount() float64 {
// 	return O.amount
// }

func main() {
	// order := Order{
	// 	id:        "12345",
	// 	amount:    100.50,
	// 	status:    "Pending",
	// 	createdAt: time.Now(),
	// }

	// order.changeStatus("Delivered")
	// fmt.Println(order.status)
	// fmt.Println(order.getStatus())
	// fmt.Println(order.getAmount())
	// // fmt.Println("Order ID:", order.ID)
	// // fmt.Println("Order Amount:", order.Amount)
	// // fmt.Println("Order Status:", order.status)
	// // fmt.Println("Order Created At:", order.createdAt)

	// language := struct {
	// 	name   string
	// 	isgood bool
	// }{"Go", true}
	// fmt.Println(language) // Go

}
