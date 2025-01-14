package main

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TimerOld struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Alert                string   `protobuf:"bytes,2,opt,name=alert,proto3" json:"alert,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Month                int32    `protobuf:"zigzag32,4,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int32    `protobuf:"zigzag32,5,opt,name=day,proto3" json:"day,omitempty"`
	Week                 int32    `protobuf:"zigzag32,6,opt,name=week,proto3" json:"week,omitempty"`
	Hour                 int32    `protobuf:"zigzag32,7,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute               int32    `protobuf:"zigzag32,8,opt,name=minute,proto3" json:"minute,omitempty"`
	Grpid                uint64   `protobuf:"varint,9,opt,name=grpid,proto3" json:"grpid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimerOld) Reset()         { *m = TimerOld{} }
func (m *TimerOld) String() string { return proto.CompactTextString(m) }
func (*TimerOld) ProtoMessage()    {}
func (*TimerOld) Descriptor() ([]byte, []int) {
	return fileDescriptor_old, []int{0}
}
func (m *TimerOld) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimerOld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Timer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimerOld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timer.Merge(m, src)
}
func (m *TimerOld) XXX_Size() int {
	return m.Size()
}
func (m *TimerOld) XXX_DiscardUnknown() {
	xxx_messageInfo_Timer.DiscardUnknown(m)
}

func (m *TimerOld) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *TimerOld) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *TimerOld) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *TimerOld) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *TimerOld) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *TimerOld) GetWeek() int32 {
	if m != nil {
		return m.Week
	}
	return 0
}

func (m *TimerOld) GetHour() int32 {
	if m != nil {
		return m.Hour
	}
	return 0
}

func (m *TimerOld) GetMinute() int32 {
	if m != nil {
		return m.Minute
	}
	return 0
}

func (m *TimerOld) GetGrpid() uint64 {
	if m != nil {
		return m.Grpid
	}
	return 0
}

