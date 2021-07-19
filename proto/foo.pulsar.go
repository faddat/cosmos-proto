package examples

import (
	errors "errors"
	fmt "fmt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	bits "math/bits"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var (
	// Interface guards to verify each message implements proto message interface
	_ protoreflect.Message = &Bar{}
	_ protoreflect.Message = &Hello{}
)

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baz string `protobuf:"bytes,1,opt,name=baz,proto3" json:"baz,omitempty" yaml:"baz"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (x *Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Bar) GetBaz() string {
	if x != nil {
		return x.Baz
	}
	var y string
	return y
}

// returns the fast methods for the message
func (x Bar) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(*Bar)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("size error: Bar does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*Bar)
			if !ok {
				return protoiface.UnmarshalOutput{}, errors.New("marshal error: Bar does not implement the protoreflect.Message interface")
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x Bar) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x Bar) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x Bar) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x Bar) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x Bar) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x Bar) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x Bar) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *Bar) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.Name() {
	case "baz":
		return protoreflect.ValueOfString(x.Baz)
	default:
		panic(fmt.Errorf("message cosmos.proto.Bar does not contain field %s", descriptor.Name()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x Bar) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x Bar) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x Bar) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x Bar) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x Bar) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x Bar) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x Bar) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x Bar) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	World    string `protobuf:"bytes,1,opt,name=world,proto3" json:"world,omitempty" yaml:"world"`
	Universe bool   `protobuf:"varint,2,opt,name=universe,proto3" json:"universe,omitempty" yaml:"universe"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (x *Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Hello) GetWorld() string {
	if x != nil {
		return x.World
	}
	var y string
	return y
}

func (x *Hello) GetUniverse() bool {
	if x != nil {
		return x.Universe
	}
	var y bool
	return y
}

// returns the fast methods for the message
func (x Hello) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(*Hello)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("size error: Hello does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*Hello)
			if !ok {
				return protoiface.UnmarshalOutput{}, errors.New("marshal error: Hello does not implement the protoreflect.Message interface")
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x Hello) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x Hello) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x Hello) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x Hello) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x Hello) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x Hello) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x Hello) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *Hello) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.Name() {
	case "world":
		return protoreflect.ValueOfString(x.World)
	case "universe":
		return protoreflect.ValueOfBool(x.Universe)
	default:
		panic(fmt.Errorf("message cosmos.proto.Hello does not contain field %s", descriptor.Name()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x Hello) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x Hello) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x Hello) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x Hello) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x Hello) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x Hello) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x Hello) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x Hello) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

var File_foo_proto protoreflect.FileDescriptor

var file_foo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x42, 0x61, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62,
	0x61, 0x7a, 0x22, 0x39, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_foo_proto_goTypes = []interface{}{
	(*Bar)(nil),   // 0: cosmos.proto.Bar
	(*Hello)(nil), // 1: cosmos.proto.Hello
}
var file_foo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_foo_proto_init() }
func file_foo_proto_init() {
	if File_foo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foo_proto_goTypes,
		DependencyIndexes: file_foo_proto_depIdxs,
		MessageInfos:      file_foo_proto_msgTypes,
	}.Build()
	File_foo_proto = out.File
	file_foo_proto_rawDesc = nil
	file_foo_proto_goTypes = nil
	file_foo_proto_depIdxs = nil
}
func (m *Bar) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bar) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Baz) > 0 {
		i -= len(m.Baz)
		copy(dAtA[i:], m.Baz)
		i = encodeVarint(dAtA, i, uint64(len(m.Baz)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Hello) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hello) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Hello) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Universe {
		i--
		if m.Universe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.World) > 0 {
		i -= len(m.World)
		copy(dAtA[i:], m.World)
		i = encodeVarint(dAtA, i, uint64(len(m.World)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Baz)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func (m *Hello) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.World)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Universe {
		n += 2
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bar) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: Bar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Baz", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Baz = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Hello) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: Hello: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hello: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field World", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.World = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Universe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
			m.Universe = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
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
					return 0, ErrIntOverflow
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
					return 0, ErrIntOverflow
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
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)
