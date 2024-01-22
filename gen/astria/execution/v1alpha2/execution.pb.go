// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: astria/execution/v1alpha2/execution.proto

package executionv1alpha2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The set of information which deterministic driver of block production
// must know about a given rollup Block
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The block number
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// The hash of the block
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The hash from the parent block
	ParentBlockHash []byte `protobuf:"bytes,3,opt,name=parent_block_hash,json=parentBlockHash,proto3" json:"parent_block_hash,omitempty"`
	// Timestamp on the block, standardized to google protobuf standard.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Block) GetParentBlockHash() []byte {
	if x != nil {
		return x.ParentBlockHash
	}
	return nil
}

func (x *Block) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// Fields which are indexed for finding blocks on a blockchain.
type BlockIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*BlockIdentifier_BlockNumber
	//	*BlockIdentifier_BlockHash
	Identifier isBlockIdentifier_Identifier `protobuf_oneof:"identifier"`
}

func (x *BlockIdentifier) Reset() {
	*x = BlockIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockIdentifier) ProtoMessage() {}

func (x *BlockIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockIdentifier.ProtoReflect.Descriptor instead.
func (*BlockIdentifier) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{1}
}

func (m *BlockIdentifier) GetIdentifier() isBlockIdentifier_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *BlockIdentifier) GetBlockNumber() uint32 {
	if x, ok := x.GetIdentifier().(*BlockIdentifier_BlockNumber); ok {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockIdentifier) GetBlockHash() []byte {
	if x, ok := x.GetIdentifier().(*BlockIdentifier_BlockHash); ok {
		return x.BlockHash
	}
	return nil
}

type isBlockIdentifier_Identifier interface {
	isBlockIdentifier_Identifier()
}

type BlockIdentifier_BlockNumber struct {
	BlockNumber uint32 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3,oneof"`
}

type BlockIdentifier_BlockHash struct {
	BlockHash []byte `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3,oneof"`
}

func (*BlockIdentifier_BlockNumber) isBlockIdentifier_Identifier() {}

func (*BlockIdentifier_BlockHash) isBlockIdentifier_Identifier() {}

// Used in GetBlock to find a single block.
type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *BlockIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockRequest) GetIdentifier() *BlockIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

