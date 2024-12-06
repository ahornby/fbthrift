#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations


# EXPERIMENTAL - DO NOT USE !!!
# See `experimental_generate_mutable_types` documentation in thrift compiler

#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import typing as _typing

import module.thrift_mutable_types as _fbthrift_current_module
import enum

import folly.iobuf as _fbthrift_iobuf
import thrift.python.types as _fbthrift_python_types
import thrift.python.mutable_types as _fbthrift_python_mutable_types
import thrift.python.mutable_exceptions as _fbthrift_python_mutable_exceptions
import thrift.python.mutable_containers as _fbthrift_python_mutable_containers

from module.thrift_enums import *


class _fbthrift_compatible_with_ComplexUnion:
    pass


class ComplexUnion(_fbthrift_python_mutable_types.MutableUnion, _fbthrift_compatible_with_ComplexUnion):

    @property
    def intValue(self) -> int: ...
    @intValue.setter
    def intValue(self, value: int): ...


    @property
    def intListValue(self) -> _fbthrift_python_mutable_containers.MutableList[int]: ...
    @intListValue.setter
    def intListValue(self, value: _fbthrift_python_mutable_containers.MutableList[int] | _fbthrift_python_mutable_types._ThriftListWrapper): ...


    @property
    def stringListValue(self) -> _fbthrift_python_mutable_containers.MutableList[str]: ...
    @stringListValue.setter
    def stringListValue(self, value: _fbthrift_python_mutable_containers.MutableList[str] | _fbthrift_python_mutable_types._ThriftListWrapper): ...


    @property
    def stringValue(self) -> str: ...
    @stringValue.setter
    def stringValue(self, value: str): ...


    @property
    def typedefValue(self) -> _fbthrift_python_mutable_containers.MutableMap[int, str]: ...
    @typedefValue.setter
    def typedefValue(self, value: _fbthrift_python_mutable_containers.MutableMap[int, str] | _fbthrift_python_mutable_types._ThriftMapWrapper): ...


    @property
    def stringRef(self) -> str: ...
    @stringRef.setter
    def stringRef(self, value: str): ...

    def __init__(
        self, *,
        intValue: _typing.Optional[int]=...,
        intListValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[int] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        stringListValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[str] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        stringValue: _typing.Optional[str]=...,
        typedefValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[int, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        stringRef: _typing.Optional[str]=...
    ) -> None: ...



    class FbThriftUnionFieldEnum(enum.Enum):
        EMPTY: ComplexUnion.FbThriftUnionFieldEnum = ...
        intValue: ComplexUnion.FbThriftUnionFieldEnum = ...
        intListValue: ComplexUnion.FbThriftUnionFieldEnum = ...
        stringListValue: ComplexUnion.FbThriftUnionFieldEnum = ...
        stringValue: ComplexUnion.FbThriftUnionFieldEnum = ...
        typedefValue: ComplexUnion.FbThriftUnionFieldEnum = ...
        stringRef: ComplexUnion.FbThriftUnionFieldEnum = ...

    fbthrift_current_value: _typing.Final[_typing.Union[None, int, _fbthrift_python_mutable_containers.MutableList[int], _fbthrift_python_mutable_containers.MutableList[str], str, _fbthrift_python_mutable_containers.MutableMap[int, str], str]]
    fbthrift_current_field: _typing.Final[FbThriftUnionFieldEnum]
    def get_type(self) -> FbThriftUnionFieldEnum: ...
    def _to_python(self) -> "module.thrift_types.ComplexUnion": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.ComplexUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.ComplexUnion": ...  # type: ignore


class _fbthrift_compatible_with_ListUnion:
    pass


class ListUnion(_fbthrift_python_mutable_types.MutableUnion, _fbthrift_compatible_with_ListUnion):

    @property
    def intListValue(self) -> _fbthrift_python_mutable_containers.MutableList[int]: ...
    @intListValue.setter
    def intListValue(self, value: _fbthrift_python_mutable_containers.MutableList[int] | _fbthrift_python_mutable_types._ThriftListWrapper): ...


    @property
    def stringListValue(self) -> _fbthrift_python_mutable_containers.MutableList[str]: ...
    @stringListValue.setter
    def stringListValue(self, value: _fbthrift_python_mutable_containers.MutableList[str] | _fbthrift_python_mutable_types._ThriftListWrapper): ...

    def __init__(
        self, *,
        intListValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[int] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        stringListValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[str] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> None: ...



    class FbThriftUnionFieldEnum(enum.Enum):
        EMPTY: ListUnion.FbThriftUnionFieldEnum = ...
        intListValue: ListUnion.FbThriftUnionFieldEnum = ...
        stringListValue: ListUnion.FbThriftUnionFieldEnum = ...

    fbthrift_current_value: _typing.Final[_typing.Union[None, _fbthrift_python_mutable_containers.MutableList[int], _fbthrift_python_mutable_containers.MutableList[str]]]
    fbthrift_current_field: _typing.Final[FbThriftUnionFieldEnum]
    def get_type(self) -> FbThriftUnionFieldEnum: ...
    def _to_python(self) -> "module.thrift_types.ListUnion": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.ListUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.ListUnion": ...  # type: ignore


class _fbthrift_compatible_with_DataUnion:
    pass


class DataUnion(_fbthrift_python_mutable_types.MutableUnion, _fbthrift_compatible_with_DataUnion):

    @property
    def binaryData(self) -> bytes: ...
    @binaryData.setter
    def binaryData(self, value: bytes): ...


    @property
    def stringData(self) -> str: ...
    @stringData.setter
    def stringData(self, value: str): ...

    def __init__(
        self, *,
        binaryData: _typing.Optional[bytes]=...,
        stringData: _typing.Optional[str]=...
    ) -> None: ...



    class FbThriftUnionFieldEnum(enum.Enum):
        EMPTY: DataUnion.FbThriftUnionFieldEnum = ...
        binaryData: DataUnion.FbThriftUnionFieldEnum = ...
        stringData: DataUnion.FbThriftUnionFieldEnum = ...

    fbthrift_current_value: _typing.Final[_typing.Union[None, bytes, str]]
    fbthrift_current_field: _typing.Final[FbThriftUnionFieldEnum]
    def get_type(self) -> FbThriftUnionFieldEnum: ...
    def _to_python(self) -> "module.thrift_types.DataUnion": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.DataUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.DataUnion": ...  # type: ignore


class _fbthrift_compatible_with_Val:
    pass


class Val(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_Val):

    @property
    def strVal(self) -> str: ...
    @strVal.setter
    def strVal(self, value: str): ...


    @property
    def intVal(self) -> int: ...
    @intVal.setter
    def intVal(self, value: int): ...


    @property
    def typedefValue(self) -> _fbthrift_python_mutable_containers.MutableMap[int, str]: ...
    @typedefValue.setter
    def typedefValue(self, value: _fbthrift_python_mutable_containers.MutableMap[int, str] | _fbthrift_python_mutable_types._ThriftMapWrapper): ...

    def __init__(
        self, *,
        strVal: _typing.Optional[str]=...,
        intVal: _typing.Optional[int]=...,
        typedefValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[int, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        strVal: _typing.Optional[str]=...,
        intVal: _typing.Optional[int]=...,
        typedefValue: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[int, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str, int, _fbthrift_python_mutable_containers.MutableMap[int, str]]]]: ...
    def _to_python(self) -> "module.thrift_types.Val": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.Val": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.Val": ...  # type: ignore


class _fbthrift_compatible_with_ValUnion:
    pass


class ValUnion(_fbthrift_python_mutable_types.MutableUnion, _fbthrift_compatible_with_ValUnion):

    @property
    def v1(self) -> _fbthrift_current_module.Val: ...
    @v1.setter
    def v1(self, value: _fbthrift_current_module.Val): ...


    @property
    def v2(self) -> _fbthrift_current_module.Val: ...
    @v2.setter
    def v2(self, value: _fbthrift_current_module.Val): ...

    def __init__(
        self, *,
        v1: _typing.Optional[_fbthrift_compatible_with_Val]=...,
        v2: _typing.Optional[_fbthrift_compatible_with_Val]=...
    ) -> None: ...



    class FbThriftUnionFieldEnum(enum.Enum):
        EMPTY: ValUnion.FbThriftUnionFieldEnum = ...
        v1: ValUnion.FbThriftUnionFieldEnum = ...
        v2: ValUnion.FbThriftUnionFieldEnum = ...

    fbthrift_current_value: _typing.Final[_typing.Union[None, _fbthrift_current_module.Val, _fbthrift_current_module.Val]]
    fbthrift_current_field: _typing.Final[FbThriftUnionFieldEnum]
    def get_type(self) -> FbThriftUnionFieldEnum: ...
    def _to_python(self) -> "module.thrift_types.ValUnion": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.ValUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.ValUnion": ...  # type: ignore


class _fbthrift_compatible_with_VirtualComplexUnion:
    pass


class VirtualComplexUnion(_fbthrift_python_mutable_types.MutableUnion, _fbthrift_compatible_with_VirtualComplexUnion):

    @property
    def thingOne(self) -> str: ...
    @thingOne.setter
    def thingOne(self, value: str): ...


    @property
    def thingTwo(self) -> str: ...
    @thingTwo.setter
    def thingTwo(self, value: str): ...

    def __init__(
        self, *,
        thingOne: _typing.Optional[str]=...,
        thingTwo: _typing.Optional[str]=...
    ) -> None: ...



    class FbThriftUnionFieldEnum(enum.Enum):
        EMPTY: VirtualComplexUnion.FbThriftUnionFieldEnum = ...
        thingOne: VirtualComplexUnion.FbThriftUnionFieldEnum = ...
        thingTwo: VirtualComplexUnion.FbThriftUnionFieldEnum = ...

    fbthrift_current_value: _typing.Final[_typing.Union[None, str, str]]
    fbthrift_current_field: _typing.Final[FbThriftUnionFieldEnum]
    def get_type(self) -> FbThriftUnionFieldEnum: ...
    def _to_python(self) -> "module.thrift_types.VirtualComplexUnion": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.VirtualComplexUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.VirtualComplexUnion": ...  # type: ignore


class _fbthrift_compatible_with_NonCopyableStruct:
    pass


class NonCopyableStruct(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_NonCopyableStruct):

    @property
    def num(self) -> int: ...
    @num.setter
    def num(self, value: int): ...

    def __init__(
        self, *,
        num: _typing.Optional[int]=...
    ) -> None: ...

    def __call__(
        self, *,
        num: _typing.Optional[int]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[int]]]: ...
    def _to_python(self) -> "module.thrift_types.NonCopyableStruct": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.NonCopyableStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.NonCopyableStruct": ...  # type: ignore


class _fbthrift_compatible_with_NonCopyableUnion:
    pass


class NonCopyableUnion(_fbthrift_python_mutable_types.MutableUnion, _fbthrift_compatible_with_NonCopyableUnion):

    @property
    def s(self) -> _fbthrift_current_module.NonCopyableStruct: ...
    @s.setter
    def s(self, value: _fbthrift_current_module.NonCopyableStruct): ...

    def __init__(
        self, *,
        s: _typing.Optional[_fbthrift_compatible_with_NonCopyableStruct]=...
    ) -> None: ...



    class FbThriftUnionFieldEnum(enum.Enum):
        EMPTY: NonCopyableUnion.FbThriftUnionFieldEnum = ...
        s: NonCopyableUnion.FbThriftUnionFieldEnum = ...

    fbthrift_current_value: _typing.Final[_typing.Union[None, _fbthrift_current_module.NonCopyableStruct]]
    fbthrift_current_field: _typing.Final[FbThriftUnionFieldEnum]
    def get_type(self) -> FbThriftUnionFieldEnum: ...
    def _to_python(self) -> "module.thrift_types.NonCopyableUnion": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "module.types.NonCopyableUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.NonCopyableUnion": ...  # type: ignore

containerTypedef = _typing.Dict[int, str]
