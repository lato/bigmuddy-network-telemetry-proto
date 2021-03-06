// Code generated by protoc-gen-go.
// source: ipv6_nd_neighbor_entry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_neighbor_interfaces_neighbor_interface_host_addresses_host_address is a generated protocol buffer package.

It is generated from these files:
	ipv6_nd_neighbor_entry.proto

It has these top-level messages:
	Ipv6NdNeighborEntry_KEYS
	Ipv6NdNeighborEntry
	BagTimespec
*/
package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_neighbor_interfaces_neighbor_interface_host_addresses_host_address

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

// IPv6 ND neighbor entry
type Ipv6NdNeighborEntry_KEYS struct {
	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	HostAddress   string `protobuf:"bytes,3,opt,name=host_address,json=hostAddress" json:"host_address,omitempty"`
}

func (m *Ipv6NdNeighborEntry_KEYS) Reset()                    { *m = Ipv6NdNeighborEntry_KEYS{} }
func (m *Ipv6NdNeighborEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdNeighborEntry_KEYS) ProtoMessage()               {}
func (*Ipv6NdNeighborEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6NdNeighborEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdNeighborEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6NdNeighborEntry_KEYS) GetHostAddress() string {
	if m != nil {
		return m.HostAddress
	}
	return ""
}

type Ipv6NdNeighborEntry struct {
	// Last time of reachability
	LastReachedTime *BagTimespec `protobuf:"bytes,50,opt,name=last_reached_time,json=lastReachedTime" json:"last_reached_time,omitempty"`
	// Current state
	ReachabilityState string `protobuf:"bytes,51,opt,name=reachability_state,json=reachabilityState" json:"reachability_state,omitempty"`
	// Link-Layer Address
	LinkLayerAddress string `protobuf:"bytes,52,opt,name=link_layer_address,json=linkLayerAddress" json:"link_layer_address,omitempty"`
	// Preferred media encap type
	Encapsulation string `protobuf:"bytes,53,opt,name=encapsulation" json:"encapsulation,omitempty"`
	// Selected media encap
	SelectedEncapsulation string `protobuf:"bytes,54,opt,name=selected_encapsulation,json=selectedEncapsulation" json:"selected_encapsulation,omitempty"`
	// Neighbor origin
	OriginEncapsulation string `protobuf:"bytes,55,opt,name=origin_encapsulation,json=originEncapsulation" json:"origin_encapsulation,omitempty"`
	// Interface name
	InterfaceName string `protobuf:"bytes,56,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Location where the neighbor entry exists
	Location string `protobuf:"bytes,57,opt,name=location" json:"location,omitempty"`
	// IsRouter
	IsRouter bool `protobuf:"varint,58,opt,name=is_router,json=isRouter" json:"is_router,omitempty"`
}

func (m *Ipv6NdNeighborEntry) Reset()                    { *m = Ipv6NdNeighborEntry{} }
func (m *Ipv6NdNeighborEntry) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdNeighborEntry) ProtoMessage()               {}
func (*Ipv6NdNeighborEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6NdNeighborEntry) GetLastReachedTime() *BagTimespec {
	if m != nil {
		return m.LastReachedTime
	}
	return nil
}

func (m *Ipv6NdNeighborEntry) GetReachabilityState() string {
	if m != nil {
		return m.ReachabilityState
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetLinkLayerAddress() string {
	if m != nil {
		return m.LinkLayerAddress
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetSelectedEncapsulation() string {
	if m != nil {
		return m.SelectedEncapsulation
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetOriginEncapsulation() string {
	if m != nil {
		return m.OriginEncapsulation
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetIsRouter() bool {
	if m != nil {
		return m.IsRouter
	}
	return false
}

// Timespec specifying the number of seconds since the base time of 00:00:00 GMT, 1 January 1970.
type BagTimespec struct {
	// Number of seconds
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
}

func (m *BagTimespec) Reset()                    { *m = BagTimespec{} }
func (m *BagTimespec) String() string            { return proto.CompactTextString(m) }
func (*BagTimespec) ProtoMessage()               {}
func (*BagTimespec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BagTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6NdNeighborEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_interfaces.neighbor_interface.host_addresses.host_address.ipv6_nd_neighbor_entry_KEYS")
	proto.RegisterType((*Ipv6NdNeighborEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_interfaces.neighbor_interface.host_addresses.host_address.ipv6_nd_neighbor_entry")
	proto.RegisterType((*BagTimespec)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_interfaces.neighbor_interface.host_addresses.host_address.bag_timespec")
}

func init() { proto.RegisterFile("ipv6_nd_neighbor_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0x54, 0x00, 0x41, 0xd6, 0xbb, 0x05, 0xd6, 0xc0, 0x2a, 0x62, 0x39, 0x94, 0x15, 0x48, 0x39,
	0x40, 0x24, 0x5a, 0x5a, 0x3e, 0x6e, 0x1c, 0x7a, 0x02, 0x71, 0x48, 0xb9, 0x70, 0xb2, 0x1c, 0xfb,
	0xd1, 0x5a, 0x24, 0x76, 0xe4, 0xe7, 0x56, 0x44, 0x5c, 0xb9, 0xf0, 0x27, 0xf8, 0x19, 0xfc, 0x3e,
	0x64, 0xa7, 0x89, 0x1a, 0xda, 0xfb, 0x5e, 0x2a, 0xcd, 0x9b, 0x79, 0x7e, 0xd3, 0xd1, 0x84, 0x3c,
	0x51, 0xf5, 0x76, 0xce, 0xb4, 0x64, 0x1a, 0xd4, 0x6a, 0x5d, 0x18, 0xcb, 0x40, 0x3b, 0xdb, 0x64,
	0xb5, 0x35, 0xce, 0xd0, 0x9f, 0x42, 0xa1, 0x30, 0x4c, 0x19, 0x64, 0x3f, 0x2c, 0xeb, 0xa4, 0xa6,
	0x06, 0x9b, 0xb5, 0xc0, 0x48, 0x60, 0xd2, 0x6b, 0xb6, 0x60, 0x9b, 0xcc, 0x43, 0x0c, 0xbf, 0x59,
	0xff, 0x9c, 0xd2, 0x0e, 0xec, 0x37, 0x2e, 0x3c, 0x71, 0x30, 0xcb, 0xd6, 0x06, 0x1d, 0xe3, 0x52,
	0x5a, 0x40, 0x04, 0x1c, 0xc0, 0xab, 0x5f, 0x11, 0xb9, 0x3c, 0xee, 0x8e, 0x7d, 0x5c, 0x7c, 0x5d,
	0xd2, 0x4b, 0x72, 0x12, 0xee, 0x6b, 0x5e, 0x41, 0x12, 0x8d, 0xa3, 0xf4, 0x24, 0x8f, 0xfd, 0xe0,
	0x33, 0xaf, 0x80, 0x3e, 0x27, 0x77, 0xfb, 0x2b, 0xad, 0xe2, 0x46, 0x50, 0x8c, 0xfa, 0x69, 0x90,
	0x3d, 0x25, 0x67, 0xfb, 0x37, 0x93, 0x9b, 0x41, 0x74, 0xea, 0x67, 0x1f, 0x76, 0x36, 0xfe, 0xdc,
	0x22, 0x17, 0xc7, 0x6d, 0xd0, 0xbf, 0x11, 0x39, 0x2f, 0x39, 0x3a, 0x66, 0x81, 0x8b, 0x35, 0x48,
	0xe6, 0x54, 0x05, 0xc9, 0x64, 0x1c, 0xa5, 0xa7, 0x93, 0xdf, 0x51, 0x76, 0x8d, 0xe1, 0x65, 0x05,
	0x5f, 0x05, 0x37, 0x58, 0x83, 0xc8, 0xef, 0x79, 0x93, 0x79, 0xeb, 0xf1, 0x8b, 0xaa, 0x80, 0xbe,
	0x24, 0x34, 0x58, 0xe6, 0x85, 0x2a, 0x95, 0x6b, 0x18, 0x3a, 0xee, 0x20, 0x99, 0x86, 0x3f, 0x7f,
	0xbe, 0xcf, 0x2c, 0x3d, 0x41, 0x5f, 0x10, 0x5a, 0x2a, 0xfd, 0x9d, 0x95, 0xbc, 0x01, 0xdb, 0x67,
	0xf5, 0x3a, 0xc8, 0xef, 0x7b, 0xe6, 0x93, 0x27, 0x76, 0x81, 0xd1, 0x67, 0x64, 0x04, 0x5a, 0xf0,
	0x1a, 0x37, 0x25, 0x77, 0xca, 0xe8, 0x64, 0xd6, 0x26, 0x3f, 0x18, 0xd2, 0x19, 0xb9, 0x40, 0x28,
	0x41, 0x38, 0x90, 0x6c, 0x28, 0x9f, 0x07, 0xf9, 0xa3, 0x8e, 0x5d, 0x0c, 0xd6, 0x5e, 0x91, 0x87,
	0xc6, 0xaa, 0x95, 0xd2, 0xff, 0x2d, 0xbd, 0x09, 0x4b, 0x0f, 0x5a, 0x6e, 0xb8, 0x72, 0x58, 0x85,
	0xb7, 0xc7, 0xaa, 0xf0, 0x98, 0xc4, 0xa5, 0x11, 0xed, 0x6b, 0xef, 0xda, 0x36, 0x75, 0xd8, 0x57,
	0x4d, 0x21, 0xb3, 0x66, 0xe3, 0xc0, 0x26, 0xef, 0xc7, 0x51, 0x1a, 0xe7, 0xb1, 0xc2, 0x3c, 0xe0,
	0xab, 0x94, 0x9c, 0xed, 0xa7, 0x4d, 0x13, 0x72, 0x07, 0x41, 0x18, 0x2d, 0x31, 0xb4, 0x72, 0x94,
	0x77, 0xb0, 0xb8, 0x1d, 0xbe, 0xaa, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x94, 0xfa,
	0xf7, 0x75, 0x03, 0x00, 0x00,
}
