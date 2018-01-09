// Code generated by protoc-gen-go.
// source: zebra.proto
// DO NOT EDIT!

/*
Package zebra is a generated protocol buffer package.

It is generated from these files:
	zebra.proto

It has these top-level messages:
	InterfaceRequest
	InterfaceUpdate
	RouterIdRequest
	RouterIdUpdate
	RedistRequest
	Prefix
	Nexthop
	Address
	Route
*/
package zebra

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AFI int32

const (
	AFI_AFI_IP  AFI = 0
	AFI_AFI_IP6 AFI = 1
	AFI_AFI_MAX AFI = 2
)

var AFI_name = map[int32]string{
	0: "AFI_IP",
	1: "AFI_IP6",
	2: "AFI_MAX",
}
var AFI_value = map[string]int32{
	"AFI_IP":  0,
	"AFI_IP6": 1,
	"AFI_MAX": 2,
}

func (x AFI) String() string {
	return proto.EnumName(AFI_name, int32(x))
}
func (AFI) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RouteType int32

const (
	RouteType_RIB_UNKNOWN   RouteType = 0
	RouteType_RIB_KERNEL    RouteType = 1
	RouteType_RIB_CONNECTED RouteType = 2
	RouteType_RIB_STATIC    RouteType = 3
	RouteType_RIB_RIP       RouteType = 4
	RouteType_RIB_OSPF      RouteType = 5
	RouteType_RIB_ISIS      RouteType = 6
	RouteType_RIB_BGP       RouteType = 7
)

var RouteType_name = map[int32]string{
	0: "RIB_UNKNOWN",
	1: "RIB_KERNEL",
	2: "RIB_CONNECTED",
	3: "RIB_STATIC",
	4: "RIB_RIP",
	5: "RIB_OSPF",
	6: "RIB_ISIS",
	7: "RIB_BGP",
}
var RouteType_value = map[string]int32{
	"RIB_UNKNOWN":   0,
	"RIB_KERNEL":    1,
	"RIB_CONNECTED": 2,
	"RIB_STATIC":    3,
	"RIB_RIP":       4,
	"RIB_OSPF":      5,
	"RIB_ISIS":      6,
	"RIB_BGP":       7,
}

