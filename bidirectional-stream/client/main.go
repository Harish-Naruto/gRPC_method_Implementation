package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "learn.com/grpc/gen/go/proto"
)

func main() {
	conn, err := grpc.NewClient(":9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	client := user.NewUserStreamClient(conn)
	names := []string{
		"john","sam","tom","jerry",
	}
	stream,err := client.ComminucateHello(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg sync.WaitGroup

	for _,i:=range(names){
		time.Sleep(3*time.Second)
		if err :=stream.Send(&user.HelloRequest{
			Name: i,
		});err != nil{
			log.Fatal(err.Error())
		}
		log.Println("data send")
		wg.Add(1)
		go func() {
			defer wg.Done()
			res,err:= stream.Recv()
			if err == io.EOF{
				return
			} 
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Println("recieved data: ",res.Title)
		}()
	
	}
	wg.Wait()
	log.Println("send stream close signal")
	stream.CloseSend()
}