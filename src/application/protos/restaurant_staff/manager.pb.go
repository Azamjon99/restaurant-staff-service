// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager.proto

package restaurant_staff

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ManagerProfile struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	EntityId             string   `protobuf:"bytes,2,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl             string   `protobuf:"bytes,4,opt,name=image_url,proto3" json:"image_url,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,5,opt,name=phone_number,proto3" json:"phone_number,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagerProfile) Reset()         { *m = ManagerProfile{} }
func (m *ManagerProfile) String() string { return proto.CompactTextString(m) }
func (*ManagerProfile) ProtoMessage()    {}
func (*ManagerProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{0}
}

func (m *ManagerProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagerProfile.Unmarshal(m, b)
}
func (m *ManagerProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagerProfile.Marshal(b, m, deterministic)
}
func (m *ManagerProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagerProfile.Merge(m, src)
}
func (m *ManagerProfile) XXX_Size() int {
	return xxx_messageInfo_ManagerProfile.Size(m)
}
func (m *ManagerProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagerProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ManagerProfile proto.InternalMessageInfo

func (m *ManagerProfile) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

func (m *ManagerProfile) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *ManagerProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ManagerProfile) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *ManagerProfile) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *ManagerProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ManagerProfile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ManagerProfile) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ManagerRestaurantAssignment struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ManagerId            string   `protobuf:"bytes,2,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	RestaurantId         string   `protobuf:"bytes,3,opt,name=restaurant_id,proto3" json:"restaurant_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagerRestaurantAssignment) Reset()         { *m = ManagerRestaurantAssignment{} }
func (m *ManagerRestaurantAssignment) String() string { return proto.CompactTextString(m) }
func (*ManagerRestaurantAssignment) ProtoMessage()    {}
func (*ManagerRestaurantAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{1}
}

