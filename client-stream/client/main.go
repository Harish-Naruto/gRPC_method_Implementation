package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "learn.com/grpc/gen/go/proto"
)

func main() {
	conn,err := grpc.NewClient(":9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	client := user.NewUserStreamClient(conn)
	stream,err := client.ReceiveHeartbeat(context.Background())
	
	//sending data for 4 second
	for i := 0; i < 4; i++ {
		select{
		case <- stream.Context().Done():
			log.Fatal("stream closed")
		default:
			time.Sleep(1*time.Second)
			stream.Send(&user.HelloRequest{
				Name: "client",
			})
		}
	}
	//this close stream from client after receiving data
	// waits till data is recieved by client
	res,err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(res.Title)

}