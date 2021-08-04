package generator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

/*
	╭━━━┳╮╱╭┳╮╱╱╭━━━┳━━━┳━━━╮
	┃╭━╮┃┃╱┃┃┃╱╱┃╭━╮┃╭━╮┃╭━╮┃
	┃╰━╯┃┃╱┃┃┃╱╱┃╰━━┫┃╱┃┃╰━╯┃
	┃╭━━┫┃╱┃┃┃╱╭╋━━╮┃╰━╯┃╭╮╭╯
	┃┃╱╱┃╰━╯┃╰━╯┃╰━╯┃╭━╮┃┃┃╰╮
	╰╯╱╱╰━━━┻━━━┻━━━┻╯╱╰┻╯╰━╯

	- many bytes, such speed. -
*/

type structTags [][2]string

func (tags structTags) String() string {
	if len(tags) == 0 {
		return ""
	}
	var ss []string
	for _, tag := range tags {
		// NOTE: When quoting the value, we need to make sure the backtick
		// character does not appear. Convert all cases to the escaped hex form.
		key := tag[0]
		val := strings.Replace(strconv.Quote(tag[1]), "`", `\x60`, -1)
		ss = append(ss, fmt.Sprintf("%s:%s", key, val))
	}
	return "`" + strings.Join(ss, " ") + "`"
}

// Standard library dependencies.
const (
	base64Package  = protogen.GoImportPath("encoding/base64")
	mathPackage    = protogen.GoImportPath("math")
	reflectPackage = protogen.GoImportPath("reflect")
	sortPackage    = protogen.GoImportPath("sort")
	stringsPackage = protogen.GoImportPath("strings")
	syncPackage    = protogen.GoImportPath("sync")
	timePackage    = protogen.GoImportPath("time")
	utf8Package    = protogen.GoImportPath("unicode/utf8")
)

// Protobuf library dependencies.
//
// These are declared as an interface type so that they can be more easily
// patched to support unique build environments that impose restrictions
// on the dependencies of generated source code.
var (
	FILENAME             string
	anySeen              bool
	protoPackage         goImportPath = protogen.GoImportPath("google.golang.org/protobuf/proto")
	protoifacePackage    goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoiface")
	protoimplPackage     goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoimpl")
	protojsonPackage     goImportPath = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	protoreflectPackage  goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	protoregistryPackage goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoregistry")
	errorsPackage        goImportPath = protogen.GoImportPath("errors")
)

type goImportPath interface {
	String() string
	Ident(string) protogen.GoIdent
}

func GenerateProtocGenGo(plugin *protogen.Plugin, g *GeneratedFile, file *protogen.File) *GeneratedFile {
	FILENAME = strings.ReplaceAll(file.GeneratedFilenamePrefix, "/", "_") // FIXME(fdymylja): produces bad output
	f := newFileInfo(file)
	genPackage(g, file.GoPackageName, file)
	genTopLevelVars(g, file.Messages)
	for _, enum := range f.allEnums {
		genEnum(g, f, enum)
	}
	for i, msg := range file.Messages {
		genMsgStruct(g, msg, i)
		g.P()
	}
	// genFileProtoTypes(g, file)
	genReflectFileDescriptor(plugin, g, f)
	return g
}

// trailingComment is like protogen.Comments, but lacks a trailing newline.
type trailingComment protogen.Comments

