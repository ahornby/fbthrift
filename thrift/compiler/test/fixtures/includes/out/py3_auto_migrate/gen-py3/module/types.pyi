#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/includes/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum as _python_std_enum
import folly.iobuf as _fbthrift_iobuf
import thrift.py3.types
import thrift.python.types
import thrift.py3.exceptions
import typing as _typing

import sys
import itertools
import includes.types as _includes_types


class MyStruct(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        MyIncludedField: bool
        MyOtherIncludedField: bool
        MyIncludedInt: bool
        pass

    MyIncludedField: _typing.Final[_includes_types.Included] = ...
    MyOtherIncludedField: _typing.Final[_includes_types.Included] = ...
    MyIncludedInt: _typing.Final[int] = ...

    def __init__(
        self, *,
        MyIncludedField: _typing.Optional[_includes_types.Included]=None,
        MyOtherIncludedField: _typing.Optional[_includes_types.Included]=None,
        MyIncludedInt: _typing.Optional[int]=None
    ) -> None: ...

    def __call__(
        self, *,
        MyIncludedField: _typing.Union[_includes_types.Included, None]=None,
        MyOtherIncludedField: _typing.Union[_includes_types.Included, None]=None,
        MyIncludedInt: _typing.Union[int, None]=None
    ) -> MyStruct: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['MyStruct'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'MyStruct') -> bool: ...
    def __gt__(self, other: 'MyStruct') -> bool: ...
    def __le__(self, other: 'MyStruct') -> bool: ...
    def __ge__(self, other: 'MyStruct') -> bool: ...

    def _to_python(self) -> "module.thrift_types.MyStruct": ...   # type: ignore
    def _to_py3(self) -> MyStruct: ...
    def _to_py_deprecated(self) -> "module.ttypes.MyStruct": ...   # type: ignore

