// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: common/idl/domain/user/user_dto.proto

//protoc --go-grpc_out=. --go_out=. ./common/idl/domain/user/user_dto.proto

package user

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

type UserDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      uint64          `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Device      *DeviceDTO      `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	Setting     *SettingDTO     `protobuf:"bytes,3,opt,name=setting,proto3" json:"setting,omitempty"`
	Information *InformationDTO `protobuf:"bytes,4,opt,name=information,proto3" json:"information,omitempty"`
	Pprofile    *ProfileDTO     `protobuf:"bytes,5,opt,name=pprofile,proto3" json:"pprofile,omitempty"`
}

func (x *UserDTO) Reset() {
	*x = UserDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDTO) ProtoMessage() {}

func (x *UserDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDTO.ProtoReflect.Descriptor instead.
func (*UserDTO) Descriptor() ([]byte, []int) {
	return file_common_idl_domain_user_user_dto_proto_rawDescGZIP(), []int{0}
}

func (x *UserDTO) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserDTO) GetDevice() *DeviceDTO {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *UserDTO) GetSetting() *SettingDTO {
	if x != nil {
		return x.Setting
	}
	return nil
}

func (x *UserDTO) GetInformation() *InformationDTO {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *UserDTO) GetPprofile() *ProfileDTO {
	if x != nil {
		return x.Pprofile
	}
	return nil
}

type DeviceDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID   uint64 `protobuf:"varint,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Os         string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	AppVersion string `protobuf:"bytes,3,opt,name=appVersion,proto3" json:"appVersion,omitempty"`
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Model      string `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *DeviceDTO) Reset() {
	*x = DeviceDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDTO) ProtoMessage() {}

func (x *DeviceDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDTO.ProtoReflect.Descriptor instead.
func (*DeviceDTO) Descriptor() ([]byte, []int) {
	return file_common_idl_domain_user_user_dto_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceDTO) GetDeviceID() uint64 {
	if x != nil {
		return x.DeviceID
	}
	return 0
}

func (x *DeviceDTO) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *DeviceDTO) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *DeviceDTO) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeviceDTO) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type SettingDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FontSize            string `protobuf:"bytes,1,opt,name=fontSize,proto3" json:"fontSize,omitempty"`
	DarkMode            bool   `protobuf:"varint,2,opt,name=darkMode,proto3" json:"darkMode,omitempty"`
	ReceiveNotification bool   `protobuf:"varint,3,opt,name=receiveNotification,proto3" json:"receiveNotification,omitempty"`
	Language            string `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Notifications       bool   `protobuf:"varint,5,opt,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *SettingDTO) Reset() {
	*x = SettingDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingDTO) ProtoMessage() {}

func (x *SettingDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingDTO.ProtoReflect.Descriptor instead.
func (*SettingDTO) Descriptor() ([]byte, []int) {
	return file_common_idl_domain_user_user_dto_proto_rawDescGZIP(), []int{2}
}

func (x *SettingDTO) GetFontSize() string {
	if x != nil {
		return x.FontSize
	}
	return ""
}

func (x *SettingDTO) GetDarkMode() bool {
	if x != nil {
		return x.DarkMode
	}
	return false
}

func (x *SettingDTO) GetReceiveNotification() bool {
	if x != nil {
		return x.ReceiveNotification
	}
	return false
}

func (x *SettingDTO) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *SettingDTO) GetNotifications() bool {
	if x != nil {
		return x.Notifications
	}
	return false
}

// 个人资料信息，在个人信息页面展示
type InformationDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname  string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar    string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *InformationDTO) Reset() {
	*x = InformationDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformationDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformationDTO) ProtoMessage() {}

func (x *InformationDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformationDTO.ProtoReflect.Descriptor instead.
func (*InformationDTO) Descriptor() ([]byte, []int) {
	return file_common_idl_domain_user_user_dto_proto_rawDescGZIP(), []int{3}
}

func (x *InformationDTO) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *InformationDTO) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *InformationDTO) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// 异步任务构建用户画像
type ProfileDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Age      int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Gender   string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Tags     string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ProfileDTO) Reset() {
	*x = ProfileDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileDTO) ProtoMessage() {}

func (x *ProfileDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_idl_domain_user_user_dto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileDTO.ProtoReflect.Descriptor instead.
func (*ProfileDTO) Descriptor() ([]byte, []int) {
	return file_common_idl_domain_user_user_dto_proto_rawDescGZIP(), []int{4}
}

func (x *ProfileDTO) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ProfileDTO) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *ProfileDTO) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *ProfileDTO) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

var File_common_idl_domain_user_user_dto_proto protoreflect.FileDescriptor

var file_common_idl_domain_user_user_dto_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xdc, 0x01,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x54, 0x4f, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x54,
	0x4f, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c,
	0x0a, 0x08, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44,
	0x54, 0x4f, 0x52, 0x08, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x81, 0x01, 0x0a,
	0x09, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x54, 0x4f, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x6f, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x72, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x61, 0x72, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x62, 0x0a, 0x0e, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x66, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_idl_domain_user_user_dto_proto_rawDescOnce sync.Once
	file_common_idl_domain_user_user_dto_proto_rawDescData = file_common_idl_domain_user_user_dto_proto_rawDesc
)

func file_common_idl_domain_user_user_dto_proto_rawDescGZIP() []byte {
	file_common_idl_domain_user_user_dto_proto_rawDescOnce.Do(func() {
		file_common_idl_domain_user_user_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_idl_domain_user_user_dto_proto_rawDescData)
	})
	return file_common_idl_domain_user_user_dto_proto_rawDescData
}

var file_common_idl_domain_user_user_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_idl_domain_user_user_dto_proto_goTypes = []interface{}{
	(*UserDTO)(nil),        // 0: user.UserDTO
	(*DeviceDTO)(nil),      // 1: user.DeviceDTO
	(*SettingDTO)(nil),     // 2: user.SettingDTO
	(*InformationDTO)(nil), // 3: user.InformationDTO
	(*ProfileDTO)(nil),     // 4: user.ProfileDTO
}
var file_common_idl_domain_user_user_dto_proto_depIdxs = []int32{
	1, // 0: user.UserDTO.device:type_name -> user.DeviceDTO
	2, // 1: user.UserDTO.setting:type_name -> user.SettingDTO
	3, // 2: user.UserDTO.information:type_name -> user.InformationDTO
	4, // 3: user.UserDTO.pprofile:type_name -> user.ProfileDTO
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_idl_domain_user_user_dto_proto_init() }
func file_common_idl_domain_user_user_dto_proto_init() {
	if File_common_idl_domain_user_user_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_idl_domain_user_user_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDTO); i {
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
		file_common_idl_domain_user_user_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDTO); i {
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
		file_common_idl_domain_user_user_dto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingDTO); i {
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
		file_common_idl_domain_user_user_dto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformationDTO); i {
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
		file_common_idl_domain_user_user_dto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileDTO); i {
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
			RawDescriptor: file_common_idl_domain_user_user_dto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_idl_domain_user_user_dto_proto_goTypes,
		DependencyIndexes: file_common_idl_domain_user_user_dto_proto_depIdxs,
		MessageInfos:      file_common_idl_domain_user_user_dto_proto_msgTypes,
	}.Build()
	File_common_idl_domain_user_user_dto_proto = out.File
	file_common_idl_domain_user_user_dto_proto_rawDesc = nil
	file_common_idl_domain_user_user_dto_proto_goTypes = nil
	file_common_idl_domain_user_user_dto_proto_depIdxs = nil
}
