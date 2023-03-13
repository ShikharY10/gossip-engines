// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: schema/schema.proto

package schema

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

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Payload) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DeliveryPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  []byte `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	TargetId string `protobuf:"bytes,2,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
}

func (x *DeliveryPacket) Reset() {
	*x = DeliveryPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryPacket) ProtoMessage() {}

func (x *DeliveryPacket) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryPacket.ProtoReflect.Descriptor instead.
func (*DeliveryPacket) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryPacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *DeliveryPacket) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

//SubTYpe////////////////////////////////////////////////////
type PushNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *PushNotification) Reset() {
	*x = PushNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushNotification) ProtoMessage() {}

func (x *PushNotification) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushNotification.ProtoReflect.Descriptor instead.
func (*PushNotification) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{2}
}

func (x *PushNotification) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PushNotification) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type MakePartnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  string `protobuf:"bytes,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	TransferId string `protobuf:"bytes,2,opt,name=TransferId,proto3" json:"TransferId,omitempty"`
	PayloadKey string `protobuf:"bytes,3,opt,name=PayloadKey,proto3" json:"PayloadKey,omitempty"`
	SenderId   string `protobuf:"bytes,4,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	TargetId   string `protobuf:"bytes,5,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	PublicKey  string `protobuf:"bytes,6,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Token      string `protobuf:"bytes,7,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *MakePartnerRequest) Reset() {
	*x = MakePartnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakePartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakePartnerRequest) ProtoMessage() {}

func (x *MakePartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakePartnerRequest.ProtoReflect.Descriptor instead.
func (*MakePartnerRequest) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{3}
}

func (x *MakePartnerRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *MakePartnerRequest) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *MakePartnerRequest) GetPayloadKey() string {
	if x != nil {
		return x.PayloadKey
	}
	return ""
}

func (x *MakePartnerRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *MakePartnerRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *MakePartnerRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *MakePartnerRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type MakePartnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseId string `protobuf:"bytes,1,opt,name=ResponseId,proto3" json:"ResponseId,omitempty"`
	TransferId string `protobuf:"bytes,2,opt,name=TransferId,proto3" json:"TransferId,omitempty"`
	PayloadKey string `protobuf:"bytes,3,opt,name=PayloadKey,proto3" json:"PayloadKey,omitempty"`
	SenderId   string `protobuf:"bytes,4,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	TargetId   string `protobuf:"bytes,5,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	IsAccepted bool   `protobuf:"varint,6,opt,name=IsAccepted,proto3" json:"IsAccepted,omitempty"`
	AesKey     string `protobuf:"bytes,7,opt,name=AesKey,proto3" json:"AesKey,omitempty"`
	Token      string `protobuf:"bytes,8,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *MakePartnerResponse) Reset() {
	*x = MakePartnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakePartnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakePartnerResponse) ProtoMessage() {}

func (x *MakePartnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakePartnerResponse.ProtoReflect.Descriptor instead.
func (*MakePartnerResponse) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{4}
}

func (x *MakePartnerResponse) GetResponseId() string {
	if x != nil {
		return x.ResponseId
	}
	return ""
}

func (x *MakePartnerResponse) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *MakePartnerResponse) GetPayloadKey() string {
	if x != nil {
		return x.PayloadKey
	}
	return ""
}

func (x *MakePartnerResponse) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *MakePartnerResponse) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *MakePartnerResponse) GetIsAccepted() bool {
	if x != nil {
		return x.IsAccepted
	}
	return false
}

func (x *MakePartnerResponse) GetAesKey() string {
	if x != nil {
		return x.AesKey
	}
	return ""
}

func (x *MakePartnerResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type NewMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string `protobuf:"bytes,1,opt,name=TransferId,proto3" json:"TransferId,omitempty"`
	PayloadKey string `protobuf:"bytes,2,opt,name=PayloadKey,proto3" json:"PayloadKey,omitempty"`
	Data       string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	SenderId   string `protobuf:"bytes,4,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	TargetId   string `protobuf:"bytes,5,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
}

func (x *NewMessage) Reset() {
	*x = NewMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMessage) ProtoMessage() {}

func (x *NewMessage) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMessage.ProtoReflect.Descriptor instead.
func (*NewMessage) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{5}
}

func (x *NewMessage) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *NewMessage) GetPayloadKey() string {
	if x != nil {
		return x.PayloadKey
	}
	return ""
}

