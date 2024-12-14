
# EXPERIMENTAL - DO NOT USE !!!
# See `experimental_generate_mutable_types` documentation in thrift compiler

#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import folly.iobuf as _fbthrift_iobuf

from abc import ABCMeta as _fbthrift_ABCMeta
import foo.thrift_abstract_types as _fbthrift_abstract_types
import thrift.python.types as _fbthrift_python_types
import thrift.python.mutable_types as _fbthrift_python_mutable_types
import thrift.python.mutable_exceptions as _fbthrift_python_mutable_exceptions
import thrift.python.mutable_typeinfos as _fbthrift_python_mutable_typeinfos



class Fields(metaclass=_fbthrift_python_mutable_types.MutableStructMeta):
    _fbthrift_SPEC = (
        _fbthrift_python_types.FieldInfo(
            100,  # id
            _fbthrift_python_types.FieldQualifier.Unqualified, # qualifier
            "injected_field",  # name
            "injected_field",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_string,  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
            8, # IDL type (see BaseTypeEnum)
        ),
        _fbthrift_python_types.FieldInfo(
            101,  # id
            _fbthrift_python_types.FieldQualifier.Optional, # qualifier
            "injected_structured_annotation_field",  # name
            "injected_structured_annotation_field",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_string,  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
            8, # IDL type (see BaseTypeEnum)
        ),
        _fbthrift_python_types.FieldInfo(
            102,  # id
            _fbthrift_python_types.FieldQualifier.Optional, # qualifier
            "injected_unstructured_annotation_field",  # name
            "injected_unstructured_annotation_field",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_string,  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
            8, # IDL type (see BaseTypeEnum)
        ),
    )

    @staticmethod
    def __get_thrift_name__() -> str:
        return "foo.Fields"

    @staticmethod
    def __get_thrift_uri__():
        return None

    @staticmethod
    def __get_metadata__():
        raise NotImplementedError(f"__get_metadata__() is not yet implemented for mutable thrift-python structs: {type(self)}")


    def _to_python(self):
        import thrift.python.converter
        import importlib
        immutable_types = importlib.import_module("foo.thrift_types")
        return thrift.python.converter.to_python_struct(immutable_types.Fields, self)

    def _to_mutable_python(self):
        return self

    def _to_py3(self):
        import importlib
        py3_types = importlib.import_module("foo.types")
        import thrift.py3.converter
        return thrift.py3.converter.to_py3_struct(py3_types.Fields, self)

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        try:
            py_deprecated_types = importlib.import_module("foo.ttypes")
            return thrift.util.converter.to_py_struct(py_deprecated_types.Fields, self)
        except ModuleNotFoundError:
            py_asyncio_types = importlib.import_module("foo.ttypes")
            return thrift.util.converter.to_py_struct(py_asyncio_types.Fields, self)

_fbthrift_ABCMeta.register(_fbthrift_abstract_types.Fields, Fields)
_fbthrift_Fields = Fields


_fbthrift_all_enums = [
]


_fbthrift_all_structs = [
    Fields,
]
_fbthrift_python_mutable_types.fill_specs(*_fbthrift_all_structs)
