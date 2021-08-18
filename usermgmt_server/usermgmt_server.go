package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "github.com/sachithKay/simple-grpc/usermgmt"
	"math/rand"
	"net"
)

const port = ":5001"

type UserManagementServer struct {
	 pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, user *pb.NewUser) (*pb.User, error) {

	fmt.Printf("message received with user %v", user)
	userId := int32(rand.Int31n(1000))
	return &pb.User{Name: user.GetName(), Age: user.GetAge(), Id: userId}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Errorf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Errorf("failed to serve: %v", err)
		return
	}

}