func genEnum(g *GeneratedFile, f *fileInfo, e *enumInfo) {
	// Enum type declaration.
	g.Annotate(e.GoIdent.GoName, e.Location)
	leadingComments := appendDeprecationSuffix(e.Comments.Leading,
		e.Desc.Options().(*descriptorpb.EnumOptions).GetDeprecated())
	g.P(leadingComments,
		"type ", e.GoIdent, " int32")

	// Enum value constants.
	g.P("const (")
	for _, value := range e.Values {
		g.Annotate(value.GoIdent.GoName, value.Location)
		leadingComments := appendDeprecationSuffix(value.Comments.Leading,
			value.Desc.Options().(*descriptorpb.EnumValueOptions).GetDeprecated())
		g.P(leadingComments,
			value.GoIdent, " ", e.GoIdent, " = ", value.Desc.Number(),
			trailingComment(value.Comments.Trailing))
	}
	g.P(")")
	g.P()

	// Enum value maps.
	g.P("// Enum value maps for ", e.GoIdent, ".")
	g.P("var (")
	g.P(e.GoIdent.GoName+"_name", " = map[int32]string{")
	for _, value := range e.Values {
		duplicate := ""
		if value.Desc != e.Desc.Values().ByNumber(value.Desc.Number()) {
			duplicate = "// Duplicate value: "
		}
		g.P(duplicate, value.Desc.Number(), ": ", strconv.Quote(string(value.Desc.Name())), ",")
	}
	g.P("}")
	g.P(e.GoIdent.GoName+"_value", " = map[string]int32{")
	for _, value := range e.Values {
		g.P(strconv.Quote(string(value.Desc.Name())), ": ", value.Desc.Number(), ",")
	}
	g.P("}")
	g.P(")")
	g.P()

	// Enum method.
	//
	// NOTE: A pointer value is needed to represent presence in proto2.
	// Since a proto2 message can reference a proto3 enum, it is useful to
	// always generate this method (even on proto3 enums) to support that case.
	g.P("func (x ", e.GoIdent, ") Enum() *", e.GoIdent, " {")
	g.P("p := new(", e.GoIdent, ")")
	g.P("*p = x")
	g.P("return p")
	g.P("}")
	g.P()

	// String method.
	g.P("func (x ", e.GoIdent, ") String() string {")
	g.P("return ", protoimplPackage.Ident("X"), ".EnumStringOf(x.Descriptor(), ", protoreflectPackage.Ident("EnumNumber"), "(x))")
	g.P("}")
	g.P()

	genEnumReflectMethods(g, f, e)

	// UnmarshalJSON method.
	if e.genJSONMethod && e.Desc.Syntax() == protoreflect.Proto2 {
		g.P("// Deprecated: Do not use.")
		g.P("func (x *", e.GoIdent, ") UnmarshalJSON(b []byte) error {")
		g.P("num, err := ", protoimplPackage.Ident("X"), ".UnmarshalJSONEnum(x.Descriptor(), b)")
		g.P("if err != nil {")
		g.P("return err")
		g.P("}")
		g.P("*x = ", e.GoIdent, "(num)")
		g.P("return nil")
		g.P("}")
		g.P()
	}

	// EnumDescriptor method.
	if e.genRawDescMethod {
		var indexes []string
		for i := 1; i < len(e.Location.Path); i += 2 {
			indexes = append(indexes, strconv.Itoa(int(e.Location.Path[i])))
		}
		g.P("// Deprecated: Use ", e.GoIdent, ".Descriptor instead.")
		g.P("func (", e.GoIdent, ") EnumDescriptor() ([]byte, []int) {")
		g.P("return ", rawDescVarName(f), "GZIP(), []int{", strings.Join(indexes, ","), "}")
		g.P("}")
		g.P()
		f.needRawDesc = true
	}
}

func genEnumReflectMethods(g *GeneratedFile, f *fileInfo, e *enumInfo) {
	idx := f.allEnumsByPtr[e]
	typesVar := enumTypesVarName(f)

	// Descriptor method.
	g.P("func (", e.GoIdent, ") Descriptor() ", protoreflectPackage.Ident("EnumDescriptor"), " {")
	g.P("return ", typesVar, "[", idx, "].Descriptor()")
	g.P("}")
	g.P()

	// Type method.
	g.P("func (", e.GoIdent, ") Type() ", protoreflectPackage.Ident("EnumType"), " {")
	g.P("return &", typesVar, "[", idx, "]")
	g.P("}")
	g.P()

	// Number method.
	g.P("func (x ", e.GoIdent, ") Number() ", protoreflectPackage.Ident("EnumNumber"), " {")
	g.P("return ", protoreflectPackage.Ident("EnumNumber"), "(x)")
	g.P("}")
	g.P()
}

func genPackage(g *GeneratedFile, packageName protogen.GoPackageName, file *protogen.File) {
	/*
		g.P("// Code generated by Pulsar \U0001FA90. DO NOT EDIT.")
		if bi, ok := debug.ReadBuildInfo(); ok {
			g.P("// Pulsar version: ", bi.Main.Version)
		}
		g.P("// source: ", file.Desc.Path())
		g.P()
	*/ // TODO(fdymylja): remove me
	g.P("package ", packageName)
	g.P()
}

func genTopLevelVars(g *GeneratedFile, msgs []*protogen.Message) {
	g.P("const (")
	g.P("// Verify that this generated code is sufficiently up-to-date.")
	g.P("_ = ", protoimplPackage.Ident("EnforceVersion"), "(", protoimpl.GenVersion, " - ", protoimplPackage.Ident("MinVersion"), ")")
	g.P("// Verify that runtime/protoimpl is sufficiently up-to-date.")
	g.P("_ = ", protoimplPackage.Ident("EnforceVersion"), "(", protoimplPackage.Ident("MaxVersion"), " - ", protoimpl.GenVersion, ")")
	g.P()
	g.P(")")
	g.P()
	g.P("var (")
	g.P("// Interface guards to verify each message implements proto message interface")
	for _, msg := range msgs {
		g.P("_ ", protoreflectPackage.Ident("Message"), " = &", msg.GoIdent.GoName, "{}")
	}
	g.P(")")
	g.P()
}

func genMsgStruct(g *GeneratedFile, msg *protogen.Message, index int) {
	g.P(msg.Comments.Leading, "type ", msg.GoIdent.GoName, " struct {")
	genStdFields(g, msg)
	g.P()
	for _, field := range msg.Fields {
		genField(g, field)
	}
	g.P("}")
	genReset(g, msg, index) // not sure wtf this is for
	g.P()
	genString(g, msg)
	g.P()
	genProtoMessage(g, msg)
	g.P()
	genProtoReflect(g, msg, index)
	g.P()
	genGetters(g, msg)
	g.P()
	genProtoMessageFunctions(g, msg)
}

