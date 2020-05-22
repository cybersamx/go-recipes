// Modified from the official gRPC example: google.golang.org/grpc/examples/helloworld/helloworld.

package main

import (
	"context"
	"github.com/cybersamx/go-recipes/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syreclabs.com/go/faker"
	"syscall"
	"time"
)

const (
	port    = ":50051"
	address = "localhost:50051"
	duration = 5 * time.Second
)

type server struct{}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &hello.HelloReply{Message: "Hello " + in.Name}, nil
}

func runServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runClient(ctx context.Context, address, name string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hello.NewGreeterClient(conn)

	// Contact the server and print out its response.
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	cctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.SayHello(cctx, &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func main() {
	// Handle system signals.
	sigChan := make(chan os.Signal, 1)
	ctx := context.Background()
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// Seed the randomizer for generating random fake names.
	faker.Seed(time.Now().UnixNano())

	// Goroutine: server.
	go runServer(port)

	// Goroutine: client.
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(duration):
				runClient(ctx, address, faker.Name().Name())
			}
		}
	}()

	sig := <-sigChan
	log.Printf("received signal %d, terminating...\n", sig)
	os.Exit(0)
}
