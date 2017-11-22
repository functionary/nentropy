// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: state.proto

package multiraftbase

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RaftTruncatedState struct {
	// The highest index that has been removed from the log.
	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The term corresponding to 'index'.
	Term uint64 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (m *RaftTruncatedState) Reset()                    { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string            { return proto.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()               {}
func (*RaftTruncatedState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

func (m *RaftTruncatedState) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RaftTruncatedState) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type ReplicaState struct {
	RaftAppliedIndex uint64           `protobuf:"varint,1,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	Desc             *GroupDescriptor `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
	// The truncation state of the Raft log.
	TruncatedState *RaftTruncatedState `protobuf:"bytes,3,opt,name=truncated_state,json=truncatedState" json:"truncated_state,omitempty"`
}

func (m *ReplicaState) Reset()                    { *m = ReplicaState{} }
func (m *ReplicaState) String() string            { return proto.CompactTextString(m) }
func (*ReplicaState) ProtoMessage()               {}
func (*ReplicaState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{1} }

func (m *ReplicaState) GetRaftAppliedIndex() uint64 {
	if m != nil {
		return m.RaftAppliedIndex
	}
	return 0
}

func (m *ReplicaState) GetDesc() *GroupDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *ReplicaState) GetTruncatedState() *RaftTruncatedState {
	if m != nil {
		return m.TruncatedState
	}
	return nil
}

type PgInfo struct {
	ReplicaState `protobuf:"bytes,1,opt,name=state,embedded=state" json:"state"`
	// The highest (and last) index in the Raft log.
	LastIndex  uint64 `protobuf:"varint,2,opt,name=lastIndex,proto3" json:"lastIndex,omitempty"`
	NumPending uint64 `protobuf:"varint,3,opt,name=num_pending,json=numPending,proto3" json:"num_pending,omitempty"`
	NumDropped uint64 `protobuf:"varint,5,opt,name=num_dropped,json=numDropped,proto3" json:"num_dropped,omitempty"`
	// raft_log_size may be initially inaccurate after a server restart.
	// See storage.Replica.mu.raftLogSize.
	RaftLogSize int64 `protobuf:"varint,6,opt,name=raft_log_size,json=raftLogSize,proto3" json:"raft_log_size,omitempty"`
	// Approximately the amount of quota available.
	ApproximateProposalQuota int64 `protobuf:"varint,7,opt,name=approximate_proposal_quota,json=approximateProposalQuota,proto3" json:"approximate_proposal_quota,omitempty"`
}

func (m *PgInfo) Reset()                    { *m = PgInfo{} }
func (m *PgInfo) String() string            { return proto.CompactTextString(m) }
func (*PgInfo) ProtoMessage()               {}
func (*PgInfo) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{2} }

func (m *PgInfo) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *PgInfo) GetNumPending() uint64 {
	if m != nil {
		return m.NumPending
	}
	return 0
}

func (m *PgInfo) GetNumDropped() uint64 {
	if m != nil {
		return m.NumDropped
	}
	return 0
}

func (m *PgInfo) GetRaftLogSize() int64 {
	if m != nil {
		return m.RaftLogSize
	}
	return 0
}

func (m *PgInfo) GetApproximateProposalQuota() int64 {
	if m != nil {
		return m.ApproximateProposalQuota
	}
	return 0
}

func init() {
	proto.RegisterType((*RaftTruncatedState)(nil), "multiraftbase.RaftTruncatedState")
	proto.RegisterType((*ReplicaState)(nil), "multiraftbase.ReplicaState")
	proto.RegisterType((*PgInfo)(nil), "multiraftbase.PgInfo")
}
func (this *RaftTruncatedState) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RaftTruncatedState)
	if !ok {
		that2, ok := that.(RaftTruncatedState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Term != that1.Term {
		return false
	}
	return true
}
func (this *PgInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PgInfo)
	if !ok {
		that2, ok := that.(PgInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ReplicaState.Equal(&that1.ReplicaState) {
		return false
	}
	if this.LastIndex != that1.LastIndex {
		return false
	}
	if this.NumPending != that1.NumPending {
		return false
	}
	if this.NumDropped != that1.NumDropped {
		return false
	}
	if this.RaftLogSize != that1.RaftLogSize {
		return false
	}
	if this.ApproximateProposalQuota != that1.ApproximateProposalQuota {
		return false
	}
	return true
}
func (m *RaftTruncatedState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftTruncatedState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Index))
	}
	if m.Term != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Term))
	}
	return i, nil
}

