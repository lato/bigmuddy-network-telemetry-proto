// Code generated by protoc-gen-go.
// source: bgp_vrf_rt_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rt_entries_rt_entry is a generated protocol buffer package.

It is generated from these files:
	bgp_vrf_rt_bag.proto

It has these top-level messages:
	BgpVrfRtBag_KEYS
	BgpVrfRtBag
	BgpEdmRtEntry_
	BgpEdmVrfRtTbl_
	BgpEdmVrfRtTblList_
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rt_entries_rt_entry

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

type BgpVrfRtBag_KEYS struct {
	InstanceName          string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	RouteTargetLowerBytes uint32 `protobuf:"varint,2,opt,name=route_target_lower_bytes,json=routeTargetLowerBytes" json:"route_target_lower_bytes,omitempty"`
	RouteTargetUpperBytes uint32 `protobuf:"varint,3,opt,name=route_target_upper_bytes,json=routeTargetUpperBytes" json:"route_target_upper_bytes,omitempty"`
	AfName                string `protobuf:"bytes,4,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	RequestId             uint32 `protobuf:"varint,5,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
}

func (m *BgpVrfRtBag_KEYS) Reset()                    { *m = BgpVrfRtBag_KEYS{} }
func (m *BgpVrfRtBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpVrfRtBag_KEYS) ProtoMessage()               {}
func (*BgpVrfRtBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpVrfRtBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpVrfRtBag_KEYS) GetRouteTargetLowerBytes() uint32 {
	if m != nil {
		return m.RouteTargetLowerBytes
	}
	return 0
}

func (m *BgpVrfRtBag_KEYS) GetRouteTargetUpperBytes() uint32 {
	if m != nil {
		return m.RouteTargetUpperBytes
	}
	return 0
}

func (m *BgpVrfRtBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpVrfRtBag_KEYS) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type BgpVrfRtBag struct {
	Rt  *BgpEdmRtEntry_ `protobuf:"bytes,50,opt,name=rt" json:"rt,omitempty"`
	Afs string          `protobuf:"bytes,51,opt,name=afs" json:"afs,omitempty"`
}

func (m *BgpVrfRtBag) Reset()                    { *m = BgpVrfRtBag{} }
func (m *BgpVrfRtBag) String() string            { return proto.CompactTextString(m) }
func (*BgpVrfRtBag) ProtoMessage()               {}
func (*BgpVrfRtBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpVrfRtBag) GetRt() *BgpEdmRtEntry_ {
	if m != nil {
		return m.Rt
	}
	return nil
}

func (m *BgpVrfRtBag) GetAfs() string {
	if m != nil {
		return m.Afs
	}
	return ""
}

type BgpEdmRtEntry_ struct {
	RouteTarget string `protobuf:"bytes,1,opt,name=route_target,json=routeTarget" json:"route_target,omitempty"`
}

func (m *BgpEdmRtEntry_) Reset()                    { *m = BgpEdmRtEntry_{} }
func (m *BgpEdmRtEntry_) String() string            { return proto.CompactTextString(m) }
func (*BgpEdmRtEntry_) ProtoMessage()               {}
func (*BgpEdmRtEntry_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpEdmRtEntry_) GetRouteTarget() string {
	if m != nil {
		return m.RouteTarget
	}
	return ""
}

type BgpEdmVrfRtTbl_ struct {
	Vrf string `protobuf:"bytes,1,opt,name=vrf" json:"vrf,omitempty"`
	Afi int32  `protobuf:"zigzag32,2,opt,name=afi" json:"afi,omitempty"`
}

func (m *BgpEdmVrfRtTbl_) Reset()                    { *m = BgpEdmVrfRtTbl_{} }
func (m *BgpEdmVrfRtTbl_) String() string            { return proto.CompactTextString(m) }
func (*BgpEdmVrfRtTbl_) ProtoMessage()               {}
func (*BgpEdmVrfRtTbl_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BgpEdmVrfRtTbl_) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *BgpEdmVrfRtTbl_) GetAfi() int32 {
	if m != nil {
		return m.Afi
	}
	return 0
}

