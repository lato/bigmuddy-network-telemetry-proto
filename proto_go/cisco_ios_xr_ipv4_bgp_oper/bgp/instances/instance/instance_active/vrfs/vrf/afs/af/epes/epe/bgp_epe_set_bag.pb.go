// Code generated by protoc-gen-go.
// source: bgp_epe_set_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_afs_af_epes_epe is a generated protocol buffer package.

It is generated from these files:
	bgp_epe_set_bag.proto

It has these top-level messages:
	BgpEpeSetBag_KEYS
	BgpEpeSetBag
	IPV4TunnelAddressType
	IPV4MDTAddressType
	RTConstraintAddressType
	IPV6AddressType
	BgpIpv4SrpolicyAddrT
	BgpIpv6SrpolicyAddrT
	BgpL2VpnAddrT
	L2VPNEVPNAddressType
	BgpL2VpnMspwAddrT
	IPV6MVPNAddressType
	IPV4MVPNAddressType
	LS_LSAddressType
	IPv4FlowspecAddressType
	IPv6FlowspecAddressType
	BgpAddrtype
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_afs_af_epes_epe

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

type BgpEpeSetBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName       string `protobuf:"bytes,3,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	EpeKeyLength uint32 `protobuf:"varint,4,opt,name=epe_key_length,json=epeKeyLength" json:"epe_key_length,omitempty"`
	EpeSetKey    string `protobuf:"bytes,5,opt,name=epe_set_key,json=epeSetKey" json:"epe_set_key,omitempty"`
}

func (m *BgpEpeSetBag_KEYS) Reset()                    { *m = BgpEpeSetBag_KEYS{} }
func (m *BgpEpeSetBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpEpeSetBag_KEYS) ProtoMessage()               {}
func (*BgpEpeSetBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpEpeSetBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpEpeSetBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpEpeSetBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpEpeSetBag_KEYS) GetEpeKeyLength() uint32 {
	if m != nil {
		return m.EpeKeyLength
	}
	return 0
}

func (m *BgpEpeSetBag_KEYS) GetEpeSetKey() string {
	if m != nil {
		return m.EpeSetKey
	}
	return ""
}

type BgpEpeSetBag struct {
	// Key of EPE object
	EpeKey []uint32 `protobuf:"varint,50,rep,packed,name=epe_key,json=epeKey" json:"epe_key,omitempty"`
	// EPE key length in bits
	EpeKeyLength uint32 `protobuf:"varint,51,opt,name=epe_key_length,json=epeKeyLength" json:"epe_key_length,omitempty"`
	// EPE object's version
	EpeVersion uint32 `protobuf:"varint,52,opt,name=epe_version,json=epeVersion" json:"epe_version,omitempty"`
	// EPE object's flags
	EpeFlags uint32 `protobuf:"varint,53,opt,name=epe_flags,json=epeFlags" json:"epe_flags,omitempty"`
	// Local AS Number
	EpeLocalAsn uint32 `protobuf:"varint,54,opt,name=epe_local_asn,json=epeLocalAsn" json:"epe_local_asn,omitempty"`
	// Remote AS Number
	EpeRemoteAsn uint32 `protobuf:"varint,55,opt,name=epe_remote_asn,json=epeRemoteAsn" json:"epe_remote_asn,omitempty"`
	// Remote router id
	EpeRemoteRouterId uint32 `protobuf:"varint,56,opt,name=epe_remote_router_id,json=epeRemoteRouterId" json:"epe_remote_router_id,omitempty"`
	// Local router id
	EpeLocalRouterId uint32 `protobuf:"varint,57,opt,name=epe_local_router_id,json=epeLocalRouterId" json:"epe_local_router_id,omitempty"`
	// Local address
	EpeLocalAddress *BgpAddrtype `protobuf:"bytes,58,opt,name=epe_local_address,json=epeLocalAddress" json:"epe_local_address,omitempty"`
	// Nexthop address
	EpeNextHop *BgpAddrtype `protobuf:"bytes,59,opt,name=epe_next_hop,json=epeNextHop" json:"epe_next_hop,omitempty"`
	// List of firsthops of EPE
	FirstHopList []*BgpAddrtype `protobuf:"bytes,60,rep,name=first_hop_list,json=firstHopList" json:"first_hop_list,omitempty"`
	// List of nexthop ID of EPE
	NexthopIdList []uint32 `protobuf:"varint,61,rep,packed,name=nexthop_id_list,json=nexthopIdList" json:"nexthop_id_list,omitempty"`
	// List of ifhandle of EPE
	IfhandleList []uint32 `protobuf:"varint,62,rep,packed,name=ifhandle_list,json=ifhandleList" json:"ifhandle_list,omitempty"`
	// Label assigned to the RPCnext-hop set
	Label uint32 `protobuf:"varint,63,opt,name=label" json:"label,omitempty"`
	// Refcount
	RefCount uint32 `protobuf:"varint,64,opt,name=ref_count,json=refCount" json:"ref_count,omitempty"`
}

func (m *BgpEpeSetBag) Reset()                    { *m = BgpEpeSetBag{} }
func (m *BgpEpeSetBag) String() string            { return proto.CompactTextString(m) }
func (*BgpEpeSetBag) ProtoMessage()               {}
func (*BgpEpeSetBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpEpeSetBag) GetEpeKey() []uint32 {
	if m != nil {
		return m.EpeKey
	}
	return nil
}

func (m *BgpEpeSetBag) GetEpeKeyLength() uint32 {
	if m != nil {
		return m.EpeKeyLength
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeVersion() uint32 {
	if m != nil {
		return m.EpeVersion
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeFlags() uint32 {
	if m != nil {
		return m.EpeFlags
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeLocalAsn() uint32 {
	if m != nil {
		return m.EpeLocalAsn
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeRemoteAsn() uint32 {
	if m != nil {
		return m.EpeRemoteAsn
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeRemoteRouterId() uint32 {
	if m != nil {
		return m.EpeRemoteRouterId
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeLocalRouterId() uint32 {
	if m != nil {
		return m.EpeLocalRouterId
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeLocalAddress() *BgpAddrtype {
	if m != nil {
		return m.EpeLocalAddress
	}
	return nil
}

func (m *BgpEpeSetBag) GetEpeNextHop() *BgpAddrtype {
	if m != nil {
		return m.EpeNextHop
	}
	return nil
}

func (m *BgpEpeSetBag) GetFirstHopList() []*BgpAddrtype {
	if m != nil {
		return m.FirstHopList
	}
	return nil
}

func (m *BgpEpeSetBag) GetNexthopIdList() []uint32 {
	if m != nil {
		return m.NexthopIdList
	}
	return nil
}

func (m *BgpEpeSetBag) GetIfhandleList() []uint32 {
	if m != nil {
		return m.IfhandleList
	}
	return nil
}

func (m *BgpEpeSetBag) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *BgpEpeSetBag) GetRefCount() uint32 {
	if m != nil {
		return m.RefCount
	}
	return 0
}

// IPV4Tunnel Address type
type IPV4TunnelAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4TunnelAddressType) Reset()                    { *m = IPV4TunnelAddressType{} }
func (m *IPV4TunnelAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4TunnelAddressType) ProtoMessage()               {}
func (*IPV4TunnelAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IPV4TunnelAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4MDT Address type
type IPV4MDTAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4MDTAddressType) Reset()                    { *m = IPV4MDTAddressType{} }
func (m *IPV4MDTAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4MDTAddressType) ProtoMessage()               {}
func (*IPV4MDTAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IPV4MDTAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4 RTConstraint Address type
type RTConstraintAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *RTConstraintAddressType) Reset()                    { *m = RTConstraintAddressType{} }
func (m *RTConstraintAddressType) String() string            { return proto.CompactTextString(m) }
func (*RTConstraintAddressType) ProtoMessage()               {}
func (*RTConstraintAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RTConstraintAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV6 Address type
type IPV6AddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6AddressType) Reset()                    { *m = IPV6AddressType{} }
func (m *IPV6AddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6AddressType) ProtoMessage()               {}
func (*IPV6AddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IPV6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpIpv4SrpolicyAddrT struct {
	Ipv4SrpolicyAddress []byte `protobuf:"bytes,1,opt,name=ipv4_srpolicy_address,json=ipv4SrpolicyAddress,proto3" json:"ipv4_srpolicy_address,omitempty"`
}

func (m *BgpIpv4SrpolicyAddrT) Reset()                    { *m = BgpIpv4SrpolicyAddrT{} }
func (m *BgpIpv4SrpolicyAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpIpv4SrpolicyAddrT) ProtoMessage()               {}
func (*BgpIpv4SrpolicyAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BgpIpv4SrpolicyAddrT) GetIpv4SrpolicyAddress() []byte {
	if m != nil {
		return m.Ipv4SrpolicyAddress
	}
	return nil
}

type BgpIpv6SrpolicyAddrT struct {
	Ipv6SrpolicyAddress []byte `protobuf:"bytes,1,opt,name=ipv6_srpolicy_address,json=ipv6SrpolicyAddress,proto3" json:"ipv6_srpolicy_address,omitempty"`
}

func (m *BgpIpv6SrpolicyAddrT) Reset()                    { *m = BgpIpv6SrpolicyAddrT{} }
func (m *BgpIpv6SrpolicyAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpIpv6SrpolicyAddrT) ProtoMessage()               {}
func (*BgpIpv6SrpolicyAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BgpIpv6SrpolicyAddrT) GetIpv6SrpolicyAddress() []byte {
	if m != nil {
		return m.Ipv6SrpolicyAddress
	}
	return nil
}

type BgpL2VpnAddrT struct {
	L2VpnAddress []byte `protobuf:"bytes,1,opt,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
}

