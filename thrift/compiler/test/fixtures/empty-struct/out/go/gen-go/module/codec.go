// Autogenerated by Thrift for thrift/compiler/test/fixtures/empty-struct/src/module.thrift
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
    premadeCodecTypeSpec_module_Empty = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "module.Empty",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Empty",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewEmpty() },
},

        }
    }()
    premadeCodecTypeSpec_module_Nada = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "module.Nada",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Nada",
    IsUnion:    true,
    NewFunc:    func() thrift.Struct { return NewNada() },
},

        }
    }()
)

// Premade struct specs
var (
    premadeStructSpec_Empty = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Empty",
    ScopedName:           "module.Empty",
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
    premadeStructSpec_Nada = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Nada",
    ScopedName:           "module.Nada",
    IsUnion:              true,
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
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Empty)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Nada)
    return fbthriftResults
}()

var premadeCodecSpecsMap = func() map[string]*thrift.TypeSpec {
    fbthriftTypeSpecsMap := make(map[string]*thrift.TypeSpec)
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Empty.FullName] = premadeCodecTypeSpec_module_Empty
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Nada.FullName] = premadeCodecTypeSpec_module_Nada
    return fbthriftTypeSpecsMap
}()

// GetMetadataThriftType (INTERNAL USE ONLY).
// Returns metadata TypeSpec for a given full type name.
func GetCodecTypeSpec(fullName string) *thrift.TypeSpec {
    return premadeCodecSpecsMap[fullName]
}
