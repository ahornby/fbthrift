#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/basic/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from __future__ import annotations

from collections.abc import Mapping, Sequence, Set
import enum
import importlib

import typing as _typing
import thrift.py3.types
import thrift.python.types
import test.fixtures.basic.module.thrift_enums as _fbthrift_python_enums
import test.fixtures.basic.module.thrift_types as _fbthrift_python_types



def get_types_reflection():
    return importlib.import_module(
        "test.fixtures.basic.module.types_reflection"
    )

_fbthrift__module_name__ = "test.fixtures.basic.module.types"

__all__ = []

### Enums ###

class MyEnum(thrift.py3.types.CompiledEnum, int):
    MyValue1 = 0
    MyValue2 = 1

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_MyEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.MyEnum"

    def _to_python(self):
        return _fbthrift_python_enums.MyEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    @staticmethod
    def from_python(python_enum: _fbthrift_python_enums.MyEnum) -> MyEnum:
        if isinstance(python_enum, thrift.python.types.BadEnum):
            return thrift.python.types.BadEnum(MyEnum, int(python_enum))
        return python_enum._to_py3()


    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_
__all__.append("MyEnum")


class HackEnum(thrift.py3.types.CompiledEnum, int):
    Value1 = 0
    Value2 = 1

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_HackEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.HackEnum"

    def _to_python(self):
        return _fbthrift_python_enums.HackEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    @staticmethod
    def from_python(python_enum: _fbthrift_python_enums.HackEnum) -> HackEnum:
        if isinstance(python_enum, thrift.python.types.BadEnum):
            return thrift.python.types.BadEnum(HackEnum, int(python_enum))
        return python_enum._to_py3()


    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_
__all__.append("HackEnum")


### Union Enums ###

class __MyUnionType(enum.Enum):
    myEnum = 1
    myStruct = 2
    myDataItem = 3
    floatSet = 4
    EMPTY = 0

    __module__ = _fbthrift__module_name__
    __slots__ = ()
__all__.append("__MyUnionType")


class __UnionToBeRenamedType(enum.Enum):
    reserved_field = 1
    EMPTY = 0

    __module__ = _fbthrift__module_name__
    __slots__ = ()
__all__.append("__UnionToBeRenamedType")


### Containers ###
class Set__float(thrift.py3.types.Set):
    __module__ = _fbthrift__module_name__
    __slots__ = ()

    def __init__(self, items=None, private_ctor_token=None) -> None:
        if private_ctor_token is thrift.py3.types._fbthrift_set_private_ctor:
            _py_obj = items
        elif isinstance(items, Set__float):
            _py_obj = frozenset(items)
        elif items is None:
            _py_obj = frozenset()
        else:
            check_method = Set__float._check_item_type_or_raise
            _py_obj = frozenset(check_method(item) for item in items)

        super().__init__(_py_obj, Set__float)

    @staticmethod
    def _check_item_type_or_raise(item):
        if not (
            isinstance(item, (float, int))
        ):
            raise TypeError(f"{item!r} is not of type float")
        return item

    @staticmethod
    def _check_item_type_or_none(item):
        if item is None:
            return None
        if isinstance(item, float):
            return item

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Set__float()

    @staticmethod
    def from_python(python_set: thrift.python.types.Set) -> Set__float:
        _items = frozenset(python_set)
        return Set__float(
            items=_items,
            private_ctor_token=thrift.py3.types._fbthrift_set_private_ctor,
        )


Set.register(Set__float)


__all__.append("Set__float")

class List__i32(thrift.py3.types.List):
    __module__ = _fbthrift__module_name__
    __slots__ = ()

    def __init__(self, items=None, private_ctor_token=None) -> None:
        if private_ctor_token is thrift.py3.types._fbthrift_list_private_ctor:
            _py_obj = items
        elif isinstance(items, List__i32):
            _py_obj = list(items)
        elif items is None:
            _py_obj = []
        else:
            check_method = List__i32._check_item_type_or_raise
            _py_obj = [check_method(item) for item in items]

        super().__init__(_py_obj, List__i32)

    @staticmethod
    def _check_item_type_or_raise(item):
        if not (
            isinstance(item, int)
        ):
            raise TypeError(f"{item!r} is not of type int")
        return item

    @staticmethod
    def _check_item_type_or_none(item):
        if item is None:
            return None
        if isinstance(item, int):
            return item

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__i32()

    @staticmethod
    def from_python(python_list: thrift.python.types.List) -> List__i32:
        _items = list(python_list)
        return List__i32(
            items=_items,
            private_ctor_token=thrift.py3.types._fbthrift_list_private_ctor,
        )


