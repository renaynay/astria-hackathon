package main

import (
	astriaGrpc "buf.build/gen/go/astria/astria/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
	"github.com/renaynay/astria-hackathon/messenger"
	"google.golang.org/grpc"
)

func main() {
	println("hello, world!")
	m := messenger.NewMessenger()

	execution_server := messenger.NewExecutionServiceServerV1Alpha2(m)
	server := grpc.NewServer()
	astriaGrpc.RegisterExecutionServiceServer(server, execution_server)

}