func genProtoMessage(g *GeneratedFile, msg *protogen.Message) {
	g.P("func (x *", msg.GoIdent.GoName, ") ProtoMessage() {}")
}

func genReset(g *GeneratedFile, msg *protogen.Message, index int) {
	g.P("func (x *", msg.GoIdent.GoName, ") Reset() {")
	g.P("*x = ", msg.GoIdent.GoName, "{}")
	g.P("if ", protoimplPackage.Ident("UnsafeEnabled"), " {")
	g.P("mi := &file_", FILENAME, "_proto_msgTypes[", index, "]")
	g.P("ms := ", protoimplPackage.Ident("X.MessageStateOf"), "(", protoimplPackage.Ident("Pointer"), "(x))")
	g.P("ms.StoreMessageInfo(mi)")
	g.P("}")
	g.P("}")
}

func genStdFields(g *GeneratedFile, msg *protogen.Message) {
	g.P("state ", protoimplPackage.Ident("MessageState"))
	g.P("sizeCache ", protoimplPackage.Ident("SizeCache"))
	g.P("unknownFields ", protoimplPackage.Ident("UnknownFields"))
}

func genField(g *GeneratedFile, field *protogen.Field) {
	if oneof := field.Oneof; oneof != nil && !oneof.Desc.IsSynthetic() {
		// It would be a bit simpler to iterate over the oneofs below,
		// but generating the field here keeps the contents of the Go
		// struct in the same order as the contents of the source
		// .proto file.
		if oneof.Fields[0] != field {
			return // only generate for first appearance
		}
		tags := structTags{
			{"protobuf_oneof", string(oneof.Desc.Name())},
		}

		g.Annotate(field.GoIdent.GoName+"."+oneof.GoName, oneof.Location)
		leadingComments := oneof.Comments.Leading
		if leadingComments != "" {
			leadingComments += "\n"
		}
		ss := []string{fmt.Sprintf(" Types that are assignable to %s:\n", oneof.GoName)}
		for _, field := range oneof.Fields {
			ss = append(ss, "\t*"+field.GoIdent.GoName+"\n")
		}
		leadingComments += protogen.Comments(strings.Join(ss, ""))
		g.P(leadingComments, oneof.GoName, " ", oneOfInterfaceName(oneof), tags)
		return
	}

	goType, isPointer := getType(g, field)
	if isPointer {
		goType = "*" + goType
	}
	tags := structTags{
		{"protobuf", fieldProtobufTagValue(field)},
		{"json", fieldJSONTagValue(field)},
		{"yaml", fieldYAMLTagValue(field)}, // TODO: do we need this still?
	}
	g.P(field.Comments.Leading, field.GoName, " ", goType, tags)
}

