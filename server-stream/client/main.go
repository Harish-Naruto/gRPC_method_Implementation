package main

import (
	"context"
	"io"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "learn.com/grpc/gen/go/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error: ",err.Error())
	}
	defer conn.Close()
	client := user.NewUserStreamClient(conn)
	
	ctx ,cancle:= context.WithTimeout(context.Background(),5*time.Second)
	defer cancle()

	stream, err := client.SendHeartbeat(ctx,&user.HelloRequest{
		Name: "Client",
	})
	if err != nil {
		log.Fatal("error: ",err.Error())
	}
	go func() {
		defer cancle()
		for{
			val,err := stream.Recv()
			if err == io.EOF {
				log.Println("stream ended")
				return
			}
			if err != nil{
				log.Fatal("Stream error : ",err.Error())
			}
			log.Printf("%s", val.Title)
		}
	}()
	
	// wait till this context is canceled or done
	<-ctx.Done()

}