func (x RouteType) String() string {
	return proto.EnumName(RouteType_name, int32(x))
}
func (RouteType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type RouteSubType int32

const (
	RouteSubType_RIB_SUB_OSPF_IA         RouteSubType = 0
	RouteSubType_RIB_SUB_OSPF_NSSA_1     RouteSubType = 1
	RouteSubType_RIB_SUB_OSPF_NSSA_2     RouteSubType = 2
	RouteSubType_RIB_SUB_OSPF_EXTERNAL_1 RouteSubType = 3
	RouteSubType_RIB_SUB_OSPF_EXTERNAL_2 RouteSubType = 4
	RouteSubType_RIB_SUB_BGP_IBGP        RouteSubType = 5
	RouteSubType_RIB_SUB_BGP_EBGP        RouteSubType = 6
	RouteSubType_RIB_SUB_BGP_CONFED      RouteSubType = 7
	RouteSubType_RIB_SUB_ISIS_L1         RouteSubType = 8
	RouteSubType_RIB_SUB_ISIS_L2         RouteSubType = 9
	RouteSubType_RIB_SUB_ISIS_IA         RouteSubType = 10
)

var RouteSubType_name = map[int32]string{
	0:  "RIB_SUB_OSPF_IA",
	1:  "RIB_SUB_OSPF_NSSA_1",
	2:  "RIB_SUB_OSPF_NSSA_2",
	3:  "RIB_SUB_OSPF_EXTERNAL_1",
	4:  "RIB_SUB_OSPF_EXTERNAL_2",
	5:  "RIB_SUB_BGP_IBGP",
	6:  "RIB_SUB_BGP_EBGP",
	7:  "RIB_SUB_BGP_CONFED",
	8:  "RIB_SUB_ISIS_L1",
	9:  "RIB_SUB_ISIS_L2",
	10: "RIB_SUB_ISIS_IA",
}
var RouteSubType_value = map[string]int32{
	"RIB_SUB_OSPF_IA":         0,
	"RIB_SUB_OSPF_NSSA_1":     1,
	"RIB_SUB_OSPF_NSSA_2":     2,
	"RIB_SUB_OSPF_EXTERNAL_1": 3,
	"RIB_SUB_OSPF_EXTERNAL_2": 4,
	"RIB_SUB_BGP_IBGP":        5,
	"RIB_SUB_BGP_EBGP":        6,
	"RIB_SUB_BGP_CONFED":      7,
	"RIB_SUB_ISIS_L1":         8,
	"RIB_SUB_ISIS_L2":         9,
	"RIB_SUB_ISIS_IA":         10,
}

func (x RouteSubType) String() string {
	return proto.EnumName(RouteSubType_name, int32(x))
}
func (RouteSubType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Op int32

const (
	Op_NoOperation              Op = 0
	Op_InterfaceSubscribe       Op = 1
	Op_InterfaceUnsubscribe     Op = 2
	Op_RouterIdSubscribe        Op = 3
	Op_RouterIdUnsubscribe      Op = 4
	Op_RedistSubscribe          Op = 5
	Op_RedistUnsubscribe        Op = 6
	Op_RedistDefaultSubscribe   Op = 7
	Op_RedistDefaultUnsubscribe Op = 8
	Op_RouteAdd                 Op = 9
	Op_RouteDelete              Op = 10
	Op_InterfaceAdd             Op = 11
	Op_InterfaceDelete          Op = 12
	Op_InterfaceAddrAdd         Op = 13
	Op_InterfaceAddrDelete      Op = 14
	Op_InterfaceUp              Op = 15
	Op_InterfaceDown            Op = 16
	Op_InterfaceFlagChange      Op = 17
	Op_InterfaceNameChange      Op = 18
	Op_InterfaceMtuChange       Op = 19
)

var Op_name = map[int32]string{
	0:  "NoOperation",
	1:  "InterfaceSubscribe",
	2:  "InterfaceUnsubscribe",
	3:  "RouterIdSubscribe",
	4:  "RouterIdUnsubscribe",
	5:  "RedistSubscribe",
	6:  "RedistUnsubscribe",
	7:  "RedistDefaultSubscribe",
	8:  "RedistDefaultUnsubscribe",
	9:  "RouteAdd",
	10: "RouteDelete",
	11: "InterfaceAdd",
	12: "InterfaceDelete",
	13: "InterfaceAddrAdd",
	14: "InterfaceAddrDelete",
	15: "InterfaceUp",
	16: "InterfaceDown",
	17: "InterfaceFlagChange",
	18: "InterfaceNameChange",
	19: "InterfaceMtuChange",
}
var Op_value = map[string]int32{
	"NoOperation":              0,
	"InterfaceSubscribe":       1,
	"InterfaceUnsubscribe":     2,
	"RouterIdSubscribe":        3,
	"RouterIdUnsubscribe":      4,
	"RedistSubscribe":          5,
	"RedistUnsubscribe":        6,
	"RedistDefaultSubscribe":   7,
	"RedistDefaultUnsubscribe": 8,
	"RouteAdd":                 9,
	"RouteDelete":              10,
	"InterfaceAdd":             11,
	"InterfaceDelete":          12,
	"InterfaceAddrAdd":         13,
	"InterfaceAddrDelete":      14,
	"InterfaceUp":              15,
	"InterfaceDown":            16,
	"InterfaceFlagChange":      17,
	"InterfaceNameChange":      18,
	"InterfaceMtuChange":       19,
}

func (x Op) String() string {
	return proto.EnumName(Op_name, int32(x))
}
func (Op) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type InterfaceRequest struct {
	Op    Op     `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId uint32 `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
}

func (m *InterfaceRequest) Reset()                    { *m = InterfaceRequest{} }
func (m *InterfaceRequest) String() string            { return proto.CompactTextString(m) }
func (*InterfaceRequest) ProtoMessage()               {}
func (*InterfaceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InterfaceRequest) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *InterfaceRequest) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

type InterfaceUpdate struct {
	Op       Op         `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId    uint32     `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Index    uint32     `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Mtu      uint32     `protobuf:"varint,5,opt,name=mtu" json:"mtu,omitempty"`
	Metric   uint32     `protobuf:"varint,6,opt,name=metric" json:"metric,omitempty"`
	AddrIpv4 []*Address `protobuf:"bytes,7,rep,name=addr_ipv4,json=addrIpv4" json:"addr_ipv4,omitempty"`
	AddrIpv6 []*Address `protobuf:"bytes,8,rep,name=addr_ipv6,json=addrIpv6" json:"addr_ipv6,omitempty"`
}

func (m *InterfaceUpdate) Reset()                    { *m = InterfaceUpdate{} }
func (m *InterfaceUpdate) String() string            { return proto.CompactTextString(m) }
func (*InterfaceUpdate) ProtoMessage()               {}
func (*InterfaceUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InterfaceUpdate) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *InterfaceUpdate) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *InterfaceUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceUpdate) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InterfaceUpdate) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfaceUpdate) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *InterfaceUpdate) GetAddrIpv4() []*Address {
	if m != nil {
		return m.AddrIpv4
	}
	return nil
}

func (m *InterfaceUpdate) GetAddrIpv6() []*Address {
	if m != nil {
		return m.AddrIpv6
	}
	return nil
}

type RouterIdRequest struct {
	Op    Op     `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId uint32 `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
}

func (m *RouterIdRequest) Reset()                    { *m = RouterIdRequest{} }
func (m *RouterIdRequest) String() string            { return proto.CompactTextString(m) }
func (*RouterIdRequest) ProtoMessage()               {}
func (*RouterIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RouterIdRequest) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *RouterIdRequest) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

type RouterIdUpdate struct {
	VrfId    uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	RouterId []byte `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
}

func (m *RouterIdUpdate) Reset()                    { *m = RouterIdUpdate{} }
func (m *RouterIdUpdate) String() string            { return proto.CompactTextString(m) }
func (*RouterIdUpdate) ProtoMessage()               {}
func (*RouterIdUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouterIdUpdate) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *RouterIdUpdate) GetRouterId() []byte {
	if m != nil {
		return m.RouterId
	}
	return nil
}

