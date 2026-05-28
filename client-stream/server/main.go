package main

import (
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	user "learn.com/grpc/gen/go/proto"
)

type userService struct {
	user.UnimplementedUserStreamServer
}

func (userService *userService) ReceiveHeartbeat(stream grpc.ClientStreamingServer[user.HelloRequest, user.HelloResponce]) error {
	var temp string
	for {
		data,err := stream.Recv()
		if err == io.EOF{
			//send data and close stream
			log.Println("data processing.....")
			time.Sleep(4*time.Second)
			stream.SendAndClose(&user.HelloResponce{
				Title: temp,
			})
			log.Print("Stream closed")
			return nil
		} 
		if err != nil {
			return err
		}
		log.Printf("data recieved: %s",data.Name)
		temp = temp + " " + data.Name
	}  
}

// same as before
func main() {
	lis,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatal(err.Error())
	}
	gRPCServer := grpc.NewServer()
	us := &userService{}
	user.RegisterUserStreamServer(gRPCServer,us)
	log.Println("gRPC server started at port 9000")
	if err := gRPCServer.Serve(lis);err!= nil {
		log.Fatal(err.Error())
	}
}