Sequence.register(List__i32)


__all__.append("List__i32")

class Set__string(thrift.py3.types.Set):
    __module__ = _fbthrift__module_name__
    __slots__ = ()

    def __init__(self, items=None, private_ctor_token=None) -> None:
        if private_ctor_token is thrift.py3.types._fbthrift_set_private_ctor:
            _py_obj = items
        elif isinstance(items, Set__string):
            _py_obj = frozenset(items)
        elif items is None:
            _py_obj = frozenset()
        else:
            if isinstance(items, str):
                raise TypeError("If you really want to pass a string into a _typing.AbstractSet[str] field, explicitly convert it first.")
            check_method = Set__string._check_item_type_or_raise
            _py_obj = frozenset(check_method(item) for item in items)

        super().__init__(_py_obj, Set__string)

    @staticmethod
    def _check_item_type_or_raise(item):
        if not (
            isinstance(item, str)
        ):
            raise TypeError(f"{item!r} is not of type str")
        return item

    @staticmethod
    def _check_item_type_or_none(item):
        if item is None:
            return None
        if isinstance(item, str):
            return item

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Set__string()

    @staticmethod
    def from_python(python_set: thrift.python.types.Set) -> Set__string:
        _items = frozenset(python_set)
        return Set__string(
            items=_items,
            private_ctor_token=thrift.py3.types._fbthrift_set_private_ctor,
        )


Set.register(Set__string)


__all__.append("Set__string")

class Map__string_i64(thrift.py3.types.Map):
    __module__ = _fbthrift__module_name__
    __slots__ = ()

    _FBTHRIFT_USE_SORTED_REPR = True

    def __init__(self, items=None, private_ctor_token=None) -> None:
        if private_ctor_token is thrift.py3.types._fbthrift_map_private_ctor:
            _py_obj = items
        elif isinstance(items, Map__string_i64):
            _py_obj = dict(items)
        elif items is None:
            _py_obj = dict()
        else:
            check_key = Map__string_i64._check_key_type_or_raise
            check_val = Map__string_i64._check_val_type_or_raise
            _py_obj = {check_key(k) : check_val(v) for k, v in items.items()}

        super().__init__(_py_obj, Map__string_i64)

    @staticmethod
    def _check_key_type_or_raise(key):
        if not (
            isinstance(key, str)
        ):
            raise TypeError(f"{key!r} is not of type str")
        return key

    @staticmethod
    def _check_key_type_or_none(key):
        if key is None:
            return None
        if isinstance(key, str):
            return key

    @staticmethod
    def _check_val_type_or_raise(item):
        if not (
            isinstance(item, int)
        ):
            raise TypeError(f"{item!r} is not of type int")
        return item

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_i64()

    @staticmethod
    def from_python(python_map: thrift.python.types.Map) -> Map__string_i64:
        _keys = python_map.keys()
        _values = python_map.values()
        return Map__string_i64(
            items=dict(zip(_keys, _values)),
            private_ctor_token=thrift.py3.types._fbthrift_map_private_ctor,
        )


Mapping.register(Map__string_i64)

__all__.append("Map__string_i64")