// Used in BatchGetBlocks, will find all or none based on the list of
// identifiers.
type BatchGetBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiers []*BlockIdentifier `protobuf:"bytes,1,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
}

func (x *BatchGetBlocksRequest) Reset() {
	*x = BatchGetBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetBlocksRequest) ProtoMessage() {}

func (x *BatchGetBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetBlocksRequest.ProtoReflect.Descriptor instead.
func (*BatchGetBlocksRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetBlocksRequest) GetIdentifiers() []*BlockIdentifier {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

// The list of blocks in response to BatchGetBlocks.
type BatchGetBlocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BatchGetBlocksResponse) Reset() {
	*x = BatchGetBlocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetBlocksResponse) ProtoMessage() {}

func (x *BatchGetBlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetBlocksResponse.ProtoReflect.Descriptor instead.
func (*BatchGetBlocksResponse) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{4}
}

func (x *BatchGetBlocksResponse) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

// ExecuteBlockRequest contains all the information needed to create a new rollup
// block.
//
// This information comes from previous rollup blocks, as well as from sequencer
// blocks.
type ExecuteBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of previous block, which new block will be created on top of.
	PrevBlockHash []byte `protobuf:"bytes,1,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	// List of transactions to include in the new block.
	Transactions [][]byte `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// Timestamp to be used for new block.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ExecuteBlockRequest) Reset() {
	*x = ExecuteBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteBlockRequest) ProtoMessage() {}

func (x *ExecuteBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteBlockRequest.ProtoReflect.Descriptor instead.
func (*ExecuteBlockRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{5}
}

func (x *ExecuteBlockRequest) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *ExecuteBlockRequest) GetTransactions() [][]byte {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ExecuteBlockRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// The CommitmentState holds the block at each stage of sequencer commitment
// level
//
// A Valid CommitmentState:
//   - Block numbers are such that soft >= firm.
//   - No blocks ever decrease in block number.
//   - The chain defined by soft is the head of the canonical chain the firm block
//     must belong to.
type CommitmentState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Soft commitment is the rollup block matching latest sequencer block.
	Soft *Block `protobuf:"bytes,1,opt,name=soft,proto3" json:"soft,omitempty"`
	// Firm commitment is achieved when data has been seen in DA.
	Firm *Block `protobuf:"bytes,2,opt,name=firm,proto3" json:"firm,omitempty"`
}

func (x *CommitmentState) Reset() {
	*x = CommitmentState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitmentState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitmentState) ProtoMessage() {}

func (x *CommitmentState) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitmentState.ProtoReflect.Descriptor instead.
func (*CommitmentState) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{6}
}

func (x *CommitmentState) GetSoft() *Block {
	if x != nil {
		return x.Soft
	}
	return nil
}

func (x *CommitmentState) GetFirm() *Block {
	if x != nil {
		return x.Firm
	}
	return nil
}

// There is only one CommitmentState object, so the request is empty.
type GetCommitmentStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCommitmentStateRequest) Reset() {
	*x = GetCommitmentStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommitmentStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommitmentStateRequest) ProtoMessage() {}

func (x *GetCommitmentStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommitmentStateRequest.ProtoReflect.Descriptor instead.
func (*GetCommitmentStateRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{7}
}

// The CommitmentState to set, must include complete state.
type UpdateCommitmentStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitmentState *CommitmentState `protobuf:"bytes,1,opt,name=commitment_state,json=commitmentState,proto3" json:"commitment_state,omitempty"`
}

func (x *UpdateCommitmentStateRequest) Reset() {
	*x = UpdateCommitmentStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommitmentStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommitmentStateRequest) ProtoMessage() {}

func (x *UpdateCommitmentStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1alpha2_execution_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommitmentStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommitmentStateRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1alpha2_execution_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCommitmentStateRequest) GetCommitmentState() *CommitmentState {
	if x != nil {
		return x.CommitmentState
	}
	return nil
}

var File_astria_execution_v1alpha2_execution_proto protoreflect.FileDescriptor

var file_astria_execution_v1alpha2_execution_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x73, 0x74,
	0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a,
	0x11, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x65, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x42, 0x0c, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x15, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x22, 0x52, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x73, 0x74,
	0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x70, 0x72, 0x65, 0x76, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x7d, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x6f, 0x66, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x73, 0x6f, 0x66, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x66,
	0x69, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x73, 0x74, 0x72,
	0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x66, 0x69, 0x72,
	0x6d, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x75,
	0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69,
	0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0xbb, 0x04, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x75, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x30, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69,
	0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x2e, 0x61, 0x73,
	0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x73,
	0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x76, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72,
	0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x7c, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x37,
	0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x87, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x73, 0x74, 0x72,
	0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6e, 0x61, 0x79, 0x6e, 0x61, 0x79, 0x2f, 0x61, 0x73, 0x74,
	0x72, 0x69, 0x61, 0x2d, 0x68, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x61, 0x73,
	0x74, 0x72, 0x69, 0x61, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa,
	0x02, 0x19, 0x41, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x19, 0x41, 0x73,
	0x74, 0x72, 0x69, 0x61, 0x5c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x25, 0x41, 0x73, 0x74, 0x72, 0x69, 0x61,
	0x5c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1b, 0x41, 0x73, 0x74, 0x72, 0x69, 0x61, 0x3a, 0x3a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_astria_execution_v1alpha2_execution_proto_rawDescOnce sync.Once
	file_astria_execution_v1alpha2_execution_proto_rawDescData = file_astria_execution_v1alpha2_execution_proto_rawDesc
)

func file_astria_execution_v1alpha2_execution_proto_rawDescGZIP() []byte {
	file_astria_execution_v1alpha2_execution_proto_rawDescOnce.Do(func() {
		file_astria_execution_v1alpha2_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_astria_execution_v1alpha2_execution_proto_rawDescData)
	})
	return file_astria_execution_v1alpha2_execution_proto_rawDescData
}

var file_astria_execution_v1alpha2_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_astria_execution_v1alpha2_execution_proto_goTypes = []interface{}{
	(*Block)(nil),                        // 0: astria.execution.v1alpha2.Block
	(*BlockIdentifier)(nil),              // 1: astria.execution.v1alpha2.BlockIdentifier
	(*GetBlockRequest)(nil),              // 2: astria.execution.v1alpha2.GetBlockRequest
	(*BatchGetBlocksRequest)(nil),        // 3: astria.execution.v1alpha2.BatchGetBlocksRequest
	(*BatchGetBlocksResponse)(nil),       // 4: astria.execution.v1alpha2.BatchGetBlocksResponse
	(*ExecuteBlockRequest)(nil),          // 5: astria.execution.v1alpha2.ExecuteBlockRequest
	(*CommitmentState)(nil),              // 6: astria.execution.v1alpha2.CommitmentState
	(*GetCommitmentStateRequest)(nil),    // 7: astria.execution.v1alpha2.GetCommitmentStateRequest
	(*UpdateCommitmentStateRequest)(nil), // 8: astria.execution.v1alpha2.UpdateCommitmentStateRequest
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
}
var file_astria_execution_v1alpha2_execution_proto_depIdxs = []int32{
	9,  // 0: astria.execution.v1alpha2.Block.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 1: astria.execution.v1alpha2.GetBlockRequest.identifier:type_name -> astria.execution.v1alpha2.BlockIdentifier
	1,  // 2: astria.execution.v1alpha2.BatchGetBlocksRequest.identifiers:type_name -> astria.execution.v1alpha2.BlockIdentifier
	0,  // 3: astria.execution.v1alpha2.BatchGetBlocksResponse.blocks:type_name -> astria.execution.v1alpha2.Block
	9,  // 4: astria.execution.v1alpha2.ExecuteBlockRequest.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 5: astria.execution.v1alpha2.CommitmentState.soft:type_name -> astria.execution.v1alpha2.Block
	0,  // 6: astria.execution.v1alpha2.CommitmentState.firm:type_name -> astria.execution.v1alpha2.Block
	6,  // 7: astria.execution.v1alpha2.UpdateCommitmentStateRequest.commitment_state:type_name -> astria.execution.v1alpha2.CommitmentState
	2,  // 8: astria.execution.v1alpha2.ExecutionService.GetBlock:input_type -> astria.execution.v1alpha2.GetBlockRequest
	3,  // 9: astria.execution.v1alpha2.ExecutionService.BatchGetBlocks:input_type -> astria.execution.v1alpha2.BatchGetBlocksRequest
	5,  // 10: astria.execution.v1alpha2.ExecutionService.ExecuteBlock:input_type -> astria.execution.v1alpha2.ExecuteBlockRequest
	7,  // 11: astria.execution.v1alpha2.ExecutionService.GetCommitmentState:input_type -> astria.execution.v1alpha2.GetCommitmentStateRequest
	8,  // 12: astria.execution.v1alpha2.ExecutionService.UpdateCommitmentState:input_type -> astria.execution.v1alpha2.UpdateCommitmentStateRequest
	0,  // 13: astria.execution.v1alpha2.ExecutionService.GetBlock:output_type -> astria.execution.v1alpha2.Block
	4,  // 14: astria.execution.v1alpha2.ExecutionService.BatchGetBlocks:output_type -> astria.execution.v1alpha2.BatchGetBlocksResponse
	0,  // 15: astria.execution.v1alpha2.ExecutionService.ExecuteBlock:output_type -> astria.execution.v1alpha2.Block
	6,  // 16: astria.execution.v1alpha2.ExecutionService.GetCommitmentState:output_type -> astria.execution.v1alpha2.CommitmentState
	6,  // 17: astria.execution.v1alpha2.ExecutionService.UpdateCommitmentState:output_type -> astria.execution.v1alpha2.CommitmentState
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_astria_execution_v1alpha2_execution_proto_init() }
func file_astria_execution_v1alpha2_execution_proto_init() {
	if File_astria_execution_v1alpha2_execution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_astria_execution_v1alpha2_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetBlocksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetBlocksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitmentState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommitmentStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_astria_execution_v1alpha2_execution_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommitmentStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_astria_execution_v1alpha2_execution_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BlockIdentifier_BlockNumber)(nil),
		(*BlockIdentifier_BlockHash)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_astria_execution_v1alpha2_execution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_astria_execution_v1alpha2_execution_proto_goTypes,
		DependencyIndexes: file_astria_execution_v1alpha2_execution_proto_depIdxs,
		MessageInfos:      file_astria_execution_v1alpha2_execution_proto_msgTypes,
	}.Build()
	File_astria_execution_v1alpha2_execution_proto = out.File
	file_astria_execution_v1alpha2_execution_proto_rawDesc = nil
	file_astria_execution_v1alpha2_execution_proto_goTypes = nil
	file_astria_execution_v1alpha2_execution_proto_depIdxs = nil
}
