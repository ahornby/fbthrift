// Autogenerated by Thrift for thrift/annotation/rust.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package rust


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
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Name *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Copy *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_RequestContext *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Arc *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Box *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Exhaustive *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Ord *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_NewType *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Type *thrift.TypeSpec = nil
    premadeCodecTypeSpec_bool *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Serde *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Mod *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Adapter *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_Derive *thrift.TypeSpec = nil
    premadeCodecTypeSpec_rust_ServiceExn *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_rust_Name = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewName() },
},

    }
    premadeCodecTypeSpec_rust_Copy = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewCopy() },
},

    }
    premadeCodecTypeSpec_rust_RequestContext = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewRequestContext() },
},

    }
    premadeCodecTypeSpec_rust_Arc = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewArc() },
},

    }
    premadeCodecTypeSpec_rust_Box = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewBox() },
},

    }
    premadeCodecTypeSpec_rust_Exhaustive = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewExhaustive() },
},

    }
    premadeCodecTypeSpec_rust_Ord = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewOrd() },
},

    }
    premadeCodecTypeSpec_rust_NewType = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewNewType_() },
},

    }
    premadeCodecTypeSpec_rust_Type = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewType() },
},

    }
    premadeCodecTypeSpec_bool = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_BOOL,
},

    }
    premadeCodecTypeSpec_rust_Serde = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewSerde() },
},

    }
    premadeCodecTypeSpec_rust_Mod = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMod() },
},

    }
    premadeCodecTypeSpec_rust_Adapter = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewAdapter() },
},

    }
    premadeCodecTypeSpec_list_string = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.STRING,
	ElementTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_rust_Derive = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewDerive() },
},

    }
    premadeCodecTypeSpec_rust_ServiceExn = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewServiceExn() },
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Name *thrift.StructSpec = nil
    premadeStructSpec_Copy *thrift.StructSpec = nil
    premadeStructSpec_RequestContext *thrift.StructSpec = nil
    premadeStructSpec_Arc *thrift.StructSpec = nil
    premadeStructSpec_Box *thrift.StructSpec = nil
    premadeStructSpec_Exhaustive *thrift.StructSpec = nil
    premadeStructSpec_Ord *thrift.StructSpec = nil
    premadeStructSpec_NewType *thrift.StructSpec = nil
    premadeStructSpec_Type *thrift.StructSpec = nil
    premadeStructSpec_Serde *thrift.StructSpec = nil
    premadeStructSpec_Mod *thrift.StructSpec = nil
    premadeStructSpec_Adapter *thrift.StructSpec = nil
    premadeStructSpec_Derive *thrift.StructSpec = nil
    premadeStructSpec_ServiceExn *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Name = &thrift.StructSpec{
    Name:                 "Name",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "name",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "name": 0,
    },
}
    premadeStructSpec_Copy = &thrift.StructSpec{
    Name:                 "Copy",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_RequestContext = &thrift.StructSpec{
    Name:                 "RequestContext",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_Arc = &thrift.StructSpec{
    Name:                 "Arc",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_Box = &thrift.StructSpec{
    Name:                 "Box",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_Exhaustive = &thrift.StructSpec{
    Name:                 "Exhaustive",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_Ord = &thrift.StructSpec{
    Name:                 "Ord",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_NewType = &thrift.StructSpec{
    Name:                 "NewType",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_Type = &thrift.StructSpec{
    Name:                 "Type",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "name",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "name": 0,
    },
}
    premadeStructSpec_Serde = &thrift.StructSpec{
    Name:                 "Serde",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "enabled",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "enabled": 0,
    },
}
    premadeStructSpec_Mod = &thrift.StructSpec{
    Name:                 "Mod",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "name",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "name": 0,
    },
}
    premadeStructSpec_Adapter = &thrift.StructSpec{
    Name:                 "Adapter",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "name",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "name": 0,
    },
}
    premadeStructSpec_Derive = &thrift.StructSpec{
    Name:                 "Derive",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "derives",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "derives": 0,
    },
}
    premadeStructSpec_ServiceExn = &thrift.StructSpec{
    Name:                 "ServiceExn",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "anyhow_to_application_exn",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "anyhow_to_application_exn": 0,
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
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "string", premadeCodecTypeSpec_string })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Name", premadeCodecTypeSpec_rust_Name })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Copy", premadeCodecTypeSpec_rust_Copy })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.RequestContext", premadeCodecTypeSpec_rust_RequestContext })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Arc", premadeCodecTypeSpec_rust_Arc })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Box", premadeCodecTypeSpec_rust_Box })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Exhaustive", premadeCodecTypeSpec_rust_Exhaustive })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Ord", premadeCodecTypeSpec_rust_Ord })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.NewType", premadeCodecTypeSpec_rust_NewType })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Type", premadeCodecTypeSpec_rust_Type })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "bool", premadeCodecTypeSpec_bool })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Serde", premadeCodecTypeSpec_rust_Serde })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Mod", premadeCodecTypeSpec_rust_Mod })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Adapter", premadeCodecTypeSpec_rust_Adapter })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.Derive", premadeCodecTypeSpec_rust_Derive })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "rust.ServiceExn", premadeCodecTypeSpec_rust_ServiceExn })

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