class Map__string_List__i32(thrift.py3.types.Map):
    __module__ = _fbthrift__module_name__
    __slots__ = ()

    _FBTHRIFT_USE_SORTED_REPR = True

    def __init__(self, items=None, private_ctor_token=None) -> None:
        if private_ctor_token is thrift.py3.types._fbthrift_map_private_ctor:
            _py_obj = items
        elif isinstance(items, Map__string_List__i32):
            _py_obj = dict(items)
        elif items is None:
            _py_obj = dict()
        else:
            check_key = Map__string_List__i32._check_key_type_or_raise
            check_val = Map__string_List__i32._check_val_type_or_raise
            _py_obj = {check_key(k) : check_val(v) for k, v in items.items()}

        super().__init__(_py_obj, Map__string_List__i32)

    @staticmethod
    def _check_key_type_or_raise(key):
        if not (
            isinstance(key, str)
        ):
            raise TypeError(f"{key!r} is not of type str")
        return key

    @staticmethod
    def _check_key_type_or_none(key):
        if key is None:
            return None
        if isinstance(key, str):
            return key

    @staticmethod
    def _check_val_type_or_raise(item):
        if item is None:
            raise TypeError("None is not of the type _typing.Sequence[int]")
        if not isinstance(item, List__i32):
            item = List__i32(item)
        return item

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_List__i32()

    @staticmethod
    def from_python(python_map: thrift.python.types.Map) -> Map__string_List__i32:
        _keys = python_map.keys()
        _values = (
            List__i32.from_python(item)
            for item in python_map.values()
        )
        return Map__string_List__i32(
            items=dict(zip(_keys, _values)),
            private_ctor_token=thrift.py3.types._fbthrift_map_private_ctor,
        )


Mapping.register(Map__string_List__i32)

__all__.append("Map__string_List__i32")


### Structured Types ###
class MyStruct(thrift.py3.types.Struct):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
        "_fbthrift_inner__MyDataField",
        "_fbthrift_inner__myEnum",
        "_fbthrift_inner__floatSet",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.MyStruct
    _FBTHRIFT__FIELD_NAMES = (
        "MyIntField",
        "MyStringField",
        "MyDataField",
        "myEnum",
        "oneway",
        "readonly",
        "idempotent",
        "floatSet",
        "no_hack_codegen_field",
    )
    _fbthrift__inner : _fbthrift_python_types.MyStruct
    _fbthrift_inner__MyDataField : MyDataItem | None
    _fbthrift_inner__myEnum : MyEnum | None
    _fbthrift_inner__floatSet : _typing.AbstractSet[float] | None


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.MyStruct(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> MyStruct:
        instance = super().__new__(cls)
        instance._fbthrift_inner__MyDataField = None
        instance._fbthrift_inner__myEnum = None
        instance._fbthrift_inner__floatSet = None
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.MyStruct) -> MyStruct:
        inst = MyStruct.__new__(MyStruct)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> MyStruct:
        return self

    def _to_python(self) -> _fbthrift_python_types.MyStruct:
        return self._fbthrift__inner


    @property
    def MyIntField(self) -> int:
        return self._fbthrift__inner.MyIntField

    @property
    def MyStringField(self) -> str:
        return self._fbthrift__inner.MyStringField

    @property
    def MyDataField(self) -> MyDataItem:
        if self._fbthrift_inner__MyDataField is None:
            __python_val = self._fbthrift__inner.MyDataField
            self._fbthrift_inner__MyDataField = MyDataItem.from_python(__python_val)

        return self._fbthrift_inner__MyDataField

    @property
    def myEnum(self) -> MyEnum:
        if self._fbthrift_inner__myEnum is None:
            __python_val = self._fbthrift__inner.myEnum
            self._fbthrift_inner__myEnum = MyEnum.from_python(__python_val)

        return self._fbthrift_inner__myEnum

    @property
    def oneway(self) -> bool:
        return self._fbthrift__inner.oneway

    @property
    def readonly(self) -> bool:
        return self._fbthrift__inner.readonly

    @property
    def idempotent(self) -> bool:
        return self._fbthrift__inner.idempotent

    @property
    def floatSet(self) -> _typing.AbstractSet[float]:
        if self._fbthrift_inner__floatSet is None:
            __python_val = self._fbthrift__inner.floatSet
            self._fbthrift_inner__floatSet = Set__float.from_python(__python_val)

        return self._fbthrift_inner__floatSet

    @property
    def no_hack_codegen_field(self) -> str:
        return self._fbthrift__inner.no_hack_codegen_field




__all__.append("MyStruct")