type BgpEdmVrfRtTblList_ struct {
	Tables []*BgpEdmVrfRtTbl_ `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
}

func (m *BgpEdmVrfRtTblList_) Reset()                    { *m = BgpEdmVrfRtTblList_{} }
func (m *BgpEdmVrfRtTblList_) String() string            { return proto.CompactTextString(m) }
func (*BgpEdmVrfRtTblList_) ProtoMessage()               {}
func (*BgpEdmVrfRtTblList_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BgpEdmVrfRtTblList_) GetTables() []*BgpEdmVrfRtTbl_ {
	if m != nil {
		return m.Tables
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpVrfRtBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rt_entries.rt_entry.bgp_vrf_rt_bag_KEYS")
	proto.RegisterType((*BgpVrfRtBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rt_entries.rt_entry.bgp_vrf_rt_bag")
	proto.RegisterType((*BgpEdmRtEntry_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rt_entries.rt_entry.bgp_edm_rt_entry_")
	proto.RegisterType((*BgpEdmVrfRtTbl_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rt_entries.rt_entry.bgp_edm_vrf_rt_tbl_")
	proto.RegisterType((*BgpEdmVrfRtTblList_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rt_entries.rt_entry.bgp_edm_vrf_rt_tbl_list_")
}

func init() { proto.RegisterFile("bgp_vrf_rt_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbd, 0x4e, 0xe3, 0x40,
	0x14, 0x85, 0x35, 0xc9, 0x6e, 0x56, 0x99, 0x24, 0xab, 0x8d, 0x77, 0x57, 0x3b, 0xcd, 0x4a, 0xc6,
	0x34, 0xae, 0x5c, 0x24, 0x08, 0x44, 0x8b, 0x44, 0x81, 0x40, 0x14, 0x86, 0x14, 0x54, 0x57, 0x63,
	0x67, 0x6c, 0x06, 0x39, 0xb6, 0x99, 0xb9, 0x31, 0x49, 0xcb, 0x73, 0x20, 0xf1, 0x72, 0x3c, 0x08,
	0x9a, 0xb1, 0x63, 0x25, 0x0a, 0x6d, 0xba, 0xa3, 0x7b, 0xe6, 0xf3, 0xb9, 0x3f, 0xa6, 0x7f, 0xa2,
	0xb4, 0x84, 0x4a, 0x25, 0xa0, 0x10, 0x22, 0x9e, 0x06, 0xa5, 0x2a, 0xb0, 0x70, 0x66, 0xb1, 0xd4,
	0x71, 0x01, 0xb2, 0xd0, 0xb0, 0x52, 0x20, 0xcb, 0xea, 0x04, 0xcc, 0xbb, 0xa2, 0x14, 0x2a, 0x88,
	0xd2, 0x32, 0x90, 0xb9, 0x46, 0x9e, 0xc7, 0x42, 0xb7, 0xaa, 0x15, 0xc0, 0x63, 0x94, 0x95, 0x08,
	0x14, 0x82, 0xc8, 0x51, 0x49, 0xa1, 0x37, 0x72, 0xed, 0x7d, 0x10, 0xfa, 0x7b, 0x37, 0x0f, 0xae,
	0x2f, 0x1f, 0xee, 0x9c, 0x63, 0x3a, 0x6a, 0xf1, 0x9c, 0x2f, 0x04, 0x23, 0x2e, 0xf1, 0xfb, 0xe1,
	0x70, 0x53, 0xbc, 0xe5, 0x0b, 0xe1, 0x9c, 0x51, 0xa6, 0x8a, 0x25, 0x0a, 0x40, 0xae, 0x52, 0x81,
	0x90, 0x15, 0x2f, 0x42, 0x41, 0xb4, 0x46, 0xa1, 0x59, 0xc7, 0x25, 0xfe, 0x28, 0xfc, 0x6b, 0xfd,
	0x7b, 0x6b, 0xdf, 0x18, 0xf7, 0xc2, 0x98, 0x7b, 0xe0, 0xb2, 0x2c, 0x5b, 0xb0, 0xbb, 0x07, 0xce,
	0x8c, 0x5b, 0x83, 0xff, 0xe8, 0x0f, 0x9e, 0xd4, 0x0d, 0x7d, 0xb3, 0x0d, 0xf5, 0x78, 0x62, 0x5b,
	0xf9, 0x4f, 0xa9, 0x12, 0xcf, 0x4b, 0xa1, 0x11, 0xe4, 0x9c, 0x7d, 0xb7, 0xdf, 0xe8, 0x37, 0x95,
	0xab, 0xb9, 0xf7, 0x46, 0xe8, 0xcf, 0xdd, 0x31, 0x9d, 0x15, 0xed, 0x28, 0x64, 0x13, 0x97, 0xf8,
	0x83, 0xc9, 0x63, 0x70, 0x90, 0xed, 0x1a, 0x14, 0xc4, 0x7c, 0x01, 0x9b, 0x02, 0x84, 0x1d, 0x85,
	0xce, 0x2f, 0xda, 0xe5, 0x89, 0x66, 0x53, 0x3b, 0x80, 0x91, 0xde, 0x29, 0x1d, 0xef, 0x3d, 0x75,
	0x8e, 0xe8, 0x70, 0x7b, 0x49, 0xcd, 0x05, 0x06, 0x5b, 0x8b, 0xf1, 0xce, 0xeb, 0xe3, 0x19, 0xae,
	0x99, 0x0c, 0xa3, 0x0c, 0x4c, 0x40, 0xa5, 0x92, 0x06, 0x30, 0xb2, 0x8e, 0x94, 0xf6, 0x28, 0x63,
	0x13, 0x29, 0xbd, 0x77, 0x42, 0xd9, 0x17, 0x6c, 0x26, 0x35, 0x82, 0xf3, 0x4a, 0x68, 0x0f, 0x79,
	0x94, 0x09, 0xcd, 0x88, 0xdb, 0xf5, 0x07, 0x93, 0xa7, 0x03, 0x2f, 0x68, 0xab, 0x83, 0xb0, 0x49,
	0x8e, 0x7a, 0xf6, 0xc7, 0x9f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xbe, 0xa8, 0xc1, 0x10,
	0x03, 0x00, 0x00,
}