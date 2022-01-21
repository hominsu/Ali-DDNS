// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/domain_record/v1/domain_record.proto

package v1

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

type DRRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
}

func (x *DRRequest) Reset() {
	*x = DRRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DRRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DRRequest) ProtoMessage() {}

func (x *DRRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DRRequest.ProtoReflect.Descriptor instead.
func (*DRRequest) Descriptor() ([]byte, []int) {
	return file_api_domain_record_v1_domain_record_proto_rawDescGZIP(), []int{0}
}

func (x *DRRequest) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

type DRResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainRecords string `protobuf:"bytes,1,opt,name=domain_records,json=domainRecords,proto3" json:"domain_records,omitempty"`
}

func (x *DRResponse) Reset() {
	*x = DRResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DRResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DRResponse) ProtoMessage() {}

func (x *DRResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DRResponse.ProtoReflect.Descriptor instead.
func (*DRResponse) Descriptor() ([]byte, []int) {
	return file_api_domain_record_v1_domain_record_proto_rawDescGZIP(), []int{1}
}

func (x *DRResponse) GetDomainRecords() string {
	if x != nil {
		return x.DomainRecords
	}
	return ""
}

type UpdateDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	RecordId   string `protobuf:"bytes,2,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	Rr         string `protobuf:"bytes,3,opt,name=rr,proto3" json:"rr,omitempty"`
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Value      string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateDomainRequest) Reset() {
	*x = UpdateDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainRequest) ProtoMessage() {}

func (x *UpdateDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainRequest.ProtoReflect.Descriptor instead.
func (*UpdateDomainRequest) Descriptor() ([]byte, []int) {
	return file_api_domain_record_v1_domain_record_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDomainRequest) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *UpdateDomainRequest) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *UpdateDomainRequest) GetRr() string {
	if x != nil {
		return x.Rr
	}
	return ""
}

func (x *UpdateDomainRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateDomainRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	RecordId  string `protobuf:"bytes,2,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (x *UpdateDomainResponse) Reset() {
	*x = UpdateDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainResponse) ProtoMessage() {}

func (x *UpdateDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_record_v1_domain_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainResponse.ProtoReflect.Descriptor instead.
func (*UpdateDomainResponse) Descriptor() ([]byte, []int) {
	return file_api_domain_record_v1_domain_record_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDomainResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UpdateDomainResponse) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

var File_api_domain_record_v1_domain_record_proto protoreflect.FileDescriptor

var file_api_domain_record_v1_domain_record_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a, 0x09,
	0x44, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0a, 0x44, 0x52,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x8d, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x72, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x52, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x49, 0x64, 0x32, 0xc6, 0x01, 0x0a, 0x0d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x52, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x52, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20,
	0x41, 0x6c, 0x69, 0x2d, 0x44, 0x44, 0x4e, 0x53, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_domain_record_v1_domain_record_proto_rawDescOnce sync.Once
	file_api_domain_record_v1_domain_record_proto_rawDescData = file_api_domain_record_v1_domain_record_proto_rawDesc
)

func file_api_domain_record_v1_domain_record_proto_rawDescGZIP() []byte {
	file_api_domain_record_v1_domain_record_proto_rawDescOnce.Do(func() {
		file_api_domain_record_v1_domain_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_domain_record_v1_domain_record_proto_rawDescData)
	})
	return file_api_domain_record_v1_domain_record_proto_rawDescData
}

var file_api_domain_record_v1_domain_record_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_domain_record_v1_domain_record_proto_goTypes = []interface{}{
	(*DRRequest)(nil),            // 0: domain_record.v1.DRRequest
	(*DRResponse)(nil),           // 1: domain_record.v1.DRResponse
	(*UpdateDomainRequest)(nil),  // 2: domain_record.v1.UpdateDomainRequest
	(*UpdateDomainResponse)(nil), // 3: domain_record.v1.UpdateDomainResponse
}
var file_api_domain_record_v1_domain_record_proto_depIdxs = []int32{
	0, // 0: domain_record.v1.DomainService.GetDomainRecord:input_type -> domain_record.v1.DRRequest
	2, // 1: domain_record.v1.DomainService.UpdateDomainRecord:input_type -> domain_record.v1.UpdateDomainRequest
	1, // 2: domain_record.v1.DomainService.GetDomainRecord:output_type -> domain_record.v1.DRResponse
	3, // 3: domain_record.v1.DomainService.UpdateDomainRecord:output_type -> domain_record.v1.UpdateDomainResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_domain_record_v1_domain_record_proto_init() }
func file_api_domain_record_v1_domain_record_proto_init() {
	if File_api_domain_record_v1_domain_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_domain_record_v1_domain_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DRRequest); i {
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
		file_api_domain_record_v1_domain_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DRResponse); i {
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
		file_api_domain_record_v1_domain_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainRequest); i {
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
		file_api_domain_record_v1_domain_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainResponse); i {
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
			RawDescriptor: file_api_domain_record_v1_domain_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_domain_record_v1_domain_record_proto_goTypes,
		DependencyIndexes: file_api_domain_record_v1_domain_record_proto_depIdxs,
		MessageInfos:      file_api_domain_record_v1_domain_record_proto_msgTypes,
	}.Build()
	File_api_domain_record_v1_domain_record_proto = out.File
	file_api_domain_record_v1_domain_record_proto_rawDesc = nil
	file_api_domain_record_v1_domain_record_proto_goTypes = nil
	file_api_domain_record_v1_domain_record_proto_depIdxs = nil
}