type RedistRequest struct {
	Op    Op        `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	Afi   AFI       `protobuf:"varint,2,opt,name=afi,enum=zebra.AFI" json:"afi,omitempty"`
	VrfId uint32    `protobuf:"varint,3,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	Type  RouteType `protobuf:"varint,4,opt,name=type,enum=zebra.RouteType" json:"type,omitempty"`
}

func (m *RedistRequest) Reset()                    { *m = RedistRequest{} }
func (m *RedistRequest) String() string            { return proto.CompactTextString(m) }
func (*RedistRequest) ProtoMessage()               {}
func (*RedistRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RedistRequest) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *RedistRequest) GetAfi() AFI {
	if m != nil {
		return m.Afi
	}
	return AFI_AFI_IP
}

func (m *RedistRequest) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *RedistRequest) GetType() RouteType {
	if m != nil {
		return m.Type
	}
	return RouteType_RIB_UNKNOWN
}

type Prefix struct {
	Addr   []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Length uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
}

func (m *Prefix) Reset()                    { *m = Prefix{} }
func (m *Prefix) String() string            { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()               {}
func (*Prefix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Prefix) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Prefix) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Nexthop struct {
	Addr    []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Ifindex uint32 `protobuf:"varint,2,opt,name=ifindex" json:"ifindex,omitempty"`
}

func (m *Nexthop) Reset()                    { *m = Nexthop{} }
func (m *Nexthop) String() string            { return proto.CompactTextString(m) }
func (*Nexthop) ProtoMessage()               {}
func (*Nexthop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Nexthop) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Nexthop) GetIfindex() uint32 {
	if m != nil {
		return m.Ifindex
	}
	return 0
}

type Address struct {
	Addr  *Prefix `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Flags uint32  `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Address) GetAddr() *Prefix {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Address) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type Route struct {
	Op       Op           `protobuf:"varint,1,opt,name=op,enum=zebra.Op" json:"op,omitempty"`
	VrfId    uint32       `protobuf:"varint,2,opt,name=vrf_id,json=vrfId" json:"vrf_id,omitempty"`
	Prefix   *Prefix      `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	Type     RouteType    `protobuf:"varint,4,opt,name=type,enum=zebra.RouteType" json:"type,omitempty"`
	SubType  RouteSubType `protobuf:"varint,5,opt,name=sub_type,json=subType,enum=zebra.RouteSubType" json:"sub_type,omitempty"`
	Distance uint32       `protobuf:"varint,6,opt,name=distance" json:"distance,omitempty"`
	Metric   uint32       `protobuf:"varint,7,opt,name=metric" json:"metric,omitempty"`
	Tag      uint32       `protobuf:"varint,8,opt,name=tag" json:"tag,omitempty"`
	Color    []string     `protobuf:"bytes,9,rep,name=color" json:"color,omitempty"`
	Nexthops []*Nexthop   `protobuf:"bytes,10,rep,name=nexthops" json:"nexthops,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Route) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_NoOperation
}

