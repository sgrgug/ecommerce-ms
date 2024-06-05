package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sgrgug/ecommerce-ms/api/protos/product"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:60000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createProduct, err := c.CreateProduct(ctx, &pb.CreateProductRequest{
		Product: &pb.Product{
			Name:        "Product 3",
			Description: "Product 3 description",
			Price:       100,
			Quantity:    10,
			Category:    "Category 3",
		},
	})
	if err != nil {
		log.Fatalf("could not create product: %v", err)
	}
	log.Printf("Product created: %v", createProduct.Product)
}