func getType(g *GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		anySeen = true
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := getType(g, field.Message.Fields[0])
		valType, _ := getType(g, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

func fieldProtobufTagValue(field *protogen.Field) string {
	var enumName string
	if field.Desc.Kind() == protoreflect.EnumKind {
		enumName = protoimpl.X.LegacyEnumName(field.Enum.Desc)
	}
	return Marshal(field.Desc, enumName)
}

func fieldJSONTagValue(field *protogen.Field) string {
	return string(field.Desc.Name()) + ",omitempty"
}

func fieldYAMLTagValue(field *protogen.Field) string {
	return string(field.Desc.Name())
}

//func fieldYAMLTagValue(field *protogen.Field) string {
//	return "not impl"
//}

func genString(g *GeneratedFile, msg *protogen.Message) {
	g.P("func (x *", msg.GoIdent.GoName, ") String() string {")
	g.P("return ", protoimplPackage.Ident("X.MessageStringOf(x)"))
	g.P("}")
}

// oneOfInterfaceName returns the name of the interface type implemented by
// the oneof field value types.
func oneOfInterfaceName(oneof *protogen.Oneof) string {
	return "is" + oneof.GoIdent.GoName
}

func genGetters(g *GeneratedFile, msg *protogen.Message) {
	for _, field := range msg.Fields {
		goType, pointer := getType(g, field)
		// Getter for parent oneof.
		if oneof := field.Oneof; oneof != nil && oneof.Fields[0] == field && !oneof.Desc.IsSynthetic() {
			g.P("type ", oneOfInterfaceName(oneof), " interface {")
			g.P(oneOfInterfaceName(oneof), "()")
			g.P("}")
			g.P()
			g.Annotate(msg.GoIdent.GoName+".Get"+oneof.GoName, oneof.Location)
			g.P("func (m *", msg.GoIdent.GoName, ") Get", oneof.GoName, "() ", oneOfInterfaceName(oneof), " {")
			g.P("if m != nil {")
			g.P("return m.", oneof.GoName)
			g.P("}")
			g.P("return nil")
			g.P("}")
			g.P()

			for _, v := range oneof.Fields {
				g.P("type ", v.GoIdent, " struct {")
				g.P(v.GoName, " ", goType)
				g.P("}")
				g.P()
			}

			for _, one := range oneof.Fields {
				g.P("func (*", one.GoIdent, ") ", oneOfInterfaceName(oneof), "() {}")
				g.P()
			}
		}

		// Getter for message field.
		g.Annotate(msg.GoIdent.GoName+".Get"+field.GoName, field.Location)
		defaultValue := "var y " + goType
		switch {
		case field.Desc.IsWeak():
			g.P(field.Comments.Leading, "func (x *", msg.GoIdent, ") Get", field.GoName, "() ", protoPackage.Ident("Message"), "{")
			g.P("var w ", protoimplPackage.Ident("WeakFields"))
			g.P("if x != nil {")
			g.P("w = x.weakFields")
			g.P("}")
			g.P("return ", protoimplPackage.Ident("X"), ".GetWeak(w, ", field.Desc.Number(), ", ", strconv.Quote(string(field.Message.Desc.FullName())), ")")
			g.P("}")
		case field.Oneof != nil && !field.Oneof.Desc.IsSynthetic():
			g.P(field.Comments.Leading, "func (x *", msg.GoIdent, ") Get", field.GoName, "() ", goType, " {")
			g.P("if x, ok := x.Get", field.Oneof.GoName, "().(*", field.GoIdent, "); ok {")
			g.P("return x.", field.GoName)
			g.P("}")
			g.P(defaultValue)
			g.P("return y")
			g.P("}")
		default:
			g.P(field.Comments.Leading, "func (x *", msg.GoIdent, ") Get", field.GoName, "() ", goType, " {")
			if !field.Desc.HasPresence() || defaultValue == "nil" {
				g.P("if x != nil {")
			} else {
				g.P("if x != nil && x.", field.GoName, " != nil {")
			}
			star := ""
			if pointer {
				star = "*"
			}
			g.P("return ", star, " x.", field.GoName)
			g.P("}")
			g.P(defaultValue)
			g.P("return y")
			g.P("}")
		}
		g.P()
	}
}

func genProtoReflect(g *GeneratedFile, msg *protogen.Message, index int) {
	genFunctionSignature(g, true, "x", msg.GoIdent.GoName, "ProtoReflect", protoreflectPackage.Ident("Message"), []args{})
	g.P("mi := &file_", FILENAME, "_proto_msgTypes[", index, "]")
	g.P("if ", protoimplPackage.Ident("UnsafeEnabled"), " && x != nil {")
	g.P("ms := ", protoimplPackage.Ident("X.MessageStateOf"), "(", protoimplPackage.Ident("Pointer"), "(x))")
	g.P("if ms.LoadMessageInfo() == nil {")
	g.P("ms.StoreMessageInfo(mi)")
	g.P("}")
	g.P("return ms")
	g.P("}")
	g.P("return mi.MessageOf(x)")
	g.P("}")
}

type args struct {
	name string
	typ3 string
}

func genFunctionSignature(g *GeneratedFile, isPointer bool, receiverVar, receiverStruct, funcName string, returns protogen.GoIdent, argz []args) {
	if isPointer {
		receiverStruct = "*" + receiverStruct
	}
	var arguments string
	if len(argz) == 1 {
		arguments = fmt.Sprintf("%v %v", argz[0].name, argz[0].typ3)
	} else if len(argz) > 1 {
		for _, v := range argz {
			arguments += fmt.Sprintf("%v %v, ", v.name, v.typ3)
		}
		// get rid of trailing comma/space
		arguments = arguments[:len(arguments)-2]
	}

	g.P("func (", receiverVar, " ", receiverStruct, ") ", funcName, "(", arguments, ")", returns, " {")
}

// --------------------------------------------------------------------------------------------------------------------
// METHODS FOR STRUCT FIELD TAGGING
// --------------------------------------------------------------------------------------------------------------------

// Format is the serialization format used to represent the default value.
type Format int

const (
	_ Format = iota

	// Descriptor uses the serialization format that protoc uses with the
	// google.protobuf.FieldDescriptorProto.default_value field.
	Descriptor

	// GoTag uses the historical serialization format in Go struct field tags.
	GoTag
)

// Marshal encodes the protoreflect.FieldDescriptor as a tag.
//
// The enumName must be provided if the kind is an enum.
// Historically, the formulation of the enum "name" was the proto package
// dot-concatenated with the generated Go identifier for the enum type.
// Depending on the context on how Marshal is called, there are different ways
// through which that information is determined. As such it is the caller's
// responsibility to provide a function to obtain that information.
func Marshal(fd protoreflect.FieldDescriptor, enumName string) string {
	var tag []string
	switch fd.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind, protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Int64Kind, protoreflect.Uint64Kind:
		tag = append(tag, "varint")
	case protoreflect.Sint32Kind:
		tag = append(tag, "zigzag32")
	case protoreflect.Sint64Kind:
		tag = append(tag, "zigzag64")
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.FloatKind:
		tag = append(tag, "fixed32")
	case protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.DoubleKind:
		tag = append(tag, "fixed64")
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind:
		tag = append(tag, "bytes")
	case protoreflect.GroupKind:
		tag = append(tag, "group")
	}
	tag = append(tag, strconv.Itoa(int(fd.Number())))
	switch fd.Cardinality() {
	case protoreflect.Optional:
		tag = append(tag, "opt")
	case protoreflect.Required:
		tag = append(tag, "req")
	case protoreflect.Repeated:
		tag = append(tag, "rep")
	}
	if fd.IsPacked() {
		tag = append(tag, "packed")
	}
	name := string(fd.Name())
	if fd.Kind() == protoreflect.GroupKind {
		// The name of the FieldDescriptor for a group field is
		// lowercased. To find the original capitalization, we
		// look in the field's MessageType.
		name = string(fd.Message().Name())
	}
	tag = append(tag, "name="+name)
	if jsonName := fd.JSONName(); jsonName != "" && jsonName != name && !fd.IsExtension() {
		// NOTE: The jsonName != name condition is suspect, but it preserve
		// the exact same semantics from the previous generator.
		tag = append(tag, "json="+jsonName)
	}
	if fd.IsWeak() {
		tag = append(tag, "weak="+string(fd.Message().FullName()))
	}
	// The previous implementation does not tag extension fields as proto3,
	// even when the field is defined in a proto3 file. Match that behavior
	// for consistency.
	if fd.Syntax() == protoreflect.Proto3 && !fd.IsExtension() {
		tag = append(tag, "proto3")
	}
	if fd.Kind() == protoreflect.EnumKind && enumName != "" {
		tag = append(tag, "enum="+enumName)
	}
	if fd.ContainingOneof() != nil {
		tag = append(tag, "oneof")
	}
	// This must appear last in the tag, since commas in strings aren't escaped.
	if fd.HasDefault() {
		def, _ := DefValMarshal(fd.Default(), fd.DefaultEnumValue(), fd.Kind(), GoTag)
		tag = append(tag, "def="+def)
	}
	return strings.Join(tag, ",")
}

// DefValMarshal serializes v as the default string according to the given kind k.
// When specifying the Descriptor format for an enum kind, the associated
// enum value descriptor must be provided.
func DefValMarshal(v protoreflect.Value, ev protoreflect.EnumValueDescriptor, k protoreflect.Kind, f Format) (string, error) {
	switch k {
	case protoreflect.BoolKind:
		if f == GoTag {
			if v.Bool() {
				return "1", nil
			} else {
				return "0", nil
			}
		} else {
			if v.Bool() {
				return "true", nil
			} else {
				return "false", nil
			}
		}
	case protoreflect.EnumKind:
		if f == GoTag {
			return strconv.FormatInt(int64(v.Enum()), 10), nil
		} else {
			return string(ev.Name()), nil
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return strconv.FormatInt(v.Int(), 10), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return strconv.FormatUint(v.Uint(), 10), nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		f := v.Float()
		switch {
		case math.IsInf(f, -1):
			return "-inf", nil
		case math.IsInf(f, +1):
			return "inf", nil
		case math.IsNaN(f):
			return "nan", nil
		default:
			if k == protoreflect.FloatKind {
				return strconv.FormatFloat(f, 'g', -1, 32), nil
			} else {
				return strconv.FormatFloat(f, 'g', -1, 64), nil
			}
		}
	case protoreflect.StringKind:
		// String values are serialized as is without any escaping.
		return v.String(), nil
	case protoreflect.BytesKind:
		if s, ok := marshalBytes(v.Bytes()); ok {
			return s, nil
		}
	}
	return "", errors.New(fmt.Sprintf("could not format value for %v: %v", k, v))
}

// marshalBytes serializes bytes by using C escaping.
// To match the exact output of protoc, this is identical to the
// CEscape function in strutil.cc of the protoc source code.
func marshalBytes(b []byte) (string, bool) {
	var s []byte
	for _, c := range b {
		switch c {
		case '\n':
			s = append(s, `\n`...)
		case '\r':
			s = append(s, `\r`...)
		case '\t':
			s = append(s, `\t`...)
		case '"':
			s = append(s, `\"`...)
		case '\'':
			s = append(s, `\'`...)
		case '\\':
			s = append(s, `\\`...)
		default:
			if printableASCII := c >= 0x20 && c <= 0x7e; printableASCII {
				s = append(s, c)
			} else {
				s = append(s, fmt.Sprintf(`\%03o`, c)...)
			}
		}
	}
	return string(s), true
}

type fileInfo struct {
	*protogen.File

	allEnums      []*enumInfo
	allMessages   []*messageInfo
	allExtensions []*extensionInfo

	allEnumsByPtr         map[*enumInfo]int    // value is index into allEnums
	allMessagesByPtr      map[*messageInfo]int // value is index into allMessages
	allMessageFieldsByPtr map[*messageInfo]*structFields

	// needRawDesc specifies whether the generator should emit logic to provide
	// the legacy raw descriptor in GZIP'd form.
	// This is updated by enum and message generation logic as necessary,
	// and checked at the end of file generation.
	needRawDesc bool
}
type enumInfo struct {
	*protogen.Enum
	genJSONMethod    bool
	genRawDescMethod bool
}

type messageInfo struct {
	*protogen.Message
	genRawDescMethod  bool
	genExtRangeMethod bool
	isTracked         bool
	hasWeak           bool
}

type extensionInfo struct {
	*protogen.Extension
}

type structFields struct {
	count      int
	unexported map[int]string
}

func fileVarName(f *protogen.File, suffix string) string {
	prefix := f.GoDescriptorIdent.GoName
	_, n := utf8.DecodeRuneInString(prefix)
	prefix = strings.ToLower(prefix[:n]) + prefix[n:]
	return prefix + "_" + suffix
}
func rawDescVarName(f *fileInfo) string {
	return fileVarName(f.File, "rawDesc")
}
func goTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "goTypes")
}
func depIdxsVarName(f *fileInfo) string {
	return fileVarName(f.File, "depIdxs")
}
func enumTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "enumTypes")
}
func messageTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "msgTypes")
}
func extensionTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "extTypes")
}
func initFuncName(f *protogen.File) string {
	return fileVarName(f, "init")
}

