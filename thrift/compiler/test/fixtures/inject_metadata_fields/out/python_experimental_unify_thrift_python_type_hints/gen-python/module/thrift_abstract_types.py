
# EXPERIMENTAL - DO NOT USE !!!
# See `experimental_unify_thrift_python_type_hints` documentation in thrift compiler

#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import abc as _abc
import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import foo.thrift_abstract_types


class Fields(_abc.ABC):
    @property
    @_abc.abstractmethod
    def injected_field(self) -> str: ...
    @_abc.abstractmethod
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str]]]: ...
    @_abc.abstractmethod
    def _to_mutable_python(self) -> "module.thrift_mutable_types.Fields": ...  # type: ignore
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.Fields": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.Fields": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.Fields": ...  # type: ignore

class FieldsInjectedToEmptyStruct(_abc.ABC):
    @property
    @_abc.abstractmethod
    def injected_field(self) -> str: ...
    @_abc.abstractmethod
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str]]]: ...
    @_abc.abstractmethod
    def _to_mutable_python(self) -> "module.thrift_mutable_types.FieldsInjectedToEmptyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.FieldsInjectedToEmptyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.FieldsInjectedToEmptyStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.FieldsInjectedToEmptyStruct": ...  # type: ignore

class FieldsInjectedToStruct(_abc.ABC):
    @property
    @_abc.abstractmethod
    def injected_field(self) -> str: ...
    @property
    @_abc.abstractmethod
    def string_field(self) -> str: ...
    @_abc.abstractmethod
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str, str]]]: ...
    @_abc.abstractmethod
    def _to_mutable_python(self) -> "module.thrift_mutable_types.FieldsInjectedToStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.FieldsInjectedToStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.FieldsInjectedToStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.FieldsInjectedToStruct": ...  # type: ignore

class FieldsInjectedWithIncludedStruct(_abc.ABC):
    @property
    @_abc.abstractmethod
    def injected_unstructured_annotation_field(self) -> _typing.Optional[str]: ...
    @property
    @_abc.abstractmethod
    def injected_structured_annotation_field(self) -> _typing.Optional[str]: ...
    @property
    @_abc.abstractmethod
    def injected_field(self) -> str: ...
    @property
    @_abc.abstractmethod
    def string_field(self) -> str: ...
    @_abc.abstractmethod
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str, str, str, str]]]: ...
    @_abc.abstractmethod
    def _to_mutable_python(self) -> "module.thrift_mutable_types.FieldsInjectedWithIncludedStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_python(self) -> "module.thrift_types.FieldsInjectedWithIncludedStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py3(self) -> "module.types.FieldsInjectedWithIncludedStruct": ...  # type: ignore
    @_abc.abstractmethod
    def _to_py_deprecated(self) -> "module.ttypes.FieldsInjectedWithIncludedStruct": ...  # type: ignore
