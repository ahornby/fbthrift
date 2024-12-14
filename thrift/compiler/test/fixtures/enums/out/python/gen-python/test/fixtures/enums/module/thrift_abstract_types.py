

#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import abc as _abc
import typing as _typing

_fbthrift_property = property


import folly.iobuf as _fbthrift_iobuf
import thrift.python.abstract_types as _fbthrift_python_abstract_types

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

class SomeStruct(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def reasonable(self) -> _fbthrift_Metasyntactic: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def fine(self) -> _fbthrift_Metasyntactic: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def questionable(self) -> _fbthrift_Metasyntactic: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def tags(self) -> _typing.AbstractSet[int]: ...
    @_abc.abstractmethod
    def _to_python(self) -> "test.fixtures.enums.module.thrift_types.SomeStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "test.fixtures.enums.module.types.SomeStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.SomeStruct": ...  # type: ignore
_fbthrift_SomeStruct = SomeStruct
class MyStruct(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def me2_3(self) -> _fbthrift_MyEnum2: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def me3_n3(self) -> _fbthrift_MyEnum3: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def me1_t1(self) -> _fbthrift_MyEnum1: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def me1_t2(self) -> _fbthrift_MyEnum1: ...
    @_abc.abstractmethod
    def _to_python(self) -> "test.fixtures.enums.module.thrift_types.MyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "test.fixtures.enums.module.types.MyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.MyStruct": ...  # type: ignore
_fbthrift_MyStruct = MyStruct
