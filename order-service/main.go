package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	orderID := 101
	for {
		time.Sleep(5 * time.Second)

		order := Order{
			OrderID:   orderID,
			ProductID: 51,
			Quantity:  1,
		}

		data, _ := json.Marshal(order)

		nc.Publish("orders.created", data)

		fmt.Printf("Order %d placed for Product %d\n", order.OrderID, order.ProductID)

		orderID++
	}
}
