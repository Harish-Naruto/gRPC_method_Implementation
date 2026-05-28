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
	conn, err := grpc.NewClient(
		":9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := user.NewUserStreamClient(conn)

	stream, err := client.ComminucateHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	names := []string{
		"john",
		"sam",
		"tom",
		"jerry",
	}

	var wg sync.WaitGroup

	// RECEIVE LOOP
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				log.Println("server closed stream")
				return
			}

			if err != nil {
				log.Println("recv error:", err)
				return
			}

			log.Println("received:", res.Title)
		}
	}()

	// SEND LOOP
	for _, name := range names {

		req := &user.HelloRequest{
			Name: name,
		}

		log.Println("sending:", name)

		if err := stream.Send(req); err != nil {
			log.Println("send error:", err)
			break
		}

		time.Sleep(2 * time.Second)
	}

	// IMPORTANT:
	// tells server client finished sending
	if err := stream.CloseSend(); err != nil {
		log.Println("close send error:", err)
	}

	log.Println("client finished sending")

	wg.Wait()
}