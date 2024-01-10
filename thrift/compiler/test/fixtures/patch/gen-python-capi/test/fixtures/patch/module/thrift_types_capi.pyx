
#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from cpython.ref cimport PyObject
from libcpp.utility cimport move as __move
from folly.iobuf cimport (
    from_unique_ptr as __IOBuf_from_unique_ptr,
    IOBuf as __IOBuf,
)
from thrift.python.serializer import (
    deserialize as __deserialize,
    Protocol as __Protocol,
    serialize_iobuf as __serialize_iobuf,
)
import test.fixtures.patch.module.thrift_types as __thrift_types

cdef api int can_extract__test__fixtures__patch__module__MyData(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyData) else 0


cdef api object init__test__fixtures__patch__module__MyData(object data):
    return __thrift_types.MyData._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataWithCustomDefault(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataWithCustomDefault) else 0


cdef api object init__test__fixtures__patch__module__MyDataWithCustomDefault(object data):
    return __thrift_types.MyDataWithCustomDefault._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__InnerUnion(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.InnerUnion) else 0


cdef api object init__test__fixtures__patch__module__InnerUnion(object data):
    return __thrift_types.InnerUnion._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyUnion(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyUnion) else 0


cdef api object init__test__fixtures__patch__module__MyUnion(object data):
    return __thrift_types.MyUnion._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStruct(object data):
    return __thrift_types.MyStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__LateDefStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.LateDefStruct) else 0


cdef api object init__test__fixtures__patch__module__LateDefStruct(object data):
    return __thrift_types.LateDefStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__Recursive(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.Recursive) else 0


cdef api object init__test__fixtures__patch__module__Recursive(object data):
    return __thrift_types.Recursive._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__Bar(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.Bar) else 0


cdef api object init__test__fixtures__patch__module__Bar(object data):
    return __thrift_types.Bar._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__Loop(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.Loop) else 0


cdef api object init__test__fixtures__patch__module__Loop(object data):
    return __thrift_types.Loop._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFields(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFields) else 0