type TimersMapOld struct {
	Timers               map[string]*TimerOld `protobuf:"bytes,1,rep,name=timers,proto3" json:"timers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimersMapOld) Reset()         { *m = TimersMapOld{} }
func (m *TimersMapOld) String() string { return proto.CompactTextString(m) }
func (*TimersMapOld) ProtoMessage()    {}
func (*TimersMapOld) Descriptor() ([]byte, []int) {
	return fileDescriptor_old, []int{1}
}
func (m *TimersMapOld) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimersMapOld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimersMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimersMapOld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimersMap.Merge(m, src)
}
func (m *TimersMapOld) XXX_Size() int {
	return m.Size()
}
func (m *TimersMapOld) XXX_DiscardUnknown() {
	xxx_messageInfo_TimersMap.DiscardUnknown(m)
}

func (m *TimersMapOld) GetTimers() map[string]*TimerOld {
	if m != nil {
		return m.Timers
	}
	return nil
}

func init() {
	proto.RegisterType((*TimerOld)(nil), "timer.Timer")
	proto.RegisterType((*TimersMapOld)(nil), "timer.TimersMap")
	proto.RegisterMapType((map[string]*TimerOld)(nil), "timer.TimersMap.TimersEntry")
}

func init() { proto.RegisterFile("timer.proto", fileDescriptor_old) }

var fileDescriptor_old = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x9d, 0xe6, 0xc3, 0x66, 0xe3, 0xa1, 0x2e, 0x22, 0x83, 0x48, 0x08, 0x39, 0xe5, 0xd4,
	0x43, 0xf5, 0x20, 0x1e, 0x05, 0xf1, 0xe4, 0x65, 0xf1, 0x05, 0x52, 0xba, 0xd8, 0xd0, 0x7c, 0xb1,
	0xdd, 0x28, 0x79, 0x05, 0x9f, 0xc0, 0x17, 0x12, 0x3c, 0xfa, 0x08, 0x12, 0x5f, 0x44, 0x66, 0x36,
	0x94, 0xde, 0x7e, 0xff, 0x5f, 0xfe, 0x64, 0x67, 0x46, 0xc4, 0xb6, 0xac, 0xb5, 0x59, 0x76, 0xa6,
	0xb5, 0xad, 0x0c, 0x38, 0x64, 0x5f, 0x20, 0x82, 0x17, 0x22, 0x79, 0x29, 0x42, 0xdd, 0x14, 0xeb,
	0x4a, 0x23, 0xa4, 0x90, 0xcf, 0xd5, 0x94, 0xe4, 0x85, 0x08, 0x8a, 0x4a, 0x1b, 0x8b, 0xb3, 0x14,
	0xf2, 0x48, 0xb9, 0x20, 0x17, 0xc2, 0xeb, 0x4d, 0x85, 0x1e, 0x3b, 0x42, 0xea, 0xd5, 0x6d, 0x63,
	0xb7, 0xe8, 0xa7, 0x90, 0x9f, 0x2b, 0x17, 0xa8, 0xb7, 0x29, 0x06, 0x0c, 0xd8, 0x11, 0x4a, 0x29,
	0xfc, 0x77, 0xad, 0x77, 0x18, 0xb2, 0x62, 0x26, 0xb7, 0x6d, 0x7b, 0x83, 0xa7, 0xce, 0x11, 0xd3,
	0x3c, 0x75, 0xd9, 0xf4, 0x56, 0xe3, 0x9c, 0xed, 0x94, 0xe8, 0x9d, 0x57, 0xd3, 0x95, 0x1b, 0x8c,
	0x52, 0xc8, 0x7d, 0xe5, 0x42, 0xf6, 0x01, 0x22, 0xe2, 0x3d, 0xf6, 0xcf, 0x45, 0x27, 0x6f, 0x45,
	0xc8, 0xeb, 0xed, 0x11, 0x52, 0x2f, 0x8f, 0x57, 0xd7, 0x4b, 0xb7, 0xfa, 0xa1, 0x31, 0xd1, 0x63,
	0x63, 0xcd, 0xa0, 0xa6, 0xee, 0xd5, 0x93, 0x88, 0x8f, 0x34, 0x8d, 0xbe, 0xd3, 0x03, 0x5f, 0x23,
	0x52, 0x84, 0x32, 0x13, 0xc1, 0x5b, 0x51, 0xf5, 0x9a, 0x4f, 0x11, 0xaf, 0xce, 0x8e, 0xff, 0xaa,
	0xdc, 0xa7, 0xfb, 0xd9, 0x1d, 0x3c, 0x2c, 0xbe, 0xc7, 0x04, 0x7e, 0xc6, 0x04, 0x7e, 0xc7, 0x04,
	0x3e, 0xff, 0x92, 0x93, 0x75, 0xc8, 0x47, 0xbf, 0xf9, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xc7,
	0xad, 0xf8, 0x83, 0x01, 0x00, 0x00,
}

func (m *TimerOld) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimerOld) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimerOld) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Grpid != 0 {
		i = encodeVarintTimerOld(dAtA, i, uint64(m.Grpid))
		i--
		dAtA[i] = 0x48
	}
	if m.Minute != 0 {
		i = encodeVarintTimerOld(dAtA, i, uint64((uint32(m.Minute)<<1)^uint32((m.Minute>>31))))
		i--
		dAtA[i] = 0x40
	}
	if m.Hour != 0 {
		i = encodeVarintTimerOld(dAtA, i, uint64((uint32(m.Hour)<<1)^uint32((m.Hour>>31))))
		i--
		dAtA[i] = 0x38
	}
	if m.Week != 0 {
		i = encodeVarintTimerOld(dAtA, i, uint64((uint32(m.Week)<<1)^uint32((m.Week>>31))))
		i--
		dAtA[i] = 0x30
	}
	if m.Day != 0 {
		i = encodeVarintTimerOld(dAtA, i, uint64((uint32(m.Day)<<1)^uint32((m.Day>>31))))
		i--
		dAtA[i] = 0x28
	}
	if m.Month != 0 {
		i = encodeVarintTimerOld(dAtA, i, uint64((uint32(m.Month)<<1)^uint32((m.Month>>31))))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintTimerOld(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Alert) > 0 {
		i -= len(m.Alert)
		copy(dAtA[i:], m.Alert)
		i = encodeVarintTimerOld(dAtA, i, uint64(len(m.Alert)))
		i--
		dAtA[i] = 0x12
	}
	if m.Enable {
		i--
		if m.Enable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TimersMapOld) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimersMapOld) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimersMapOld) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Timers) > 0 {
		for k := range m.Timers {
			v := m.Timers[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTimerOld(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTimerOld(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTimerOld(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimerOld(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimerOld(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimerOld) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enable {
		n += 2
	}
	l = len(m.Alert)
	if l > 0 {
		n += 1 + l + sovTimerOld(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovTimerOld(uint64(l))
	}
	if m.Month != 0 {
		n += 1 + sozTimerOld(uint64(m.Month))
	}
	if m.Day != 0 {
		n += 1 + sozTimerOld(uint64(m.Day))
	}
	if m.Week != 0 {
		n += 1 + sozTimerOld(uint64(m.Week))
	}
	if m.Hour != 0 {
		n += 1 + sozTimerOld(uint64(m.Hour))
	}
	if m.Minute != 0 {
		n += 1 + sozTimerOld(uint64(m.Minute))
	}
	if m.Grpid != 0 {
		n += 1 + sovTimerOld(uint64(m.Grpid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TimersMapOld) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Timers) > 0 {
		for k, v := range m.Timers {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTimerOld(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovTimerOld(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTimerOld(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTimerOld(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimerOld(x uint64) (n int) {
	return sovTimerOld(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimerOld) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimerOld
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
			return fmt.Errorf("proto: Timer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enable = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
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
				return ErrInvalidLengthTimerOld
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimerOld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
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
				return ErrInvalidLengthTimerOld
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimerOld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Month", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Month = v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Day = v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Week", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Week = v
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hour", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Hour = v
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minute", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Minute = v
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grpid", wireType)
			}
			m.Grpid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Grpid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimerOld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimerOld
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimersMapOld) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimerOld
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
			return fmt.Errorf("proto: TimersMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimersMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimerOld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimerOld
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTimerOld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timers == nil {
				m.Timers = make(map[string]*TimerOld)
			}
			var mapkey string
			var mapvalue *TimerOld
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTimerOld
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTimerOld
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTimerOld
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTimerOld
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTimerOld
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTimerOld
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTimerOld
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &TimerOld{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTimerOld(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTimerOld
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Timers[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimerOld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimerOld
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimerOld(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimerOld
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
					return 0, ErrIntOverflowTimerOld
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
					return 0, ErrIntOverflowTimerOld
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
				return 0, ErrInvalidLengthTimerOld
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimerOld
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimerOld
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimerOld        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimerOld          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimerOld = fmt.Errorf("proto: unexpected end of group")
)