func genReflectFileDescriptor(gen *protogen.Plugin, g *GeneratedFile, f *fileInfo) {
	g.P("var ", f.GoDescriptorIdent, " ", protoreflectPackage.Ident("FileDescriptor"))
	g.P()

	genFileDescriptor(gen, g, f)
	if len(f.allEnums) > 0 {
		g.P("var ", enumTypesVarName(f), " = make([]", protoimplPackage.Ident("EnumInfo"), ",", len(f.allEnums), ")")
	}
	if len(f.allMessages) > 0 {
		g.P("var ", messageTypesVarName(f), " = make([]", protoimplPackage.Ident("MessageInfo"), ",", len(f.allMessages), ")")
	}

	// Generate a unique list of Go types for all declarations and dependencies,
	// and the associated index into the type list for all dependencies.
	var goTypes []string
	var depIdxs []string
	seen := map[protoreflect.FullName]int{}
	genDep := func(name protoreflect.FullName, depSource string) {
		if depSource != "" {
			line := fmt.Sprintf("%d, // %d: %s -> %s", seen[name], len(depIdxs), depSource, name)
			depIdxs = append(depIdxs, line)
		}
	}
	genEnum := func(e *protogen.Enum, depSource string) {
		if e != nil {
			name := e.Desc.FullName()
			if _, ok := seen[name]; !ok {
				line := fmt.Sprintf("(%s)(0), // %d: %s", g.QualifiedGoIdent(e.GoIdent), len(goTypes), name)
				goTypes = append(goTypes, line)
				seen[name] = len(seen)
			}
			if depSource != "" {
				genDep(name, depSource)
			}
		}
	}
	genMessage := func(m *protogen.Message, depSource string) {
		if m != nil {
			name := m.Desc.FullName()
			if _, ok := seen[name]; !ok {
				line := fmt.Sprintf("(*%s)(nil), // %d: %s", g.QualifiedGoIdent(m.GoIdent), len(goTypes), name)
				if m.Desc.IsMapEntry() {
					// Map entry messages have no associated Go type.
					line = fmt.Sprintf("nil, // %d: %s", len(goTypes), name)
				}
				goTypes = append(goTypes, line)
				seen[name] = len(seen)
			}
			if depSource != "" {
				genDep(name, depSource)
			}
		}
	}

	// This ordering is significant.
	// See filetype.TypeBuilder.DependencyIndexes.
	type offsetEntry struct {
		start int
		name  string
	}
	var depOffsets []offsetEntry
	for _, enum := range f.allEnums {
		genEnum(enum.Enum, "")
	}
	for _, message := range f.allMessages {
		genMessage(message.Message, "")
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "field type_name"})
	for _, message := range f.allMessages {
		for _, field := range message.Fields {
			if field.Desc.IsWeak() {
				continue
			}
			source := string(field.Desc.FullName())
			genEnum(field.Enum, source+":type_name")
			genMessage(field.Message, source+":type_name")
		}
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "extension extendee"})
	for _, extension := range f.allExtensions {
		source := string(extension.Desc.FullName())
		genMessage(extension.Extendee, source+":extendee")
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "extension type_name"})
	for _, extension := range f.allExtensions {
		source := string(extension.Desc.FullName())
		genEnum(extension.Enum, source+":type_name")
		genMessage(extension.Message, source+":type_name")
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "method input_type"})
	for _, service := range f.Services {
		for _, method := range service.Methods {
			source := string(method.Desc.FullName())
			genMessage(method.Input, source+":input_type")
		}
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "method output_type"})
	for _, service := range f.Services {
		for _, method := range service.Methods {
			source := string(method.Desc.FullName())
			genMessage(method.Output, source+":output_type")
		}
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), ""})
	for i := len(depOffsets) - 2; i >= 0; i-- {
		curr, next := depOffsets[i], depOffsets[i+1]
		depIdxs = append(depIdxs, fmt.Sprintf("%d, // [%d:%d] is the sub-list for %s",
			curr.start, curr.start, next.start, curr.name))
	}
	if len(depIdxs) > math.MaxInt32 {
		panic("too many dependencies") // sanity check
	}

	g.P("var ", goTypesVarName(f), " = []interface{}{")
	for _, s := range goTypes {
		g.P(s)
	}
	g.P("}")

	g.P("var ", depIdxsVarName(f), " = []int32{")
	for _, s := range depIdxs {
		g.P(s)
	}
	g.P("}")

	g.P("func init() { ", initFuncName(f.File), "() }")

	g.P("func ", initFuncName(f.File), "() {")
	g.P("if ", f.GoDescriptorIdent, " != nil {")
	g.P("return")
	g.P("}")

	// Ensure that initialization functions for different files in the same Go
	// package run in the correct order: Call the init funcs for every .proto file
	// imported by this one that is in the same Go package.
	for i, imps := 0, f.Desc.Imports(); i < imps.Len(); i++ {
		impFile := gen.FilesByPath[imps.Get(i).Path()]
		if impFile.GoImportPath != f.GoImportPath {
			continue
		}
		g.P(initFuncName(impFile), "()")
	}

	if len(f.allMessages) > 0 {
		// Populate MessageInfo.Exporters.
		g.P("if !", protoimplPackage.Ident("UnsafeEnabled"), " {")
		for _, message := range f.allMessages {
			if sf := f.allMessageFieldsByPtr[message]; len(sf.unexported) > 0 {
				idx := f.allMessagesByPtr[message]
				typesVar := messageTypesVarName(f)

				g.P(typesVar, "[", idx, "].Exporter = func(v interface{}, i int) interface{} {")
				g.P("switch v := v.(*", message.GoIdent, "); i {")
				for i := 0; i < sf.count; i++ {
					if name := sf.unexported[i]; name != "" {
						g.P("case ", i, ": return &v.", name)
					}
				}
				g.P("default: return nil")
				g.P("}")
				g.P("}")
			}
		}
		g.P("}")

		// Populate MessageInfo.OneofWrappers.
		for _, message := range f.allMessages {
			if len(message.Oneofs) > 0 {
				idx := f.allMessagesByPtr[message]
				typesVar := messageTypesVarName(f)

				// Associate the wrapper types by directly passing them to the MessageInfo.
				g.P(typesVar, "[", idx, "].OneofWrappers = []interface{} {")
				for _, oneof := range message.Oneofs {
					if !oneof.Desc.IsSynthetic() {
						for _, field := range oneof.Fields {
							g.P("(*", field.GoIdent, ")(nil),")
						}
					}
				}
				g.P("}")
			}
		}
	}

	g.P("type x struct{}")
	g.P("out := ", protoimplPackage.Ident("TypeBuilder"), "{")
	g.P("File: ", protoimplPackage.Ident("DescBuilder"), "{")
	g.P("GoPackagePath: ", reflectPackage.Ident("TypeOf"), "(x{}).PkgPath(),")
	g.P("RawDescriptor: ", rawDescVarName(f), ",")
	g.P("NumEnums: ", len(f.allEnums), ",")
	g.P("NumMessages: ", len(f.allMessages), ",")
	g.P("NumExtensions: ", len(f.allExtensions), ",")
	g.P("NumServices: ", len(f.Services), ",")
	g.P("},")
	g.P("GoTypes: ", goTypesVarName(f), ",")
	g.P("DependencyIndexes: ", depIdxsVarName(f), ",")
	if len(f.allEnums) > 0 {
		g.P("EnumInfos: ", enumTypesVarName(f), ",")
	}
	if len(f.allMessages) > 0 {
		g.P("MessageInfos: ", messageTypesVarName(f), ",")
	}
	if len(f.allExtensions) > 0 {
		g.P("ExtensionInfos: ", extensionTypesVarName(f), ",")
	}
	g.P("}.Build()")
	g.P(f.GoDescriptorIdent, " = out.File")

	// Set inputs to nil to allow GC to reclaim resources.
	g.P(rawDescVarName(f), " = nil")
	g.P(goTypesVarName(f), " = nil")
	g.P(depIdxsVarName(f), " = nil")
	g.P("}")
}

