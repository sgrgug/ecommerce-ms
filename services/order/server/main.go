package main

import (
	"context"
	"log"
	"net"

	"github.com/glebarez/sqlite"
	pb "github.com/sgrgug/ecommerce-ms/api/protos/order"
	"github.com/sgrgug/ecommerce-ms/services/order/model"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	db *gorm.DB
}

func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	newOrder := model.Order{
		User_Id:    uint(req.UserId),
		Product_Id: uint(req.ProductId),
		Quantity:   req.Quantity,
		Price:      req.Price,
	}

	if err := o.db.Create(&newOrder).Error; err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:        int64(newOrder.ID),
			UserId:    int64(newOrder.User_Id),
			ProductId: int64(newOrder.Product_Id),
			Quantity:  newOrder.Quantity,
			Price:     newOrder.Price,
		}}, nil
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
