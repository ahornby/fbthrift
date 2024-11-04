// Autogenerated by Thrift for thrift/compiler/test/fixtures/constants/src/module.thrift
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
    premadeCodecTypeSpec_module_EmptyEnum *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_City *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Company *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_double *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Internship *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Range *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_struct1 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_struct2 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_struct3 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_byte *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_struct4 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_union1 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_union2 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyCompany *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyStringIdentifier *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyIntIdentifier *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_string_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyMapIdentifier *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_module_EmptyEnum = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_City = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_Company = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_double = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_DOUBLE,
},

    }
    premadeCodecTypeSpec_module_Internship = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewInternship() },
},

    }
    premadeCodecTypeSpec_module_Range = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewRange() },
},

    }
    premadeCodecTypeSpec_module_struct1 = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewStruct1() },
},

    }
    premadeCodecTypeSpec_list_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.I32,
	ElementTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_module_struct2 = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewStruct2() },
},

    }
    premadeCodecTypeSpec_module_struct3 = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewStruct3() },
},

    }
    premadeCodecTypeSpec_byte = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_BYTE,
},

    }
    premadeCodecTypeSpec_module_struct4 = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewStruct4() },
},

    }
    premadeCodecTypeSpec_module_union1 = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewUnion1() },
},

    }
    premadeCodecTypeSpec_module_union2 = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewUnion2() },
},

    }
    premadeCodecTypeSpec_module_MyCompany = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_module_Company,
},

    }
    premadeCodecTypeSpec_module_MyStringIdentifier = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_module_MyIntIdentifier = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_map_string_string = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_string,
	ValueTypeSpec: premadeCodecTypeSpec_string,
    KeyWireType:   thrift.STRING,
	ValueWireType: thrift.STRING,
},

    }
    premadeCodecTypeSpec_module_MyMapIdentifier = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_map_string_string,
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Internship *thrift.StructSpec = nil
    premadeStructSpec_Range *thrift.StructSpec = nil
    premadeStructSpec_struct1 *thrift.StructSpec = nil
    premadeStructSpec_struct2 *thrift.StructSpec = nil
    premadeStructSpec_struct3 *thrift.StructSpec = nil
    premadeStructSpec_struct4 *thrift.StructSpec = nil
    premadeStructSpec_union1 *thrift.StructSpec = nil
    premadeStructSpec_union2 *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Internship = &thrift.StructSpec{
    Name:                 "Internship",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "weeks",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "title",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "employer",
            ReflectIndex:         2,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Company,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.DOUBLE),
            Name:                 "compensation",
            ReflectIndex:         3,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_double,
            MustBeSetToSerialize: true,
        },        {
            ID:                   5,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "school",
            ReflectIndex:         4,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
        5: 4,
    },
    FieldSpecNameToIndex: map[string]int{
        "weeks": 0,
        "title": 1,
        "employer": 2,
        "compensation": 3,
        "school": 4,
    },
}
    premadeStructSpec_Range = &thrift.StructSpec{
    Name:                 "Range",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "min",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "max",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
    },
    FieldSpecNameToIndex: map[string]int{
        "min": 0,
        "max": 1,
    },
}
    premadeStructSpec_struct1 = &thrift.StructSpec{
    Name:                 "struct1",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "a",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "b",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
    },
    FieldSpecNameToIndex: map[string]int{
        "a": 0,
        "b": 1,
    },
}
    premadeStructSpec_struct2 = &thrift.StructSpec{
    Name:                 "struct2",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "a",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "b",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "c",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_struct1,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "d",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
    },
    FieldSpecNameToIndex: map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
        "d": 3,
    },
}
    premadeStructSpec_struct3 = &thrift.StructSpec{
    Name:                 "struct3",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "a",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "b",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "c",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_struct2,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
    },
    FieldSpecNameToIndex: map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
    },
}
    premadeStructSpec_struct4 = &thrift.StructSpec{
    Name:                 "struct4",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "a",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.DOUBLE),
            Name:                 "b",
            ReflectIndex:         1,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_double,
            MustBeSetToSerialize: true,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.BYTE),
            Name:                 "c",
            ReflectIndex:         2,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_byte,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
    },
    FieldSpecNameToIndex: map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
    },
}
    premadeStructSpec_union1 = &thrift.StructSpec{
    Name:                 "union1",
    IsUnion:              true,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "i",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.DOUBLE),
            Name:                 "d",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_double,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
    },
    FieldSpecNameToIndex: map[string]int{
        "i": 0,
        "d": 1,
    },
}
    premadeStructSpec_union2 = &thrift.StructSpec{
    Name:                 "union2",
    IsUnion:              true,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "i",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.DOUBLE),
            Name:                 "d",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_double,
            MustBeSetToSerialize: true,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "s",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_struct1,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "u",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_union1,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
    },
    FieldSpecNameToIndex: map[string]int{
        "i": 0,
        "d": 1,
        "s": 2,
        "u": 3,
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
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.EmptyEnum", premadeCodecTypeSpec_module_EmptyEnum })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.City", premadeCodecTypeSpec_module_City })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.Company", premadeCodecTypeSpec_module_Company })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "i32", premadeCodecTypeSpec_i32 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "string", premadeCodecTypeSpec_string })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "double", premadeCodecTypeSpec_double })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.Internship", premadeCodecTypeSpec_module_Internship })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.Range", premadeCodecTypeSpec_module_Range })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.struct1", premadeCodecTypeSpec_module_struct1 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.struct2", premadeCodecTypeSpec_module_struct2 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.struct3", premadeCodecTypeSpec_module_struct3 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "byte", premadeCodecTypeSpec_byte })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.struct4", premadeCodecTypeSpec_module_struct4 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.union1", premadeCodecTypeSpec_module_union1 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.union2", premadeCodecTypeSpec_module_union2 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyCompany", premadeCodecTypeSpec_module_MyCompany })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyStringIdentifier", premadeCodecTypeSpec_module_MyStringIdentifier })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyIntIdentifier", premadeCodecTypeSpec_module_MyIntIdentifier })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyMapIdentifier", premadeCodecTypeSpec_module_MyMapIdentifier })

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