func genFileDescriptor(gen *protogen.Plugin, g *GeneratedFile, f *fileInfo) {
	descProto := proto.Clone(f.Proto).(*descriptorpb.FileDescriptorProto)
	descProto.SourceCodeInfo = nil // drop source code information

	b, err := proto.MarshalOptions{AllowPartial: true, Deterministic: true}.Marshal(descProto)
	if err != nil {
		gen.Error(err)
		return
	}

	g.P("var ", rawDescVarName(f), " = []byte{")
	for len(b) > 0 {
		n := 16
		if n > len(b) {
			n = len(b)
		}

		s := ""
		for _, c := range b[:n] {
			s += fmt.Sprintf("0x%02x,", c)
		}
		g.P(s)

		b = b[n:]
	}
	g.P("}")
	g.P()

	if f.needRawDesc {
		onceVar := rawDescVarName(f) + "Once"
		dataVar := rawDescVarName(f) + "Data"
		g.P("var (")
		g.P(onceVar, " ", syncPackage.Ident("Once"))
		g.P(dataVar, " = ", rawDescVarName(f))
		g.P(")")
		g.P()

		g.P("func ", rawDescVarName(f), "GZIP() []byte {")
		g.P(onceVar, ".Do(func() {")
		g.P(dataVar, " = ", protoimplPackage.Ident("X"), ".CompressGZIP(", dataVar, ")")
		g.P("})")
		g.P("return ", dataVar)
		g.P("}")
		g.P()
	}
}