func (m *ReplicaState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintState(dAtA, i, uint64(m.RaftAppliedIndex))
	}
	if m.Desc != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Desc.Size()))
		n1, err := m.Desc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.TruncatedState != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintState(dAtA, i, uint64(m.TruncatedState.Size()))
		n2, err := m.TruncatedState.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *PgInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PgInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintState(dAtA, i, uint64(m.ReplicaState.Size()))
	n3, err := m.ReplicaState.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.LastIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintState(dAtA, i, uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintState(dAtA, i, uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintState(dAtA, i, uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintState(dAtA, i, uint64(m.RaftLogSize))
	}
	if m.ApproximateProposalQuota != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintState(dAtA, i, uint64(m.ApproximateProposalQuota))
	}
	return i, nil
}

func encodeFixed64State(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32State(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedRaftTruncatedState(r randyState, easy bool) *RaftTruncatedState {
	this := &RaftTruncatedState{}
	this.Index = uint64(uint64(r.Uint32()))
	this.Term = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyState interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneState(r randyState) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringState(r randyState) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneState(r)
	}
	return string(tmps)
}
func randUnrecognizedState(r randyState, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldState(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldState(dAtA []byte, r randyState, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateState(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateState(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateState(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateState(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateState(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateState(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateState(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *RaftTruncatedState) Size() (n int) {
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovState(uint64(m.Index))
	}
	if m.Term != 0 {
		n += 1 + sovState(uint64(m.Term))
	}
	return n
}

func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.RaftAppliedIndex))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.TruncatedState != nil {
		l = m.TruncatedState.Size()
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *PgInfo) Size() (n int) {
	var l int
	_ = l
	l = m.ReplicaState.Size()
	n += 1 + l + sovState(uint64(l))
	if m.LastIndex != 0 {
		n += 1 + sovState(uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		n += 1 + sovState(uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		n += 1 + sovState(uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		n += 1 + sovState(uint64(m.RaftLogSize))
	}
	if m.ApproximateProposalQuota != 0 {
		n += 1 + sovState(uint64(m.ApproximateProposalQuota))
	}
	return n
}

func sovState(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftTruncatedState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: RaftTruncatedState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftTruncatedState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func (m *ReplicaState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ReplicaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &GroupDescriptor{}
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TruncatedState == nil {
				m.TruncatedState = &RaftTruncatedState{}
			}
			if err := m.TruncatedState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func (m *PgInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: PgInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PgInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReplicaState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPending", wireType)
			}
			m.NumPending = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumPending |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumDropped", wireType)
			}
			m.NumDropped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumDropped |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftLogSize", wireType)
			}
			m.RaftLogSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftLogSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApproximateProposalQuota", wireType)
			}
			m.ApproximateProposalQuota = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApproximateProposalQuota |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowState
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
				next, err := skipState(dAtA[start:])
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
	ErrInvalidLengthState = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("state.proto", fileDescriptorState) }

var fileDescriptorState = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd7, 0xbb, 0x49, 0x29, 0x0e, 0xbb, 0x54, 0xd6, 0x1e, 0xa2, 0x82, 0x92, 0x25, 0xa7,
	0x3d, 0xa0, 0x22, 0x2d, 0xb7, 0x85, 0x0b, 0xab, 0x95, 0x60, 0x2b, 0x0e, 0xc5, 0xe5, 0x1e, 0xb9,
	0xb1, 0x1b, 0x59, 0x4a, 0x62, 0xe3, 0x4c, 0xa4, 0xaa, 0xaf, 0xc1, 0x85, 0x63, 0x1f, 0x83, 0x03,
	0x0f, 0xd0, 0x63, 0x9f, 0xa0, 0x42, 0xe5, 0xc2, 0x63, 0xa0, 0xd8, 0x2d, 0xb4, 0xe5, 0x36, 0xfe,
	0xe7, 0x9f, 0x4f, 0xf3, 0x4f, 0x82, 0x83, 0x1a, 0x18, 0x88, 0x81, 0x36, 0x0a, 0x14, 0x39, 0x2f,
	0x9b, 0x02, 0xa4, 0x61, 0x53, 0x98, 0xb0, 0x5a, 0xf4, 0x2f, 0x4a, 0x01, 0x8c, 0x33, 0x60, 0xae,
	0xdd, 0xbf, 0xcc, 0x55, 0xae, 0x6c, 0xf9, 0xaa, 0xad, 0x9c, 0x9a, 0x7c, 0xc0, 0x84, 0xb2, 0x29,
	0x7c, 0x36, 0x4d, 0x95, 0x31, 0x10, 0x7c, 0xdc, 0x02, 0xc9, 0x25, 0xf6, 0x65, 0xc5, 0xc5, 0x2c,
	0x44, 0x57, 0xe8, 0xda, 0xa3, 0xee, 0x41, 0x08, 0xf6, 0x40, 0x98, 0x32, 0x3c, 0xb5, 0xa2, 0xad,
	0x6f, 0xbb, 0xdf, 0x17, 0x31, 0xfa, 0xbd, 0x88, 0x51, 0xf2, 0x03, 0xe1, 0x27, 0x54, 0xe8, 0x42,
	0x66, 0xcc, 0x41, 0x5e, 0x62, 0xd2, 0x2e, 0x93, 0x32, 0xad, 0x0b, 0x29, 0x78, 0xba, 0x4f, 0xec,
	0xb5, 0x9d, 0x77, 0xae, 0xf1, 0x60, 0xe1, 0x37, 0xd8, 0xe3, 0xa2, 0xce, 0x2c, 0x3c, 0xb8, 0x89,
	0x06, 0x07, 0x61, 0x06, 0xef, 0x8d, 0x6a, 0xf4, 0xbd, 0xa8, 0x33, 0x23, 0x35, 0x28, 0x43, 0xad,
	0x97, 0x0c, 0xf1, 0x53, 0xd8, 0x2d, 0x9e, 0xda, 0x53, 0x84, 0x67, 0x76, 0xfc, 0xc5, 0xd1, 0xf8,
	0xff, 0x11, 0xe9, 0x05, 0x1c, 0xbc, 0x93, 0xaf, 0xa7, 0xb8, 0x33, 0xca, 0x1f, 0xaa, 0xa9, 0x22,
	0x6f, 0xb0, 0xef, 0x60, 0xc8, 0xc2, 0x9e, 0x1d, 0xc3, 0xf6, 0x42, 0xde, 0x75, 0x97, 0xeb, 0xf8,
	0x64, 0xb5, 0x8e, 0x11, 0x75, 0x33, 0xe4, 0x39, 0x7e, 0x5c, 0xb0, 0x1a, 0x6c, 0xa8, 0xed, 0xa5,
	0xfe, 0x09, 0x24, 0xc6, 0x41, 0xd5, 0x94, 0xa9, 0x16, 0x15, 0x97, 0x55, 0x6e, 0xb7, 0xf5, 0x28,
	0xae, 0x9a, 0x72, 0xe4, 0x94, 0x9d, 0x81, 0x1b, 0xa5, 0xb5, 0xe0, 0xa1, 0xff, 0xd7, 0x70, 0xef,
	0x14, 0x92, 0xe0, 0x73, 0x7b, 0xd5, 0x42, 0xe5, 0x69, 0x2d, 0xe7, 0x22, 0xec, 0x5c, 0xa1, 0xeb,
	0x33, 0x1a, 0xb4, 0xe2, 0x47, 0x95, 0x8f, 0xe5, 0x5c, 0x90, 0xb7, 0xb8, 0xcf, 0xb4, 0x36, 0x6a,
	0x26, 0x4b, 0x06, 0x22, 0xd5, 0x46, 0x69, 0x55, 0xb3, 0x22, 0xfd, 0xd2, 0x28, 0x60, 0xe1, 0x23,
	0x3b, 0x10, 0xee, 0x39, 0x46, 0x5b, 0xc3, 0xa7, 0xb6, 0x7f, 0xeb, 0xb5, 0x9f, 0x73, 0xe8, 0x75,
	0xbd, 0x9e, 0x7f, 0xd7, 0x5b, 0x6e, 0x22, 0xb4, 0xda, 0x44, 0xe8, 0xe7, 0x26, 0x42, 0xdf, 0x7e,
	0x45, 0x27, 0x93, 0x8e, 0xfd, 0x6f, 0x5e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xe4, 0xb4,
	0x09, 0x7b, 0x02, 0x00, 0x00,
}
