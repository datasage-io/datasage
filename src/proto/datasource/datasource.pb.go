// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: datasource.proto

package datasource

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//Request for List Datasource
type ListDatasourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     int64  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	User     string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ListDatasourceRequest) Reset() {
	*x = ListDatasourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasourceRequest) ProtoMessage() {}

func (x *ListDatasourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasourceRequest.ProtoReflect.Descriptor instead.
func (*ListDatasourceRequest) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{0}
}

func (x *ListDatasourceRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ListDatasourceRequest) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ListDatasourceRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ListDatasourceRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

//Response for List Datasource
type ListDatasourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListAllDatasources []*ListAllDatasource `protobuf:"bytes,1,rep,name=list_all_datasources,json=listAllDatasources,proto3" json:"list_all_datasources,omitempty"`
}

func (x *ListDatasourceResponse) Reset() {
	*x = ListDatasourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasourceResponse) ProtoMessage() {}

func (x *ListDatasourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasourceResponse.ProtoReflect.Descriptor instead.
func (*ListDatasourceResponse) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{1}
}

func (x *ListDatasourceResponse) GetListAllDatasources() []*ListAllDatasource {
	if x != nil {
		return x.ListAllDatasources
	}
	return nil
}

//Request for Delete Datasource
type DeleteDatasourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDatasourceRequest) Reset() {
	*x = DeleteDatasourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatasourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatasourceRequest) ProtoMessage() {}

func (x *DeleteDatasourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatasourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatasourceRequest) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDatasourceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//Response for Delete Datasource
type DeleteDatasourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteDatasourceResponse) Reset() {
	*x = DeleteDatasourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatasourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatasourceResponse) ProtoMessage() {}

func (x *DeleteDatasourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatasourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteDatasourceResponse) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDatasourceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListAllDatasource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataDomain    string               `protobuf:"bytes,1,opt,name=data_domain,json=dataDomain,proto3" json:"data_domain,omitempty"`
	Name          string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DsDescription string               `protobuf:"bytes,3,opt,name=ds_description,json=dsDescription,proto3" json:"ds_description,omitempty"`
	DsType        string               `protobuf:"bytes,4,opt,name=ds_type,json=dsType,proto3" json:"ds_type,omitempty"`
	DsVersion     string               `protobuf:"bytes,5,opt,name=ds_version,json=dsVersion,proto3" json:"ds_version,omitempty"`
	DsKey         string               `protobuf:"bytes,6,opt,name=ds_key,json=dsKey,proto3" json:"ds_key,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Deleted       *timestamp.Timestamp `protobuf:"bytes,8,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *ListAllDatasource) Reset() {
	*x = ListAllDatasource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllDatasource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllDatasource) ProtoMessage() {}

func (x *ListAllDatasource) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllDatasource.ProtoReflect.Descriptor instead.
func (*ListAllDatasource) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllDatasource) GetDataDomain() string {
	if x != nil {
		return x.DataDomain
	}
	return ""
}

func (x *ListAllDatasource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListAllDatasource) GetDsDescription() string {
	if x != nil {
		return x.DsDescription
	}
	return ""
}

func (x *ListAllDatasource) GetDsType() string {
	if x != nil {
		return x.DsType
	}
	return ""
}

func (x *ListAllDatasource) GetDsVersion() string {
	if x != nil {
		return x.DsVersion
	}
	return ""
}

func (x *ListAllDatasource) GetDsKey() string {
	if x != nil {
		return x.DsKey
	}
	return ""
}

func (x *ListAllDatasource) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListAllDatasource) GetDeleted() *timestamp.Timestamp {
	if x != nil {
		return x.Deleted
	}
	return nil
}

var File_datasource_proto protoreflect.FileDescriptor

var file_datasource_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6f, 0x0a, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x69, 0x0a, 0x16, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x14, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x17, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xaf, 0x02, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x73, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xc6,
	0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x69,
	0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_datasource_proto_rawDescOnce sync.Once
	file_datasource_proto_rawDescData = file_datasource_proto_rawDesc
)

func file_datasource_proto_rawDescGZIP() []byte {
	file_datasource_proto_rawDescOnce.Do(func() {
		file_datasource_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasource_proto_rawDescData)
	})
	return file_datasource_proto_rawDescData
}

var file_datasource_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_datasource_proto_goTypes = []interface{}{
	(*ListDatasourceRequest)(nil),    // 0: datasource.listDatasourceRequest
	(*ListDatasourceResponse)(nil),   // 1: datasource.listDatasourceResponse
	(*DeleteDatasourceRequest)(nil),  // 2: datasource.deleteDatasourceRequest
	(*DeleteDatasourceResponse)(nil), // 3: datasource.deleteDatasourceResponse
	(*ListAllDatasource)(nil),        // 4: datasource.ListAllDatasource
	(*timestamp.Timestamp)(nil),      // 5: google.protobuf.Timestamp
}
var file_datasource_proto_depIdxs = []int32{
	4, // 0: datasource.listDatasourceResponse.list_all_datasources:type_name -> datasource.ListAllDatasource
	5, // 1: datasource.ListAllDatasource.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: datasource.ListAllDatasource.deleted:type_name -> google.protobuf.Timestamp
	0, // 3: datasource.Datasource.ListDatasources:input_type -> datasource.listDatasourceRequest
	2, // 4: datasource.Datasource.DeleteDatasources:input_type -> datasource.deleteDatasourceRequest
	1, // 5: datasource.Datasource.ListDatasources:output_type -> datasource.listDatasourceResponse
	3, // 6: datasource.Datasource.DeleteDatasources:output_type -> datasource.deleteDatasourceResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_datasource_proto_init() }
func file_datasource_proto_init() {
	if File_datasource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasourceRequest); i {
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
		file_datasource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasourceResponse); i {
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
		file_datasource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatasourceRequest); i {
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
		file_datasource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatasourceResponse); i {
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
		file_datasource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllDatasource); i {
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
			RawDescriptor: file_datasource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datasource_proto_goTypes,
		DependencyIndexes: file_datasource_proto_depIdxs,
		MessageInfos:      file_datasource_proto_msgTypes,
	}.Build()
	File_datasource_proto = out.File
	file_datasource_proto_rawDesc = nil
	file_datasource_proto_goTypes = nil
	file_datasource_proto_depIdxs = nil
}
