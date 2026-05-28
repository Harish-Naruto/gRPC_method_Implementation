package main

import (
	"fmt"
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

func (userService *userService) ComminucateHello( stream grpc.BidiStreamingServer[user.HelloRequest, user.HelloResponce]) error {
	for {
		req ,err := stream.Recv()
		if err == io.EOF{
			log.Println("data stream ended")
			return nil
		}
		if err!= nil {
			log.Fatal(err.Error())
			return err
		}
		log.Println("data received: ",req.Name)
		go func() {
			log.Println("data processing")
			time.Sleep(5*time.Second)
			
			//add mutex to make this safe for concurrency Send is not safe for concurrency
			if err := stream.Send(&user.HelloResponce{
				Title: fmt.Sprint("hello, ",req.Name),
			}); err != nil {
				log.Fatal(err.Error())
			}
		}()		
	}

}


func main() {
	lis,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatal(err.Error())
	}
	gRPCServer := grpc.NewServer()
	us := &userService{}
	user.RegisterUserStreamServer(gRPCServer,us)
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}