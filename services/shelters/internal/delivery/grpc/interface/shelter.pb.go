// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: shelter.proto

package shelter

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

type ShelterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name        string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	OwnedBreeds []string             `protobuf:"bytes,4,rep,name=ownedBreeds,proto3" json:"ownedBreeds,omitempty"`
}

func (x *ShelterResponse) Reset() {
	*x = ShelterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShelterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShelterResponse) ProtoMessage() {}

func (x *ShelterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShelterResponse.ProtoReflect.Descriptor instead.
func (*ShelterResponse) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{0}
}

func (x *ShelterResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShelterResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ShelterResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShelterResponse) GetOwnedBreeds() []string {
	if x != nil {
		return x.OwnedBreeds
	}
	return nil
}

type CreateShelterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnedBreeds []string `protobuf:"bytes,2,rep,name=ownedBreeds,proto3" json:"ownedBreeds,omitempty"`
}

func (x *CreateShelterRequest) Reset() {
	*x = CreateShelterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelterRequest) ProtoMessage() {}

func (x *CreateShelterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelterRequest.ProtoReflect.Descriptor instead.
func (*CreateShelterRequest) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShelterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateShelterRequest) GetOwnedBreeds() []string {
	if x != nil {
		return x.OwnedBreeds
	}
	return nil
}

type UpdateShelterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnedBreeds []string `protobuf:"bytes,2,rep,name=ownedBreeds,proto3" json:"ownedBreeds,omitempty"`
}

func (x *UpdateShelterRequest) Reset() {
	*x = UpdateShelterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShelterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShelterRequest) ProtoMessage() {}

func (x *UpdateShelterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShelterRequest.ProtoReflect.Descriptor instead.
func (*UpdateShelterRequest) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateShelterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateShelterRequest) GetOwnedBreeds() []string {
	if x != nil {
		return x.OwnedBreeds
	}
	return nil
}

type UpdateShelterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *ShelterResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateShelterResponse) Reset() {
	*x = UpdateShelterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShelterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShelterResponse) ProtoMessage() {}

func (x *UpdateShelterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShelterResponse.ProtoReflect.Descriptor instead.
func (*UpdateShelterResponse) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateShelterResponse) GetResponse() *ShelterResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type CreateShelterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *ShelterResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateShelterResponse) Reset() {
	*x = CreateShelterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelterResponse) ProtoMessage() {}

func (x *CreateShelterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelterResponse.ProtoReflect.Descriptor instead.
func (*CreateShelterResponse) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{4}
}

func (x *CreateShelterResponse) GetResponse() *ShelterResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteShelterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteShelterRequest) Reset() {
	*x = DeleteShelterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelterRequest) ProtoMessage() {}

func (x *DeleteShelterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelterRequest.ProtoReflect.Descriptor instead.
func (*DeleteShelterRequest) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShelterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteShelterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteShelterResponse) Reset() {
	*x = DeleteShelterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelterResponse) ProtoMessage() {}

func (x *DeleteShelterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelterResponse.ProtoReflect.Descriptor instead.
func (*DeleteShelterResponse) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteShelterResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetShelterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetShelterRequest) Reset() {
	*x = GetShelterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelterRequest) ProtoMessage() {}

func (x *GetShelterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelterRequest.ProtoReflect.Descriptor instead.
func (*GetShelterRequest) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{7}
}

func (x *GetShelterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetShelterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *ShelterResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetShelterResponse) Reset() {
	*x = GetShelterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelterResponse) ProtoMessage() {}

func (x *GetShelterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelterResponse.ProtoReflect.Descriptor instead.
func (*GetShelterResponse) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{8}
}

func (x *GetShelterResponse) GetResponse() *ShelterResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type ShelterFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Sort         string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	SortSafelist []string `protobuf:"bytes,4,rep,name=sortSafelist,proto3" json:"sortSafelist,omitempty"`
}

func (x *ShelterFilters) Reset() {
	*x = ShelterFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShelterFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShelterFilters) ProtoMessage() {}

func (x *ShelterFilters) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShelterFilters.ProtoReflect.Descriptor instead.
func (*ShelterFilters) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{9}
}

func (x *ShelterFilters) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ShelterFilters) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ShelterFilters) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ShelterFilters) GetSortSafelist() []string {
	if x != nil {
		return x.SortSafelist
	}
	return nil
}

type GetSheltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnedBreeds []string        `protobuf:"bytes,2,rep,name=ownedBreeds,proto3" json:"ownedBreeds,omitempty"`
	Filters     *ShelterFilters `protobuf:"bytes,3,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *GetSheltersRequest) Reset() {
	*x = GetSheltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSheltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSheltersRequest) ProtoMessage() {}

func (x *GetSheltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSheltersRequest.ProtoReflect.Descriptor instead.
func (*GetSheltersRequest) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{10}
}

func (x *GetSheltersRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSheltersRequest) GetOwnedBreeds() []string {
	if x != nil {
		return x.OwnedBreeds
	}
	return nil
}

func (x *GetSheltersRequest) GetFilters() *ShelterFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

type GetSheltersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shelters []*ShelterResponse `protobuf:"bytes,1,rep,name=shelters,proto3" json:"shelters,omitempty"`
}

func (x *GetSheltersResponse) Reset() {
	*x = GetSheltersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shelter_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSheltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSheltersResponse) ProtoMessage() {}

func (x *GetSheltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shelter_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSheltersResponse.ProtoReflect.Descriptor instead.
func (*GetSheltersResponse) Descriptor() ([]byte, []int) {
	return file_shelter_proto_rawDescGZIP(), []int{11}
}

func (x *GetSheltersResponse) GetShelters() []*ShelterResponse {
	if x != nil {
		return x.Shelters
	}
	return nil
}

var File_shelter_proto protoreflect.FileDescriptor

var file_shelter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0f, 0x53, 0x68,
	0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x22, 0x4c,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77,
	0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x22, 0x4c, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65,
	0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x77, 0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x78, 0x0a, 0x0e, 0x53, 0x68,
	0x65, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x66, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x66, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x7d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x72, 0x65, 0x65, 0x64, 0x73,
	0x12, 0x31, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x65, 0x6c,
	0x74, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x68,
	0x65, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x32, 0x95, 0x03, 0x0a, 0x0e, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x68,
	0x65, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x73,
	0x68, 0x65, 0x6c, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shelter_proto_rawDescOnce sync.Once
	file_shelter_proto_rawDescData = file_shelter_proto_rawDesc
)

func file_shelter_proto_rawDescGZIP() []byte {
	file_shelter_proto_rawDescOnce.Do(func() {
		file_shelter_proto_rawDescData = protoimpl.X.CompressGZIP(file_shelter_proto_rawDescData)
	})
	return file_shelter_proto_rawDescData
}

var file_shelter_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_shelter_proto_goTypes = []interface{}{
	(*ShelterResponse)(nil),       // 0: shelter.ShelterResponse
	(*CreateShelterRequest)(nil),  // 1: shelter.CreateShelterRequest
	(*UpdateShelterRequest)(nil),  // 2: shelter.UpdateShelterRequest
	(*UpdateShelterResponse)(nil), // 3: shelter.UpdateShelterResponse
	(*CreateShelterResponse)(nil), // 4: shelter.CreateShelterResponse
	(*DeleteShelterRequest)(nil),  // 5: shelter.DeleteShelterRequest
	(*DeleteShelterResponse)(nil), // 6: shelter.DeleteShelterResponse
	(*GetShelterRequest)(nil),     // 7: shelter.GetShelterRequest
	(*GetShelterResponse)(nil),    // 8: shelter.GetShelterResponse
	(*ShelterFilters)(nil),        // 9: shelter.ShelterFilters
	(*GetSheltersRequest)(nil),    // 10: shelter.GetSheltersRequest
	(*GetSheltersResponse)(nil),   // 11: shelter.GetSheltersResponse
	(*timestamp.Timestamp)(nil),   // 12: google.protobuf.Timestamp
}
var file_shelter_proto_depIdxs = []int32{
	12, // 0: shelter.ShelterResponse.created_at:type_name -> google.protobuf.Timestamp
	0,  // 1: shelter.UpdateShelterResponse.response:type_name -> shelter.ShelterResponse
	0,  // 2: shelter.CreateShelterResponse.response:type_name -> shelter.ShelterResponse
	0,  // 3: shelter.GetShelterResponse.response:type_name -> shelter.ShelterResponse
	9,  // 4: shelter.GetSheltersRequest.filters:type_name -> shelter.ShelterFilters
	0,  // 5: shelter.GetSheltersResponse.shelters:type_name -> shelter.ShelterResponse
	1,  // 6: shelter.ShelterService.CreateShelter:input_type -> shelter.CreateShelterRequest
	2,  // 7: shelter.ShelterService.UpdateShelter:input_type -> shelter.UpdateShelterRequest
	5,  // 8: shelter.ShelterService.DeleteShelter:input_type -> shelter.DeleteShelterRequest
	7,  // 9: shelter.ShelterService.GetShelter:input_type -> shelter.GetShelterRequest
	10, // 10: shelter.ShelterService.GetShelters:input_type -> shelter.GetSheltersRequest
	4,  // 11: shelter.ShelterService.CreateShelter:output_type -> shelter.CreateShelterResponse
	0,  // 12: shelter.ShelterService.UpdateShelter:output_type -> shelter.ShelterResponse
	6,  // 13: shelter.ShelterService.DeleteShelter:output_type -> shelter.DeleteShelterResponse
	8,  // 14: shelter.ShelterService.GetShelter:output_type -> shelter.GetShelterResponse
	11, // 15: shelter.ShelterService.GetShelters:output_type -> shelter.GetSheltersResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_shelter_proto_init() }
func file_shelter_proto_init() {
	if File_shelter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shelter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShelterResponse); i {
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
		file_shelter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelterRequest); i {
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
		file_shelter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateShelterRequest); i {
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
		file_shelter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateShelterResponse); i {
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
		file_shelter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelterResponse); i {
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
		file_shelter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelterRequest); i {
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
		file_shelter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelterResponse); i {
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
		file_shelter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelterRequest); i {
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
		file_shelter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelterResponse); i {
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
		file_shelter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShelterFilters); i {
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
		file_shelter_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSheltersRequest); i {
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
		file_shelter_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSheltersResponse); i {
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
			RawDescriptor: file_shelter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shelter_proto_goTypes,
		DependencyIndexes: file_shelter_proto_depIdxs,
		MessageInfos:      file_shelter_proto_msgTypes,
	}.Build()
	File_shelter_proto = out.File
	file_shelter_proto_rawDesc = nil
	file_shelter_proto_goTypes = nil
	file_shelter_proto_depIdxs = nil
}
