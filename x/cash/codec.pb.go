// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/cash/codec.proto

package cash

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import x "github.com/iov-one/weave/x"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Set may contain Coin of many different currencies.
// It handles adding and subtracting sets of currencies.
type Set struct {
	Coins                []*x.Coin `protobuf:"bytes,1,rep,name=coins" json:"coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Set) Reset()         { *m = Set{} }
func (m *Set) String() string { return proto.CompactTextString(m) }
func (*Set) ProtoMessage()    {}
func (*Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_bf507ad2773f489b, []int{0}
}
func (m *Set) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Set.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Set.Merge(dst, src)
}
func (m *Set) XXX_Size() int {
	return m.Size()
}
func (m *Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Set proto.InternalMessageInfo

func (m *Set) GetCoins() []*x.Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

// SendMsg is a request to move these coins from the given
// source to the given destination address.
// memo is an optional human-readable message
// ref is optional binary data, that can refer to another
// eg. tx hash
type SendMsg struct {
	Src    []byte  `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Dest   []byte  `protobuf:"bytes,2,opt,name=dest,proto3" json:"dest,omitempty"`
	Amount *x.Coin `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	// max length 128 character
	Memo string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
	// max length 64 bytes
	Ref                  []byte   `protobuf:"bytes,5,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsg) Reset()         { *m = SendMsg{} }
func (m *SendMsg) String() string { return proto.CompactTextString(m) }
func (*SendMsg) ProtoMessage()    {}
func (*SendMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_bf507ad2773f489b, []int{1}
}
func (m *SendMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SendMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsg.Merge(dst, src)
}
func (m *SendMsg) XXX_Size() int {
	return m.Size()
}
func (m *SendMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsg proto.InternalMessageInfo

func (m *SendMsg) GetSrc() []byte {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *SendMsg) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *SendMsg) GetAmount() *x.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *SendMsg) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *SendMsg) GetRef() []byte {
	if m != nil {
		return m.Ref
	}
	return nil
}

// FeeInfo records who pays what fees to have this
// message processed
type FeeInfo struct {
	Payer                []byte   `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Fees                 *x.Coin  `protobuf:"bytes,2,opt,name=fees" json:"fees,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeeInfo) Reset()         { *m = FeeInfo{} }
func (m *FeeInfo) String() string { return proto.CompactTextString(m) }
func (*FeeInfo) ProtoMessage()    {}
func (*FeeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_bf507ad2773f489b, []int{2}
}
func (m *FeeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FeeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeInfo.Merge(dst, src)
}
func (m *FeeInfo) XXX_Size() int {
	return m.Size()
}
func (m *FeeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FeeInfo proto.InternalMessageInfo

func (m *FeeInfo) GetPayer() []byte {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (m *FeeInfo) GetFees() *x.Coin {
	if m != nil {
		return m.Fees
	}
	return nil
}

func init() {
	proto.RegisterType((*Set)(nil), "cash.Set")
	proto.RegisterType((*SendMsg)(nil), "cash.SendMsg")
	proto.RegisterType((*FeeInfo)(nil), "cash.FeeInfo")
}
func (m *Set) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Set) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, msg := range m.Coins {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SendMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Src) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Src)))
		i += copy(dAtA[i:], m.Src)
	}
	if len(m.Dest) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Dest)))
		i += copy(dAtA[i:], m.Dest)
	}
	if m.Amount != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Amount.Size()))
		n1, err := m.Amount.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	if len(m.Ref) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Ref)))
		i += copy(dAtA[i:], m.Ref)
	}
	return i, nil
}

