// Code generated by protoc-gen-go.
// source: ipv6_rib_edm_route.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_subscriber_non_as_protocol_routes_protocol_route is a generated protocol buffer package.

It is generated from these files:
	ipv6_rib_edm_route.proto

It has these top-level messages:
	Ipv6RibEdmRoute_KEYS
	Ipv6RibEdmRoute
	Ipv6RibEdmAddr
	Ipv6RibEdmPath
	Ipv6RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_subscriber_non_as_protocol_routes_protocol_route

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

// Information of a rib route head and rib proto route
type Ipv6RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *Ipv6RibEdmRoute_KEYS) Reset()                    { *m = Ipv6RibEdmRoute_KEYS{} }
func (m *Ipv6RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv6RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv6RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath *Ipv6RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv6RibEdmRoute) Reset()                    { *m = Ipv6RibEdmRoute{} }
func (m *Ipv6RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute) ProtoMessage()               {}
func (*Ipv6RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePath() *Ipv6RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

type Ipv6RibEdmAddr struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Ipv6RibEdmAddr) Reset()                    { *m = Ipv6RibEdmAddr{} }
func (m *Ipv6RibEdmAddr) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAddr) ProtoMessage()               {}
func (*Ipv6RibEdmAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6RibEdmAddr) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Information of a rib path
type Ipv6RibEdmPath struct {
	// Next path
	Ipv6RibEdmPath []*Ipv6RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv6_rib_edm_path,json=ipv6RibEdmPath" json:"ipv6_rib_edm_path,omitempty"`
}

func (m *Ipv6RibEdmPath) Reset()                    { *m = Ipv6RibEdmPath{} }
func (m *Ipv6RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPath) ProtoMessage()               {}
func (*Ipv6RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv6RibEdmPath) GetIpv6RibEdmPath() []*Ipv6RibEdmPathItem {
	if m != nil {
		return m.Ipv6RibEdmPath
	}
	return nil
}

type Ipv6RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop" json:"v6_nexthop,omitempty"`
	// V6 Infosource
	V6InformationSource string `protobuf:"bytes,4,opt,name=v6_information_source,json=v6InformationSource" json:"v6_information_source,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,5,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,6,opt,name=metric" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,7,opt,name=load_metric,json=loadMetric" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,8,opt,name=flags64" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,9,opt,name=flags" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,10,opt,name=private_flags,json=privateFlags" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,11,opt,name=looped" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,12,opt,name=next_hop_table_id,json=nextHopTableId" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,13,opt,name=next_hop_vrf_name,json=nextHopVrfName" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,14,opt,name=next_hop_table_name,json=nextHopTableName" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,15,opt,name=next_hop_afi,json=nextHopAfi" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,16,opt,name=next_hop_safi,json=nextHopSafi" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,17,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,18,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,19,opt,name=pathid" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,20,opt,name=backup_pathid,json=backupPathid" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,21,opt,name=ref_cnt_of_backup,json=refCntOfBackup" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,22,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,23,opt,name=mvpn_present,json=mvpnPresent" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,24,opt,name=pathrt_present,json=pathrtPresent" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,25,opt,name=vrfimportrt_present,json=vrfimportrtPresent" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,26,opt,name=sourceasrt_present,json=sourceasrtPresent" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,27,opt,name=sourcerd_present,json=sourcerdPresent" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,28,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,29,opt,name=next_hop_id,json=nextHopId" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,30,opt,name=next_hop_id_refcount,json=nextHopIdRefcount" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,31,opt,name=ospf_area_id,json=ospfAreaId" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr []*Ipv6RibEdmAddr `protobuf:"bytes,32,rep,name=remote_backup_addr,json=remoteBackupAddr" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,33,opt,name=has_labelstk,json=hasLabelstk" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,34,opt,name=num_labels,json=numLabels" json:"num_labels,omitempty"`
	// Outgoing label stack for this path
	Labelstk []uint32 `protobuf:"varint,35,rep,packed,name=labelstk" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,36,opt,name=binding_label,json=bindingLabel" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,37,opt,name=nhid_feid,json=nhidFeid" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid uint64 `protobuf:"varint,38,opt,name=mpls_feid,json=mplsFeid" json:"mpls_feid,omitempty"`
}

