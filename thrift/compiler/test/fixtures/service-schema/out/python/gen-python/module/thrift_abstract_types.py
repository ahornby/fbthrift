

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
import include.thrift_abstract_types as _fbthrift__include__thrift_abstract_types

from module.thrift_enums import *

from module.thrift_enums import (
    Result as _fbthrift_Result,
    _fbthrift_compatible_with_Result,
)

class CustomException(_fbthrift_python_abstract_types.AbstractGeneratedError):
    @_fbthrift_property
    def name(self) -> str: ...
    @_fbthrift_property
    def result(self) -> _fbthrift_Result: ...
    def _to_python(self) -> "module.thrift_types.CustomException": ...  # type: ignore
    def _to_py3(self) -> "module.types.CustomException": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.CustomException": ...  # type: ignore
_fbthrift_CustomException = CustomException
