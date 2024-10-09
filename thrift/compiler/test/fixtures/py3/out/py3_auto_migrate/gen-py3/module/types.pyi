#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/py3/src/module.thrift
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


class AnEnum(thrift.py3.types.Enum):
    NOTSET: AnEnum = ...
    ONE: AnEnum = ...
    TWO: AnEnum = ...
    THREE: AnEnum = ...
    FOUR: AnEnum = ...
    def _to_python(self) -> "module.thrift_types.AnEnum": ...   # type: ignore
    def _to_py3(self) -> AnEnum: ...
    def _to_py_deprecated(self) -> int: ...


class AnEnumRenamed(thrift.py3.types.Enum):
    name_: AnEnumRenamed = ...
    value_: AnEnumRenamed = ...
    renamed_: AnEnumRenamed = ...
    def _to_python(self) -> "module.thrift_types.AnEnumRenamed": ...   # type: ignore
    def _to_py3(self) -> AnEnumRenamed: ...
    def _to_py_deprecated(self) -> int: ...


class Flags(thrift.py3.types.Flag):
    flag_A: Flags = ...
    flag_B: Flags = ...
    flag_C: Flags = ...
    flag_D: Flags = ...
    def _to_python(self) -> "module.thrift_types.Flags": ...   # type: ignore
    def _to_py3(self) -> Flags: ...
    def _to_py_deprecated(self) -> int: ...


class SimpleException(thrift.py3.exceptions.GeneratedError, _typing.Hashable):
    class __fbthrift_IsSet:
        err_code: bool
        pass

    type: Final["SimpleException.Type"]
    err_code: Final[int] = ...

    def __init__(
        self, *,
        err_code: _typing.Optional[int]=None
    ) -> None: ...

    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'SimpleException') -> bool: ...
    def __gt__(self, other: 'SimpleException') -> bool: ...
    def __le__(self, other: 'SimpleException') -> bool: ...
    def __ge__(self, other: 'SimpleException') -> bool: ...

    def _to_python(self) -> "module.thrift_types.SimpleException": ...   # type: ignore
    def _to_py3(self) -> SimpleException: ...
    def _to_py_deprecated(self) -> "module.ttypes.SimpleException": ...   # type: ignore

