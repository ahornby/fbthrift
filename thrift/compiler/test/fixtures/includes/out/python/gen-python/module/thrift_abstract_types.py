

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
import includes.thrift_abstract_types as _fbthrift__includes__thrift_abstract_types

class MyStruct(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def MyIncludedField(self) -> _fbthrift__includes__thrift_abstract_types.Included: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def MyOtherIncludedField(self) -> _fbthrift__includes__thrift_abstract_types.Included: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def MyIncludedInt(self) -> int: ...
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.MyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.MyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.MyStruct": ...  # type: ignore
_fbthrift_MyStruct = MyStruct
