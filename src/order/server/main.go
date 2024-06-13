package main

import (
	"context"
	"log"
	"net"

	"github.com/glebarez/sqlite"
	pb "github.com/sgrgug/ecommerce-ms/src/order/genproto"
	"github.com/sgrgug/ecommerce-ms/src/order/model"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	db *gorm.DB
}

func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	newOrder := model.Order{
		User_Id:    req.UserId,
		Product_Id: req.ProductId,
		Quantity:   req.Quantity,
		Price:      req.Price,
	}

	if err := o.db.Create(&newOrder).Error; err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:        newOrder.ID,
			UserId:    newOrder.User_Id,
			ProductId: newOrder.Product_Id,
			Quantity:  newOrder.Quantity,
			Price:     newOrder.Price,
		},
	}, nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Order{})

	lis, err := net.Listen("tcp", ":60002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &OrderService{db: db})

	log.Printf("Order service is running on port 60002")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
