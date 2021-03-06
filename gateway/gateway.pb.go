// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: gateway/gateway.proto

package gateway

import (
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

type GetIpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRequestId int64    `protobuf:"varint,1,opt,name=userRequestId,proto3" json:"userRequestId,omitempty"`
	IpDevId       []string `protobuf:"bytes,2,rep,name=ipDevId,proto3" json:"ipDevId,omitempty"`
}

func (x *GetIpRequest) Reset() {
	*x = GetIpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIpRequest) ProtoMessage() {}

func (x *GetIpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIpRequest.ProtoReflect.Descriptor instead.
func (*GetIpRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GetIpRequest) GetUserRequestId() int64 {
	if x != nil {
		return x.UserRequestId
	}
	return 0
}

func (x *GetIpRequest) GetIpDevId() []string {
	if x != nil {
		return x.IpDevId
	}
	return nil
}

type IpInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyIp   int64  `protobuf:"varint,1,opt,name=proxyIp,proto3" json:"proxyIp,omitempty"`
	ProxyPort int32  `protobuf:"varint,2,opt,name=proxyPort,proto3" json:"proxyPort,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Province  int32  `protobuf:"varint,5,opt,name=province,proto3" json:"province,omitempty"`
	City      int32  `protobuf:"varint,6,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *IpInfo) Reset() {
	*x = IpInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpInfo) ProtoMessage() {}

func (x *IpInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpInfo.ProtoReflect.Descriptor instead.
func (*IpInfo) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *IpInfo) GetProxyIp() int64 {
	if x != nil {
		return x.ProxyIp
	}
	return 0
}

func (x *IpInfo) GetProxyPort() int32 {
	if x != nil {
		return x.ProxyPort
	}
	return 0
}

func (x *IpInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IpInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IpInfo) GetProvince() int32 {
	if x != nil {
		return x.Province
	}
	return 0
}

func (x *IpInfo) GetCity() int32 {
	if x != nil {
		return x.City
	}
	return 0
}

type GetIpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRequestId int64   `protobuf:"varint,1,opt,name=userRequestId,proto3" json:"userRequestId,omitempty"`
	IpList        *IpInfo `protobuf:"bytes,2,opt,name=ipList,proto3" json:"ipList,omitempty"`
}

func (x *GetIpResponse) Reset() {
	*x = GetIpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIpResponse) ProtoMessage() {}

func (x *GetIpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIpResponse.ProtoReflect.Descriptor instead.
func (*GetIpResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GetIpResponse) GetUserRequestId() int64 {
	if x != nil {
		return x.UserRequestId
	}
	return 0
}

func (x *GetIpResponse) GetIpList() *IpInfo {
	if x != nil {
		return x.IpList
	}
	return nil
}

type ConnDevRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRequestId int64  `protobuf:"varint,1,opt,name=userRequestId,proto3" json:"userRequestId,omitempty"`
	GetIpCid      int64  `protobuf:"varint,2,opt,name=GetIpCid,proto3" json:"GetIpCid,omitempty"`
	GatewayAddr   string `protobuf:"bytes,3,opt,name=gatewayAddr,proto3" json:"gatewayAddr,omitempty"`
}

func (x *ConnDevRequest) Reset() {
	*x = ConnDevRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnDevRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnDevRequest) ProtoMessage() {}

func (x *ConnDevRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnDevRequest.ProtoReflect.Descriptor instead.
func (*ConnDevRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *ConnDevRequest) GetUserRequestId() int64 {
	if x != nil {
		return x.UserRequestId
	}
	return 0
}

func (x *ConnDevRequest) GetGetIpCid() int64 {
	if x != nil {
		return x.GetIpCid
	}
	return 0
}

func (x *ConnDevRequest) GetGatewayAddr() string {
	if x != nil {
		return x.GatewayAddr
	}
	return ""
}

type ConnDevResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRequestId int64  `protobuf:"varint,1,opt,name=userRequestId,proto3" json:"userRequestId,omitempty"`
	GetIpCid      int64  `protobuf:"varint,2,opt,name=GetIpCid,proto3" json:"GetIpCid,omitempty"`
	ProxyAddr     string `protobuf:"bytes,3,opt,name=proxyAddr,proto3" json:"proxyAddr,omitempty"`
}

func (x *ConnDevResponse) Reset() {
	*x = ConnDevResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnDevResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnDevResponse) ProtoMessage() {}

