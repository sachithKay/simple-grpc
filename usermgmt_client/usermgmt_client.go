package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "sachithk.com/go-grpctesting/usermgmt"
	"time"
)

const address = "localhost:5001"

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Errorf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Alice"] = 43
	new_users["Bob"] = 30
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			fmt.Errorf("could not create user: %v", err)
			return
		}
		fmt.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d \n`, r.GetName(), r.GetAge(), r.GetId())

	}

}
