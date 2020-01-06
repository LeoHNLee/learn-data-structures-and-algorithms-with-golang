package main

import "fmt"

// Queue Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// New method initializes with Order with priority, quantity, product, customerName
func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

// Add method adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

// main function
func main() {
	var queue Queue
	queue = make(Queue, 0)
	var order1 *Order = &Order{}
	priority1 := 2
	quantity1 := 20
	product1 := "Computer"
	customerName1 := "Leohn"
	order1.New(priority1, quantity1, product1, customerName1)
	var order2 *Order = &Order{}
	priority2 := 1
	quantity2 := 10
	product2 := "Monitor"
	customerName2 := "Lee"
	order2.New(priority2, quantity2, product2, customerName2)
	queue.Add(order1)
	queue.Add(order2)
	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
