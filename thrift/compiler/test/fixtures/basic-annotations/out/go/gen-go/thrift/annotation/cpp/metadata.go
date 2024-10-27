// Autogenerated by Thrift for thrift/annotation/cpp.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package cpp

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
    premadeThriftType_cpp_RefType *metadata.ThriftType = nil
    premadeThriftType_cpp_EnumUnderlyingType *metadata.ThriftType = nil
    premadeThriftType_string *metadata.ThriftType = nil
    premadeThriftType_cpp_Name *metadata.ThriftType = nil
    premadeThriftType_cpp_Type *metadata.ThriftType = nil
    premadeThriftType_cpp_Ref *metadata.ThriftType = nil
    premadeThriftType_bool *metadata.ThriftType = nil
    premadeThriftType_cpp_Lazy *metadata.ThriftType = nil
    premadeThriftType_cpp_DisableLazyChecksum *metadata.ThriftType = nil
    premadeThriftType_cpp_Adapter *metadata.ThriftType = nil
    premadeThriftType_cpp_PackIsset *metadata.ThriftType = nil
    premadeThriftType_cpp_MinimizePadding *metadata.ThriftType = nil
    premadeThriftType_cpp_ScopedEnumAsUnionType *metadata.ThriftType = nil
    premadeThriftType_cpp_FieldInterceptor *metadata.ThriftType = nil
    premadeThriftType_cpp_UseOpEncode *metadata.ThriftType = nil
    premadeThriftType_cpp_EnumType *metadata.ThriftType = nil
    premadeThriftType_cpp_Frozen2Exclude *metadata.ThriftType = nil
    premadeThriftType_cpp_Frozen2RequiresCompleteContainerParams *metadata.ThriftType = nil
    premadeThriftType_cpp_ProcessInEbThreadUnsafe *metadata.ThriftType = nil
    premadeThriftType_cpp_RuntimeAnnotation *metadata.ThriftType = nil
    premadeThriftType_cpp_UseCursorSerialization *metadata.ThriftType = nil
    premadeThriftType_cpp_GenerateDeprecatedHeaderClientMethods *metadata.ThriftType = nil
)

// Premade Thrift type initializer
var premadeThriftTypesInitOnce = sync.OnceFunc(func() {
    premadeThriftType_cpp_RefType = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("cpp.RefType"),
    )
    premadeThriftType_cpp_EnumUnderlyingType = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("cpp.EnumUnderlyingType"),
    )
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
    )
    premadeThriftType_cpp_Name = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Name"),
    )
    premadeThriftType_cpp_Type = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Type"),
    )
    premadeThriftType_cpp_Ref = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Ref"),
    )
    premadeThriftType_bool = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BOOL_TYPE.Ptr(),
    )
    premadeThriftType_cpp_Lazy = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Lazy"),
    )
    premadeThriftType_cpp_DisableLazyChecksum = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.DisableLazyChecksum"),
    )
    premadeThriftType_cpp_Adapter = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Adapter"),
    )
    premadeThriftType_cpp_PackIsset = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.PackIsset"),
    )
    premadeThriftType_cpp_MinimizePadding = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.MinimizePadding"),
    )
    premadeThriftType_cpp_ScopedEnumAsUnionType = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.ScopedEnumAsUnionType"),
    )
    premadeThriftType_cpp_FieldInterceptor = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.FieldInterceptor"),
    )
    premadeThriftType_cpp_UseOpEncode = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.UseOpEncode"),
    )
    premadeThriftType_cpp_EnumType = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.EnumType"),
    )
    premadeThriftType_cpp_Frozen2Exclude = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Frozen2Exclude"),
    )
    premadeThriftType_cpp_Frozen2RequiresCompleteContainerParams = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.Frozen2RequiresCompleteContainerParams"),
    )
    premadeThriftType_cpp_ProcessInEbThreadUnsafe = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.ProcessInEbThreadUnsafe"),
    )
    premadeThriftType_cpp_RuntimeAnnotation = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.RuntimeAnnotation"),
    )
    premadeThriftType_cpp_UseCursorSerialization = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.UseCursorSerialization"),
    )
    premadeThriftType_cpp_GenerateDeprecatedHeaderClientMethods = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("cpp.GenerateDeprecatedHeaderClientMethods"),
    )
})

