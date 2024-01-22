package messenger

import (
	"context"
	"errors"

	astriaGrpc "buf.build/gen/go/astria/astria/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
	astriaPb "buf.build/gen/go/astria/astria/protocolbuffers/go/astria/execution/v1alpha2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ExecutionServiceServerV1Alpha2 struct {
	astriaGrpc.ExecutionServiceServer
	m *Messenger
}

func NewExecutionServiceServerV1Alpha2(m *Messenger) *ExecutionServiceServerV1Alpha2 {
	return &ExecutionServiceServerV1Alpha2{
		m: m,
	}
}

func (s *ExecutionServiceServerV1Alpha2) getSingleBlock(height uint32) (*astriaPb.Block, error) {
	if height > uint32(len(s.m.Blocks)) {
		return nil, errors.New("block not found")
	}

	block := s.m.Blocks[height]
	timestamp := timestamppb.New(block.timestamp)

	return &astriaPb.Block{
		Number:          height,
		Hash:            block.hash,
		ParentBlockHash: s.m.Blocks[height-1].hash,
		Timestamp:       timestamp,
	}, nil
}

func (s *ExecutionServiceServerV1Alpha2) GetBlock(ctx context.Context, req *astriaPb.GetBlockRequest) (*astriaPb.Block, error) {
	switch req.Identifier.Identifier.(type) {
	case *astriaPb.BlockIdentifier_BlockNumber:
		block, err := s.getSingleBlock(uint32(req.Identifier.GetBlockNumber()))
		if err != nil {
			return nil, err
		}
		return block, nil
	default:
		return nil, errors.New("invalid identifier")
	}
}

func (s *ExecutionServiceServerV1Alpha2) BatchGetBlocks(ctx context.Context, req *astriaPb.BatchGetBlocksRequest) (*astriaPb.BatchGetBlocksResponse, error) {
	res := &astriaPb.BatchGetBlocksResponse{
		Blocks: []*astriaPb.Block{},
	}
	for _, id := range req.Identifiers {
		switch id.Identifier.(type) {
		case *astriaPb.BlockIdentifier_BlockNumber:
			block, err := s.getSingleBlock(uint32(id.GetBlockNumber()))
			if err != nil {
				return nil, err
			}
			res.Blocks = append(res.Blocks, block)
		}
	}
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
