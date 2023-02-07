package main

import (
	"buf.build/gen/go/semichkin/moduletwo/grpc/go/proto/someservice/v1/someservicev1grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	runGrpc(50500)
}

func runGrpc(port int) {
	address := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen", "err", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	someservicev1grpc.RegisterSomeServiceServer(
		s,
		nil,
	)

	log.Println("server listening", "addr", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve", "err", err)
	}
}
