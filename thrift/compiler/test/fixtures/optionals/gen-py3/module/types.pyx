#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cpython.object cimport PyTypeObject, Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, Py_GE
from libcpp.memory cimport shared_ptr, make_shared, unique_ptr, make_unique
from libcpp.string cimport string
from libcpp cimport bool as cbool
from libcpp.iterator cimport inserter as cinserter
from cpython cimport bool as pbool
from cython.operator cimport dereference as deref, preincrement as inc, address as ptr_address
import thrift.py3.types
from thrift.py3.types import _IsSet as _fbthrift_IsSet
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
cimport thrift.py3.std_libcpp as std_libcpp
cimport thrift.py3.serializer as serializer
import folly.iobuf as _fbthrift_iobuf
from folly.optional cimport cOptional
from folly.memory cimport to_shared_ptr as __to_shared_ptr
from folly.range cimport Range as __cRange

import sys
from collections.abc import Sequence, Set, Mapping, Iterable
import weakref as __weakref
import builtins as _builtins

cimport module.types_reflection as _types_reflection


cdef __EnumData __Animal_enum_data  = __EnumData._fbthrift_create(thrift.py3.types.createEnumData[cAnimal](), Animal)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __AnimalMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __Animal_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __Animal_enum_data.get_all_names()

    def __len__(cls):
        return __Animal_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __Animal_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class Animal(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __Animal_enum_data.get_by_name(name)


    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        EnumMetadata[cAnimal].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Animal"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return python_types.Animal(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value


__SetMetaClass(<PyTypeObject*> Animal, <PyTypeObject*> __AnimalMeta)



@__cython.auto_pickle(False)
cdef class Color(thrift.py3.types.Struct):
    def __init__(Color self, **kwargs):
        self._cpp_obj = make_shared[cColor]()
        self._fields_setter = _fbthrift_types_fields.__Color_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(Color self, **kwargs):
        if not kwargs:
            return self
        cdef Color __fbthrift_inst = Color.__new__(Color)
        __fbthrift_inst._cpp_obj = make_shared[cColor](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Color_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Color", {
          "red": deref(self._cpp_obj).red_ref().has_value(),
          "green": deref(self._cpp_obj).green_ref().has_value(),
          "blue": deref(self._cpp_obj).blue_ref().has_value(),
          "alpha": deref(self._cpp_obj).alpha_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cColor] cpp_obj):
        __fbthrift_inst = <Color>Color.__new__(Color)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline red_impl(self):

        return deref(self._cpp_obj).red_ref().value()

    @property
    def red(self):
        return self.red_impl()

    cdef inline green_impl(self):

        return deref(self._cpp_obj).green_ref().value()

    @property
    def green(self):
        return self.green_impl()

    cdef inline blue_impl(self):

        return deref(self._cpp_obj).blue_ref().value()

    @property
    def blue(self):
        return self.blue_impl()

    cdef inline alpha_impl(self):

        return deref(self._cpp_obj).alpha_ref().value()

    @property
    def alpha(self):
        return self.alpha_impl()


    def __hash__(Color self):
        return super().__hash__()

    def __repr__(Color self):
        return super().__repr__()

    def __str__(Color self):
        return super().__str__()


    def __copy__(Color self):
        cdef shared_ptr[cColor] cpp_obj = make_shared[cColor](
            deref(self._cpp_obj)
        )
        return Color._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cColor](
            self._cpp_obj,
            (<Color>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__Color()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cColor].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Color"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cColor](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 4

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Color self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cColor](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Color self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cColor]()
        with nogil:
            needed = serializer.cdeserialize[cColor](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.Color, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Color, self)
