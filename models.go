package main

import (
	"fmt"
	timeext "time"
)

type DateTime struct {
	timeext.Time
}

type Node struct {

	// control address
	ControlAddress string `json:"controladdress"`

	// o s
	// Example: rhel
	// Required: true
	OS string `json:"os"`

	// offline time
	OfflineTime DateTime `json:"offlinetime"`

	// credential
	Credential string `json:"credential"`

	// name
	// Example: rv1
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// zone
	// Example: z1
	Zone string `json:"zone"`
}

func (r Node) Info() string {
	return fmt.Sprintf(`
Name  : %s
OS    : %s
Status: %s
State : %s
Zone  : %s
`, r.Name, r.OS, r.Status, r.State, r.Zone)
}

type Volume struct {

	// consistency group
	ConsistencyGroup string `json:"consistencygroup"`

	// content snapshot
	ContentSnapshot string `json:"contentsnapshot"`

	// content volume
	ContentVolume string `json:"contentvolume"`

	// Volume name
	// Example: vol1
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// replication controller
	ReplicationController string `json:"replicationcontroller"`

	// replication node
	ReplicationNode string `json:"replicationnode"`

	// Replication Volume ID
	// Read Only: true
	// Min Length: 1
	ReplicationVolumeGroupID string `json:"replicationvolumegroupid"`

	// Replication Volume Group name
	// Example: vg_1
	// Min Length: 1
	ReplicationVolumeGroupName string `json:"replicationvolumegroupname"`

	// Volume ID
	// Read Only: true
	// Min Length: 1
	VolumeGroupID string `json:"volumegroupid"`

	// Volume Group name
	// Example: vg_1
	// Min Length: 1
	VolumeGroupName string `json:"volumegroupname"`

	// Volume ID
	// Read Only: true
	// Min Length: 1
	VolumeID string `json:"volumeid"`

	// controller
	Controller string `json:"controller"`

	// node
	Node *string `json:"node"`

	// policy
	// Required: true
	// Min Length: 1
	Policy string `json:"policy"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// Size, GiB
	// Example: 10
	// Required: true
	// Maximum: 65536
	// Minimum: 1
	Size int64 `json:"size"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// type
	// Required: true
	// Enum: [file block]
	Type string `json:"type"`

	// zone
	Zone *string `json:"zone"`

	// zonereplica
	Zonereplica *string `json:"zonereplica"`
}

func (r Volume) Info() string {
	return fmt.Sprintf(`
Name        : %s
VolumeID    : %s	
Size        : %dGB
Node        : %s
Policy      : %s
State       : %s
Status      : %s
Type        : %s
Zone        : %v
`, r.Name, r.VolumeID, r.Size, *r.Node, r.Policy, r.State, r.Status, r.Type, r.Zone)
}

type Attachment struct {

	// node
	// Required: true
	Node string `json:"node"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// protocol
	// Required: true
	// Enum: [local iscsi fc nvmeof]
	Protocol string `json:"protocol"`

	// Snapshot ID
	// Read Only: true
	SnapshotID string `json:"snapshotid"`

	// Snapshot Name
	SnapshotName string `json:"snapshotname"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// Volume ID
	// Read Only: true
	VolumeID string `json:"volumeid"`

	// Volume name
	// Example: vol1
	VolumeName string `json:"volumename"`
}

func (r Attachment) Info() string {
	return fmt.Sprintf(`
SnapshotID  : %s
SnapshotName: %s
Node        : %s
Protocol    : %s	
State       : %s
Status      : %s
VolumeID    : %s
VolumeName  : %v
`, r.SnapshotID, r.SnapshotName, r.Node, r.Protocol, r.State, r.Status, r.VolumeID, r.VolumeName)
}

