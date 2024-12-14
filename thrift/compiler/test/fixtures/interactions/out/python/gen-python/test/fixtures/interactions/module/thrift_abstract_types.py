

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
import test.fixtures.another_interactions.shared.thrift_abstract_types as _fbthrift__test__fixtures__another_interactions__shared__thrift_abstract_types

class CustomException(_fbthrift_python_abstract_types.AbstractGeneratedError):
    @_fbthrift_property
    def message(self) -> str: ...
    def _to_python(self) -> "test.fixtures.interactions.module.thrift_types.CustomException": ...  # type: ignore
    def _to_py3(self) -> "test.fixtures.interactions.module.types.CustomException": ...  # type: ignore
    def _to_py_deprecated(self) -> "test.fixtures.interactions.ttypes.CustomException": ...  # type: ignore
_fbthrift_CustomException = CustomException
class ShouldBeBoxed(_abc.ABC):
    @_fbthrift_property
    @_abc.abstractmethod
    def sessionId(self) -> str: ...
    @_abc.abstractmethod
    def _to_python(self) -> "test.fixtures.interactions.module.thrift_types.ShouldBeBoxed": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "test.fixtures.interactions.module.types.ShouldBeBoxed": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "test.fixtures.interactions.ttypes.ShouldBeBoxed": ...  # type: ignore
_fbthrift_ShouldBeBoxed = ShouldBeBoxed
