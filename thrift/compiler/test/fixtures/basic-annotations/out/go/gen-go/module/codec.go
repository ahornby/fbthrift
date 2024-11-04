// Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-annotations/src/module.thrift
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
    premadeCodecTypeSpec_module_MyEnum *thrift.TypeSpec = nil
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyStructNestedAnnotation *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyUnion *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i64 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_list_string_6884 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyStruct *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_SecretStruct *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyException *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_AwesomeStruct *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_FantasticStruct *thrift.TypeSpec = nil
    premadeCodecTypeSpec_void *thrift.TypeSpec = nil
    premadeCodecTypeSpec_bool *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_module_MyEnum = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_module_MyStructNestedAnnotation = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyStructNestedAnnotation() },
},

    }
    premadeCodecTypeSpec_module_MyUnion = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyUnion() },
},

    }
    premadeCodecTypeSpec_i64 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I64,
},

    }
    premadeCodecTypeSpec_list_string = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.STRING,
	ElementTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_module_list_string_6884 = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_list_string,
},

    }
    premadeCodecTypeSpec_module_MyStruct = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyStruct() },
},

    }
    premadeCodecTypeSpec_module_SecretStruct = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewSecretStruct() },
},

    }
    premadeCodecTypeSpec_module_MyException = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyException() },
},

    }
    premadeCodecTypeSpec_module_AwesomeStruct = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_module_MyStruct,
},

    }
    premadeCodecTypeSpec_module_FantasticStruct = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_module_MyStruct,
},

    }
    premadeCodecTypeSpec_void = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_VOID,
},

    }
    premadeCodecTypeSpec_bool = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_BOOL,
},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_MyStructNestedAnnotation *thrift.StructSpec = nil
    premadeStructSpec_MyUnion *thrift.StructSpec = nil
    premadeStructSpec_MyException *thrift.StructSpec = nil
    premadeStructSpec_MyStruct *thrift.StructSpec = nil
    premadeStructSpec_SecretStruct *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePing *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePing *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceGetRandomData *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceGetRandomData *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceHasDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceHasDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceGoGetDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceGoGetDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePutDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePutDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceLobDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceLobDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceGoDoNothing *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceGoDoNothing *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePrioParentPing *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePrioParentPing *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePrioParentPong *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePrioParentPong *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePrioChildPang *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePrioChildPang *thrift.StructSpec = nil
    premadeStructSpec_reqBadServiceBar *thrift.StructSpec = nil
    premadeStructSpec_respBadServiceBar *thrift.StructSpec = nil
    premadeStructSpec_reqFooBarBazServiceFooStructured *thrift.StructSpec = nil
    premadeStructSpec_respFooBarBazServiceFooStructured *thrift.StructSpec = nil
    premadeStructSpec_reqFooBarBazServiceBarNonStructured *thrift.StructSpec = nil
    premadeStructSpec_respFooBarBazServiceBarNonStructured *thrift.StructSpec = nil
    premadeStructSpec_reqFooBarBazServiceBaz *thrift.StructSpec = nil
    premadeStructSpec_respFooBarBazServiceBaz *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_MyStructNestedAnnotation = &thrift.StructSpec{
    Name:                 "MyStructNestedAnnotation",
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
    premadeStructSpec_MyUnion = &thrift.StructSpec{
    Name:                 "MyUnion",
    IsUnion:              true,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_MyException = &thrift.StructSpec{
    Name:                 "MyException",
    IsUnion:              false,
    IsException:          true,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_MyStruct = &thrift.StructSpec{
    Name:                 "MyStruct",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "abstract",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "major",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "annotation_with_quote",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "class_",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   5,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "annotation_with_trailing_comma",
            ReflectIndex:         4,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   6,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "empty_annotations",
            ReflectIndex:         5,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   7,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "my_enum",
            ReflectIndex:         6,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnum,
            MustBeSetToSerialize: false,
        },        {
            ID:                   8,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "cpp_type_annotation",
            ReflectIndex:         7,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_list_string_6884,
            MustBeSetToSerialize: false,
        },        {
            ID:                   9,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "my_union",
            ReflectIndex:         8,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyUnion,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
        5: 4,
        6: 5,
        7: 6,
        8: 7,
        9: 8,
    },
    FieldSpecNameToIndex: map[string]int{
        "abstract": 0,
        "major": 1,
        "annotation_with_quote": 2,
        "class_": 3,
        "annotation_with_trailing_comma": 4,
        "empty_annotations": 5,
        "my_enum": 6,
        "cpp_type_annotation": 7,
        "my_union": 8,
    },
}
    premadeStructSpec_SecretStruct = &thrift.StructSpec{
    Name:                 "SecretStruct",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "password",
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
        "id": 0,
        "password": 1,
    },
}
    premadeStructSpec_reqMyServicePing = &thrift.StructSpec{
    Name:                 "reqMyServicePing",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respMyServicePing = &thrift.StructSpec{
    Name:                 "respMyServicePing",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myExcept",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyException,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "myExcept": 0,
    },
}
    premadeStructSpec_reqMyServiceGetRandomData = &thrift.StructSpec{
    Name:                 "reqMyServiceGetRandomData",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respMyServiceGetRandomData = &thrift.StructSpec{
    Name:                 "respMyServiceGetRandomData",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
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
    premadeStructSpec_reqMyServiceHasDataById = &thrift.StructSpec{
    Name:                 "reqMyServiceHasDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "id": 0,
    },
}
    premadeStructSpec_respMyServiceHasDataById = &thrift.StructSpec{
    Name:                 "respMyServiceHasDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        0: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "success": 0,
    },
}
    premadeStructSpec_reqMyServiceGoGetDataById = &thrift.StructSpec{
    Name:                 "reqMyServiceGoGetDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "id": 0,
    },
}
    premadeStructSpec_respMyServiceGoGetDataById = &thrift.StructSpec{
    Name:                 "respMyServiceGoGetDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
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
    premadeStructSpec_reqMyServicePutDataById = &thrift.StructSpec{
    Name:                 "reqMyServicePutDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "data",
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
        "id": 0,
        "data": 1,
    },
}
    premadeStructSpec_respMyServicePutDataById = &thrift.StructSpec{
    Name:                 "respMyServicePutDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqMyServiceLobDataById = &thrift.StructSpec{
    Name:                 "reqMyServiceLobDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "data",
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
        "id": 0,
        "data": 1,
    },
}
    premadeStructSpec_respMyServiceLobDataById = &thrift.StructSpec{
    Name:                 "respMyServiceLobDataById",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqMyServiceGoDoNothing = &thrift.StructSpec{
    Name:                 "reqMyServiceGoDoNothing",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respMyServiceGoDoNothing = &thrift.StructSpec{
    Name:                 "respMyServiceGoDoNothing",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqMyServicePrioParentPing = &thrift.StructSpec{
    Name:                 "reqMyServicePrioParentPing",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respMyServicePrioParentPing = &thrift.StructSpec{
    Name:                 "respMyServicePrioParentPing",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqMyServicePrioParentPong = &thrift.StructSpec{
    Name:                 "reqMyServicePrioParentPong",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respMyServicePrioParentPong = &thrift.StructSpec{
    Name:                 "respMyServicePrioParentPong",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqMyServicePrioChildPang = &thrift.StructSpec{
    Name:                 "reqMyServicePrioChildPang",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respMyServicePrioChildPang = &thrift.StructSpec{
    Name:                 "respMyServicePrioChildPang",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqBadServiceBar = &thrift.StructSpec{
    Name:                 "reqBadServiceBar",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respBadServiceBar = &thrift.StructSpec{
    Name:                 "respBadServiceBar",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        0: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "success": 0,
    },
}
    premadeStructSpec_reqFooBarBazServiceFooStructured = &thrift.StructSpec{
    Name:                 "reqFooBarBazServiceFooStructured",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respFooBarBazServiceFooStructured = &thrift.StructSpec{
    Name:                 "respFooBarBazServiceFooStructured",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqFooBarBazServiceBarNonStructured = &thrift.StructSpec{
    Name:                 "reqFooBarBazServiceBarNonStructured",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respFooBarBazServiceBarNonStructured = &thrift.StructSpec{
    Name:                 "respFooBarBazServiceBarNonStructured",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_reqFooBarBazServiceBaz = &thrift.StructSpec{
    Name:                 "reqFooBarBazServiceBaz",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
    },
    FieldSpecIDToIndex:   map[int16]int{
    },
    FieldSpecNameToIndex: map[string]int{
    },
}
    premadeStructSpec_respFooBarBazServiceBaz = &thrift.StructSpec{
    Name:                 "respFooBarBazServiceBaz",
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
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyEnum", premadeCodecTypeSpec_module_MyEnum })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "string", premadeCodecTypeSpec_string })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyStructNestedAnnotation", premadeCodecTypeSpec_module_MyStructNestedAnnotation })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyUnion", premadeCodecTypeSpec_module_MyUnion })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "i64", premadeCodecTypeSpec_i64 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.list_string_6884", premadeCodecTypeSpec_module_list_string_6884 })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyStruct", premadeCodecTypeSpec_module_MyStruct })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.SecretStruct", premadeCodecTypeSpec_module_SecretStruct })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.MyException", premadeCodecTypeSpec_module_MyException })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.AwesomeStruct", premadeCodecTypeSpec_module_AwesomeStruct })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "module.FantasticStruct", premadeCodecTypeSpec_module_FantasticStruct })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "void", premadeCodecTypeSpec_void })
        codecSpecsWithFullName = append(codecSpecsWithFullName, codecSpecWithFullName{ "bool", premadeCodecTypeSpec_bool })
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
