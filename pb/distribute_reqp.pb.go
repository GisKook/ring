// Code generated by protoc-gen-go.
// source: distribute_reqp.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeReqp struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// 序列号
	Serail string `protobuf:"bytes,2,opt,name=serail" json:"serail,omitempty"`
	// 1 终端被动上报 3 超圈检测
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *DistributeReqp) Reset()                    { *m = DistributeReqp{} }
func (m *DistributeReqp) String() string            { return proto.CompactTextString(m) }
func (*DistributeReqp) ProtoMessage()               {}
func (*DistributeReqp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *DistributeReqp) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeReqp) GetSerail() string {
	if m != nil {
		return m.Serail
	}
	return ""
}

func (m *DistributeReqp) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributeReqp)(nil), "Carrier.DistributeReqp")
}

func init() { proto.RegisterFile("distribute_reqp.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x8d, 0x2f, 0x4a, 0x2d, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x52, 0xb2, 0xe1, 0xe2, 0x73, 0x81,
	0xab, 0x08, 0x4a, 0x2d, 0x2c, 0x10, 0xe2, 0xe1, 0x62, 0xc9, 0xcc, 0x4d, 0xcd, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62, 0x2b, 0x4e, 0x2d, 0x4a, 0xcc, 0xcc, 0x91, 0x60, 0x02,
	0xf3, 0x79, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x41, 0xbc, 0x24, 0x36, 0xb0, 0x69,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x65, 0x26, 0x84, 0x66, 0x00, 0x00, 0x00,
}