@__cython.auto_pickle(False)
cdef class Vehicle(thrift.py3.types.Struct):
    def __init__(Vehicle self, **kwargs):
        self._cpp_obj = make_shared[cVehicle]()
        self._fields_setter = _fbthrift_types_fields.__Vehicle_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(Vehicle self, **kwargs):
        if not kwargs:
            return self
        cdef Vehicle __fbthrift_inst = Vehicle.__new__(Vehicle)
        __fbthrift_inst._cpp_obj = make_shared[cVehicle](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Vehicle_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Vehicle", {
          "color": deref(self._cpp_obj).color_ref().has_value(),
          "licensePlate": deref(self._cpp_obj).licensePlate_ref().has_value(),
          "description": deref(self._cpp_obj).description_ref().has_value(),
          "name": deref(self._cpp_obj).name_ref().has_value(),
          "hasAC": deref(self._cpp_obj).hasAC_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cVehicle] cpp_obj):
        __fbthrift_inst = <Vehicle>Vehicle.__new__(Vehicle)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline color_impl(self):

        if self.__fbthrift_cached_color is None:
            self.__fbthrift_cached_color = Color._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).color_ref().ref(), self._cpp_obj))
        return self.__fbthrift_cached_color

    @property
    def color(self):
        return self.color_impl()

    cdef inline licensePlate_impl(self):
        if not deref(self._cpp_obj).licensePlate_ref().has_value():
            return None

        return (<bytes>deref(self._cpp_obj).licensePlate_ref().value_unchecked()).decode('UTF-8')

    @property
    def licensePlate(self):
        return self.licensePlate_impl()

    cdef inline description_impl(self):
        if not deref(self._cpp_obj).description_ref().has_value():
            return None

        return (<bytes>deref(self._cpp_obj).description_ref().value_unchecked()).decode('UTF-8')

    @property
    def description(self):
        return self.description_impl()

    cdef inline name_impl(self):
        if not deref(self._cpp_obj).name_ref().has_value():
            return None

        return (<bytes>deref(self._cpp_obj).name_ref().value_unchecked()).decode('UTF-8')

    @property
    def name(self):
        return self.name_impl()

    cdef inline hasAC_impl(self):

        return <pbool> deref(self._cpp_obj).hasAC_ref().value_unchecked()

    @property
    def hasAC(self):
        return self.hasAC_impl()


    def __hash__(Vehicle self):
        return super().__hash__()

    def __repr__(Vehicle self):
        return super().__repr__()

    def __str__(Vehicle self):
        return super().__str__()


    def __copy__(Vehicle self):
        cdef shared_ptr[cVehicle] cpp_obj = make_shared[cVehicle](
            deref(self._cpp_obj)
        )
        return Vehicle._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cVehicle](
            self._cpp_obj,
            (<Vehicle>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__Vehicle()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cVehicle].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Vehicle"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cVehicle](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 5

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Vehicle self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cVehicle](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Vehicle self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cVehicle]()
        with nogil:
            needed = serializer.cdeserialize[cVehicle](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.Vehicle, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Vehicle, self)
