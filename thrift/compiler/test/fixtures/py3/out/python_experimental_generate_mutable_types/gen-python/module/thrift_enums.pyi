#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import __static__

import fbcode.thrift.python.types as _fbthrift_python_types

class _fbthrift_compatible_with_AnEnum:
    pass


class AnEnum(_fbthrift_python_types.Enum, int, _fbthrift_compatible_with_AnEnum):
    NOTSET: AnEnum = ...
    ONE: AnEnum = ...
    TWO: AnEnum = ...
    THREE: AnEnum = ...
    FOUR: AnEnum = ...
    def _to_python(self) -> AnEnum: ...
    def _to_py3(self) -> "module.types.AnEnum": ...  # type: ignore
    def _to_py_deprecated(self) -> int: ...

class _fbthrift_compatible_with_AnEnumRenamed:
    pass


class AnEnumRenamed(_fbthrift_python_types.Enum, int, _fbthrift_compatible_with_AnEnumRenamed):
    name_: AnEnumRenamed = ...
    value_: AnEnumRenamed = ...
    renamed_: AnEnumRenamed = ...
    def _to_python(self) -> AnEnumRenamed: ...
    def _to_py3(self) -> "module.types.AnEnumRenamed": ...  # type: ignore
    def _to_py_deprecated(self) -> int: ...

class _fbthrift_compatible_with_Flags:
    pass


class Flags(_fbthrift_python_types.Flag, _fbthrift_compatible_with_Flags):
    flag_A: Flags = ...
    flag_B: Flags = ...
    flag_C: Flags = ...
    flag_D: Flags = ...
    def _to_python(self) -> Flags: ...
    def _to_py3(self) -> "module.types.Flags": ...  # type: ignore
    def _to_py_deprecated(self) -> int: ...
