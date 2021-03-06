// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person.proto

package example

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_NO_STATUS Status = 0
	Status_PENDING   Status = 1
	Status_ACTIVE    Status = 2
	Status_EXPIRED   Status = 3
)

var Status_name = map[int32]string{
	0: "NO_STATUS",
	1: "PENDING",
	2: "ACTIVE",
	3: "EXPIRED",
}

var Status_value = map[string]int32{
	"NO_STATUS": 0,
	"PENDING":   1,
	"ACTIVE":    2,
	"EXPIRED":   3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

type CreatePersonRequest struct {
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=example.Status" json:"status,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *CreatePersonRequest) Reset()         { *m = CreatePersonRequest{} }
func (m *CreatePersonRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePersonRequest) ProtoMessage()    {}
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}
func (m *CreatePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePersonRequest.Unmarshal(m, b)
}
func (m *CreatePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePersonRequest.Marshal(b, m, deterministic)
}
func (m *CreatePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePersonRequest.Merge(m, src)
}
func (m *CreatePersonRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePersonRequest.Size(m)
}
func (m *CreatePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePersonRequest proto.InternalMessageInfo

func (m *CreatePersonRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_NO_STATUS
}

func (m *CreatePersonRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePersonRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type GetPersonRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetPersonRequest) Reset()         { *m = GetPersonRequest{} }
func (m *GetPersonRequest) String() string { return proto.CompactTextString(m) }
func (*GetPersonRequest) ProtoMessage()    {}
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}
func (m *GetPersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPersonRequest.Unmarshal(m, b)
}
func (m *GetPersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPersonRequest.Marshal(b, m, deterministic)
}
func (m *GetPersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPersonRequest.Merge(m, src)
}
func (m *GetPersonRequest) XXX_Size() int {
	return xxx_messageInfo_GetPersonRequest.Size(m)
}
func (m *GetPersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPersonRequest proto.InternalMessageInfo

func (m *GetPersonRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetPersonsRequest struct {
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=example.Status" json:"status,omitempty"`
	Skip   int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	Limit  int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *GetPersonsRequest) Reset()         { *m = GetPersonsRequest{} }
func (m *GetPersonsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPersonsRequest) ProtoMessage()    {}
func (*GetPersonsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{2}
}
func (m *GetPersonsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPersonsRequest.Unmarshal(m, b)
}
func (m *GetPersonsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPersonsRequest.Marshal(b, m, deterministic)
}
func (m *GetPersonsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPersonsRequest.Merge(m, src)
}
func (m *GetPersonsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPersonsRequest.Size(m)
}
func (m *GetPersonsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPersonsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPersonsRequest proto.InternalMessageInfo

func (m *GetPersonsRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_NO_STATUS
}

func (m *GetPersonsRequest) GetSkip() int64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *GetPersonsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type UpdatePersonRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=example.Status" json:"status,omitempty"`
}

func (m *UpdatePersonRequest) Reset()         { *m = UpdatePersonRequest{} }
func (m *UpdatePersonRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePersonRequest) ProtoMessage()    {}
func (*UpdatePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{3}
}
func (m *UpdatePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePersonRequest.Unmarshal(m, b)
}
func (m *UpdatePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePersonRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePersonRequest.Merge(m, src)
}
func (m *UpdatePersonRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePersonRequest.Size(m)
}
func (m *UpdatePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePersonRequest proto.InternalMessageInfo

func (m *UpdatePersonRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdatePersonRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePersonRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdatePersonRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_NO_STATUS
}

type DeletePersonRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *DeletePersonRequest) Reset()         { *m = DeletePersonRequest{} }
func (m *DeletePersonRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePersonRequest) ProtoMessage()    {}
func (*DeletePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{4}
}
func (m *DeletePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePersonRequest.Unmarshal(m, b)
}
func (m *DeletePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePersonRequest.Marshal(b, m, deterministic)
}
func (m *DeletePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePersonRequest.Merge(m, src)
}
func (m *DeletePersonRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePersonRequest.Size(m)
}
func (m *DeletePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePersonRequest proto.InternalMessageInfo

func (m *DeletePersonRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Person struct {
	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    Status               `protobuf:"varint,2,opt,name=status,proto3,enum=example.Status" json:"status,omitempty"`
	Name      string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email     string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{5}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_NO_STATUS
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Person) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Persons struct {
	Items []*Person `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *Persons) Reset()         { *m = Persons{} }
func (m *Persons) String() string { return proto.CompactTextString(m) }
func (*Persons) ProtoMessage()    {}
func (*Persons) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{6}
}
func (m *Persons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Persons.Unmarshal(m, b)
}
func (m *Persons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Persons.Marshal(b, m, deterministic)
}
func (m *Persons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Persons.Merge(m, src)
}
func (m *Persons) XXX_Size() int {
	return xxx_messageInfo_Persons.Size(m)
}
func (m *Persons) XXX_DiscardUnknown() {
	xxx_messageInfo_Persons.DiscardUnknown(m)
}

var xxx_messageInfo_Persons proto.InternalMessageInfo

func (m *Persons) GetItems() []*Person {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterEnum("example.Status", Status_name, Status_value)
	proto.RegisterType((*CreatePersonRequest)(nil), "example.CreatePersonRequest")
	proto.RegisterType((*GetPersonRequest)(nil), "example.GetPersonRequest")
	proto.RegisterType((*GetPersonsRequest)(nil), "example.GetPersonsRequest")
	proto.RegisterType((*UpdatePersonRequest)(nil), "example.UpdatePersonRequest")
	proto.RegisterType((*DeletePersonRequest)(nil), "example.DeletePersonRequest")
	proto.RegisterType((*Person)(nil), "example.Person")
	proto.RegisterType((*Persons)(nil), "example.Persons")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0xed, 0x9f, 0xdd, 0x12, 0x2e, 0xb0, 0xdb, 0x1d, 0x36, 0xbf, 0xf0, 0xab, 0x06, 0x49, 0xb3,
	0x46, 0x62, 0x42, 0x31, 0xe8, 0xc3, 0x1a, 0x35, 0x06, 0x16, 0xdc, 0xf0, 0x82, 0xa4, 0xb0, 0xc6,
	0xc4, 0x7f, 0x29, 0x74, 0xe8, 0x36, 0xb6, 0xb4, 0xb6, 0x03, 0x59, 0x62, 0xfc, 0x0a, 0xc6, 0x8f,
	0xe6, 0xa3, 0x3e, 0xf9, 0xc0, 0xb7, 0xf0, 0xc9, 0x74, 0xa6, 0x40, 0x05, 0x54, 0xb2, 0x6f, 0x73,
	0x7b, 0xee, 0xb9, 0x73, 0xee, 0x9d, 0x73, 0x0b, 0x59, 0x1f, 0x07, 0xa1, 0x37, 0xd6, 0xfc, 0xc0,
	0x23, 0x1e, 0x4a, 0xe1, 0x2b, 0xc3, 0xf5, 0x1d, 0xac, 0xd4, 0x2d, 0x9b, 0x5c, 0x4e, 0x06, 0xda,
	0xd0, 0x73, 0xab, 0x78, 0x3c, 0xf5, 0x66, 0x7e, 0xe0, 0x5d, 0xcd, 0xaa, 0x34, 0x6b, 0x58, 0xb1,
	0xf0, 0xb8, 0x32, 0x35, 0x1c, 0xdb, 0x34, 0x08, 0xae, 0x6e, 0x1c, 0x58, 0x2d, 0xa5, 0x92, 0x28,
	0x61, 0x79, 0x96, 0xc7, 0xc8, 0x83, 0xc9, 0x88, 0x46, 0x34, 0xa0, 0xa7, 0x38, 0xfd, 0x86, 0xe5,
	0x79, 0x96, 0x83, 0x57, 0x59, 0xd8, 0xf5, 0xc9, 0x2c, 0x06, 0x6f, 0xad, 0x83, 0xc4, 0x76, 0x71,
	0x48, 0x0c, 0xd7, 0x67, 0x09, 0xea, 0x25, 0xe4, 0xcf, 0x02, 0x6c, 0x10, 0xdc, 0xa5, 0xed, 0xe8,
	0xf8, 0xc3, 0x04, 0x87, 0x04, 0xdd, 0x01, 0x29, 0x24, 0x06, 0x99, 0x84, 0x05, 0xbe, 0xc4, 0x97,
	0x0f, 0x6a, 0x87, 0x5a, 0xdc, 0xa0, 0xd6, 0xa3, 0x9f, 0xf5, 0x18, 0x46, 0x08, 0xf6, 0xc6, 0x86,
	0x8b, 0x0b, 0x42, 0x89, 0x2f, 0xa7, 0x75, 0x7a, 0x46, 0xc7, 0xb0, 0x8f, 0x5d, 0xc3, 0x76, 0x0a,
	0x22, 0xfd, 0xc8, 0x02, 0xf5, 0x14, 0xe4, 0x73, 0x4c, 0x7e, 0xbf, 0xe6, 0x04, 0x04, 0xdb, 0xa4,
	0x57, 0xa4, 0x1b, 0xc7, 0x3f, 0x1b, 0x47, 0xc1, 0x61, 0x2d, 0xf7, 0xf6, 0x95, 0x51, 0x19, 0xbd,
	0x36, 0xdf, 0x7c, 0xac, 0x3d, 0xf8, 0x74, 0xa2, 0x0b, 0xb6, 0xa9, 0x8e, 0xe0, 0x68, 0xc9, 0x0c,
	0xaf, 0xa3, 0x30, 0x7c, 0x6f, 0xfb, 0x54, 0xa1, 0xa8, 0xd3, 0x73, 0xa4, 0xd0, 0xb1, 0x5d, 0x9b,
	0x50, 0x85, 0xa2, 0xce, 0x02, 0xf5, 0x33, 0x0f, 0xf9, 0x0b, 0xdf, 0xdc, 0x18, 0xc6, 0x4e, 0x2a,
	0x77, 0x9f, 0x44, 0x42, 0xfa, 0xde, 0x5f, 0xa5, 0xab, 0x8f, 0x20, 0xdf, 0xc4, 0x0e, 0xbe, 0x96,
	0x1e, 0xf5, 0x07, 0x0f, 0x12, 0xe3, 0xa1, 0x83, 0x15, 0x81, 0x4a, 0x5d, 0x09, 0x10, 0x76, 0x7b,
	0x5d, 0x71, 0x5b, 0x4f, 0x7b, 0xc9, 0x9e, 0x4e, 0x21, 0x3d, 0xa4, 0x3e, 0x32, 0xeb, 0xa4, 0xb0,
	0x5f, 0xe2, 0xcb, 0x99, 0x9a, 0xa2, 0x31, 0xf3, 0x69, 0x0b, 0xf3, 0x69, 0xfd, 0x85, 0xf9, 0xf4,
	0x55, 0x72, 0xc4, 0x9c, 0xd0, 0xa1, 0x47, 0x4c, 0xe9, 0xdf, 0xcc, 0x65, 0xb2, 0x7a, 0x0f, 0x52,
	0xb1, 0x29, 0xd0, 0x6d, 0xd8, 0xb7, 0x09, 0x76, 0x23, 0x33, 0x88, 0xe5, 0x4c, 0xa2, 0xa1, 0x78,
	0x72, 0x0c, 0xbd, 0xfb, 0x04, 0x24, 0xd6, 0x21, 0xca, 0x41, 0xba, 0xf3, 0xfc, 0x5d, 0xaf, 0x5f,
	0xef, 0x5f, 0xf4, 0x64, 0x0e, 0x65, 0x20, 0xd5, 0x6d, 0x75, 0x9a, 0xed, 0xce, 0xb9, 0xcc, 0x23,
	0x00, 0xa9, 0x7e, 0xd6, 0x6f, 0xbf, 0x68, 0xc9, 0x42, 0x04, 0xb4, 0x5e, 0x76, 0xdb, 0x7a, 0xab,
	0x29, 0x8b, 0xb5, 0xef, 0x02, 0xe4, 0x58, 0xc1, 0x1e, 0x0e, 0xa6, 0xf6, 0x10, 0xa3, 0xa7, 0x90,
	0x4d, 0xae, 0x0f, 0xba, 0xb9, 0xbc, 0x78, 0xcb, 0x56, 0x29, 0xeb, 0xb2, 0x54, 0x0e, 0x3d, 0x84,
	0xf4, 0xd2, 0xdb, 0xe8, 0xff, 0x25, 0xbe, 0xbe, 0x29, 0xdb, 0xa8, 0x8f, 0x01, 0x56, 0x6b, 0x81,
	0x94, 0x4d, 0xee, 0x62, 0x57, 0x14, 0x79, 0x8d, 0x1c, 0xaa, 0x1c, 0x7a, 0x06, 0xd9, 0xa4, 0xd7,
	0x13, 0xca, 0xb7, 0xac, 0x80, 0xf2, 0xdf, 0xc6, 0x8b, 0xb4, 0xa2, 0xbf, 0x0c, 0xab, 0x93, 0xf4,
	0x68, 0xa2, 0xce, 0x16, 0xeb, 0xfe, 0xb9, 0x4e, 0x23, 0xfb, 0x65, 0x5e, 0xe4, 0xbe, 0xce, 0x8b,
	0xdc, 0xb7, 0x79, 0x91, 0x1b, 0x48, 0x14, 0xbf, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x04, 0x60,
	0xe8, 0x65, 0x66, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonServiceClient interface {
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*Person, error)
	GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*Person, error)
	GetPersons(ctx context.Context, in *GetPersonsRequest, opts ...grpc.CallOption) (*Persons, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type personServiceClient struct {
	cc *grpc.ClientConn
}

func NewPersonServiceClient(cc *grpc.ClientConn) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/example.PersonService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/example.PersonService/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetPersons(ctx context.Context, in *GetPersonsRequest, opts ...grpc.CallOption) (*Persons, error) {
	out := new(Persons)
	err := c.cc.Invoke(ctx, "/example.PersonService/GetPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/example.PersonService/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/example.PersonService/DeletePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
type PersonServiceServer interface {
	CreatePerson(context.Context, *CreatePersonRequest) (*Person, error)
	GetPerson(context.Context, *GetPersonRequest) (*Person, error)
	GetPersons(context.Context, *GetPersonsRequest) (*Persons, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*empty.Empty, error)
	DeletePerson(context.Context, *DeletePersonRequest) (*empty.Empty, error)
}

// UnimplementedPersonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (*UnimplementedPersonServiceServer) CreatePerson(ctx context.Context, req *CreatePersonRequest) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (*UnimplementedPersonServiceServer) GetPerson(ctx context.Context, req *GetPersonRequest) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (*UnimplementedPersonServiceServer) GetPersons(ctx context.Context, req *GetPersonsRequest) (*Persons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersons not implemented")
}
func (*UnimplementedPersonServiceServer) UpdatePerson(ctx context.Context, req *UpdatePersonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (*UnimplementedPersonServiceServer) DeletePerson(ctx context.Context, req *DeletePersonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerson not implemented")
}

func RegisterPersonServiceServer(s *grpc.Server, srv PersonServiceServer) {
	s.RegisterService(&_PersonService_serviceDesc, srv)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PersonService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PersonService/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPerson(ctx, req.(*GetPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PersonService/GetPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPersons(ctx, req.(*GetPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PersonService/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PersonService/DeletePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).DeletePerson(ctx, req.(*DeletePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _PersonService_GetPerson_Handler,
		},
		{
			MethodName: "GetPersons",
			Handler:    _PersonService_GetPersons_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _PersonService_UpdatePerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _PersonService_DeletePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