func (x *ConnDevResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnDevResponse.ProtoReflect.Descriptor instead.
func (*ConnDevResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *ConnDevResponse) GetUserRequestId() int64 {
	if x != nil {
		return x.UserRequestId
	}
	return 0
}

func (x *ConnDevResponse) GetGetIpCid() int64 {
	if x != nil {
		return x.GetIpCid
	}
	return 0
}

func (x *ConnDevResponse) GetProxyAddr() string {
	if x != nil {
		return x.ProxyAddr
	}
	return ""
}

type RemoteConnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S5ConnCid int64  `protobuf:"varint,1,opt,name=s5ConnCid,proto3" json:"s5ConnCid,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *RemoteConnRequest) Reset() {
	*x = RemoteConnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConnRequest) ProtoMessage() {}

func (x *RemoteConnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConnRequest.ProtoReflect.Descriptor instead.
func (*RemoteConnRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *RemoteConnRequest) GetS5ConnCid() int64 {
	if x != nil {
		return x.S5ConnCid
	}
	return 0
}

func (x *RemoteConnRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type RemoteConnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S5ConnCid     int64 `protobuf:"varint,1,opt,name=s5ConnCid,proto3" json:"s5ConnCid,omitempty"`
	RemoteConnCid int64 `protobuf:"varint,2,opt,name=remoteConnCid,proto3" json:"remoteConnCid,omitempty"`
}

func (x *RemoteConnResponse) Reset() {
	*x = RemoteConnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConnResponse) ProtoMessage() {}

func (x *RemoteConnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConnResponse.ProtoReflect.Descriptor instead.
func (*RemoteConnResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *RemoteConnResponse) GetS5ConnCid() int64 {
	if x != nil {
		return x.S5ConnCid
	}
	return 0
}

func (x *RemoteConnResponse) GetRemoteConnCid() int64 {
	if x != nil {
		return x.RemoteConnCid
	}
	return 0
}

type RemoteDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteConnCid int64  `protobuf:"varint,1,opt,name=remoteConnCid,proto3" json:"remoteConnCid,omitempty"`
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RemoteDataRequest) Reset() {
	*x = RemoteDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteDataRequest) ProtoMessage() {}

func (x *RemoteDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteDataRequest.ProtoReflect.Descriptor instead.
func (*RemoteDataRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *RemoteDataRequest) GetRemoteConnCid() int64 {
	if x != nil {
		return x.RemoteConnCid
	}
	return 0
}

func (x *RemoteDataRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoteDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S5ConnCid int64  `protobuf:"varint,1,opt,name=s5ConnCid,proto3" json:"s5ConnCid,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RemoteDataResponse) Reset() {
	*x = RemoteDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteDataResponse) ProtoMessage() {}

func (x *RemoteDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteDataResponse.ProtoReflect.Descriptor instead.
func (*RemoteDataResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *RemoteDataResponse) GetS5ConnCid() int64 {
	if x != nil {
		return x.S5ConnCid
	}
	return 0
}

func (x *RemoteDataResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_gateway_gateway_proto protoreflect.FileDescriptor

var file_gateway_gateway_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x22, 0x4e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x70, 0x44, 0x65, 0x76, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x70, 0x44, 0x65, 0x76, 0x49, 0x64,
	0x22, 0xa8, 0x01, 0x0a, 0x06, 0x49, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x49, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x5e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x49, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x49, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x74, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x6e, 0x44, 0x65, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x70, 0x43, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x47, 0x65, 0x74, 0x49, 0x70, 0x43, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x22, 0x71, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x44, 0x65, 0x76, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x49, 0x70, 0x43, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x47, 0x65,
	0x74, 0x49, 0x70, 0x43, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x22, 0x45, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x35, 0x43,
	0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x35,
	0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x58, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x35, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x35, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x43, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x35,
	0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x35, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gateway_gateway_proto_rawDescOnce sync.Once
	file_gateway_gateway_proto_rawDescData = file_gateway_gateway_proto_rawDesc
)

func file_gateway_gateway_proto_rawDescGZIP() []byte {
	file_gateway_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_gateway_proto_rawDescData)
	})
	return file_gateway_gateway_proto_rawDescData
}

var file_gateway_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gateway_gateway_proto_goTypes = []interface{}{
	(*GetIpRequest)(nil),       // 0: gateway.GetIpRequest
	(*IpInfo)(nil),             // 1: gateway.IpInfo
	(*GetIpResponse)(nil),      // 2: gateway.GetIpResponse
	(*ConnDevRequest)(nil),     // 3: gateway.ConnDevRequest
	(*ConnDevResponse)(nil),    // 4: gateway.ConnDevResponse
	(*RemoteConnRequest)(nil),  // 5: gateway.RemoteConnRequest
	(*RemoteConnResponse)(nil), // 6: gateway.RemoteConnResponse
	(*RemoteDataRequest)(nil),  // 7: gateway.RemoteDataRequest
	(*RemoteDataResponse)(nil), // 8: gateway.RemoteDataResponse
}
var file_gateway_gateway_proto_depIdxs = []int32{
	1, // 0: gateway.GetIpResponse.ipList:type_name -> gateway.IpInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gateway_gateway_proto_init() }
func file_gateway_gateway_proto_init() {
	if File_gateway_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIpRequest); i {
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
		file_gateway_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpInfo); i {
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
		file_gateway_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIpResponse); i {
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
		file_gateway_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnDevRequest); i {
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
		file_gateway_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnDevResponse); i {
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
		file_gateway_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConnRequest); i {
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
		file_gateway_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConnResponse); i {
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
		file_gateway_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteDataRequest); i {
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
		file_gateway_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteDataResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_gateway_proto_msgTypes,
	}.Build()
	File_gateway_gateway_proto = out.File
	file_gateway_gateway_proto_rawDesc = nil
	file_gateway_gateway_proto_goTypes = nil
	file_gateway_gateway_proto_depIdxs = nil
}