func (m *Ipv6RibEdmPathItem) Reset()                    { *m = Ipv6RibEdmPathItem{} }
func (m *Ipv6RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPathItem) ProtoMessage()               {}
func (*Ipv6RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Ipv6RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetV6InformationSource() string {
	if m != nil {
		return m.V6InformationSource
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetRemoteBackupAddr() []*Ipv6RibEdmAddr {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored" json:"mirrored,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,5,opt,name=redist_only,json=redistOnly" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,6,opt,name=label" json:"label,omitempty"`
	// Distance
	Distance uint32 `protobuf:"varint,7,opt,name=distance" json:"distance,omitempty"`
}

func (m *RibEdmLocalLabel) Reset()                    { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string            { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()               {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.subscriber.non_as.protocol_routes.protocol_route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.subscriber.non_as.protocol_routes.protocol_route.ipv6_rib_edm_route")
	proto.RegisterType((*Ipv6RibEdmAddr)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.subscriber.non_as.protocol_routes.protocol_route.ipv6_rib_edm_addr")
	proto.RegisterType((*Ipv6RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.subscriber.non_as.protocol_routes.protocol_route.ipv6_rib_edm_path")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.subscriber.non_as.protocol_routes.protocol_route.ipv6_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.subscriber.non_as.protocol_routes.protocol_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xcd, 0x72, 0x1b, 0xb9,
	0x11, 0xae, 0xb1, 0xd6, 0x94, 0x04, 0x89, 0xb2, 0x38, 0xd2, 0x4a, 0xd0, 0x6a, 0xd7, 0x4b, 0x6b,
	0xe3, 0x94, 0x94, 0x2a, 0x71, 0x53, 0xb2, 0xa3, 0x24, 0xce, 0xaf, 0x2c, 0xcb, 0x36, 0x63, 0xd9,
	0x52, 0x68, 0x97, 0xab, 0x72, 0x42, 0x81, 0x03, 0x0c, 0x89, 0xf2, 0xcc, 0x60, 0x0c, 0x80, 0x63,
	0xe9, 0x9a, 0x63, 0x9e, 0x26, 0xef, 0x90, 0x4b, 0xae, 0xb9, 0xe7, 0x96, 0x73, 0x2e, 0xa9, 0x3c,
	0x40, 0x0a, 0xdd, 0x98, 0x21, 0xf5, 0x93, 0x07, 0xf0, 0x5e, 0x6c, 0xf5, 0xf7, 0x7d, 0xe8, 0x01,
	0xba, 0xd1, 0xdd, 0x20, 0xa1, 0xaa, 0xac, 0x0e, 0x99, 0x51, 0x43, 0x26, 0x45, 0xce, 0x8c, 0x9e,
	0x38, 0xd9, 0x2b, 0x8d, 0x76, 0x3a, 0xfe, 0x6b, 0x94, 0x28, 0x9b, 0x68, 0xa6, 0xb4, 0x65, 0x17,
	0x86, 0xa9, 0x12, 0x54, 0x20, 0xd7, 0xa5, 0x34, 0xbd, 0x66, 0xa1, 0x75, 0x62, 0x78, 0xd9, 0xab,
	0x4c, 0x6a, 0xfd, 0x3f, 0x3d, 0x9e, 0xda, 0x1e, 0x4f, 0x7b, 0xd6, 0xff, 0x6f, 0x79, 0xda, 0x0b,
	0x0b, 0xc1, 0x35, 0x73, 0x7c, 0x98, 0x49, 0x56, 0xf0, 0x5c, 0xda, 0xff, 0x47, 0xe0, 0xe7, 0x13,
	0x9d, 0xf5, 0xec, 0x64, 0x68, 0x13, 0xa3, 0x86, 0xd2, 0xf4, 0x0a, 0x5d, 0x30, 0x6e, 0x1b, 0x0a,
	0x17, 0x5d, 0xb7, 0x77, 0xfe, 0x11, 0x91, 0xcd, 0x9b, 0xe7, 0x61, 0xaf, 0x4e, 0xfe, 0xf4, 0x36,
	0xde, 0x22, 0x0b, 0x95, 0x49, 0xe1, 0x0b, 0x34, 0xea, 0x46, 0xbb, 0x8b, 0x83, 0xf9, 0xca, 0xa4,
	0x6f, 0x78, 0x2e, 0xe3, 0x4d, 0x32, 0xcf, 0x03, 0x73, 0x07, 0x98, 0x16, 0x47, 0x62, 0x8b, 0x2c,
	0xd8, 0x9a, 0x99, 0xc3, 0x35, 0x36, 0x50, 0xbb, 0x64, 0xf5, 0xfa, 0xc6, 0xe9, 0x17, 0x20, 0x59,
	0x01, 0xfc, 0x9d, 0x87, 0x41, 0x49, 0xc9, 0x3c, 0x17, 0xc2, 0x48, 0x6b, 0xe9, 0x5d, 0xf4, 0x11,
	0xcc, 0xf8, 0x3b, 0xd2, 0x2e, 0x8d, 0x4c, 0xd5, 0x05, 0xcb, 0x64, 0x31, 0x72, 0x63, 0xda, 0xea,
	0x46, 0xbb, 0xed, 0xc1, 0x32, 0x82, 0xa7, 0x80, 0xed, 0xfc, 0x99, 0x90, 0xf8, 0xe6, 0x99, 0xe2,
	0x0d, 0xd2, 0x42, 0x19, 0x3d, 0xc0, 0x2d, 0xa3, 0x75, 0xd3, 0xe7, 0xa3, 0x9b, 0x3e, 0xbd, 0x08,
	0x37, 0x5f, 0x49, 0x63, 0x95, 0x2e, 0xe8, 0x63, 0x14, 0x01, 0xf8, 0x1e, 0xb1, 0xf8, 0x5b, 0xb2,
	0xd4, 0x84, 0x57, 0x09, 0xfa, 0x33, 0x90, 0x90, 0x1a, 0xea, 0x0b, 0xfc, 0x54, 0x10, 0xc0, 0xf9,
	0x0f, 0x61, 0x27, 0xcb, 0x35, 0x08, 0xa7, 0xff, 0x8a, 0x2c, 0xa8, 0xc2, 0x3a, 0x5e, 0x24, 0x92,
	0xfe, 0x1c, 0xf8, 0xc6, 0x8e, 0xb7, 0xc9, 0x62, 0x92, 0x29, 0x59, 0x38, 0xef, 0xff, 0x17, 0xe0,
	0x7f, 0x01, 0x81, 0xbe, 0x88, 0xbf, 0x21, 0x24, 0x04, 0xf8, 0xb2, 0x94, 0xf4, 0x97, 0xc0, 0x2e,
	0x62, 0x68, 0x2f, 0x4b, 0xf0, 0x5b, 0x1a, 0xa5, 0x8d, 0x72, 0x97, 0xf4, 0x09, 0x2e, 0xad, 0x6d,
	0x48, 0x5b, 0x25, 0x70, 0xe1, 0xaf, 0x80, 0x9b, 0xb7, 0x95, 0x80, 0x65, 0xeb, 0xe4, 0x6e, 0x9a,
	0xf1, 0x91, 0xa5, 0xbf, 0x06, 0x1c, 0x8d, 0xf8, 0x21, 0x59, 0x91, 0x17, 0x4e, 0x16, 0x42, 0x0a,
	0x86, 0xf4, 0x6f, 0xba, 0xd1, 0xee, 0x17, 0x83, 0x76, 0x8d, 0x3e, 0x07, 0xd9, 0x2a, 0x99, 0x73,
	0x7c, 0x44, 0x7f, 0x0b, 0x4b, 0xfd, 0x9f, 0x7e, 0x17, 0x42, 0x85, 0xd3, 0xfd, 0x0e, 0x77, 0x51,
	0xdb, 0xf1, 0x3e, 0x89, 0x85, 0x0a, 0x01, 0x66, 0x8d, 0xea, 0xf7, 0xa0, 0xea, 0x34, 0xcc, 0xb3,
	0x5a, 0xbe, 0x41, 0x5a, 0xb9, 0x74, 0x46, 0x25, 0xf4, 0x08, 0x24, 0xc1, 0x82, 0x34, 0x70, 0x37,
	0xb6, 0x2c, 0xd1, 0x93, 0xc2, 0xd1, 0xa7, 0x21, 0x0d, 0x1e, 0x3a, 0xf6, 0x88, 0xff, 0x0e, 0x77,
	0xce, 0xa8, 0xa1, 0x0f, 0x96, 0x12, 0xb2, 0x70, 0x3e, 0x26, 0xc7, 0xf8, 0x9d, 0x86, 0xe9, 0x07,
	0xc2, 0x67, 0xcd, 0x19, 0x9e, 0xa6, 0x2a, 0x61, 0xaa, 0x10, 0xf2, 0x82, 0x3e, 0xc3, 0xdc, 0x07,
	0xb0, 0xef, 0xb1, 0x78, 0xaf, 0xbe, 0xdd, 0xa5, 0x91, 0x89, 0x14, 0xd2, 0xef, 0xfc, 0x04, 0x74,
	0xf7, 0x00, 0x3f, 0x6f, 0x60, 0x9f, 0xc4, 0x8f, 0xda, 0xb2, 0x91, 0xd1, 0x93, 0x92, 0x3e, 0xc7,
	0x18, 0x7c, 0xd4, 0xf6, 0x85, 0xb7, 0x7d, 0x26, 0xd2, 0x4c, 0x7f, 0x62, 0x3e, 0x6c, 0x2f, 0x30,
	0x13, 0xde, 0x7e, 0xc7, 0x47, 0x7e, 0x5d, 0xfa, 0x49, 0xb0, 0x24, 0xe3, 0xd6, 0xd2, 0x97, 0xb8,
	0x2e, 0xfd, 0x24, 0x8e, 0xbd, 0xed, 0xc9, 0x52, 0x25, 0xe1, 0xc8, 0xfd, 0x90, 0x5e, 0x95, 0xe0,
	0x81, 0x37, 0x48, 0x8b, 0x27, 0x4e, 0x55, 0x92, 0xfe, 0xa1, 0x1b, 0xed, 0x2e, 0x0c, 0x82, 0x15,
	0x7f, 0x4d, 0x16, 0x9b, 0xb0, 0xd2, 0x57, 0x40, 0x4d, 0x81, 0xf8, 0xa7, 0x64, 0x7d, 0x9a, 0x0e,
	0xb8, 0xa2, 0x78, 0x69, 0x4f, 0xe1, 0x52, 0x4e, 0x53, 0x75, 0xee, 0x29, 0xb8, 0xba, 0xdb, 0x04,
	0xef, 0x1b, 0xe3, 0x23, 0x49, 0x5f, 0xe3, 0x26, 0x00, 0x38, 0x1a, 0x49, 0x9f, 0x16, 0x24, 0x33,
	0x3e, 0x94, 0x19, 0x7d, 0x83, 0x69, 0x01, 0xe8, 0xd4, 0x23, 0xbe, 0xec, 0xeb, 0xbd, 0x9c, 0xe1,
	0xc9, 0xab, 0x69, 0x61, 0xb9, 0x61, 0xd6, 0xd4, 0xde, 0x39, 0x5c, 0x35, 0xe2, 0x86, 0x59, 0x5d,
	0x79, 0x3f, 0x21, 0x1d, 0xf4, 0x9d, 0x6b, 0xa1, 0xd2, 0x4b, 0xe6, 0x54, 0x2e, 0xe9, 0x1f, 0x41,
	0x86, 0xe1, 0x7f, 0x0d, 0xf8, 0x3b, 0x95, 0xcb, 0xf8, 0x5f, 0x51, 0x5d, 0x27, 0xfe, 0x4a, 0xd0,
	0x41, 0x37, 0xda, 0x5d, 0x3a, 0xf8, 0x5b, 0xd4, 0xfb, 0xdc, 0x7a, 0x77, 0xef, 0x4a, 0x8f, 0xf3,
	0x67, 0x09, 0xd5, 0x7e, 0xce, 0xdd, 0x78, 0x67, 0x8f, 0x74, 0xae, 0xf0, 0xbe, 0x83, 0xfa, 0x5a,
	0xae, 0x78, 0x36, 0xa9, 0xdb, 0x39, 0x1a, 0x3b, 0x7f, 0xb9, 0x73, 0x4d, 0xeb, 0x7d, 0xc5, 0xff,
	0x8d, 0x6e, 0x41, 0x69, 0xd4, 0x9d, 0xdb, 0x5d, 0x3a, 0xf8, 0xfb, 0x0f, 0x21, 0x5a, 0x4c, 0x39,
	0x99, 0x0f, 0x56, 0x3c, 0x3e, 0x50, 0xc3, 0x13, 0x91, 0x43, 0xdc, 0xfe, 0xbd, 0x4c, 0x36, 0x6e,
	0x97, 0xce, 0x8e, 0xa5, 0xe8, 0xea, 0x58, 0xda, 0x27, 0xb1, 0x2a, 0x52, 0x6d, 0x72, 0xee, 0x7c,
	0xad, 0x58, 0x3d, 0x31, 0x49, 0x3d, 0x19, 0x3b, 0x33, 0xcc, 0x5b, 0x20, 0x7c, 0xa3, 0xae, 0x0e,
	0x59, 0x21, 0x2f, 0xdc, 0x58, 0x97, 0x61, 0x4c, 0x2e, 0x56, 0x87, 0x6f, 0x10, 0x88, 0x0f, 0xc8,
	0x97, 0xd5, 0x21, 0xbb, 0xc5, 0x21, 0x4e, 0xcb, 0xb5, 0xea, 0xb0, 0x7f, 0xc3, 0xe5, 0x43, 0xb2,
	0xa2, 0x0a, 0x27, 0x4d, 0xca, 0x93, 0x30, 0x5a, 0x71, 0x72, 0xb6, 0x1b, 0x14, 0x0a, 0x74, 0xda,
	0x32, 0x5b, 0xd7, 0x5b, 0x66, 0xa6, 0xb9, 0x60, 0x81, 0x9c, 0xc7, 0xda, 0xf4, 0xd0, 0x6b, 0x14,
	0x50, 0x32, 0x0f, 0x6d, 0xfe, 0xf0, 0x31, 0x5d, 0x80, 0xb2, 0xaa, 0xcd, 0xe9, 0x7c, 0x58, 0x9c,
	0x9d, 0x0f, 0x30, 0xe9, 0x54, 0xc5, 0x9d, 0x0c, 0xe3, 0x81, 0xd4, 0x43, 0x15, 0x40, 0x9c, 0x0e,
	0x1b, 0xa4, 0x95, 0x69, 0x5d, 0x4a, 0x41, 0x97, 0xb0, 0x2d, 0xa1, 0x15, 0xef, 0x91, 0x8e, 0x0f,
	0x0e, 0x1b, 0xeb, 0x32, 0xe4, 0x5e, 0x09, 0xba, 0x0c, 0x0e, 0x56, 0x3c, 0xf1, 0x52, 0x97, 0xf0,
	0x58, 0xe8, 0x5f, 0x95, 0x36, 0x8f, 0x95, 0x36, 0xbe, 0x2a, 0x82, 0xf4, 0x7d, 0x78, 0xb3, 0xec,
	0x93, 0xb5, 0x6b, 0x5e, 0x41, 0xbc, 0x02, 0xe2, 0xd5, 0x59, 0xbf, 0x20, 0xef, 0x92, 0xe5, 0x46,
	0xce, 0x53, 0x45, 0xef, 0x61, 0x4c, 0x82, 0xee, 0x28, 0x55, 0xf1, 0x0e, 0x69, 0x37, 0x0a, 0xeb,
	0x25, 0xab, 0x20, 0x59, 0x0a, 0x92, 0xb7, 0x3c, 0x55, 0xd7, 0x9b, 0x5e, 0xe7, 0x46, 0xd3, 0xdb,
	0x26, 0x8b, 0x6e, 0x52, 0x14, 0x12, 0x5e, 0x0c, 0x31, 0xb6, 0x4c, 0x04, 0xfa, 0x02, 0x9e, 0x2c,
	0xdc, 0x8d, 0x95, 0xa0, 0x6b, 0x98, 0x2e, 0xb4, 0x7c, 0x74, 0x87, 0x3c, 0xf9, 0x30, 0x29, 0x59,
	0xa0, 0xd7, 0x31, 0xba, 0x08, 0x9e, 0xa3, 0x68, 0x8f, 0x74, 0x8c, 0x4c, 0x59, 0x52, 0x38, 0xa6,
	0x53, 0x86, 0x14, 0xfd, 0x12, 0xa3, 0x68, 0x64, 0x7a, 0x5c, 0xb8, 0xb3, 0xf4, 0x29, 0xa0, 0xf1,
	0x31, 0xb9, 0x5f, 0x4c, 0xf2, 0xa1, 0x34, 0x5e, 0xd9, 0xcc, 0xf5, 0x44, 0xe7, 0xf9, 0xa4, 0x50,
	0x4e, 0x49, 0x4b, 0x37, 0x60, 0xdd, 0x36, 0xaa, 0xce, 0xd2, 0x93, 0xa0, 0x39, 0x9e, 0x4a, 0xe2,
	0x07, 0x64, 0x39, 0xaf, 0x4a, 0x3f, 0x29, 0xa4, 0x95, 0x85, 0xa3, 0x9b, 0x90, 0xd3, 0x25, 0x8f,
	0x9d, 0x23, 0xe4, 0x6f, 0xa9, 0xdf, 0xb0, 0x71, 0x8d, 0x88, 0x82, 0xa8, 0x8d, 0x68, 0x2d, 0xfb,
	0x9e, 0xac, 0x55, 0x26, 0x55, 0x79, 0xa9, 0x8d, 0x9b, 0xd1, 0x6e, 0x81, 0x36, 0x9e, 0xa1, 0xea,
	0x05, 0xfb, 0x24, 0xc6, 0x12, 0xe1, 0x76, 0x46, 0xff, 0x15, 0xe8, 0x3b, 0x53, 0xa6, 0x96, 0xef,
	0x91, 0x55, 0x04, 0x8d, 0x68, 0xc4, 0xdb, 0x20, 0xbe, 0x57, 0xe3, 0xb5, 0xf4, 0x09, 0xd9, 0xb2,
	0x72, 0x94, 0xcb, 0xc2, 0x49, 0x51, 0x57, 0x6c, 0xb3, 0xe6, 0x6b, 0x58, 0xb3, 0xd9, 0x08, 0x42,
	0x01, 0xd7, 0x6b, 0xef, 0x93, 0xa5, 0xe6, 0x7e, 0x28, 0x41, 0xbf, 0xc1, 0x07, 0x59, 0xb8, 0x1d,
	0x7d, 0x11, 0x7f, 0x4f, 0xd6, 0x67, 0x78, 0x66, 0x64, 0x8a, 0xd3, 0xfb, 0x3e, 0x3e, 0x44, 0x1a,
	0xe1, 0x20, 0x10, 0xfe, 0x4a, 0x6a, 0x5b, 0xa6, 0x8c, 0x1b, 0xc9, 0xbd, 0xc7, 0x6f, 0xe1, 0xea,
	0x12, 0x8f, 0x1d, 0x19, 0xc9, 0xfb, 0x22, 0xfe, 0x4f, 0x44, 0x62, 0x23, 0x73, 0xed, 0x64, 0x48,
	0x38, 0xf4, 0x7d, 0xda, 0x85, 0xae, 0xfd, 0xd9, 0xcf, 0x38, 0x7f, 0x96, 0xc1, 0x2a, 0x9e, 0x0f,
	0x6f, 0xee, 0x91, 0x9f, 0x6a, 0x0f, 0xc8, 0xf2, 0x98, 0x5b, 0xac, 0x30, 0xeb, 0x3e, 0xd0, 0x07,
	0x78, 0xf1, 0xc6, 0xdc, 0x9e, 0x06, 0xc8, 0x77, 0xdc, 0x62, 0x92, 0x07, 0x09, 0xdd, 0x09, 0x99,
	0x98, 0xe4, 0x28, 0xf0, 0x8f, 0xd2, 0x66, 0xf5, 0x77, 0xdd, 0x39, 0x5f, 0x83, 0xb5, 0x0d, 0xb5,
	0xa6, 0x0a, 0xa1, 0x8a, 0x51, 0xa8, 0xe1, 0x1f, 0x85, 0x5a, 0x43, 0xb0, 0xa9, 0xe2, 0x62, 0xac,
	0x04, 0x4b, 0xa5, 0x12, 0xf4, 0x21, 0x34, 0xc8, 0x05, 0x0f, 0x3c, 0x97, 0x4a, 0x78, 0x32, 0x2f,
	0x33, 0x8b, 0xe4, 0x8f, 0x91, 0xf4, 0x80, 0x27, 0x77, 0xfe, 0x19, 0x91, 0xb5, 0xfa, 0x7c, 0x99,
	0x4e, 0x78, 0x86, 0x5f, 0xb9, 0xf9, 0x53, 0x21, 0xba, 0xe5, 0xa7, 0xc2, 0x95, 0x9f, 0x03, 0x77,
	0xae, 0xfd, 0x1c, 0x58, 0x27, 0x77, 0xad, 0xe3, 0x19, 0xfe, 0x0e, 0x6b, 0x0f, 0xd0, 0xf0, 0x47,
	0xcd, 0x95, 0x31, 0xda, 0x48, 0x01, 0xf3, 0xa4, 0x3d, 0x68, 0x6c, 0x68, 0x56, 0xd2, 0xbf, 0xbb,
	0x99, 0x2e, 0xb2, 0x4b, 0x98, 0x20, 0xbe, 0x59, 0x01, 0x74, 0x56, 0x64, 0x97, 0xde, 0x25, 0xc6,
	0x00, 0xa7, 0x07, 0x1a, 0x57, 0x9e, 0xf4, 0xf3, 0x57, 0x9f, 0xf4, 0xc3, 0x16, 0xec, 0xf7, 0xd1,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xd5, 0x93, 0x21, 0x35, 0x0f, 0x00, 0x00,
}
