#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import test.fixtures.enums.module.thrift_abstract_types as _fbthrift_python_abstract_types
import thrift.python.types as _fbthrift_python_types
import thrift.python.exceptions as _fbthrift_python_exceptions

from test.fixtures.enums.module.thrift_enums import *

from test.fixtures.enums.module.thrift_enums import (
    Metasyntactic as _fbthrift_Metasyntactic,
    _fbthrift_compatible_with_Metasyntactic,
    MyEnum1 as _fbthrift_MyEnum1,
    _fbthrift_compatible_with_MyEnum1,
    MyEnum2 as _fbthrift_MyEnum2,
    _fbthrift_compatible_with_MyEnum2,
    MyEnum3 as _fbthrift_MyEnum3,
    _fbthrift_compatible_with_MyEnum3,
    MyEnum4 as _fbthrift_MyEnum4,
    _fbthrift_compatible_with_MyEnum4,
    MyBitmaskEnum1 as _fbthrift_MyBitmaskEnum1,
    _fbthrift_compatible_with_MyBitmaskEnum1,
    MyBitmaskEnum2 as _fbthrift_MyBitmaskEnum2,
    _fbthrift_compatible_with_MyBitmaskEnum2,
)


class _fbthrift_compatible_with_SomeStruct:
    pass


class SomeStruct(_fbthrift_python_types.Struct, _fbthrift_compatible_with_SomeStruct, _fbthrift_python_abstract_types.SomeStruct):
    reasonable: _typing.Final[_fbthrift_Metasyntactic] = ...
    fine: _typing.Final[_fbthrift_Metasyntactic] = ...
    questionable: _typing.Final[_fbthrift_Metasyntactic] = ...
    tags: _typing.Final[_typing.AbstractSet[int]] = ...
    def __init__(
        self, *,
        reasonable: _typing.Optional[_fbthrift_compatible_with_Metasyntactic]=...,
        fine: _typing.Optional[_fbthrift_compatible_with_Metasyntactic]=...,
        questionable: _typing.Optional[_fbthrift_compatible_with_Metasyntactic]=...,
        tags: _typing.Optional[_typing.AbstractSet[int]]=...
    ) -> None: ...

    def __call__(
        self, *,
        reasonable: _typing.Optional[_fbthrift_compatible_with_Metasyntactic]=...,
        fine: _typing.Optional[_fbthrift_compatible_with_Metasyntactic]=...,
        questionable: _typing.Optional[_fbthrift_compatible_with_Metasyntactic]=...,
        tags: _typing.Optional[_typing.AbstractSet[int]]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_Metasyntactic, _fbthrift_Metasyntactic, _fbthrift_Metasyntactic, _typing.AbstractSet[int]]]]: ...
    def _to_python(self) -> _typing.Self: ...
    def _to_mutable_python(self) -> "test.fixtures.enums.module.thrift_mutable_types.SomeStruct": ...  # type: ignore
    def _to_py3(self) -> "test.fixtures.enums.module.types.SomeStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.SomeStruct": ...  # type: ignore
_fbthrift_SomeStruct = SomeStruct

class _fbthrift_compatible_with_MyStruct:
    pass


class MyStruct(_fbthrift_python_types.Struct, _fbthrift_compatible_with_MyStruct, _fbthrift_python_abstract_types.MyStruct):
    me2_3: _typing.Final[_fbthrift_MyEnum2] = ...
    me3_n3: _typing.Final[_fbthrift_MyEnum3] = ...
    me1_t1: _typing.Final[_fbthrift_MyEnum1] = ...
    me1_t2: _typing.Final[_fbthrift_MyEnum1] = ...
    def __init__(
        self, *,
        me2_3: _typing.Optional[_fbthrift_compatible_with_MyEnum2]=...,
        me3_n3: _typing.Optional[_fbthrift_compatible_with_MyEnum3]=...,
        me1_t1: _typing.Optional[_fbthrift_compatible_with_MyEnum1]=...,
        me1_t2: _typing.Optional[_fbthrift_compatible_with_MyEnum1]=...
    ) -> None: ...

    def __call__(
        self, *,
        me2_3: _typing.Optional[_fbthrift_compatible_with_MyEnum2]=...,
        me3_n3: _typing.Optional[_fbthrift_compatible_with_MyEnum3]=...,
        me1_t1: _typing.Optional[_fbthrift_compatible_with_MyEnum1]=...,
        me1_t2: _typing.Optional[_fbthrift_compatible_with_MyEnum1]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_MyEnum2, _fbthrift_MyEnum3, _fbthrift_MyEnum1, _fbthrift_MyEnum1]]]: ...
    def _to_python(self) -> _typing.Self: ...
    def _to_mutable_python(self) -> "test.fixtures.enums.module.thrift_mutable_types.MyStruct": ...  # type: ignore
    def _to_py3(self) -> "test.fixtures.enums.module.types.MyStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.MyStruct": ...  # type: ignore
_fbthrift_MyStruct = MyStruct
