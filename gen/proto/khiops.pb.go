// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: khiops.proto

package khiops

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BatchRequest_Tool int32

const (
	BatchRequest_KHIOPS              BatchRequest_Tool = 0
	BatchRequest_KHIOPS_COCLUSTERING BatchRequest_Tool = 1
)

// Enum value maps for BatchRequest_Tool.
var (
	BatchRequest_Tool_name = map[int32]string{
		0: "KHIOPS",
		1: "KHIOPS_COCLUSTERING",
	}
	BatchRequest_Tool_value = map[string]int32{
		"KHIOPS":              0,
		"KHIOPS_COCLUSTERING": 1,
	}
)

func (x BatchRequest_Tool) Enum() *BatchRequest_Tool {
	p := new(BatchRequest_Tool)
	*p = x
	return p
}

func (x BatchRequest_Tool) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BatchRequest_Tool) Descriptor() protoreflect.EnumDescriptor {
	return file_khiops_proto_enumTypes[0].Descriptor()
}

func (BatchRequest_Tool) Type() protoreflect.EnumType {
	return &file_khiops_proto_enumTypes[0]
}

func (x BatchRequest_Tool) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchRequest_Tool.Descriptor instead.
func (BatchRequest_Tool) EnumDescriptor() ([]byte, []int) {
	return file_khiops_proto_rawDescGZIP(), []int{0, 0}
}

type BatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tool BatchRequest_Tool `protobuf:"varint,1,opt,name=tool,proto3,enum=com.orange.khiops.BatchRequest_Tool" json:"tool,omitempty"`
	// Types that are assignable to Param:
	//	*BatchRequest_Scenario
	//	*BatchRequest_ScenarioPath
	//	*BatchRequest_BinaryScenario
	Param isBatchRequest_Param `protobuf_oneof:"param"`
}

func (x *BatchRequest) Reset() {
	*x = BatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khiops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequest) ProtoMessage() {}

func (x *BatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_khiops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequest.ProtoReflect.Descriptor instead.
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return file_khiops_proto_rawDescGZIP(), []int{0}
}

func (x *BatchRequest) GetTool() BatchRequest_Tool {
	if x != nil {
		return x.Tool
	}
	return BatchRequest_KHIOPS
}

func (m *BatchRequest) GetParam() isBatchRequest_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *BatchRequest) GetScenario() string {
	if x, ok := x.GetParam().(*BatchRequest_Scenario); ok {
		return x.Scenario
	}
	return ""
}

func (x *BatchRequest) GetScenarioPath() string {
	if x, ok := x.GetParam().(*BatchRequest_ScenarioPath); ok {
		return x.ScenarioPath
	}
	return ""
}

func (x *BatchRequest) GetBinaryScenario() []byte {
	if x, ok := x.GetParam().(*BatchRequest_BinaryScenario); ok {
		return x.BinaryScenario
	}
	return nil
}

type isBatchRequest_Param interface {
	isBatchRequest_Param()
}

type BatchRequest_Scenario struct {
	Scenario string `protobuf:"bytes,2,opt,name=scenario,proto3,oneof"`
}

type BatchRequest_ScenarioPath struct {
	ScenarioPath string `protobuf:"bytes,3,opt,name=scenario_path,json=scenarioPath,proto3,oneof"`
}

type BatchRequest_BinaryScenario struct {
	BinaryScenario []byte `protobuf:"bytes,4,opt,name=binary_scenario,json=binaryScenario,proto3,oneof"`
}

func (*BatchRequest_Scenario) isBatchRequest_Param() {}

func (*BatchRequest_ScenarioPath) isBatchRequest_Param() {}

func (*BatchRequest_BinaryScenario) isBatchRequest_Param() {}

type BatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khiops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_khiops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_khiops_proto_rawDescGZIP(), []int{1}
}

func (x *BatchResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BatchResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *BatchResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_khiops_proto protoreflect.FileDescriptor

var file_khiops_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x68, 0x69, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6b, 0x68, 0x69, 0x6f, 0x70,
	0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6b, 0x68,
	0x69, 0x6f, 0x70, 0x73, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08,
	0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x25, 0x0a, 0x0d, 0x73, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x29, 0x0a, 0x0f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x2b, 0x0a, 0x04,
	0x54, 0x6f, 0x6f, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x48, 0x49, 0x4f, 0x50, 0x53, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x4b, 0x48, 0x49, 0x4f, 0x50, 0x53, 0x5f, 0x43, 0x4f, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x22, 0x55, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xa8, 0x01, 0x0a, 0x06, 0x4b, 0x68,
	0x69, 0x6f, 0x70, 0x73, 0x12, 0x9d, 0x01, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6b,
	0x68, 0x69, 0x6f, 0x70, 0x73, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3a, 0x12, 0x17, 0x52, 0x75, 0x6e, 0x20,
	0x61, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x1a, 0x1f, 0x52, 0x75, 0x6e, 0x20, 0x61, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x20, 0x6f, 0x66, 0x20, 0x4b, 0x68, 0x69, 0x6f, 0x70, 0x73, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x32, 0xfb, 0x04, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x7f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x74, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f,
	0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7a, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7d, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x0d, 0x57, 0x61, 0x69, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x69, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x77, 0x61,
	0x69, 0x74, 0x42, 0x7f, 0x5a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x6b, 0x68, 0x69, 0x6f, 0x70, 0x73, 0x3b, 0x6b, 0x68, 0x69, 0x6f, 0x70, 0x73, 0x92, 0x41,
	0x62, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x02, 0x01, 0x72, 0x55, 0x0a, 0x22,
	0x4b, 0x68, 0x69, 0x6f, 0x70, 0x73, 0x20, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x20, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x6b,
	0x68, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x6b, 0x68, 0x69,
	0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_khiops_proto_rawDescOnce sync.Once
	file_khiops_proto_rawDescData = file_khiops_proto_rawDesc
)

