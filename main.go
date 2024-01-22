package main

import (
	astriaGrpc "buf.build/gen/go/astria/astria/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
)

type ExecutionServiceServerV1Alpha2 struct {
	astriaGrpc.ExecutionServiceServer
}

func main() {
	println("hello, world!")
}