type Connectivity struct {

	// media protocol
	// Required: true
	MediaProtocol string `json:"mediaprotocol"`

	// replication bandwidth
	// Minimum: 0
	ReplicationBandwidth int64 `json:"replicationbandwidth"`

	// replication protocol
	// Required: true
	ReplicationProtocol string `json:"replicationprotocol"`

	// system types1
	// Required: true
	// Pattern: ^(?:(?:windows|rhel|sles|amzn|ubuntu|aws|gcloud|swift)[, ])*(?:windows|rhel|sles|amzn|ubuntu|aws|gcloud|swift)$
	SystemTypes1 string `json:"systemtypes1"`

	// system types2
	// Required: true
	// Pattern: ^(?:(?:windows|rhel|sles|amzn|ubuntu|aws|gcloud|swift)[, ])*(?:windows|rhel|sles|amzn|ubuntu|aws|gcloud|swift)$
	SystemTypes2 string `json:"systemtypes2"`

	// name
	// Required: true
	Name string `json:"name"`

	// zones1
	// Required: true
	Zones1 string `json:"zones1"`

	// zones2
	// Required: true
	Zones2 string `json:"zones2"`
}
type Job struct {

	// end time
	EndTime DateTime `json:"endtime"`

	// ID
	ID int64 `json:"id"`

	// start time
	StartTime DateTime `json:"starttime"`

	// args
	Args map[string]string `json:"args"`

	// object
	Object string `json:"object"`

	// progress
	Progress int64 `json:"progress"`

	// state
	State string `json:"state"`

	// status
	Status string `json:"status"`

	// type
	Type string `json:"type"`
}

func (r Job) Info() string {
	return fmt.Sprintf(`
ID          : %s
Args        : %v
StartTime   : %v
EndTime     : %v
State       : %s
Status      : %s
Type        : %s
`, r.ID, r.Args, r.StartTime, r.EndTime, r.State, r.Status, r.Type)
}

type Media struct {

	// bandwidth read
	BandwidthRead int64 `json:"bandwidthread"`

	// bandwidth write
	BandwidthWrite int64 `json:"bandwidthwrite"`

	// free bandwidth read
	FreeBandwidthRead int64 `json:"freebandwidthread"`

	// free bandwidth write
	FreeBandwidthWrite int64 `json:"freebandwidthwrite"`

	// free i o p s read
	FreeIOPSRead int64 `json:"freeiopsread"`

	// free i o p s write
	FreeIOPSWrite int64 `json:"freeiopswrite"`

	// free size
	FreeSize int64 `json:"freesize"`

	// i o p s read
	IOPSRead int64 `json:"iopsread"`

	// i o p s write
	IOPSWrite int64 `json:"iopswrite"`

	// latency read
	LatencyRead int64 `json:"latencyread"`

	// latency write
	LatencyWrite int64 `json:"latencywrite"`

	// media ID
	MediaID string `json:"mediaid"`

	// offline time
	OfflineTime DateTime `json:"offlinetime"`

	// s e d
	SED bool `json:"sed"`

	// sector size
	SectorSize int64 `json:"sectorsize"`

	// zero on init
	ZeroOnInit bool `json:"zerooninit"`

	// assignment
	// Read Only: true
	Assignment string `json:"assignment"`

	// bus
	Bus string `json:"bus"`

	// firmware
	Firmware string `json:"firmware"`

	// location
	Location string `json:"location"`

	// media
	Media string `json:"media"`

	// model
	Model string `json:"model"`

	// node
	Node string `json:"node"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// size
	Size int64 `json:"size"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// zone
	Zone string `json:"zone"`
}

func (m Media) Info() string {
	return fmt.Sprintf(`
MediaID: %s
Size   : %dGB
Media  : %s
Node   : %s
State  : %s
Status : %s
Assignment: %s
BandwidthRead     : %-8d FreeBandwidthRead  : %-8d  IOPSRead : %-8d FreeIOPSRead : %-8d  LatencyRead :  %-8d
BandwidthWrite    : %-8d FreeBandwidthWrite : %-8d  IOPSWrite: %-8d FreeIOPSWrite: %-8d  LatencyWrite:  %-8d
BandwidthRead Used: %-8d BandwidthWrite Used: %-8d  
IOPSRead Used     : %-8d IOPSWrite Used     : %-8d
Zone : %s
Model: %s
`, m.MediaID, m.Size, m.Media, m.Node, m.State, m.Status, m.Assignment,
		m.BandwidthRead, m.FreeBandwidthRead, m.IOPSRead, m.FreeIOPSRead, m.LatencyRead,
		m.BandwidthWrite, m.FreeBandwidthWrite, m.IOPSWrite, m.FreeIOPSWrite, m.LatencyWrite,
		m.BandwidthRead-m.FreeBandwidthRead, m.BandwidthWrite-m.FreeBandwidthWrite,
		m.IOPSRead-m.FreeIOPSRead, m.IOPSWrite-m.FreeIOPSWrite,
		m.Zone, m.Model,
	)
}

