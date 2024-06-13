package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sgrgug/ecommerce-ms/src/user/genproto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:60001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createUser, err := c.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Username: "User 4",
			Email:    "user4@gmail.com",
			Password: "password",
		},
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("User created: %v", createUser.User)
}
