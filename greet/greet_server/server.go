package main
 
import (
    "log"
    "net"
    "fmt"
    "google.golang.org/grpc"
    "github.com/gsafcik/grpc-go-course/greet/greetpb"
)

type server struct {}

func main() {
    fmt.Println("Hello world")

    lis, err := net.Listen("tcp", "0.0.0.0:50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    greetpb.RegisterGreetServiceServer(s, &server{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to server: %v", err)
    }


}