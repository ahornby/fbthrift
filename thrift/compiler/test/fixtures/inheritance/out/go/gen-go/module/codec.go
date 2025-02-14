// Autogenerated by Thrift for thrift/compiler/test/fixtures/inheritance/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module


import (
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.VOID

// Premade codec specs
var (
    premadeCodecTypeSpec_void = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "void",
            CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_VOID,
},

        }
    }()
)

// Premade struct specs
var (
    premadeStructSpec_reqMyRootDoRoot = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "reqMyRootDoRoot",
    ScopedName:           "module.reqMyRootDoRoot",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    }()
    premadeStructSpec_respMyRootDoRoot = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "respMyRootDoRoot",
    ScopedName:           "module.respMyRootDoRoot",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    }()
    premadeStructSpec_reqMyNodeDoMid = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "reqMyNodeDoMid",
    ScopedName:           "module.reqMyNodeDoMid",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    }()
    premadeStructSpec_respMyNodeDoMid = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "respMyNodeDoMid",
    ScopedName:           "module.respMyNodeDoMid",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    }()
    premadeStructSpec_reqMyLeafDoLeaf = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "reqMyLeafDoLeaf",
    ScopedName:           "module.reqMyLeafDoLeaf",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    }()
    premadeStructSpec_respMyLeafDoLeaf = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "respMyLeafDoLeaf",
    ScopedName:           "module.respMyLeafDoLeaf",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    }()
)

// Premade slice of all struct specs
var premadeStructSpecs = func() []*thrift.StructSpec {
    fbthriftResults := make([]*thrift.StructSpec, 0)
    return fbthriftResults
}()

var premadeCodecSpecsMap = func() map[string]*thrift.TypeSpec {
    fbthriftTypeSpecsMap := make(map[string]*thrift.TypeSpec)
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_void.FullName] = premadeCodecTypeSpec_void
    return fbthriftTypeSpecsMap
}()

// GetMetadataThriftType (INTERNAL USE ONLY).
// Returns metadata TypeSpec for a given full type name.
func GetCodecTypeSpec(fullName string) *thrift.TypeSpec {
    return premadeCodecSpecsMap[fullName]
}
