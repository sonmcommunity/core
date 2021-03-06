// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ask_plan.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	ask_plan.proto
	benchmarks.proto
	bigint.proto
	capabilities.proto
	container.proto
	dwh.proto
	geoip.proto
	insonmnia.proto
	marketplace.proto
	net.proto
	node.proto
	optimus.proto
	relay.proto
	rendezvous.proto
	timestamp.proto
	volume.proto
	worker.proto

It has these top-level messages:
	AskPlanCPU
	AskPlanGPU
	AskPlanRAM
	AskPlanStorage
	AskPlanNetwork
	AskPlanResources
	AskPlan
	Benchmark
	BigInt
	CPUDevice
	CPU
	RAMDevice
	RAM
	GPUDevice
	GPU
	NetFlags
	Network
	StorageDevice
	Storage
	Registry
	ContainerRestartPolicy
	NetworkSpec
	Container
	SortingOption
	DealsRequest
	DWHDealsReply
	DWHDeal
	DealConditionsRequest
	DealConditionsReply
	OrdersRequest
	MatchingOrdersRequest
	DWHOrdersReply
	DWHOrder
	DealCondition
	DWHWorker
	ProfilesRequest
	ProfilesReply
	Profile
	BlacklistRequest
	BlacklistReply
	BlacklistsContainingUserReply
	ValidatorsRequest
	ValidatorsReply
	DWHValidator
	Validator
	DealChangeRequestsReply
	DealChangeRequest
	DealPayment
	ChangeRequestsRequest
	WorkersRequest
	WorkersReply
	Certificate
	MaxMinUint64
	MaxMinBig
	MaxMinTimestamp
	CmpUint64
	BlacklistQuery
	DWHStatsReply
	OrdersByIDsRequest
	GeoIPCountry
	GeoIP
	Empty
	ID
	NumericID
	EthID
	TaskID
	Count
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	TaskLogsRequest
	TaskLogsChunk
	Chunk
	Progress
	Duration
	EthAddress
	DataSize
	DataSizeRate
	Price
	ErrorByID
	ErrorByStringID
	OrderIDs
	GetOrdersReply
	Benchmarks
	Deal
	Order
	BidNetwork
	BidResources
	BidOrder
	Addr
	SocketAddr
	Endpoints
	JoinNetworkRequest
	TaskListRequest
	QuickBuyRequest
	DealFinishRequest
	DealsFinishRequest
	DealsPurgeRequest
	DealsReply
	OpenDealRequest
	WorkerRemoveRequest
	WorkerListReply
	BalanceReply
	TokenTransferRequest
	NPPMetricsReply
	NamedMetrics
	NamedMetric
	PredictSupplierRequest
	PredictSupplierReply
	HandshakeRequest
	DiscoverResponse
	HandshakeResponse
	RelayClusterReply
	RelayMetrics
	NetMetrics
	RelayInfo
	RelayMeeting
	ConnectRequest
	PublishRequest
	RendezvousReply
	RendezvousState
	RendezvousMeeting
	ResolveMetaReply
	Timestamp
	Volume
	TaskTag
	TaskSpec
	StartTaskRequest
	WorkerJoinNetworkRequest
	StartTaskReply
	StatusReply
	AskPlansReply
	TaskListReply
	DevicesReply
	PullTaskRequest
	DealInfoReply
	TaskStatusReply
	TaskPool
	AskPlanPool
	SchedulerData
	SalesmanData
	DebugStateReply
	PurgeTasksRequest
	WorkerMetricsRequest
	WorkerMetricsResponse
*/
package sonm

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

type AskPlan_Status int32

const (
	AskPlan_ACTIVE           AskPlan_Status = 0
	AskPlan_PENDING_DELETION AskPlan_Status = 1
)

var AskPlan_Status_name = map[int32]string{
	0: "ACTIVE",
	1: "PENDING_DELETION",
}
var AskPlan_Status_value = map[string]int32{
	"ACTIVE":           0,
	"PENDING_DELETION": 1,
}

