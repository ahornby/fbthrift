#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-annotations/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import thrift.py3.builder


import module.types as _module_types


_fbthrift_struct_type__MyStructNestedAnnotation = _module_types.MyStructNestedAnnotation
class MyStructNestedAnnotation_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__MyStructNestedAnnotation

    def __init__(self):
        self.name: _typing.Optional[str] = None

    def __iter__(self):
        yield "name", self.name

_fbthrift_struct_type__SecretStruct = _module_types.SecretStruct
class SecretStruct_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__SecretStruct

    def __init__(self):
        self.id: _typing.Optional[int] = None
        self.password: _typing.Optional[str] = None

    def __iter__(self):
        yield "id", self.id
        yield "password", self.password

