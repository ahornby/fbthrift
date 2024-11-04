// Autogenerated by Thrift for thrift/compiler/test/fixtures/params/src/module.thrift
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
    premadeCodecTypeSpec_void *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_i32_list_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_set_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_i32_set_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_i32_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_map_i32_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_set_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_i32_map_i32_set_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_map_i32_map_i32_set_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_list_map_i32_map_i32_set_i32 *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_void = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_VOID,
},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_list_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.I32,
	ElementTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_map_i32_list_i32 = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_i32,
	ValueTypeSpec: premadeCodecTypeSpec_list_i32,
    KeyWireType:   thrift.I32,
	ValueWireType: thrift.LIST,
},

    }
    premadeCodecTypeSpec_set_i32 = &thrift.TypeSpec{
        CodecSetSpec: &thrift.CodecSetSpec{
    ElementWireType: thrift.I32,
	ElementTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_map_i32_set_i32 = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_i32,
	ValueTypeSpec: premadeCodecTypeSpec_set_i32,
    KeyWireType:   thrift.I32,
	ValueWireType: thrift.SET,
},

    }
    premadeCodecTypeSpec_map_i32_i32 = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_i32,
	ValueTypeSpec: premadeCodecTypeSpec_i32,
    KeyWireType:   thrift.I32,
	ValueWireType: thrift.I32,
},

    }
    premadeCodecTypeSpec_list_map_i32_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.MAP,
	ElementTypeSpec: premadeCodecTypeSpec_map_i32_i32,
},

    }
    premadeCodecTypeSpec_list_set_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.SET,
	ElementTypeSpec: premadeCodecTypeSpec_set_i32,
},

    }
    premadeCodecTypeSpec_map_i32_map_i32_set_i32 = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_i32,
	ValueTypeSpec: premadeCodecTypeSpec_map_i32_set_i32,
    KeyWireType:   thrift.I32,
	ValueWireType: thrift.MAP,
},

    }
    premadeCodecTypeSpec_list_map_i32_map_i32_set_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.MAP,
	ElementTypeSpec: premadeCodecTypeSpec_map_i32_map_i32_set_i32,
},

    }
    premadeCodecTypeSpec_list_list_map_i32_map_i32_set_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.LIST,
	ElementTypeSpec: premadeCodecTypeSpec_list_map_i32_map_i32_set_i32,
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_reqNestedContainersMapList *thrift.StructSpec = nil
    premadeStructSpec_respNestedContainersMapList *thrift.StructSpec = nil
    premadeStructSpec_reqNestedContainersMapSet *thrift.StructSpec = nil
    premadeStructSpec_respNestedContainersMapSet *thrift.StructSpec = nil
    premadeStructSpec_reqNestedContainersListMap *thrift.StructSpec = nil
    premadeStructSpec_respNestedContainersListMap *thrift.StructSpec = nil
    premadeStructSpec_reqNestedContainersListSet *thrift.StructSpec = nil
    premadeStructSpec_respNestedContainersListSet *thrift.StructSpec = nil
    premadeStructSpec_reqNestedContainersTurtles *thrift.StructSpec = nil
    premadeStructSpec_respNestedContainersTurtles *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_reqNestedContainersMapList = &thrift.StructSpec{
    Name:                 "reqNestedContainersMapList",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.MAP),
            Name:                 "foo",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_i32_list_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "foo": 0,
    },
}
    premadeStructSpec_respNestedContainersMapList = &thrift.StructSpec{
    Name:                 "respNestedContainersMapList",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqNestedContainersMapSet = &thrift.StructSpec{
    Name:                 "reqNestedContainersMapSet",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.MAP),
            Name:                 "foo",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_i32_set_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "foo": 0,
    },
}
    premadeStructSpec_respNestedContainersMapSet = &thrift.StructSpec{
    Name:                 "respNestedContainersMapSet",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqNestedContainersListMap = &thrift.StructSpec{
    Name:                 "reqNestedContainersListMap",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "foo",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_map_i32_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "foo": 0,
    },
}
    premadeStructSpec_respNestedContainersListMap = &thrift.StructSpec{
    Name:                 "respNestedContainersListMap",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqNestedContainersListSet = &thrift.StructSpec{
    Name:                 "reqNestedContainersListSet",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "foo",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_set_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "foo": 0,
    },
}
    premadeStructSpec_respNestedContainersListSet = &thrift.StructSpec{
    Name:                 "respNestedContainersListSet",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqNestedContainersTurtles = &thrift.StructSpec{
    Name:                 "reqNestedContainersTurtles",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "foo",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_list_map_i32_map_i32_set_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "foo": 0,
    },
}
    premadeStructSpec_respNestedContainersTurtles = &thrift.StructSpec{
    Name:                 "respNestedContainersTurtles",
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
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "void", premadeCodecTypeSpec_void })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "i32", premadeCodecTypeSpec_i32 })

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
