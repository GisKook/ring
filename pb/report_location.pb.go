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
func (*WifiCell) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
func (*StationCell) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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

type LocationExtra struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// serial
	Serial string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	// terminal status  "00" 正常 ”x1" 拆卸 "1x" 低电
	TerminalStatus string `protobuf:"bytes,3,opt,name=terminal_status" json:"terminal_status,omitempty"`
	// time  格式为YYMMDD-HHMMSS
	Time string `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	// 电量
	Batt string `protobuf:"bytes,5,opt,name=batt" json:"batt,omitempty"`
	// 0 终端主动上报 1 终端被动上报 2 终端位置补报 3 超圈检测
	PosReason string `protobuf:"bytes,6,opt,name=pos_reason" json:"pos_reason,omitempty"`
	// 定位类型 0=GPS位置，1=基站位置,2=wifi定位
	PosType string `protobuf:"bytes,7,opt,name=pos_type" json:"pos_type,omitempty"`
}

func (m *LocationExtra) Reset()                    { *m = LocationExtra{} }
func (m *LocationExtra) String() string            { return proto.CompactTextString(m) }
func (*LocationExtra) ProtoMessage()               {}
func (*LocationExtra) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *LocationExtra) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *LocationExtra) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *LocationExtra) GetTerminalStatus() string {
	if m != nil {
		return m.TerminalStatus
	}
	return ""
}

func (m *LocationExtra) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *LocationExtra) GetBatt() string {
	if m != nil {
		return m.Batt
	}
	return ""
}

func (m *LocationExtra) GetPosReason() string {
	if m != nil {
		return m.PosReason
	}
	return ""
}

func (m *LocationExtra) GetPosType() string {
	if m != nil {
		return m.PosType
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
func (*LocationInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

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
func (*ReportLocation) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

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
	proto.RegisterType((*LocationExtra)(nil), "Carrier.LocationExtra")
	proto.RegisterType((*LocationInfo)(nil), "Carrier.LocationInfo")
	proto.RegisterType((*ReportLocation)(nil), "Carrier.ReportLocation")
}

func init() { proto.RegisterFile("report_location.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0x55, 0xd2, 0xdf, 0x4d, 0xfa, 0x67, 0x81, 0xf0, 0x11, 0x55, 0x1c, 0x2a, 0x0e, 0x3d,
	0x00, 0x6f, 0x50, 0x71, 0x40, 0xe2, 0x44, 0x0f, 0x1c, 0x23, 0x37, 0xd9, 0x56, 0x2b, 0x39, 0x76,
	0x64, 0x6f, 0x05, 0xbc, 0x01, 0x8f, 0x8d, 0x62, 0x27, 0x11, 0xe2, 0xb8, 0xa3, 0x6f, 0x3d, 0xb3,
	0x63, 0xb8, 0x71, 0x58, 0x5b, 0xc7, 0xb9, 0xb6, 0x85, 0x62, 0xb2, 0x66, 0x57, 0x3b, 0xcb, 0x56,
	0x4c, 0xf6, 0xca, 0x39, 0x42, 0xb7, 0xd9, 0xc2, 0xf4, 0x83, 0x4e, 0xb4, 0x47, 0xad, 0x45, 0x0a,
	0x49, 0xa5, 0x0a, 0x39, 0xb8, 0x1b, 0x6c, 0x67, 0x62, 0x09, 0x13, 0x4f, 0xe6, 0x6c, 0x94, 0x96,
	0x57, 0x8d, 0xb0, 0x79, 0x86, 0xf4, 0xc0, 0xe1, 0x8d, 0x0e, 0xd6, 0x3d, 0x9c, 0x42, 0x52, 0x50,
	0x19, 0xc1, 0x66, 0x28, 0x8f, 0x95, 0x4c, 0xc2, 0xd6, 0xcf, 0x00, 0xe6, 0x6f, 0xad, 0xf7, 0xcb,
	0x17, 0x3b, 0x25, 0x32, 0x18, 0x52, 0x85, 0xd4, 0x6e, 0x2e, 0x60, 0xec, 0xd1, 0x51, 0xe7, 0x22,
	0x6e, 0x61, 0xc9, 0xe8, 0x2a, 0x32, 0x4a, 0xe7, 0x9e, 0x15, 0x5f, 0x7c, 0x7c, 0xa8, 0x59, 0x63,
	0xaa, 0x50, 0x0e, 0xbb, 0xe9, 0xa8, 0x98, 0xe5, 0x28, 0x4c, 0x02, 0xa0, 0xb6, 0x3e, 0x77, 0xa8,
	0xbc, 0x35, 0x72, 0x1c, 0xb4, 0x15, 0x4c, 0x1b, 0x8d, 0xbf, 0x6b, 0x94, 0x93, 0x10, 0xc5, 0x42,
	0xd6, 0x25, 0x79, 0x35, 0x27, 0x2b, 0xe6, 0x30, 0xc2, 0x26, 0x51, 0x48, 0x92, 0x89, 0x7b, 0x98,
	0x7d, 0xd2, 0x89, 0xf2, 0x02, 0x75, 0x13, 0x26, 0xd9, 0xa6, 0x8f, 0xeb, 0x5d, 0x5b, 0xd3, 0xae,
	0xef, 0xe8, 0x01, 0x32, 0x1f, 0x5b, 0x88, 0x60, 0x12, 0xc0, 0xeb, 0x1e, 0xfc, 0x53, 0xd1, 0xe6,
	0x00, 0x8b, 0xf7, 0xd0, 0x7e, 0x67, 0xfb, 0xdf, 0x72, 0x0d, 0x33, 0x6d, 0xcd, 0x99, 0xf8, 0x52,
	0x62, 0x7b, 0xff, 0x0a, 0xa6, 0x5a, 0x71, 0x54, 0xe2, 0xe1, 0x73, 0x18, 0xf9, 0x1a, 0xb1, 0x8c,
	0x97, 0x1f, 0xc7, 0xe1, 0x03, 0x9f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x26, 0x86, 0xd6,
	0xd9, 0x01, 0x00, 0x00,
}
