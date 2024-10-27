// Autogenerated by Thrift for thrift/annotation/thrift.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package thrift

import (
    "maps"
    "sync"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = maps.Copy[map[int]int, map[int]int]
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
    premadeThriftType_thrift_RpcPriority *metadata.ThriftType = nil
    premadeThriftType_thrift_Experimental *metadata.ThriftType = nil
    premadeThriftType_i32 *metadata.ThriftType = nil
    premadeThriftType_list_i32 *metadata.ThriftType = nil
    premadeThriftType_map_i32_i32 *metadata.ThriftType = nil
    premadeThriftType_thrift_ReserveIds *metadata.ThriftType = nil
    premadeThriftType_bool *metadata.ThriftType = nil
    premadeThriftType_thrift_RequiresBackwardCompatibility *metadata.ThriftType = nil
    premadeThriftType_thrift_TerseWrite *metadata.ThriftType = nil
    premadeThriftType_thrift_Box *metadata.ThriftType = nil
    premadeThriftType_thrift_Mixin *metadata.ThriftType = nil
    premadeThriftType_thrift_SerializeInFieldIdOrder *metadata.ThriftType = nil
    premadeThriftType_thrift_BitmaskEnum *metadata.ThriftType = nil
    premadeThriftType_thrift_ExceptionMessage *metadata.ThriftType = nil
    premadeThriftType_thrift_InternBox *metadata.ThriftType = nil
    premadeThriftType_thrift_Serial *metadata.ThriftType = nil
    premadeThriftType_string *metadata.ThriftType = nil
    premadeThriftType_thrift_Uri *metadata.ThriftType = nil
    premadeThriftType_thrift_Priority *metadata.ThriftType = nil
    premadeThriftType_map_string_string *metadata.ThriftType = nil
    premadeThriftType_thrift_DeprecatedUnvalidatedAnnotations *metadata.ThriftType = nil
    premadeThriftType_thrift_AllowReservedIdentifier *metadata.ThriftType = nil
    premadeThriftType_thrift_AllowReservedFilename *metadata.ThriftType = nil
)

// Premade Thrift type initializer
var premadeThriftTypesInitOnce = sync.OnceFunc(func() {
    premadeThriftType_thrift_RpcPriority = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("thrift.RpcPriority"),
    )
    premadeThriftType_thrift_Experimental = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.Experimental"),
    )
    premadeThriftType_i32 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I32_TYPE.Ptr(),
    )
    premadeThriftType_list_i32 = metadata.NewThriftType().SetTList(
        metadata.NewThriftListType().
            SetValueType(premadeThriftType_i32),
    )
    premadeThriftType_map_i32_i32 = metadata.NewThriftType().SetTMap(
        metadata.NewThriftMapType().
            SetKeyType(premadeThriftType_i32).
            SetValueType(premadeThriftType_i32),
    )
    premadeThriftType_thrift_ReserveIds = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.ReserveIds"),
    )
    premadeThriftType_bool = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BOOL_TYPE.Ptr(),
    )
    premadeThriftType_thrift_RequiresBackwardCompatibility = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.RequiresBackwardCompatibility"),
    )
    premadeThriftType_thrift_TerseWrite = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.TerseWrite"),
    )
    premadeThriftType_thrift_Box = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.Box"),
    )
    premadeThriftType_thrift_Mixin = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.Mixin"),
    )
    premadeThriftType_thrift_SerializeInFieldIdOrder = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.SerializeInFieldIdOrder"),
    )
    premadeThriftType_thrift_BitmaskEnum = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.BitmaskEnum"),
    )
    premadeThriftType_thrift_ExceptionMessage = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.ExceptionMessage"),
    )
    premadeThriftType_thrift_InternBox = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.InternBox"),
    )
    premadeThriftType_thrift_Serial = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.Serial"),
    )
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
    )
    premadeThriftType_thrift_Uri = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.Uri"),
    )
    premadeThriftType_thrift_Priority = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.Priority"),
    )
    premadeThriftType_map_string_string = metadata.NewThriftType().SetTMap(
        metadata.NewThriftMapType().
            SetKeyType(premadeThriftType_string).
            SetValueType(premadeThriftType_string),
    )
    premadeThriftType_thrift_DeprecatedUnvalidatedAnnotations = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.DeprecatedUnvalidatedAnnotations"),
    )
    premadeThriftType_thrift_AllowReservedIdentifier = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.AllowReservedIdentifier"),
    )
    premadeThriftType_thrift_AllowReservedFilename = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("thrift.AllowReservedFilename"),
    )
})

