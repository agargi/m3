// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/query/generated/proto/admin/namespace.proto

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package admin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import namespace1 "github.com/m3db/m3/src/dbnode/generated/proto/namespace"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NamespaceGetResponse struct {
	Registry *namespace1.Registry `protobuf:"bytes,1,opt,name=registry" json:"registry,omitempty"`
}

func (m *NamespaceGetResponse) Reset()                    { *m = NamespaceGetResponse{} }
func (m *NamespaceGetResponse) String() string            { return proto.CompactTextString(m) }
func (*NamespaceGetResponse) ProtoMessage()               {}
func (*NamespaceGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *NamespaceGetResponse) GetRegistry() *namespace1.Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type NamespaceAddRequest struct {
	Name    string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options *namespace1.NamespaceOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *NamespaceAddRequest) Reset()                    { *m = NamespaceAddRequest{} }
func (m *NamespaceAddRequest) String() string            { return proto.CompactTextString(m) }
func (*NamespaceAddRequest) ProtoMessage()               {}
func (*NamespaceAddRequest) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *NamespaceAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamespaceAddRequest) GetOptions() *namespace1.NamespaceOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type NamespaceSchemaAddRequest struct {
	// Name is the namespace name.
	// Add schema to non-existent namespace will get 404.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// MsgName is the name of the top level proto message.
	MsgName string `protobuf:"bytes,2,opt,name=msgName,proto3" json:"msgName,omitempty"`
	// ProtoName is the name of the top level proto file.
	// Proto file content will be looked up from protoMap, so the name must corresponds to the map key.
	ProtoName string `protobuf:"bytes,3,opt,name=protoName,proto3" json:"protoName,omitempty"`
	// ProtoMap is a map of name to proto strings.
	// Except the top level proto file, other imported proto files' key must be exactly the same
	// as how they are imported in the import statement.
	// E.g. If import.proto is imported using as below
	// import "mainpkg/imported.proto";
	// Then the map key for imported.proto must be "mainpkg/imported.proto"
	// See src/dbnode/namespame/kvadmin test for example.
	ProtoMap map[string]string `protobuf:"bytes,4,rep,name=protoMap" json:"protoMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *NamespaceSchemaAddRequest) Reset()         { *m = NamespaceSchemaAddRequest{} }
func (m *NamespaceSchemaAddRequest) String() string { return proto.CompactTextString(m) }
func (*NamespaceSchemaAddRequest) ProtoMessage()    {}
func (*NamespaceSchemaAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorNamespace, []int{2}
}

func (m *NamespaceSchemaAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamespaceSchemaAddRequest) GetMsgName() string {
	if m != nil {
		return m.MsgName
	}
	return ""
}

func (m *NamespaceSchemaAddRequest) GetProtoName() string {
	if m != nil {
		return m.ProtoName
	}
	return ""
}

func (m *NamespaceSchemaAddRequest) GetProtoMap() map[string]string {
	if m != nil {
		return m.ProtoMap
	}
	return nil
}

type NamespaceSchemaAddResponse struct {
	DeployID string `protobuf:"bytes,1,opt,name=deployID,proto3" json:"deployID,omitempty"`
}

func (m *NamespaceSchemaAddResponse) Reset()         { *m = NamespaceSchemaAddResponse{} }
func (m *NamespaceSchemaAddResponse) String() string { return proto.CompactTextString(m) }
func (*NamespaceSchemaAddResponse) ProtoMessage()    {}
func (*NamespaceSchemaAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorNamespace, []int{3}
}

func (m *NamespaceSchemaAddResponse) GetDeployID() string {
	if m != nil {
		return m.DeployID
	}
	return ""
}

type NamespaceSchemaResetRequest struct {
	// Name is the namespace name.
	// Reset schema to non-existent namespace will get 404.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *NamespaceSchemaResetRequest) Reset()         { *m = NamespaceSchemaResetRequest{} }
func (m *NamespaceSchemaResetRequest) String() string { return proto.CompactTextString(m) }
func (*NamespaceSchemaResetRequest) ProtoMessage()    {}
func (*NamespaceSchemaResetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorNamespace, []int{4}
}

func (m *NamespaceSchemaResetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NamespaceSchemaResetResponse struct {
}

func (m *NamespaceSchemaResetResponse) Reset()         { *m = NamespaceSchemaResetResponse{} }
func (m *NamespaceSchemaResetResponse) String() string { return proto.CompactTextString(m) }
func (*NamespaceSchemaResetResponse) ProtoMessage()    {}
func (*NamespaceSchemaResetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorNamespace, []int{5}
}

func init() {
	proto.RegisterType((*NamespaceGetResponse)(nil), "admin.NamespaceGetResponse")
	proto.RegisterType((*NamespaceAddRequest)(nil), "admin.NamespaceAddRequest")
	proto.RegisterType((*NamespaceSchemaAddRequest)(nil), "admin.NamespaceSchemaAddRequest")
	proto.RegisterType((*NamespaceSchemaAddResponse)(nil), "admin.NamespaceSchemaAddResponse")
	proto.RegisterType((*NamespaceSchemaResetRequest)(nil), "admin.NamespaceSchemaResetRequest")
	proto.RegisterType((*NamespaceSchemaResetResponse)(nil), "admin.NamespaceSchemaResetResponse")
}
func (m *NamespaceGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Registry != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Registry.Size()))
		n1, err := m.Registry.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *NamespaceAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Options != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Options.Size()))
		n2, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *NamespaceSchemaAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceSchemaAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.MsgName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.MsgName)))
		i += copy(dAtA[i:], m.MsgName)
	}
	if len(m.ProtoName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.ProtoName)))
		i += copy(dAtA[i:], m.ProtoName)
	}
	if len(m.ProtoMap) > 0 {
		for k, _ := range m.ProtoMap {
			dAtA[i] = 0x22
			i++
			v := m.ProtoMap[k]
			mapSize := 1 + len(k) + sovNamespace(uint64(len(k))) + 1 + len(v) + sovNamespace(uint64(len(v)))
			i = encodeVarintNamespace(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *NamespaceSchemaAddResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceSchemaAddResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DeployID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.DeployID)))
		i += copy(dAtA[i:], m.DeployID)
	}
	return i, nil
}

func (m *NamespaceSchemaResetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceSchemaResetRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *NamespaceSchemaResetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceSchemaResetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NamespaceGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Registry != nil {
		l = m.Registry.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceAddRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceSchemaAddRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.MsgName)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.ProtoName)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if len(m.ProtoMap) > 0 {
		for k, v := range m.ProtoMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovNamespace(uint64(len(k))) + 1 + len(v) + sovNamespace(uint64(len(v)))
			n += mapEntrySize + 1 + sovNamespace(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NamespaceSchemaAddResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.DeployID)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceSchemaResetRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceSchemaResetResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovNamespace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamespaceGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Registry == nil {
				m.Registry = &namespace1.Registry{}
			}
			if err := m.Registry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &namespace1.NamespaceOptions{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceSchemaAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceSchemaAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceSchemaAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtoName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProtoMap == nil {
				m.ProtoMap = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNamespace
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespace
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNamespace
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNamespace
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthNamespace
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNamespace(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNamespace
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ProtoMap[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceSchemaAddResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceSchemaAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceSchemaAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeployID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceSchemaResetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceSchemaResetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceSchemaResetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceSchemaResetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceSchemaResetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceSchemaResetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
				return 0, ErrInvalidLengthNamespace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamespace
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
				next, err := skipNamespace(dAtA[start:])
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
	ErrInvalidLengthNamespace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/query/generated/proto/admin/namespace.proto", fileDescriptorNamespace)
}

var fileDescriptorNamespace = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x1b, 0xff, 0x54, 0x3d, 0x52, 0x90, 0xd1, 0x8b, 0x34, 0x4a, 0x90, 0x5c, 0x79, 0x95,
	0xa1, 0x4a, 0x41, 0xda, 0xab, 0x4a, 0x8b, 0xb4, 0xd0, 0x3f, 0x4c, 0x5f, 0xa0, 0x49, 0xe6, 0x10,
	0x43, 0x4d, 0x26, 0x66, 0x92, 0x85, 0xbc, 0xc5, 0x3e, 0xd6, 0x5e, 0xee, 0x23, 0x2c, 0xee, 0x0b,
	0xec, 0x23, 0x2c, 0x19, 0x93, 0xb8, 0xeb, 0x2a, 0x7b, 0x37, 0xe7, 0x7c, 0xe7, 0xfb, 0x7d, 0x33,
	0x67, 0x60, 0xe5, 0x07, 0xe9, 0x26, 0x73, 0x6d, 0x4f, 0x84, 0x34, 0x5c, 0x70, 0x97, 0x86, 0x0b,
	0x2a, 0x13, 0x8f, 0xee, 0x32, 0x4c, 0x72, 0xea, 0x63, 0x84, 0x89, 0x93, 0x22, 0xa7, 0x71, 0x22,
	0x52, 0x41, 0x1d, 0x1e, 0x06, 0x11, 0x8d, 0x9c, 0x10, 0x65, 0xec, 0x78, 0x68, 0xab, 0x2e, 0x69,
	0xab, 0xb6, 0xb1, 0xbe, 0x80, 0xe2, 0x6e, 0x24, 0x38, 0xbe, 0x60, 0xd5, 0x94, 0x53, 0x9e, 0xb5,
	0x86, 0xd1, 0xaf, 0xaa, 0xb5, 0xc6, 0x94, 0xa1, 0x8c, 0x45, 0x24, 0x91, 0x50, 0xe8, 0x26, 0xe8,
	0x07, 0x32, 0x4d, 0x72, 0x5d, 0x9b, 0x6a, 0xb3, 0xfe, 0x7c, 0x68, 0x1f, 0xbd, 0xac, 0x94, 0x58,
	0x3d, 0x64, 0xfd, 0x83, 0x61, 0x0d, 0xfa, 0xc2, 0x39, 0xc3, 0x5d, 0x86, 0x32, 0x25, 0x04, 0x5a,
	0x85, 0x4d, 0x31, 0x7a, 0x4c, 0x9d, 0xc9, 0x47, 0xe8, 0x88, 0x38, 0x0d, 0x44, 0x24, 0xf5, 0x86,
	0x42, 0x8f, 0x9f, 0xa0, 0x6b, 0xc8, 0xef, 0xc3, 0x08, 0xab, 0x66, 0xad, 0x07, 0x0d, 0xde, 0xd7,
	0xea, 0x5f, 0x6f, 0x83, 0xa1, 0xf3, 0x4a, 0x90, 0x0e, 0x9d, 0x50, 0xfa, 0x85, 0x47, 0x05, 0xf5,
	0x58, 0x55, 0x92, 0x09, 0xf4, 0xd4, 0xfb, 0x95, 0xd6, 0x54, 0xda, 0xb1, 0x41, 0x7e, 0x40, 0x57,
	0x15, 0x3f, 0x9d, 0x58, 0x6f, 0x4d, 0x9b, 0xb3, 0xfe, 0xdc, 0xb6, 0xd5, 0xde, 0xed, 0x8b, 0xf9,
	0xf6, 0x9f, 0xd2, 0xf0, 0x2d, 0x52, 0x7b, 0xa9, 0xfc, 0xc6, 0x67, 0x78, 0xf7, 0x4c, 0x22, 0x03,
	0x68, 0xfe, 0xc7, 0xbc, 0xbc, 0x67, 0x71, 0x24, 0x23, 0x68, 0x5f, 0x39, 0xdb, 0xac, 0xba, 0xe4,
	0xa1, 0xf8, 0xd4, 0x58, 0x6a, 0xd6, 0x12, 0x8c, 0x73, 0x89, 0xe5, 0x1f, 0x19, 0xd0, 0xe5, 0x18,
	0x6f, 0x45, 0xfe, 0xfd, 0x6b, 0x89, 0xab, 0x6b, 0xeb, 0x03, 0x8c, 0x4f, 0x9c, 0x0c, 0x65, 0xf1,
	0xbf, 0x17, 0xb7, 0x65, 0x99, 0x30, 0x39, 0x6f, 0x39, 0xc4, 0xad, 0x06, 0x37, 0x7b, 0x53, 0xbb,
	0xdd, 0x9b, 0xda, 0xdd, 0xde, 0xd4, 0xae, 0xef, 0xcd, 0x37, 0xee, 0x5b, 0xf5, 0xca, 0xc5, 0x63,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xa6, 0x90, 0x2e, 0xd9, 0x02, 0x00, 0x00,
}