@__cython.auto_pickle(False)
cdef class Person(thrift.py3.types.Struct):
    def __init__(Person self, **kwargs):
        self._cpp_obj = make_shared[cPerson]()
        self._fields_setter = _fbthrift_types_fields.__Person_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(Person self, **kwargs):
        if not kwargs:
            return self
        cdef Person __fbthrift_inst = Person.__new__(Person)
        __fbthrift_inst._cpp_obj = make_shared[cPerson](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Person_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Person", {
          "id": deref(self._cpp_obj).id_ref().has_value(),
          "name": deref(self._cpp_obj).name_ref().has_value(),
          "age": deref(self._cpp_obj).age_ref().has_value(),
          "address": deref(self._cpp_obj).address_ref().has_value(),
          "favoriteColor": deref(self._cpp_obj).favoriteColor_ref().has_value(),
          "friends": deref(self._cpp_obj).friends_ref().has_value(),
          "bestFriend": deref(self._cpp_obj).bestFriend_ref().has_value(),
          "petNames": deref(self._cpp_obj).petNames_ref().has_value(),
          "afraidOfAnimal": deref(self._cpp_obj).afraidOfAnimal_ref().has_value(),
          "vehicles": deref(self._cpp_obj).vehicles_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cPerson] cpp_obj):
        __fbthrift_inst = <Person>Person.__new__(Person)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline id_impl(self):

        return deref(self._cpp_obj).id_ref().value()

    @property
    def id(self):
        return self.id_impl()

    cdef inline name_impl(self):

        return (<bytes>deref(self._cpp_obj).name_ref().value()).decode('UTF-8')

    @property
    def name(self):
        return self.name_impl()

    cdef inline age_impl(self):
        if not deref(self._cpp_obj).age_ref().has_value():
            return None

        return deref(self._cpp_obj).age_ref().value_unchecked()

    @property
    def age(self):
        return self.age_impl()

    cdef inline address_impl(self):
        if not deref(self._cpp_obj).address_ref().has_value():
            return None

        return (<bytes>deref(self._cpp_obj).address_ref().value_unchecked()).decode('UTF-8')

    @property
    def address(self):
        return self.address_impl()

    cdef inline favoriteColor_impl(self):
        if not deref(self._cpp_obj).favoriteColor_ref().has_value():
            return None

        if self.__fbthrift_cached_favoriteColor is None:
            self.__fbthrift_cached_favoriteColor = Color._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).favoriteColor_ref().ref_unchecked(), self._cpp_obj))
        return self.__fbthrift_cached_favoriteColor

    @property
    def favoriteColor(self):
        return self.favoriteColor_impl()

    cdef inline friends_impl(self):
        if not deref(self._cpp_obj).friends_ref().has_value():
            return None

        if self.__fbthrift_cached_friends is None:
            self.__fbthrift_cached_friends = Set__i64._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).friends_ref().ref_unchecked(), self._cpp_obj))
        return self.__fbthrift_cached_friends

    @property
    def friends(self):
        return self.friends_impl()

    cdef inline bestFriend_impl(self):
        if not deref(self._cpp_obj).bestFriend_ref().has_value():
            return None

        return deref(self._cpp_obj).bestFriend_ref().value_unchecked()

    @property
    def bestFriend(self):
        return self.bestFriend_impl()

    cdef inline petNames_impl(self):
        if not deref(self._cpp_obj).petNames_ref().has_value():
            return None

        if self.__fbthrift_cached_petNames is None:
            self.__fbthrift_cached_petNames = Map__Animal_string._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).petNames_ref().ref_unchecked(), self._cpp_obj))
        return self.__fbthrift_cached_petNames

    @property
    def petNames(self):
        return self.petNames_impl()

    cdef inline afraidOfAnimal_impl(self):
        if not deref(self._cpp_obj).afraidOfAnimal_ref().has_value():
            return None

        if self.__fbthrift_cached_afraidOfAnimal is None:
            self.__fbthrift_cached_afraidOfAnimal = translate_cpp_enum_to_python(Animal, <int>(deref(self._cpp_obj).afraidOfAnimal_ref().value_unchecked()))
        return self.__fbthrift_cached_afraidOfAnimal

    @property
    def afraidOfAnimal(self):
        return self.afraidOfAnimal_impl()

    cdef inline vehicles_impl(self):
        if not deref(self._cpp_obj).vehicles_ref().has_value():
            return None

        if self.__fbthrift_cached_vehicles is None:
            self.__fbthrift_cached_vehicles = List__Vehicle._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).vehicles_ref().ref_unchecked(), self._cpp_obj))
        return self.__fbthrift_cached_vehicles

    @property
    def vehicles(self):
        return self.vehicles_impl()


    def __hash__(Person self):
        return super().__hash__()

    def __repr__(Person self):
        return super().__repr__()

    def __str__(Person self):
        return super().__str__()


    def __copy__(Person self):
        cdef shared_ptr[cPerson] cpp_obj = make_shared[cPerson](
            deref(self._cpp_obj)
        )
        return Person._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cPerson](
            self._cpp_obj,
            (<Person>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__Person()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cPerson].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Person"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cPerson](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 10

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Person self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cPerson](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Person self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cPerson]()
        with nogil:
            needed = serializer.cdeserialize[cPerson](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.Person, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Person, self)