func newFileInfo(file *protogen.File) *fileInfo {
	f := &fileInfo{File: file}

	// Collect all enums, messages, and extensions in "flattened ordering".
	// See filetype.TypeBuilder.
	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}
	initEnumInfos := func(enums []*protogen.Enum) {
		for _, enum := range enums {
			f.allEnums = append(f.allEnums, newEnumInfo(f, enum))
		}
	}
	initMessageInfos := func(messages []*protogen.Message) {
		for _, message := range messages {
			f.allMessages = append(f.allMessages, newMessageInfo(f, message))
		}
	}
	initExtensionInfos := func(extensions []*protogen.Extension) {
		for _, extension := range extensions {
			f.allExtensions = append(f.allExtensions, newExtensionInfo(f, extension))
		}
	}
	initEnumInfos(f.Enums)
	initMessageInfos(f.Messages)
	initExtensionInfos(f.Extensions)
	walkMessages(f.Messages, func(m *protogen.Message) {
		initEnumInfos(m.Enums)
		initMessageInfos(m.Messages)
		initExtensionInfos(m.Extensions)
	})

	// Derive a reverse mapping of enum and message pointers to their index
	// in allEnums and allMessages.
	if len(f.allEnums) > 0 {
		f.allEnumsByPtr = make(map[*enumInfo]int)
		for i, e := range f.allEnums {
			f.allEnumsByPtr[e] = i
		}
	}
	if len(f.allMessages) > 0 {
		f.allMessagesByPtr = make(map[*messageInfo]int)
		f.allMessageFieldsByPtr = make(map[*messageInfo]*structFields)
		for i, m := range f.allMessages {
			f.allMessagesByPtr[m] = i
			f.allMessageFieldsByPtr[m] = new(structFields)
		}
	}

	return f
}

