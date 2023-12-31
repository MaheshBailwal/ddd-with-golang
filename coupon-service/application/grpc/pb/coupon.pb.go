// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: coupon.proto

package pb

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

type CouponProfileDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	DeptNo      int32  `protobuf:"varint,2,opt,name=dept_no,json=deptNo,proto3" json:"dept_no,omitempty"`
	CouponPluno *int32 `protobuf:"varint,3,opt,name=coupon_pluno,json=couponPluno,proto3,oneof" json:"coupon_pluno,omitempty"`
	CouponUpc   int32  `protobuf:"varint,4,opt,name=coupon_upc,json=couponUpc,proto3" json:"coupon_upc,omitempty"`
	Value       int32  `protobuf:"varint,5,opt,name=Value,proto3" json:"Value,omitempty"`
	ModifyUser  int32  `protobuf:"varint,6,opt,name=modify_user,json=modifyUser,proto3" json:"modify_user,omitempty"`
	ModifyDate  int32  `protobuf:"varint,7,opt,name=modify_date,json=modifyDate,proto3" json:"modify_date,omitempty"`
}

func (x *CouponProfileDetail) Reset() {
	*x = CouponProfileDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponProfileDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponProfileDetail) ProtoMessage() {}

func (x *CouponProfileDetail) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponProfileDetail.ProtoReflect.Descriptor instead.
func (*CouponProfileDetail) Descriptor() ([]byte, []int) {
	return file_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *CouponProfileDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CouponProfileDetail) GetDeptNo() int32 {
	if x != nil {
		return x.DeptNo
	}
	return 0
}

func (x *CouponProfileDetail) GetCouponPluno() int32 {
	if x != nil && x.CouponPluno != nil {
		return *x.CouponPluno
	}
	return 0
}

func (x *CouponProfileDetail) GetCouponUpc() int32 {
	if x != nil {
		return x.CouponUpc
	}
	return 0
}

func (x *CouponProfileDetail) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CouponProfileDetail) GetModifyUser() int32 {
	if x != nil {
		return x.ModifyUser
	}
	return 0
}

func (x *CouponProfileDetail) GetModifyDate() int32 {
	if x != nil {
		return x.ModifyDate
	}
	return 0
}

type CouponProfileDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponProfile []*CouponProfileDetail `protobuf:"bytes,1,rep,name=CouponProfile,proto3" json:"CouponProfile,omitempty"`
}

func (x *CouponProfileDetails) Reset() {
	*x = CouponProfileDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponProfileDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponProfileDetails) ProtoMessage() {}

func (x *CouponProfileDetails) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponProfileDetails.ProtoReflect.Descriptor instead.
func (*CouponProfileDetails) Descriptor() ([]byte, []int) {
	return file_coupon_proto_rawDescGZIP(), []int{1}
}

func (x *CouponProfileDetails) GetCouponProfile() []*CouponProfileDetail {
	if x != nil {
		return x.CouponProfile
	}
	return nil
}

type EmptyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyParams) Reset() {
	*x = EmptyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyParams) ProtoMessage() {}

func (x *EmptyParams) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyParams.ProtoReflect.Descriptor instead.
func (*EmptyParams) Descriptor() ([]byte, []int) {
	return file_coupon_proto_rawDescGZIP(), []int{2}
}

var File_coupon_proto protoreflect.FileDescriptor

var file_coupon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xee, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65,
	0x70, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65, 0x70,
	0x74, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x70, 0x6c,
	0x75, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x50, 0x6c, 0x75, 0x6e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x70, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x61,
	0x74, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x70, 0x6c,
	0x75, 0x6e, 0x6f, 0x22, 0x55, 0x0a, 0x14, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0x46, 0x0a, 0x0d, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x69, 0x6e, 0x76, 0x61, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coupon_proto_rawDescOnce sync.Once
	file_coupon_proto_rawDescData = file_coupon_proto_rawDesc
)

func file_coupon_proto_rawDescGZIP() []byte {
	file_coupon_proto_rawDescOnce.Do(func() {
		file_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_coupon_proto_rawDescData)
	})
	return file_coupon_proto_rawDescData
}

var file_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coupon_proto_goTypes = []interface{}{
	(*CouponProfileDetail)(nil),  // 0: pb.CouponProfileDetail
	(*CouponProfileDetails)(nil), // 1: pb.CouponProfileDetails
	(*EmptyParams)(nil),          // 2: pb.EmptyParams
}
var file_coupon_proto_depIdxs = []int32{
	0, // 0: pb.CouponProfileDetails.CouponProfile:type_name -> pb.CouponProfileDetail
	2, // 1: pb.CouponService.GetAll:input_type -> pb.EmptyParams
	1, // 2: pb.CouponService.GetAll:output_type -> pb.CouponProfileDetails
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coupon_proto_init() }
func file_coupon_proto_init() {
	if File_coupon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coupon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponProfileDetail); i {
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
		file_coupon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponProfileDetails); i {
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
		file_coupon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyParams); i {
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
	file_coupon_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coupon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coupon_proto_goTypes,
		DependencyIndexes: file_coupon_proto_depIdxs,
		MessageInfos:      file_coupon_proto_msgTypes,
	}.Build()
	File_coupon_proto = out.File
	file_coupon_proto_rawDesc = nil
	file_coupon_proto_goTypes = nil
	file_coupon_proto_depIdxs = nil
}