type MediaProfile struct {

	// bandwidth read
	BandwidthRead int64 `json:"bandwidthread"`

	// bandwidth write
	BandwidthWrite int64 `json:"bandwidthwrite"`

	// free bandwidth read
	FreeBandwidthRead int64 `json:"freebandwidthread"`

	// free bandwidth write
	FreeBandwidthWrite int64 `json:"freebandwidthwrite"`

	// free i o p s read
	FreeIOPSRead int64 `json:"freeiopsread"`

	// free i o p s write
	FreeIOPSWrite int64 `json:"freeiopswrite"`

	// i o p s read
	IOPSRead int64 `json:"iopsread"`

	// i o p s write
	IOPSWrite int64 `json:"iopswrite"`

	// latency read
	LatencyRead int64 `json:"latencyread"`

	// latency write
	LatencyWrite int64 `json:"latencywrite"`
}

type Network struct {

	// IP end
	// Required: true
	// Pattern: ^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3})|(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$
	IPEnd string `json:"ipend"`

	// IP start
	// Required: true
	// Pattern: ^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3})|(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$
	IPStart string `json:"ipstart"`

	// name
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// type
	// Required: true
	// Enum: [management storage]
	Type string `json:"type"`

	// zone
	// Min Length: 1
	Zone string `json:"zone"`
}

type Policy struct {

	// Enter the maximum read bandwidth that a volume is expected to sustain. Read Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	BandwidthRead int64 `json:"bandwidthread"`

	// Enter the maximum write bandwidth that a volume is expected to sustain. Write Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	BandwidthWrite int64 `json:"bandwidthwrite"`

	// Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes.
	// Required: true
	// Enum: [capacity balanced performance]
	CapacityOptimization string `json:"capacityoptimization"`

	// Enter how much logical capacity is reserved up-front for the applications to use. If more capacity is needed for the volume, it will be allocated based on availability of media. Capacities that are reserved can be used for the volume itself and for its snapshots. For example – Use 0% for thin-provisioned volumes, 130% for thick-provisioned volumes with estimated 30% of space for snapshots. Valid values are 0%-500%, default is 20%.
	// Example: 20
	// Maximum: 500
	// Minimum: 0
	CapacityReservation int64 `json:"capacityreservation"`

	// Enter the percentage of the volume’s capacity that is expected to be “cold” (i.e. expected to only have infrequent reads). Default is 0%. Values that are greater than 0 give Volumez the option to use more economic media with more relaxed read performance requirements. Valid values: Integers in the range of 0..100.
	// Maximum: 100
	// Minimum: 0
	ColdData int64 `json:"colddata"`

	// Setting this value to “Yes” directs Volumez to over-provision volumes in a way that even after having a failure, the volumes will have the desired performance. Setting this value to “No” directs Volumez to provision volumes according to the desired performance, however in a case of failure – performance may be impacted. The default value is “No”.
	FailurePerformance bool `json:"failureperformance"`

	// Enter the maximum read IOPS that a volume is expected to sustain (assuming 8K reads). Read IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Example: 1000
	// Minimum: 0
	IOPSRead int64 `json:"iopsread"`

	// Enter the maximum write IOPS that a volume is expected to sustain (assuming 8K writes). Write IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Example: 1000
	// Minimum: 0
	IOPSWrite int64 `json:"iopswrite"`

	// Enter the maximum read IOPS that a volume is expected to sustain. Read latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	LatencyRead int64 `json:"latencyread"`

	//  If not all the reads are hot (i.e., Percentage of Cold Reads is >0) – Enter the more relaxed constraints for read latencies of cold data.  Valid values: non-negative integer number, that is larger than “Read Latency”.
	// Minimum: 0
	LatencyReadCold int64 `json:"latencyreadcold"`

	// Enter the maximum latency that a volume is expected to sustain. Write latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	LatencyWrite int64 `json:"latencywrite"`

	// Setting this value to “Yes” directs Volumez to prefer volume configurations where reads are usually happening from disks that are in the same zone as the application. This saves east-west network traffic across zones, however more media per zone will be required to achieve read-IOPs requirements. Set this value to “Yes” if you have network constraints (bandwidth or cost) across your zones; otherwise set it to “No”.
	LocalZoneRead bool `json:"localzoneread"`

	// Specifies the maximum bandwidth that Volumez is allowed to consume for replication of this volume (MB/s). 0 means no bandwidth limitation.
	ReplicationBandwidth int64 `json:"replicationbandwidth"`

	// Enter how many seconds are allowed for the replica to stay behind the primary storage. 0 means synchronous replication. Valid values are 0..3600, default value is 0. Max value: 3600. (1 hour).
	// Maximum: 3600
	// Minimum: 0
	ReplicationRPO int64 `json:"replicationrpo"`

	//  Enter how many media failures (e.g. disk, memory card) the system is required to sustain, and still serve volumes of this policy. A value of “0” means any disk failure will result data unavailability and loss. Valid values are 0..3, default value is 2.
	// Example: 2
	// Minimum: 0
	ResiliencyMedia int64 `json:"resiliencymedia"`

	// Enter how many Volumez node (e.g. EC2 instance, server) failures the system is required to sustain, and still serve volumes of this policy. This is different than “Media failures” because sometimes multiple media copies may end on a single node. A value of “0” means any node failure will result data unavailability and loss. Valid values are 0..3, default value is 1.
	// Example: 1
	// Minimum: 0
	ResiliencyNode int64 `json:"resiliencynode"`

	// Enter how many regions (e.g. AWS regions zones, DataCenters across continents) failures the system is required to sustain, and still serve volumes of this policy. Note: regions are assumed to reside across WAN distance, with some bandwidth limitations. Valid values are 0 and 1, default value is 0.
	// Example: 1
	// Maximum: 1
	// Minimum: 0
	ResiliencyRegion int64 `json:"resiliencyregion"`

	// Enter how many zones (e.g. AWS availability zones, DataCenters Buildings) failures the system is required to sustain, and still serve volumes of this policy. Note: zones are assumed to be within the same metro distance, and resiliency to zone failures means cross-zone network traffic. Valid values are 0..2, default value is 1.
	// Example: 1
	// Minimum: 0
	ResiliencyZone int64 `json:"resiliencyzone"`

	// snapshot day
	SnapshotDay int64 `json:"snapshotday"`

	// snapshot frequency
	SnapshotFrequency string `json:"snapshotfrequency"`

	// snapshot hour
	SnapshotHour int64 `json:"snapshothour"`

	// snapshot keep
	SnapshotKeep int64 `json:"snapshotkeep"`

	// snapshot minute
	SnapshotMinute int64 `json:"snapshotminute"`

	// Enter “YES” to encrypt the data in server where the application is running. Note: No change is needed in the applications themselves, however encryption will consume some CPU cycles on the application server. Default value NO.
	Encryption bool `json:"encryption"`

	// Enter “YES” to direct Volumez to activate the “Device Mapper integrity” protection for the volume. This capability provides strong integrity checking. Note: No change is needed in the applications themselves, however Data Integrity will consume non-negligible CPU cycles on the application server. Default value: NO.
	Integrity bool `json:"integrity"`

	// A name for the policy. The name can be any non-empty string that does not contain a white space.
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// Enter “YES” to direct Volumez to only use media with disk encryption capabilities. Note that specifying “NO” can still use such media, however it is not a must to use it. Default value: NO.
	Sed bool `json:"sed"`
}

