package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	user "learn.com/grpc/gen/go/proto"
)

type userService struct {
	user.UnimplementedUserStreamServer
}
 
func (userService *userService) SendHeartbeat(req *user.HelloRequest,stream grpc.ServerStreamingServer[user.HelloResponce]) error {
  	
	name:= req.Name

	// stream data for 3 second
	for i:=0;i<3;i++ {
		select{
		case <- stream.Context().Done():  //this trigger when client cancle context or get disconnected
			log.Println("connection closed")
			return status.Error(codes.Canceled,"stream has ended")
		default:
			time.Sleep(1*time.Second)
			//diff between Send and sendMsg for stream
			err := stream.Send(&user.HelloResponce{
				Title: fmt.Sprint("hello ",name),
			})
			if err!=nil {
				return status.Error(codes.Aborted,err.Error())
			}
		}
	}
	return nil
}
func main() {
	lis ,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatal("error while listing on port 9000: ",err.Error())
	}

	gRpcServer := grpc.NewServer()
	us := &userService{}
	user.RegisterUserStreamServer(gRpcServer,us)
	log.Println("gRPC server started on Port 9000")
	if err := gRpcServer.Serve(lis); err!= nil {
		log.Fatal("error: ",err.Error())
	}
}