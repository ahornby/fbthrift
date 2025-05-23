#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/list/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum as _python_std_enum
import folly.iobuf as _fbthrift_iobuf
import thrift.py3.types
import thrift.python.types
import thrift.py3.exceptions
import typing as _typing

import sys
import itertools
import module.thrift_types


_List__stringT = _typing.TypeVar('_List__stringT', bound=_typing.Sequence[str])


class List__string(_typing.Sequence[str], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[str]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[str]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> str: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[str]: ...
    def __add__(self, other: _typing.Sequence[str]) -> 'List__string': ...
    def __radd__(self, other: _typing.Sequence[str]) -> 'List__string': ...
    def __reversed__(self) -> _typing.Iterator[str]: ...
    #pyre-ignore[14]: no idea what pyre is on about
    def __iter__(self) -> _typing.Iterator[str]: ...


class Map__i64_List__string(_typing.Mapping[int, _typing.Sequence[str]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[int, _typing.Sequence[str]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[int, _typing.Sequence[str]]: ...
    def __getitem__(self, key: int) -> _typing.Sequence[str]: ...
    def __iter__(self) -> _typing.Iterator[int]: ...


TEST_MAP: Map__i64_List__string = ...
