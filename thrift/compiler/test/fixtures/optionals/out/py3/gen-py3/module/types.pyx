#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/optionals/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cpython.object cimport PyTypeObject
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
cimport thrift.python.exceptions
import thrift.python.converter
from thrift.python.types import EnumMeta as __EnumMeta
from thrift.python.std_libcpp cimport sv_to_str as __sv_to_str, string_view as __cstring_view
from thrift.python.types cimport BadEnum as __BadEnum
from thrift.py3.types cimport (
    richcmp as __richcmp,
    init_unicode_from_cpp as __init_unicode_from_cpp,
    set_iter as __set_iter,
    map_iter as __map_iter,
    map_contains as __map_contains,
    map_getitem as __map_getitem,
    reference_shared_ptr as __reference_shared_ptr,
    get_field_name_by_index as __get_field_name_by_index,
    reset_field as __reset_field,
    translate_cpp_enum_to_python,
    const_pointer_cast,
    make_const_shared,
    constant_shared_ptr,
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
import importlib

import module.thrift_types as _fbthrift_python_types
from module.types_impl_FBTHRIFT_ONLY_DO_NOT_USE import (
    Animal,
)

from module.containers_FBTHRIFT_ONLY_DO_NOT_USE import (
    Set__i64,
    List__Vehicle,
)


cdef object get_types_reflection():
    return importlib.import_module(
        "module.types_reflection"
    )

@__cython.auto_pickle(False)
cdef class Color(thrift.py3.types.Struct):
    def __init__(Color self, **kwargs):
        self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cColor]()
        self._fields_setter = _fbthrift_types_fields.__Color_FieldsSetter._fbthrift_create(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get())
        super().__init__(**kwargs)

    def __call__(Color self, **kwargs):
        if not kwargs:
            return self
        cdef Color __fbthrift_inst = Color.__new__(Color)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cColor](deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Color_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Color", {
          "red": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).red_ref().has_value(),
          "green": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).green_ref().has_value(),
          "blue": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).blue_ref().has_value(),
          "alpha": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).alpha_ref().has_value(),
        })

    @staticmethod
    cdef _create_FBTHRIFT_ONLY_DO_NOT_USE(shared_ptr[_module_cbindings.cColor] cpp_obj):
        __fbthrift_inst = <Color>Color.__new__(Color)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline red_impl(self):
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).red_ref().value()

    @property
    def red(self):
        return self.red_impl()

    cdef inline green_impl(self):
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).green_ref().value()

    @property
    def green(self):
        return self.green_impl()

    cdef inline blue_impl(self):
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).blue_ref().value()

    @property
    def blue(self):
        return self.blue_impl()

    cdef inline alpha_impl(self):
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).alpha_ref().value()

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
        cdef shared_ptr[_module_cbindings.cColor] cpp_obj = make_shared[_module_cbindings.cColor](
            deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        )
        return Color._create_FBTHRIFT_ONLY_DO_NOT_USE(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[_module_cbindings.cColor](
            self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE,
            (<Color>other)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Color()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        _module_cbindings.StructMetadata[_module_cbindings.cColor].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Color"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[_module_cbindings.cColor](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 4

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Color self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[_module_cbindings.cColor](self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Color self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cColor]()
        with nogil:
            needed = serializer.cdeserialize[_module_cbindings.cColor](buf, self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get(), proto)
        return needed


    def _to_python(self):
        return thrift.python.converter.to_python_struct(
            _fbthrift_python_types.Color,
            self,
        )

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Color, self)

@__cython.auto_pickle(False)
cdef class Vehicle(thrift.py3.types.Struct):
    def __init__(Vehicle self, **kwargs):
        self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cVehicle]()
        self._fields_setter = _fbthrift_types_fields.__Vehicle_FieldsSetter._fbthrift_create(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get())
        super().__init__(**kwargs)

    def __call__(Vehicle self, **kwargs):
        if not kwargs:
            return self
        cdef Vehicle __fbthrift_inst = Vehicle.__new__(Vehicle)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cVehicle](deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Vehicle_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Vehicle", {
          "color": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).color_ref().has_value(),
          "licensePlate": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).licensePlate_ref().has_value(),
          "description": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).description_ref().has_value(),
          "name": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).name_ref().has_value(),
          "hasAC": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).hasAC_ref().has_value(),
        })

    @staticmethod
    cdef _create_FBTHRIFT_ONLY_DO_NOT_USE(shared_ptr[_module_cbindings.cVehicle] cpp_obj):
        __fbthrift_inst = <Vehicle>Vehicle.__new__(Vehicle)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline color_impl(self):
        if self.__fbthrift_cached_color is None:
            self.__fbthrift_cached_color = Color._create_FBTHRIFT_ONLY_DO_NOT_USE(__reference_shared_ptr(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).color_ref().ref(), self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
        return self.__fbthrift_cached_color

    @property
    def color(self):
        return self.color_impl()

    cdef inline licensePlate_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).licensePlate_ref().has_value():
            return None
        return (<bytes>deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).licensePlate_ref().value_unchecked()).decode('UTF-8')

    @property
    def licensePlate(self):
        return self.licensePlate_impl()

    cdef inline description_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).description_ref().has_value():
            return None
        return (<bytes>deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).description_ref().value_unchecked()).decode('UTF-8')

    @property
    def description(self):
        return self.description_impl()

    cdef inline name_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).name_ref().has_value():
            return None
        return (<bytes>deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).name_ref().value_unchecked()).decode('UTF-8')

    @property
    def name(self):
        return self.name_impl()

    cdef inline hasAC_impl(self):
        return <pbool> deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).hasAC_ref().value_unchecked()

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
        cdef shared_ptr[_module_cbindings.cVehicle] cpp_obj = make_shared[_module_cbindings.cVehicle](
            deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        )
        return Vehicle._create_FBTHRIFT_ONLY_DO_NOT_USE(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[_module_cbindings.cVehicle](
            self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE,
            (<Vehicle>other)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Vehicle()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        _module_cbindings.StructMetadata[_module_cbindings.cVehicle].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Vehicle"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[_module_cbindings.cVehicle](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 5

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Vehicle self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[_module_cbindings.cVehicle](self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Vehicle self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cVehicle]()
        with nogil:
            needed = serializer.cdeserialize[_module_cbindings.cVehicle](buf, self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get(), proto)
        return needed


    def _to_python(self):
        return thrift.python.converter.to_python_struct(
            _fbthrift_python_types.Vehicle,
            self,
        )

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Vehicle, self)

@__cython.auto_pickle(False)
cdef class Person(thrift.py3.types.Struct):
    def __init__(Person self, **kwargs):
        self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cPerson]()
        self._fields_setter = _fbthrift_types_fields.__Person_FieldsSetter._fbthrift_create(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get())
        super().__init__(**kwargs)

    def __call__(Person self, **kwargs):
        if not kwargs:
            return self
        cdef Person __fbthrift_inst = Person.__new__(Person)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cPerson](deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Person_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Person", {
          "id": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).id_ref().has_value(),
          "name": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).name_ref().has_value(),
          "age": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).age_ref().has_value(),
          "address": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).address_ref().has_value(),
          "favoriteColor": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).favoriteColor_ref().has_value(),
          "friends": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).friends_ref().has_value(),
          "bestFriend": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).bestFriend_ref().has_value(),
          "petNames": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).petNames_ref().has_value(),
          "afraidOfAnimal": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).afraidOfAnimal_ref().has_value(),
          "vehicles": deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).vehicles_ref().has_value(),
        })

    @staticmethod
    cdef _create_FBTHRIFT_ONLY_DO_NOT_USE(shared_ptr[_module_cbindings.cPerson] cpp_obj):
        __fbthrift_inst = <Person>Person.__new__(Person)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline id_impl(self):
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).id_ref().value()

    @property
    def id(self):
        return self.id_impl()

    cdef inline name_impl(self):
        return (<bytes>deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).name_ref().value()).decode('UTF-8')

    @property
    def name(self):
        return self.name_impl()

    cdef inline age_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).age_ref().has_value():
            return None
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).age_ref().value_unchecked()

    @property
    def age(self):
        return self.age_impl()

    cdef inline address_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).address_ref().has_value():
            return None
        return (<bytes>deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).address_ref().value_unchecked()).decode('UTF-8')

    @property
    def address(self):
        return self.address_impl()

    cdef inline favoriteColor_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).favoriteColor_ref().has_value():
            return None
        if self.__fbthrift_cached_favoriteColor is None:
            self.__fbthrift_cached_favoriteColor = Color._create_FBTHRIFT_ONLY_DO_NOT_USE(__reference_shared_ptr(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).favoriteColor_ref().ref_unchecked(), self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
        return self.__fbthrift_cached_favoriteColor

    @property
    def favoriteColor(self):
        return self.favoriteColor_impl()

    cdef inline friends_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).friends_ref().has_value():
            return None
        if self.__fbthrift_cached_friends is None:
            self.__fbthrift_cached_friends = Set__i64__from_cpp(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).friends_ref().ref_unchecked())
        return self.__fbthrift_cached_friends

    @property
    def friends(self):
        return self.friends_impl()

    cdef inline bestFriend_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).bestFriend_ref().has_value():
            return None
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).bestFriend_ref().value_unchecked()

    @property
    def bestFriend(self):
        return self.bestFriend_impl()

    cdef inline petNames_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).petNames_ref().has_value():
            return None
        if self.__fbthrift_cached_petNames is None:
            self.__fbthrift_cached_petNames = Map__Animal_string._create_FBTHRIFT_ONLY_DO_NOT_USE(__reference_shared_ptr(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).petNames_ref().ref_unchecked(), self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
        return self.__fbthrift_cached_petNames

    @property
    def petNames(self):
        return self.petNames_impl()

    cdef inline afraidOfAnimal_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).afraidOfAnimal_ref().has_value():
            return None
        if self.__fbthrift_cached_afraidOfAnimal is None:
            self.__fbthrift_cached_afraidOfAnimal = translate_cpp_enum_to_python(Animal, <int>(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).afraidOfAnimal_ref().value_unchecked()))
        return self.__fbthrift_cached_afraidOfAnimal

    @property
    def afraidOfAnimal(self):
        return self.afraidOfAnimal_impl()

    cdef inline vehicles_impl(self):
        if not deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).vehicles_ref().has_value():
            return None
        if self.__fbthrift_cached_vehicles is None:
            self.__fbthrift_cached_vehicles = List__Vehicle__from_cpp(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).vehicles_ref().ref_unchecked())
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
        cdef shared_ptr[_module_cbindings.cPerson] cpp_obj = make_shared[_module_cbindings.cPerson](
            deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        )
        return Person._create_FBTHRIFT_ONLY_DO_NOT_USE(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[_module_cbindings.cPerson](
            self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE,
            (<Person>other)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Person()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        _module_cbindings.StructMetadata[_module_cbindings.cPerson].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Person"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[_module_cbindings.cPerson](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 10

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Person self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[_module_cbindings.cPerson](self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Person self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = make_shared[_module_cbindings.cPerson]()
        with nogil:
            needed = serializer.cdeserialize[_module_cbindings.cPerson](buf, self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE.get(), proto)
        return needed


    def _to_python(self):
        return thrift.python.converter.to_python_struct(
            _fbthrift_python_types.Person,
            self,
        )

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Person, self)


cdef cset[cint64_t] Set__i64__make_instance(object items) except *:
    cdef cset[cint64_t] c_inst
    if items is not None:
        for item in items:
            if not isinstance(item, int):
                raise TypeError(f"{item!r} is not of type int")
            item = <cint64_t> item
            c_inst.insert(item)
    return cmove(c_inst)

cdef object Set__i64__from_cpp(const cset[cint64_t]& c_set) except *:
    cdef list py_items = []
    cdef __set_iter[cset[cint64_t]] iter = __set_iter[cset[cint64_t]](c_set)
    cdef cint64_t citem = 0
    for i in range(c_set.size()):
        iter.genNextItem(citem)
        py_items.append(citem)
    return Set__i64(frozenset(py_items), thrift.py3.types._fbthrift_set_private_ctor)

@__cython.auto_pickle(False)
@__cython.final
cdef class Map__Animal_string(thrift.py3.types.Map):
    def __init__(self, items=None):
        if isinstance(items, Map__Animal_string):
            self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = (<Map__Animal_string> items)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
        else:
            self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = Map__Animal_string__make_instance(items)

    @staticmethod
    cdef _create_FBTHRIFT_ONLY_DO_NOT_USE(shared_ptr[cmap[_module_cbindings.cAnimal,string]] c_items):
        __fbthrift_inst = <Map__Animal_string>Map__Animal_string.__new__(Map__Animal_string)
        __fbthrift_inst._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE = cmove(c_items)
        return __fbthrift_inst

    def __copy__(Map__Animal_string self):
        cdef shared_ptr[cmap[_module_cbindings.cAnimal,string]] cpp_obj = make_shared[cmap[_module_cbindings.cAnimal,string]](
            deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        )
        return Map__Animal_string._create_FBTHRIFT_ONLY_DO_NOT_USE(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).size()

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
        cdef _module_cbindings.cAnimal ckey = <_module_cbindings.cAnimal><int>key
        if not __map_contains(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE, ckey):
            raise err
        cdef string citem
        __map_getitem(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE, ckey, citem)
        return bytes(citem).decode('UTF-8')

    def __iter__(self):
        if not self:
            return
        cdef __map_iter[cmap[_module_cbindings.cAnimal,string]] itr = __map_iter[cmap[_module_cbindings.cAnimal,string]](self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        cdef _module_cbindings.cAnimal citem
        for i in range(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).size()):
            itr.genNextKey(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE, citem)
            yield translate_cpp_enum_to_python(Animal, <int> citem)

    def __contains__(self, key):
        key = self._check_key_type(key)
        if key is None:
            return False
        cdef _module_cbindings.cAnimal ckey = <_module_cbindings.cAnimal><int>key
        return __map_contains(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE, ckey)

    def values(self):
        if not self:
            return
        cdef __map_iter[cmap[_module_cbindings.cAnimal,string]] itr = __map_iter[cmap[_module_cbindings.cAnimal,string]](self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        cdef string citem
        for i in range(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).size()):
            itr.genNextValue(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE, citem)
            yield bytes(citem).decode('UTF-8')

    def items(self):
        if not self:
            return
        cdef __map_iter[cmap[_module_cbindings.cAnimal,string]] itr = __map_iter[cmap[_module_cbindings.cAnimal,string]](self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE)
        cdef _module_cbindings.cAnimal ckey
        cdef string citem
        for i in range(deref(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE).size()):
            itr.genNextItem(self._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE, ckey, citem)
            yield (translate_cpp_enum_to_python(Animal, <int> ckey), bytes(citem).decode('UTF-8'))

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__Animal_string()

Mapping.register(Map__Animal_string)

cdef shared_ptr[cmap[_module_cbindings.cAnimal,string]] Map__Animal_string__make_instance(object items) except *:
    cdef shared_ptr[cmap[_module_cbindings.cAnimal,string]] c_inst = make_shared[cmap[_module_cbindings.cAnimal,string]]()
    if items is not None:
        for key, item in items.items():
            if not isinstance(key, Animal):
                raise TypeError(f"{key!r} is not of type Animal")
            if not isinstance(item, str):
                raise TypeError(f"{item!r} is not of type str")

            deref(c_inst)[<_module_cbindings.cAnimal><int>key] = item.encode('UTF-8')
    return cmove(c_inst)

cdef object Map__Animal_string__from_cpp(const cmap[_module_cbindings.cAnimal,string]& c_map) except *:
    cdef dict py_items = {}
    cdef __map_iter[cmap[_module_cbindings.cAnimal,string]] iter = __map_iter[cmap[_module_cbindings.cAnimal,string]](c_map)
    cdef _module_cbindings.cAnimal ckey
    cdef string cval
    for i in range(c_map.size()):
        iter.genNextKeyVal(ckey, cval)
        py_items[translate_cpp_enum_to_python(Animal, <int> ckey)] = __init_unicode_from_cpp(cval)
    return Map__Animal_string(py_items)


cdef vector[_module_cbindings.cVehicle] List__Vehicle__make_instance(object items) except *:
    cdef vector[_module_cbindings.cVehicle] c_inst
    if items is not None:
        for item in items:
            if not isinstance(item, Vehicle):
                raise TypeError(f"{item!r} is not of type Vehicle")
            c_inst.push_back(deref((<Vehicle>item)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE))
    return cmove(c_inst)

cdef object List__Vehicle__from_cpp(const vector[_module_cbindings.cVehicle]& c_vec) except *:
    cdef list py_list = []
    cdef int idx = 0
    for idx in range(c_vec.size()):
        py_list.append(Vehicle._create_FBTHRIFT_ONLY_DO_NOT_USE(make_shared[_module_cbindings.cVehicle](c_vec[idx])))
    return List__Vehicle(py_list, thrift.py3.types._fbthrift_list_private_ctor)


PersonID = int
