// Code generated by protoc-gen-go.
// source: snmp_traps_info.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_trap_infos_trap_info is a generated protocol buffer package.

It is generated from these files:
	snmp_traps_info.proto

It has these top-level messages:
	SnmpTrapsInfo_KEYS
	SnmpTrapsInfo
	SnmpTrapOidInfo
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_trap_infos_trap_info

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

type SnmpTrapsInfo_KEYS struct {
	TrapHost string `protobuf:"bytes,1,opt,name=trap_host,json=trapHost" json:"trap_host,omitempty"`
	Port     uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *SnmpTrapsInfo_KEYS) Reset()                    { *m = SnmpTrapsInfo_KEYS{} }
func (m *SnmpTrapsInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpTrapsInfo_KEYS) ProtoMessage()               {}
func (*SnmpTrapsInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SnmpTrapsInfo_KEYS) GetTrapHost() string {
	if m != nil {
		return m.TrapHost
	}
	return ""
}

func (m *SnmpTrapsInfo_KEYS) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type SnmpTrapsInfo struct {
	// NMS/Host address
	Host string `protobuf:"bytes,50,opt,name=host" json:"host,omitempty"`
	// udp port number
	Port uint32 `protobuf:"varint,51,opt,name=port" json:"port,omitempty"`
	// Total number of OID's sent
	TrapOidCount uint32 `protobuf:"varint,52,opt,name=trap_oid_count,json=trapOidCount" json:"trap_oid_count,omitempty"`
	// Per trap OID statistics
	TrapOiDinfo []*SnmpTrapOidInfo `protobuf:"bytes,53,rep,name=trap_oi_dinfo,json=trapOiDinfo" json:"trap_oi_dinfo,omitempty"`
}

func (m *SnmpTrapsInfo) Reset()                    { *m = SnmpTrapsInfo{} }
func (m *SnmpTrapsInfo) String() string            { return proto.CompactTextString(m) }
func (*SnmpTrapsInfo) ProtoMessage()               {}
func (*SnmpTrapsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpTrapsInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SnmpTrapsInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SnmpTrapsInfo) GetTrapOidCount() uint32 {
	if m != nil {
		return m.TrapOidCount
	}
	return 0
}

func (m *SnmpTrapsInfo) GetTrapOiDinfo() []*SnmpTrapOidInfo {
	if m != nil {
		return m.TrapOiDinfo
	}
	return nil
}

type SnmpTrapOidInfo struct {
	// TRAP OID
	TrapOid string `protobuf:"bytes,1,opt,name=trap_oid,json=trapOid" json:"trap_oid,omitempty"`
	// Number of traps sent
	Count uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	// Number of Traps Dropped
	DropCount uint32 `protobuf:"varint,3,opt,name=drop_count,json=dropCount" json:"drop_count,omitempty"`
	// Num of times retry
	RetryCount uint32 `protobuf:"varint,4,opt,name=retry_count,json=retryCount" json:"retry_count,omitempty"`
	// Timestamp of latest successfully sent
	LastsentTime string `protobuf:"bytes,5,opt,name=lastsent_time,json=lastsentTime" json:"lastsent_time,omitempty"`
	// Timestamp of latest droped
	LasrdropTime string `protobuf:"bytes,6,opt,name=lasrdrop_time,json=lasrdropTime" json:"lasrdrop_time,omitempty"`
}

func (m *SnmpTrapOidInfo) Reset()                    { *m = SnmpTrapOidInfo{} }
func (m *SnmpTrapOidInfo) String() string            { return proto.CompactTextString(m) }
func (*SnmpTrapOidInfo) ProtoMessage()               {}
func (*SnmpTrapOidInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SnmpTrapOidInfo) GetTrapOid() string {
	if m != nil {
		return m.TrapOid
	}
	return ""
}

func (m *SnmpTrapOidInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SnmpTrapOidInfo) GetDropCount() uint32 {
	if m != nil {
		return m.DropCount
	}
	return 0
}

func (m *SnmpTrapOidInfo) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *SnmpTrapOidInfo) GetLastsentTime() string {
	if m != nil {
		return m.LastsentTime
	}
	return ""
}

func (m *SnmpTrapOidInfo) GetLasrdropTime() string {
	if m != nil {
		return m.LasrdropTime
	}
	return ""
}

