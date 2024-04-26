#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cpython.object cimport PyTypeObject, Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, Py_GE
from libcpp.memory cimport shared_ptr, make_shared, unique_ptr
from libcpp.optional cimport optional as __optional
from libcpp.string cimport string
from libcpp cimport bool as cbool
from libcpp.iterator cimport inserter as cinserter
from libcpp.utility cimport move as cmove
from cpython cimport bool as pbool
from cython.operator cimport dereference as deref, preincrement as inc, address as ptr_address
import thrift.py3.types
from thrift.py3.types import _IsSet as _fbthrift_IsSet
from thrift.py3.types cimport make_unique
cimport thrift.py3.types
cimport thrift.py3.exceptions
from thrift.py3.std_libcpp cimport sv_to_str as __sv_to_str, string_view as __cstring_view
from thrift.py3.types cimport (
    cSetOp as __cSetOp,
    richcmp as __richcmp,
    set_op as __set_op,
    setcmp as __setcmp,
    list_index as __list_index,
    list_count as __list_count,
    list_slice as __list_slice,
    list_getitem as __list_getitem,
    set_iter as __set_iter,
    map_iter as __map_iter,
    map_contains as __map_contains,
    map_getitem as __map_getitem,
    reference_shared_ptr as __reference_shared_ptr,
    get_field_name_by_index as __get_field_name_by_index,
    reset_field as __reset_field,
    translate_cpp_enum_to_python,
    SetMetaClass as __SetMetaClass,
    const_pointer_cast,
    constant_shared_ptr,
    NOTSET as __NOTSET,
    EnumData as __EnumData,
    EnumFlagsData as __EnumFlagsData,
    UnionTypeEnumData as __UnionTypeEnumData,
    createEnumDataForUnionType as __createEnumDataForUnionType,
)
cimport thrift.py3.serializer as serializer
from thrift.python.protocol cimport Protocol as __Protocol
import folly.iobuf as _fbthrift_iobuf
from folly.optional cimport cOptional
from folly.memory cimport to_shared_ptr as __to_shared_ptr
from folly.range cimport Range as __cRange

import sys
from collections.abc import Sequence, Set, Mapping, Iterable
import weakref as __weakref
import builtins as _builtins