@__cython.auto_pickle(False)
cdef class Set__i64(thrift.py3.types.Set):
    def __init__(self, items=None):
        if isinstance(items, Set__i64):
            self._cpp_obj = (<Set__i64> items)._cpp_obj
        else:
            self._cpp_obj = Set__i64._make_instance(items)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cset[cint64_t]] c_items):
        __fbthrift_inst = <Set__i64>Set__i64.__new__(Set__i64)
        __fbthrift_inst._cpp_obj = cmove(c_items)
        return __fbthrift_inst

    def __copy__(Set__i64 self):
        cdef shared_ptr[cset[cint64_t]] cpp_obj = make_shared[cset[cint64_t]](
            deref(self._cpp_obj)
        )
        return Set__i64._fbthrift_create(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj).size()

    @staticmethod
    cdef shared_ptr[cset[cint64_t]] _make_instance(object items) except *:
        cdef shared_ptr[cset[cint64_t]] c_inst = make_shared[cset[cint64_t]]()
        if items is not None:
            for item in items:
                if not isinstance(item, int):
                    raise TypeError(f"{item!r} is not of type int")
                item = <cint64_t> item
                deref(c_inst).insert(item)
        return c_inst

    def __contains__(self, item):
        if not self or item is None:
            return False
        if not isinstance(item, int):
            return False
        return pbool(deref(self._cpp_obj).count(item))


    def __iter__(self):
        if not self:
            return
        cdef __set_iter[cset[cint64_t]] itr = __set_iter[cset[cint64_t]](self._cpp_obj)
        cdef cint64_t citem = 0
        for i in range(deref(self._cpp_obj).size()):
            itr.genNext(self._cpp_obj, citem)
            yield citem

    def __hash__(self):
        return super().__hash__()

    def __richcmp__(self, other, int op):
        if isinstance(other, Set__i64):
            # C level comparisons
            return __setcmp(
                self._cpp_obj,
                (<Set__i64> other)._cpp_obj,
                op,
            )
        return self._fbthrift_py_richcmp(other, op)

    cdef _fbthrift_do_set_op(self, other, __cSetOp op):
        if not isinstance(other, Set__i64):
            other = Set__i64(other)
        cdef shared_ptr[cset[cint64_t]] result
        return Set__i64._fbthrift_create(__set_op[cset[cint64_t]](
            self._cpp_obj,
            (<Set__i64>other)._cpp_obj,
            op,
        ))

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__Set__i64()


Set.register(Set__i64)

@__cython.auto_pickle(False)
cdef class Map__Animal_string(thrift.py3.types.Map):
    def __init__(self, items=None):
        if isinstance(items, Map__Animal_string):
            self._cpp_obj = (<Map__Animal_string> items)._cpp_obj
        else:
            self._cpp_obj = Map__Animal_string._make_instance(items)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cmap[cAnimal,string]] c_items):
        __fbthrift_inst = <Map__Animal_string>Map__Animal_string.__new__(Map__Animal_string)
        __fbthrift_inst._cpp_obj = cmove(c_items)
        return __fbthrift_inst

    def __copy__(Map__Animal_string self):
        cdef shared_ptr[cmap[cAnimal,string]] cpp_obj = make_shared[cmap[cAnimal,string]](
            deref(self._cpp_obj)
        )
        return Map__Animal_string._fbthrift_create(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj).size()

    @staticmethod
    cdef shared_ptr[cmap[cAnimal,string]] _make_instance(object items) except *:
        cdef shared_ptr[cmap[cAnimal,string]] c_inst = make_shared[cmap[cAnimal,string]]()
        if items is not None:
            for key, item in items.items():
                if not isinstance(key, Animal):
                    raise TypeError(f"{key!r} is not of type Animal")
                if not isinstance(item, str):
                    raise TypeError(f"{item!r} is not of type str")

                deref(c_inst)[<cAnimal><int>key] = item.encode('UTF-8')
        return c_inst

    cdef _check_key_type(self, key):
        if not self or key is None:
            return
        if isinstance(key, Animal):
            return key

    def __getitem__(self, key):
        err = KeyError(f'{key}')
        key = self._check_key_type(key)
        if key is None:
            raise err
        cdef cAnimal ckey = <cAnimal><int>key
        if not __map_contains(self._cpp_obj, ckey):
            raise err
        cdef string citem
        __map_getitem(self._cpp_obj, ckey, citem)
        return bytes(citem).decode('UTF-8')

    def __iter__(self):
        if not self:
            return
        cdef __map_iter[cmap[cAnimal,string]] itr = __map_iter[cmap[cAnimal,string]](self._cpp_obj)
        cdef cAnimal citem
        for i in range(deref(self._cpp_obj).size()):
            itr.genNextKey(self._cpp_obj, citem)
            yield translate_cpp_enum_to_python(Animal, <int> citem)

    def __contains__(self, key):
        key = self._check_key_type(key)
        if key is None:
            return False
        cdef cAnimal ckey = <cAnimal><int>key
        return __map_contains(self._cpp_obj, ckey)

    def values(self):
        if not self:
            return
        cdef __map_iter[cmap[cAnimal,string]] itr = __map_iter[cmap[cAnimal,string]](self._cpp_obj)
        cdef string citem
        for i in range(deref(self._cpp_obj).size()):
            itr.genNextValue(self._cpp_obj, citem)
            yield bytes(citem).decode('UTF-8')

    def items(self):
        if not self:
            return
        cdef __map_iter[cmap[cAnimal,string]] itr = __map_iter[cmap[cAnimal,string]](self._cpp_obj)
        cdef cAnimal ckey
        cdef string citem
        for i in range(deref(self._cpp_obj).size()):
            itr.genNextItem(self._cpp_obj, ckey, citem)
            yield (translate_cpp_enum_to_python(Animal, <int> ckey), bytes(citem).decode('UTF-8'))

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__Map__Animal_string()

