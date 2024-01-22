package execution

import (
	"context"

	astriaGrpc "buf.build/gen/go/astria/astria/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
	astriaPb "buf.build/gen/go/astria/astria/protocolbuffers/go/astria/execution/v1alpha2"
)

type ExecutionServiceServerV1Alpha2 struct {
	astriaGrpc.ExecutionServiceServer
}

func NewExecutionServiceServerV1Alpha2() *ExecutionServiceServerV1Alpha2 {
	return &ExecutionServiceServerV1Alpha2{}
}

func (s *ExecutionServiceServerV1Alpha2) GetBlock(ctx context.Context, req *astriaPb.GetBlockRequest) (*astriaPb.Block, error) {
	println("GetBlock called", "request", req)
	res := &astriaPb.Block{
		Number:          uint32(0),
		Hash:            []byte{0x0},
		ParentBlockHash: []byte{0x0},
	}
	println("GetBlock completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) BatchGetBlocks(ctx context.Context, req *astriaPb.BatchGetBlocksRequest) (*astriaPb.BatchGetBlocksResponse, error) {
	println("BatchGetBlocks called", "request", req)
	res := &astriaPb.BatchGetBlocksResponse{
		Blocks: []*astriaPb.Block{
			{
				Number:          uint32(0),
				Hash:            []byte{0x0},
				ParentBlockHash: []byte{0x0},
			},
		},
	}
	println("BatchGetBlocks completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) ExecuteBlock(ctx context.Context, req *astriaPb.ExecuteBlockRequest) (*astriaPb.Block, error) {
	println("ExecuteBlock called", "request", req)
	res := &astriaPb.Block{
		Number:          uint32(0),
		Hash:            []byte{0x0},
		ParentBlockHash: []byte{0x0},
	}
	println("ExecuteBlock completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) GetCommitmentState(ctx context.Context, req *astriaPb.GetCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	println("GetCommitmentState called", "request", req)
	res := &astriaPb.CommitmentState{
		Soft: &astriaPb.Block{
			Number:          uint32(0),
			Hash:            []byte{0x0},
			ParentBlockHash: []byte{0x0},
		},
		Firm: &astriaPb.Block{
			Number:          uint32(0),
			Hash:            []byte{0x0},
			ParentBlockHash: []byte{0x0},
		},
	}
	println("GetCommitmentState completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) UpdateCommitmentState(ctx context.Context, req *astriaPb.UpdateCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	println("UpdateCommitmentState called", "request", req)
	println("UpdateCommitmentState completed", "request", req)
	return req.CommitmentState, nil
}

func (s *ExecutionServiceServerV1Alpha2) getBlockFromIdentifier(identifier *astriaPb.BlockIdentifier) (*astriaPb.Block, error) {
	println("getBlockFromIdentifier called", "identifier", identifier)

	res := &astriaPb.Block{
		Number:          uint32(0),
		Hash:            []byte{0x0},
		ParentBlockHash: []byte{0x0},
	}
	println("getBlockFromIdentifier completed", "identifier", identifier, "response", res)
	return res, nil
}
