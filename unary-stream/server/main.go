package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	user "learn.com/grpc/gen/go/proto"
)

type userService struct{
	user.UnimplementedSendHelloServer
}


func (u *userService) SendHello(ctx context.Context, req *user.HelloRequest) (*user.HelloResponce,error) {
	return &user.HelloResponce{
		Title: fmt.Sprintf("hello welcome to grpc, %s",req.Name),
	},nil
}

func main() {
	lis,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatal("error: ",err)
	}
	grpcserver := grpc.NewServer()
	user.RegisterSendHelloServer(grpcserver,&userService{})

	if err := grpcserver.Serve(lis); err!= nil {
		log.Fatal("error: ",err)
	}
	
}