cdef __EnumData __Void_enum_data  = __EnumData._fbthrift_create(thrift.py3.types.createEnumData[cVoid](), Void)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __VoidMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __Void_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __Void_enum_data.get_all_names()

    def __len__(cls):
        return __Void_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __Void_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class Void(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __Void_enum_data.get_by_name(name)


    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        EnumMetadata[cVoid].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "standard.Void"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "apache.thrift.type.standard.thrift_types"
        )
        return python_types.Void(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value


__SetMetaClass(<PyTypeObject*> Void, <PyTypeObject*> __VoidMeta)


cdef __EnumData __StandardProtocol_enum_data  = __EnumData._fbthrift_create(thrift.py3.types.createEnumData[cStandardProtocol](), StandardProtocol)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __StandardProtocolMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __StandardProtocol_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __StandardProtocol_enum_data.get_all_names()

    def __len__(cls):
        return __StandardProtocol_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __StandardProtocol_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class StandardProtocol(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __StandardProtocol_enum_data.get_by_name(name)


    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        EnumMetadata[cStandardProtocol].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "standard.StandardProtocol"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "apache.thrift.type.standard.thrift_types"
        )
        return python_types.StandardProtocol(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value


__SetMetaClass(<PyTypeObject*> StandardProtocol, <PyTypeObject*> __StandardProtocolMeta)



cdef __UnionTypeEnumData __TypeUri_union_type_enum_data  = __UnionTypeEnumData._fbthrift_create(
    __createEnumDataForUnionType[cTypeUri](),
    __TypeUriType,
)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __TypeUri_Union_TypeMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __TypeUri_union_type_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __TypeUri_union_type_enum_data.get_all_names()

    def __len__(cls):
        return __TypeUri_union_type_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __TypeUri_union_type_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class __TypeUriType(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __TypeUri_union_type_enum_data.get_by_name(name)


__SetMetaClass(<PyTypeObject*> __TypeUriType, <PyTypeObject*> __TypeUri_Union_TypeMeta)


cdef __UnionTypeEnumData __TypeName_union_type_enum_data  = __UnionTypeEnumData._fbthrift_create(
    __createEnumDataForUnionType[cTypeName](),
    __TypeNameType,
)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __TypeName_Union_TypeMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __TypeName_union_type_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __TypeName_union_type_enum_data.get_all_names()

    def __len__(cls):
        return __TypeName_union_type_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __TypeName_union_type_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class __TypeNameType(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __TypeName_union_type_enum_data.get_by_name(name)


__SetMetaClass(<PyTypeObject*> __TypeNameType, <PyTypeObject*> __TypeName_Union_TypeMeta)




@__cython.auto_pickle(False)
cdef class TypeUri(thrift.py3.types.Union):
    Type = __TypeUriType

    def __init__(
        self, *,
        str uri=None,
        bytes typeHashPrefixSha2_256=None,
        str scopedName=None
    ):
        self._cpp_obj = __to_shared_ptr(cmove(TypeUri._make_instance(
          NULL,
          uri,
          typeHashPrefixSha2_256,
          scopedName,
        )))
        self._load_cache()

    @staticmethod
    def fromValue(value):
        if value is None:
            return TypeUri()
        if isinstance(value, str):
            return TypeUri(uri=value)
        if isinstance(value, bytes):
            return TypeUri(typeHashPrefixSha2_256=value)
        if isinstance(value, str):
            return TypeUri(scopedName=value)
        raise ValueError(f"Unable to derive correct union field for value: {value}")

    @staticmethod
    cdef unique_ptr[cTypeUri] _make_instance(
        cTypeUri* base_instance,
        str uri,
        bytes typeHashPrefixSha2_256,
        str scopedName
    ) except *:
        cdef unique_ptr[cTypeUri] c_inst = make_unique[cTypeUri]()
        cdef bint any_set = False
        if uri is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_uri(uri.encode('UTF-8'))
            any_set = True
        if typeHashPrefixSha2_256 is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_typeHashPrefixSha2_256(typeHashPrefixSha2_256)
            any_set = True
        if scopedName is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_scopedName(scopedName.encode('UTF-8'))
            any_set = True
        # in C++ you don't have to call move(), but this doesn't translate
        # into a C++ return statement, so you do here
        return cmove(c_inst)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cTypeUri] cpp_obj):
        __fbthrift_inst = <TypeUri>TypeUri.__new__(TypeUri)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        __fbthrift_inst._load_cache()
        return __fbthrift_inst

    @property
    def uri(self):
        if self.type.value != 1:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not uri')
        return self.value

    @property
    def typeHashPrefixSha2_256(self):
        if self.type.value != 2:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not typeHashPrefixSha2_256')
        return self.value

    @property
    def scopedName(self):
        if self.type.value != 3:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not scopedName')
        return self.value


    def __hash__(TypeUri self):
        return  super().__hash__()

    cdef _load_cache(TypeUri self):
        self.type = TypeUri.Type(<int>(deref(self._cpp_obj).getType()))
        cdef int type = self.type.value
        if type == 0:    # Empty
            self.value = None
        elif type == 1:
            self.value = bytes(deref(self._cpp_obj).get_uri()).decode('UTF-8')
        elif type == 2:
            self.value = deref(self._cpp_obj).get_typeHashPrefixSha2_256()
        elif type == 3:
            self.value = bytes(deref(self._cpp_obj).get_scopedName()).decode('UTF-8')

    def __copy__(TypeUri self):
        cdef shared_ptr[cTypeUri] cpp_obj = make_shared[cTypeUri](
            deref(self._cpp_obj)
        )
        return TypeUri._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cTypeUri](
            self._cpp_obj,
            (<TypeUri>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        import importlib
        types_reflection = importlib.import_module(
            "apache.thrift.type.standard.types_reflection"
        )
        return types_reflection.get_reflection__TypeUri()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cTypeUri].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "standard.TypeUri"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cTypeUri](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 3

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(TypeUri self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cTypeUri](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(TypeUri self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cTypeUri]()
        with nogil:
            needed = serializer.cdeserialize[cTypeUri](buf, self._cpp_obj.get(), proto)
        # force a cache reload since the underlying data's changed
        self._load_cache()
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "apache.thrift.type.standard.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.TypeUri, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("thrift.lib.thrift.standard.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.TypeUri, self)


@__cython.auto_pickle(False)
cdef class TypeName(thrift.py3.types.Union):
    Type = __TypeNameType

    def __init__(
        self, *,
        Void boolType=None,
        Void byteType=None,
        Void i16Type=None,
        Void i32Type=None,
        Void i64Type=None,
        Void floatType=None,
        Void doubleType=None,
        Void stringType=None,
        Void binaryType=None,
        TypeUri enumType=None,
        TypeUri typedefType=None,
        TypeUri structType=None,
        TypeUri unionType=None,
        TypeUri exceptionType=None,
        Void listType=None,
        Void setType=None,
        Void mapType=None
    ):
        self._cpp_obj = __to_shared_ptr(cmove(TypeName._make_instance(
          NULL,
          boolType,
          byteType,
          i16Type,
          i32Type,
          i64Type,
          floatType,
          doubleType,
          stringType,
          binaryType,
          enumType,
          typedefType,
          structType,
          unionType,
          exceptionType,
          listType,
          setType,
          mapType,
        )))
        self._load_cache()

    @staticmethod
    def fromValue(value):
        if value is None:
            return TypeName()
        if isinstance(value, Void):
            return TypeName(boolType=value)
        if isinstance(value, Void):
            return TypeName(byteType=value)
        if isinstance(value, Void):
            return TypeName(i16Type=value)
        if isinstance(value, Void):
            return TypeName(i32Type=value)
        if isinstance(value, Void):
            return TypeName(i64Type=value)
        if isinstance(value, Void):
            return TypeName(floatType=value)
        if isinstance(value, Void):
            return TypeName(doubleType=value)
        if isinstance(value, Void):
            return TypeName(stringType=value)
        if isinstance(value, Void):
            return TypeName(binaryType=value)
        if isinstance(value, TypeUri):
            return TypeName(enumType=value)
        if isinstance(value, TypeUri):
            return TypeName(typedefType=value)
        if isinstance(value, TypeUri):
            return TypeName(structType=value)
        if isinstance(value, TypeUri):
            return TypeName(unionType=value)
        if isinstance(value, TypeUri):
            return TypeName(exceptionType=value)
        if isinstance(value, Void):
            return TypeName(listType=value)
        if isinstance(value, Void):
            return TypeName(setType=value)
        if isinstance(value, Void):
            return TypeName(mapType=value)
        raise ValueError(f"Unable to derive correct union field for value: {value}")

    @staticmethod
    cdef unique_ptr[cTypeName] _make_instance(
        cTypeName* base_instance,
        Void boolType,
        Void byteType,
        Void i16Type,
        Void i32Type,
        Void i64Type,
        Void floatType,
        Void doubleType,
        Void stringType,
        Void binaryType,
        TypeUri enumType,
        TypeUri typedefType,
        TypeUri structType,
        TypeUri unionType,
        TypeUri exceptionType,
        Void listType,
        Void setType,
        Void mapType
    ) except *:
        cdef unique_ptr[cTypeName] c_inst = make_unique[cTypeName]()
        cdef bint any_set = False
        if boolType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_boolType(<cVoid><int>boolType)
            any_set = True
        if byteType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_byteType(<cVoid><int>byteType)
            any_set = True
        if i16Type is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_i16Type(<cVoid><int>i16Type)
            any_set = True
        if i32Type is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_i32Type(<cVoid><int>i32Type)
            any_set = True
        if i64Type is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_i64Type(<cVoid><int>i64Type)
            any_set = True
        if floatType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_floatType(<cVoid><int>floatType)
            any_set = True
        if doubleType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_doubleType(<cVoid><int>doubleType)
            any_set = True
        if stringType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_stringType(<cVoid><int>stringType)
            any_set = True
        if binaryType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_binaryType(<cVoid><int>binaryType)
            any_set = True
        if enumType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_enumType(deref((<TypeUri?> enumType)._cpp_obj))
            any_set = True
        if typedefType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_typedefType(deref((<TypeUri?> typedefType)._cpp_obj))
            any_set = True
        if structType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_structType(deref((<TypeUri?> structType)._cpp_obj))
            any_set = True
        if unionType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_unionType(deref((<TypeUri?> unionType)._cpp_obj))
            any_set = True
        if exceptionType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_exceptionType(deref((<TypeUri?> exceptionType)._cpp_obj))
            any_set = True
        if listType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_listType(<cVoid><int>listType)
            any_set = True
        if setType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_setType(<cVoid><int>setType)
            any_set = True
        if mapType is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).set_mapType(<cVoid><int>mapType)
            any_set = True
        # in C++ you don't have to call move(), but this doesn't translate
        # into a C++ return statement, so you do here
        return cmove(c_inst)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cTypeName] cpp_obj):
        __fbthrift_inst = <TypeName>TypeName.__new__(TypeName)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        __fbthrift_inst._load_cache()
        return __fbthrift_inst

    @property
    def boolType(self):
        if self.type.value != 1:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not boolType')
        return self.value

    @property
    def byteType(self):
        if self.type.value != 2:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not byteType')
        return self.value

    @property
    def i16Type(self):
        if self.type.value != 3:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not i16Type')
        return self.value

    @property
    def i32Type(self):
        if self.type.value != 4:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not i32Type')
        return self.value

    @property
    def i64Type(self):
        if self.type.value != 5:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not i64Type')
        return self.value

    @property
    def floatType(self):
        if self.type.value != 6:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not floatType')
        return self.value

    @property
    def doubleType(self):
        if self.type.value != 7:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not doubleType')
        return self.value

    @property
    def stringType(self):
        if self.type.value != 8:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not stringType')
        return self.value

    @property
    def binaryType(self):
        if self.type.value != 9:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not binaryType')
        return self.value

    @property
    def enumType(self):
        if self.type.value != 10:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not enumType')
        return self.value

    @property
    def typedefType(self):
        if self.type.value != 17:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not typedefType')
        return self.value

    @property
    def structType(self):
        if self.type.value != 11:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not structType')
        return self.value

    @property
    def unionType(self):
        if self.type.value != 12:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not unionType')
        return self.value

    @property
    def exceptionType(self):
        if self.type.value != 13:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not exceptionType')
        return self.value

    @property
    def listType(self):
        if self.type.value != 14:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not listType')
        return self.value

    @property
    def setType(self):
        if self.type.value != 15:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not setType')
        return self.value

    @property
    def mapType(self):
        if self.type.value != 16:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not mapType')
        return self.value


    def __hash__(TypeName self):
        return  super().__hash__()

    cdef _load_cache(TypeName self):
        self.type = TypeName.Type(<int>(deref(self._cpp_obj).getType()))
        cdef int type = self.type.value
        if type == 0:    # Empty
            self.value = None
        elif type == 1:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_boolType())
        elif type == 2:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_byteType())
        elif type == 3:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_i16Type())
        elif type == 4:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_i32Type())
        elif type == 5:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_i64Type())
        elif type == 6:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_floatType())
        elif type == 7:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_doubleType())
        elif type == 8:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_stringType())
        elif type == 9:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_binaryType())
        elif type == 10:
            self.value = TypeUri._fbthrift_create(make_shared[cTypeUri](deref(self._cpp_obj).get_enumType()))
        elif type == 17:
            self.value = TypeUri._fbthrift_create(make_shared[cTypeUri](deref(self._cpp_obj).get_typedefType()))
        elif type == 11:
            self.value = TypeUri._fbthrift_create(make_shared[cTypeUri](deref(self._cpp_obj).get_structType()))
        elif type == 12:
            self.value = TypeUri._fbthrift_create(make_shared[cTypeUri](deref(self._cpp_obj).get_unionType()))
        elif type == 13:
            self.value = TypeUri._fbthrift_create(make_shared[cTypeUri](deref(self._cpp_obj).get_exceptionType()))
        elif type == 14:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_listType())
        elif type == 15:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_setType())
        elif type == 16:
            self.value = translate_cpp_enum_to_python(Void, <int>deref(self._cpp_obj).get_mapType())

    def __copy__(TypeName self):
        cdef shared_ptr[cTypeName] cpp_obj = make_shared[cTypeName](
            deref(self._cpp_obj)
        )
        return TypeName._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cTypeName](
            self._cpp_obj,
            (<TypeName>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        import importlib
        types_reflection = importlib.import_module(
            "apache.thrift.type.standard.types_reflection"
        )
        return types_reflection.get_reflection__TypeName()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cTypeName].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "standard.TypeName"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cTypeName](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 17

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(TypeName self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cTypeName](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(TypeName self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cTypeName]()
        with nogil:
            needed = serializer.cdeserialize[cTypeName](buf, self._cpp_obj.get(), proto)
        # force a cache reload since the underlying data's changed
        self._load_cache()
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "apache.thrift.type.standard.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.TypeName, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("thrift.lib.thrift.standard.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.TypeName, self)
ByteString = bytes
ByteBuffer = _fbthrift_iobuf.IOBuf
