#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/req-opt/src/module.thrift
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
        myInteger: bool
        myString: bool
        myBools: bool
        myNumbers: bool
        pass

    type: Final["Foo.Type"]
    myInteger: Final[int] = ...
    myString: Final[_typing.Optional[str]] = ...
    myBools: Final[_typing.Sequence[bool]] = ...
    myNumbers: Final[_typing.Sequence[int]] = ...

    def __init__(
        self, *,
        myInteger: _typing.Optional[int]=None,
        myString: _typing.Optional[str]=None,
        myBools: _typing.Optional[_typing.Sequence[bool]]=None,
        myNumbers: _typing.Optional[_typing.Sequence[int]]=None
    ) -> None: ...

    def __call__(
        self, *,
        myInteger: _typing.Union[int, None]=None,
        myString: _typing.Union[str, None]=None,
        myBools: _typing.Union[_typing.Sequence[bool], None]=None,
        myNumbers: _typing.Union[_typing.Sequence[int], None]=None
    ) -> Foo: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['Foo'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'Foo') -> bool: ...
    def __gt__(self, other: 'Foo') -> bool: ...
    def __le__(self, other: 'Foo') -> bool: ...
    def __ge__(self, other: 'Foo') -> bool: ...

    def _to_python(self) -> "module.thrift_types.Foo": ...   # type: ignore
    def _to_py3(self) -> Foo: ...
    def _to_py_deprecated(self) -> "module.ttypes.Foo": ...   # type: ignore

_List__boolT = _typing.TypeVar('_List__boolT', bound=_typing.Sequence[bool])


class List__bool(_typing.Sequence[bool], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[bool]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[bool]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> bool: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[bool]: ...
    def __add__(self, other: _typing.Sequence[bool]) -> 'List__bool': ...
    def __radd__(self, other: _List__boolT) -> _List__boolT: ...
    def __reversed__(self) -> _typing.Iterator[bool]: ...
    def __iter__(self) -> _typing.Iterator[bool]: ...


_List__i32T = _typing.TypeVar('_List__i32T', bound=_typing.Sequence[int])


class List__i32(_typing.Sequence[int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[int]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> int: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[int]: ...
    def __add__(self, other: _typing.Sequence[int]) -> 'List__i32': ...
    def __radd__(self, other: _List__i32T) -> _List__i32T: ...
    def __reversed__(self) -> _typing.Iterator[int]: ...
    def __iter__(self) -> _typing.Iterator[int]: ...