func (m *ManagerRestaurantAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagerRestaurantAssignment.Unmarshal(m, b)
}
func (m *ManagerRestaurantAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagerRestaurantAssignment.Marshal(b, m, deterministic)
}
func (m *ManagerRestaurantAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagerRestaurantAssignment.Merge(m, src)
}
func (m *ManagerRestaurantAssignment) XXX_Size() int {
	return xxx_messageInfo_ManagerRestaurantAssignment.Size(m)
}
func (m *ManagerRestaurantAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagerRestaurantAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ManagerRestaurantAssignment proto.InternalMessageInfo

func (m *ManagerRestaurantAssignment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ManagerRestaurantAssignment) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

func (m *ManagerRestaurantAssignment) GetRestaurantId() string {
	if m != nil {
		return m.RestaurantId
	}
	return ""
}

func (m *ManagerRestaurantAssignment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type RegisterManagerRequest struct {
	EntityId             string   `protobuf:"bytes,1,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
	RestaurantId         string   `protobuf:"bytes,2,opt,name=restaurant_id,proto3" json:"restaurant_id,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterManagerRequest) Reset()         { *m = RegisterManagerRequest{} }
func (m *RegisterManagerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterManagerRequest) ProtoMessage()    {}
func (*RegisterManagerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{2}
}

func (m *RegisterManagerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterManagerRequest.Unmarshal(m, b)
}
func (m *RegisterManagerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterManagerRequest.Marshal(b, m, deterministic)
}
func (m *RegisterManagerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterManagerRequest.Merge(m, src)
}
func (m *RegisterManagerRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterManagerRequest.Size(m)
}
func (m *RegisterManagerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterManagerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterManagerRequest proto.InternalMessageInfo

func (m *RegisterManagerRequest) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *RegisterManagerRequest) GetRestaurantId() string {
	if m != nil {
		return m.RestaurantId
	}
	return ""
}

func (m *RegisterManagerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterManagerRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterManagerResponse struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterManagerResponse) Reset()         { *m = RegisterManagerResponse{} }
func (m *RegisterManagerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterManagerResponse) ProtoMessage()    {}
func (*RegisterManagerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{3}
}

func (m *RegisterManagerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterManagerResponse.Unmarshal(m, b)
}
func (m *RegisterManagerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterManagerResponse.Marshal(b, m, deterministic)
}
func (m *RegisterManagerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterManagerResponse.Merge(m, src)
}
func (m *RegisterManagerResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterManagerResponse.Size(m)
}
func (m *RegisterManagerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterManagerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterManagerResponse proto.InternalMessageInfo

func (m *RegisterManagerResponse) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

type ChangeManagerPasswordRequest struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	CurrentPassword      string   `protobuf:"bytes,2,opt,name=current_password,proto3" json:"current_password,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=new_password,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeManagerPasswordRequest) Reset()         { *m = ChangeManagerPasswordRequest{} }
func (m *ChangeManagerPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeManagerPasswordRequest) ProtoMessage()    {}
func (*ChangeManagerPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{4}
}

func (m *ChangeManagerPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeManagerPasswordRequest.Unmarshal(m, b)
}
func (m *ChangeManagerPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeManagerPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ChangeManagerPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeManagerPasswordRequest.Merge(m, src)
}
func (m *ChangeManagerPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeManagerPasswordRequest.Size(m)
}
func (m *ChangeManagerPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeManagerPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeManagerPasswordRequest proto.InternalMessageInfo

func (m *ChangeManagerPasswordRequest) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

func (m *ChangeManagerPasswordRequest) GetCurrentPassword() string {
	if m != nil {
		return m.CurrentPassword
	}
	return ""
}

func (m *ChangeManagerPasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ChangeManagerPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeManagerPasswordResponse) Reset()         { *m = ChangeManagerPasswordResponse{} }
func (m *ChangeManagerPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ChangeManagerPasswordResponse) ProtoMessage()    {}
func (*ChangeManagerPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{5}
}

func (m *ChangeManagerPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeManagerPasswordResponse.Unmarshal(m, b)
}
func (m *ChangeManagerPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeManagerPasswordResponse.Marshal(b, m, deterministic)
}
func (m *ChangeManagerPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeManagerPasswordResponse.Merge(m, src)
}
func (m *ChangeManagerPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ChangeManagerPasswordResponse.Size(m)
}
func (m *ChangeManagerPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeManagerPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeManagerPasswordResponse proto.InternalMessageInfo

type LoginManagerRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginManagerRequest) Reset()         { *m = LoginManagerRequest{} }
func (m *LoginManagerRequest) String() string { return proto.CompactTextString(m) }
func (*LoginManagerRequest) ProtoMessage()    {}
func (*LoginManagerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{6}
}

func (m *LoginManagerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginManagerRequest.Unmarshal(m, b)
}
func (m *LoginManagerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginManagerRequest.Marshal(b, m, deterministic)
}
func (m *LoginManagerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginManagerRequest.Merge(m, src)
}
func (m *LoginManagerRequest) XXX_Size() int {
	return xxx_messageInfo_LoginManagerRequest.Size(m)
}
func (m *LoginManagerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginManagerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginManagerRequest proto.InternalMessageInfo

func (m *LoginManagerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginManagerRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginManagerResponse struct {
	Profile              *ManagerProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Token                string          `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoginManagerResponse) Reset()         { *m = LoginManagerResponse{} }
func (m *LoginManagerResponse) String() string { return proto.CompactTextString(m) }
func (*LoginManagerResponse) ProtoMessage()    {}
func (*LoginManagerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{7}
}

func (m *LoginManagerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginManagerResponse.Unmarshal(m, b)
}
func (m *LoginManagerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginManagerResponse.Marshal(b, m, deterministic)
}
func (m *LoginManagerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginManagerResponse.Merge(m, src)
}
func (m *LoginManagerResponse) XXX_Size() int {
	return xxx_messageInfo_LoginManagerResponse.Size(m)
}
func (m *LoginManagerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginManagerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginManagerResponse proto.InternalMessageInfo

func (m *LoginManagerResponse) GetProfile() *ManagerProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *LoginManagerResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetManagerProfileRequest struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagerProfileRequest) Reset()         { *m = GetManagerProfileRequest{} }
func (m *GetManagerProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetManagerProfileRequest) ProtoMessage()    {}
func (*GetManagerProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{8}
}

func (m *GetManagerProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagerProfileRequest.Unmarshal(m, b)
}
func (m *GetManagerProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagerProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetManagerProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagerProfileRequest.Merge(m, src)
}
func (m *GetManagerProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetManagerProfileRequest.Size(m)
}
func (m *GetManagerProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagerProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagerProfileRequest proto.InternalMessageInfo

func (m *GetManagerProfileRequest) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

type GetManagerProfileResponse struct {
	Profile              *ManagerProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetManagerProfileResponse) Reset()         { *m = GetManagerProfileResponse{} }
func (m *GetManagerProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetManagerProfileResponse) ProtoMessage()    {}
func (*GetManagerProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{9}
}

func (m *GetManagerProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagerProfileResponse.Unmarshal(m, b)
}
func (m *GetManagerProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagerProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetManagerProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagerProfileResponse.Merge(m, src)
}
func (m *GetManagerProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetManagerProfileResponse.Size(m)
}
func (m *GetManagerProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagerProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagerProfileResponse proto.InternalMessageInfo

func (m *GetManagerProfileResponse) GetProfile() *ManagerProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateManagerProfileRequest struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,proto3" json:"phone_number,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ImageUrl             string   `protobuf:"bytes,5,opt,name=image_url,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateManagerProfileRequest) Reset()         { *m = UpdateManagerProfileRequest{} }
func (m *UpdateManagerProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateManagerProfileRequest) ProtoMessage()    {}
func (*UpdateManagerProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{10}
}

func (m *UpdateManagerProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateManagerProfileRequest.Unmarshal(m, b)
}
func (m *UpdateManagerProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateManagerProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateManagerProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateManagerProfileRequest.Merge(m, src)
}
func (m *UpdateManagerProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateManagerProfileRequest.Size(m)
}
func (m *UpdateManagerProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateManagerProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateManagerProfileRequest proto.InternalMessageInfo

func (m *UpdateManagerProfileRequest) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

func (m *UpdateManagerProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateManagerProfileRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UpdateManagerProfileRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateManagerProfileRequest) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type UpdateManagerProfileResponse struct {
	Profile              *ManagerProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateManagerProfileResponse) Reset()         { *m = UpdateManagerProfileResponse{} }
func (m *UpdateManagerProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateManagerProfileResponse) ProtoMessage()    {}
func (*UpdateManagerProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{11}
}

func (m *UpdateManagerProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateManagerProfileResponse.Unmarshal(m, b)
}
func (m *UpdateManagerProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateManagerProfileResponse.Marshal(b, m, deterministic)
}
func (m *UpdateManagerProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateManagerProfileResponse.Merge(m, src)
}
func (m *UpdateManagerProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateManagerProfileResponse.Size(m)
}
func (m *UpdateManagerProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateManagerProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateManagerProfileResponse proto.InternalMessageInfo

func (m *UpdateManagerProfileResponse) GetProfile() *ManagerProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type CreateManagerRestaurantAssignmentRequest struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	RestaurantId         string   `protobuf:"bytes,2,opt,name=restaurant_id,proto3" json:"restaurant_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateManagerRestaurantAssignmentRequest) Reset() {
	*m = CreateManagerRestaurantAssignmentRequest{}
}
func (m *CreateManagerRestaurantAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateManagerRestaurantAssignmentRequest) ProtoMessage()    {}
func (*CreateManagerRestaurantAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{12}
}

func (m *CreateManagerRestaurantAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateManagerRestaurantAssignmentRequest.Unmarshal(m, b)
}
func (m *CreateManagerRestaurantAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateManagerRestaurantAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *CreateManagerRestaurantAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateManagerRestaurantAssignmentRequest.Merge(m, src)
}
func (m *CreateManagerRestaurantAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateManagerRestaurantAssignmentRequest.Size(m)
}
func (m *CreateManagerRestaurantAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateManagerRestaurantAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateManagerRestaurantAssignmentRequest proto.InternalMessageInfo

func (m *CreateManagerRestaurantAssignmentRequest) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

func (m *CreateManagerRestaurantAssignmentRequest) GetRestaurantId() string {
	if m != nil {
		return m.RestaurantId
	}
	return ""
}

type CreateManagerRestaurantAssignmentResponse struct {
	AssigmentId          int64    `protobuf:"varint,1,opt,name=assigment_id,proto3" json:"assigment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateManagerRestaurantAssignmentResponse) Reset() {
	*m = CreateManagerRestaurantAssignmentResponse{}
}
func (m *CreateManagerRestaurantAssignmentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateManagerRestaurantAssignmentResponse) ProtoMessage()    {}
func (*CreateManagerRestaurantAssignmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{13}
}

func (m *CreateManagerRestaurantAssignmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateManagerRestaurantAssignmentResponse.Unmarshal(m, b)
}
func (m *CreateManagerRestaurantAssignmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateManagerRestaurantAssignmentResponse.Marshal(b, m, deterministic)
}
func (m *CreateManagerRestaurantAssignmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateManagerRestaurantAssignmentResponse.Merge(m, src)
}
func (m *CreateManagerRestaurantAssignmentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateManagerRestaurantAssignmentResponse.Size(m)
}
func (m *CreateManagerRestaurantAssignmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateManagerRestaurantAssignmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateManagerRestaurantAssignmentResponse proto.InternalMessageInfo

func (m *CreateManagerRestaurantAssignmentResponse) GetAssigmentId() int64 {
	if m != nil {
		return m.AssigmentId
	}
	return 0
}

type RemoveManagerRestaurantAssignmentRequest struct {
	AssigmentId          int64    `protobuf:"varint,1,opt,name=assigment_id,proto3" json:"assigment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveManagerRestaurantAssignmentRequest) Reset() {
	*m = RemoveManagerRestaurantAssignmentRequest{}
}
func (m *RemoveManagerRestaurantAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveManagerRestaurantAssignmentRequest) ProtoMessage()    {}
func (*RemoveManagerRestaurantAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{14}
}

func (m *RemoveManagerRestaurantAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveManagerRestaurantAssignmentRequest.Unmarshal(m, b)
}
func (m *RemoveManagerRestaurantAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveManagerRestaurantAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *RemoveManagerRestaurantAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveManagerRestaurantAssignmentRequest.Merge(m, src)
}
func (m *RemoveManagerRestaurantAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveManagerRestaurantAssignmentRequest.Size(m)
}
func (m *RemoveManagerRestaurantAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveManagerRestaurantAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveManagerRestaurantAssignmentRequest proto.InternalMessageInfo

func (m *RemoveManagerRestaurantAssignmentRequest) GetAssigmentId() int64 {
	if m != nil {
		return m.AssigmentId
	}
	return 0
}

type RemoveManagerRestaurantAssignmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveManagerRestaurantAssignmentResponse) Reset() {
	*m = RemoveManagerRestaurantAssignmentResponse{}
}
func (m *RemoveManagerRestaurantAssignmentResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveManagerRestaurantAssignmentResponse) ProtoMessage()    {}
func (*RemoveManagerRestaurantAssignmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{15}
}

func (m *RemoveManagerRestaurantAssignmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveManagerRestaurantAssignmentResponse.Unmarshal(m, b)
}
func (m *RemoveManagerRestaurantAssignmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveManagerRestaurantAssignmentResponse.Marshal(b, m, deterministic)
}
func (m *RemoveManagerRestaurantAssignmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveManagerRestaurantAssignmentResponse.Merge(m, src)
}
func (m *RemoveManagerRestaurantAssignmentResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveManagerRestaurantAssignmentResponse.Size(m)
}
func (m *RemoveManagerRestaurantAssignmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveManagerRestaurantAssignmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveManagerRestaurantAssignmentResponse proto.InternalMessageInfo

type GetManagerRestaurantRequest struct {
	ManagerId            string   `protobuf:"bytes,1,opt,name=manager_id,proto3" json:"manager_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagerRestaurantRequest) Reset()         { *m = GetManagerRestaurantRequest{} }
func (m *GetManagerRestaurantRequest) String() string { return proto.CompactTextString(m) }
func (*GetManagerRestaurantRequest) ProtoMessage()    {}
func (*GetManagerRestaurantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{16}
}

func (m *GetManagerRestaurantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagerRestaurantRequest.Unmarshal(m, b)
}
func (m *GetManagerRestaurantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagerRestaurantRequest.Marshal(b, m, deterministic)
}
func (m *GetManagerRestaurantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagerRestaurantRequest.Merge(m, src)
}
func (m *GetManagerRestaurantRequest) XXX_Size() int {
	return xxx_messageInfo_GetManagerRestaurantRequest.Size(m)
}
func (m *GetManagerRestaurantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagerRestaurantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagerRestaurantRequest proto.InternalMessageInfo

func (m *GetManagerRestaurantRequest) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

type GetManagerRestaurantResponse struct {
	RestaurantId         string   `protobuf:"bytes,1,opt,name=restaurant_id,proto3" json:"restaurant_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagerRestaurantResponse) Reset()         { *m = GetManagerRestaurantResponse{} }
func (m *GetManagerRestaurantResponse) String() string { return proto.CompactTextString(m) }
func (*GetManagerRestaurantResponse) ProtoMessage()    {}
func (*GetManagerRestaurantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{17}
}

func (m *GetManagerRestaurantResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagerRestaurantResponse.Unmarshal(m, b)
}
func (m *GetManagerRestaurantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagerRestaurantResponse.Marshal(b, m, deterministic)
}
func (m *GetManagerRestaurantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagerRestaurantResponse.Merge(m, src)
}
func (m *GetManagerRestaurantResponse) XXX_Size() int {
	return xxx_messageInfo_GetManagerRestaurantResponse.Size(m)
}
func (m *GetManagerRestaurantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagerRestaurantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagerRestaurantResponse proto.InternalMessageInfo

func (m *GetManagerRestaurantResponse) GetRestaurantId() string {
	if m != nil {
		return m.RestaurantId
	}
	return ""
}

func init() {
	proto.RegisterType((*ManagerProfile)(nil), "ManagerProfile")
	proto.RegisterType((*ManagerRestaurantAssignment)(nil), "ManagerRestaurantAssignment")
	proto.RegisterType((*RegisterManagerRequest)(nil), "RegisterManagerRequest")
	proto.RegisterType((*RegisterManagerResponse)(nil), "RegisterManagerResponse")
	proto.RegisterType((*ChangeManagerPasswordRequest)(nil), "ChangeManagerPasswordRequest")
	proto.RegisterType((*ChangeManagerPasswordResponse)(nil), "ChangeManagerPasswordResponse")
	proto.RegisterType((*LoginManagerRequest)(nil), "LoginManagerRequest")
	proto.RegisterType((*LoginManagerResponse)(nil), "LoginManagerResponse")
	proto.RegisterType((*GetManagerProfileRequest)(nil), "GetManagerProfileRequest")
	proto.RegisterType((*GetManagerProfileResponse)(nil), "GetManagerProfileResponse")
	proto.RegisterType((*UpdateManagerProfileRequest)(nil), "UpdateManagerProfileRequest")
	proto.RegisterType((*UpdateManagerProfileResponse)(nil), "UpdateManagerProfileResponse")
	proto.RegisterType((*CreateManagerRestaurantAssignmentRequest)(nil), "CreateManagerRestaurantAssignmentRequest")
	proto.RegisterType((*CreateManagerRestaurantAssignmentResponse)(nil), "CreateManagerRestaurantAssignmentResponse")
	proto.RegisterType((*RemoveManagerRestaurantAssignmentRequest)(nil), "RemoveManagerRestaurantAssignmentRequest")
	proto.RegisterType((*RemoveManagerRestaurantAssignmentResponse)(nil), "RemoveManagerRestaurantAssignmentResponse")
	proto.RegisterType((*GetManagerRestaurantRequest)(nil), "GetManagerRestaurantRequest")
	proto.RegisterType((*GetManagerRestaurantResponse)(nil), "GetManagerRestaurantResponse")
}

func init() {
	proto.RegisterFile("manager.proto", fileDescriptor_cde9ec64f0d2c859)
}

var fileDescriptor_cde9ec64f0d2c859 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0xa4, 0x6d, 0x86, 0xb6, 0x20, 0x13, 0x81, 0x21, 0x29, 0xa0, 0x15, 0x87, 0x04,
	0xa4, 0x20, 0xc1, 0xa9, 0x07, 0x0e, 0x50, 0x44, 0x55, 0x89, 0x2f, 0x59, 0x42, 0x48, 0x5c, 0xac,
	0x6d, 0x3d, 0x71, 0x2d, 0xe2, 0xb5, 0xd9, 0x5d, 0x53, 0xf1, 0x27, 0xb8, 0xf0, 0x2b, 0xf8, 0x8b,
	0x9c, 0xaa, 0xf5, 0x6e, 0xec, 0xf8, 0xa3, 0x8d, 0x7b, 0xcb, 0xbe, 0xd9, 0x99, 0x7d, 0xef, 0x79,
	0x66, 0x02, 0x7b, 0x31, 0x65, 0x34, 0x44, 0x3e, 0x4f, 0x79, 0x22, 0x13, 0xf2, 0xdf, 0x82, 0xfd,
	0x8f, 0x1a, 0xf9, 0xc2, 0x93, 0x45, 0xb4, 0x44, 0xe7, 0x11, 0x80, 0xb9, 0xe3, 0x47, 0x81, 0x6b,
	0x3d, 0xb1, 0xa6, 0x43, 0x6f, 0x0d, 0x71, 0x26, 0x30, 0x44, 0x26, 0x23, 0xf9, 0x5b, 0x85, 0xed,
	0x3c, 0x5c, 0x02, 0x8e, 0x03, 0x7d, 0x46, 0x63, 0x74, 0x7b, 0x79, 0x20, 0xff, 0xad, 0x32, 0xa2,
	0x98, 0x86, 0xe8, 0x67, 0x7c, 0xe9, 0xf6, 0x75, 0x46, 0x01, 0x38, 0x04, 0x76, 0xd3, 0xf3, 0x84,
	0xa1, 0xcf, 0xb2, 0xf8, 0x14, 0xb9, 0x3b, 0xc8, 0x2f, 0x54, 0x30, 0x67, 0x04, 0x03, 0x8c, 0x69,
	0xb4, 0x74, 0xb7, 0xf2, 0xa0, 0x3e, 0x28, 0xa6, 0x67, 0x1c, 0xa9, 0xc4, 0xc0, 0xa7, 0xd2, 0xdd,
	0xd6, 0x4c, 0x4b, 0x44, 0xc5, 0xb3, 0x34, 0x58, 0xc5, 0x77, 0x74, 0xbc, 0x44, 0xc8, 0x5f, 0x0b,
	0xc6, 0x46, 0xbc, 0x87, 0x42, 0xd2, 0x8c, 0x53, 0x26, 0xdf, 0x08, 0x11, 0x85, 0x2c, 0x46, 0x26,
	0x9d, 0x7d, 0xb0, 0x8d, 0x03, 0x3d, 0xcf, 0x8e, 0x82, 0x9a, 0x33, 0x76, 0xc3, 0x99, 0xa7, 0xb0,
	0xc7, 0x8b, 0x3a, 0xea, 0x8a, 0x36, 0xa1, 0x0a, 0xd6, 0x58, 0xf7, 0xeb, 0xac, 0xc9, 0x1f, 0x0b,
	0xee, 0x79, 0x18, 0x46, 0x42, 0x22, 0x2f, 0xd8, 0xfd, 0xcc, 0x50, 0xc8, 0xaa, 0xf5, 0x56, 0xdd,
	0xfa, 0xc6, 0xf3, 0x76, 0xdb, 0xf3, 0x85, 0x95, 0xbd, 0x75, 0x2b, 0x1f, 0xc2, 0x4e, 0x4a, 0x85,
	0xb8, 0x48, 0x78, 0x60, 0x28, 0x15, 0x67, 0x72, 0x08, 0xf7, 0x1b, 0x7c, 0x44, 0x9a, 0x30, 0xb1,
	0xb1, 0x57, 0x94, 0x96, 0xc9, 0xd1, 0x39, 0x65, 0x21, 0xae, 0x9a, 0xcc, 0x14, 0x5d, 0x29, 0xda,
	0xd4, 0x6c, 0xcf, 0xe0, 0xce, 0x59, 0xc6, 0x39, 0x32, 0xe9, 0x17, 0xfc, 0xb4, 0xac, 0x06, 0xae,
	0x1a, 0x89, 0xe1, 0x45, 0x79, 0x4f, 0x0b, 0xac, 0x60, 0xe4, 0x31, 0x1c, 0x5c, 0xc1, 0x47, 0x2b,
	0x22, 0xc7, 0x70, 0xf7, 0x43, 0x12, 0x46, 0xac, 0xe6, 0x7c, 0xe1, 0x9a, 0x75, 0x95, 0x6b, 0x76,
	0xcd, 0xb5, 0x6f, 0x30, 0xaa, 0x16, 0x32, 0x96, 0xcd, 0x60, 0x3b, 0xd5, 0x93, 0x96, 0xd7, 0xba,
	0xf5, 0xf2, 0xf6, 0xbc, 0x3a, 0x80, 0xde, 0x2a, 0xae, 0x1e, 0x95, 0xc9, 0x0f, 0x64, 0xa6, 0xb6,
	0x3e, 0x90, 0x43, 0x70, 0x8f, 0x51, 0xd6, 0x72, 0x0c, 0xcd, 0x83, 0x16, 0x3b, 0x87, 0x06, 0x39,
	0x09, 0xc8, 0x7b, 0x78, 0xd0, 0x92, 0x7a, 0x63, 0x62, 0xe4, 0x9f, 0x05, 0xe3, 0xaf, 0xf9, 0x1c,
	0xb5, 0xd3, 0xd8, 0xf4, 0x55, 0x57, 0x4b, 0xc2, 0x5e, 0x5b, 0x12, 0xf5, 0x35, 0xd0, 0xbb, 0x6e,
	0x0d, 0xf4, 0xd7, 0xbf, 0x42, 0x65, 0xbd, 0x0c, 0x6a, 0xeb, 0x85, 0x9c, 0xc0, 0xa4, 0x9d, 0xea,
	0xcd, 0x65, 0xa7, 0x30, 0x3d, 0xca, 0xe7, 0xf4, 0x9a, 0xa5, 0xd1, 0xd5, 0x82, 0x4e, 0xc3, 0x4a,
	0x3e, 0xc3, 0xac, 0xc3, 0x8b, 0x46, 0x09, 0x81, 0x5d, 0xaa, 0x50, 0x05, 0xfa, 0xc5, 0xe2, 0xaa,
	0x60, 0xe4, 0x13, 0x4c, 0x3d, 0x8c, 0x93, 0x5f, 0x5d, 0x24, 0x74, 0xa9, 0xf7, 0x1c, 0x66, 0x1d,
	0xea, 0x99, 0xd9, 0x7a, 0x0d, 0xe3, 0xb2, 0xfd, 0xca, 0x9b, 0x1d, 0x2d, 0x23, 0xef, 0x60, 0xd2,
	0x9e, 0x6e, 0xf4, 0x37, 0x2c, 0xb5, 0x5a, 0x2c, 0x7d, 0x3b, 0xfa, 0xee, 0xcc, 0x5f, 0xac, 0x41,
	0x42, 0xd2, 0xc5, 0xe2, 0x74, 0x2b, 0xff, 0x3b, 0x7c, 0x75, 0x19, 0x00, 0x00, 0xff, 0xff, 0xd4,
	0x87, 0x58, 0x70, 0x1f, 0x07, 0x00, 0x00,
}
