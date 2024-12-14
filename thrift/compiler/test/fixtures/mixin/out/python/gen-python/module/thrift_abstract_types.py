

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

class Mixin1(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def field1(self) -> str: ...
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.Mixin1": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.Mixin1": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.Mixin1": ...  # type: ignore
_fbthrift_Mixin1 = Mixin1
class Mixin2(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def m1(self) -> _fbthrift_Mixin1: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def field2(self) -> _typing.Optional[str]: ...
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.Mixin2": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.Mixin2": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.Mixin2": ...  # type: ignore
_fbthrift_Mixin2 = Mixin2
class Mixin3Base(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def field3(self) -> str: ...
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.Mixin3Base": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.Mixin3Base": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.Mixin3Base": ...  # type: ignore
_fbthrift_Mixin3Base = Mixin3Base
class Foo(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def field4(self) -> str: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def m2(self) -> _fbthrift_Mixin2: ...
    @_fbthrift_property
    @_abc.abstractmethod
    def m3(self) -> _fbthrift_Mixin3Base: ...
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.Foo": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.Foo": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.Foo": ...  # type: ignore
_fbthrift_Foo = Foo

Mixin3 = _fbthrift_Mixin3Base