func (m *FeeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Payer) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Payer)))
		i += copy(dAtA[i:], m.Payer)
	}
	if m.Fees != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Fees.Size()))
		n2, err := m.Fees.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Set) Size() (n int) {
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *SendMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Src)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Dest)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *FeeInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Payer)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Fees != nil {
		l = m.Fees.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Set) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Set: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Set: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, &x.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *SendMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Src", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Src = append(m.Src[:0], dAtA[iNdEx:postIndex]...)
			if m.Src == nil {
				m.Src = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dest = append(m.Dest[:0], dAtA[iNdEx:postIndex]...)
			if m.Dest == nil {
				m.Dest = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Amount == nil {
				m.Amount = &x.Coin{}
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = append(m.Ref[:0], dAtA[iNdEx:postIndex]...)
			if m.Ref == nil {
				m.Ref = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *FeeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payer = append(m.Payer[:0], dAtA[iNdEx:postIndex]...)
			if m.Payer == nil {
				m.Payer = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fees == nil {
				m.Fees = &x.Coin{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/cash/codec.proto", fileDescriptor_codec_bf507ad2773f489b) }

var fileDescriptor_codec_bf507ad2773f489b = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0x5d, 0x81, 0x23, 0x8e, 0x16, 0x97, 0x8d, 0xc5, 0x46, 0x23, 0x12, 0x62, 0x0c, 0x8d,
	0x90, 0x68, 0x6b, 0xa5, 0x89, 0x89, 0x85, 0x0d, 0xf7, 0x04, 0xdc, 0x32, 0xdc, 0x51, 0xb0, 0x73,
	0xd9, 0xdd, 0x43, 0x7c, 0x0b, 0x1f, 0xcb, 0xd2, 0x47, 0x30, 0xf8, 0x22, 0x17, 0x96, 0x2b, 0xe8,
	0xfe, 0xf9, 0x32, 0xf3, 0x7f, 0x19, 0xe0, 0x7d, 0x2e, 0x4b, 0xb3, 0xcd, 0x25, 0x55, 0x28, 0xb3,
	0x9d, 0x26, 0x4b, 0xdc, 0x1f, 0xc9, 0xd5, 0xfd, 0xa6, 0xb1, 0xdb, 0xfd, 0x3a, 0x93, 0xd4, 0xe6,
	0x0d, 0x75, 0x0f, 0xa4, 0x30, 0xff, 0xc4, 0xb2, 0xc3, 0xbc, 0x9f, 0x6f, 0x27, 0x77, 0xe0, 0xad,
	0xd0, 0xf2, 0x1b, 0x08, 0x24, 0x35, 0xca, 0x08, 0x16, 0x7b, 0xe9, 0xf9, 0x63, 0x98, 0xf5, 0xd9,
	0x2b, 0x35, 0xaa, 0x98, 0x68, 0xd2, 0x41, 0xb8, 0x42, 0x55, 0x7d, 0x98, 0x0d, 0x5f, 0x82, 0x67,
	0xb4, 0x14, 0x2c, 0x66, 0xe9, 0x45, 0x31, 0x46, 0xce, 0xc1, 0xaf, 0xd0, 0x58, 0x71, 0xea, 0x90,
	0xcb, 0xfc, 0x16, 0x16, 0x65, 0x4b, 0x7b, 0x65, 0x85, 0x17, 0xb3, 0x79, 0xe1, 0x11, 0x8f, 0x47,
	0x2d, 0xb6, 0x24, 0xfc, 0x98, 0xa5, 0x67, 0x85, 0xcb, 0x63, 0xb5, 0xc6, 0x5a, 0x04, 0x53, 0xb5,
	0xc6, 0x3a, 0x79, 0x86, 0xf0, 0x0d, 0xf1, 0x5d, 0xd5, 0xc4, 0x2f, 0x21, 0xd8, 0x95, 0x5f, 0xa8,
	0x8f, 0xe6, 0x69, 0xe0, 0xd7, 0xe0, 0xd7, 0x88, 0xc6, 0xb9, 0x67, 0x16, 0x07, 0x5f, 0x96, 0x3f,
	0x43, 0xc4, 0x7e, 0x87, 0x88, 0xfd, 0x0d, 0x11, 0xfb, 0xfe, 0x8f, 0x4e, 0xd6, 0x0b, 0xf7, 0xf4,
	0xd3, 0x21, 0x00, 0x00, 0xff, 0xff, 0xce, 0x64, 0x21, 0xca, 0x38, 0x01, 0x00, 0x00,
}