var premadeThriftTypesMapOnce = sync.OnceValue(
    func() map[string]*metadata.ThriftType {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return map[string]*metadata.ThriftType{
            "thrift.RpcPriority": premadeThriftType_thrift_RpcPriority,
            "thrift.Experimental": premadeThriftType_thrift_Experimental,
            "i32": premadeThriftType_i32,
            "thrift.ReserveIds": premadeThriftType_thrift_ReserveIds,
            "bool": premadeThriftType_bool,
            "thrift.RequiresBackwardCompatibility": premadeThriftType_thrift_RequiresBackwardCompatibility,
            "thrift.TerseWrite": premadeThriftType_thrift_TerseWrite,
            "thrift.Box": premadeThriftType_thrift_Box,
            "thrift.Mixin": premadeThriftType_thrift_Mixin,
            "thrift.SerializeInFieldIdOrder": premadeThriftType_thrift_SerializeInFieldIdOrder,
            "thrift.BitmaskEnum": premadeThriftType_thrift_BitmaskEnum,
            "thrift.ExceptionMessage": premadeThriftType_thrift_ExceptionMessage,
            "thrift.InternBox": premadeThriftType_thrift_InternBox,
            "thrift.Serial": premadeThriftType_thrift_Serial,
            "string": premadeThriftType_string,
            "thrift.Uri": premadeThriftType_thrift_Uri,
            "thrift.Priority": premadeThriftType_thrift_Priority,
            "thrift.DeprecatedUnvalidatedAnnotations": premadeThriftType_thrift_DeprecatedUnvalidatedAnnotations,
            "thrift.AllowReservedIdentifier": premadeThriftType_thrift_AllowReservedIdentifier,
            "thrift.AllowReservedFilename": premadeThriftType_thrift_AllowReservedFilename,
        }
    },
)

var structMetadatasOnce = sync.OnceValue(
    func() []*metadata.ThriftStruct {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return []*metadata.ThriftStruct{
            metadata.NewThriftStruct().
    SetName("thrift.Experimental").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.ReserveIds").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("ids").
    SetIsOptional(false).
    SetType(premadeThriftType_list_i32),
            metadata.NewThriftField().
    SetId(2).
    SetName("id_ranges").
    SetIsOptional(false).
    SetType(premadeThriftType_map_i32_i32),
        },
    ),
            metadata.NewThriftStruct().
    SetName("thrift.RequiresBackwardCompatibility").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("field_name").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
        },
    ),
            metadata.NewThriftStruct().
    SetName("thrift.TerseWrite").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.Box").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.Mixin").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.SerializeInFieldIdOrder").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.BitmaskEnum").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.ExceptionMessage").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.InternBox").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.Serial").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.Uri").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("value").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
            metadata.NewThriftStruct().
    SetName("thrift.Priority").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("level").
    SetIsOptional(false).
    SetType(premadeThriftType_thrift_RpcPriority),
        },
    ),
            metadata.NewThriftStruct().
    SetName("thrift.DeprecatedUnvalidatedAnnotations").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("items").
    SetIsOptional(false).
    SetType(premadeThriftType_map_string_string),
        },
    ),
            metadata.NewThriftStruct().
    SetName("thrift.AllowReservedIdentifier").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("thrift.AllowReservedFilename").
    SetIsUnion(false),
        }
    },
)

