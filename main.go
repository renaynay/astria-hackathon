package main

import (
	astriaGrpc "buf.build/gen/go/astria/astria/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
	execution "github.com/renaynay/astria-hackathon/src"
	"google.golang.org/grpc"
)

func main() {
	println("hello, world!")
	execution_server := execution.NewExecutionServiceServerV1Alpha2()
	server := grpc.NewServer()
	astriaGrpc.RegisterExecutionServiceServer(server, execution_server)

}
