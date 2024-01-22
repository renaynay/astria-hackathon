package main

import (
	executionv1a2 "github.com/renaynay/astria-hackathon/gen/astria/execution/v1alpha2"
)

type ExecutionServiceServerV1Alpha2 struct {
	executionv1a2.UnimplementedExecutionServiceServer
}

func main() {
	println("hello, world!")
}
