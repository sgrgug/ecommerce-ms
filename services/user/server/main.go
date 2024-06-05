package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sgrgug/ecommerce-ms/api/protos/user"
	"github.com/sgrgug/ecommerce-ms/services/user/model"
	"google.golang.org/grpc"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := model.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
	}

	if err := u.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:       int64(newUser.ID),
			Username: newUser.Username,
			Email:    newUser.Email,
		}}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user model.User

	if err := u.db.First(&user, req.Id).Error; err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:       int64(user.ID),
			Username: user.Username,
			Email:    user.Email,
		}}, nil
}

func main() {

	dsn := "root:@tcp(localhost:3306)/e_ws_user_service?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.User{})

	lis, err := net.Listen("tcp", ":60001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserService{db: db})

	log.Println("Starting server on port :60001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
