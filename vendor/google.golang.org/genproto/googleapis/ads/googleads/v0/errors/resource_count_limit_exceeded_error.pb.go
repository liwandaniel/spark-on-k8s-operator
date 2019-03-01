// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/resource_count_limit_exceeded_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible resource count limit exceeded errors.
type ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError int32

const (
	// Enum unspecified.
	ResourceCountLimitExceededErrorEnum_UNSPECIFIED ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 0
	// The received error code is not known in this version.
	ResourceCountLimitExceededErrorEnum_UNKNOWN ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 1
	// Indicates that this request would exceed the number of allowed resources
	// for the Google Ads account. The exact resource type and limit being
	// checked
	// can be inferred from accountLimitType.
	ResourceCountLimitExceededErrorEnum_ACCOUNT_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 2
	// Indicates that this request would exceed the number of allowed resources
	// in a Campaign. The exact resource type and limit being checked can be
	// inferred from accountLimitType, and the numeric id of the
	// Campaign involved is given by enclosingId.
	ResourceCountLimitExceededErrorEnum_CAMPAIGN_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 3
	// Indicates that this request would exceed the number of allowed resources
	// in an ad group. The exact resource type and limit being checked can be
	// inferred from accountLimitType, and the numeric id of the
	// ad group involved is given by enclosingId.
	ResourceCountLimitExceededErrorEnum_ADGROUP_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 4
	// Indicates that this request would exceed the number of allowed resources
	// in an ad group ad. The exact resource type and limit being checked can
	// be inferred from accountLimitType, and the enclosingId
	// contains the ad group id followed by the ad id, separated by a single
	// comma (,).
	ResourceCountLimitExceededErrorEnum_AD_GROUP_AD_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 5
	// Indicates that this request would exceed the number of allowed resources
	// in an ad group criterion. The exact resource type and limit being checked
	// can be inferred from accountLimitType, and the
	// enclosingId contains the ad group id followed by the
	// criterion id, separated by a single comma (,).
	ResourceCountLimitExceededErrorEnum_AD_GROUP_CRITERION_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 6
	// Indicates that this request would exceed the number of allowed resources
	// in this shared set. The exact resource type and limit being checked can
	// be inferred from accountLimitType, and the numeric id of the
	// shared set involved is given by enclosingId.
	ResourceCountLimitExceededErrorEnum_SHARED_SET_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 7
	// Exceeds a limit related to a matching function.
	ResourceCountLimitExceededErrorEnum_MATCHING_FUNCTION_LIMIT ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 8
	// The response for this request would exceed the maximum number of rows
	// that can be returned.
	ResourceCountLimitExceededErrorEnum_RESPONSE_ROW_LIMIT_EXCEEDED ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 9
)

var ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ACCOUNT_LIMIT",
	3: "CAMPAIGN_LIMIT",
	4: "ADGROUP_LIMIT",
	5: "AD_GROUP_AD_LIMIT",
	6: "AD_GROUP_CRITERION_LIMIT",
	7: "SHARED_SET_LIMIT",
	8: "MATCHING_FUNCTION_LIMIT",
	9: "RESPONSE_ROW_LIMIT_EXCEEDED",
}
var ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"ACCOUNT_LIMIT":               2,
	"CAMPAIGN_LIMIT":              3,
	"ADGROUP_LIMIT":               4,
	"AD_GROUP_AD_LIMIT":           5,
	"AD_GROUP_CRITERION_LIMIT":    6,
	"SHARED_SET_LIMIT":            7,
	"MATCHING_FUNCTION_LIMIT":     8,
	"RESPONSE_ROW_LIMIT_EXCEEDED": 9,
}

func (x ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) String() string {
	return proto.EnumName(ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_name, int32(x))
}
func (ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_resource_count_limit_exceeded_error_78e157e8f7acfe78, []int{0, 0}
}

// Container for enum describing possible resource count limit exceeded errors.
type ResourceCountLimitExceededErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceCountLimitExceededErrorEnum) Reset()         { *m = ResourceCountLimitExceededErrorEnum{} }
func (m *ResourceCountLimitExceededErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ResourceCountLimitExceededErrorEnum) ProtoMessage()    {}
func (*ResourceCountLimitExceededErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_count_limit_exceeded_error_78e157e8f7acfe78, []int{0}
}
func (m *ResourceCountLimitExceededErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceCountLimitExceededErrorEnum.Unmarshal(m, b)
}
func (m *ResourceCountLimitExceededErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceCountLimitExceededErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ResourceCountLimitExceededErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceCountLimitExceededErrorEnum.Merge(dst, src)
}
func (m *ResourceCountLimitExceededErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ResourceCountLimitExceededErrorEnum.Size(m)
}
func (m *ResourceCountLimitExceededErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceCountLimitExceededErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceCountLimitExceededErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ResourceCountLimitExceededErrorEnum)(nil), "google.ads.googleads.v0.errors.ResourceCountLimitExceededErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError", ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_name, ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/resource_count_limit_exceeded_error.proto", fileDescriptor_resource_count_limit_exceeded_error_78e157e8f7acfe78)
}

