// Code generated by protoc-gen-go.
// source: report_location.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WifiCell struct {
	Mac     string `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	Singnal string `protobuf:"bytes,2,opt,name=singnal" json:"singnal,omitempty"`
}

func (m *WifiCell) Reset()                    { *m = WifiCell{} }
func (m *WifiCell) String() string            { return proto.CompactTextString(m) }
func (*WifiCell) ProtoMessage()               {}
func (*WifiCell) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

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
func (*StationCell) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

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

// 将数据上传给位置解析服务的结构
// 如果解析成功 则 将位置转发，使用message ReportLocation 如果失败则返回给请求方
type LocationInfo struct {
	// 附加数据  LocationExtra 序列化后的结果
	Extra []byte `protobuf:"bytes,1,opt,name=extra,proto3" json:"extra,omitempty"`
	// Wifi位置
	WifiCell []*WifiCell `protobuf:"bytes,2,rep,name=wifi_cell" json:"wifi_cell,omitempty"`
	// 基站位置
	StationCell []*StationCell `protobuf:"bytes,3,rep,name=station_cell" json:"station_cell,omitempty"`
}

func (m *LocationInfo) Reset()                    { *m = LocationInfo{} }
func (m *LocationInfo) String() string            { return proto.CompactTextString(m) }
func (*LocationInfo) ProtoMessage()               {}
func (*LocationInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *LocationInfo) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *LocationInfo) GetWifiCell() []*WifiCell {
	if m != nil {
		return m.WifiCell
	}
	return nil
}

func (m *LocationInfo) GetStationCell() []*StationCell {
	if m != nil {
		return m.StationCell
	}
	return nil
}

// 最终抛给上层服务的结构,考虑到解析位置服务的独立性，将其他的数据都作为Extra数据进行处理
type ReportLocation struct {
	// 附加数据  LocationExtra 序列化后的结果
	Extra     []byte `protobuf:"bytes,1,opt,name=extra,proto3" json:"extra,omitempty"`
	Longitude string `protobuf:"bytes,2,opt,name=longitude" json:"longitude,omitempty"`
	Latitude  string `protobuf:"bytes,3,opt,name=latitude" json:"latitude,omitempty"`
	Speed     string `protobuf:"bytes,4,opt,name=speed" json:"speed,omitempty"`
}

func (m *ReportLocation) Reset()                    { *m = ReportLocation{} }
func (m *ReportLocation) String() string            { return proto.CompactTextString(m) }
func (*ReportLocation) ProtoMessage()               {}
func (*ReportLocation) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ReportLocation) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *ReportLocation) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *ReportLocation) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *ReportLocation) GetSpeed() string {
	if m != nil {
		return m.Speed
	}
	return ""
}

func init() {
	proto.RegisterType((*WifiCell)(nil), "Carrier.WifiCell")
	proto.RegisterType((*StationCell)(nil), "Carrier.StationCell")
	proto.RegisterType((*LocationInfo)(nil), "Carrier.LocationInfo")
	proto.RegisterType((*ReportLocation)(nil), "Carrier.ReportLocation")
}

func init() { proto.RegisterFile("report_location.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x55, 0x0c, 0xb4, 0xb9, 0x84, 0x8f, 0x5a, 0x20, 0x65, 0x44, 0x11, 0x43, 0xc4, 0x90,
	0x01, 0xf8, 0x07, 0x9d, 0x90, 0x98, 0xe8, 0xc0, 0x58, 0xb9, 0xc9, 0xa5, 0x3a, 0xe9, 0xea, 0x8b,
	0x1c, 0x23, 0xf8, 0xf9, 0xc8, 0x76, 0x1a, 0x21, 0xc6, 0x3b, 0x3d, 0xaf, 0xdf, 0xc7, 0x07, 0xf7,
	0x0e, 0x07, 0x71, 0x7e, 0xc7, 0xd2, 0x1a, 0x4f, 0x62, 0x9b, 0xc1, 0x89, 0x17, 0xbd, 0xdc, 0x18,
	0xe7, 0x08, 0x5d, 0x55, 0xc3, 0xea, 0x93, 0x7a, 0xda, 0x20, 0xb3, 0xce, 0x41, 0x1d, 0x4d, 0x5b,
	0x2e, 0x1e, 0x16, 0x75, 0xa6, 0x6f, 0x60, 0x39, 0x92, 0x3d, 0x58, 0xc3, 0xe5, 0x59, 0x58, 0x54,
	0xaf, 0x90, 0x6f, 0x7d, 0x7c, 0xe3, 0x04, 0xf3, 0x0c, 0xe7, 0xa0, 0x5a, 0xea, 0x12, 0x18, 0x86,
	0x6e, 0x7f, 0x2c, 0x55, 0x4c, 0x09, 0x14, 0xef, 0x53, 0xf5, 0x9b, 0xed, 0x45, 0x5f, 0xc1, 0x05,
	0xfe, 0x78, 0x67, 0x62, 0xb0, 0xd0, 0x8f, 0x90, 0x7d, 0x53, 0x4f, 0xbb, 0x16, 0x39, 0xf4, 0xa8,
	0x3a, 0x7f, 0x5e, 0x37, 0x93, 0x5b, 0x33, 0x8b, 0x3d, 0x41, 0x31, 0xa6, 0xea, 0x04, 0xaa, 0x08,
	0xde, 0xcd, 0xe0, 0x1f, 0xaf, 0x6a, 0x0b, 0xd7, 0x1f, 0xf1, 0xcb, 0xa7, 0xda, 0xff, 0x95, 0x6b,
	0xc8, 0x58, 0xec, 0x81, 0xfc, 0x57, 0x87, 0x93, 0xf1, 0x2d, 0xac, 0xd8, 0xf8, 0xb4, 0x89, 0xda,
	0x21, 0x33, 0x0e, 0x88, 0x5d, 0x79, 0x1e, 0xc6, 0xfd, 0x65, 0xbc, 0xda, 0xcb, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd9, 0xe5, 0xee, 0x7d, 0x4e, 0x01, 0x00, 0x00,
}
