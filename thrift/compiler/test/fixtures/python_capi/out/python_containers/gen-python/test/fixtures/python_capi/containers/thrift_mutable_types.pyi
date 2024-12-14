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

import folly.iobuf as _fbthrift_iobuf
import test.fixtures.python_capi.containers.thrift_abstract_types as _fbthrift_python_abstract_types
import thrift.python.types as _fbthrift_python_types
import thrift.python.mutable_types as _fbthrift_python_mutable_types
import thrift.python.mutable_exceptions as _fbthrift_python_mutable_exceptions
import thrift.python.mutable_containers as _fbthrift_python_mutable_containers


class _fbthrift_compatible_with_TemplateLists:
    pass


class TemplateLists(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_TemplateLists, _fbthrift_python_abstract_types.TemplateLists):

    @property
    def std_string(self) -> _typing.Optional[_fbthrift_python_mutable_containers.MutableList[str]]: ...
    @std_string.setter
    def std_string(self, value: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[str]] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...


    @property
    def deque_string(self) -> _fbthrift_python_mutable_containers.MutableList[bytes]: ...
    @deque_string.setter
    def deque_string(self, value: _fbthrift_python_mutable_containers.MutableList[bytes] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...


    @property
    def small_vector_iobuf(self) -> _fbthrift_python_mutable_containers.MutableList[_fbthrift_iobuf.IOBuf]: ...
    @small_vector_iobuf.setter
    def small_vector_iobuf(self, value: _fbthrift_python_mutable_containers.MutableList[_fbthrift_iobuf.IOBuf] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...


    @property
    def nested_small_vector(self) -> _fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]]: ...
    @nested_small_vector.setter
    def nested_small_vector(self, value: _fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...


    @property
    def small_vector_tensor(self) -> _fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]]]: ...
    @small_vector_tensor.setter
    def small_vector_tensor(self, value: _fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]]] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...

    def __init__(
        self, *,
        std_string: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[str] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        deque_string: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[bytes] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        small_vector_iobuf: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_iobuf.IOBuf] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        nested_small_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        small_vector_tensor: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]]] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        std_string: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[str] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        deque_string: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[bytes] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        small_vector_iobuf: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_iobuf.IOBuf] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        nested_small_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]] | _fbthrift_python_mutable_types._ThriftListWrapper]=...,
        small_vector_tensor: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]]] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_python_mutable_containers.MutableList[str], _fbthrift_python_mutable_containers.MutableList[bytes], _fbthrift_python_mutable_containers.MutableList[_fbthrift_iobuf.IOBuf], _fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]], _fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[_fbthrift_python_mutable_containers.MutableList[str]]]]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.TemplateLists": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.TemplateLists": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.TemplateLists": ...  # type: ignore
_fbthrift_TemplateLists = TemplateLists

class _fbthrift_compatible_with_TemplateSets:
    pass