class Containers(thrift.py3.types.Struct):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
        "_fbthrift_inner__I32List",
        "_fbthrift_inner__StringSet",
        "_fbthrift_inner__StringToI64Map",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.Containers
    _FBTHRIFT__FIELD_NAMES = (
        "I32List",
        "StringSet",
        "StringToI64Map",
    )
    _fbthrift__inner : _fbthrift_python_types.Containers
    _fbthrift_inner__I32List : _typing.Sequence[int] | None
    _fbthrift_inner__StringSet : _typing.AbstractSet[str] | None
    _fbthrift_inner__StringToI64Map : _typing.Mapping[str, int] | None


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.Containers(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> Containers:
        instance = super().__new__(cls)
        instance._fbthrift_inner__I32List = None
        instance._fbthrift_inner__StringSet = None
        instance._fbthrift_inner__StringToI64Map = None
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.Containers) -> Containers:
        inst = Containers.__new__(Containers)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> Containers:
        return self

    def _to_python(self) -> _fbthrift_python_types.Containers:
        return self._fbthrift__inner


    @property
    def I32List(self) -> _typing.Sequence[int]:
        if self._fbthrift_inner__I32List is None:
            __python_val = self._fbthrift__inner.I32List
            self._fbthrift_inner__I32List = List__i32.from_python(__python_val)

        return self._fbthrift_inner__I32List

    @property
    def StringSet(self) -> _typing.AbstractSet[str]:
        if self._fbthrift_inner__StringSet is None:
            __python_val = self._fbthrift__inner.StringSet
            self._fbthrift_inner__StringSet = Set__string.from_python(__python_val)

        return self._fbthrift_inner__StringSet

    @property
    def StringToI64Map(self) -> _typing.Mapping[str, int]:
        if self._fbthrift_inner__StringToI64Map is None:
            __python_val = self._fbthrift__inner.StringToI64Map
            self._fbthrift_inner__StringToI64Map = Map__string_i64.from_python(__python_val)

        return self._fbthrift_inner__StringToI64Map




__all__.append("Containers")

class MyDataItem(thrift.py3.types.Struct):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.MyDataItem
    _FBTHRIFT__FIELD_NAMES = (
    )
    _fbthrift__inner : _fbthrift_python_types.MyDataItem


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.MyDataItem(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> MyDataItem:
        instance = super().__new__(cls)
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.MyDataItem) -> MyDataItem:
        inst = MyDataItem.__new__(MyDataItem)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> MyDataItem:
        return self

    def _to_python(self) -> _fbthrift_python_types.MyDataItem:
        return self._fbthrift__inner





__all__.append("MyDataItem")

class MyUnion(thrift.py3.types.Union):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
        "_fbthrift_inner__myEnum",
        "_fbthrift_inner__myStruct",
        "_fbthrift_inner__myDataItem",
        "_fbthrift_inner__floatSet",
        "_fbthrift_inner__type",
        "_fbthrift_inner__value",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.MyUnion
    _FBTHRIFT__FIELD_NAMES = (
        "myEnum",
        "myStruct",
        "myDataItem",
        "floatSet",
    )
    Type = _fbthrift_python_types.MyUnion.Type
    _fbthrift__inner : _fbthrift_python_types.MyUnion
    _fbthrift_inner__myEnum : MyEnum | None
    _fbthrift_inner__myStruct : MyStruct | None
    _fbthrift_inner__myDataItem : MyDataItem | None
    _fbthrift_inner__floatSet : _typing.AbstractSet[float] | None
    _fbthrift_inner__type: Type
    _fbthrift_inner__value: MyEnum | MyStruct | MyDataItem | _typing.AbstractSet[float] | None


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.MyUnion(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> MyUnion:
        instance = super().__new__(cls)
        instance._fbthrift_inner__myEnum = None
        instance._fbthrift_inner__myStruct = None
        instance._fbthrift_inner__myDataItem = None
        instance._fbthrift_inner__floatSet = None
        instance._fbthrift_inner__type = None
        instance._fbthrift_inner__value = None
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.MyUnion) -> MyUnion:
        inst = MyUnion.__new__(MyUnion)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> MyUnion:
        return self

    def _to_python(self) -> _fbthrift_python_types.MyUnion:
        return self._fbthrift__inner


    @property
    def myEnum(self) -> MyEnum:
        if self._fbthrift_inner__myEnum is None:
            __python_val = self._fbthrift__inner.myEnum
            if __python_val is None:
                return None
            self._fbthrift_inner__myEnum = MyEnum.from_python(__python_val)

        return self._fbthrift_inner__myEnum

    @property
    def myStruct(self) -> MyStruct:
        if self._fbthrift_inner__myStruct is None:
            __python_val = self._fbthrift__inner.myStruct
            if __python_val is None:
                return None
            self._fbthrift_inner__myStruct = MyStruct.from_python(__python_val)

        return self._fbthrift_inner__myStruct

    @property
    def myDataItem(self) -> MyDataItem:
        if self._fbthrift_inner__myDataItem is None:
            __python_val = self._fbthrift__inner.myDataItem
            if __python_val is None:
                return None
            self._fbthrift_inner__myDataItem = MyDataItem.from_python(__python_val)

        return self._fbthrift_inner__myDataItem

    @property
    def floatSet(self) -> _typing.AbstractSet[float]:
        if self._fbthrift_inner__floatSet is None:
            __python_val = self._fbthrift__inner.floatSet
            if __python_val is None:
                return None
            self._fbthrift_inner__floatSet = Set__float.from_python(__python_val)

        return self._fbthrift_inner__floatSet

    @property
    def type(self) -> _fbthrift_python_types.MyUnion.Type:
        return self._fbthrift__inner.type

    @property
    def value(self) -> MyEnum | MyStruct | MyDataItem | _typing.AbstractSet[float] | None:
        match self._fbthrift__inner.type:
            case _fbthrift_python_types.MyUnion.Type.myEnum:
                return self.myEnum
            case _fbthrift_python_types.MyUnion.Type.myStruct:
                return self.myStruct
            case _fbthrift_python_types.MyUnion.Type.myDataItem:
                return self.myDataItem
            case _fbthrift_python_types.MyUnion.Type.floatSet:
                return self.floatSet
            case _:
                return self._fbthrift__inner.value