func (x AskPlan_Status) String() string {
	return proto.EnumName(AskPlan_Status_name, int32(x))
}
func (AskPlan_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type AskPlanCPU struct {
	CorePercents uint64 `protobuf:"varint,1,opt,name=core_percents,json=corePercents" json:"core_percents,omitempty"`
}

func (m *AskPlanCPU) Reset()                    { *m = AskPlanCPU{} }
func (m *AskPlanCPU) String() string            { return proto.CompactTextString(m) }
func (*AskPlanCPU) ProtoMessage()               {}
func (*AskPlanCPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AskPlanCPU) GetCorePercents() uint64 {
	if m != nil {
		return m.CorePercents
	}
	return 0
}

type AskPlanGPU struct {
	Indexes []uint64 `protobuf:"varint,1,rep,packed,name=indexes" json:"indexes,omitempty"`
	Hashes  []string `protobuf:"bytes,2,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *AskPlanGPU) Reset()                    { *m = AskPlanGPU{} }
func (m *AskPlanGPU) String() string            { return proto.CompactTextString(m) }
func (*AskPlanGPU) ProtoMessage()               {}
func (*AskPlanGPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AskPlanGPU) GetIndexes() []uint64 {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *AskPlanGPU) GetHashes() []string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type AskPlanRAM struct {
	Size *DataSize `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
}

func (m *AskPlanRAM) Reset()                    { *m = AskPlanRAM{} }
func (m *AskPlanRAM) String() string            { return proto.CompactTextString(m) }
func (*AskPlanRAM) ProtoMessage()               {}
func (*AskPlanRAM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AskPlanRAM) GetSize() *DataSize {
	if m != nil {
		return m.Size
	}
	return nil
}

type AskPlanStorage struct {
	Size *DataSize `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
}

func (m *AskPlanStorage) Reset()                    { *m = AskPlanStorage{} }
func (m *AskPlanStorage) String() string            { return proto.CompactTextString(m) }
func (*AskPlanStorage) ProtoMessage()               {}
func (*AskPlanStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AskPlanStorage) GetSize() *DataSize {
	if m != nil {
		return m.Size
	}
	return nil
}

type AskPlanNetwork struct {
	ThroughputIn  *DataSizeRate `protobuf:"bytes,1,opt,name=throughputIn" json:"throughputIn,omitempty"`
	ThroughputOut *DataSizeRate `protobuf:"bytes,2,opt,name=throughputOut" json:"throughputOut,omitempty"`
	NetFlags      *NetFlags     `protobuf:"bytes,3,opt,name=netFlags" json:"netFlags,omitempty"`
}

func (m *AskPlanNetwork) Reset()                    { *m = AskPlanNetwork{} }
func (m *AskPlanNetwork) String() string            { return proto.CompactTextString(m) }
func (*AskPlanNetwork) ProtoMessage()               {}
func (*AskPlanNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AskPlanNetwork) GetThroughputIn() *DataSizeRate {
	if m != nil {
		return m.ThroughputIn
	}
	return nil
}

func (m *AskPlanNetwork) GetThroughputOut() *DataSizeRate {
	if m != nil {
		return m.ThroughputOut
	}
	return nil
}

func (m *AskPlanNetwork) GetNetFlags() *NetFlags {
	if m != nil {
		return m.NetFlags
	}
	return nil
}

type AskPlanResources struct {
	CPU     *AskPlanCPU     `protobuf:"bytes,1,opt,name=CPU" json:"CPU,omitempty"`
	RAM     *AskPlanRAM     `protobuf:"bytes,2,opt,name=RAM" json:"RAM,omitempty"`
	Storage *AskPlanStorage `protobuf:"bytes,3,opt,name=storage" json:"storage,omitempty"`
	GPU     *AskPlanGPU     `protobuf:"bytes,4,opt,name=GPU" json:"GPU,omitempty"`
	Network *AskPlanNetwork `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
}

func (m *AskPlanResources) Reset()                    { *m = AskPlanResources{} }
func (m *AskPlanResources) String() string            { return proto.CompactTextString(m) }
func (*AskPlanResources) ProtoMessage()               {}
func (*AskPlanResources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AskPlanResources) GetCPU() *AskPlanCPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *AskPlanResources) GetRAM() *AskPlanRAM {
	if m != nil {
		return m.RAM
	}
	return nil
}

func (m *AskPlanResources) GetStorage() *AskPlanStorage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *AskPlanResources) GetGPU() *AskPlanGPU {
	if m != nil {
		return m.GPU
	}
	return nil
}

func (m *AskPlanResources) GetNetwork() *AskPlanNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

