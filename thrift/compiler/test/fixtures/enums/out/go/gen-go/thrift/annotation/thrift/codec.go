// Autogenerated by Thrift for thrift/annotation/thrift.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package thrift


import (
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.VOID

// Premade codec specs
var (
    premadeCodecTypeSpec_thrift_RpcPriority = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.RpcPriority",
            CodecEnumSpec: &thrift.CodecEnumSpec{
    ScopedName: "thrift.RpcPriority",
},

        }
    }()
    premadeCodecTypeSpec_thrift_Experimental = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.Experimental",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.Experimental",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewExperimental() },
},

        }
    }()
    premadeCodecTypeSpec_i32 = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "i32",
            CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

        }
    }()
    premadeCodecTypeSpec_list_i32 = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "list<i32>",
            CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.I32,
	ElementTypeSpec: premadeCodecTypeSpec_i32,
},

        }
    }()
    premadeCodecTypeSpec_map_i32_i32 = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "map<i32, i32>",
            CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_i32,
	ValueTypeSpec: premadeCodecTypeSpec_i32,
    KeyWireType:   thrift.I32,
	ValueWireType: thrift.I32,
},

        }
    }()
    premadeCodecTypeSpec_thrift_ReserveIds = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.ReserveIds",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.ReserveIds",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewReserveIds() },
},

        }
    }()
    premadeCodecTypeSpec_bool = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "bool",
            CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_BOOL,
},

        }
    }()
    premadeCodecTypeSpec_thrift_RequiresBackwardCompatibility = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.RequiresBackwardCompatibility",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.RequiresBackwardCompatibility",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewRequiresBackwardCompatibility() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_TerseWrite = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.TerseWrite",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.TerseWrite",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewTerseWrite() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_Box = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.Box",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.Box",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewBox() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_Mixin = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.Mixin",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.Mixin",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewMixin() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_SerializeInFieldIdOrder = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.SerializeInFieldIdOrder",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.SerializeInFieldIdOrder",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewSerializeInFieldIdOrder() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_BitmaskEnum = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.BitmaskEnum",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.BitmaskEnum",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewBitmaskEnum() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_ExceptionMessage = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.ExceptionMessage",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.ExceptionMessage",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewExceptionMessage() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_InternBox = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.InternBox",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.InternBox",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewInternBox() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_Serial = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.Serial",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.Serial",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewSerial() },
},

        }
    }()
    premadeCodecTypeSpec_string = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "string",
            CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

        }
    }()
    premadeCodecTypeSpec_thrift_Uri = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.Uri",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.Uri",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewUri() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_Priority = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.Priority",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.Priority",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewPriority() },
},

        }
    }()
    premadeCodecTypeSpec_map_string_string = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "map<string, string>",
            CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_string,
	ValueTypeSpec: premadeCodecTypeSpec_string,
    KeyWireType:   thrift.STRING,
	ValueWireType: thrift.STRING,
},

        }
    }()
    premadeCodecTypeSpec_thrift_DeprecatedUnvalidatedAnnotations = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.DeprecatedUnvalidatedAnnotations",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.DeprecatedUnvalidatedAnnotations",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewDeprecatedUnvalidatedAnnotations() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_AllowReservedIdentifier = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.AllowReservedIdentifier",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.AllowReservedIdentifier",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewAllowReservedIdentifier() },
},

        }
    }()
    premadeCodecTypeSpec_thrift_AllowReservedFilename = func() *thrift.TypeSpec {
        return &thrift.TypeSpec{
            FullName: "thrift.AllowReservedFilename",
            CodecStructSpec: &thrift.CodecStructSpec{
    ScopedName: "thrift.AllowReservedFilename",
    IsUnion:    false,
    NewFunc:    func() thrift.Struct { return NewAllowReservedFilename() },
},

        }
    }()
)

