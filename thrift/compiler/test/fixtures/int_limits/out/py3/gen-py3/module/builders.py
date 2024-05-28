#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import thrift.py3.builder


import module.types as _module_types


_fbthrift_struct_type__Limits = _module_types.Limits
class Limits_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__Limits

    def __init__(self):
        self.max_i64_field: _typing.Optional[int] = None
        self.min_i64_field: _typing.Optional[int] = None
        self.max_i32_field: _typing.Optional[int] = None
        self.min_i32_field: _typing.Optional[int] = None
        self.max_i16_field: _typing.Optional[int] = None
        self.min_i16_field: _typing.Optional[int] = None
        self.max_byte_field: _typing.Optional[int] = None
        self.min_byte_field: _typing.Optional[int] = None

    def __iter__(self):
        yield "max_i64_field", self.max_i64_field
        yield "min_i64_field", self.min_i64_field
        yield "max_i32_field", self.max_i32_field
        yield "min_i32_field", self.min_i32_field
        yield "max_i16_field", self.max_i16_field
        yield "min_i16_field", self.min_i16_field
        yield "max_byte_field", self.max_byte_field
        yield "min_byte_field", self.min_byte_field

