// Autogenerated by Thrift for thrift/compiler/test/fixtures/enums/src/module.thrift
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
    premadeCodecTypeSpec_module_Metasyntactic *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyEnum1 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyEnum2 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyEnum3 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyEnum4 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyBitmaskEnum1 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyBitmaskEnum2 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_set_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_SomeStruct *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyStruct *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_module_Metasyntactic = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_MyEnum1 = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_MyEnum2 = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_MyEnum3 = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_MyEnum4 = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_MyBitmaskEnum1 = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_MyBitmaskEnum2 = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_set_i32 = &thrift.TypeSpec{
        CodecSetSpec: &thrift.CodecSetSpec{
    ElementWireType: thrift.I32,
	ElementTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_module_SomeStruct = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewSomeStruct() },
},

    }
    premadeCodecTypeSpec_module_MyStruct = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyStruct() },
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_SomeStruct *thrift.StructSpec = nil
    premadeStructSpec_MyStruct *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_SomeStruct = &thrift.StructSpec{
    Name:                 "SomeStruct",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "reasonable",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Metasyntactic,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "fine",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Metasyntactic,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "questionable",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Metasyntactic,
            MustBeSetToSerialize: false,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.SET),
            Name:                 "tags",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_set_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
    },
    FieldSpecNameToIndex: map[string]int{
        "reasonable": 0,
        "fine": 1,
        "questionable": 2,
        "tags": 3,
    },
}
    premadeStructSpec_MyStruct = &thrift.StructSpec{
    Name:                 "MyStruct",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "me2_3",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnum2,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "me3_n3",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnum3,
            MustBeSetToSerialize: false,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "me1_t1",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnum1,
            MustBeSetToSerialize: false,
        },        {
            ID:                   6,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "me1_t2",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnum1,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        4: 2,
        6: 3,
    },
    FieldSpecNameToIndex: map[string]int{
        "me2_3": 0,
        "me3_n3": 1,
        "me1_t1": 2,
        "me1_t2": 3,
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
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.Metasyntactic", premadeCodecTypeSpec_module_Metasyntactic })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyEnum1", premadeCodecTypeSpec_module_MyEnum1 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyEnum2", premadeCodecTypeSpec_module_MyEnum2 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyEnum3", premadeCodecTypeSpec_module_MyEnum3 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyEnum4", premadeCodecTypeSpec_module_MyEnum4 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyBitmaskEnum1", premadeCodecTypeSpec_module_MyBitmaskEnum1 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyBitmaskEnum2", premadeCodecTypeSpec_module_MyBitmaskEnum2 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "i32", premadeCodecTypeSpec_i32 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.SomeStruct", premadeCodecTypeSpec_module_SomeStruct })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyStruct", premadeCodecTypeSpec_module_MyStruct })

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