// Premade struct specs
var (
    premadeStructSpec_Experimental = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Experimental",
    ScopedName:           "thrift.Experimental",
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
    premadeStructSpec_ReserveIds = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "ReserveIds",
    ScopedName:           "thrift.ReserveIds",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.LIST,
            Name:                 "ids",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.MAP,
            Name:                 "id_ranges",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_i32_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
        2: 1,
    },
    FieldSpecNameToIndex: map[string]int{
        "ids": 0,
        "id_ranges": 1,
    },
}
    }()
    premadeStructSpec_RequiresBackwardCompatibility = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "RequiresBackwardCompatibility",
    ScopedName:           "thrift.RequiresBackwardCompatibility",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.BOOL,
            Name:                 "field_name",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "field_name": 0,
    },
}
    }()
    premadeStructSpec_TerseWrite = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "TerseWrite",
    ScopedName:           "thrift.TerseWrite",
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
    premadeStructSpec_Box = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Box",
    ScopedName:           "thrift.Box",
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
    premadeStructSpec_Mixin = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Mixin",
    ScopedName:           "thrift.Mixin",
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
    premadeStructSpec_SerializeInFieldIdOrder = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "SerializeInFieldIdOrder",
    ScopedName:           "thrift.SerializeInFieldIdOrder",
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
    premadeStructSpec_BitmaskEnum = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "BitmaskEnum",
    ScopedName:           "thrift.BitmaskEnum",
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
    premadeStructSpec_ExceptionMessage = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "ExceptionMessage",
    ScopedName:           "thrift.ExceptionMessage",
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
    premadeStructSpec_InternBox = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "InternBox",
    ScopedName:           "thrift.InternBox",
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
    premadeStructSpec_Serial = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Serial",
    ScopedName:           "thrift.Serial",
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
    premadeStructSpec_Uri = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Uri",
    ScopedName:           "thrift.Uri",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.STRING,
            Name:                 "value",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "value": 0,
    },
}
    }()
    premadeStructSpec_Priority = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "Priority",
    ScopedName:           "thrift.Priority",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.I32,
            Name:                 "level",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_thrift_RpcPriority,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "level": 0,
    },
}
    }()
    premadeStructSpec_DeprecatedUnvalidatedAnnotations = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "DeprecatedUnvalidatedAnnotations",
    ScopedName:           "thrift.DeprecatedUnvalidatedAnnotations",
    IsUnion:              false,
    IsException:          false,
    FieldSpecs:           []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.MAP,
            Name:                 "items",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_string_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex:   map[int16]int{
        1: 0,
    },
    FieldSpecNameToIndex: map[string]int{
        "items": 0,
    },
}
    }()
    premadeStructSpec_AllowReservedIdentifier = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "AllowReservedIdentifier",
    ScopedName:           "thrift.AllowReservedIdentifier",
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
    premadeStructSpec_AllowReservedFilename = func() *thrift.StructSpec {
        return &thrift.StructSpec{
    Name:                 "AllowReservedFilename",
    ScopedName:           "thrift.AllowReservedFilename",
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
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Experimental)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_ReserveIds)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_RequiresBackwardCompatibility)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_TerseWrite)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Box)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Mixin)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_SerializeInFieldIdOrder)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_BitmaskEnum)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_ExceptionMessage)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_InternBox)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Serial)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Uri)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_Priority)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_DeprecatedUnvalidatedAnnotations)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_AllowReservedIdentifier)
    fbthriftResults = append(fbthriftResults, premadeStructSpec_AllowReservedFilename)
    return fbthriftResults
}()

var premadeCodecSpecsMap = func() map[string]*thrift.TypeSpec {
    fbthriftTypeSpecsMap := make(map[string]*thrift.TypeSpec)
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_RpcPriority.FullName] = premadeCodecTypeSpec_thrift_RpcPriority
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_Experimental.FullName] = premadeCodecTypeSpec_thrift_Experimental
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_i32.FullName] = premadeCodecTypeSpec_i32
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_ReserveIds.FullName] = premadeCodecTypeSpec_thrift_ReserveIds
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_bool.FullName] = premadeCodecTypeSpec_bool
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_RequiresBackwardCompatibility.FullName] = premadeCodecTypeSpec_thrift_RequiresBackwardCompatibility
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_TerseWrite.FullName] = premadeCodecTypeSpec_thrift_TerseWrite
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_Box.FullName] = premadeCodecTypeSpec_thrift_Box
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_Mixin.FullName] = premadeCodecTypeSpec_thrift_Mixin
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_SerializeInFieldIdOrder.FullName] = premadeCodecTypeSpec_thrift_SerializeInFieldIdOrder
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_BitmaskEnum.FullName] = premadeCodecTypeSpec_thrift_BitmaskEnum
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_ExceptionMessage.FullName] = premadeCodecTypeSpec_thrift_ExceptionMessage
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_InternBox.FullName] = premadeCodecTypeSpec_thrift_InternBox
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_Serial.FullName] = premadeCodecTypeSpec_thrift_Serial
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_string.FullName] = premadeCodecTypeSpec_string
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_Uri.FullName] = premadeCodecTypeSpec_thrift_Uri
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_Priority.FullName] = premadeCodecTypeSpec_thrift_Priority
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_DeprecatedUnvalidatedAnnotations.FullName] = premadeCodecTypeSpec_thrift_DeprecatedUnvalidatedAnnotations
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_AllowReservedIdentifier.FullName] = premadeCodecTypeSpec_thrift_AllowReservedIdentifier
    fbthriftTypeSpecsMap[premadeCodecTypeSpec_thrift_AllowReservedFilename.FullName] = premadeCodecTypeSpec_thrift_AllowReservedFilename
    return fbthriftTypeSpecsMap
}()

// GetMetadataThriftType (INTERNAL USE ONLY).
// Returns metadata TypeSpec for a given full type name.
func GetCodecTypeSpec(fullName string) *thrift.TypeSpec {
    return premadeCodecSpecsMap[fullName]
}
