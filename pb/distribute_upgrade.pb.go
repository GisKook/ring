// Code generated by protoc-gen-go.
// source: distribute_upgrade.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeUpgrade struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// 序列号
	Serial string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	Ip     string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Port   string `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	User   string `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Passwd string `protobuf:"bytes,6,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *DistributeUpgrade) Reset()                    { *m = DistributeUpgrade{} }
func (m *DistributeUpgrade) String() string            { return proto.CompactTextString(m) }
func (*DistributeUpgrade) ProtoMessage()               {}
func (*DistributeUpgrade) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *DistributeUpgrade) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeUpgrade) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *DistributeUpgrade) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DistributeUpgrade) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *DistributeUpgrade) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *DistributeUpgrade) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributeUpgrade)(nil), "Carrier.DistributeUpgrade")
}

func init() { proto.RegisterFile("distribute_upgrade.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0xcd, 0xb1, 0x0a, 0x02, 0x31,
	0x0c, 0xc6, 0x71, 0xee, 0xac, 0x15, 0xc3, 0x21, 0xd8, 0x29, 0xa3, 0x38, 0x39, 0xb9, 0xf8, 0x08,
	0xfa, 0x0a, 0xce, 0xd2, 0xa3, 0x41, 0x02, 0x4a, 0x43, 0xd2, 0xe2, 0xeb, 0x8b, 0x95, 0xde, 0x96,
	0x1f, 0xfc, 0xc3, 0x07, 0x98, 0xd8, 0x8a, 0xf2, 0x5c, 0x0b, 0x3d, 0xaa, 0x3c, 0x35, 0x26, 0x3a,
	0x8b, 0xe6, 0x92, 0xc3, 0xe6, 0x1a, 0x55, 0x99, 0xf4, 0xc8, 0xb0, 0xbf, 0x2d, 0xd1, 0xfd, 0xdf,
	0x84, 0x09, 0x1c, 0xbf, 0x89, 0x71, 0x38, 0x0c, 0xa7, 0x6d, 0xd8, 0x81, 0x37, 0x52, 0x8e, 0x2f,
	0x1c, 0x9b, 0x01, 0x46, 0x16, 0x5c, 0xb5, 0x7b, 0x02, 0x27, 0x59, 0x0b, 0xba, 0xae, 0x6a, 0xa4,
	0xb8, 0xee, 0x7f, 0x12, 0xcd, 0x3e, 0x09, 0xfd, 0xcf, 0xb3, 0x6f, 0xd3, 0x97, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0xdb, 0x5b, 0xa7, 0x96, 0x00, 0x00, 0x00,
}
