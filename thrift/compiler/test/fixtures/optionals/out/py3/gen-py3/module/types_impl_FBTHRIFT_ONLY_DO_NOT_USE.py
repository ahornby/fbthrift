#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/optionals/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum
import thrift.py3.types
import module.thrift_enums as _fbthrift_python_enums

_fbthrift__module_name__ = "module.types"



class Animal(thrift.py3.types.CompiledEnum, int):
    DOG = 1
    CAT = 2
    TARANTULA = 3

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_Animal()

    @staticmethod
    def __get_thrift_name__():
        return "module.Animal"

    def _to_python(self):
        return _fbthrift_python_enums.Animal(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_


    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_

