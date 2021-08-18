package main

import (
	"context"
	"fmt"
	choreo "github.com/wso2-enterprise/choreo-runtime/pkg/apis"
	"github.com/wso2-enterprise/choreo-runtime/pkg/auth0"
	grpcmiddleware "github.com/wso2-enterprise/choreo-runtime/pkg/grpc/middleware"
	"google.golang.org/grpc"
)

const address = "localhost:9990"

func main() {

	fmt.Printf("start \n")
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcmiddleware.ChainUnaryClient(
			grpcmiddleware.UnaryUserContextClientInterceptor(),
		)),
		grpc.WithStreamInterceptor(grpcmiddleware.ChainStreamClient(
			grpcmiddleware.StreamUserContextClientInterceptor(),
		)),
		)
	if err != nil {
		fmt.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	fmt.Printf("dialed \n")

	userclient := choreo.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = auth0.WithUserContext(ctx, auth0.SystemUserId)

	fmt.Printf("im here")
	request := &choreo.DeleteOrganizationRequest{OrganizationName: "hopefultoucan"}

	fmt.Printf("im here.... \n")
	response, err := userclient.DeleteOrganization(ctx, request)
	if err != nil {
		fmt.Errorf("shit happened %v", err)

	}

	fmt.Printf("im here....mklmm")
	fmt.Printf("Response %v", response)
}
