#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import module.thrift_abstract_types as _fbthrift_python_abstract_types
import thrift.python.types as _fbthrift_python_types
import thrift.python.exceptions as _fbthrift_python_exceptions

import includes.thrift_types as _fbthrift__includes__thrift_types


class _fbthrift_compatible_with_MyStruct:
    pass


class MyStruct(_fbthrift_python_types.Struct, _fbthrift_compatible_with_MyStruct, _fbthrift_python_abstract_types.MyStruct):
    MyIncludedField: _typing.Final[_fbthrift__includes__thrift_types.Included] = ...
    MyOtherIncludedField: _typing.Final[_fbthrift__includes__thrift_types.Included] = ...
    MyIncludedInt: _typing.Final[int] = ...
    def __init__(
        self, *,
        MyIncludedField: _typing.Optional[_fbthrift__includes__thrift_types._fbthrift_compatible_with_Included]=...,
        MyOtherIncludedField: _typing.Optional[_fbthrift__includes__thrift_types._fbthrift_compatible_with_Included]=...,
        MyIncludedInt: _typing.Optional[int]=...
    ) -> None: ...

    def __call__(
        self, *,
        MyIncludedField: _typing.Optional[_fbthrift__includes__thrift_types._fbthrift_compatible_with_Included]=...,
        MyOtherIncludedField: _typing.Optional[_fbthrift__includes__thrift_types._fbthrift_compatible_with_Included]=...,
        MyIncludedInt: _typing.Optional[int]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift__includes__thrift_types.Included, _fbthrift__includes__thrift_types.Included, int]]]: ...
    def _to_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.MyStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.MyStruct": ...  # type: ignore
_fbthrift_MyStruct = MyStruct