func (x *NewMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *NewMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *NewMessage) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

type RemovePartnerNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemovedId string `protobuf:"bytes,1,opt,name=RemovedId,proto3" json:"RemovedId,omitempty"`
	RemoverID string `protobuf:"bytes,2,opt,name=RemoverID,proto3" json:"RemoverID,omitempty"`
}

func (x *RemovePartnerNotification) Reset() {
	*x = RemovePartnerNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePartnerNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePartnerNotification) ProtoMessage() {}

func (x *RemovePartnerNotification) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePartnerNotification.ProtoReflect.Descriptor instead.
func (*RemovePartnerNotification) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{6}
}

func (x *RemovePartnerNotification) GetRemovedId() string {
	if x != nil {
		return x.RemovedId
	}
	return ""
}

func (x *RemovePartnerNotification) GetRemoverID() string {
	if x != nil {
		return x.RemoverID
	}
	return ""
}

//SubType/////////////////////////////////////////////////////
type Messaging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Messaging) Reset() {
	*x = Messaging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Messaging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Messaging) ProtoMessage() {}

func (x *Messaging) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Messaging.ProtoReflect.Descriptor instead.
func (*Messaging) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{7}
}

func (x *Messaging) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Messaging) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PayloadAcknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string `protobuf:"bytes,1,opt,name=TransferId,proto3" json:"TransferId,omitempty"`
	PayloadKey string `protobuf:"bytes,2,opt,name=PayloadKey,proto3" json:"PayloadKey,omitempty"`
}

func (x *PayloadAcknowledgement) Reset() {
	*x = PayloadAcknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadAcknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadAcknowledgement) ProtoMessage() {}

func (x *PayloadAcknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadAcknowledgement.ProtoReflect.Descriptor instead.
func (*PayloadAcknowledgement) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{8}
}

func (x *PayloadAcknowledgement) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *PayloadAcknowledgement) GetPayloadKey() string {
	if x != nil {
		return x.PayloadKey
	}
	return ""
}

type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{9}
}

func (x *Partner) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Partner) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_schema_schema_proto protoreflect.FileDescriptor

var file_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x31, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xfb, 0x01, 0x0a, 0x13, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x19,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x4b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_schema_proto_rawDescOnce sync.Once
	file_schema_schema_proto_rawDescData = file_schema_schema_proto_rawDesc
)

func file_schema_schema_proto_rawDescGZIP() []byte {
	file_schema_schema_proto_rawDescOnce.Do(func() {
		file_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_schema_proto_rawDescData)
	})
	return file_schema_schema_proto_rawDescData
}

var file_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_schema_schema_proto_goTypes = []interface{}{
	(*Payload)(nil),                   // 0: main.Payload
	(*DeliveryPacket)(nil),            // 1: main.DeliveryPacket
	(*PushNotification)(nil),          // 2: main.PushNotification
	(*MakePartnerRequest)(nil),        // 3: main.MakePartnerRequest
	(*MakePartnerResponse)(nil),       // 4: main.MakePartnerResponse
	(*NewMessage)(nil),                // 5: main.NewMessage
	(*RemovePartnerNotification)(nil), // 6: main.RemovePartnerNotification
	(*Messaging)(nil),                 // 7: main.Messaging
	(*PayloadAcknowledgement)(nil),    // 8: main.PayloadAcknowledgement
	(*Partner)(nil),                   // 9: main.Partner
}
var file_schema_schema_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_schema_proto_init() }
func file_schema_schema_proto_init() {
	if File_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryPacket); i {
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
		file_schema_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushNotification); i {
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
		file_schema_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakePartnerRequest); i {
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
		file_schema_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakePartnerResponse); i {
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
		file_schema_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMessage); i {
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
		file_schema_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePartnerNotification); i {
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
		file_schema_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Messaging); i {
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
		file_schema_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadAcknowledgement); i {
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
		file_schema_schema_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
			RawDescriptor: file_schema_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_schema_proto_goTypes,
		DependencyIndexes: file_schema_schema_proto_depIdxs,
		MessageInfos:      file_schema_schema_proto_msgTypes,
	}.Build()
	File_schema_schema_proto = out.File
	file_schema_schema_proto_rawDesc = nil
	file_schema_schema_proto_goTypes = nil
	file_schema_schema_proto_depIdxs = nil
}