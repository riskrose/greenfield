// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/sp/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventCreateStorageProvider is emitted when there is a storage provider created
type EventCreateStorageProvider struct {
	// sp_address is the operator address of the storage provider
	SpAddress string `protobuf:"bytes,1,opt,name=sp_address,json=spAddress,proto3" json:"sp_address,omitempty"`
	// funding_address is the funding account address of the storage provider
	FundingAddress string `protobuf:"bytes,2,opt,name=funding_address,json=fundingAddress,proto3" json:"funding_address,omitempty"`
	// seal_address is the account address for SealObject Tx
	SealAddress string `protobuf:"bytes,3,opt,name=seal_address,json=sealAddress,proto3" json:"seal_address,omitempty"`
	// approval_address is the account address for approve create bucket/object signature
	ApprovalAddress string `protobuf:"bytes,4,opt,name=approval_address,json=approvalAddress,proto3" json:"approval_address,omitempty"`
	// endpoint is the domain name address used by SP to provide storage services
	Endpoint string `protobuf:"bytes,5,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// total_deposit is the token coin that the storage provider deposit to the storage module
	TotalDeposit string `protobuf:"bytes,6,opt,name=total_deposit,json=totalDeposit,proto3" json:"total_deposit,omitempty"`
}

func (m *EventCreateStorageProvider) Reset()         { *m = EventCreateStorageProvider{} }
func (m *EventCreateStorageProvider) String() string { return proto.CompactTextString(m) }
func (*EventCreateStorageProvider) ProtoMessage()    {}
func (*EventCreateStorageProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_685cbfa50fdf0841, []int{0}
}
func (m *EventCreateStorageProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateStorageProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateStorageProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateStorageProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateStorageProvider.Merge(m, src)
}
func (m *EventCreateStorageProvider) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateStorageProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateStorageProvider.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateStorageProvider proto.InternalMessageInfo

func (m *EventCreateStorageProvider) GetSpAddress() string {
	if m != nil {
		return m.SpAddress
	}
	return ""
}

func (m *EventCreateStorageProvider) GetFundingAddress() string {
	if m != nil {
		return m.FundingAddress
	}
	return ""
}

func (m *EventCreateStorageProvider) GetSealAddress() string {
	if m != nil {
		return m.SealAddress
	}
	return ""
}

func (m *EventCreateStorageProvider) GetApprovalAddress() string {
	if m != nil {
		return m.ApprovalAddress
	}
	return ""
}

func (m *EventCreateStorageProvider) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *EventCreateStorageProvider) GetTotalDeposit() string {
	if m != nil {
		return m.TotalDeposit
	}
	return ""
}

// EventEditStorageProvider is emitted when SP's metadata is edited.
type EventEditStorageProvider struct {
	// old_endpoint is the service endpoint of the storage provider before edit
	OldEndpoint string `protobuf:"bytes,1,opt,name=old_endpoint,json=oldEndpoint,proto3" json:"old_endpoint,omitempty"`
	// new_endpoint is the service endpoint of the storage provider after edit
	NewEndpoint string `protobuf:"bytes,2,opt,name=new_endpoint,json=newEndpoint,proto3" json:"new_endpoint,omitempty"`
}

func (m *EventEditStorageProvider) Reset()         { *m = EventEditStorageProvider{} }
func (m *EventEditStorageProvider) String() string { return proto.CompactTextString(m) }
func (*EventEditStorageProvider) ProtoMessage()    {}
func (*EventEditStorageProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_685cbfa50fdf0841, []int{1}
}
func (m *EventEditStorageProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEditStorageProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEditStorageProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEditStorageProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEditStorageProvider.Merge(m, src)
}
func (m *EventEditStorageProvider) XXX_Size() int {
	return m.Size()
}
func (m *EventEditStorageProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEditStorageProvider.DiscardUnknown(m)
}

var xxx_messageInfo_EventEditStorageProvider proto.InternalMessageInfo

func (m *EventEditStorageProvider) GetOldEndpoint() string {
	if m != nil {
		return m.OldEndpoint
	}
	return ""
}

func (m *EventEditStorageProvider) GetNewEndpoint() string {
	if m != nil {
		return m.NewEndpoint
	}
	return ""
}

// EventDeposit is emitted when sp deposit tokens.
type EventDeposit struct {
	// sp_address is the account address of storage provider
	SpAddress string `protobuf:"bytes,1,opt,name=sp_address,json=spAddress,proto3" json:"sp_address,omitempty"`
	// deposit is the token coin deposited this message
	Deposit string `protobuf:"bytes,2,opt,name=deposit,proto3" json:"deposit,omitempty"`
	// total_deposit is the total token coins this storage provider deposited
	TotalDeposit string `protobuf:"bytes,3,opt,name=total_deposit,json=totalDeposit,proto3" json:"total_deposit,omitempty"`
}

func (m *EventDeposit) Reset()         { *m = EventDeposit{} }
func (m *EventDeposit) String() string { return proto.CompactTextString(m) }
func (*EventDeposit) ProtoMessage()    {}
func (*EventDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_685cbfa50fdf0841, []int{2}
}
func (m *EventDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDeposit.Merge(m, src)
}
func (m *EventDeposit) XXX_Size() int {
	return m.Size()
}
func (m *EventDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_EventDeposit proto.InternalMessageInfo

func (m *EventDeposit) GetSpAddress() string {
	if m != nil {
		return m.SpAddress
	}
	return ""
}

func (m *EventDeposit) GetDeposit() string {
	if m != nil {
		return m.Deposit
	}
	return ""
}

func (m *EventDeposit) GetTotalDeposit() string {
	if m != nil {
		return m.TotalDeposit
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreateStorageProvider)(nil), "bnbchain.greenfield.sp.EventCreateStorageProvider")
	proto.RegisterType((*EventEditStorageProvider)(nil), "bnbchain.greenfield.sp.EventEditStorageProvider")
	proto.RegisterType((*EventDeposit)(nil), "bnbchain.greenfield.sp.EventDeposit")
}

func init() { proto.RegisterFile("greenfield/sp/events.proto", fileDescriptor_685cbfa50fdf0841) }

var fileDescriptor_685cbfa50fdf0841 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x8e, 0xda, 0x40,
	0x10, 0xc6, 0x31, 0x24, 0x24, 0x2c, 0x4e, 0x88, 0xac, 0x28, 0x72, 0x5c, 0x58, 0x81, 0x34, 0x51,
	0x24, 0xec, 0x22, 0x45, 0x8a, 0x54, 0xfc, 0xeb, 0x23, 0xe8, 0xd2, 0x38, 0x36, 0x3b, 0x98, 0x95,
	0xcc, 0xee, 0x6a, 0x77, 0x81, 0xe4, 0x05, 0xae, 0xbe, 0x87, 0xb9, 0x87, 0xb8, 0x12, 0x5d, 0x75,
	0xd2, 0x35, 0x27, 0x78, 0x91, 0x93, 0xd7, 0x5e, 0x83, 0x8e, 0x02, 0xe9, 0xca, 0x99, 0xfd, 0x7e,
	0xdf, 0xa7, 0x99, 0x1d, 0xe4, 0xa5, 0x02, 0x80, 0x2e, 0x08, 0x64, 0x38, 0x94, 0x3c, 0x84, 0x0d,
	0x50, 0x25, 0x03, 0x2e, 0x98, 0x62, 0xce, 0xa7, 0x84, 0x26, 0xf3, 0x65, 0x4c, 0x68, 0x70, 0x14,
	0x05, 0x92, 0x7b, 0x9f, 0xe7, 0x4c, 0xae, 0x98, 0x8c, 0xb4, 0x2a, 0x2c, 0x8a, 0x02, 0xe9, 0x3d,
	0xd4, 0x91, 0x37, 0xc9, 0x3d, 0x46, 0x02, 0x62, 0x05, 0x33, 0xc5, 0x44, 0x9c, 0xc2, 0x6f, 0xc1,
	0x36, 0x04, 0x83, 0x70, 0x7e, 0x22, 0x24, 0x79, 0x14, 0x63, 0x2c, 0x40, 0x4a, 0xd7, 0xfa, 0x62,
	0x7d, 0x6b, 0x0d, 0xdd, 0xbb, 0x9b, 0xfe, 0xc7, 0xd2, 0x64, 0x50, 0xbc, 0xcc, 0x94, 0x20, 0x34,
	0x9d, 0xb6, 0x24, 0x2f, 0x1b, 0xce, 0x00, 0x75, 0x16, 0x6b, 0x8a, 0x09, 0x4d, 0x2b, 0xba, 0x7e,
	0x81, 0x7e, 0x5f, 0x02, 0xc6, 0xe2, 0x17, 0xb2, 0x25, 0xc4, 0x59, 0xc5, 0x37, 0x2e, 0xf0, 0xed,
	0x5c, 0x6d, 0xe0, 0x11, 0xfa, 0x10, 0x73, 0x2e, 0xd8, 0xe6, 0xc4, 0xe0, 0xd5, 0x05, 0x83, 0x8e,
	0x21, 0x8c, 0x89, 0x87, 0xde, 0x02, 0xc5, 0x9c, 0x11, 0xaa, 0xdc, 0xd7, 0x39, 0x3c, 0xad, 0x6a,
	0xe7, 0x2b, 0x7a, 0xa7, 0x98, 0x8a, 0xb3, 0x08, 0x03, 0x67, 0x92, 0x28, 0xb7, 0xa9, 0x05, 0xb6,
	0x6e, 0x8e, 0x8b, 0x5e, 0xef, 0x2f, 0x72, 0xf5, 0x72, 0x27, 0x98, 0xa8, 0xe7, 0xab, 0xed, 0x22,
	0x9b, 0x65, 0x38, 0xaa, 0x02, 0xf4, 0x72, 0xa7, 0x6d, 0x96, 0xe1, 0x89, 0xc9, 0xe8, 0x22, 0x9b,
	0xc2, 0xf6, 0x28, 0xa9, 0x17, 0x12, 0x0a, 0x5b, 0x23, 0xe9, 0x5d, 0x59, 0xc8, 0xd6, 0x11, 0x65,
	0xe4, 0xcb, 0x7f, 0xcc, 0x45, 0x6f, 0xcc, 0x28, 0x45, 0x8e, 0x29, 0xcf, 0x47, 0x6d, 0x9c, 0x8f,
	0x3a, 0x1c, 0xdf, 0xee, 0x7d, 0x6b, 0xb7, 0xf7, 0xad, 0xc7, 0xbd, 0x6f, 0x5d, 0x1f, 0xfc, 0xda,
	0xee, 0xe0, 0xd7, 0xee, 0x0f, 0x7e, 0xed, 0xcf, 0xf7, 0x94, 0xa8, 0xe5, 0x3a, 0x09, 0xe6, 0x6c,
	0x15, 0x26, 0x34, 0xe9, 0xeb, 0x0b, 0x0d, 0x4f, 0xce, 0xf8, 0x5f, 0x7e, 0xc8, 0xea, 0x3f, 0x07,
	0x99, 0x34, 0xf5, 0x55, 0xfe, 0x78, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x8f, 0xc2, 0xf1, 0xe6,
	0x02, 0x00, 0x00,
}

func (m *EventCreateStorageProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateStorageProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateStorageProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalDeposit) > 0 {
		i -= len(m.TotalDeposit)
		copy(dAtA[i:], m.TotalDeposit)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TotalDeposit)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ApprovalAddress) > 0 {
		i -= len(m.ApprovalAddress)
		copy(dAtA[i:], m.ApprovalAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ApprovalAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SealAddress) > 0 {
		i -= len(m.SealAddress)
		copy(dAtA[i:], m.SealAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SealAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FundingAddress) > 0 {
		i -= len(m.FundingAddress)
		copy(dAtA[i:], m.FundingAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.FundingAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpAddress) > 0 {
		i -= len(m.SpAddress)
		copy(dAtA[i:], m.SpAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SpAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventEditStorageProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEditStorageProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEditStorageProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewEndpoint) > 0 {
		i -= len(m.NewEndpoint)
		copy(dAtA[i:], m.NewEndpoint)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewEndpoint)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OldEndpoint) > 0 {
		i -= len(m.OldEndpoint)
		copy(dAtA[i:], m.OldEndpoint)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.OldEndpoint)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalDeposit) > 0 {
		i -= len(m.TotalDeposit)
		copy(dAtA[i:], m.TotalDeposit)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TotalDeposit)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpAddress) > 0 {
		i -= len(m.SpAddress)
		copy(dAtA[i:], m.SpAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SpAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateStorageProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.FundingAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SealAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ApprovalAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TotalDeposit)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventEditStorageProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OldEndpoint)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewEndpoint)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TotalDeposit)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateStorageProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCreateStorageProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateStorageProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FundingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SealAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SealAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApprovalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDeposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventEditStorageProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventEditStorageProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEditStorageProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDeposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
