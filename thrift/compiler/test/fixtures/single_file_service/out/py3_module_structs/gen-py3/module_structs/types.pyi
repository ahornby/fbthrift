#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/single_file_service/src/module_structs.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum as _python_std_enum
import folly.iobuf as _fbthrift_iobuf
import thrift.py3.types
import thrift.py3.exceptions
import typing as _typing
from typing_extensions import Final

import sys
import itertools


__property__ = property


class Foo(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        value: bool
        pass

    type: Final["Foo.Type"]
    value: Final[_typing.Optional[int]] = ...

    def __init__(
        self, *,
        value: _typing.Optional[int]=None
    ) -> None: ...

    def __call__(
        self, *,
        value: _typing.Union[int, None]=None
    ) -> Foo: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['Foo'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'Foo') -> bool: ...
    def __gt__(self, other: 'Foo') -> bool: ...
    def __le__(self, other: 'Foo') -> bool: ...
    def __ge__(self, other: 'Foo') -> bool: ...

    def _to_python(self) -> "module_structs.thrift_types.Foo": ...   # type: ignore
    def _to_py3(self) -> Foo: ...
    def _to_py_deprecated(self) -> "module_structs.ttypes.Foo": ...   # type: ignore

