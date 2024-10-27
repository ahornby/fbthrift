#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/interactions/src/module.thrift
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
import thrift.py3.stream
import test.fixtures.another_interactions.shared.types as _test_fixtures_another_interactions_shared_types


class CustomException(thrift.py3.exceptions.GeneratedError, _typing.Hashable):
    class __fbthrift_IsSet:
        message: bool
        pass

    message: _typing.Final[str] = ...

    def __init__(
        self, *,
        message: _typing.Optional[str]=None
    ) -> None: ...

    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'CustomException') -> bool: ...
    def __gt__(self, other: 'CustomException') -> bool: ...
    def __le__(self, other: 'CustomException') -> bool: ...
    def __ge__(self, other: 'CustomException') -> bool: ...

    def _to_python(self) -> "test.fixtures.interactions.module.thrift_types.CustomException": ...   # type: ignore
    def _to_py3(self) -> CustomException: ...
    def _to_py_deprecated(self) -> "test.fixtures.interactions.ttypes.CustomException": ...   # type: ignore


class ClientBufferedStream__bool(thrift.py3.stream.ClientBufferedStream[bool]):
    def __aiter__(self) -> _typing.AsyncIterator[bool]: ...
    async def __anext__(self) -> bool: ...

class ServerStream__bool(thrift.py3.stream.ServerStream[bool]):
    pass

class ServerPublisher_cbool(thrift.py3.stream.ServerPublisher):
    def complete(self) -> None: ...
    def send(self, item: pbool) -> None: ...

class ClientBufferedStream__i32(thrift.py3.stream.ClientBufferedStream[int]):
    def __aiter__(self) -> _typing.AsyncIterator[int]: ...
    async def __anext__(self) -> int: ...

class ServerStream__i32(thrift.py3.stream.ServerStream[int]):
    pass

class ServerPublisher_cint32_t(thrift.py3.stream.ServerPublisher):
    def complete(self) -> None: ...
    def send(self, item: cint32_t) -> None: ...

class ResponseAndClientBufferedStream__i32_i32(thrift.py3.stream.ResponseAndClientBufferedStream[int, int]):
    def __iter__(self) -> _typing.Tuple[
        int,
        ClientBufferedStream__i32,
    ]: ...

class ResponseAndServerStream__i32_i32(thrift.py3.stream.ResponseAndServerStream[int, int]):
    pass

