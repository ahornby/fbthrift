#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import apache.thrift.metadata.thrift_types as _fbthrift_metadata
import thrift.python.types as _fbthrift_python_types
import typing as _std_python_typing

class _fbthrift_compatible_with_EmptyEnum:
    pass


class EmptyEnum(_fbthrift_python_types.Enum, int, _fbthrift_compatible_with_EmptyEnum):
    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.EmptyEnum"

    @staticmethod
    def __get_thrift_uri__() -> _std_python_typing.Optional[str]:
        return None

    @staticmethod
    def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
        return gen_metadata_enum_EmptyEnum()

    def _to_python(self) -> "EmptyEnum":
        return self

    def _to_py3(self) -> "module.types.EmptyEnum": # type: ignore
        import importlib
        py3_types = importlib.import_module("module.types")
        return py3_types.EmptyEnum(self.value)

    def _to_py_deprecated(self) -> int:
        return self.value
import typing as _std_python_typing

class _fbthrift_compatible_with_City:
    pass


class City(_fbthrift_python_types.Enum, int, _fbthrift_compatible_with_City):
    NYC = 0
    MPK = 1
    SEA = 2
    LON = 3
    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.City"

    @staticmethod
    def __get_thrift_uri__() -> _std_python_typing.Optional[str]:
        return None

    @staticmethod
    def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
        return gen_metadata_enum_City()

    def _to_python(self) -> "City":
        return self

    def _to_py3(self) -> "module.types.City": # type: ignore
        import importlib
        py3_types = importlib.import_module("module.types")
        return py3_types.City(self.value)

    def _to_py_deprecated(self) -> int:
        return self.value
import typing as _std_python_typing

class _fbthrift_compatible_with_Company:
    pass


class Company(_fbthrift_python_types.Enum, int, _fbthrift_compatible_with_Company):
    FACEBOOK = 0
    WHATSAPP = 1
    OCULUS = 2
    INSTAGRAM = 3
    __FRIEND__FEED = 4
    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.Company"

    @staticmethod
    def __get_thrift_uri__() -> _std_python_typing.Optional[str]:
        return None

    @staticmethod
    def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
        return gen_metadata_enum_Company()

    def _to_python(self) -> "Company":
        return self

    def _to_py3(self) -> "module.types.Company": # type: ignore
        import importlib
        py3_types = importlib.import_module("module.types")
        return py3_types.Company(self.value)

    def _to_py_deprecated(self) -> int:
        return self.value

def _fbthrift_gen_metadata_enum_EmptyEnum(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.EmptyEnum"

    if qualified_name in metadata_struct.enums:
        return metadata_struct
    elements = {
    }
    structured_annotations = [
    ]
    enum_dict = dict(metadata_struct.enums)
    enum_dict[qualified_name] = _fbthrift_metadata.ThriftEnum(name=qualified_name, elements=elements, structured_annotations=structured_annotations)
    new_struct = metadata_struct(enums=enum_dict)

    return new_struct

def gen_metadata_enum_EmptyEnum() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_enum_EmptyEnum(
        _fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={})
    )

def _fbthrift_gen_metadata_enum_City(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.City"

    if qualified_name in metadata_struct.enums:
        return metadata_struct
    elements = {
        0: "NYC",
        1: "MPK",
        2: "SEA",
        3: "LON",
    }
    structured_annotations = [
    ]
    enum_dict = dict(metadata_struct.enums)
    enum_dict[qualified_name] = _fbthrift_metadata.ThriftEnum(name=qualified_name, elements=elements, structured_annotations=structured_annotations)
    new_struct = metadata_struct(enums=enum_dict)

    return new_struct

def gen_metadata_enum_City() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_enum_City(
        _fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={})
    )

def _fbthrift_gen_metadata_enum_Company(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.Company"

    if qualified_name in metadata_struct.enums:
        return metadata_struct
    elements = {
        0: "FACEBOOK",
        1: "WHATSAPP",
        2: "OCULUS",
        3: "INSTAGRAM",
        4: "__FRIEND__FEED",
    }
    structured_annotations = [
    ]
    enum_dict = dict(metadata_struct.enums)
    enum_dict[qualified_name] = _fbthrift_metadata.ThriftEnum(name=qualified_name, elements=elements, structured_annotations=structured_annotations)
    new_struct = metadata_struct(enums=enum_dict)

    return new_struct

def gen_metadata_enum_Company() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_enum_Company(
        _fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={})
    )

