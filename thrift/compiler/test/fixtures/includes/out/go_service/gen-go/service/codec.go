// Autogenerated by Thrift for thrift/compiler/test/fixtures/includes/src/service.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package service


import (
    "reflect"
    "sync"

    module "module"
    includes "includes"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

var _ = module.GoUnusedProtection__
var _ = includes.GoUnusedProtection__
// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = reflect.Ptr

// Premade codec specs
var (
    premadeCodecTypeSpec_service_IncludesIncluded *thrift.TypeSpec = nil
    premadeCodecTypeSpec_service_IncludesTransitiveFoo *thrift.TypeSpec = nil
    premadeCodecTypeSpec_void *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_service_IncludesIncluded = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: includes.GetCodecTypeSpec("includes.Included"),
},

    }
    premadeCodecTypeSpec_service_IncludesTransitiveFoo = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: includes.GetCodecTypeSpec("includes.TransitiveFoo"),
},

    }
    premadeCodecTypeSpec_void = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_VOID,
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_reqMyServiceQuery *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceQuery *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceHasArgDocs *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceHasArgDocs *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_reqMyServiceQuery = &thrift.StructSpec{
    Name:                 "reqMyServiceQuery",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "s",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        module.GetCodecTypeSpec("module.MyStruct"),
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "i",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        includes.GetCodecTypeSpec("includes.Included"),
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
    },
    FieldSpecNameToIndex: map[string]int{
        "s": 0,
        "i": 1,
    },
}
    premadeStructSpec_respMyServiceQuery = &thrift.StructSpec{
    Name:                 "respMyServiceQuery",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqMyServiceHasArgDocs = &thrift.StructSpec{
    Name:                 "reqMyServiceHasArgDocs",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "s",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        module.GetCodecTypeSpec("module.MyStruct"),
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "i",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        includes.GetCodecTypeSpec("includes.Included"),
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
    },
    FieldSpecNameToIndex: map[string]int{
        "s": 0,
        "i": 1,
    },
}
    premadeStructSpec_respMyServiceHasArgDocs = &thrift.StructSpec{
    Name:                 "respMyServiceHasArgDocs",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
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
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "service.IncludesIncluded", premadeCodecTypeSpec_service_IncludesIncluded })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "service.IncludesTransitiveFoo", premadeCodecTypeSpec_service_IncludesTransitiveFoo })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "void", premadeCodecTypeSpec_void })

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