cdef api object init__test__fixtures__patch__module__RefFields(object data):
    return __thrift_types.RefFields._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyDataPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyDataPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyDataPatch(object data):
    return __thrift_types.MyDataPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyDataFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyDataFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyDataFieldPatch(object data):
    return __thrift_types.MyDataFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyDataEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyDataEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyDataEnsureStruct(object data):
    return __thrift_types.MyDataEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataWithCustomDefaultPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataWithCustomDefaultPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyDataWithCustomDefaultPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyDataWithCustomDefaultPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataWithCustomDefaultPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyDataWithCustomDefaultPatch(object data):
    return __thrift_types.MyDataWithCustomDefaultPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataWithCustomDefaultFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataWithCustomDefaultFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyDataWithCustomDefaultFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyDataWithCustomDefaultFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataWithCustomDefaultFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyDataWithCustomDefaultFieldPatch(object data):
    return __thrift_types.MyDataWithCustomDefaultFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyDataWithCustomDefaultEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyDataWithCustomDefaultEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyDataWithCustomDefaultEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyDataWithCustomDefaultEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyDataWithCustomDefaultEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyDataWithCustomDefaultEnsureStruct(object data):
    return __thrift_types.MyDataWithCustomDefaultEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__InnerUnionPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.InnerUnionPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__InnerUnionPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__InnerUnionPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.InnerUnionPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__InnerUnionPatch(object data):
    return __thrift_types.InnerUnionPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__InnerUnionFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.InnerUnionFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__InnerUnionFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__InnerUnionFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.InnerUnionFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__InnerUnionFieldPatch(object data):
    return __thrift_types.InnerUnionFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyUnionPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyUnionPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyUnionPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyUnionPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyUnionPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyUnionPatch(object data):
    return __thrift_types.MyUnionPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyUnionFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyUnionFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyUnionFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyUnionFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyUnionFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyUnionFieldPatch(object data):
    return __thrift_types.MyUnionFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructPatch(object data):
    return __thrift_types.MyStructPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField10Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField10Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField10Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField10Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField10Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField10Patch(object data):
    return __thrift_types.MyStructField10Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField23Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField23Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField23Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField23Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField23Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField23Patch(object data):
    return __thrift_types.MyStructField23Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField26Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField26Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField26Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField26Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField26Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField26Patch(object data):
    return __thrift_types.MyStructField26Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField27Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField27Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField27Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField27Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField27Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField27Patch(object data):
    return __thrift_types.MyStructField27Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField28Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField28Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField28Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField28Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField28Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField28Patch(object data):
    return __thrift_types.MyStructField28Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField29Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField29Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField29Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField29Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField29Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField29Patch(object data):
    return __thrift_types.MyStructField29Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField30Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField30Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField30Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField30Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField30Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField30Patch(object data):
    return __thrift_types.MyStructField30Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructField30Patch1(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructField30Patch1) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructField30Patch1(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructField30Patch1(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructField30Patch1,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructField30Patch1(object data):
    return __thrift_types.MyStructField30Patch1._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructFieldPatch(object data):
    return __thrift_types.MyStructFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyStructEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyStructEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__MyStructEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__MyStructEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.MyStructEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__MyStructEnsureStruct(object data):
    return __thrift_types.MyStructEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__LateDefStructPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.LateDefStructPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__LateDefStructPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__LateDefStructPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.LateDefStructPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__LateDefStructPatch(object data):
    return __thrift_types.LateDefStructPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__LateDefStructFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.LateDefStructFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__LateDefStructFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__LateDefStructFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.LateDefStructFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__LateDefStructFieldPatch(object data):
    return __thrift_types.LateDefStructFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__LateDefStructEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.LateDefStructEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__LateDefStructEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__LateDefStructEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.LateDefStructEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__LateDefStructEnsureStruct(object data):
    return __thrift_types.LateDefStructEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RecursivePatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RecursivePatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RecursivePatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RecursivePatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RecursivePatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RecursivePatch(object data):
    return __thrift_types.RecursivePatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RecursiveField1Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RecursiveField1Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RecursiveField1Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RecursiveField1Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RecursiveField1Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RecursiveField1Patch(object data):
    return __thrift_types.RecursiveField1Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RecursiveFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RecursiveFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RecursiveFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RecursiveFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RecursiveFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RecursiveFieldPatch(object data):
    return __thrift_types.RecursiveFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RecursiveEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RecursiveEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RecursiveEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RecursiveEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RecursiveEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RecursiveEnsureStruct(object data):
    return __thrift_types.RecursiveEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__BarPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.BarPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__BarPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__BarPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.BarPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__BarPatch(object data):
    return __thrift_types.BarPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__BarFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.BarFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__BarFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__BarFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.BarFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__BarFieldPatch(object data):
    return __thrift_types.BarFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__BarEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.BarEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__BarEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__BarEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.BarEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__BarEnsureStruct(object data):
    return __thrift_types.BarEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__LoopPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.LoopPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__LoopPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__LoopPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.LoopPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__LoopPatch(object data):
    return __thrift_types.LoopPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsPatch(object data):
    return __thrift_types.RefFieldsPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField1Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField1Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField1Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField1Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField1Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField1Patch(object data):
    return __thrift_types.RefFieldsField1Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField2Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField2Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField2Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField2Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField2Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField2Patch(object data):
    return __thrift_types.RefFieldsField2Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField3Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField3Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField3Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField3Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField3Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField3Patch(object data):
    return __thrift_types.RefFieldsField3Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField4Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField4Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField4Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField4Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField4Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField4Patch(object data):
    return __thrift_types.RefFieldsField4Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField5Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField5Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField5Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField5Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField5Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField5Patch(object data):
    return __thrift_types.RefFieldsField5Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField6Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField6Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField6Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField6Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField6Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField6Patch(object data):
    return __thrift_types.RefFieldsField6Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsField7Patch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsField7Patch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsField7Patch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsField7Patch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsField7Patch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsField7Patch(object data):
    return __thrift_types.RefFieldsField7Patch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsFieldPatch(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsFieldPatch) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsFieldPatch(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsFieldPatch(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsFieldPatch,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsFieldPatch(object data):
    return __thrift_types.RefFieldsFieldPatch._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__RefFieldsEnsureStruct(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.RefFieldsEnsureStruct) else 0

cdef api __cIOBuf* extract__test__fixtures__patch__module__RefFieldsEnsureStruct(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__fixtures__patch__module__RefFieldsEnsureStruct(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.RefFieldsEnsureStruct,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__fixtures__patch__module__RefFieldsEnsureStruct(object data):
    return __thrift_types.RefFieldsEnsureStruct._fbthrift_create(data)

cdef api int can_extract__test__fixtures__patch__module__MyEnum(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.MyEnum) else 0

cdef api object construct__test__fixtures__patch__module__MyEnum(int64_t __val):
    return __thrift_types.MyEnum(__val)

