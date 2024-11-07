// Autogenerated by Thrift for thrift/compiler/test/fixtures/exceptions/src/module.thrift
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
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Fiery *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Serious *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_ComplexFieldNames *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_CustomFieldNames *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_ExceptionWithPrimitiveField *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_ExceptionWithStructuredAnnotation *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Banal *thrift.TypeSpec = nil
    premadeCodecTypeSpec_void *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        FullName: "string",
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_module_Fiery = &thrift.TypeSpec{
        FullName: "module.Fiery",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Fiery",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewFiery() },
},

    }
    premadeCodecTypeSpec_module_Serious = &thrift.TypeSpec{
        FullName: "module.Serious",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Serious",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewSerious() },
},

    }
    premadeCodecTypeSpec_module_ComplexFieldNames = &thrift.TypeSpec{
        FullName: "module.ComplexFieldNames",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.ComplexFieldNames",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewComplexFieldNames() },
},

    }
    premadeCodecTypeSpec_module_CustomFieldNames = &thrift.TypeSpec{
        FullName: "module.CustomFieldNames",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.CustomFieldNames",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewCustomFieldNames() },
},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        FullName: "i32",
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_module_ExceptionWithPrimitiveField = &thrift.TypeSpec{
        FullName: "module.ExceptionWithPrimitiveField",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.ExceptionWithPrimitiveField",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewExceptionWithPrimitiveField() },
},

    }
    premadeCodecTypeSpec_module_ExceptionWithStructuredAnnotation = &thrift.TypeSpec{
        FullName: "module.ExceptionWithStructuredAnnotation",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.ExceptionWithStructuredAnnotation",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewExceptionWithStructuredAnnotation() },
},

    }
    premadeCodecTypeSpec_module_Banal = &thrift.TypeSpec{
        FullName: "module.Banal",
        CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "module.Banal",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewBanal() },
},

    }
    premadeCodecTypeSpec_void = &thrift.TypeSpec{
        FullName: "void",
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_VOID,
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Fiery *thrift.StructSpec = nil
    premadeStructSpec_Serious *thrift.StructSpec = nil
    premadeStructSpec_ComplexFieldNames *thrift.StructSpec = nil
    premadeStructSpec_CustomFieldNames *thrift.StructSpec = nil
    premadeStructSpec_ExceptionWithPrimitiveField *thrift.StructSpec = nil
    premadeStructSpec_ExceptionWithStructuredAnnotation *thrift.StructSpec = nil
    premadeStructSpec_Banal *thrift.StructSpec = nil
    premadeStructSpec_reqRaiserDoBland *thrift.StructSpec = nil
    premadeStructSpec_respRaiserDoBland *thrift.StructSpec = nil
    premadeStructSpec_reqRaiserDoRaise *thrift.StructSpec = nil
    premadeStructSpec_respRaiserDoRaise *thrift.StructSpec = nil
    premadeStructSpec_reqRaiserGet200 *thrift.StructSpec = nil
    premadeStructSpec_respRaiserGet200 *thrift.StructSpec = nil
    premadeStructSpec_reqRaiserGet500 *thrift.StructSpec = nil
    premadeStructSpec_respRaiserGet500 *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Fiery = &thrift.StructSpec{
    Name:                 "Fiery",
    ScopedName:           "module.Fiery",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "message",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "message": 0,
    },
}
    premadeStructSpec_Serious = &thrift.StructSpec{
    Name:                 "Serious",
    ScopedName:           "module.Serious",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "sonnet",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "sonnet": 0,
    },
}
    premadeStructSpec_ComplexFieldNames = &thrift.StructSpec{
    Name:                 "ComplexFieldNames",
    ScopedName:           "module.ComplexFieldNames",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "error_message",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.STRING,
            Name:                 "internal_error_message",
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
        "error_message": 0,
        "internal_error_message": 1,
    },
}
    premadeStructSpec_CustomFieldNames = &thrift.StructSpec{
    Name:                 "CustomFieldNames",
    ScopedName:           "module.CustomFieldNames",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "error_message",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.STRING,
            Name:                 "internal_error_message",
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
        "error_message": 0,
        "internal_error_message": 1,
    },
}
    premadeStructSpec_ExceptionWithPrimitiveField = &thrift.StructSpec{
    Name:                 "ExceptionWithPrimitiveField",
    ScopedName:           "module.ExceptionWithPrimitiveField",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "message",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.I32,
            Name:                 "error_code",
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
        "message": 0,
        "error_code": 1,
    },
}
    premadeStructSpec_ExceptionWithStructuredAnnotation = &thrift.StructSpec{
    Name:                 "ExceptionWithStructuredAnnotation",
    ScopedName:           "module.ExceptionWithStructuredAnnotation",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "message_field",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.I32,
            Name:                 "error_code",
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
        "message_field": 0,
        "error_code": 1,
    },
}
    premadeStructSpec_Banal = &thrift.StructSpec{
    Name:                 "Banal",
    ScopedName:           "module.Banal",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqRaiserDoBland = &thrift.StructSpec{
    Name:                 "reqRaiserDoBland",
    ScopedName:           "module.reqRaiserDoBland",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respRaiserDoBland = &thrift.StructSpec{
    Name:                 "respRaiserDoBland",
    ScopedName:           "module.respRaiserDoBland",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqRaiserDoRaise = &thrift.StructSpec{
    Name:                 "reqRaiserDoRaise",
    ScopedName:           "module.reqRaiserDoRaise",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respRaiserDoRaise = &thrift.StructSpec{
    Name:                 "respRaiserDoRaise",
    ScopedName:           "module.respRaiserDoRaise",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRUCT,
            Name:                 "b",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Banal,
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.STRUCT,
            Name:                 "f",
            ReflectIndex:         1,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Fiery,
            MustBeSetToSerialize: true,
        },        {
            ID:                   3,
            WireType:             thrift.STRUCT,
            Name:                 "s",
            ReflectIndex:         2,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Serious,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
    },
    FieldSpecNameToIndex: map[string]int{
        "b": 0,
        "f": 1,
        "s": 2,
    },
}
    premadeStructSpec_reqRaiserGet200 = &thrift.StructSpec{
    Name:                 "reqRaiserGet200",
    ScopedName:           "module.reqRaiserGet200",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respRaiserGet200 = &thrift.StructSpec{
    Name:                 "respRaiserGet200",
    ScopedName:           "module.respRaiserGet200",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.STRING,
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        0: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "success": 0,
    },
}
    premadeStructSpec_reqRaiserGet500 = &thrift.StructSpec{
    Name:                 "reqRaiserGet500",
    ScopedName:           "module.reqRaiserGet500",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respRaiserGet500 = &thrift.StructSpec{
    Name:                 "respRaiserGet500",
    ScopedName:           "module.respRaiserGet500",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.STRING,
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },        {
            ID:                   1,
            WireType:             thrift.STRUCT,
            Name:                 "f",
            ReflectIndex:         1,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Fiery,
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.STRUCT,
            Name:                 "b",
            ReflectIndex:         2,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Banal,
            MustBeSetToSerialize: true,
        },        {
            ID:                   3,
            WireType:             thrift.STRUCT,
            Name:                 "s",
            ReflectIndex:         3,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_Serious,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        0: 0,
        1: 1,
        2: 2,
        3: 3,
    },
    FieldSpecNameToIndex: map[string]int{
        "success": 0,
        "f": 1,
        "b": 2,
        "s": 3,
    },
}
})

var premadeCodecSpecsMapOnce = sync.OnceValue(
    func() map[string]*thrift.TypeSpec {
        // Relies on premade codec specs initialization
        premadeCodecSpecsInitOnce()

        fbthriftTypeSpecsMap := make(map[string]*thrift.TypeSpec)
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_string.FullName] = premadeCodecTypeSpec_string
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Fiery.FullName] = premadeCodecTypeSpec_module_Fiery
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Serious.FullName] = premadeCodecTypeSpec_module_Serious
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_ComplexFieldNames.FullName] = premadeCodecTypeSpec_module_ComplexFieldNames
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_CustomFieldNames.FullName] = premadeCodecTypeSpec_module_CustomFieldNames
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_i32.FullName] = premadeCodecTypeSpec_i32
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_ExceptionWithPrimitiveField.FullName] = premadeCodecTypeSpec_module_ExceptionWithPrimitiveField
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_ExceptionWithStructuredAnnotation.FullName] = premadeCodecTypeSpec_module_ExceptionWithStructuredAnnotation
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_module_Banal.FullName] = premadeCodecTypeSpec_module_Banal
        fbthriftTypeSpecsMap[premadeCodecTypeSpec_void.FullName] = premadeCodecTypeSpec_void
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
