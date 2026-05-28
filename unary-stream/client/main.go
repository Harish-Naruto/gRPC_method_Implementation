package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "learn.com/grpc/gen/go/proto"
)

func main() {
	conn,err := grpc.NewClient("localhost:9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
		log.Fatal("error in client: ",err)
	}
	defer conn.Close()

	clnt := user.NewSendHelloClient(conn)

	res,err := clnt.SendHello(context.Background(),&user.HelloRequest{
		Name: "Client",
	})
	if err!= nil {
		log.Println("error in request: ",err)
	}
	log.Println(res)

}