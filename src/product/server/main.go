package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/sgrgug/ecommerce-ms/src/product/genproto"
	"github.com/sgrgug/ecommerce-ms/src/product/model"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedProductServiceServer
	db *gorm.DB
}

func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	newProduct := model.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       req.Product.Price,
		Quantity:    req.Product.Quantity,
		Category:    req.Product.Category,
	}

	if err := s.db.Create(&newProduct).Error; err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:          newProduct.ID,
			Name:        newProduct.Name,
			Description: newProduct.Description,
			Price:       newProduct.Price,
			Quantity:    newProduct.Quantity,
			Category:    newProduct.Category,
		}}, nil
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	var product model.Product

	if err := s.db.First(&product, req.Id).Error; err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Category:    product.Category,
		}}, nil
}

func main() {

	dsn := "host=localhost user=postgres password=postgresql dbname=e_ws_product_service port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Product{})

	lis, err := net.Listen("tcp", ":60000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{db: db})

	log.Printf("Server is running on port %s", os.Getenv("PRODUCT_SERVICE_SERVER_PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