func (m *Route) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Route) GetPrefix() *Prefix {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Route) GetType() RouteType {
	if m != nil {
		return m.Type
	}
	return RouteType_RIB_UNKNOWN
}

func (m *Route) GetSubType() RouteSubType {
	if m != nil {
		return m.SubType
	}
	return RouteSubType_RIB_SUB_OSPF_IA
}

func (m *Route) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Route) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Route) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Route) GetColor() []string {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *Route) GetNexthops() []*Nexthop {
	if m != nil {
		return m.Nexthops
	}
	return nil
}

func init() {
	proto.RegisterType((*InterfaceRequest)(nil), "zebra.InterfaceRequest")
	proto.RegisterType((*InterfaceUpdate)(nil), "zebra.InterfaceUpdate")
	proto.RegisterType((*RouterIdRequest)(nil), "zebra.RouterIdRequest")
	proto.RegisterType((*RouterIdUpdate)(nil), "zebra.RouterIdUpdate")
	proto.RegisterType((*RedistRequest)(nil), "zebra.RedistRequest")
	proto.RegisterType((*Prefix)(nil), "zebra.Prefix")
	proto.RegisterType((*Nexthop)(nil), "zebra.Nexthop")
	proto.RegisterType((*Address)(nil), "zebra.Address")
	proto.RegisterType((*Route)(nil), "zebra.Route")
	proto.RegisterEnum("zebra.AFI", AFI_name, AFI_value)
	proto.RegisterEnum("zebra.RouteType", RouteType_name, RouteType_value)
	proto.RegisterEnum("zebra.RouteSubType", RouteSubType_name, RouteSubType_value)
	proto.RegisterEnum("zebra.Op", Op_name, Op_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Zebra service

type ZebraClient interface {
	InterfaceService(ctx context.Context, opts ...grpc.CallOption) (Zebra_InterfaceServiceClient, error)
	RouterIdService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouterIdServiceClient, error)
	RedistService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RedistServiceClient, error)
	RouteIPv4Service(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouteIPv4ServiceClient, error)
	RouteIPv6Service(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouteIPv6ServiceClient, error)
}

type zebraClient struct {
	cc *grpc.ClientConn
}

func NewZebraClient(cc *grpc.ClientConn) ZebraClient {
	return &zebraClient{cc}
}

func (c *zebraClient) InterfaceService(ctx context.Context, opts ...grpc.CallOption) (Zebra_InterfaceServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[0], c.cc, "/zebra.Zebra/InterfaceService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraInterfaceServiceClient{stream}
	return x, nil
}

type Zebra_InterfaceServiceClient interface {
	Send(*InterfaceRequest) error
	Recv() (*InterfaceUpdate, error)
	grpc.ClientStream
}

type zebraInterfaceServiceClient struct {
	grpc.ClientStream
}

func (x *zebraInterfaceServiceClient) Send(m *InterfaceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraInterfaceServiceClient) Recv() (*InterfaceUpdate, error) {
	m := new(InterfaceUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RouterIdService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouterIdServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[1], c.cc, "/zebra.Zebra/RouterIdService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRouterIdServiceClient{stream}
	return x, nil
}

type Zebra_RouterIdServiceClient interface {
	Send(*RouterIdRequest) error
	Recv() (*RouterIdUpdate, error)
	grpc.ClientStream
}

type zebraRouterIdServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRouterIdServiceClient) Send(m *RouterIdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRouterIdServiceClient) Recv() (*RouterIdUpdate, error) {
	m := new(RouterIdUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RedistService(ctx context.Context, opts ...grpc.CallOption) (Zebra_RedistServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[2], c.cc, "/zebra.Zebra/RedistService", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRedistServiceClient{stream}
	return x, nil
}

type Zebra_RedistServiceClient interface {
	Send(*RedistRequest) error
	Recv() (*Route, error)
	grpc.ClientStream
}

type zebraRedistServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRedistServiceClient) Send(m *RedistRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRedistServiceClient) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RouteIPv4Service(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouteIPv4ServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[3], c.cc, "/zebra.Zebra/RouteIPv4Service", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRouteIPv4ServiceClient{stream}
	return x, nil
}

type Zebra_RouteIPv4ServiceClient interface {
	Send(*Route) error
	Recv() (*Route, error)
	grpc.ClientStream
}

type zebraRouteIPv4ServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRouteIPv4ServiceClient) Send(m *Route) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRouteIPv4ServiceClient) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *zebraClient) RouteIPv6Service(ctx context.Context, opts ...grpc.CallOption) (Zebra_RouteIPv6ServiceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Zebra_serviceDesc.Streams[4], c.cc, "/zebra.Zebra/RouteIPv6Service", opts...)
	if err != nil {
		return nil, err
	}
	x := &zebraRouteIPv6ServiceClient{stream}
	return x, nil
}

