package execution

import (
	"context"

	executionv1a2 "github.com/renaynay/astria-hackathon/gen/astria/execution/v1alpha2"
)

type ExecutionServiceServerV1Alpha2 struct {
	executionv1a2.UnimplementedExecutionServiceServer
}

func NewExecutionServiceServerV1Alpha2() *ExecutionServiceServerV1Alpha2 {
	return &ExecutionServiceServerV1Alpha2{}
}

func (s *ExecutionServiceServerV1Alpha2) GetBlock(ctx context.Context, req *executionv1a2.GetBlockRequest) (*executionv1a2.Block, error) {
	println("GetBlock called", "request", req)
	res := &executionv1a2.Block{
		Number:          uint32(0),
		Hash:            []byte{0x0},
		ParentBlockHash: []byte{0x0},
	}
	println("GetBlock completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) BatchGetBlocks(ctx context.Context, req *executionv1a2.BatchGetBlocksRequest) (*executionv1a2.BatchGetBlocksResponse, error) {
	println("BatchGetBlocks called", "request", req)
	res := &executionv1a2.BatchGetBlocksResponse{
		Blocks: []*executionv1a2.Block{
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

func (s *ExecutionServiceServerV1Alpha2) ExecuteBlock(ctx context.Context, req *executionv1a2.ExecuteBlockRequest) (*executionv1a2.Block, error) {
	println("ExecuteBlock called", "request", req)
	res := &executionv1a2.Block{
		Number:          uint32(0),
		Hash:            []byte{0x0},
		ParentBlockHash: []byte{0x0},
	}
	println("ExecuteBlock completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) GetCommitmentState(ctx context.Context, req *executionv1a2.GetCommitmentStateRequest) (*executionv1a2.CommitmentState, error) {
	println("GetCommitmentState called", "request", req)
	res := &executionv1a2.CommitmentState{
		Soft: &executionv1a2.Block{
			Number:          uint32(0),
			Hash:            []byte{0x0},
			ParentBlockHash: []byte{0x0},
		},
		Firm: &executionv1a2.Block{
			Number:          uint32(0),
			Hash:            []byte{0x0},
			ParentBlockHash: []byte{0x0},
		},
	}
	println("GetCommitmentState completed", "request", req, "response", res)
	return res, nil
}

func (s *ExecutionServiceServerV1Alpha2) UpdateCommitmentState(ctx context.Context, req *executionv1a2.UpdateCommitmentStateRequest) (*executionv1a2.CommitmentState, error) {
	println("UpdateCommitmentState called", "request", req)
	println("UpdateCommitmentState completed", "request", req)
	return req.CommitmentState, nil
}

func (s *ExecutionServiceServerV1Alpha2) getBlockFromIdentifier(identifier *executionv1a2.BlockIdentifier) (*executionv1a2.Block, error) {
	println("getBlockFromIdentifier called", "identifier", identifier)

	res := &executionv1a2.Block{
		Number:          uint32(0),
		Hash:            []byte{0x0},
		ParentBlockHash: []byte{0x0},
	}
	println("getBlockFromIdentifier completed", "identifier", identifier, "response", res)
	return res, nil
}