class TemplateSets(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_TemplateSets, _fbthrift_python_abstract_types.TemplateSets):

    @property
    def std_set(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @std_set.setter
    def std_set(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...


    @property
    def std_unordered(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @std_unordered.setter
    def std_unordered(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...


    @property
    def folly_fast(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @folly_fast.setter
    def folly_fast(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...


    @property
    def folly_node(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @folly_node.setter
    def folly_node(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...


    @property
    def folly_value(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @folly_value.setter
    def folly_value(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...


    @property
    def folly_vector(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @folly_vector.setter
    def folly_vector(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...


    @property
    def folly_sorted_vector(self) -> _fbthrift_python_mutable_containers.MutableSet[str]: ...
    @folly_sorted_vector.setter
    def folly_sorted_vector(self, value: _fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper) -> None: ...

    def __init__(
        self, *,
        std_set: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        std_unordered: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_fast: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_node: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_value: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_sorted_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        std_set: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        std_unordered: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_fast: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_node: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_value: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...,
        folly_sorted_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableSet[str] | _fbthrift_python_mutable_types._ThriftSetWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_python_mutable_containers.MutableSet[str], _fbthrift_python_mutable_containers.MutableSet[str], _fbthrift_python_mutable_containers.MutableSet[str], _fbthrift_python_mutable_containers.MutableSet[str], _fbthrift_python_mutable_containers.MutableSet[str], _fbthrift_python_mutable_containers.MutableSet[str], _fbthrift_python_mutable_containers.MutableSet[str]]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.TemplateSets": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.TemplateSets": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.TemplateSets": ...  # type: ignore
_fbthrift_TemplateSets = TemplateSets

class _fbthrift_compatible_with_TemplateMaps:
    pass


class TemplateMaps(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_TemplateMaps, _fbthrift_python_abstract_types.TemplateMaps):

    @property
    def std_map(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @std_map.setter
    def std_map(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...


    @property
    def std_unordered(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @std_unordered.setter
    def std_unordered(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...


    @property
    def folly_fast(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @folly_fast.setter
    def folly_fast(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...


    @property
    def folly_node(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @folly_node.setter
    def folly_node(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...


    @property
    def folly_value(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @folly_value.setter
    def folly_value(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...


    @property
    def folly_vector(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @folly_vector.setter
    def folly_vector(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...


    @property
    def folly_sorted_vector(self) -> _fbthrift_python_mutable_containers.MutableMap[str, str]: ...
    @folly_sorted_vector.setter
    def folly_sorted_vector(self, value: _fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper) -> None: ...

    def __init__(
        self, *,
        std_map: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        std_unordered: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_fast: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_node: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_value: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_sorted_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        std_map: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        std_unordered: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_fast: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_node: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_value: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...,
        folly_sorted_vector: _typing.Optional[_fbthrift_python_mutable_containers.MutableMap[str, str] | _fbthrift_python_mutable_types._ThriftMapWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_python_mutable_containers.MutableMap[str, str], _fbthrift_python_mutable_containers.MutableMap[str, str], _fbthrift_python_mutable_containers.MutableMap[str, str], _fbthrift_python_mutable_containers.MutableMap[str, str], _fbthrift_python_mutable_containers.MutableMap[str, str], _fbthrift_python_mutable_containers.MutableMap[str, str], _fbthrift_python_mutable_containers.MutableMap[str, str]]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.TemplateMaps": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.TemplateMaps": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.TemplateMaps": ...  # type: ignore
_fbthrift_TemplateMaps = TemplateMaps

class _fbthrift_compatible_with_TWrapped:
    pass


class TWrapped(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_TWrapped, _fbthrift_python_abstract_types.TWrapped):

    @property
    def fieldA(self) -> str: ...
    @fieldA.setter
    def fieldA(self, value: str) -> None: ...


    @property
    def fieldB(self) -> bytes: ...
    @fieldB.setter
    def fieldB(self, value: bytes) -> None: ...

    def __init__(
        self, *,
        fieldA: _typing.Optional[str]=...,
        fieldB: _typing.Optional[bytes]=...
    ) -> None: ...

    def __call__(
        self, *,
        fieldA: _typing.Optional[str]=...,
        fieldB: _typing.Optional[bytes]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str, bytes]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.TWrapped": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.TWrapped": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.TWrapped": ...  # type: ignore
_fbthrift_TWrapped = TWrapped

class _fbthrift_compatible_with_IndirectionA:
    pass


class IndirectionA(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_IndirectionA, _fbthrift_python_abstract_types.IndirectionA):

    @property
    def lst(self) -> _fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped]: ...
    @lst.setter
    def lst(self, value: _fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...

    def __init__(
        self, *,
        lst: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_compatible_with_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        lst: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_compatible_with_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped]]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.IndirectionA": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.IndirectionA": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.IndirectionA": ...  # type: ignore
_fbthrift_IndirectionA = IndirectionA

class _fbthrift_compatible_with_IndirectionB:
    pass


class IndirectionB(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_IndirectionB, _fbthrift_python_abstract_types.IndirectionB):

    @property
    def lst(self) -> _fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped]: ...
    @lst.setter
    def lst(self, value: _fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...

    def __init__(
        self, *,
        lst: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_compatible_with_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        lst: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_compatible_with_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped]]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.IndirectionB": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.IndirectionB": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.IndirectionB": ...  # type: ignore
_fbthrift_IndirectionB = IndirectionB

class _fbthrift_compatible_with_IndirectionC:
    pass


class IndirectionC(_fbthrift_python_mutable_types.MutableStruct, _fbthrift_compatible_with_IndirectionC, _fbthrift_python_abstract_types.IndirectionC):

    @property
    def lst(self) -> _fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped]: ...
    @lst.setter
    def lst(self, value: _fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper) -> None: ...

    def __init__(
        self, *,
        lst: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_compatible_with_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> None: ...

    def __call__(
        self, *,
        lst: _typing.Optional[_fbthrift_python_mutable_containers.MutableList[_fbthrift_compatible_with_TWrapped] | _fbthrift_python_mutable_types._ThriftListWrapper]=...
    ) -> _typing.Self: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_fbthrift_python_mutable_containers.MutableList[_fbthrift_TWrapped]]]]: ...
    def _to_python(self) -> "test.fixtures.python_capi.containers.thrift_types.IndirectionC": ...  # type: ignore
    def _to_mutable_python(self) -> _typing.Self: ...
    def _to_py3(self) -> "test.fixtures.python_capi.containers.types.IndirectionC": ...  # type: ignore
    def _to_py_deprecated(self) -> "containers.ttypes.IndirectionC": ...  # type: ignore
_fbthrift_IndirectionC = IndirectionC

IOBuf = _fbthrift_iobuf.IOBuf
small_vector_iobuf = _typing.List[_fbthrift_iobuf.IOBuf]
fbvector_string = _typing.List[str]
fbvector_fbvector_string = _typing.List[_fbthrift_python_mutable_containers.MutableList[str]]
CppWrapper = _fbthrift_TWrapped
ListOfWrapped = _typing.List[_fbthrift_TWrapped]
VecOfWrapped = _typing.List[_fbthrift_TWrapped]
ListOfWrappedAlias = _typing.List[_fbthrift_TWrapped]