type AskPlan struct {
	ID                  string            `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	OrderID             *BigInt           `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	DealID              *BigInt           `protobuf:"bytes,3,opt,name=dealID" json:"dealID,omitempty"`
	Duration            *Duration         `protobuf:"bytes,4,opt,name=duration" json:"duration,omitempty"`
	Price               *Price            `protobuf:"bytes,5,opt,name=price" json:"price,omitempty"`
	Blacklist           *EthAddress       `protobuf:"bytes,6,opt,name=blacklist" json:"blacklist,omitempty"`
	Counterparty        *EthAddress       `protobuf:"bytes,7,opt,name=counterparty" json:"counterparty,omitempty"`
	Identity            IdentityLevel     `protobuf:"varint,8,opt,name=identity,enum=sonm.IdentityLevel" json:"identity,omitempty"`
	Tag                 []byte            `protobuf:"bytes,9,opt,name=tag,proto3" json:"tag,omitempty"`
	Resources           *AskPlanResources `protobuf:"bytes,10,opt,name=resources" json:"resources,omitempty"`
	Status              AskPlan_Status    `protobuf:"varint,11,opt,name=status,enum=sonm.AskPlan_Status" json:"status,omitempty"`
	CreateTime          *Timestamp        `protobuf:"bytes,12,opt,name=createTime" json:"createTime,omitempty"`
	LastOrderPlacedTime *Timestamp        `protobuf:"bytes,13,opt,name=lastOrderPlacedTime" json:"lastOrderPlacedTime,omitempty"`
}

func (m *AskPlan) Reset()                    { *m = AskPlan{} }
func (m *AskPlan) String() string            { return proto.CompactTextString(m) }
func (*AskPlan) ProtoMessage()               {}
func (*AskPlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AskPlan) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AskPlan) GetOrderID() *BigInt {
	if m != nil {
		return m.OrderID
	}
	return nil
}

func (m *AskPlan) GetDealID() *BigInt {
	if m != nil {
		return m.DealID
	}
	return nil
}

func (m *AskPlan) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AskPlan) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *AskPlan) GetBlacklist() *EthAddress {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *AskPlan) GetCounterparty() *EthAddress {
	if m != nil {
		return m.Counterparty
	}
	return nil
}

func (m *AskPlan) GetIdentity() IdentityLevel {
	if m != nil {
		return m.Identity
	}
	return IdentityLevel_UNKNOWN
}

func (m *AskPlan) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *AskPlan) GetResources() *AskPlanResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *AskPlan) GetStatus() AskPlan_Status {
	if m != nil {
		return m.Status
	}
	return AskPlan_ACTIVE
}

func (m *AskPlan) GetCreateTime() *Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *AskPlan) GetLastOrderPlacedTime() *Timestamp {
	if m != nil {
		return m.LastOrderPlacedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*AskPlanCPU)(nil), "sonm.AskPlanCPU")
	proto.RegisterType((*AskPlanGPU)(nil), "sonm.AskPlanGPU")
	proto.RegisterType((*AskPlanRAM)(nil), "sonm.AskPlanRAM")
	proto.RegisterType((*AskPlanStorage)(nil), "sonm.AskPlanStorage")
	proto.RegisterType((*AskPlanNetwork)(nil), "sonm.AskPlanNetwork")
	proto.RegisterType((*AskPlanResources)(nil), "sonm.AskPlanResources")
	proto.RegisterType((*AskPlan)(nil), "sonm.AskPlan")
	proto.RegisterEnum("sonm.AskPlan_Status", AskPlan_Status_name, AskPlan_Status_value)
}

