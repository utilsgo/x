package types

import (
	"go/types"
	"reflect"
)

func FromRType(rtype reflect.Type) *RType {
	return &RType{Type: rtype}
}

type RType struct {
	reflect.Type
}

func (rtype *RType) Unwrap() any {
	return rtype.Type
}

func (rtype *RType) Method(i int) Method {
	return &RMethod{Method: rtype.Type.Method(i)}
}

func (rtype *RType) MethodByName(name string) (Method, bool) {
	if m, ok := rtype.Type.MethodByName(name); ok {
		return &RMethod{Method: m}, true
	}
	return nil, false
}

func (rtype *RType) In(i int) Type {
	return FromRType(rtype.Type.In(i))
}

func (rtype *RType) Out(i int) Type {
	return FromRType(rtype.Type.Out(i))
}

func (rtype *RType) Implements(u Type) bool {
	switch x := u.(type) {
	case *RType:
		return rtype.Type.Implements(x.Type)
	case *TType:
		if rtype.PkgPath() == "" {
			return false
		}
		if i, ok := x.Type.(*types.Interface); ok {
			return types.Implements(NewTypesTypeFromReflectType(rtype.Type), i)
		}
	}
	return false
}

func (rtype *RType) AssignableTo(u Type) bool {
	return rtype.Type.AssignableTo(u.(*RType).Type)
}

func (rtype *RType) ConvertibleTo(u Type) bool {
	return rtype.Type.ConvertibleTo(u.(*RType).Type)
}

func (rtype *RType) Field(i int) StructField {
	return &RStructField{
		StructField: rtype.Type.Field(i),
	}
}

func (rtype *RType) FieldByName(name string) (StructField, bool) {
	if sf, ok := rtype.Type.FieldByName(name); ok {
		return &RStructField{
			StructField: sf,
		}, true
	}
	return nil, false
}

func (rtype *RType) FieldByNameFunc(match func(string) bool) (StructField, bool) {
	if sf, ok := rtype.Type.FieldByNameFunc(match); ok {
		return &RStructField{
			StructField: sf,
		}, true
	}
	return nil, false
}

func (rtype *RType) Key() Type {
	return FromRType(rtype.Type.Key())
}

func (rtype *RType) Elem() Type {
	return FromRType(rtype.Type.Elem())
}

func (rtype *RType) String() string {
	return typeString(rtype)
}

type RStructField struct {
	StructField reflect.StructField
}

func (f *RStructField) PkgPath() string {
	return f.StructField.PkgPath
}

func (f *RStructField) Name() string {
	return f.StructField.Name
}

func (f *RStructField) Tag() reflect.StructTag {
	return f.StructField.Tag
}

func (f *RStructField) Type() Type {
	return FromRType(f.StructField.Type)
}

func (f *RStructField) Anonymous() bool {
	return f.StructField.Anonymous
}

type RMethod struct {
	Method reflect.Method
}

func (m *RMethod) PkgPath() string {
	return m.Method.PkgPath
}

func (m *RMethod) Name() string {
	return m.Method.Name
}

func (m *RMethod) Type() Type {
	return FromRType(m.Method.Type)
}