var fileDescriptor_resource_count_limit_exceeded_error_78e157e8f7acfe78 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x69, 0x06, 0x1b, 0x78, 0x02, 0x32, 0x0b, 0x04, 0xd2, 0xd0, 0x26, 0x15, 0xce, 0x4e,
	0x24, 0x6e, 0xe6, 0xe4, 0xd9, 0x5e, 0x6a, 0xb1, 0x3a, 0x51, 0xfe, 0x74, 0x08, 0x55, 0xb2, 0x4a,
	0x63, 0x45, 0x95, 0xda, 0x7a, 0x8a, 0xdb, 0x89, 0x33, 0x8f, 0xc2, 0x71, 0x8f, 0xc2, 0xa3, 0xf0,
	0x12, 0xa0, 0xc4, 0x49, 0x38, 0xb1, 0x9d, 0xf2, 0xd5, 0x57, 0x1f, 0x7f, 0x1c, 0xfd, 0x7e, 0x06,
	0x93, 0xca, 0x98, 0x6a, 0xad, 0x83, 0x45, 0x69, 0x03, 0x17, 0x9b, 0x74, 0x1b, 0x06, 0xba, 0xae,
	0x4d, 0x6d, 0x83, 0x5a, 0x5b, 0xb3, 0xaf, 0x97, 0x5a, 0x2d, 0xcd, 0x7e, 0xbb, 0x53, 0xeb, 0xd5,
	0x66, 0xb5, 0x53, 0xfa, 0xfb, 0x52, 0xeb, 0x52, 0x97, 0xaa, 0x85, 0xd0, 0x4d, 0x6d, 0x76, 0x06,
	0x9e, 0xb9, 0xe3, 0x68, 0x51, 0x5a, 0x34, 0x98, 0xd0, 0x6d, 0x88, 0x9c, 0x69, 0x7c, 0xe7, 0x81,
	0xf7, 0x69, 0x67, 0xa3, 0x8d, 0xec, 0xaa, 0x71, 0xf1, 0x4e, 0xc5, 0x1b, 0x88, 0x6f, 0xf7, 0x9b,
	0xf1, 0x0f, 0x0f, 0x9c, 0x3f, 0xc0, 0xc1, 0x97, 0xe0, 0xb8, 0x90, 0x59, 0xc2, 0xa9, 0xb8, 0x14,
	0x9c, 0xf9, 0x8f, 0xe0, 0x31, 0x38, 0x2a, 0xe4, 0x67, 0x19, 0x5f, 0x4b, 0x7f, 0x04, 0x4f, 0xc0,
	0x73, 0x42, 0x69, 0x5c, 0xc8, 0x5c, 0x5d, 0x89, 0xa9, 0xc8, 0x7d, 0x0f, 0x42, 0xf0, 0x82, 0x92,
	0x69, 0x42, 0x44, 0x24, 0xbb, 0xee, 0xa0, 0xc5, 0x58, 0x94, 0xc6, 0x45, 0xd2, 0x55, 0x8f, 0xe1,
	0x6b, 0x70, 0x42, 0x98, 0x72, 0x1d, 0x61, 0x5d, 0xfd, 0x04, 0xbe, 0x03, 0x6f, 0x87, 0x9a, 0xa6,
	0x22, 0xe7, 0xa9, 0x88, 0x7b, 0xcf, 0x21, 0x7c, 0x05, 0xfc, 0x6c, 0x42, 0x52, 0xce, 0x54, 0xc6,
	0xfb, 0x1b, 0x8f, 0xe0, 0x29, 0x78, 0x33, 0x25, 0x39, 0x9d, 0x08, 0x19, 0xa9, 0xcb, 0x42, 0xd2,
	0xfc, 0xdf, 0x91, 0xa7, 0xf0, 0x1c, 0x9c, 0xa6, 0x3c, 0x4b, 0x62, 0x99, 0x71, 0x95, 0xc6, 0xd7,
	0xae, 0x57, 0xfc, 0x0b, 0xe5, 0x9c, 0x71, 0xe6, 0x3f, 0xbb, 0xf8, 0x33, 0x02, 0xe3, 0xa5, 0xd9,
	0xa0, 0xfb, 0x67, 0x7a, 0xf1, 0xe1, 0x81, 0x41, 0x25, 0xcd, 0x66, 0x92, 0xd1, 0x57, 0xd6, 0x79,
	0x2a, 0xb3, 0x5e, 0x6c, 0x2b, 0x64, 0xea, 0x2a, 0xa8, 0xf4, 0xb6, 0xdd, 0x5b, 0xbf, 0xf5, 0x9b,
	0x95, 0xfd, 0xdf, 0x23, 0xf8, 0xe4, 0x3e, 0x3f, 0xbd, 0x83, 0x88, 0x90, 0x3b, 0xef, 0x2c, 0x72,
	0x32, 0x52, 0x5a, 0xe4, 0x62, 0x93, 0x66, 0x21, 0x6a, 0xaf, 0xb4, 0xbf, 0x7a, 0x60, 0x4e, 0x4a,
	0x3b, 0x1f, 0x80, 0xf9, 0x2c, 0x9c, 0x3b, 0xe0, 0xb7, 0x37, 0x76, 0x2d, 0xc6, 0xa4, 0xb4, 0x18,
	0x0f, 0x08, 0xc6, 0xb3, 0x10, 0x63, 0x07, 0x7d, 0x3b, 0x6c, 0xff, 0xee, 0xe3, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x38, 0xc5, 0xbd, 0x3c, 0xa1, 0x02, 0x00, 0x00,
}