func (r Policy) Info() string {
	return fmt.Sprintf(`
CapacityOptimization: %s CapacityReservation: %d ColdData: %d
BandwidthRead     : %-8d IOPSRead : %-8d LatencyRead :  %-8d
BandwidthWrite    : %-8d IOPSWrite: %-8d LatencyWrite:  %-8d
Encryption : %v   
`, r.CapacityOptimization, r.CapacityReservation, r.ColdData,
		r.BandwidthRead, r.IOPSRead, r.LatencyRead,
		r.BandwidthWrite, r.IOPSWrite, r.LatencyWrite,
		r.Encryption,
	)
}

type Snapshot struct {

	// consistency
	// Required: true
	// Enum: [crash application]
	Consistency string `json:"consistency"`

	// consistency group
	ConsistencyGroup bool `json:"consistencygroup"`

	// policy
	Policy bool `json:"policy"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// Snapshot name
	// Required: true
	// Min Length: 1
	SnapName string `json:"name"`

	// Snapshot ID
	// Read Only: true
	// Min Length: 1
	SnapshotID string `json:"snapshotid"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// time
	Time DateTime `json:"time"`

	// used
	Used int64 `json:"used"`

	// Volume ID
	// Read Only: true
	// Min Length: 1
	VolumeID string `json:"volumeid"`

	// Volume name
	// Example: vol1
	VolumeName string `json:"volumename"`
}

func (r Snapshot) Info() string {
	return fmt.Sprintf(`
SnapshotID  : %s SnapName    : %s VolumeName  : %s
State       : %s Status      : %s Used        : %d
`, r.SnapshotID, r.SnapName, r.VolumeName, r.State, r.Status, r.Used)
}