var premadeThriftTypesMapOnce = sync.OnceValue(
    func() map[string]*metadata.ThriftType {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return map[string]*metadata.ThriftType{
            "cpp.RefType": premadeThriftType_cpp_RefType,
            "cpp.EnumUnderlyingType": premadeThriftType_cpp_EnumUnderlyingType,
            "string": premadeThriftType_string,
            "cpp.Name": premadeThriftType_cpp_Name,
            "cpp.Type": premadeThriftType_cpp_Type,
            "cpp.Ref": premadeThriftType_cpp_Ref,
            "bool": premadeThriftType_bool,
            "cpp.Lazy": premadeThriftType_cpp_Lazy,
            "cpp.DisableLazyChecksum": premadeThriftType_cpp_DisableLazyChecksum,
            "cpp.Adapter": premadeThriftType_cpp_Adapter,
            "cpp.PackIsset": premadeThriftType_cpp_PackIsset,
            "cpp.MinimizePadding": premadeThriftType_cpp_MinimizePadding,
            "cpp.ScopedEnumAsUnionType": premadeThriftType_cpp_ScopedEnumAsUnionType,
            "cpp.FieldInterceptor": premadeThriftType_cpp_FieldInterceptor,
            "cpp.UseOpEncode": premadeThriftType_cpp_UseOpEncode,
            "cpp.EnumType": premadeThriftType_cpp_EnumType,
            "cpp.Frozen2Exclude": premadeThriftType_cpp_Frozen2Exclude,
            "cpp.Frozen2RequiresCompleteContainerParams": premadeThriftType_cpp_Frozen2RequiresCompleteContainerParams,
            "cpp.ProcessInEbThreadUnsafe": premadeThriftType_cpp_ProcessInEbThreadUnsafe,
            "cpp.RuntimeAnnotation": premadeThriftType_cpp_RuntimeAnnotation,
            "cpp.UseCursorSerialization": premadeThriftType_cpp_UseCursorSerialization,
            "cpp.GenerateDeprecatedHeaderClientMethods": premadeThriftType_cpp_GenerateDeprecatedHeaderClientMethods,
        }
    },
)

var structMetadatasOnce = sync.OnceValue(
    func() []*metadata.ThriftStruct {
        // Relies on premade Thrift types initialization
        premadeThriftTypesInitOnce()
        return []*metadata.ThriftStruct{
            metadata.NewThriftStruct().
    SetName("cpp.Name").
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
    SetName("cpp.Type").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("name").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(2).
    SetName("template").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.Ref").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("type").
    SetIsOptional(false).
    SetType(premadeThriftType_cpp_RefType),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.Lazy").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("ref").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.DisableLazyChecksum").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.Adapter").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("name").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(2).
    SetName("adaptedType").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(3).
    SetName("underlyingName").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(4).
    SetName("extraNamespace").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(5).
    SetName("moveOnly").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.PackIsset").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("atomic").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.MinimizePadding").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.ScopedEnumAsUnionType").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.FieldInterceptor").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("name").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(2).
    SetName("noinline").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.UseOpEncode").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.EnumType").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("type").
    SetIsOptional(false).
    SetType(premadeThriftType_cpp_EnumUnderlyingType),
        },
    ),
            metadata.NewThriftStruct().
    SetName("cpp.Frozen2Exclude").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.Frozen2RequiresCompleteContainerParams").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.ProcessInEbThreadUnsafe").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.RuntimeAnnotation").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.UseCursorSerialization").
    SetIsUnion(false),
            metadata.NewThriftStruct().
    SetName("cpp.GenerateDeprecatedHeaderClientMethods").
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
    SetName("cpp.RefType").
    SetElements(
        map[int32]string{
            0: "Unique",
            1: "Shared",
            2: "SharedMutable",
        },
    ),
            metadata.NewThriftEnum().
    SetName("cpp.EnumUnderlyingType").
    SetElements(
        map[int32]string{
            0: "I8",
            1: "U8",
            2: "I16",
            3: "U16",
            4: "U32",
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