__all__.append("MyUnion")

class MyException(thrift.py3.types.Struct):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
        "_fbthrift_inner__myStruct",
        "_fbthrift_inner__myUnion",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.MyException
    _FBTHRIFT__FIELD_NAMES = (
        "MyIntField",
        "MyStringField",
        "myStruct",
        "myUnion",
    )
    _fbthrift__inner : _fbthrift_python_types.MyException
    _fbthrift_inner__myStruct : MyStruct | None
    _fbthrift_inner__myUnion : MyUnion | None


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.MyException(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> MyException:
        instance = super().__new__(cls)
        instance._fbthrift_inner__myStruct = None
        instance._fbthrift_inner__myUnion = None
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.MyException) -> MyException:
        inst = MyException.__new__(MyException)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> MyException:
        return self

    def _to_python(self) -> _fbthrift_python_types.MyException:
        return self._fbthrift__inner


    @property
    def MyIntField(self) -> int:
        return self._fbthrift__inner.MyIntField

    @property
    def MyStringField(self) -> str:
        return self._fbthrift__inner.MyStringField

    @property
    def myStruct(self) -> MyStruct:
        if self._fbthrift_inner__myStruct is None:
            __python_val = self._fbthrift__inner.myStruct
            self._fbthrift_inner__myStruct = MyStruct.from_python(__python_val)

        return self._fbthrift_inner__myStruct

    @property
    def myUnion(self) -> MyUnion:
        if self._fbthrift_inner__myUnion is None:
            __python_val = self._fbthrift__inner.myUnion
            self._fbthrift_inner__myUnion = MyUnion.from_python(__python_val)

        return self._fbthrift_inner__myUnion




__all__.append("MyException")

