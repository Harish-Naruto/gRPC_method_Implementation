// SERVER

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	user "learn.com/grpc/gen/go/proto"
)

type userService struct {
	user.UnimplementedUserStreamServer
}

func (s *userService) ComminucateHello(
	stream grpc.BidiStreamingServer[user.HelloRequest, user.HelloResponce],
) error {

	log.Println("client connected")

	var wg sync.WaitGroup

	// protects concurrent stream.Send()
	var mu sync.Mutex

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("client finished sending")
			break
		}

		if err != nil {
			log.Println("recv error:", err)
			return err
		}

		log.Println("received:", req.Name)

		name := req.Name

		wg.Add(1)

		go func(name string) {
			defer wg.Done()

			log.Println("processing:", name)

			time.Sleep(5 * time.Second)

			res := &user.HelloResponce{
				Title: fmt.Sprintf("hello %s", name),
			}

			// stream.Send is NOT concurrency safe
			mu.Lock()
			defer mu.Unlock()

			if err := stream.Send(res); err != nil {
				log.Println("send error:", err)
				return
			}

			log.Println("response sent:", name)

		}(name)
	}

	// wait all workers
	wg.Wait()

	log.Println("closing stream")

	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	user.RegisterUserStreamServer(
		grpcServer,
		&userService{},
	)

	log.Println("gRPC server started on :9000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}