func init() { proto.RegisterFile("ask_plan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5f, 0x6f, 0xda, 0x3a,
	0x18, 0xc6, 0x0f, 0x7f, 0x0a, 0xe5, 0x2d, 0xa5, 0x1c, 0xb7, 0xaa, 0xa2, 0x5e, 0x71, 0x72, 0x8e,
	0x8e, 0x50, 0x35, 0xd1, 0xad, 0xab, 0xa6, 0x5d, 0x4d, 0x62, 0x85, 0xa1, 0x48, 0x2d, 0x8d, 0x5c,
	0xd8, 0x6d, 0x65, 0x92, 0x57, 0x60, 0x11, 0x9c, 0xc8, 0x36, 0xdb, 0xda, 0x0f, 0xb5, 0x8f, 0xb4,
	0x9b, 0x7d, 0x91, 0xc9, 0x89, 0x03, 0xcd, 0x44, 0xa5, 0xdd, 0x25, 0xcf, 0xf3, 0x7b, 0xec, 0xf7,
	0x7d, 0x6d, 0x19, 0x5a, 0x4c, 0x2d, 0x1f, 0x92, 0x88, 0x89, 0x5e, 0x22, 0x63, 0x1d, 0x93, 0xaa,
	0x8a, 0xc5, 0xea, 0xac, 0x39, 0xe3, 0x73, 0x2e, 0x74, 0xa6, 0x9d, 0x91, 0x80, 0x25, 0x6c, 0xc6,
	0x23, 0xae, 0x39, 0x2a, 0xab, 0x1d, 0x71, 0x61, 0x48, 0xc1, 0x99, 0x15, 0xfe, 0x5e, 0x31, 0xb9,
	0x44, 0x9d, 0x44, 0x2c, 0xc0, 0x9c, 0xd1, 0x7c, 0x85, 0x4a, 0xb3, 0x55, 0x92, 0x09, 0xee, 0x1b,
	0x80, 0xbe, 0x5a, 0xfa, 0x11, 0x13, 0xd7, 0xfe, 0x94, 0xfc, 0x0b, 0x87, 0x41, 0x2c, 0xf1, 0x21,
	0x41, 0x19, 0xa0, 0xd0, 0xca, 0x29, 0x75, 0x4a, 0xdd, 0x2a, 0x6d, 0x1a, 0xd1, 0xb7, 0x9a, 0xfb,
	0x61, 0x13, 0x19, 0xf9, 0x53, 0xe2, 0x40, 0x9d, 0x8b, 0x10, 0xbf, 0xa1, 0x81, 0x2b, 0xdd, 0x2a,
	0xcd, 0x7f, 0xc9, 0x29, 0xd4, 0x16, 0x4c, 0x2d, 0x50, 0x39, 0xe5, 0x4e, 0xa5, 0xdb, 0xa0, 0xf6,
	0xcf, 0x7d, 0xbd, 0xc9, 0xd3, 0xfe, 0x2d, 0x71, 0xa1, 0xaa, 0xf8, 0x13, 0xa6, 0x3b, 0x1d, 0x5c,
	0xb6, 0x7a, 0xa6, 0x85, 0xde, 0x80, 0x69, 0x76, 0xcf, 0x9f, 0x90, 0xa6, 0x9e, 0x7b, 0x05, 0x2d,
	0x9b, 0xb8, 0xd7, 0xb1, 0x64, 0x73, 0xfc, 0xa3, 0xd4, 0xf7, 0xd2, 0x26, 0x36, 0x46, 0xfd, 0x35,
	0x96, 0x4b, 0xf2, 0x0e, 0x9a, 0x7a, 0x21, 0xe3, 0xf5, 0x7c, 0x91, 0xac, 0xb5, 0x27, 0x6c, 0x9c,
	0xfc, 0x16, 0x67, 0x1a, 0x69, 0x81, 0x23, 0xef, 0xe1, 0x70, 0xfb, 0x7f, 0xb7, 0xd6, 0x4e, 0xf9,
	0xc5, 0x60, 0x11, 0x24, 0xe7, 0xb0, 0x2f, 0x50, 0x7f, 0x8a, 0xd8, 0x5c, 0x39, 0x95, 0xe7, 0xc5,
	0x8e, 0xad, 0x4a, 0x37, 0xbe, 0xfb, 0xa3, 0x04, 0xed, 0x7c, 0x32, 0xa8, 0xe2, 0xb5, 0x0c, 0x50,
	0x11, 0x17, 0x2a, 0xd7, 0xfe, 0xd4, 0x56, 0xda, 0xce, 0xb2, 0xdb, 0x13, 0xa3, 0xc6, 0x34, 0x0c,
	0xed, 0xdf, 0xda, 0xa2, 0x8a, 0x0c, 0xed, 0xdf, 0x52, 0x63, 0x92, 0x1e, 0xd4, 0x55, 0x36, 0x3c,
	0x5b, 0xc7, 0x49, 0x81, 0xb3, 0x83, 0xa5, 0x39, 0x64, 0xd6, 0x1c, 0xf9, 0x53, 0xa7, 0xba, 0x63,
	0xcd, 0x91, 0xd9, 0xd7, 0x9c, 0x7d, 0x0f, 0xea, 0x22, 0x9b, 0xac, 0xb3, 0xb7, 0x63, 0x4d, 0x3b,
	0x75, 0x9a, 0x43, 0xee, 0xcf, 0x2a, 0xd4, 0xad, 0x47, 0x5a, 0x50, 0xf6, 0x06, 0x69, 0x5b, 0x0d,
	0x5a, 0xf6, 0x06, 0xe4, 0x7f, 0xa8, 0xc7, 0x32, 0x44, 0xe9, 0x0d, 0x6c, 0x1f, 0xcd, 0x6c, 0xad,
	0x8f, 0x7c, 0xee, 0x09, 0x4d, 0x73, 0x93, 0xfc, 0x07, 0xb5, 0x10, 0x59, 0xe4, 0x0d, 0x6c, 0x1b,
	0x45, 0xcc, 0x7a, 0x66, 0xec, 0xe1, 0x5a, 0x32, 0xcd, 0x63, 0x61, 0x5b, 0xc8, 0xef, 0x88, 0x55,
	0xe9, 0xc6, 0x27, 0xff, 0xc0, 0x5e, 0x22, 0x79, 0x80, 0xb6, 0x87, 0x83, 0x0c, 0xf4, 0x8d, 0x44,
	0x33, 0x87, 0xf4, 0xa0, 0x31, 0x8b, 0x58, 0xb0, 0x8c, 0xb8, 0xd2, 0x4e, 0xed, 0xf9, 0x48, 0x86,
	0x7a, 0xd1, 0x0f, 0x43, 0x89, 0x4a, 0xd1, 0x2d, 0x42, 0xae, 0xa0, 0x19, 0xc4, 0x6b, 0xa1, 0x51,
	0x26, 0x4c, 0xea, 0x47, 0xa7, 0xfe, 0x42, 0xa4, 0x40, 0x91, 0x0b, 0xd8, 0xe7, 0x21, 0x0a, 0xcd,
	0xf5, 0xa3, 0xb3, 0xdf, 0x29, 0x75, 0x5b, 0x97, 0xc7, 0x59, 0xc2, 0xb3, 0xea, 0x0d, 0x7e, 0xc1,
	0x88, 0x6e, 0x20, 0xd2, 0x86, 0x8a, 0x66, 0x73, 0xa7, 0xd1, 0x29, 0x75, 0x9b, 0xd4, 0x7c, 0x92,
	0x2b, 0x68, 0xc8, 0xfc, 0xea, 0x38, 0x90, 0xee, 0x7a, 0x5a, 0xbc, 0x0f, 0xb9, 0x4b, 0xb7, 0x20,
	0x79, 0x05, 0x35, 0xa5, 0x99, 0x5e, 0x2b, 0xe7, 0x20, 0xdd, 0xb6, 0x78, 0x8c, 0xbd, 0xfb, 0xd4,
	0xa3, 0x96, 0x21, 0x17, 0x00, 0x81, 0x44, 0xa6, 0x71, 0xc2, 0x57, 0xe8, 0x34, 0xd3, 0x4d, 0x8e,
	0xb2, 0xc4, 0x24, 0x7f, 0x5d, 0xe8, 0x33, 0x84, 0xf4, 0xe1, 0x38, 0x62, 0x4a, 0xdf, 0x99, 0x13,
	0xf4, 0xcd, 0x63, 0x14, 0xa6, 0xc9, 0xc3, 0xdd, 0xc9, 0x5d, 0xac, 0x7b, 0x0e, 0xb5, 0xac, 0x0a,
	0x02, 0x50, 0xeb, 0x5f, 0x4f, 0xbc, 0xcf, 0xc3, 0xf6, 0x5f, 0xe4, 0x04, 0xda, 0xfe, 0x70, 0x3c,
	0xf0, 0xc6, 0xa3, 0x87, 0xc1, 0xf0, 0x66, 0x38, 0xf1, 0xee, 0xc6, 0xed, 0xd2, 0xac, 0x96, 0xbe,
	0x6c, 0x6f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x87, 0x32, 0x79, 0x4d, 0x48, 0x05, 0x00, 0x00,
}
