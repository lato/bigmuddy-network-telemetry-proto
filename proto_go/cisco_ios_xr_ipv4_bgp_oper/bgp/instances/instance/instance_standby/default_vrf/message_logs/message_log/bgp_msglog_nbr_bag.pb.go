// Code generated by protoc-gen-go.
// source: bgp_msglog_nbr_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_message_logs_message_log is a generated protocol buffer package.

It is generated from these files:
	bgp_msglog_nbr_bag.proto

It has these top-level messages:
	BgpMsglogNbrBag_KEYS
	BgpMsglogNbrBag
	BgpTimespec
	BgpNbrMsg_
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_message_logs_message_log

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

// BGP Message logging Neighbor information
type BgpMsglogNbrBag_KEYS struct {
	InstanceName    string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	Direction       uint32 `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *BgpMsglogNbrBag_KEYS) Reset()                    { *m = BgpMsglogNbrBag_KEYS{} }
func (m *BgpMsglogNbrBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpMsglogNbrBag_KEYS) ProtoMessage()               {}
func (*BgpMsglogNbrBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpMsglogNbrBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpMsglogNbrBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *BgpMsglogNbrBag_KEYS) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type BgpMsglogNbrBag struct {
	// Array of Neighbor Messages in one direction
	NeighborMessages []*BgpNbrMsg_ `protobuf:"bytes,50,rep,name=neighbor_messages,json=neighborMessages" json:"neighbor_messages,omitempty"`
}

func (m *BgpMsglogNbrBag) Reset()                    { *m = BgpMsglogNbrBag{} }
func (m *BgpMsglogNbrBag) String() string            { return proto.CompactTextString(m) }
func (*BgpMsglogNbrBag) ProtoMessage()               {}
func (*BgpMsglogNbrBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpMsglogNbrBag) GetNeighborMessages() []*BgpNbrMsg_ {
	if m != nil {
		return m.NeighborMessages
	}
	return nil
}

type BgpTimespec struct {
	// Seconds part of time value
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Nanoseconds part of time value
	Nanoseconds uint32 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *BgpTimespec) Reset()                    { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string            { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()               {}
func (*BgpTimespec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type BgpNbrMsg_ struct {
	// Type of messages Received
	MessageTypeReceived uint32 `protobuf:"varint,1,opt,name=message_type_received,json=messageTypeReceived" json:"message_type_received,omitempty"`
	// Total logged messages count
	TotalLoggedMessageCount uint32 `protobuf:"varint,2,opt,name=total_logged_message_count,json=totalLoggedMessageCount" json:"total_logged_message_count,omitempty"`
	// message received time: time elapsed since 00:00:00 UTC, January 1, 1970
	MessageTimestamp *BgpTimespec `protobuf:"bytes,3,opt,name=message_timestamp,json=messageTimestamp" json:"message_timestamp,omitempty"`
	// Message data length
	MessageDataLength uint32 `protobuf:"varint,4,opt,name=message_data_length,json=messageDataLength" json:"message_data_length,omitempty"`
	// Raw Message data in binary format
	LoggedMessageData []uint32 `protobuf:"varint,5,rep,packed,name=logged_message_data,json=loggedMessageData" json:"logged_message_data,omitempty"`
}

func (m *BgpNbrMsg_) Reset()                    { *m = BgpNbrMsg_{} }
func (m *BgpNbrMsg_) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrMsg_) ProtoMessage()               {}
func (*BgpNbrMsg_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BgpNbrMsg_) GetMessageTypeReceived() uint32 {
	if m != nil {
		return m.MessageTypeReceived
	}
	return 0
}

func (m *BgpNbrMsg_) GetTotalLoggedMessageCount() uint32 {
	if m != nil {
		return m.TotalLoggedMessageCount
	}
	return 0
}

func (m *BgpNbrMsg_) GetMessageTimestamp() *BgpTimespec {
	if m != nil {
		return m.MessageTimestamp
	}
	return nil
}

func (m *BgpNbrMsg_) GetMessageDataLength() uint32 {
	if m != nil {
		return m.MessageDataLength
	}
	return 0
}

func (m *BgpNbrMsg_) GetLoggedMessageData() []uint32 {
	if m != nil {
		return m.LoggedMessageData
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpMsglogNbrBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.message_logs.message_log.bgp_msglog_nbr_bag_KEYS")
	proto.RegisterType((*BgpMsglogNbrBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.message_logs.message_log.bgp_msglog_nbr_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.message_logs.message_log.bgp_timespec")
	proto.RegisterType((*BgpNbrMsg_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.message_logs.message_log.bgp_nbr_msg_")
}

func init() { proto.RegisterFile("bgp_msglog_nbr_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x3d, 0x8f, 0xd4, 0x30,
	0x14, 0x54, 0x2e, 0x7c, 0xe8, 0x7c, 0xb7, 0xe2, 0x30, 0x42, 0x67, 0x21, 0x8a, 0x28, 0x34, 0x4b,
	0x93, 0x62, 0xa1, 0xa3, 0x42, 0x40, 0x03, 0x07, 0x45, 0xa0, 0xa1, 0x7a, 0x72, 0xec, 0x77, 0x3e,
	0x4b, 0x89, 0x6d, 0xd9, 0xde, 0x15, 0xdb, 0x23, 0xf1, 0x3f, 0xf8, 0x17, 0xfc, 0x3b, 0x64, 0x5f,
	0x1c, 0x56, 0x5c, 0xbf, 0xd5, 0xda, 0x33, 0xe3, 0xf7, 0x66, 0x66, 0x15, 0xc2, 0x06, 0xe5, 0x60,
	0x0a, 0x6a, 0xb4, 0x0a, 0xcc, 0xe0, 0x61, 0xe0, 0xaa, 0x73, 0xde, 0x46, 0x4b, 0x95, 0xd0, 0x41,
	0x58, 0xd0, 0x36, 0xc0, 0x0f, 0x0f, 0xda, 0xed, 0x5e, 0x43, 0xd2, 0x5a, 0x87, 0xbe, 0x1b, 0x94,
	0xeb, 0xb4, 0x09, 0x91, 0x1b, 0x81, 0x61, 0x39, 0x2d, 0x07, 0x48, 0x3f, 0x72, 0xd8, 0x77, 0x12,
	0xaf, 0xf9, 0x76, 0x8c, 0xb0, 0xf3, 0xd7, 0xdd, 0x84, 0x21, 0x70, 0x85, 0x30, 0x5a, 0x15, 0x0e,
	0x2f, 0xed, 0xaf, 0x8a, 0x5c, 0xde, 0x75, 0x01, 0x9f, 0x3e, 0x7c, 0xff, 0x4a, 0x5f, 0x90, 0xd5,
	0x32, 0xd4, 0xf0, 0x09, 0x59, 0xd5, 0x54, 0xeb, 0xd3, 0xfe, 0xbc, 0x80, 0x5f, 0xf8, 0x84, 0xf4,
	0x25, 0xb9, 0x30, 0xa8, 0xd5, 0xcd, 0x60, 0x3d, 0x70, 0x29, 0x3d, 0x86, 0xc0, 0x4e, 0xb2, 0xee,
	0x51, 0xc1, 0xdf, 0xde, 0xc2, 0xf4, 0x39, 0x39, 0x95, 0xda, 0xa3, 0x88, 0xda, 0x1a, 0x56, 0x37,
	0xd5, 0x7a, 0xd5, 0xff, 0x03, 0xda, 0x3f, 0x15, 0xa1, 0x77, 0x9d, 0xd0, 0xdf, 0x15, 0x79, 0xbc,
	0x2c, 0x98, 0x9d, 0x07, 0xb6, 0x69, 0xea, 0xf5, 0xd9, 0x66, 0xdb, 0x1d, 0xa9, 0xa6, 0x34, 0x2c,
	0x3b, 0x9a, 0x82, 0x82, 0x7e, 0x09, 0xfc, 0x79, 0xb6, 0xd3, 0x7e, 0x24, 0xe7, 0x49, 0x11, 0xf5,
	0x84, 0xc1, 0xa1, 0xa0, 0x8c, 0x3c, 0x0c, 0x28, 0xac, 0x91, 0x21, 0x77, 0xb6, 0xea, 0xcb, 0x95,
	0x36, 0xe4, 0xcc, 0x70, 0x63, 0x0b, 0x7b, 0x92, 0xd9, 0x43, 0xa8, 0xfd, 0x59, 0xdf, 0x0e, 0x2b,
	0xeb, 0xe8, 0x86, 0x3c, 0x2d, 0x56, 0xe2, 0xde, 0x21, 0x78, 0x14, 0xa8, 0x77, 0x28, 0xe7, 0xd1,
	0x4f, 0x66, 0xf2, 0xdb, 0xde, 0x61, 0x3f, 0x53, 0xf4, 0x0d, 0x79, 0x16, 0x6d, 0xe4, 0x63, 0x32,
	0xaf, 0x50, 0x96, 0xe2, 0x40, 0xd8, 0xad, 0x89, 0xf3, 0xd6, 0xcb, 0xac, 0xb8, 0xca, 0x82, 0x39,
	0xc9, 0xbb, 0x44, 0xe7, 0xca, 0x97, 0x8d, 0x29, 0x52, 0xe4, 0x93, 0xcb, 0x7f, 0xd8, 0xb1, 0x2b,
	0x2f, 0x85, 0xf6, 0x17, 0x25, 0x64, 0xb1, 0x43, 0x3b, 0x52, 0x82, 0x83, 0xe4, 0x91, 0xc3, 0x88,
	0x46, 0xc5, 0x1b, 0x76, 0x2f, 0x47, 0x2b, 0xf6, 0xdf, 0xf3, 0xc8, 0xaf, 0x32, 0x91, 0xf4, 0xff,
	0x75, 0x91, 0x9e, 0xb1, 0xfb, 0x4d, 0x9d, 0xf4, 0xe3, 0x61, 0x0b, 0xe9, 0xd5, 0xf0, 0x20, 0x7f,
	0x88, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xec, 0x19, 0x73, 0xb1, 0xa4, 0x03, 0x00, 0x00,
}