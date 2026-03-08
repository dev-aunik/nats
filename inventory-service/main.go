package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type Order struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	inventory := map[int]int{
		51: 10, // Start with 10 items of product 51
	}

	nc.Subscribe("orders.created", func(msg *nats.Msg) {
		var order Order
		json.Unmarshal(msg.Data, &order)
		fmt.Printf("Inventory received order: %d\n", order.OrderID)

		if stock, exists := inventory[order.ProductID]; exists && stock >= order.Quantity {
			inventory[order.ProductID] -= order.Quantity
			fmt.Printf("Reducing stock for product: %d. Remaining stock: %d\n", order.ProductID, inventory[order.ProductID])
		} else {
			fmt.Printf("Failed to fulfill order %d: Product %d is out of stock!\n", order.OrderID, order.ProductID)
		}
		fmt.Println("--------------------------------------------------")
	})

	fmt.Println("Inventory service listening...")

	select {}
}