type Zebra_RouteIPv6ServiceClient interface {
	Send(*Route) error
	Recv() (*Route, error)
	grpc.ClientStream
}

type zebraRouteIPv6ServiceClient struct {
	grpc.ClientStream
}

func (x *zebraRouteIPv6ServiceClient) Send(m *Route) error {
	return x.ClientStream.SendMsg(m)
}

func (x *zebraRouteIPv6ServiceClient) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Zebra service

type ZebraServer interface {
	InterfaceService(Zebra_InterfaceServiceServer) error
	RouterIdService(Zebra_RouterIdServiceServer) error
	RedistService(Zebra_RedistServiceServer) error
	RouteIPv4Service(Zebra_RouteIPv4ServiceServer) error
	RouteIPv6Service(Zebra_RouteIPv6ServiceServer) error
}

func RegisterZebraServer(s *grpc.Server, srv ZebraServer) {
	s.RegisterService(&_Zebra_serviceDesc, srv)
}

func _Zebra_InterfaceService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).InterfaceService(&zebraInterfaceServiceServer{stream})
}

type Zebra_InterfaceServiceServer interface {
	Send(*InterfaceUpdate) error
	Recv() (*InterfaceRequest, error)
	grpc.ServerStream
}

type zebraInterfaceServiceServer struct {
	grpc.ServerStream
}

func (x *zebraInterfaceServiceServer) Send(m *InterfaceUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraInterfaceServiceServer) Recv() (*InterfaceRequest, error) {
	m := new(InterfaceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RouterIdService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RouterIdService(&zebraRouterIdServiceServer{stream})
}

type Zebra_RouterIdServiceServer interface {
	Send(*RouterIdUpdate) error
	Recv() (*RouterIdRequest, error)
	grpc.ServerStream
}

type zebraRouterIdServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRouterIdServiceServer) Send(m *RouterIdUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRouterIdServiceServer) Recv() (*RouterIdRequest, error) {
	m := new(RouterIdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RedistService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RedistService(&zebraRedistServiceServer{stream})
}

type Zebra_RedistServiceServer interface {
	Send(*Route) error
	Recv() (*RedistRequest, error)
	grpc.ServerStream
}

type zebraRedistServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRedistServiceServer) Send(m *Route) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRedistServiceServer) Recv() (*RedistRequest, error) {
	m := new(RedistRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RouteIPv4Service_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RouteIPv4Service(&zebraRouteIPv4ServiceServer{stream})
}

type Zebra_RouteIPv4ServiceServer interface {
	Send(*Route) error
	Recv() (*Route, error)
	grpc.ServerStream
}

type zebraRouteIPv4ServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRouteIPv4ServiceServer) Send(m *Route) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRouteIPv4ServiceServer) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Zebra_RouteIPv6Service_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ZebraServer).RouteIPv6Service(&zebraRouteIPv6ServiceServer{stream})
}

type Zebra_RouteIPv6ServiceServer interface {
	Send(*Route) error
	Recv() (*Route, error)
	grpc.ServerStream
}

type zebraRouteIPv6ServiceServer struct {
	grpc.ServerStream
}

func (x *zebraRouteIPv6ServiceServer) Send(m *Route) error {
	return x.ServerStream.SendMsg(m)
}