var exceptionMetadatasOnce = sync.OnceValue(
    func() []*metadata.ThriftException {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return []*metadata.ThriftException{
        }
    },
)

var enumMetadatasOnce = sync.OnceValue(
    func() []*metadata.ThriftEnum {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return []*metadata.ThriftEnum{
            metadata.NewThriftEnum().
    SetName("thrift.RpcPriority").
    SetElements(
        map[int32]string{
            0: "HIGH_IMPORTANT",
            1: "HIGH",
            2: "IMPORTANT",
            3: "NORMAL",
            4: "BEST_EFFORT",
        },
    ),
        }
    },
)

var serviceMetadatasOnce = sync.OnceValue(
    func() []*metadata.ThriftService {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return []*metadata.ThriftService{
        }
    },
)

// GetMetadataThriftType (INTERNAL USE ONLY).
// Returns metadata ThriftType for a given full type name.
func GetMetadataThriftType(fullName string) *metadata.ThriftType {
    return premadeThriftTypesMapOnce()[fullName]
}

// GetThriftMetadata returns complete Thrift metadata for current and imported packages.
func GetThriftMetadata() *metadata.ThriftMetadata {
    allEnumsMap := make(map[string]*metadata.ThriftEnum)
    allStructsMap := make(map[string]*metadata.ThriftStruct)
    allExceptionsMap := make(map[string]*metadata.ThriftException)
    allServicesMap := make(map[string]*metadata.ThriftService)

    // Add enum metadatas from the current program...
    for _, enumMetadata := range enumMetadatasOnce() {
        allEnumsMap[enumMetadata.GetName()] = enumMetadata
    }
    // Add struct metadatas from the current program...
    for _, structMetadata := range structMetadatasOnce() {
        allStructsMap[structMetadata.GetName()] = structMetadata
    }
    // Add exception metadatas from the current program...
    for _, exceptionMetadata := range exceptionMetadatasOnce() {
        allExceptionsMap[exceptionMetadata.GetName()] = exceptionMetadata
    }
    // Add service metadatas from the current program...
    for _, serviceMetadata := range serviceMetadatasOnce() {
        allServicesMap[serviceMetadata.GetName()] = serviceMetadata
    }

    // Obtain Thrift metadatas from recursively included programs...
    var recursiveThriftMetadatas []*metadata.ThriftMetadata

    // ...now merge metadatas from recursively included programs.
    for _, thriftMetadata := range recursiveThriftMetadatas {
        maps.Copy(allEnumsMap, thriftMetadata.GetEnums())
        maps.Copy(allStructsMap, thriftMetadata.GetStructs())
        maps.Copy(allExceptionsMap, thriftMetadata.GetExceptions())
        maps.Copy(allServicesMap, thriftMetadata.GetServices())
    }

    return metadata.NewThriftMetadata().
        SetEnums(allEnumsMap).
        SetStructs(allStructsMap).
        SetExceptions(allExceptionsMap).
        SetServices(allServicesMap)
}

// GetThriftMetadataForService returns Thrift metadata for the given service.
func GetThriftMetadataForService(scopedServiceName string) *metadata.ThriftMetadata {
    thriftMetadata := GetThriftMetadata()

    allServicesMap := thriftMetadata.GetServices()
    relevantServicesMap := make(map[string]*metadata.ThriftService)

    serviceMetadata := allServicesMap[scopedServiceName]
    // Visit and record all recursive parents of the target service.
    for serviceMetadata != nil {
        relevantServicesMap[serviceMetadata.GetName()] = serviceMetadata
        if serviceMetadata.IsSetParent() {
            serviceMetadata = allServicesMap[serviceMetadata.GetParent()]
        } else {
            serviceMetadata = nil
        }
    }

    thriftMetadata.SetServices(relevantServicesMap)

    return thriftMetadata
}