class OptionalRefStruct(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        optional_blob: bool
        pass

    type: Final["OptionalRefStruct.Type"]
    optional_blob: Final[_typing.Optional[_fbthrift_iobuf.IOBuf]] = ...

    def __init__(
        self, *,
        optional_blob: _typing.Optional[_fbthrift_iobuf.IOBuf]=None
    ) -> None: ...

    def __call__(
        self, *,
        optional_blob: _typing.Union[_fbthrift_iobuf.IOBuf, None]=None
    ) -> OptionalRefStruct: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['OptionalRefStruct'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'OptionalRefStruct') -> bool: ...
    def __gt__(self, other: 'OptionalRefStruct') -> bool: ...
    def __le__(self, other: 'OptionalRefStruct') -> bool: ...
    def __ge__(self, other: 'OptionalRefStruct') -> bool: ...

    def _to_python(self) -> "module.thrift_types.OptionalRefStruct": ...   # type: ignore
    def _to_py3(self) -> OptionalRefStruct: ...
    def _to_py_deprecated(self) -> "module.ttypes.OptionalRefStruct": ...   # type: ignore

class SimpleStruct(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        is_on: bool
        tiny_int: bool
        small_int: bool
        nice_sized_int: bool
        big_int: bool
        real: bool
        smaller_real: bool
        pass

    type: Final["SimpleStruct.Type"]
    is_on: Final[bool] = ...
    tiny_int: Final[int] = ...
    small_int: Final[int] = ...
    nice_sized_int: Final[int] = ...
    big_int: Final[int] = ...
    real: Final[float] = ...
    smaller_real: Final[float] = ...

    def __init__(
        self, *,
        is_on: _typing.Optional[bool]=None,
        tiny_int: _typing.Optional[int]=None,
        small_int: _typing.Optional[int]=None,
        nice_sized_int: _typing.Optional[int]=None,
        big_int: _typing.Optional[int]=None,
        real: _typing.Optional[float]=None,
        smaller_real: _typing.Optional[float]=None
    ) -> None: ...

    def __call__(
        self, *,
        is_on: _typing.Union[bool, None]=None,
        tiny_int: _typing.Union[int, None]=None,
        small_int: _typing.Union[int, None]=None,
        nice_sized_int: _typing.Union[int, None]=None,
        big_int: _typing.Union[int, None]=None,
        real: _typing.Union[float, None]=None,
        smaller_real: _typing.Union[float, None]=None
    ) -> SimpleStruct: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['SimpleStruct'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'SimpleStruct') -> bool: ...
    def __gt__(self, other: 'SimpleStruct') -> bool: ...
    def __le__(self, other: 'SimpleStruct') -> bool: ...
    def __ge__(self, other: 'SimpleStruct') -> bool: ...

    def _to_python(self) -> "module.thrift_types.SimpleStruct": ...   # type: ignore
    def _to_py3(self) -> SimpleStruct: ...
    def _to_py_deprecated(self) -> "module.ttypes.SimpleStruct": ...   # type: ignore

class HiddenTypeFieldsStruct(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        pass

    type: Final["HiddenTypeFieldsStruct.Type"]

    def __init__(
        self, 
    ) -> None: ...

    def __call__(
        self, 
    ) -> HiddenTypeFieldsStruct: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['HiddenTypeFieldsStruct'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...

    def _to_python(self) -> "module.thrift_types.HiddenTypeFieldsStruct": ...   # type: ignore
    def _to_py3(self) -> HiddenTypeFieldsStruct: ...
    def _to_py_deprecated(self) -> "module.ttypes.HiddenTypeFieldsStruct": ...   # type: ignore

class ComplexStruct(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        structOne: bool
        structTwo: bool
        an_integer: bool
        name: bool
        an_enum: bool
        some_bytes: bool
        sender: bool
        cdef_: bool
        bytes_with_cpp_type: bool
        pass

    type: Final["ComplexStruct.Type"]
    structOne: Final[SimpleStruct] = ...
    structTwo: Final[SimpleStruct] = ...
    an_integer: Final[int] = ...
    name: Final[str] = ...
    an_enum: Final[AnEnum] = ...
    some_bytes: Final[bytes] = ...
    sender: Final[str] = ...
    cdef_: Final[str] = ...
    bytes_with_cpp_type: Final[bytes] = ...

    def __init__(
        self, *,
        structOne: _typing.Optional[SimpleStruct]=None,
        structTwo: _typing.Optional[SimpleStruct]=None,
        an_integer: _typing.Optional[int]=None,
        name: _typing.Optional[str]=None,
        an_enum: _typing.Optional[AnEnum]=None,
        some_bytes: _typing.Optional[bytes]=None,
        sender: _typing.Optional[str]=None,
        cdef_: _typing.Optional[str]=None,
        bytes_with_cpp_type: _typing.Optional[bytes]=None
    ) -> None: ...

    def __call__(
        self, *,
        structOne: _typing.Union[SimpleStruct, None]=None,
        structTwo: _typing.Union[SimpleStruct, None]=None,
        an_integer: _typing.Union[int, None]=None,
        name: _typing.Union[str, None]=None,
        an_enum: _typing.Union[AnEnum, None]=None,
        some_bytes: _typing.Union[bytes, None]=None,
        sender: _typing.Union[str, None]=None,
        cdef_: _typing.Union[str, None]=None,
        bytes_with_cpp_type: _typing.Union[bytes, None]=None
    ) -> ComplexStruct: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['ComplexStruct'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'ComplexStruct') -> bool: ...
    def __gt__(self, other: 'ComplexStruct') -> bool: ...
    def __le__(self, other: 'ComplexStruct') -> bool: ...
    def __ge__(self, other: 'ComplexStruct') -> bool: ...

    def _to_python(self) -> "module.thrift_types.ComplexStruct": ...   # type: ignore
    def _to_py3(self) -> ComplexStruct: ...
    def _to_py_deprecated(self) -> "module.ttypes.ComplexStruct": ...   # type: ignore

_BinaryUnionValueType = _typing.Union[None, _fbthrift_iobuf.IOBuf]

class BinaryUnion(thrift.py3.types.Union, _typing.Hashable):
    class __fbthrift_IsSet:
        iobuf_val: bool
        pass

    type: Final["BinaryUnion.Type"]
    iobuf_val: Final[_fbthrift_iobuf.IOBuf] = ...

    def __init__(
        self, *,
        iobuf_val: _typing.Optional[_fbthrift_iobuf.IOBuf]=None
    ) -> None: ...

    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'BinaryUnion') -> bool: ...
    def __gt__(self, other: 'BinaryUnion') -> bool: ...
    def __le__(self, other: 'BinaryUnion') -> bool: ...
    def __ge__(self, other: 'BinaryUnion') -> bool: ...

    class Type(_python_std_enum.Enum):
        EMPTY: BinaryUnion.Type = ...
        iobuf_val: BinaryUnion.Type = ...

    @staticmethod
    def fromValue(value: _BinaryUnionValueType) -> BinaryUnion: ...
    @__property__
    def value(self) -> _BinaryUnionValueType: ...
    @__property__
    def type(self) -> "BinaryUnion.Type": ...

    def _to_python(self) -> "module.thrift_types.BinaryUnion": ...   # type: ignore
    def _to_py3(self) -> BinaryUnion: ...
    def _to_py_deprecated(self) -> "module.ttypes.BinaryUnion": ...   # type: ignore

class BinaryUnionStruct(thrift.py3.types.Struct, _typing.Hashable):
    class __fbthrift_IsSet:
        u: bool
        pass

    type: Final["BinaryUnionStruct.Type"]
    u: Final[BinaryUnion] = ...

    def __init__(
        self, *,
        u: _typing.Optional[BinaryUnion]=None
    ) -> None: ...

    def __call__(
        self, *,
        u: _typing.Union[BinaryUnion, None]=None
    ) -> BinaryUnionStruct: ...

    def __reduce__(self) -> _typing.Tuple[_typing.Callable, _typing.Tuple[_typing.Type['BinaryUnionStruct'], bytes]]: ...
    def __hash__(self) -> int: ...
    def __str__(self) -> str: ...
    def __repr__(self) -> str: ...
    def __lt__(self, other: 'BinaryUnionStruct') -> bool: ...
    def __gt__(self, other: 'BinaryUnionStruct') -> bool: ...
    def __le__(self, other: 'BinaryUnionStruct') -> bool: ...
    def __ge__(self, other: 'BinaryUnionStruct') -> bool: ...

    def _to_python(self) -> "module.thrift_types.BinaryUnionStruct": ...   # type: ignore
    def _to_py3(self) -> BinaryUnionStruct: ...
    def _to_py_deprecated(self) -> "module.ttypes.BinaryUnionStruct": ...   # type: ignore

_List__i16T = _typing.TypeVar('_List__i16T', bound=_typing.Sequence[int])


class List__i16(_typing.Sequence[int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[int]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> int: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[int]: ...
    def __add__(self, other: _typing.Sequence[int]) -> 'List__i16': ...
    def __radd__(self, other: _List__i16T) -> _List__i16T: ...
    def __reversed__(self) -> _typing.Iterator[int]: ...
    def __iter__(self) -> _typing.Iterator[int]: ...


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


_List__i64T = _typing.TypeVar('_List__i64T', bound=_typing.Sequence[int])


class List__i64(_typing.Sequence[int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[int]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> int: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[int]: ...
    def __add__(self, other: _typing.Sequence[int]) -> 'List__i64': ...
    def __radd__(self, other: _List__i64T) -> _List__i64T: ...
    def __reversed__(self) -> _typing.Iterator[int]: ...
    def __iter__(self) -> _typing.Iterator[int]: ...


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
    def __radd__(self, other: _List__stringT) -> _List__stringT: ...
    def __reversed__(self) -> _typing.Iterator[str]: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


_List__SimpleStructT = _typing.TypeVar('_List__SimpleStructT', bound=_typing.Sequence[SimpleStruct])


class List__SimpleStruct(_typing.Sequence[SimpleStruct], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[SimpleStruct]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[SimpleStruct]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> SimpleStruct: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[SimpleStruct]: ...
    def __add__(self, other: _typing.Sequence[SimpleStruct]) -> 'List__SimpleStruct': ...
    def __radd__(self, other: _List__SimpleStructT) -> _List__SimpleStructT: ...
    def __reversed__(self) -> _typing.Iterator[SimpleStruct]: ...
    def __iter__(self) -> _typing.Iterator[SimpleStruct]: ...


class Set__i32(_typing.AbstractSet[int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.AbstractSet[int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.AbstractSet[int]: ...
    def __contains__(self, x: object) -> bool: ...
    def union(self, other: _typing.AbstractSet[int]) -> 'Set__i32': ...
    def intersection(self, other: _typing.AbstractSet[int]) -> 'Set__i32': ...
    def difference(self, other: _typing.AbstractSet[int]) -> 'Set__i32': ...
    def symmetric_difference(self, other: _typing.AbstractSet[int]) -> 'Set__i32': ...
    def issubset(self, other: _typing.AbstractSet[int]) -> bool: ...
    def issuperset(self, other: _typing.AbstractSet[int]) -> bool: ...
    def __iter__(self) -> _typing.Iterator[int]: ...


class Set__string(_typing.AbstractSet[str], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.AbstractSet[str]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.AbstractSet[str]: ...
    def __contains__(self, x: object) -> bool: ...
    def union(self, other: _typing.AbstractSet[str]) -> 'Set__string': ...
    def intersection(self, other: _typing.AbstractSet[str]) -> 'Set__string': ...
    def difference(self, other: _typing.AbstractSet[str]) -> 'Set__string': ...
    def symmetric_difference(self, other: _typing.AbstractSet[str]) -> 'Set__string': ...
    def issubset(self, other: _typing.AbstractSet[str]) -> bool: ...
    def issuperset(self, other: _typing.AbstractSet[str]) -> bool: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


class Map__string_string(_typing.Mapping[str, str], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[str, str]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[str, str]: ...
    def __getitem__(self, key: str) -> str: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


class Map__string_SimpleStruct(_typing.Mapping[str, SimpleStruct], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[str, SimpleStruct]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[str, SimpleStruct]: ...
    def __getitem__(self, key: str) -> SimpleStruct: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


class Map__string_i16(_typing.Mapping[str, int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[str, int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[str, int]: ...
    def __getitem__(self, key: str) -> int: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


_List__List__i32T = _typing.TypeVar('_List__List__i32T', bound=_typing.Sequence[_typing.Sequence[int]])


class List__List__i32(_typing.Sequence[_typing.Sequence[int]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[_typing.Sequence[int]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[_typing.Sequence[int]]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> _typing.Sequence[int]: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[_typing.Sequence[int]]: ...
    def __add__(self, other: _typing.Sequence[_typing.Sequence[int]]) -> 'List__List__i32': ...
    def __radd__(self, other: _List__List__i32T) -> _List__List__i32T: ...
    def __reversed__(self) -> _typing.Iterator[_typing.Sequence[int]]: ...
    def __iter__(self) -> _typing.Iterator[_typing.Sequence[int]]: ...


class Map__string_i32(_typing.Mapping[str, int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[str, int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[str, int]: ...
    def __getitem__(self, key: str) -> int: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


class Map__string_Map__string_i32(_typing.Mapping[str, _typing.Mapping[str, int]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[str, _typing.Mapping[str, int]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[str, _typing.Mapping[str, int]]: ...
    def __getitem__(self, key: str) -> _typing.Mapping[str, int]: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


_List__Set__stringT = _typing.TypeVar('_List__Set__stringT', bound=_typing.Sequence[_typing.AbstractSet[str]])


class List__Set__string(_typing.Sequence[_typing.AbstractSet[str]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[_typing.AbstractSet[str]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[_typing.AbstractSet[str]]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> _typing.AbstractSet[str]: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[_typing.AbstractSet[str]]: ...
    def __add__(self, other: _typing.Sequence[_typing.AbstractSet[str]]) -> 'List__Set__string': ...
    def __radd__(self, other: _List__Set__stringT) -> _List__Set__stringT: ...
    def __reversed__(self) -> _typing.Iterator[_typing.AbstractSet[str]]: ...
    def __iter__(self) -> _typing.Iterator[_typing.AbstractSet[str]]: ...


class Map__string_List__SimpleStruct(_typing.Mapping[str, _typing.Sequence[SimpleStruct]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[str, _typing.Sequence[SimpleStruct]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[str, _typing.Sequence[SimpleStruct]]: ...
    def __getitem__(self, key: str) -> _typing.Sequence[SimpleStruct]: ...
    def __iter__(self) -> _typing.Iterator[str]: ...


_List__List__stringT = _typing.TypeVar('_List__List__stringT', bound=_typing.Sequence[_typing.Sequence[str]])


class List__List__string(_typing.Sequence[_typing.Sequence[str]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[_typing.Sequence[str]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[_typing.Sequence[str]]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> _typing.Sequence[str]: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[_typing.Sequence[str]]: ...
    def __add__(self, other: _typing.Sequence[_typing.Sequence[str]]) -> 'List__List__string': ...
    def __radd__(self, other: _List__List__stringT) -> _List__List__stringT: ...
    def __reversed__(self) -> _typing.Iterator[_typing.Sequence[str]]: ...
    def __iter__(self) -> _typing.Iterator[_typing.Sequence[str]]: ...


_List__Set__i32T = _typing.TypeVar('_List__Set__i32T', bound=_typing.Sequence[_typing.AbstractSet[int]])


class List__Set__i32(_typing.Sequence[_typing.AbstractSet[int]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[_typing.AbstractSet[int]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[_typing.AbstractSet[int]]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> _typing.AbstractSet[int]: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[_typing.AbstractSet[int]]: ...
    def __add__(self, other: _typing.Sequence[_typing.AbstractSet[int]]) -> 'List__Set__i32': ...
    def __radd__(self, other: _List__Set__i32T) -> _List__Set__i32T: ...
    def __reversed__(self) -> _typing.Iterator[_typing.AbstractSet[int]]: ...
    def __iter__(self) -> _typing.Iterator[_typing.AbstractSet[int]]: ...


_List__Map__string_stringT = _typing.TypeVar('_List__Map__string_stringT', bound=_typing.Sequence[_typing.Mapping[str, str]])


class List__Map__string_string(_typing.Sequence[_typing.Mapping[str, str]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[_typing.Mapping[str, str]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[_typing.Mapping[str, str]]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> _typing.Mapping[str, str]: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[_typing.Mapping[str, str]]: ...
    def __add__(self, other: _typing.Sequence[_typing.Mapping[str, str]]) -> 'List__Map__string_string': ...
    def __radd__(self, other: _List__Map__string_stringT) -> _List__Map__string_stringT: ...
    def __reversed__(self) -> _typing.Iterator[_typing.Mapping[str, str]]: ...
    def __iter__(self) -> _typing.Iterator[_typing.Mapping[str, str]]: ...


_List__binaryT = _typing.TypeVar('_List__binaryT', bound=_typing.Sequence[bytes])


class List__binary(_typing.Sequence[bytes], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[bytes]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[bytes]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> bytes: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[bytes]: ...
    def __add__(self, other: _typing.Sequence[bytes]) -> 'List__binary': ...
    def __radd__(self, other: _List__binaryT) -> _List__binaryT: ...
    def __reversed__(self) -> _typing.Iterator[bytes]: ...
    def __iter__(self) -> _typing.Iterator[bytes]: ...


class Set__binary(_typing.AbstractSet[bytes], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.AbstractSet[bytes]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.AbstractSet[bytes]: ...
    def __contains__(self, x: object) -> bool: ...
    def union(self, other: _typing.AbstractSet[bytes]) -> 'Set__binary': ...
    def intersection(self, other: _typing.AbstractSet[bytes]) -> 'Set__binary': ...
    def difference(self, other: _typing.AbstractSet[bytes]) -> 'Set__binary': ...
    def symmetric_difference(self, other: _typing.AbstractSet[bytes]) -> 'Set__binary': ...
    def issubset(self, other: _typing.AbstractSet[bytes]) -> bool: ...
    def issuperset(self, other: _typing.AbstractSet[bytes]) -> bool: ...
    def __iter__(self) -> _typing.Iterator[bytes]: ...


_List__AnEnumT = _typing.TypeVar('_List__AnEnumT', bound=_typing.Sequence[AnEnum])


class List__AnEnum(_typing.Sequence[AnEnum], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[AnEnum]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[AnEnum]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> AnEnum: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[AnEnum]: ...
    def __add__(self, other: _typing.Sequence[AnEnum]) -> 'List__AnEnum': ...
    def __radd__(self, other: _List__AnEnumT) -> _List__AnEnumT: ...
    def __reversed__(self) -> _typing.Iterator[AnEnum]: ...
    def __iter__(self) -> _typing.Iterator[AnEnum]: ...


class Map__i32_double(_typing.Mapping[int, float], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[int, float]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[int, float]: ...
    def __getitem__(self, key: int) -> float: ...
    def __iter__(self) -> _typing.Iterator[int]: ...


_List__Map__i32_doubleT = _typing.TypeVar('_List__Map__i32_doubleT', bound=_typing.Sequence[_typing.Mapping[int, float]])


class List__Map__i32_double(_typing.Sequence[_typing.Mapping[int, float]], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Sequence[_typing.Mapping[int, float]]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Sequence[_typing.Mapping[int, float]]: ...
    @_typing.overload
    def __getitem__(self, i: int) -> _typing.Mapping[int, float]: ...
    @_typing.overload
    def __getitem__(self, s: slice) -> _typing.Sequence[_typing.Mapping[int, float]]: ...
    def __add__(self, other: _typing.Sequence[_typing.Mapping[int, float]]) -> 'List__Map__i32_double': ...
    def __radd__(self, other: _List__Map__i32_doubleT) -> _List__Map__i32_doubleT: ...
    def __reversed__(self) -> _typing.Iterator[_typing.Mapping[int, float]]: ...
    def __iter__(self) -> _typing.Iterator[_typing.Mapping[int, float]]: ...


class Map__AnEnumRenamed_i32(_typing.Mapping[AnEnumRenamed, int], _typing.Hashable):
    def __init__(self, items: _typing.Optional[_typing.Mapping[AnEnumRenamed, int]]=None) -> None: ...
    def __len__(self) -> int: ...
    def __hash__(self) -> int: ...
    def __copy__(self) -> _typing.Mapping[AnEnumRenamed, int]: ...
    def __getitem__(self, key: AnEnumRenamed) -> int: ...
    def __iter__(self) -> _typing.Iterator[AnEnumRenamed]: ...


A_BOOL: bool = ...
A_BYTE: int = ...
THE_ANSWER: int = ...
A_NUMBER: int = ...
A_BIG_NUMBER: int = ...
A_REAL_NUMBER: float = ...
A_FAKE_NUMBER: float = ...
A_WORD: str = ...
SOME_BYTES: bytes = ...
A_STRUCT: SimpleStruct = ...
EMPTY: SimpleStruct = ...
WORD_LIST: List__string = ...
SOME_MAP: List__Map__i32_double = ...
DIGITS: Set__i32 = ...
A_CONST_MAP: Map__string_SimpleStruct = ...
ANOTHER_CONST_MAP: Map__AnEnumRenamed_i32 = ...
IOBufPtr = _fbthrift_iobuf.IOBuf
IOBuf = _fbthrift_iobuf.IOBuf
foo_bar = bytes
