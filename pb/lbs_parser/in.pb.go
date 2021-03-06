// Code generated by protoc-gen-go.
// source: in.proto
// DO NOT EDIT!

package Lbs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Report_LbsType int32

const (
	Report_UNKNOWN Report_LbsType = 0
	Report_WIFI    Report_LbsType = 1
	Report_STATION Report_LbsType = 2
)

var Report_LbsType_name = map[int32]string{
	0: "UNKNOWN",
	1: "WIFI",
	2: "STATION",
}
var Report_LbsType_value = map[string]int32{
	"UNKNOWN": 0,
	"WIFI":    1,
	"STATION": 2,
}

func (x Report_LbsType) String() string {
	return proto.EnumName(Report_LbsType_name, int32(x))
}
func (Report_LbsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type WifiCell struct {
	Mac     string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	Singnal string `protobuf:"bytes,2,opt,name=singnal" json:"singnal,omitempty"`
}

func (m *WifiCell) Reset()                    { *m = WifiCell{} }
func (m *WifiCell) String() string            { return proto.CompactTextString(m) }
func (*WifiCell) ProtoMessage()               {}
func (*WifiCell) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *WifiCell) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *WifiCell) GetSingnal() string {
	if m != nil {
		return m.Singnal
	}
	return ""
}

type StationCell struct {
	Lac string `protobuf:"bytes,1,opt,name=lac" json:"lac,omitempty"`
	Cid string `protobuf:"bytes,2,opt,name=cid" json:"cid,omitempty"`
	Dbm string `protobuf:"bytes,3,opt,name=dbm" json:"dbm,omitempty"`
}

func (m *StationCell) Reset()                    { *m = StationCell{} }
func (m *StationCell) String() string            { return proto.CompactTextString(m) }
func (*StationCell) ProtoMessage()               {}
func (*StationCell) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *StationCell) GetLac() string {
	if m != nil {
		return m.Lac
	}
	return ""
}

func (m *StationCell) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *StationCell) GetDbm() string {
	if m != nil {
		return m.Dbm
	}
	return ""
}

type Report struct {
	Header *Header        `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Type   Report_LbsType `protobuf:"varint,2,opt,name=type,enum=Lbs.Report_LbsType" json:"type,omitempty"`
	// Wifi位置
	WifiCell []*WifiCell `protobuf:"bytes,3,rep,name=wifi_cell" json:"wifi_cell,omitempty"`
	// 基站位置
	StationCell []*StationCell `protobuf:"bytes,4,rep,name=station_cell" json:"station_cell,omitempty"`
	// 附加数据  LocationExtra 序列化后的结果
	Extra []byte `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Report) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Report) GetType() Report_LbsType {
	if m != nil {
		return m.Type
	}
	return Report_UNKNOWN
}

func (m *Report) GetWifiCell() []*WifiCell {
	if m != nil {
		return m.WifiCell
	}
	return nil
}

func (m *Report) GetStationCell() []*StationCell {
	if m != nil {
		return m.StationCell
	}
	return nil
}

func (m *Report) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterType((*WifiCell)(nil), "Lbs.WifiCell")
	proto.RegisterType((*StationCell)(nil), "Lbs.StationCell")
	proto.RegisterType((*Report)(nil), "Lbs.Report")
	proto.RegisterEnum("Lbs.Report_LbsType", Report_LbsType_name, Report_LbsType_value)
}

func init() { proto.RegisterFile("in.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4f, 0x83, 0x30,
	0x18, 0x86, 0x65, 0x65, 0xc0, 0xbe, 0x6e, 0x4a, 0xea, 0x85, 0xe8, 0x05, 0x39, 0x18, 0x2e, 0x72,
	0x40, 0xff, 0x80, 0x31, 0x31, 0x12, 0x17, 0x96, 0x38, 0x0c, 0x47, 0xd3, 0x42, 0xa7, 0x4d, 0x3a,
	0x20, 0xb4, 0x89, 0xee, 0x5f, 0xfa, 0x93, 0x0c, 0x6d, 0xb6, 0x78, 0x6b, 0x9f, 0x3e, 0x6f, 0xfb,
	0xf6, 0x83, 0x40, 0x74, 0xd9, 0x30, 0xf6, 0xba, 0x27, 0x68, 0xcd, 0xd4, 0x15, 0x30, 0xaa, 0xb8,
	0x05, 0x49, 0x0a, 0x41, 0x2d, 0x76, 0xe2, 0x89, 0x4b, 0x49, 0x30, 0xa0, 0x3d, 0x6d, 0x22, 0x27,
	0x76, 0xd2, 0x05, 0xb9, 0x00, 0x5f, 0x89, 0xee, 0xb3, 0xa3, 0x32, 0x9a, 0x4d, 0x20, 0x79, 0x00,
	0xbc, 0xd5, 0x54, 0x8b, 0xbe, 0x3b, 0xca, 0xf2, 0x24, 0x63, 0x40, 0x8d, 0x68, 0xad, 0x38, 0x6d,
	0x5a, 0xb6, 0x8f, 0x90, 0x49, 0xfd, 0x3a, 0xe0, 0xbd, 0xf1, 0xa1, 0x1f, 0x35, 0xb9, 0x06, 0xef,
	0x8b, 0xd3, 0x96, 0x8f, 0x26, 0x84, 0x73, 0x9c, 0xad, 0x99, 0xca, 0x5e, 0x0c, 0x22, 0x37, 0xe0,
	0xea, 0xc3, 0xc0, 0xcd, 0x15, 0xe7, 0xf9, 0xa5, 0x39, 0xb2, 0xb9, 0x69, 0x59, 0x1d, 0x06, 0x4e,
	0x62, 0x58, 0x7c, 0x8b, 0x9d, 0xf8, 0x68, 0xb8, 0x94, 0x11, 0x8a, 0x51, 0x8a, 0xf3, 0x95, 0xf1,
	0x4e, 0x1f, 0xb8, 0x85, 0xa5, 0xb2, 0x15, 0xad, 0xe4, 0x1a, 0x29, 0x34, 0xd2, 0xff, 0xee, 0x2b,
	0x98, 0xf3, 0x1f, 0x3d, 0xd2, 0x68, 0x1e, 0x3b, 0xe9, 0x32, 0xb9, 0x03, 0xff, 0xf8, 0x06, 0x06,
	0xff, 0xbd, 0x7c, 0x2d, 0x37, 0x75, 0x19, 0x9e, 0x91, 0x00, 0xdc, 0xba, 0x78, 0x2e, 0x42, 0x67,
	0xc2, 0xdb, 0xea, 0xb1, 0x2a, 0x36, 0x65, 0x38, 0x63, 0x9e, 0x99, 0xdc, 0xfd, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x57, 0xd9, 0xc6, 0x01, 0x56, 0x01, 0x00, 0x00,
}
