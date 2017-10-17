package protocol

import (
	"github.com/giskook/ring/base"
	"github.com/giskook/ring/pb/common"
	"github.com/giskook/ring/pb/lbs_parser"
	"github.com/golang/protobuf/proto"
	"strings"
)

const (
	LOCATION_TYPE_GPS  string = "0"
	LOCATION_TYPE_CELL string = "1"
	LOCATION_TYPE_WIFI string = "2"

	LOCATION_GPS_SEP       string = ","
	LOCATION_CELL_MAIN_SEP string = "^"
	LOCATION_CELL_SUB_SEP  string = ","
	LOCATION_WIFI_MAIN_SEP string = ";"
	LOCATION_WIFI_SUB_SEP  string = ","
)

type Gps struct {
	Longitude string
	Latitude  string
	Speed     string
}

type Wifi struct {
	Mac     string
	Singnal string
}

type Cell struct {
	Lac string
	Cid string
	Dbm string
}

type ReportLocationPkg struct {
	Header        *base.Header
	Imei          string
	Serial        string
	TerminalStaus string // "00" 正常 ”x1" 拆卸 "1x" 低电
	Batt          string
	Time          string // 格式为YYMMDD-HHMMSS
	PosReason     string // 0 终端主动上报 1 终端被动上报 2 终端位置补报 3 超圈检测
	PosType       string // 0 GPS 1 cell 2 wifi
	GpsInfo       *Gps
	WifiInfo      []*Wifi
	CellInfo      []*Cell
}

func (p *ReportLocationPkg) SerializeExtra() []byte {
	extra := &Carrier.LocationExtra{
		Imei:           p.Imei,
		Serial:         p.Serial,
		TerminalStatus: p.TerminalStaus,
		Time:           p.Time,
		Batt:           p.Batt,
		PosReason:      p.PosReason,
		PosType:        p.PosType,
	}

	data, _ := proto.Marshal(extra)

	return data
}

func (p *ReportLocationPkg) SerializeLbs() []byte {
	report := &Lbs.Report{
		Header: &Lbs.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Extra: p.SerializeExtra(),
	}

	if p.PosType == LOCATION_TYPE_WIFI {
		for _, v := range p.WifiInfo {
			report.WifiCell = append(report.WifiCell, &Lbs.WifiCell{
				Mac:     v.Mac,
				Singnal: v.Singnal,
			})
		}
	} else if p.PosType == LOCATION_TYPE_CELL {
		for _, v := range p.CellInfo {
			report.StationCell = append(report.StationCell, &Lbs.StationCell{
				Lac: v.Lac,
				Cid: v.Cid,
				Dbm: v.Dbm,
			})
		}
	}

	data, _ := proto.Marshal(report)

	return data
}

func (p *ReportLocationPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOCATION,
		Location: &Carrier.ReportLocation{
			Extra:     p.SerializeExtra(),
			Longitude: p.GpsInfo.Longitude,
			Latitude:  p.GpsInfo.Latitude,
			Speed:     p.GpsInfo.Speed,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportLocation(p []string, h *base.Header, from4gps string, from4lbs string) *ReportLocationPkg {
	r := &ReportLocationPkg{
		Header:        h,
		Imei:          p[2],
		Serial:        p[3],
		TerminalStaus: p[4],
		Batt:          p[5],
		Time:          p[6],
		PosReason:     p[7],
		PosType:       p[8],
	}
	r.Header.From = from4lbs
	rest := p[9][0 : len(p[9])-2]
	if p[8] == LOCATION_TYPE_GPS {
		values := strings.Split(rest, LOCATION_GPS_SEP)
		r.Header.From = from4gps
		r.GpsInfo = &Gps{
			Longitude: values[0],
			Latitude:  values[1],
			Speed:     values[2],
		}
	} else if p[8] == LOCATION_TYPE_WIFI {
		values := strings.Split(rest, LOCATION_WIFI_MAIN_SEP)
		for _, v := range values {
			vv := strings.Split(v, LOCATION_WIFI_SUB_SEP)
			r.WifiInfo = append(r.WifiInfo, &Wifi{
				Mac:     vv[0],
				Singnal: vv[1],
			})
		}
	} else if p[8] == LOCATION_TYPE_CELL {
		values := strings.Split(rest, LOCATION_CELL_MAIN_SEP)
		for _, v := range values {
			vv := strings.Split(v, LOCATION_CELL_SUB_SEP)
			r.CellInfo = append(r.CellInfo, &Cell{
				Lac: vv[0],
				Cid: vv[1],
				Dbm: vv[2],
			})
		}
	}

	return r
}
