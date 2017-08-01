// Code generated by protoc-gen-go.
// source: bgp_updgen_subgrp_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_update_generation_sub_groups_update_generation_sub_group is a generated protocol buffer package.

It is generated from these files:
	bgp_updgen_subgrp_bag.proto

It has these top-level messages:
	BgpUpdgenSubgrpBag_KEYS
	BgpUpdgenSubgrpBag
	BgpTimespec
	BgpUpdgenStatsBag
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_update_generation_sub_groups_update_generation_sub_group

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

// BGP Update generation Sub-group information
type BgpUpdgenSubgrpBag_KEYS struct {
	InstanceName     string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName           string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	UpdateGroupIndex uint32 `protobuf:"varint,3,opt,name=update_group_index,json=updateGroupIndex" json:"update_group_index,omitempty"`
	SubGroupIndex    uint32 `protobuf:"varint,4,opt,name=sub_group_index,json=subGroupIndex" json:"sub_group_index,omitempty"`
	SubGroupId       uint32 `protobuf:"varint,5,opt,name=sub_group_id,json=subGroupId" json:"sub_group_id,omitempty"`
}

func (m *BgpUpdgenSubgrpBag_KEYS) Reset()                    { *m = BgpUpdgenSubgrpBag_KEYS{} }
func (m *BgpUpdgenSubgrpBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdgenSubgrpBag_KEYS) ProtoMessage()               {}
func (*BgpUpdgenSubgrpBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpUpdgenSubgrpBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetUpdateGroupIndex() uint32 {
	if m != nil {
		return m.UpdateGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetSubGroupIndex() uint32 {
	if m != nil {
		return m.SubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

type BgpUpdgenSubgrpBag struct {
	// ProcessID
	ProcessId uint32 `protobuf:"varint,50,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	// VRF Name
	UpdateVrfName string `protobuf:"bytes,51,opt,name=update_vrf_name,json=updateVrfName" json:"update_vrf_name,omitempty"`
	// Address family identifier
	UpdateGroupAfName string `protobuf:"bytes,52,opt,name=update_group_af_name,json=updateGroupAfName" json:"update_group_af_name,omitempty"`
	// Sub-group index
	SubGroupIndex uint32 `protobuf:"varint,53,opt,name=sub_group_index,json=subGroupIndex" json:"sub_group_index,omitempty"`
	// Sub-group ID
	SubGroupId uint32 `protobuf:"varint,54,opt,name=sub_group_id,json=subGroupId" json:"sub_group_id,omitempty"`
	// Parent Sub-group index
	ParentSubGroupIndex uint32 `protobuf:"varint,55,opt,name=parent_sub_group_index,json=parentSubGroupIndex" json:"parent_sub_group_index,omitempty"`
	// Parent Sub-group identifier
	ParentSubGroupId uint32 `protobuf:"varint,56,opt,name=parent_sub_group_id,json=parentSubGroupId" json:"parent_sub_group_id,omitempty"`
	// Update-group index
	UpdateGroupIndex uint32 `protobuf:"varint,57,opt,name=update_group_index,json=updateGroupIndex" json:"update_group_index,omitempty"`
	// Main table version
	UpdateMainTableVersion uint32 `protobuf:"varint,58,opt,name=update_main_table_version,json=updateMainTableVersion" json:"update_main_table_version,omitempty"`
	// VRF Table RIB version
	UpdateVrfTableRibVersion uint32 `protobuf:"varint,59,opt,name=update_vrf_table_rib_version,json=updateVrfTableRibVersion" json:"update_vrf_table_rib_version,omitempty"`
	// Current update limit
	CurrentUpdateLimitSubGroup uint32 `protobuf:"varint,60,opt,name=current_update_limit_sub_group,json=currentUpdateLimitSubGroup" json:"current_update_limit_sub_group,omitempty"`
	// Configured update limit
	ConfiguredUpdateLimitSubGroup uint32 `protobuf:"varint,61,opt,name=configured_update_limit_sub_group,json=configuredUpdateLimitSubGroup" json:"configured_update_limit_sub_group,omitempty"`
	// OutQueue messages
	UpdateOutQueueMessages uint32 `protobuf:"varint,62,opt,name=update_out_queue_messages,json=updateOutQueueMessages" json:"update_out_queue_messages,omitempty"`
	// OutQueue size
	UpdateOutQueueSize uint32 `protobuf:"varint,63,opt,name=update_out_queue_size,json=updateOutQueueSize" json:"update_out_queue_size,omitempty"`
	// Is update throttled
	UpdateThrottled bool `protobuf:"varint,64,opt,name=update_throttled,json=updateThrottled" json:"update_throttled,omitempty"`
	// Refresh sub-group count
	RefreshSubGroupCount uint32 `protobuf:"varint,65,opt,name=refresh_sub_group_count,json=refreshSubGroupCount" json:"refresh_sub_group_count,omitempty"`
	// Filter-group count
	FilterGroupCount uint32 `protobuf:"varint,66,opt,name=filter_group_count,json=filterGroupCount" json:"filter_group_count,omitempty"`
	// Neighbor count
	NeighborCount uint32 `protobuf:"varint,67,opt,name=neighbor_count,json=neighborCount" json:"neighbor_count,omitempty"`
	// Version
	Version uint32 `protobuf:"varint,68,opt,name=version" json:"version,omitempty"`
	// NSR version
	NsrVersion uint32 `protobuf:"varint,69,opt,name=nsr_version,json=nsrVersion" json:"nsr_version,omitempty"`
	// Pending target version
	PendingTargetVersion uint32 `protobuf:"varint,70,opt,name=pending_target_version,json=pendingTargetVersion" json:"pending_target_version,omitempty"`
	// Next resume version
	NextResumeVersion uint32 `protobuf:"varint,71,opt,name=next_resume_version,json=nextResumeVersion" json:"next_resume_version,omitempty"`
	// Refresh target version
	UpdateRefreshTargetVersion uint32 `protobuf:"varint,72,opt,name=update_refresh_target_version,json=updateRefreshTargetVersion" json:"update_refresh_target_version,omitempty"`
	// Is IO write event pending
	IoWriteEventPending bool `protobuf:"varint,73,opt,name=io_write_event_pending,json=ioWriteEventPending" json:"io_write_event_pending,omitempty"`
	// Is default route sent
	UpdateDefaultRouteSent bool `protobuf:"varint,74,opt,name=update_default_route_sent,json=updateDefaultRouteSent" json:"update_default_route_sent,omitempty"`
	// Creation time
	CreationTimestamp *BgpTimespec `protobuf:"bytes,75,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	// Merge count
	MergeCount uint32 `protobuf:"varint,76,opt,name=merge_count,json=mergeCount" json:"merge_count,omitempty"`
	// Last merge time
	LastMergeTimestamp *BgpTimespec `protobuf:"bytes,77,opt,name=last_merge_timestamp,json=lastMergeTimestamp" json:"last_merge_timestamp,omitempty"`
	// Last merged sub-group index
	LastMergedSubGroupIndex uint32 `protobuf:"varint,78,opt,name=last_merged_sub_group_index,json=lastMergedSubGroupIndex" json:"last_merged_sub_group_index,omitempty"`
	// Sending EoR has been attempted
	EoRAttempted bool `protobuf:"varint,79,opt,name=eo_r_attempted,json=eoRAttempted" json:"eo_r_attempted,omitempty"`
	// First update walk start time
	FirstUpdateWalkStartTimestamp *BgpTimespec `protobuf:"bytes,80,opt,name=first_update_walk_start_timestamp,json=firstUpdateWalkStartTimestamp" json:"first_update_walk_start_timestamp,omitempty"`
	// First update walk end time
	FirstUpdateWalkEndTimestamp *BgpTimespec `protobuf:"bytes,81,opt,name=first_update_walk_end_timestamp,json=firstUpdateWalkEndTimestamp" json:"first_update_walk_end_timestamp,omitempty"`
	// Last update walk start time
	LastUpdateWalkStartTimestamp *BgpTimespec `protobuf:"bytes,82,opt,name=last_update_walk_start_timestamp,json=lastUpdateWalkStartTimestamp" json:"last_update_walk_start_timestamp,omitempty"`
	// Last update walk end time
	LastUpdateWalkEndTimestamp *BgpTimespec `protobuf:"bytes,83,opt,name=last_update_walk_end_timestamp,json=lastUpdateWalkEndTimestamp" json:"last_update_walk_end_timestamp,omitempty"`
	// Time since last update walk end event
	LastUpdateWalkEndAge uint32 `protobuf:"varint,84,opt,name=last_update_walk_end_age,json=lastUpdateWalkEndAge" json:"last_update_walk_end_age,omitempty"`
	// Last update queued time
	LastUpdateQueuedTimestamp *BgpTimespec `protobuf:"bytes,85,opt,name=last_update_queued_timestamp,json=lastUpdateQueuedTimestamp" json:"last_update_queued_timestamp,omitempty"`
	// Time since last update queued event (in seconds)
	LastUpdateQueuedAge uint32 `protobuf:"varint,86,opt,name=last_update_queued_age,json=lastUpdateQueuedAge" json:"last_update_queued_age,omitempty"`
	// Update statistics
	UpdateStatistics *BgpUpdgenStatsBag `protobuf:"bytes,87,opt,name=update_statistics,json=updateStatistics" json:"update_statistics,omitempty"`
}

func (m *BgpUpdgenSubgrpBag) Reset()                    { *m = BgpUpdgenSubgrpBag{} }
func (m *BgpUpdgenSubgrpBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdgenSubgrpBag) ProtoMessage()               {}
func (*BgpUpdgenSubgrpBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpUpdgenSubgrpBag) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateVrfName() string {
	if m != nil {
		return m.UpdateVrfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag) GetUpdateGroupAfName() string {
	if m != nil {
		return m.UpdateGroupAfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag) GetSubGroupIndex() uint32 {
	if m != nil {
		return m.SubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetParentSubGroupIndex() uint32 {
	if m != nil {
		return m.ParentSubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetParentSubGroupId() uint32 {
	if m != nil {
		return m.ParentSubGroupId
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateGroupIndex() uint32 {
	if m != nil {
		return m.UpdateGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateMainTableVersion() uint32 {
	if m != nil {
		return m.UpdateMainTableVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateVrfTableRibVersion() uint32 {
	if m != nil {
		return m.UpdateVrfTableRibVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetCurrentUpdateLimitSubGroup() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitSubGroup
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetConfiguredUpdateLimitSubGroup() uint32 {
	if m != nil {
		return m.ConfiguredUpdateLimitSubGroup
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateOutQueueMessages() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessages
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateOutQueueSize() uint32 {
	if m != nil {
		return m.UpdateOutQueueSize
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateThrottled() bool {
	if m != nil {
		return m.UpdateThrottled
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetRefreshSubGroupCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetFilterGroupCount() uint32 {
	if m != nil {
		return m.FilterGroupCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetNeighborCount() uint32 {
	if m != nil {
		return m.NeighborCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetNsrVersion() uint32 {
	if m != nil {
		return m.NsrVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetPendingTargetVersion() uint32 {
	if m != nil {
		return m.PendingTargetVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetNextResumeVersion() uint32 {
	if m != nil {
		return m.NextResumeVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateRefreshTargetVersion() uint32 {
	if m != nil {
		return m.UpdateRefreshTargetVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetIoWriteEventPending() bool {
	if m != nil {
		return m.IoWriteEventPending
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetUpdateDefaultRouteSent() bool {
	if m != nil {
		return m.UpdateDefaultRouteSent
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetCreationTimestamp() *BgpTimespec {
	if m != nil {
		return m.CreationTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetMergeCount() uint32 {
	if m != nil {
		return m.MergeCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetLastMergeTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastMergeTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastMergedSubGroupIndex() uint32 {
	if m != nil {
		return m.LastMergedSubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetEoRAttempted() bool {
	if m != nil {
		return m.EoRAttempted
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetFirstUpdateWalkStartTimestamp() *BgpTimespec {
	if m != nil {
		return m.FirstUpdateWalkStartTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetFirstUpdateWalkEndTimestamp() *BgpTimespec {
	if m != nil {
		return m.FirstUpdateWalkEndTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateWalkStartTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateWalkStartTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateWalkEndTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateWalkEndTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateWalkEndAge() uint32 {
	if m != nil {
		return m.LastUpdateWalkEndAge
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateQueuedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateQueuedTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateQueuedAge() uint32 {
	if m != nil {
		return m.LastUpdateQueuedAge
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateStatistics() *BgpUpdgenStatsBag {
	if m != nil {
		return m.UpdateStatistics
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

// BGP Update generation common statistics information
type BgpUpdgenStatsBag struct {
	// OutQueue High Messages
	UpdateOutQueueMessagesHigh uint32 `protobuf:"varint,1,opt,name=update_out_queue_messages_high,json=updateOutQueueMessagesHigh" json:"update_out_queue_messages_high,omitempty"`
	// OutQueue Cumulative Messages
	UpdateOutQueueMessagesCumulative uint32 `protobuf:"varint,2,opt,name=update_out_queue_messages_cumulative,json=updateOutQueueMessagesCumulative" json:"update_out_queue_messages_cumulative,omitempty"`
	// OutQueue Discarded Messages
	UpdateOutQueueMessagesDiscarded uint32 `protobuf:"varint,3,opt,name=update_out_queue_messages_discarded,json=updateOutQueueMessagesDiscarded" json:"update_out_queue_messages_discarded,omitempty"`
	// OutQueue Cleared Messages
	UpdateOutQueueMessagesCleared uint32 `protobuf:"varint,4,opt,name=update_out_queue_messages_cleared,json=updateOutQueueMessagesCleared" json:"update_out_queue_messages_cleared,omitempty"`
	// OutQueue Hi Size
	UpdateOutQueueSizeHigh uint32 `protobuf:"varint,5,opt,name=update_out_queue_size_high,json=updateOutQueueSizeHigh" json:"update_out_queue_size_high,omitempty"`
	// OutQueue Cumulative Size
	UpdateOutQueueSizeCumulative uint64 `protobuf:"varint,6,opt,name=update_out_queue_size_cumulative,json=updateOutQueueSizeCumulative" json:"update_out_queue_size_cumulative,omitempty"`
	// OutQueue Discarded Size
	UpdateOutQueueSizeDiscarded uint64 `protobuf:"varint,7,opt,name=update_out_queue_size_discarded,json=updateOutQueueSizeDiscarded" json:"update_out_queue_size_discarded,omitempty"`
	// OutQueue Cleared Size
	UpdateOutQueueSizeCleared uint64 `protobuf:"varint,8,opt,name=update_out_queue_size_cleared,json=updateOutQueueSizeCleared" json:"update_out_queue_size_cleared,omitempty"`
	// Last Discarded time
	LastUpdateDiscardTimestamp *BgpTimespec `protobuf:"bytes,9,opt,name=last_update_discard_timestamp,json=lastUpdateDiscardTimestamp" json:"last_update_discard_timestamp,omitempty"`
	// Time since last Discard event (in seconds)
	LastUpdateDiscardAge uint32 `protobuf:"varint,10,opt,name=last_update_discard_age,json=lastUpdateDiscardAge" json:"last_update_discard_age,omitempty"`
	// Last Cleared time
	LastUpdateClearedTimestamp *BgpTimespec `protobuf:"bytes,11,opt,name=last_update_cleared_timestamp,json=lastUpdateClearedTimestamp" json:"last_update_cleared_timestamp,omitempty"`
	// Time since last Clear event (in seconds)
	LastUpdateCleardAge uint32 `protobuf:"varint,12,opt,name=last_update_cleard_age,json=lastUpdateCleardAge" json:"last_update_cleard_age,omitempty"`
	// Throttle Count
	UpdateThrottleCount uint32 `protobuf:"varint,13,opt,name=update_throttle_count,json=updateThrottleCount" json:"update_throttle_count,omitempty"`
	// Last Throttled time
	LastUpdateThrottleTimestamp *BgpTimespec `protobuf:"bytes,14,opt,name=last_update_throttle_timestamp,json=lastUpdateThrottleTimestamp" json:"last_update_throttle_timestamp,omitempty"`
	// Time since last Throttle event (in seconds)
	LastUpdateThrottleAge uint32 `protobuf:"varint,15,opt,name=last_update_throttle_age,json=lastUpdateThrottleAge" json:"last_update_throttle_age,omitempty"`
	// Recovery Count
	UpdateRecoveryCount uint32 `protobuf:"varint,16,opt,name=update_recovery_count,json=updateRecoveryCount" json:"update_recovery_count,omitempty"`
	// Last Recovery time
	LastUpdateRecoveryTimestamp *BgpTimespec `protobuf:"bytes,17,opt,name=last_update_recovery_timestamp,json=lastUpdateRecoveryTimestamp" json:"last_update_recovery_timestamp,omitempty"`
	// Time since last Recovery event (in seconds)
	LastUpdateRecoveryAge uint32 `protobuf:"varint,18,opt,name=last_update_recovery_age,json=lastUpdateRecoveryAge" json:"last_update_recovery_age,omitempty"`
	// Memory allocation failure count
	UpdateMemoryAllocationFailCount uint32 `protobuf:"varint,19,opt,name=update_memory_allocation_fail_count,json=updateMemoryAllocationFailCount" json:"update_memory_allocation_fail_count,omitempty"`
	// Memory allocation failure time
	LastUpdateMemoryAllocationFailTimestamp *BgpTimespec `protobuf:"bytes,20,opt,name=last_update_memory_allocation_fail_timestamp,json=lastUpdateMemoryAllocationFailTimestamp" json:"last_update_memory_allocation_fail_timestamp,omitempty"`
	// Time since last memory allocation failure event (in seconds)
	LastUpdateMemoryAllocationFailAge uint32 `protobuf:"varint,21,opt,name=last_update_memory_allocation_fail_age,json=lastUpdateMemoryAllocationFailAge" json:"last_update_memory_allocation_fail_age,omitempty"`
}

func (m *BgpUpdgenStatsBag) Reset()                    { *m = BgpUpdgenStatsBag{} }
func (m *BgpUpdgenStatsBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdgenStatsBag) ProtoMessage()               {}
func (*BgpUpdgenStatsBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesHigh() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesHigh
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesCumulative() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesCumulative
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesDiscarded() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesDiscarded
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesCleared() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesCleared
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeHigh() uint32 {
	if m != nil {
		return m.UpdateOutQueueSizeHigh
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeCumulative() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeCumulative
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeDiscarded() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeDiscarded
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeCleared() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeCleared
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateDiscardTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateDiscardTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateDiscardAge() uint32 {
	if m != nil {
		return m.LastUpdateDiscardAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateClearedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateClearedTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateCleardAge() uint32 {
	if m != nil {
		return m.LastUpdateCleardAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateThrottleCount() uint32 {
	if m != nil {
		return m.UpdateThrottleCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateThrottleTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateThrottleTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateThrottleAge() uint32 {
	if m != nil {
		return m.LastUpdateThrottleAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateRecoveryCount() uint32 {
	if m != nil {
		return m.UpdateRecoveryCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateRecoveryTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateRecoveryTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateRecoveryAge() uint32 {
	if m != nil {
		return m.LastUpdateRecoveryAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateMemoryAllocationFailCount() uint32 {
	if m != nil {
		return m.UpdateMemoryAllocationFailCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateMemoryAllocationFailTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateMemoryAllocationFailTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateMemoryAllocationFailAge() uint32 {
	if m != nil {
		return m.LastUpdateMemoryAllocationFailAge
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpUpdgenSubgrpBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_updgen_subgrp_bag_KEYS")
	proto.RegisterType((*BgpUpdgenSubgrpBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_updgen_subgrp_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_timespec")
	proto.RegisterType((*BgpUpdgenStatsBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_updgen_stats_bag")
}

func init() { proto.RegisterFile("bgp_updgen_subgrp_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0x59, 0x73, 0xd4, 0x46,
	0x10, 0x2e, 0x11, 0xce, 0xb1, 0x0d, 0x58, 0xbe, 0xe4, 0x0b, 0x2f, 0x86, 0x10, 0xa7, 0x8a, 0x6c,
	0x2a, 0x98, 0x23, 0x10, 0x42, 0x58, 0xb0, 0xb9, 0xcd, 0xb1, 0x6b, 0xa0, 0xf2, 0x34, 0x35, 0x2b,
	0xf5, 0xca, 0x53, 0x68, 0x25, 0x65, 0x66, 0xb4, 0x10, 0xfe, 0x43, 0x7e, 0x41, 0x1e, 0xf3, 0x43,
	0xf2, 0x9e, 0x54, 0xa5, 0x2a, 0x79, 0x4f, 0xe5, 0xa8, 0x54, 0xfe, 0x44, 0x5e, 0x52, 0x73, 0x49,
	0xda, 0x5d, 0xad, 0x8b, 0x47, 0xf6, 0xcd, 0x9e, 0xee, 0xaf, 0xa7, 0xbf, 0xaf, 0xdb, 0xad, 0x1e,
	0xa3, 0xe5, 0x76, 0x98, 0xe2, 0x2c, 0x0d, 0x42, 0x88, 0x31, 0xcf, 0xda, 0x21, 0x4b, 0x71, 0x9b,
	0x84, 0xf5, 0x94, 0x25, 0x22, 0x71, 0xbf, 0x73, 0x7c, 0xca, 0xfd, 0x04, 0xd3, 0x84, 0xe3, 0x37,
	0x0c, 0xd3, 0xb4, 0x77, 0x11, 0x4b, 0xff, 0x24, 0x05, 0x56, 0x6f, 0x87, 0x69, 0x9d, 0xc6, 0x5c,
	0x90, 0xd8, 0x07, 0x9e, 0xff, 0x94, 0xff, 0x80, 0x89, 0x2f, 0x68, 0x0f, 0xea, 0x01, 0x74, 0x48,
	0x16, 0x09, 0xdc, 0x63, 0x9d, 0x3a, 0xe9, 0xf0, 0x3a, 0xe9, 0xd4, 0xb3, 0x34, 0x20, 0x02, 0x70,
	0x08, 0x31, 0x30, 0x22, 0x68, 0xa2, 0x6e, 0xc6, 0x21, 0x4b, 0xb2, 0x94, 0xef, 0x67, 0x5c, 0xff,
	0xcd, 0x41, 0x4b, 0x95, 0xf9, 0xe2, 0x87, 0xdb, 0x5f, 0xb7, 0xdc, 0x33, 0x68, 0x2a, 0xbf, 0x3d,
	0x26, 0x5d, 0xf0, 0x9c, 0x9a, 0xb3, 0x71, 0xac, 0x39, 0x69, 0x0f, 0x1f, 0x93, 0x2e, 0xb8, 0x0b,
	0xe8, 0x08, 0xe9, 0x68, 0xf3, 0x01, 0x65, 0x3e, 0x4c, 0x3a, 0xca, 0x70, 0x1e, 0xb9, 0xf6, 0x6e,
	0x79, 0x19, 0xa6, 0x71, 0x00, 0x6f, 0xbc, 0x0f, 0x6a, 0xce, 0xc6, 0x54, 0xf3, 0xa4, 0xb6, 0xdc,
	0x95, 0x86, 0xfb, 0xf2, 0xdc, 0x3d, 0x87, 0x4e, 0xe4, 0x79, 0x19, 0xd7, 0x83, 0xca, 0x75, 0x8a,
	0x67, 0xed, 0x92, 0x5f, 0x0d, 0x4d, 0x96, 0xfc, 0x02, 0xef, 0x90, 0x72, 0x42, 0xb9, 0x53, 0xb0,
	0xfe, 0xd3, 0x22, 0x9a, 0xab, 0x24, 0xe5, 0xae, 0x22, 0x94, 0xb2, 0xc4, 0x07, 0xce, 0x25, 0xf2,
	0x82, 0x42, 0x1e, 0x33, 0x27, 0xf7, 0x03, 0x99, 0x82, 0x49, 0xb8, 0xc7, 0x0c, 0xa3, 0x4d, 0xc5,
	0x68, 0x4a, 0x1f, 0xbf, 0x60, 0x9a, 0xd8, 0xa7, 0x68, 0xb6, 0x8f, 0x98, 0xa5, 0x7f, 0x51, 0x39,
	0x4f, 0x97, 0xa8, 0x35, 0x34, 0xa0, 0x82, 0xdb, 0xa5, 0x77, 0xe1, 0x76, 0x79, 0x90, 0x9b, 0xbb,
	0x89, 0xe6, 0x53, 0xc2, 0x20, 0x16, 0x78, 0x30, 0xe0, 0x15, 0xe5, 0x3b, 0xa3, 0xad, 0xad, 0xbe,
	0xb0, 0x9f, 0xa0, 0x99, 0x61, 0x50, 0xe0, 0x7d, 0xae, 0x2b, 0x31, 0x80, 0x08, 0x46, 0xd4, 0xed,
	0xea, 0x88, 0xba, 0x5d, 0x45, 0x8b, 0xc6, 0xbb, 0x4b, 0x68, 0x8c, 0x05, 0x69, 0x47, 0x80, 0x7b,
	0xc0, 0x38, 0x4d, 0x62, 0xef, 0x9a, 0x02, 0xcd, 0x6b, 0x87, 0x1d, 0x42, 0xe3, 0x5d, 0x69, 0x7e,
	0xa1, 0xad, 0xee, 0x0d, 0xb4, 0x52, 0xd2, 0x5b, 0x23, 0x19, 0x6d, 0xe7, 0xe8, 0x2f, 0x14, 0xda,
	0xcb, 0xc5, 0x57, 0xe0, 0x26, 0x6d, 0x5b, 0xfc, 0x2d, 0x74, 0xca, 0xcf, 0x98, 0x22, 0x66, 0xe2,
	0x44, 0xb4, 0x4b, 0x4b, 0x2c, 0xbd, 0xeb, 0x2a, 0xc2, 0x92, 0xf1, 0x7a, 0xae, 0x9c, 0x1e, 0x49,
	0x1f, 0x4b, 0xd7, 0xbd, 0x87, 0x4e, 0xfb, 0x49, 0xdc, 0xa1, 0x61, 0xc6, 0x20, 0x18, 0x15, 0xe6,
	0x4b, 0x15, 0x66, 0xb5, 0x70, 0xac, 0x8a, 0x54, 0x08, 0x91, 0x64, 0x02, 0x7f, 0x93, 0x41, 0x06,
	0xb8, 0x0b, 0x9c, 0x93, 0x10, 0xb8, 0x77, 0xa3, 0x2c, 0xc4, 0x93, 0x4c, 0x3c, 0x93, 0xe6, 0x1d,
	0x63, 0x75, 0x3f, 0x43, 0x73, 0x43, 0x50, 0x4e, 0xdf, 0x82, 0xf7, 0x95, 0x82, 0xb9, 0xfd, 0xb0,
	0x16, 0x7d, 0x0b, 0xee, 0xc7, 0xc8, 0x94, 0x02, 0x8b, 0x3d, 0x96, 0x08, 0x11, 0x41, 0xe0, 0xdd,
	0xac, 0x39, 0x1b, 0x47, 0x9b, 0xa6, 0x87, 0x77, 0xed, 0xb1, 0x7b, 0x09, 0x2d, 0x30, 0xe8, 0x30,
	0xe0, 0x7b, 0xa5, 0xfa, 0xfb, 0x49, 0x16, 0x0b, 0xaf, 0xa1, 0xe2, 0xcf, 0x1a, 0xb3, 0xa5, 0x72,
	0x5b, 0xda, 0x64, 0x1b, 0x74, 0x68, 0x24, 0x80, 0xf5, 0x21, 0x6e, 0xe9, 0x36, 0xd0, 0x96, 0x92,
	0xf7, 0x87, 0xe8, 0x78, 0x0c, 0x34, 0xdc, 0x6b, 0x27, 0xcc, 0x78, 0xde, 0xd6, 0x1d, 0x6e, 0x4f,
	0xb5, 0x9b, 0x87, 0x8e, 0xd8, 0xea, 0x6e, 0x29, 0xbb, 0xfd, 0xd5, 0x5d, 0x43, 0x13, 0x31, 0x67,
	0x79, 0xed, 0xb7, 0x75, 0xeb, 0xc7, 0x9c, 0xd9, 0x6a, 0x5f, 0x44, 0xf3, 0x29, 0xc4, 0x01, 0x8d,
	0x43, 0x2c, 0x08, 0x0b, 0x41, 0xe4, 0xbe, 0x77, 0x34, 0x0b, 0x63, 0xdd, 0x55, 0x46, 0x8b, 0xaa,
	0xa3, 0x99, 0x18, 0xde, 0x08, 0xcc, 0x80, 0x67, 0xdd, 0xa2, 0x31, 0xef, 0x2a, 0xc8, 0xb4, 0x34,
	0x35, 0x95, 0xc5, 0xfa, 0x37, 0xd0, 0xaa, 0xd1, 0xd5, 0x6a, 0x36, 0x70, 0xd9, 0x3d, 0xdd, 0x52,
	0xda, 0xa9, 0xa9, 0x7d, 0xfa, 0xaf, 0xdc, 0x44, 0xf3, 0x34, 0xc1, 0xaf, 0x19, 0x15, 0x80, 0xa1,
	0x27, 0xbb, 0xd3, 0x64, 0xe6, 0xdd, 0x57, 0x05, 0x9a, 0xa1, 0xc9, 0x4b, 0x69, 0xdc, 0x96, 0xb6,
	0xa7, 0xda, 0x54, 0xea, 0x1e, 0x3b, 0xdf, 0x59, 0x92, 0x09, 0xc0, 0x1c, 0x62, 0xe1, 0x3d, 0x50,
	0x38, 0xd3, 0x3d, 0x5b, 0xda, 0xde, 0x94, 0xe6, 0x16, 0xc4, 0xc2, 0xfd, 0xd9, 0x41, 0xae, 0xcf,
	0x40, 0xcf, 0x76, 0x41, 0xbb, 0xc0, 0x05, 0xe9, 0xa6, 0xde, 0xc3, 0x9a, 0xb3, 0x31, 0x71, 0xe1,
	0x7b, 0xa7, 0xfe, 0x5e, 0x7d, 0x72, 0xe4, 0x95, 0x3a, 0xc9, 0x14, 0xfc, 0xe6, 0xb4, 0xcd, 0x7b,
	0xd7, 0xa6, 0x2d, 0xfb, 0xa0, 0x0b, 0x2c, 0x04, 0xd3, 0x45, 0x8f, 0x74, 0x1f, 0xa8, 0x23, 0xdd,
	0x42, 0xbf, 0x38, 0x68, 0x36, 0x22, 0x5c, 0x60, 0xed, 0x56, 0x10, 0xde, 0x19, 0x03, 0xc2, 0xae,
	0xcc, 0x7c, 0x47, 0x26, 0x5e, 0x30, 0xbe, 0x8e, 0x96, 0x0b, 0x3e, 0xc1, 0xd0, 0x60, 0x7f, 0xac,
	0x14, 0x58, 0xc8, 0x81, 0x41, 0xff, 0x70, 0x3f, 0x8b, 0x8e, 0x43, 0x82, 0x19, 0x26, 0x42, 0x40,
	0x37, 0x15, 0x10, 0x78, 0x4f, 0x54, 0xb7, 0x4c, 0x42, 0xd2, 0x6c, 0xd8, 0x33, 0xf7, 0x5f, 0x07,
	0x9d, 0xee, 0x50, 0xc6, 0xf3, 0x49, 0xf9, 0x9a, 0x44, 0xaf, 0x30, 0x17, 0x84, 0x89, 0x92, 0x82,
	0x4f, 0xc7, 0x40, 0xc1, 0x55, 0x45, 0x43, 0x0f, 0xe0, 0x97, 0x24, 0x7a, 0xd5, 0x92, 0x1c, 0x0a,
	0x31, 0xff, 0x76, 0xd0, 0xda, 0x30, 0x51, 0x88, 0x83, 0x12, 0xcd, 0x67, 0x63, 0x40, 0x73, 0x79,
	0x80, 0xe6, 0x76, 0x1c, 0x14, 0x24, 0xff, 0x71, 0x50, 0x4d, 0xb5, 0xcc, 0x7e, 0xc5, 0x6c, 0x8e,
	0x01, 0xcb, 0x15, 0xc9, 0x62, 0x64, 0x2d, 0xff, 0x74, 0xd0, 0xa9, 0x21, 0x9a, 0xfd, 0xa5, 0x6c,
	0x8d, 0x01, 0xc9, 0xa5, 0x7e, 0x92, 0x7d, 0x95, 0xbc, 0x8c, 0xbc, 0x4a, 0x86, 0x24, 0x04, 0x6f,
	0x57, 0x7f, 0xd6, 0x86, 0xd0, 0x8d, 0x10, 0xdc, 0xdf, 0x1d, 0xb4, 0x52, 0x06, 0xaa, 0x9d, 0xa1,
	0x2c, 0xcc, 0xf3, 0x31, 0x10, 0x66, 0xb1, 0xa0, 0xa6, 0x16, 0x9b, 0x92, 0x2e, 0x9b, 0x68, 0xbe,
	0x82, 0x9e, 0x54, 0xe5, 0x85, 0xde, 0x73, 0x07, 0xa1, 0x52, 0x94, 0x5f, 0x1d, 0x64, 0x96, 0x6f,
	0xf9, 0xc7, 0x20, 0x28, 0x17, 0xd4, 0xe7, 0xde, 0x4b, 0xa5, 0xc4, 0x0f, 0xef, 0xa3, 0x12, 0xf6,
	0x85, 0x22, 0x88, 0xe0, 0xf2, 0x81, 0x62, 0xd7, 0xeb, 0x56, 0x9e, 0xfd, 0xfa, 0x03, 0x34, 0x59,
	0xd6, 0x4c, 0x2e, 0x50, 0x1c, 0xfc, 0x24, 0x0e, 0xb8, 0x7a, 0x8c, 0x4d, 0x35, 0xed, 0xaf, 0x6e,
	0x0d, 0x4d, 0xc4, 0x24, 0x4e, 0xac, 0xf5, 0x80, 0xb2, 0x96, 0x8f, 0xd6, 0x7f, 0x9c, 0x46, 0xb3,
	0x55, 0xd7, 0xca, 0x45, 0x7a, 0xe4, 0xea, 0x8a, 0xf7, 0x68, 0xb8, 0x67, 0xee, 0x5a, 0xaa, 0xde,
	0x5f, 0xef, 0xd1, 0x70, 0xcf, 0x7d, 0x8c, 0xce, 0x8e, 0x8e, 0xe1, 0x67, 0xdd, 0x2c, 0x22, 0x52,
	0x43, 0x93, 0x57, 0xad, 0x3a, 0xd2, 0xed, 0xdc, 0xcf, 0x7d, 0x84, 0xce, 0x8c, 0x8e, 0x17, 0x50,
	0xee, 0x13, 0x16, 0x40, 0x60, 0x9e, 0x93, 0x6b, 0xd5, 0xe1, 0xb6, 0xac, 0x9b, 0x5c, 0xf3, 0xf7,
	0xc9, 0x2e, 0x02, 0xc2, 0x20, 0x30, 0xef, 0xcd, 0xd5, 0x11, 0xa9, 0x69, 0x27, 0xf7, 0x1a, 0x5a,
	0xaa, 0xdc, 0xd5, 0xb5, 0x4e, 0x87, 0xaa, 0xf6, 0x7c, 0xb9, 0xb0, 0x2b, 0x8d, 0xee, 0xa0, 0x5a,
	0x35, 0xb6, 0xa4, 0xcf, 0xe1, 0x9a, 0xb3, 0x71, 0xb0, 0xb9, 0x32, 0x1c, 0xa1, 0xa4, 0xcd, 0x16,
	0x5a, 0xab, 0x8e, 0x53, 0xe8, 0x72, 0x44, 0x85, 0x59, 0x1e, 0x0e, 0x53, 0x68, 0x72, 0x33, 0x5f,
	0x75, 0x07, 0xb3, 0x31, 0x7a, 0x1c, 0x55, 0x31, 0x16, 0x2b, 0x52, 0x31, 0x5a, 0xfc, 0xe1, 0xa0,
	0xd5, 0xf2, 0x9f, 0xa9, 0xb9, 0xbe, 0x34, 0x86, 0x8e, 0x8d, 0xd7, 0x7c, 0x36, 0xe2, 0x14, 0x73,
	0xe8, 0x12, 0x5a, 0xa8, 0x22, 0x28, 0x07, 0x11, 0x1a, 0x1c, 0xcf, 0x06, 0x2c, 0x27, 0xd1, 0xa0,
	0x30, 0x46, 0xd1, 0x92, 0x30, 0x13, 0xe3, 0x25, 0x8c, 0xa9, 0xf8, 0xc8, 0x01, 0xad, 0x08, 0x6a,
	0x5d, 0x26, 0x07, 0x07, 0xb4, 0xc2, 0x2a, 0x59, 0x2e, 0xe4, 0xef, 0x5c, 0xfb, 0x68, 0x35, 0x5b,
	0xfe, 0x94, 0xc6, 0xf4, 0xbf, 0x5c, 0xf5, 0xba, 0xff, 0xd7, 0xc0, 0x12, 0x90, 0x23, 0x0b, 0x2d,
	0x8f, 0x8f, 0xc3, 0x3e, 0x57, 0xe8, 0x61, 0xf9, 0x15, 0x62, 0x5e, 0xe9, 0xdf, 0x02, 0x72, 0x8a,
	0x52, 0xce, 0x13, 0x4a, 0x9a, 0xb9, 0x61, 0x78, 0xbf, 0xa0, 0x0c, 0xfc, 0xa4, 0x07, 0xec, 0x5b,
	0x23, 0xe8, 0xc9, 0xb2, 0xa0, 0x4d, 0x63, 0xab, 0x16, 0x34, 0x47, 0x16, 0x82, 0x4e, 0x8f, 0x97,
	0xa0, 0x96, 0xdf, 0x48, 0x41, 0x73, 0x8a, 0x52, 0x50, 0x77, 0x50, 0x50, 0x0b, 0x97, 0x82, 0x16,
	0x5f, 0x9d, 0x2e, 0x74, 0x13, 0x89, 0x88, 0xa2, 0xc4, 0xd7, 0x89, 0x74, 0x08, 0x8d, 0x8c, 0xbc,
	0x33, 0xe5, 0xaf, 0xce, 0x8e, 0xf2, 0x6c, 0xe4, 0x8e, 0x77, 0x08, 0x8d, 0xb4, 0xd4, 0xff, 0x39,
	0xe8, 0x7c, 0x39, 0x8f, 0x11, 0x31, 0x0b, 0xe1, 0x67, 0xc7, 0x40, 0xf8, 0x8f, 0x0a, 0xe5, 0xaa,
	0x98, 0x17, 0x45, 0x78, 0x86, 0xce, 0xbd, 0x03, 0x79, 0x59, 0x92, 0x39, 0x25, 0xe7, 0xe9, 0xfd,
	0x03, 0x37, 0x42, 0x68, 0x1f, 0x56, 0xff, 0x46, 0xdf, 0xfc, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x9b,
	0x29, 0x79, 0xef, 0x65, 0x17, 0x00, 0x00,
}