class MyExceptionWithMessage(thrift.py3.types.Struct):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
        "_fbthrift_inner__myStruct",
        "_fbthrift_inner__myUnion",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.MyExceptionWithMessage
    _FBTHRIFT__FIELD_NAMES = (
        "MyIntField",
        "MyStringField",
        "myStruct",
        "myUnion",
    )
    _fbthrift__inner : _fbthrift_python_types.MyExceptionWithMessage
    _fbthrift_inner__myStruct : MyStruct | None
    _fbthrift_inner__myUnion : MyUnion | None


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.MyExceptionWithMessage(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> MyExceptionWithMessage:
        instance = super().__new__(cls)
        instance._fbthrift_inner__myStruct = None
        instance._fbthrift_inner__myUnion = None
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.MyExceptionWithMessage) -> MyExceptionWithMessage:
        inst = MyExceptionWithMessage.__new__(MyExceptionWithMessage)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> MyExceptionWithMessage:
        return self

    def _to_python(self) -> _fbthrift_python_types.MyExceptionWithMessage:
        return self._fbthrift__inner


    @property
    def MyIntField(self) -> int:
        return self._fbthrift__inner.MyIntField

    @property
    def MyStringField(self) -> str:
        return self._fbthrift__inner.MyStringField

    @property
    def myStruct(self) -> MyStruct:
        if self._fbthrift_inner__myStruct is None:
            __python_val = self._fbthrift__inner.myStruct
            self._fbthrift_inner__myStruct = MyStruct.from_python(__python_val)

        return self._fbthrift_inner__myStruct

    @property
    def myUnion(self) -> MyUnion:
        if self._fbthrift_inner__myUnion is None:
            __python_val = self._fbthrift__inner.myUnion
            self._fbthrift_inner__myUnion = MyUnion.from_python(__python_val)

        return self._fbthrift_inner__myUnion




__all__.append("MyExceptionWithMessage")

class ReservedKeyword(thrift.py3.types.Struct):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.ReservedKeyword
    _FBTHRIFT__FIELD_NAMES = (
        "reserved_field",
    )
    _fbthrift__inner : _fbthrift_python_types.ReservedKeyword


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.ReservedKeyword(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> ReservedKeyword:
        instance = super().__new__(cls)
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.ReservedKeyword) -> ReservedKeyword:
        inst = ReservedKeyword.__new__(ReservedKeyword)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> ReservedKeyword:
        return self

    def _to_python(self) -> _fbthrift_python_types.ReservedKeyword:
        return self._fbthrift__inner


    @property
    def reserved_field(self) -> int:
        return self._fbthrift__inner.reserved_field




__all__.append("ReservedKeyword")

class UnionToBeRenamed(thrift.py3.types.Union):
    __module__ = _fbthrift__module_name__
    __slots__ = (
        "_fbthrift__inner",
        "_fbthrift_inner__type",
        "_fbthrift_inner__value",
    )
    _FBTHRIFT__PYTHON_CLASS = _fbthrift_python_types.UnionToBeRenamed
    _FBTHRIFT__FIELD_NAMES = (
        "reserved_field",
    )
    Type = _fbthrift_python_types.UnionToBeRenamed.Type
    _fbthrift__inner : _fbthrift_python_types.UnionToBeRenamed
    _fbthrift_inner__type: Type
    _fbthrift_inner__value: int | None


    def __init__(self, *args, **kwargs) -> None:
        self._fbthrift__inner = _fbthrift_python_types.UnionToBeRenamed(*args, **kwargs)

    def __new__(cls, *args, **kwargs) -> UnionToBeRenamed:
        instance = super().__new__(cls)
        instance._fbthrift_inner__type = None
        instance._fbthrift_inner__value = None
        return instance

    @staticmethod
    def from_python(thrift_python_inner: _fbthrift_python_types.UnionToBeRenamed) -> UnionToBeRenamed:
        inst = UnionToBeRenamed.__new__(UnionToBeRenamed)
        inst._fbthrift__inner = thrift_python_inner
        return inst

    def _to_py3(self) -> UnionToBeRenamed:
        return self

    def _to_python(self) -> _fbthrift_python_types.UnionToBeRenamed:
        return self._fbthrift__inner


    @property
    def reserved_field(self) -> int:
        return self._fbthrift__inner.reserved_field

    @property
    def type(self) -> _fbthrift_python_types.UnionToBeRenamed.Type:
        return self._fbthrift__inner.type

    @property
    def value(self) -> int | None:
        match self._fbthrift__inner.type:
            case _:
                return self._fbthrift__inner.value




__all__.append("UnionToBeRenamed")


### Constants
FLAG = True
OFFSET = -10
COUNT = 200
MASK = 16388846
E = 2.718281828459
DATE = "June 28, 2017"
AList = List__i32((2, 3, 5, 7, ))
ASet = Set__string(("foo", "bar", "baz", ))
AMap = Map__string_List__i32( { "foo": List__i32((1, 2, 3, 4, )), "bar": List__i32((10, 32, 54, )) })
