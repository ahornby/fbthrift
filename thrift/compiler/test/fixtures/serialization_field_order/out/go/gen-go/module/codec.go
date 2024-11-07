// Autogenerated by Thrift for thrift/compiler/test/fixtures/serialization_field_order/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module


import (
    "reflect"
    "sync"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = reflect.Ptr

// Premade codec specs
var (
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Foo *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Foo2 *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        FullName: "i32",
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_module_Foo = &thrift.TypeSpec{
        FullName: "module.Foo",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Foo",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewFoo() },
},

    }
    premadeCodecTypeSpec_module_Foo2 = &thrift.TypeSpec{
        FullName: "module.Foo2",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Foo2",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewFoo2() },
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Foo *thrift.StructSpec = nil
    premadeStructSpec_Foo2 *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Foo = &thrift.StructSpec{
    Name:                 "Foo",
    ScopedName:           "module.Foo",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.I32,
            Name:                 "field2",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.I32,
            Name:                 "field3",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.I32,
            Name:                 "field1",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
    },
    FieldSpecNameToIndex: map[string]int{
        "field2": 0,
        "field3": 1,
        "field1": 2,
    },
}
    premadeStructSpec_Foo2 = &thrift.StructSpec{
    Name:                 "Foo2",
    ScopedName:           "module.Foo2",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.I32,
            Name:                 "field2",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.I32,
            Name:                 "field3",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.I32,
            Name:                 "field1",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
    },
    FieldSpecNameToIndex: map[string]int{
        "field2": 0,
        "field3": 1,
        "field1": 2,
    },
}
})

var premadeCodecSpecsMapOnce = sync.OnceValue(
    func() map[string]*thrift.TypeSpec {
        // Relies on premade codec specs initialization
        premadeCodecSpecsInitOnce()

        fbthriftTypeSpecsMap := make(map[string]*thrift.TypeSpec)
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_i32.FullName] = premadeCodecTypeSpec_i32
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Foo.FullName] = premadeCodecTypeSpec_module_Foo
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Foo2.FullName] = premadeCodecTypeSpec_module_Foo2
        return fbthriftTypeSpecsMap
    },
)

func init() {
    premadeCodecSpecsInitOnce()
    premadeStructSpecsInitOnce()
}

// GetMetadataThriftType (INTERNAL USE ONLY).
// Returns metadata TypeSpec for a given full type name.
func GetCodecTypeSpec(fullName string) *thrift.TypeSpec {
    return premadeCodecSpecsMapOnce()[fullName]
}
