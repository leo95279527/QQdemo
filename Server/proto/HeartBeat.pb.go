// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HeartBeat.proto

package proto

import proto1 "QQdemo/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HeartBeat_ToS struct {
	Timestamp        *int64 `protobuf:"varint,1,req,name=Timestamp" json:"Timestamp,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartBeat_ToS) Reset()                    { *m = HeartBeat_ToS{} }
func (m *HeartBeat_ToS) String() string            { return proto1.CompactTextString(m) }
func (*HeartBeat_ToS) ProtoMessage()               {}
func (*HeartBeat_ToS) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *HeartBeat_ToS) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func init() {
	proto1.RegisterType((*HeartBeat_ToS)(nil), "proto.HeartBeat_ToS")
}

func init() { proto1.RegisterFile("HeartBeat.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 76 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf7, 0x48, 0x4d, 0x2c,
	0x2a, 0x71, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xba, 0x5c, 0xbc, 0x70, 0x99, 0xf8, 0x90, 0xfc, 0x60, 0x21, 0x19, 0x2e, 0xce, 0x90, 0xcc, 0xdc,
	0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x09, 0x46, 0x05, 0x26, 0x0d, 0xe6, 0x20, 0x84, 0x00, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0x8f, 0x54, 0x29, 0x47, 0x00, 0x00, 0x00,
}
