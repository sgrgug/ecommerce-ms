package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sgrgug/ecommerce-ms/src/order/genproto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:60002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createOrder, err := c.CreateOrder(ctx, &pb.CreateOrderRequest{
		UserId:    "3",
		ProductId: "5",
		Quantity:  55,
		Price:     200,
	})

	if err != nil {
		log.Fatalf("could not create product: %v", err)
	}
	log.Printf("Order created: %v", createOrder.Order)

}