func (m *BgpL2VpnAddrT) Reset()                    { *m = BgpL2VpnAddrT{} }
func (m *BgpL2VpnAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpL2VpnAddrT) ProtoMessage()               {}
func (*BgpL2VpnAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BgpL2VpnAddrT) GetL2VpnAddress() []byte {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

// L2VPN EVPN Address type
type L2VPNEVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *L2VPNEVPNAddressType) Reset()                    { *m = L2VPNEVPNAddressType{} }
func (m *L2VPNEVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*L2VPNEVPNAddressType) ProtoMessage()               {}
func (*L2VPNEVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *L2VPNEVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpL2VpnMspwAddrT struct {
	L2VpnAddress []byte `protobuf:"bytes,1,opt,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
}

func (m *BgpL2VpnMspwAddrT) Reset()                    { *m = BgpL2VpnMspwAddrT{} }
func (m *BgpL2VpnMspwAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpL2VpnMspwAddrT) ProtoMessage()               {}
func (*BgpL2VpnMspwAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BgpL2VpnMspwAddrT) GetL2VpnAddress() []byte {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

// IPV6 MVPN Address type
type IPV6MVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6MVPNAddressType) Reset()                    { *m = IPV6MVPNAddressType{} }
func (m *IPV6MVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6MVPNAddressType) ProtoMessage()               {}
func (*IPV6MVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IPV6MVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4 MVPN Address type
type IPV4MVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4MVPNAddressType) Reset()                    { *m = IPV4MVPNAddressType{} }
func (m *IPV4MVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4MVPNAddressType) ProtoMessage()               {}
func (*IPV4MVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IPV4MVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// LINKSTATE LINKSTATE Address type
type LS_LSAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *LS_LSAddressType) Reset()                    { *m = LS_LSAddressType{} }
func (m *LS_LSAddressType) String() string            { return proto.CompactTextString(m) }
func (*LS_LSAddressType) ProtoMessage()               {}
func (*LS_LSAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *LS_LSAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPv4 Flowspec Address type
type IPv4FlowspecAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPv4FlowspecAddressType) Reset()                    { *m = IPv4FlowspecAddressType{} }
func (m *IPv4FlowspecAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPv4FlowspecAddressType) ProtoMessage()               {}
func (*IPv4FlowspecAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *IPv4FlowspecAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPv6 Flowspec Address type
type IPv6FlowspecAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPv6FlowspecAddressType) Reset()                    { *m = IPv6FlowspecAddressType{} }
func (m *IPv6FlowspecAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPv6FlowspecAddressType) ProtoMessage()               {}
func (*IPv6FlowspecAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *IPv6FlowspecAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpAddrtype struct {
	Afi string `protobuf:"bytes,1,opt,name=afi" json:"afi,omitempty"`
	// IPv4 Addr
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv4 Mcast Addr
	Ipv4McastAddress string `protobuf:"bytes,3,opt,name=ipv4_mcast_address,json=ipv4McastAddress" json:"ipv4_mcast_address,omitempty"`
	// IPv4 Label Addr
	Ipv4LabelAddress string `protobuf:"bytes,4,opt,name=ipv4_label_address,json=ipv4LabelAddress" json:"ipv4_label_address,omitempty"`
	// IPv4 Tunnel
	Ipv4TunnelAddress *IPV4TunnelAddressType `protobuf:"bytes,5,opt,name=ipv4_tunnel_address,json=ipv4TunnelAddress" json:"ipv4_tunnel_address,omitempty"`
	// IPv4 MDT Addr
	Ipv4MdtAddress *IPV4MDTAddressType `protobuf:"bytes,6,opt,name=ipv4_mdt_address,json=ipv4MdtAddress" json:"ipv4_mdt_address,omitempty"`
	// IPv4 VPN Addr
	Ipv4VpnAddress string `protobuf:"bytes,7,opt,name=ipv4_vpn_address,json=ipv4VpnAddress" json:"ipv4_vpn_address,omitempty"`
	// IPv4 VPN Mcast Addr
	Ipv4VpnaMcastddress string `protobuf:"bytes,8,opt,name=ipv4_vpna_mcastddress,json=ipv4VpnaMcastddress" json:"ipv4_vpna_mcastddress,omitempty"`
	// IPV6 Addr
	Ipv6Address *IPV6AddressType `protobuf:"bytes,9,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
	// IPV6 Mcast Addr
	Ipv6McastAddress *IPV6AddressType `protobuf:"bytes,10,opt,name=ipv6_mcast_address,json=ipv6McastAddress" json:"ipv6_mcast_address,omitempty"`
	// IPv6 Label Addr
	Ipv6LabelAddress *IPV6AddressType `protobuf:"bytes,11,opt,name=ipv6_label_address,json=ipv6LabelAddress" json:"ipv6_label_address,omitempty"`
	// IPv6 VPN Addr
	Ipv6VpnAddress *IPV6AddressType `protobuf:"bytes,12,opt,name=ipv6_vpn_address,json=ipv6VpnAddress" json:"ipv6_vpn_address,omitempty"`
	// IPv6 VPN Mcast Addr
	Ipv6VpnMcastAddress *IPV6AddressType `protobuf:"bytes,13,opt,name=ipv6_vpn_mcast_address,json=ipv6VpnMcastAddress" json:"ipv6_vpn_mcast_address,omitempty"`
	// L2VPN VPLS Addr
	L2VpnvplsAddress *BgpL2VpnAddrT `protobuf:"bytes,14,opt,name=l2_vpnvpls_address,json=l2VpnvplsAddress" json:"l2_vpnvpls_address,omitempty"`
	// RT Constrt Addr
	RtConstraintAddress *RTConstraintAddressType `protobuf:"bytes,15,opt,name=rt_constraint_address,json=rtConstraintAddress" json:"rt_constraint_address,omitempty"`
	// MVPN addr
	Ipv6MvpnAddress *IPV6MVPNAddressType `protobuf:"bytes,16,opt,name=ipv6_mvpn_address,json=ipv6MvpnAddress" json:"ipv6_mvpn_address,omitempty"`
	// MVPN4 addr
	Ipv4MvpnAddress *IPV4MVPNAddressType `protobuf:"bytes,17,opt,name=ipv4_mvpn_address,json=ipv4MvpnAddress" json:"ipv4_mvpn_address,omitempty"`
	// L2VPN EVPN Addr
	L2VpnEvpnAddress *L2VPNEVPNAddressType `protobuf:"bytes,18,opt,name=l2_vpn_evpn_address,json=l2VpnEvpnAddress" json:"l2_vpn_evpn_address,omitempty"`
	// LINKSTATE LINKSTATE Addr
	LsLsAddress *LS_LSAddressType `protobuf:"bytes,19,opt,name=ls_ls_address,json=lsLsAddress" json:"ls_ls_address,omitempty"`
	// L2VPN MSPW Addr
	L2VpnMspwAddress *BgpL2VpnMspwAddrT `protobuf:"bytes,20,opt,name=l2_vpn_mspw_address,json=l2VpnMspwAddress" json:"l2_vpn_mspw_address,omitempty"`
	// IPV4 Flowspec Addr
	Ipv4FlowspecAddress *IPv4FlowspecAddressType `protobuf:"bytes,21,opt,name=ipv4_flowspec_address,json=ipv4FlowspecAddress" json:"ipv4_flowspec_address,omitempty"`
	// IPV6 Flowspec Addr
	Ipv6FlowspecAddress *IPv6FlowspecAddressType `protobuf:"bytes,22,opt,name=ipv6_flowspec_address,json=ipv6FlowspecAddress" json:"ipv6_flowspec_address,omitempty"`
	// IPV4 VPN Flowspec Addr
	Ipv4VpnFlowspecAddress *IPv4FlowspecAddressType `protobuf:"bytes,23,opt,name=ipv4_vpn_flowspec_address,json=ipv4VpnFlowspecAddress" json:"ipv4_vpn_flowspec_address,omitempty"`
	// IPV6 VPN Flowspec Addr
	Ipv6VpnFlowspecAddress *IPv6FlowspecAddressType `protobuf:"bytes,24,opt,name=ipv6_vpn_flowspec_address,json=ipv6VpnFlowspecAddress" json:"ipv6_vpn_flowspec_address,omitempty"`
	// IPV4 Policy Addr
	Ipv4SrPolicyAddress *BgpIpv4SrpolicyAddrT `protobuf:"bytes,25,opt,name=ipv4_sr_policy_address,json=ipv4SrPolicyAddress" json:"ipv4_sr_policy_address,omitempty"`
	// IPV6 Policy Addr
	Ipv6SrPolicyAddress *BgpIpv6SrpolicyAddrT `protobuf:"bytes,26,opt,name=ipv6_sr_policy_address,json=ipv6SrPolicyAddress" json:"ipv6_sr_policy_address,omitempty"`
}

func (m *BgpAddrtype) Reset()                    { *m = BgpAddrtype{} }
func (m *BgpAddrtype) String() string            { return proto.CompactTextString(m) }
func (*BgpAddrtype) ProtoMessage()               {}
func (*BgpAddrtype) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *BgpAddrtype) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4McastAddress() string {
	if m != nil {
		return m.Ipv4McastAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4LabelAddress() string {
	if m != nil {
		return m.Ipv4LabelAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4TunnelAddress() *IPV4TunnelAddressType {
	if m != nil {
		return m.Ipv4TunnelAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4MdtAddress() *IPV4MDTAddressType {
	if m != nil {
		return m.Ipv4MdtAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4VpnAddress() string {
	if m != nil {
		return m.Ipv4VpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4VpnaMcastddress() string {
	if m != nil {
		return m.Ipv4VpnaMcastddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6Address() *IPV6AddressType {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6McastAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6McastAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6LabelAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6LabelAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6VpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnMcastAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6VpnMcastAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnvplsAddress() *BgpL2VpnAddrT {
	if m != nil {
		return m.L2VpnvplsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetRtConstraintAddress() *RTConstraintAddressType {
	if m != nil {
		return m.RtConstraintAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6MvpnAddress() *IPV6MVPNAddressType {
	if m != nil {
		return m.Ipv6MvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4MvpnAddress() *IPV4MVPNAddressType {
	if m != nil {
		return m.Ipv4MvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnEvpnAddress() *L2VPNEVPNAddressType {
	if m != nil {
		return m.L2VpnEvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetLsLsAddress() *LS_LSAddressType {
	if m != nil {
		return m.LsLsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnMspwAddress() *BgpL2VpnMspwAddrT {
	if m != nil {
		return m.L2VpnMspwAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4FlowspecAddress() *IPv4FlowspecAddressType {
	if m != nil {
		return m.Ipv4FlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6FlowspecAddress() *IPv6FlowspecAddressType {
	if m != nil {
		return m.Ipv6FlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4VpnFlowspecAddress() *IPv4FlowspecAddressType {
	if m != nil {
		return m.Ipv4VpnFlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnFlowspecAddress() *IPv6FlowspecAddressType {
	if m != nil {
		return m.Ipv6VpnFlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4SrPolicyAddress() *BgpIpv4SrpolicyAddrT {
	if m != nil {
		return m.Ipv4SrPolicyAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6SrPolicyAddress() *BgpIpv6SrpolicyAddrT {
	if m != nil {
		return m.Ipv6SrPolicyAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpEpeSetBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_epe_set_bag_KEYS")
	proto.RegisterType((*BgpEpeSetBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_epe_set_bag")
	proto.RegisterType((*IPV4TunnelAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPV4TunnelAddressType")
	proto.RegisterType((*IPV4MDTAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPV4MDTAddressType")
	proto.RegisterType((*RTConstraintAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.RTConstraintAddressType")
	proto.RegisterType((*IPV6AddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPV6AddressType")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_l2vpn_addr_t")
	proto.RegisterType((*L2VPNEVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.L2VPNEVPNAddressType")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*IPV6MVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPV6MVPNAddressType")
	proto.RegisterType((*IPV4MVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPV4MVPNAddressType")
	proto.RegisterType((*LS_LSAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.LS_LSAddressType")
	proto.RegisterType((*IPv4FlowspecAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPv4FlowspecAddressType")
	proto.RegisterType((*IPv6FlowspecAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.IPv6FlowspecAddressType")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.epes.epe.bgp_addrtype")
}

func init() { proto.RegisterFile("bgp_epe_set_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x98, 0xdb, 0x6f, 0xe3, 0xc4,
	0x17, 0xc7, 0xe5, 0x5f, 0xb7, 0xb7, 0xd3, 0xa4, 0x49, 0x9d, 0x5e, 0xdc, 0xdf, 0x4a, 0x50, 0xba,
	0x08, 0x22, 0xd8, 0xcd, 0x4a, 0xd9, 0xe0, 0xe5, 0xb2, 0xdc, 0xb4, 0x74, 0xb5, 0x55, 0xd3, 0x2a,
	0x72, 0xab, 0x48, 0xf0, 0x32, 0x72, 0x93, 0x71, 0x6b, 0xd5, 0xb5, 0x07, 0xcf, 0xd4, 0x6d, 0xfe,
	0x05, 0x1e, 0x10, 0x0f, 0x20, 0x71, 0x79, 0x00, 0x9e, 0xd0, 0xbe, 0xf0, 0xc6, 0x9f, 0xc5, 0xff,
	0x80, 0xe6, 0x8c, 0xed, 0x38, 0x6e, 0x4a, 0xc3, 0x83, 0xd3, 0x97, 0x2a, 0x9e, 0xf3, 0xf5, 0x39,
	0x1f, 0x1f, 0x9f, 0x33, 0x73, 0x5c, 0x58, 0x3b, 0x3e, 0x61, 0x84, 0x32, 0x4a, 0x38, 0x15, 0xe4,
	0xd8, 0x3e, 0x69, 0xb0, 0x30, 0x10, 0x81, 0xfe, 0x55, 0xcf, 0xe5, 0xbd, 0x80, 0xb8, 0x01, 0x27,
	0x57, 0x21, 0x71, 0x59, 0xd4, 0x22, 0x52, 0x18, 0x30, 0x1a, 0x36, 0x8e, 0x4f, 0x58, 0xc3, 0xf5,
	0xb9, 0xb0, 0xfd, 0x1e, 0xe5, 0xe9, 0xaf, 0xf4, 0x07, 0xb1, 0x7b, 0xc2, 0x8d, 0x68, 0x23, 0x0a,
	0x1d, 0x2e, 0xff, 0x34, 0x6c, 0x87, 0x37, 0x6c, 0xa7, 0x41, 0x19, 0xe5, 0xf2, 0xcf, 0xf6, 0x5f,
	0x1a, 0xac, 0xe6, 0xa2, 0x92, 0xbd, 0x9d, 0x2f, 0x0f, 0xf5, 0x07, 0x50, 0x4e, 0x9d, 0xf8, 0xf6,
	0x39, 0x35, 0xb4, 0x2d, 0xad, 0xbe, 0x68, 0x95, 0x92, 0xc5, 0x03, 0xfb, 0x9c, 0xea, 0x9b, 0xb0,
	0x10, 0x85, 0x8e, 0xb2, 0xff, 0x0f, 0xed, 0xf3, 0x51, 0xe8, 0xa0, 0x69, 0x03, 0xe6, 0xed, 0xd8,
	0x32, 0x83, 0x96, 0x39, 0x5b, 0x19, 0xde, 0x84, 0x65, 0x19, 0xec, 0x8c, 0x0e, 0x88, 0x47, 0xfd,
	0x13, 0x71, 0x6a, 0xdc, 0xdb, 0xd2, 0xea, 0x65, 0xab, 0x44, 0x19, 0xdd, 0xa3, 0x83, 0x36, 0xae,
	0xe9, 0xaf, 0xc1, 0x52, 0x82, 0x74, 0x46, 0x07, 0xc6, 0x2c, 0xba, 0x58, 0xa4, 0x8c, 0x1e, 0x52,
	0xb1, 0x47, 0x07, 0xdb, 0x7f, 0xcf, 0x41, 0x25, 0xc7, 0x2d, 0x43, 0xc6, 0x9e, 0x8d, 0xe6, 0xd6,
	0x4c, 0xbd, 0x6c, 0xcd, 0x29, 0x97, 0x63, 0x42, 0x3e, 0x19, 0x13, 0xf2, 0x75, 0x15, 0x32, 0xa2,
	0x21, 0x77, 0x03, 0xdf, 0x68, 0xa1, 0x04, 0x28, 0xa3, 0x5d, 0xb5, 0xa2, 0xdf, 0x07, 0x09, 0x40,
	0x1c, 0xcf, 0x3e, 0xe1, 0xc6, 0x7b, 0x68, 0x5e, 0xa0, 0x8c, 0xbe, 0x90, 0xd7, 0xfa, 0x36, 0x94,
	0xa5, 0xd1, 0x0b, 0x7a, 0xb6, 0x47, 0x6c, 0xee, 0x1b, 0x26, 0x0a, 0xa4, 0xcb, 0xb6, 0x5c, 0xfb,
	0x9c, 0xfb, 0x09, 0x47, 0x48, 0xcf, 0x03, 0x41, 0x51, 0xf4, 0x34, 0xe5, 0xb0, 0x70, 0x51, 0xaa,
	0x1e, 0xc3, 0x6a, 0x46, 0x15, 0x06, 0x17, 0x82, 0x86, 0xc4, 0xed, 0x1b, 0xef, 0xa3, 0x76, 0x25,
	0xd5, 0x5a, 0x68, 0xd9, 0xed, 0xeb, 0x8f, 0xa0, 0x36, 0x0c, 0x3d, 0xd4, 0x7f, 0x80, 0xfa, 0x6a,
	0x02, 0x90, 0xca, 0xbf, 0xd7, 0x60, 0x25, 0x83, 0xda, 0xef, 0x87, 0x94, 0x73, 0xe3, 0xc3, 0x2d,
	0xad, 0xbe, 0xd4, 0x3c, 0x6d, 0x14, 0x57, 0x6b, 0xf2, 0x7e, 0x0c, 0x27, 0x06, 0x8c, 0x5a, 0x95,
	0x34, 0x31, 0x0a, 0x40, 0xff, 0x46, 0x03, 0x99, 0x07, 0xe2, 0xd3, 0x2b, 0x41, 0x4e, 0x03, 0x66,
	0x7c, 0x34, 0x65, 0x22, 0xf9, 0xaa, 0x0f, 0xe8, 0x95, 0x78, 0x19, 0x30, 0xfd, 0x5b, 0x0d, 0x96,
	0x1d, 0x37, 0xe4, 0x48, 0x42, 0x3c, 0x97, 0x0b, 0xe3, 0xd9, 0xd6, 0xcc, 0x54, 0x71, 0x4a, 0x18,
	0xff, 0x65, 0xc0, 0xda, 0x2e, 0x17, 0xfa, 0x5b, 0x50, 0x91, 0x89, 0x91, 0x34, 0x6e, 0x5f, 0x01,
	0x7d, 0x8c, 0x35, 0x5e, 0x8e, 0x97, 0x77, 0xfb, 0xa8, 0x93, 0x6d, 0xeb, 0x9c, 0xda, 0x7e, 0xdf,
	0xa3, 0x4a, 0xf5, 0x09, 0xaa, 0x4a, 0xc9, 0x22, 0x8a, 0x56, 0x61, 0xd6, 0xb3, 0x8f, 0xa9, 0x67,
	0x7c, 0x8a, 0x25, 0xa2, 0x2e, 0x64, 0x79, 0x87, 0xd4, 0x21, 0xbd, 0xe0, 0xc2, 0x17, 0xc6, 0x67,
	0xaa, 0xbc, 0x43, 0xea, 0x3c, 0x97, 0xd7, 0xdb, 0x8f, 0x60, 0x6d, 0xb7, 0xd3, 0x6d, 0x1d, 0x5d,
	0xf8, 0x3e, 0x4d, 0x5e, 0xd9, 0xd1, 0x80, 0x51, 0xe9, 0x2b, 0xb2, 0xbd, 0x8b, 0x64, 0x7f, 0x50,
	0x17, 0xdb, 0xef, 0x80, 0x2e, 0xe5, 0xfb, 0x5f, 0x1c, 0xdd, 0xae, 0x7d, 0x0c, 0x1b, 0xd6, 0xd1,
	0xf3, 0xc0, 0xe7, 0x22, 0xb4, 0x5d, 0x5f, 0xdc, 0x7e, 0xc3, 0xdb, 0x50, 0xd9, 0xed, 0x74, 0xcd,
	0xdb, 0x85, 0x07, 0x60, 0xc8, 0x94, 0xe2, 0x4b, 0xe2, 0x21, 0x0b, 0x3c, 0xb7, 0x37, 0xc0, 0x04,
	0x13, 0xa1, 0x37, 0x61, 0xed, 0xfa, 0xba, 0x6c, 0x04, 0xe9, 0xa1, 0x64, 0xd5, 0xa4, 0xf1, 0x30,
	0xb6, 0xc5, 0x91, 0x32, 0xfe, 0xcc, 0x1b, 0xfc, 0x99, 0xff, 0xe6, 0xcf, 0xcc, 0xfb, 0x7b, 0x0a,
	0x55, 0xe9, 0xcf, 0x6b, 0x46, 0xcc, 0x4f, 0xfc, 0x3c, 0x80, 0xf2, 0xf0, 0x7a, 0x78, 0x7f, 0x09,
	0x17, 0x93, 0x1b, 0x1f, 0xc2, 0x6a, 0xbb, 0xd9, 0xed, 0x1c, 0xec, 0x74, 0x3b, 0x07, 0xb7, 0xa7,
	0xe1, 0x99, 0x3a, 0x58, 0x94, 0xdb, 0x73, 0xce, 0x2e, 0xff, 0x53, 0xac, 0x77, 0xa1, 0x26, 0xb3,
	0xbd, 0x3f, 0x51, 0x28, 0x25, 0x6e, 0x4d, 0x26, 0xae, 0x43, 0xb5, 0x7d, 0x48, 0xda, 0x87, 0x13,
	0x95, 0xc8, 0x6e, 0x27, 0x6a, 0xbd, 0xf0, 0x82, 0x4b, 0xce, 0x68, 0x6f, 0xd2, 0x1b, 0xcc, 0xc9,
	0x6f, 0x78, 0x75, 0x1f, 0x4a, 0xd9, 0xf6, 0xd3, 0xab, 0x30, 0x63, 0x3b, 0x6e, 0x2c, 0x92, 0x3f,
	0xf5, 0x37, 0xa0, 0x84, 0x15, 0x93, 0x24, 0x4b, 0x1d, 0x78, 0x4b, 0x72, 0x2d, 0xd9, 0xc3, 0x1e,
	0x82, 0x8e, 0x92, 0xf3, 0x9e, 0xcd, 0x45, 0x2a, 0x54, 0xe7, 0x5f, 0x55, 0x5a, 0xf6, 0xa5, 0x21,
	0xaf, 0xc6, 0xf6, 0x4b, 0xd5, 0xf7, 0x86, 0xea, 0xb6, 0x34, 0x24, 0xea, 0xdf, 0x35, 0xc0, 0xa2,
	0x24, 0x02, 0x7b, 0x30, 0xd5, 0xcf, 0xe2, 0x36, 0xf9, 0x75, 0x91, 0xfb, 0xd2, 0xd8, 0xce, 0xb7,
	0x56, 0x64, 0x90, 0x91, 0x65, 0xfd, 0x47, 0x0d, 0xaa, 0x2a, 0x01, 0xfd, 0xe1, 0xe3, 0xcf, 0x21,
	0xa0, 0x5f, 0x34, 0xe0, 0xe8, 0x5e, 0x63, 0x2d, 0x63, 0xba, 0xfb, 0x69, 0xb2, 0xeb, 0x31, 0x59,
	0xb6, 0xdc, 0xe7, 0x31, 0xd5, 0xa8, 0xec, 0xa6, 0x05, 0x9f, 0xee, 0x0c, 0x11, 0xf3, 0x6d, 0xf5,
	0x26, 0x63, 0xf9, 0x02, 0xca, 0x6b, 0xb1, 0xdc, 0xde, 0x1f, 0x9a, 0xe4, 0x79, 0x51, 0xc2, 0xf6,
	0x4f, 0x5c, 0x2f, 0xe2, 0x43, 0x9f, 0x15, 0xfc, 0xd0, 0xd9, 0x3d, 0x10, 0x2b, 0x31, 0x59, 0xd0,
	0x7f, 0xd2, 0xb0, 0xb8, 0xcc, 0x5c, 0x29, 0xc2, 0xf4, 0xb1, 0x64, 0xda, 0xcd, 0x91, 0xba, 0x4f,
	0xd9, 0x46, 0x0b, 0x7f, 0xe9, 0x8e, 0xd8, 0x46, 0xba, 0xec, 0x07, 0x55, 0xc1, 0xe6, 0x48, 0x9d,
	0x94, 0xa6, 0x4f, 0x26, 0x8b, 0xd2, 0xcc, 0x14, 0xe5, 0x6f, 0x1a, 0xac, 0xa7, 0x5c, 0xa3, 0xef,
	0xb4, 0x3c, 0x7d, 0xba, 0x5a, 0x4c, 0x37, 0xf2, 0x5a, 0x7f, 0xd6, 0x40, 0xf7, 0x9a, 0x12, 0x30,
	0x62, 0x1e, 0x4f, 0xf1, 0x96, 0x11, 0xcf, 0x2b, 0x7a, 0x6e, 0xca, 0x1e, 0xa2, 0x56, 0xd5, 0x6b,
	0x76, 0x15, 0x46, 0x02, 0xf7, 0x87, 0x06, 0x6b, 0xa1, 0x20, 0xbd, 0x74, 0xcc, 0x48, 0xf9, 0x2a,
	0xc8, 0xc7, 0x8b, 0xe4, 0xbb, 0x61, 0xbc, 0xb1, 0x6a, 0xa1, 0xb8, 0x66, 0xd0, 0x7f, 0xd1, 0x60,
	0x45, 0x75, 0x6e, 0xb6, 0x04, 0xab, 0x48, 0x19, 0x14, 0xfd, 0x92, 0x73, 0x07, 0xb7, 0x55, 0xc1,
	0xe6, 0x1d, 0x4e, 0x03, 0x09, 0x5d, 0x6b, 0x94, 0x6e, 0x65, 0x2a, 0x74, 0xad, 0x71, 0x74, 0xad,
	0x2c, 0xdd, 0xaf, 0x1a, 0xd4, 0x54, 0x09, 0x12, 0x9a, 0xe5, 0xd3, 0x91, 0x8f, 0x15, 0xc9, 0x37,
	0x6e, 0x1e, 0x8b, 0xeb, 0x70, 0x27, 0x43, 0xf8, 0x9d, 0x06, 0x65, 0x8f, 0x93, 0x4c, 0x7f, 0xd4,
	0x8a, 0xef, 0x8f, 0xfc, 0x94, 0x65, 0x2d, 0x79, 0xbc, 0xcd, 0x33, 0x5b, 0x4b, 0x92, 0xb4, 0x74,
	0x38, 0x94, 0x60, 0xab, 0xc5, 0x0f, 0x16, 0x63, 0xc7, 0xd2, 0x38, 0x6b, 0xfb, 0x9c, 0x5d, 0x66,
	0xbb, 0x17, 0x23, 0x3b, 0xf1, 0x40, 0x97, 0x42, 0xae, 0x15, 0xdf, 0xbd, 0x37, 0x4c, 0x9e, 0x6a,
	0x10, 0xc8, 0x19, 0x12, 0x52, 0xf3, 0x3a, 0xe9, 0xfa, 0x54, 0x48, 0xcd, 0x9b, 0x48, 0xf3, 0x06,
	0xfd, 0x4f, 0x0d, 0x36, 0xd3, 0x89, 0xe8, 0x1a, 0xed, 0xc6, 0xdd, 0xe5, 0x75, 0x3d, 0x1e, 0xb0,
	0x6e, 0x00, 0x36, 0xc7, 0x03, 0x1b, 0x77, 0x97, 0xde, 0xf5, 0xf8, 0x34, 0xcc, 0x03, 0xbf, 0x52,
	0x67, 0xb6, 0xfc, 0xc6, 0x24, 0xb9, 0x8f, 0xc2, 0x4d, 0xa4, 0x15, 0x45, 0xf7, 0xd6, 0xb8, 0x2f,
	0xdf, 0xe4, 0xd3, 0xb6, 0x93, 0xfd, 0x14, 0x4d, 0x58, 0xcd, 0x31, 0xac, 0xff, 0x9f, 0x1a, 0xab,
	0x39, 0x96, 0xd5, 0xcc, 0xb1, 0x1e, 0xcf, 0xe1, 0xbf, 0x45, 0x9f, 0xfc, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x73, 0x4c, 0x64, 0xb6, 0x2f, 0x15, 0x00, 0x00,
}