Mapping.register(Map__Animal_string)

@__cython.auto_pickle(False)
cdef class List__Vehicle(thrift.py3.types.List):
    def __init__(self, items=None):
        if isinstance(items, List__Vehicle):
            self._cpp_obj = (<List__Vehicle> items)._cpp_obj
        else:
            self._cpp_obj = List__Vehicle._make_instance(items)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[vector[cVehicle]] c_items):
        __fbthrift_inst = <List__Vehicle>List__Vehicle.__new__(List__Vehicle)
        __fbthrift_inst._cpp_obj = cmove(c_items)
        return __fbthrift_inst

    def __copy__(List__Vehicle self):
        cdef shared_ptr[vector[cVehicle]] cpp_obj = make_shared[vector[cVehicle]](
            deref(self._cpp_obj)
        )
        return List__Vehicle._fbthrift_create(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj).size()

    @staticmethod
    cdef shared_ptr[vector[cVehicle]] _make_instance(object items) except *:
        cdef shared_ptr[vector[cVehicle]] c_inst = make_shared[vector[cVehicle]]()
        if items is not None:
            for item in items:
                if not isinstance(item, Vehicle):
                    raise TypeError(f"{item!r} is not of type Vehicle")
                deref(c_inst).push_back(deref((<Vehicle>item)._cpp_obj))
        return c_inst

    cdef _get_slice(self, slice index_obj):
        cdef int start, stop, step
        start, stop, step = index_obj.indices(deref(self._cpp_obj).size())
        return List__Vehicle._fbthrift_create(
            __list_slice[vector[cVehicle]](self._cpp_obj, start, stop, step)
        )

    cdef _get_single_item(self, size_t index):
        cdef shared_ptr[cVehicle] citem
        __list_getitem(self._cpp_obj, index, citem)
        return Vehicle._fbthrift_create(citem)

    cdef _check_item_type(self, item):
        if not self or item is None:
            return
        if isinstance(item, Vehicle):
            return item

    def index(self, item, start=0, stop=None):
        err = ValueError(f'{item} is not in list')
        item = self._check_item_type(item)
        if item is None:
            raise err
        cdef (int, int, int) indices = slice(start, stop).indices(deref(self._cpp_obj).size())
        cdef cVehicle citem = deref((<Vehicle>item)._cpp_obj)
        cdef std_libcpp.optional[size_t] found = __list_index[vector[cVehicle]](self._cpp_obj, indices[0], indices[1], citem)
        if not found.has_value():
            raise err
        return found.value()

    def count(self, item):
        item = self._check_item_type(item)
        if item is None:
            return 0
        cdef cVehicle citem = deref((<Vehicle>item)._cpp_obj)
        return __list_count[vector[cVehicle]](self._cpp_obj, citem)

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__List__Vehicle()


Sequence.register(List__Vehicle)

PersonID = int
