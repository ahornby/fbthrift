#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/req-opt/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
import typing as _typing

import folly.iobuf as _fbthrift_iobuf
import thrift.py3.builder


import module.types as _module_types


_fbthrift_struct_type__Foo = _module_types.Foo
class Foo_Builder(thrift.py3.builder.StructBuilder):
    _struct_type = _fbthrift_struct_type__Foo

    def __init__(self):
        self.myInteger: _typing.Optional[int] = None
        self.myString: _typing.Optional[str] = None
        self.myBools: _typing.Optional[list] = None
        self.myNumbers: _typing.Optional[list] = None

    def __iter__(self):
        yield "myInteger", self.myInteger
        yield "myString", self.myString
        yield "myBools", self.myBools
        yield "myNumbers", self.myNumbers

