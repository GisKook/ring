// Code generated by protoc-gen-go.
// source: distribute_tearchk.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeTearchk struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// 序列号
	Serial string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
}

func (m *DistributeTearchk) Reset()                    { *m = DistributeTearchk{} }
func (m *DistributeTearchk) String() string            { return proto.CompactTextString(m) }
func (*DistributeTearchk) ProtoMessage()               {}
func (*DistributeTearchk) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *DistributeTearchk) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeTearchk) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributeTearchk)(nil), "Carrier.DistributeTearchk")
}

func init() { proto.RegisterFile("distribute_tearchk.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x8d, 0x2f, 0x49, 0x4d, 0x2c, 0x4a, 0xce, 0xc8, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x52, 0x32, 0xe4,
	0x12, 0x74, 0x81, 0x2b, 0x0a, 0x81, 0xa8, 0x11, 0xe2, 0xe1, 0x62, 0xc9, 0xcc, 0x4d, 0xcd, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62, 0x2b, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0x91,
	0x60, 0x02, 0xf1, 0x93, 0xd8, 0xc0, 0x46, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x94, 0xd1,
	0x1c, 0x5d, 0x5e, 0x00, 0x00, 0x00,
}