func (x *zebraRouteIPv6ServiceServer) Recv() (*Route, error) {
	m := new(Route)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Zebra_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zebra.Zebra",
	HandlerType: (*ZebraServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InterfaceService",
			Handler:       _Zebra_InterfaceService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RouterIdService",
			Handler:       _Zebra_RouterIdService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RedistService",
			Handler:       _Zebra_RedistService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteIPv4Service",
			Handler:       _Zebra_RouteIPv4Service_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteIPv6Service",
			Handler:       _Zebra_RouteIPv6Service_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "zebra.proto",
}

func init() { proto.RegisterFile("zebra.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 990 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x36, 0x49, 0xf1, 0xa0, 0x91, 0x2c, 0xaf, 0xd7, 0x27, 0xfe, 0x4e, 0x2e, 0xfc, 0x0b, 0x2d,
	0x20, 0x28, 0x80, 0x51, 0xab, 0x82, 0x8b, 0x5e, 0xd2, 0x3a, 0x04, 0x44, 0x1c, 0x4a, 0xa0, 0x64,
	0x34, 0xe8, 0x0d, 0x41, 0x89, 0x2b, 0x99, 0x80, 0x4c, 0xb2, 0x24, 0xa5, 0x3a, 0xbd, 0xcf, 0x0b,
	0xf4, 0x81, 0x8a, 0xf6, 0x91, 0xfa, 0x06, 0xc5, 0x2e, 0x49, 0x69, 0x69, 0xa7, 0xad, 0x9b, 0x3b,
	0xce, 0x37, 0xdf, 0x37, 0x3b, 0x3b, 0xb3, 0xbb, 0x43, 0xa8, 0xfd, 0x42, 0x66, 0xb1, 0x7b, 0x19,
	0xc5, 0x61, 0x1a, 0x62, 0x99, 0x19, 0xcd, 0x3e, 0x20, 0x33, 0x48, 0x49, 0xbc, 0x70, 0xe7, 0xc4,
	0x26, 0x3f, 0xad, 0x49, 0x92, 0xe2, 0xff, 0x81, 0x18, 0x46, 0xba, 0x70, 0x21, 0xb4, 0x1a, 0x9d,
	0xea, 0x65, 0x26, 0x1a, 0x45, 0xb6, 0x18, 0x46, 0xf8, 0x04, 0x94, 0x4d, 0xbc, 0x70, 0x7c, 0x4f,
	0x17, 0x2f, 0x84, 0xd6, 0xbe, 0x2d, 0x6f, 0xe2, 0x85, 0xe9, 0x35, 0xff, 0x14, 0xe0, 0x60, 0x1b,
	0xe6, 0x2e, 0xf2, 0xdc, 0x94, 0xfc, 0xf7, 0x28, 0x18, 0x43, 0x25, 0x70, 0x1f, 0x88, 0x2e, 0x5d,
	0x08, 0xad, 0xaa, 0xcd, 0xbe, 0xf1, 0x31, 0xc8, 0x7e, 0xe0, 0x91, 0x47, 0xbd, 0x92, 0x31, 0x99,
	0x81, 0x11, 0x48, 0x0f, 0xe9, 0x5a, 0x97, 0x19, 0x46, 0x3f, 0xf1, 0x29, 0x28, 0x0f, 0x24, 0x8d,
	0xfd, 0xb9, 0xae, 0x30, 0x30, 0xb7, 0xf0, 0x1b, 0xa8, 0xba, 0x9e, 0x17, 0x3b, 0x7e, 0xb4, 0xe9,
	0xea, 0xea, 0x85, 0xd4, 0xaa, 0x75, 0x1a, 0x79, 0x32, 0x86, 0xe7, 0xc5, 0x24, 0x49, 0x6c, 0x8d,
	0x12, 0xcc, 0x68, 0xd3, 0xe5, 0xc9, 0xd7, 0xba, 0xf6, 0x8f, 0xe4, 0xeb, 0x66, 0x0f, 0x0e, 0xec,
	0x70, 0x9d, 0x92, 0xd8, 0xf4, 0xbe, 0xbc, 0x70, 0x7d, 0x68, 0x14, 0x41, 0xf2, 0xb2, 0xed, 0x88,
	0x02, 0x5f, 0x9b, 0x57, 0x50, 0x8d, 0x19, 0xb1, 0x08, 0x51, 0xb7, 0xb5, 0x38, 0x57, 0x36, 0x3f,
	0x09, 0xb0, 0x6f, 0x13, 0xcf, 0x4f, 0xd2, 0x17, 0x64, 0xf2, 0x1a, 0x24, 0x77, 0xe1, 0xb3, 0x18,
	0x8d, 0x0e, 0x14, 0xdb, 0x1b, 0x9a, 0x36, 0x85, 0xb9, 0xe5, 0x25, 0x7e, 0xf9, 0xaf, 0xa0, 0x92,
	0x7e, 0x8c, 0x08, 0xeb, 0x42, 0xa3, 0x83, 0x72, 0x15, 0x4b, 0x7d, 0xfa, 0x31, 0x22, 0x36, 0xf3,
	0x36, 0xbb, 0xa0, 0x8c, 0x63, 0xb2, 0xf0, 0x1f, 0x69, 0x2b, 0x69, 0xa1, 0x58, 0x06, 0x75, 0x9b,
	0x7d, 0xd3, 0x16, 0xad, 0x48, 0xb0, 0x4c, 0xef, 0xf3, 0x12, 0xe4, 0x56, 0xf3, 0x3b, 0x50, 0x2d,
	0xf2, 0x98, 0xde, 0x87, 0xd1, 0x67, 0x65, 0x3a, 0xa8, 0xfe, 0x22, 0x3b, 0x03, 0x99, 0xae, 0x30,
	0x9b, 0x37, 0xa0, 0xe6, 0x6d, 0xc1, 0xff, 0xe7, 0x84, 0xb5, 0xce, 0x7e, 0x9e, 0x5f, 0x96, 0x4c,
	0x1e, 0xe7, 0x18, 0xe4, 0xc5, 0xca, 0x5d, 0x26, 0x45, 0x03, 0x98, 0xd1, 0xfc, 0x4d, 0x04, 0x99,
	0x6d, 0xe3, 0x0b, 0xce, 0xeb, 0xd7, 0xa0, 0x44, 0x6c, 0x05, 0x56, 0xab, 0x67, 0xcb, 0xe6, 0xce,
	0x97, 0xd5, 0x0e, 0x5f, 0x82, 0x96, 0xac, 0x67, 0x0e, 0x63, 0xca, 0x8c, 0x79, 0xc4, 0x33, 0x27,
	0xeb, 0x19, 0x23, 0xab, 0x49, 0xf6, 0x81, 0xcf, 0x41, 0xa3, 0x0d, 0x77, 0x83, 0x39, 0xc9, 0x8f,
	0xfc, 0xd6, 0xe6, 0x2e, 0x83, 0x5a, 0xba, 0x0c, 0x08, 0xa4, 0xd4, 0x5d, 0xea, 0x5a, 0x76, 0x6d,
	0x52, 0x77, 0x49, 0x8b, 0x32, 0x0f, 0x57, 0x61, 0xac, 0x57, 0x2f, 0xa4, 0x56, 0xd5, 0xce, 0x0c,
	0xdc, 0x06, 0x2d, 0xc8, 0x3a, 0x92, 0xe8, 0x50, 0xba, 0x06, 0x79, 0xa3, 0xec, 0xad, 0xbf, 0xfd,
	0x06, 0x24, 0x63, 0x68, 0x62, 0x00, 0xc5, 0x18, 0x9a, 0x8e, 0x39, 0x46, 0x7b, 0xb8, 0x06, 0x6a,
	0xf6, 0x7d, 0x8d, 0x84, 0xc2, 0x78, 0x6f, 0x7c, 0x40, 0x62, 0xfb, 0x93, 0x00, 0xd5, 0xed, 0xc6,
	0xf1, 0x01, 0xd4, 0x6c, 0xf3, 0xc6, 0xb9, 0xb3, 0xde, 0x59, 0xa3, 0x1f, 0x2c, 0xb4, 0x87, 0x1b,
	0x00, 0x14, 0x78, 0x37, 0xb0, 0xad, 0xc1, 0x2d, 0x12, 0xf0, 0x21, 0xec, 0x53, 0xbb, 0x37, 0xb2,
	0xac, 0x41, 0x6f, 0x3a, 0xe8, 0x23, 0xb1, 0xa0, 0x4c, 0xa6, 0xc6, 0xd4, 0xec, 0x21, 0x89, 0x86,
	0xa7, 0xb6, 0x6d, 0x8e, 0x51, 0x05, 0xd7, 0x41, 0xa3, 0xc6, 0x68, 0x32, 0x1e, 0x22, 0xb9, 0xb0,
	0xcc, 0x89, 0x39, 0x41, 0x4a, 0x41, 0xbc, 0x79, 0x3b, 0x46, 0x6a, 0xfb, 0x57, 0x11, 0xea, 0x7c,
	0x59, 0xf1, 0x11, 0x1c, 0xb0, 0xb0, 0x77, 0x99, 0xda, 0x31, 0x0d, 0xb4, 0x87, 0xcf, 0xe0, 0xa8,
	0x04, 0x5a, 0x93, 0x89, 0xe1, 0x5c, 0x21, 0xe1, 0xf3, 0x8e, 0x0e, 0x12, 0xf1, 0x2b, 0x38, 0x2b,
	0x39, 0x06, 0x1f, 0xa6, 0x03, 0xdb, 0x32, 0x6e, 0x9d, 0x2b, 0x24, 0xfd, 0xbd, 0xb3, 0x83, 0x2a,
	0xf8, 0x18, 0x50, 0xe1, 0xbc, 0x79, 0x3b, 0x76, 0x4c, 0x9a, 0xa7, 0xfc, 0x14, 0x1d, 0x50, 0x54,
	0xc1, 0xa7, 0x80, 0x79, 0xb4, 0x37, 0xb2, 0x86, 0x83, 0x3e, 0x52, 0xf9, 0x4d, 0xd0, 0x4d, 0x3b,
	0xb7, 0x57, 0x48, 0x7b, 0x0e, 0x76, 0x50, 0xf5, 0x19, 0x68, 0x1a, 0x08, 0xda, 0xbf, 0x4b, 0x20,
	0x8e, 0x22, 0xda, 0x15, 0x2b, 0x1c, 0x45, 0x24, 0x76, 0x53, 0x3f, 0x0c, 0xd0, 0x1e, 0x5d, 0x6e,
	0xfb, 0xb6, 0x4f, 0xd6, 0xb3, 0x64, 0x1e, 0xfb, 0x33, 0x82, 0x04, 0xac, 0xc3, 0xf1, 0xee, 0xcd,
	0x0f, 0x92, 0xad, 0x47, 0xc4, 0x27, 0x70, 0x58, 0xbc, 0x6a, 0x3b, 0x81, 0xc4, 0xca, 0x56, 0x3c,
	0x76, 0x1c, 0xbf, 0xc2, 0xd2, 0x61, 0xcf, 0xd7, 0x8e, 0x2d, 0xb3, 0x20, 0x0c, 0xe4, 0xb9, 0x0a,
	0x3e, 0x87, 0xd3, 0x0c, 0xee, 0x93, 0x85, 0xbb, 0x5e, 0x71, 0x12, 0x15, 0xbf, 0x06, 0xbd, 0xe4,
	0xe3, 0x95, 0x1a, 0x3b, 0x0f, 0x74, 0x79, 0xc3, 0xf3, 0x50, 0x95, 0x1d, 0x3e, 0x6a, 0xf5, 0xc9,
	0x8a, 0xa4, 0x04, 0x01, 0x46, 0x50, 0xdf, 0x6e, 0x87, 0x52, 0x6a, 0x34, 0xad, 0x2d, 0x92, 0xd3,
	0xea, 0xb4, 0x25, 0x3c, 0x2d, 0xa6, 0xd4, 0x7d, 0xba, 0xb5, 0x12, 0x9a, 0xd3, 0x1b, 0x74, 0x19,
	0x6e, 0x30, 0xa2, 0x03, 0x7a, 0xa6, 0x77, 0x41, 0xc3, 0x9f, 0x03, 0x84, 0x4a, 0xe2, 0xe1, 0xca,
	0x5d, 0xf6, 0xee, 0xdd, 0x60, 0x49, 0xd0, 0x61, 0xc9, 0x61, 0xb9, 0x0f, 0x24, 0x77, 0xe0, 0x52,
	0x4b, 0xde, 0xa7, 0xeb, 0x1c, 0x3f, 0xea, 0xfc, 0x21, 0x82, 0xfc, 0x23, 0xbd, 0xa8, 0xd8, 0xe4,
	0xd2, 0x9c, 0x90, 0x78, 0xe3, 0xcf, 0x09, 0x3e, 0xcb, 0x2f, 0xf1, 0xd3, 0x81, 0x7f, 0x7e, 0xfa,
	0xd4, 0x91, 0xcd, 0xa2, 0xe6, 0x5e, 0x4b, 0xf8, 0x46, 0xc0, 0xc3, 0xdd, 0xa0, 0x2b, 0x22, 0x9d,
	0xf2, 0x4f, 0xd3, 0x6e, 0x00, 0x9e, 0x9f, 0x3c, 0xc1, 0x4b, 0x71, 0xbe, 0x2f, 0x86, 0x54, 0x11,
	0xe5, 0xb8, 0x60, 0xf3, 0xa3, 0xeb, 0xbc, 0xce, 0xc7, 0xc8, 0xa5, 0x5d, 0x40, 0xcc, 0x34, 0xc7,
	0x9b, 0x6e, 0xa1, 0x2e, 0xf1, 0xfe, 0x4d, 0x75, 0xfd, 0x62, 0xd5, 0x4c, 0x61, 0xff, 0x47, 0xdf,
	0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xcb, 0x28, 0xfe, 0x2e, 0x09, 0x00, 0x00,
}
