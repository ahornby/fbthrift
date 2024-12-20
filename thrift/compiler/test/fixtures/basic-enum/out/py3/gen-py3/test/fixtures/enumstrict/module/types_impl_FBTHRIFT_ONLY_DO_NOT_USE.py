#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-enum/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum
import thrift.py3.types
import test.fixtures.enumstrict.module.thrift_metadata as _fbthrift_python_metadata
try:
    import test.fixtures.enumstrict.module.thrift_types as _fbthrift_python_types
except Exception: # TODO(T205494848): fix thrift-python import failures
    _fbthrift_python_types = None

_fbthrift__module_name__ = "test.fixtures.enumstrict.module.types"



class EmptyEnum(thrift.py3.types.CompiledEnum):

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_metadata.gen_metadata_enum_EmptyEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.EmptyEnum"

    def _to_python(self):
        if _fbthrift_python_types is None:
            raise AttributeError(
                "Enum EmptyEnum doesn't define `_to_python` because couldn't import "
                "test.fixtures.enumstrict.module.thrift_types"
            )

        return _fbthrift_python_types.EmptyEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_




class MyEnum(thrift.py3.types.CompiledEnum):
    ONE = 1
    TWO = 2

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_metadata.gen_metadata_enum_MyEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.MyEnum"

    def _to_python(self):
        if _fbthrift_python_types is None:
            raise AttributeError(
                "Enum MyEnum doesn't define `_to_python` because couldn't import "
                "test.fixtures.enumstrict.module.thrift_types"
            )

        return _fbthrift_python_types.MyEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_




class MyUseIntrinsicDefaultEnum(thrift.py3.types.CompiledEnum):
    ZERO = 0
    ONE = 1
    TWO = 2

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_metadata.gen_metadata_enum_MyUseIntrinsicDefaultEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.MyUseIntrinsicDefaultEnum"

    def _to_python(self):
        if _fbthrift_python_types is None:
            raise AttributeError(
                "Enum MyUseIntrinsicDefaultEnum doesn't define `_to_python` because couldn't import "
                "test.fixtures.enumstrict.module.thrift_types"
            )

        return _fbthrift_python_types.MyUseIntrinsicDefaultEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_




class MyBigEnum(thrift.py3.types.CompiledEnum):
    UNKNOWN = 0
    ONE = 1
    TWO = 2
    THREE = 3
    FOUR = 4
    FIVE = 5
    SIX = 6
    SEVEN = 7
    EIGHT = 8
    NINE = 9
    TEN = 10
    ELEVEN = 11
    TWELVE = 12
    THIRTEEN = 13
    FOURTEEN = 14
    FIFTEEN = 15
    SIXTEEN = 16
    SEVENTEEN = 17
    EIGHTEEN = 18
    NINETEEN = 19

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_metadata.gen_metadata_enum_MyBigEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.MyBigEnum"

    def _to_python(self):
        if _fbthrift_python_types is None:
            raise AttributeError(
                "Enum MyBigEnum doesn't define `_to_python` because couldn't import "
                "test.fixtures.enumstrict.module.thrift_types"
            )

        return _fbthrift_python_types.MyBigEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_




