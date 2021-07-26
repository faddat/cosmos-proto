package fastreflection

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// fields are not populated if
// if scalar: value != zero value
// if msg value != nil
// if list len(list) != 0
// if map len(map) != 0
// if oneof: oneof != nil (if oneof is scalar do we need to check it??)
// if bytes: len(bytes) != 0
type hasGen struct {
	*protogen.GeneratedFile
	typeName string
	message  *protogen.Message
}

func (g *hasGen) genComments() {
	g.P("// Has reports whether a field is populated.")
	g.P("//")
	g.P("// Some fields have the property of nullability where it is possible to")
	g.P("// distinguish between the default value of a field and whether the field")
	g.P("// was explicitly populated with the default value. Singular message fields,")
	g.P("// member fields of a oneof, and proto2 scalar fields are nullable. Such")
	g.P("// fields are populated only if explicitly set.")
	g.P("//")
	g.P("// In other cases (aside from the nullable cases above),")
	g.P("// a proto3 scalar field is populated if it contains a non-zero value, and")
	g.P("// a repeated field is populated if it is non-empty.")
}

func (g *hasGen) generate() {
	g.genComments()
	g.P("func (x *", g.typeName, ") Has(fd ", protoreflectPkg.Ident("FieldDescriptor"), ") bool {")
	g.P("switch fd.FullName() {")
	for _, field := range g.message.Fields {
		g.genField(field)
	}
	g.P("default:")
	g.P("panic(", fmtPkg.Ident("Errorf"), "(\"message ", g.message.Desc.FullName(), " does not have field %s\", fd.Name()))")
	g.P("}")
	g.P("}")
}

func (g *hasGen) genField(field *protogen.Field) {
	g.P("case \"", field.Desc.FullName(), "\":")
	if field.Desc.HasPresence() || field.Desc.IsList() || field.Desc.IsMap() {
		g.genNullable(field)
		return
	}

	if field.Desc.Kind() == protoreflect.BytesKind {
		g.P("return len(x.", field.GoName, ") != 0")
		return
	}

	g.P("return x.", field.GoName, " != ", zeroValueForField(nil, field))
}

func (g *hasGen) genNullable(field *protogen.Field) {
	switch {
	case field.Desc.ContainingOneof() != nil:
		g.P("return x.", field.Oneof.GoName, " != nil")
	case field.Desc.IsMap(), field.Desc.IsList():
		g.P("return len(x.", field.GoName, ") != 0")
	case field.Desc.Kind() == protoreflect.MessageKind:
		g.P("return x.", field.GoName, " != nil")
	default:
		panic("unknown case")
	}
}