func file_khiops_proto_rawDescGZIP() []byte {
	file_khiops_proto_rawDescOnce.Do(func() {
		file_khiops_proto_rawDescData = protoimpl.X.CompressGZIP(file_khiops_proto_rawDescData)
	})
	return file_khiops_proto_rawDescData
}

var file_khiops_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_khiops_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_khiops_proto_goTypes = []interface{}{
	(BatchRequest_Tool)(0),                     // 0: com.orange.khiops.BatchRequest.Tool
	(*BatchRequest)(nil),                       // 1: com.orange.khiops.BatchRequest
	(*BatchResponse)(nil),                      // 2: com.orange.khiops.BatchResponse
	(*longrunning.ListOperationsRequest)(nil),  // 3: google.longrunning.ListOperationsRequest
	(*longrunning.GetOperationRequest)(nil),    // 4: google.longrunning.GetOperationRequest
	(*longrunning.DeleteOperationRequest)(nil), // 5: google.longrunning.DeleteOperationRequest
	(*longrunning.CancelOperationRequest)(nil), // 6: google.longrunning.CancelOperationRequest
	(*longrunning.WaitOperationRequest)(nil),   // 7: google.longrunning.WaitOperationRequest
	(*longrunning.Operation)(nil),              // 8: google.longrunning.Operation
	(*longrunning.ListOperationsResponse)(nil), // 9: google.longrunning.ListOperationsResponse
	(*empty.Empty)(nil),                        // 10: google.protobuf.Empty
}
var file_khiops_proto_depIdxs = []int32{
	0,  // 0: com.orange.khiops.BatchRequest.tool:type_name -> com.orange.khiops.BatchRequest.Tool
	1,  // 1: com.orange.khiops.Khiops.RunBatch:input_type -> com.orange.khiops.BatchRequest
	3,  // 2: com.orange.khiops.Operations.ListOperations:input_type -> google.longrunning.ListOperationsRequest
	4,  // 3: com.orange.khiops.Operations.GetOperation:input_type -> google.longrunning.GetOperationRequest
	5,  // 4: com.orange.khiops.Operations.DeleteOperation:input_type -> google.longrunning.DeleteOperationRequest
	6,  // 5: com.orange.khiops.Operations.CancelOperation:input_type -> google.longrunning.CancelOperationRequest
	7,  // 6: com.orange.khiops.Operations.WaitOperation:input_type -> google.longrunning.WaitOperationRequest
	8,  // 7: com.orange.khiops.Khiops.RunBatch:output_type -> google.longrunning.Operation
	9,  // 8: com.orange.khiops.Operations.ListOperations:output_type -> google.longrunning.ListOperationsResponse
	8,  // 9: com.orange.khiops.Operations.GetOperation:output_type -> google.longrunning.Operation
	10, // 10: com.orange.khiops.Operations.DeleteOperation:output_type -> google.protobuf.Empty
	10, // 11: com.orange.khiops.Operations.CancelOperation:output_type -> google.protobuf.Empty
	8,  // 12: com.orange.khiops.Operations.WaitOperation:output_type -> google.longrunning.Operation
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_khiops_proto_init() }
func file_khiops_proto_init() {
	if File_khiops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_khiops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchRequest); i {
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
		file_khiops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchResponse); i {
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
	file_khiops_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BatchRequest_Scenario)(nil),
		(*BatchRequest_ScenarioPath)(nil),
		(*BatchRequest_BinaryScenario)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_khiops_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_khiops_proto_goTypes,
		DependencyIndexes: file_khiops_proto_depIdxs,
		EnumInfos:         file_khiops_proto_enumTypes,
		MessageInfos:      file_khiops_proto_msgTypes,
	}.Build()
	File_khiops_proto = out.File
	file_khiops_proto_rawDesc = nil
	file_khiops_proto_goTypes = nil
	file_khiops_proto_depIdxs = nil
}
