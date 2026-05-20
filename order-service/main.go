package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type Order struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func main() {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	orderID := 101
	for {
		time.Sleep(5 * time.Second)

		order := Order{
			OrderID:   orderID,
			ProductID: 51,
			Quantity:  1,
		}

		data, err := json.Marshal(order)
		if err != nil {
			log.Printf("failed to encode order: %v", err)
			continue
		}

		if err := nc.Publish("orders.created", data); err != nil {
			log.Printf("failed to publish order: %v", err)
			continue
		}

		fmt.Printf("Order %d placed for Product %d\n", order.OrderID, order.ProductID)

		orderID++
	}
}