func newEnumInfo(f *fileInfo, enum *protogen.Enum) *enumInfo {
	e := &enumInfo{Enum: enum}
	e.genJSONMethod = true
	e.genRawDescMethod = true
	return e
}

func newMessageInfo(f *fileInfo, message *protogen.Message) *messageInfo {
	m := &messageInfo{Message: message}
	m.genRawDescMethod = true
	m.genExtRangeMethod = true
	m.isTracked = isTrackedMessage(m)
	for _, field := range m.Fields {
		m.hasWeak = m.hasWeak || field.Desc.IsWeak()
	}
	return m
}

// isTrackedMessage reports whether field tracking is enabled on the message.
func isTrackedMessage(m *messageInfo) (tracked bool) {
	const trackFieldUse_fieldNumber = 37383685

	// Decode the option from unknown fields to avoid a dependency on the
	// annotation proto from protoc-gen-go.
	b := m.Desc.Options().(*descriptorpb.MessageOptions).ProtoReflect().GetUnknown()
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		b = b[n:]
		if num == trackFieldUse_fieldNumber && typ == protowire.VarintType {
			v, _ := protowire.ConsumeVarint(b)
			tracked = protowire.DecodeBool(v)
		}
		m := protowire.ConsumeFieldValue(num, typ, b)
		b = b[m:]
	}
	return tracked
}

func newExtensionInfo(f *fileInfo, extension *protogen.Extension) *extensionInfo {
	x := &extensionInfo{Extension: extension}
	return x
}

// appendDeprecationSuffix optionally appends a deprecation notice as a suffix.
func appendDeprecationSuffix(prefix protogen.Comments, deprecated bool) protogen.Comments {
	if !deprecated {
		return prefix
	}
	if prefix != "" {
		prefix += "\n"
	}
	return prefix + " Deprecated: Do not use.\n"
}
