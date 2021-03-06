// Code generated by protoc-gen-go.
// source: distribute_srvset.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeSrvset struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// 序列号
	Serial string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	// 服务器地址 格式为 127.0.0.1:8000
	Srvset string `protobuf:"bytes,3,opt,name=srvset" json:"srvset,omitempty"`
}

func (m *DistributeSrvset) Reset()                    { *m = DistributeSrvset{} }
func (m *DistributeSrvset) String() string            { return proto.CompactTextString(m) }
func (*DistributeSrvset) ProtoMessage()               {}
func (*DistributeSrvset) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *DistributeSrvset) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeSrvset) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *DistributeSrvset) GetSrvset() string {
	if m != nil {
		return m.Srvset
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributeSrvset)(nil), "Carrier.DistributeSrvset")
}

func init() { proto.RegisterFile("distribute_srvset.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x8d, 0x2f, 0x2e, 0x2a, 0x2b, 0x4e, 0x2d, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x52, 0x72, 0xe0, 0x12,
	0x70, 0x81, 0xab, 0x09, 0x06, 0x2b, 0x11, 0xe2, 0xe1, 0x62, 0xc9, 0xcc, 0x4d, 0xcd, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62, 0x2b, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0x91, 0x60,
	0x82, 0xf3, 0xc1, 0xea, 0x24, 0x98, 0x41, 0xfc, 0x24, 0x36, 0xb0, 0x89, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x35, 0xea, 0xfb, 0x59, 0x6c, 0x00, 0x00, 0x00,
}