func init() {
	proto.RegisterType((*SnmpTrapsInfo_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_infos.trap_info.snmp_traps_info_KEYS")
	proto.RegisterType((*SnmpTrapsInfo)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_infos.trap_info.snmp_traps_info")
	proto.RegisterType((*SnmpTrapOidInfo)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_infos.trap_info.snmp_trap_oid_info")
}

func init() { proto.RegisterFile("snmp_traps_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0xc7, 0xd9, 0xa7, 0x2f, 0x4f, 0x3b, 0x6d, 0x15, 0x42, 0x85, 0x15, 0x11, 0x4b, 0xf5, 0xd0,
	0x53, 0x0e, 0xad, 0x7e, 0x01, 0x5f, 0x50, 0xf0, 0x20, 0x54, 0x11, 0x3c, 0x85, 0x75, 0x77, 0xd5,
	0x80, 0xbb, 0x13, 0x92, 0x08, 0xea, 0x97, 0xf4, 0x03, 0xf8, 0x65, 0x64, 0x26, 0xe9, 0x16, 0xea,
	0xd1, 0xdb, 0xe4, 0x3f, 0xbf, 0xff, 0xbc, 0x11, 0xd8, 0x71, 0x75, 0x65, 0x94, 0xb7, 0x99, 0x71,
	0x4a, 0xd7, 0x4f, 0x28, 0x8d, 0x45, 0x8f, 0xe2, 0x34, 0xd7, 0x2e, 0x47, 0xa5, 0xd1, 0xa9, 0x77,
	0xab, 0x98, 0xc9, 0x9e, 0xcb, 0xda, 0x2b, 0x34, 0xa5, 0x95, 0xf4, 0x96, 0x44, 0xdb, 0x2a, 0xf3,
	0x1a, 0x6b, 0x49, 0x7e, 0xb6, 0xbb, 0x75, 0x38, 0xbd, 0x84, 0xf1, 0x46, 0x71, 0x75, 0x7d, 0xf1,
	0x70, 0x2b, 0xf6, 0xa0, 0xcf, 0xd0, 0x0b, 0x3a, 0x9f, 0x26, 0x93, 0x64, 0xd6, 0x5f, 0xf6, 0x48,
	0xb8, 0x42, 0xe7, 0x85, 0x80, 0xb6, 0x41, 0xeb, 0xd3, 0x7f, 0x93, 0x64, 0x36, 0x5a, 0x72, 0x3c,
	0xfd, 0x4e, 0x60, 0x7b, 0xa3, 0x12, 0x71, 0xec, 0x9f, 0xb3, 0x9f, 0xe3, 0xc6, 0xbb, 0x58, 0x7b,
	0xc5, 0x11, 0x6c, 0x71, 0x33, 0xd4, 0x85, 0xca, 0xf1, 0xad, 0xf6, 0xe9, 0x31, 0x67, 0x87, 0xa4,
	0xde, 0xe8, 0xe2, 0x8c, 0x34, 0xf1, 0x09, 0xa3, 0x48, 0xa9, 0x82, 0xca, 0xa7, 0x27, 0x93, 0xd6,
	0x6c, 0x30, 0xbf, 0x97, 0x7f, 0x3f, 0x83, 0x6c, 0x26, 0xe7, 0x19, 0x48, 0x5a, 0x0e, 0x42, 0xf3,
	0x73, 0x3e, 0xd3, 0x57, 0x02, 0xe2, 0x37, 0x23, 0x76, 0xa1, 0xb7, 0x12, 0xe2, 0x91, 0xfe, 0xc7,
	0x91, 0xc5, 0x18, 0x3a, 0x61, 0x95, 0x70, 0xa4, 0xf0, 0x10, 0xfb, 0x00, 0x85, 0x45, 0x13, 0xb7,
	0x6c, 0x71, 0xaa, 0x4f, 0x4a, 0x58, 0xf1, 0x00, 0x06, 0xb6, 0xf4, 0xf6, 0x23, 0xe6, 0xdb, 0x9c,
	0x07, 0x96, 0x02, 0x70, 0x08, 0xa3, 0xd7, 0xcc, 0x79, 0x47, 0xeb, 0x79, 0x5d, 0x95, 0x69, 0x87,
	0xbb, 0x0e, 0x57, 0xe2, 0x9d, 0xae, 0xca, 0x08, 0x59, 0x6e, 0xc4, 0x50, 0xb7, 0x81, 0x58, 0x24,
	0xe8, 0xb1, 0xcb, 0x7f, 0x68, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x11, 0xaf, 0x8a, 0x5c,
	0x02, 0x00, 0x00,
}