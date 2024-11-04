// Autogenerated by Thrift for thrift/compiler/test/fixtures/namespace/src/emptyns.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package emptyns


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
    premadeCodecTypeSpec_i64 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_emptyns_Foo *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_i64 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I64,
},

    }
    premadeCodecTypeSpec_emptyns_Foo = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewFoo() },
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Foo *thrift.StructSpec = nil
    premadeStructSpec_reqTestServiceInit *thrift.StructSpec = nil
    premadeStructSpec_respTestServiceInit *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Foo = &thrift.StructSpec{
    Name:                 "Foo",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "MyInt",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "MyInt": 0,
    },
}
    premadeStructSpec_reqTestServiceInit = &thrift.StructSpec{
    Name:                 "reqTestServiceInit",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "int1",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "int1": 0,
    },
}
    premadeStructSpec_respTestServiceInit = &thrift.StructSpec{
    Name:                 "respTestServiceInit",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        0: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "success": 0,
    },
}
})

// Helper type to allow us to store codec specs in a slice at compile time,
// and put them in a map at runtime. See comment at the top of template
// about a compilation limitation that affects map literals.
type codecSpecWithFullName struct {
    fullName string
    typeSpec *thrift.TypeSpec
}

var premadeCodecSpecsMapOnce = sync.OnceValue(
    func() map[string]*thrift.TypeSpec {
        // Relies on premade codec specs initialization
        premadeCodecSpecsInitOnce()

        codecSpecsWithFullName := make([]codecSpecWithFullName, 0)
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "i64", premadeCodecTypeSpec_i64 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "emptyns.Foo", premadeCodecTypeSpec_emptyns_Foo })

        fbthriftTypeSpecsMap := make(map[string]*thrift.TypeSpec, len(codecSpecsWithFullName))
        for _, value := range codecSpecsWithFullName {
            fbthriftTypeSpecsMap[value.fullName] = value.typeSpec
        }
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
