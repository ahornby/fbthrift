#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import folly.iobuf as _fbthrift_iobuf
import thrift.python.types as _fbthrift_python_types
import thrift.python.exceptions as _fbthrift_python_exceptions


import include.thrift_types


class CustomException(metaclass=_fbthrift_python_exceptions.GeneratedErrorMeta):
    _fbthrift_SPEC = (
        _fbthrift_python_types.FieldInfo(
            1,  # id
            _fbthrift_python_types.FieldQualifier.Unqualified, # qualifier
            "name",  # name
            "name",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_string,  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
        ),
        _fbthrift_python_types.FieldInfo(
            2,  # id
            _fbthrift_python_types.FieldQualifier.Unqualified, # qualifier
            "result",  # name
            "result",  # python name (from @python.Name annotation)
            lambda: _fbthrift_python_types.EnumTypeInfo(Result),  # typeinfo
            lambda: Result.SO_SO,  # default value
            None,  # adapter info
            False, # field type is primitive
        ),
    )

    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.CustomException"

    @staticmethod
    def __get_thrift_uri__():
        return None

    @staticmethod
    def __get_metadata__():
        return _fbthrift_metadata__exception_CustomException()

    def _to_python(self):
        return self

    def _to_py3(self):
        import importlib
        py3_types = importlib.import_module("module.types")
        import thrift.py3.converter
        return thrift.py3.converter.to_py3_struct(py3_types.CustomException, self)

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        try:
            py_deprecated_types = importlib.import_module("module.ttypes")
            return thrift.util.converter.to_py_struct(py_deprecated_types.CustomException, self)
        except ModuleNotFoundError:
            py_asyncio_types = importlib.import_module("module.ttypes")
            return thrift.util.converter.to_py_struct(py_asyncio_types.CustomException, self)

# This unfortunately has to be down here to prevent circular imports
import module.thrift_metadata


class Result(_fbthrift_python_types.Enum, int):
    OK = 0
    SO_SO = 1
    GOOD = 2
    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.Result"

    @staticmethod
    def __get_thrift_uri__():
        return None

    @staticmethod
    def __get_metadata__():
        return module.thrift_metadata.gen_metadata_enum_Result()

    def _to_python(self):
        return self

    def _to_py3(self):
        import importlib
        py3_types = importlib.import_module("module.types")
        return py3_types.Result(self.value)

    def _to_py_deprecated(self):
        return self.value


_fbthrift_all_enums = [
    Result,
]

def _fbthrift_metadata__exception_CustomException():
    return module.thrift_metadata.gen_metadata_exception_CustomException()


_fbthrift_all_structs = [
    CustomException,
]
_fbthrift_python_types.fill_specs(*_fbthrift_all_structs)


_fbthrift_schema_b747839c13cb3aa5 = b"(\265/\375`\261\001\255\016\000\366W];\000ys)\202\263\223c\232\361\333\002l\235PZQ\245\204\306V\017\233z\255j\267$d\273\267\033)\212\275T\2566L\003(\341\365\352?7\317\252^G7\3438+\256\267\270I\nY\000K\000J\000#z\347\372\004\214\340\013\341\251\253\343\350\326\203a\200\nBb\0100\210Uqi\001\011\211D\241\270\030\2244\241#*\223\330\n\320VH1$5C\245\242\323\264\210\223\233\230\004\310\307U\023C\363S\244\223\3053\033?\031#%\033\257\245\235\036\321\354\363\372\252.+ \342\332\030\264?\200\0210\255DN^0\031\226\256\011'\005u]\344\304D\375Y}\327\264\325\003\"\376\323f\303$\334\206\362P\035\214\345\243\331\310ts\022es.\345\0169\325\274\2359\311\363~{\211\036!;\317\235\334y\325;\326M\311\036_\306\024nxN\342\356\216\346\372\251\272\271\236Y\005\261\374\377\207\010e\377\017,\272\321\235@\301\370\377\201k\240\2546\342\357\272\325\241X\226\343\343x\346\351\256Y\2759\022D\234\362\002\020\362\340\355\303\261\362\261\331j\351]I\255_\377\377\341\302[\230&\006\270\360&\3564q\001\"\376\3270\301\321c\343\377g\337\327\227\347\235Y\264\rW\276uu\337\345\264\344\306\272\235\277\025\010\".\377U\r\342\021\377\251j\227w\266\373\355yg\364`\374\237\003% P\"AJ\361\001b\301\366\006\367\2015\000\"\366\034eA\276\203\205\203\017\312\030Zf\316\014\327\325\354\252X\006\327(\301\011i$\346\005\270l2s7\2309\370a\254\001\006\365Bd\024\361\037\200\346\177\250\236\314K\365\026&\026;\032\212!\310\207\030\220\201\210\000\337\016\306P\014\004\r"



class _fbthrift_PrimitivesService_init_args(metaclass=_fbthrift_python_types.StructMeta):
    _fbthrift_SPEC = (
        _fbthrift_python_types.FieldInfo(
            1,  # id
            _fbthrift_python_types.FieldQualifier.Unqualified, # qualifier
            "param0",  # name
            "param0",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_i64,  # typeinfo
            None,  # default value
            None,  # adapter info
            True, # field type is primitive
        ),
        _fbthrift_python_types.FieldInfo(
            2,  # id
            _fbthrift_python_types.FieldQualifier.Unqualified, # qualifier
            "param1",  # name
            "param1",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_i64,  # typeinfo
            None,  # default value
            None,  # adapter info
            True, # field type is primitive
        ),
    )


class _fbthrift_PrimitivesService_init_result(metaclass=_fbthrift_python_types.StructMeta):
    _fbthrift_SPEC = (
        _fbthrift_python_types.FieldInfo(
            0,  # id
            _fbthrift_python_types.FieldQualifier.Optional, # qualifier
            "success",  # name
            "success", # name
            _fbthrift_python_types.typeinfo_i64,  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
        ),
    )


class _fbthrift_PrimitivesService_method_that_throws_args(metaclass=_fbthrift_python_types.StructMeta):
    _fbthrift_SPEC = (
    )


class _fbthrift_PrimitivesService_method_that_throws_result(metaclass=_fbthrift_python_types.StructMeta):
    _fbthrift_SPEC = (
        _fbthrift_python_types.FieldInfo(
            0,  # id
            _fbthrift_python_types.FieldQualifier.Optional, # qualifier
            "success",  # name
            "success", # name
            lambda: _fbthrift_python_types.EnumTypeInfo(Result),  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
        ),
        _fbthrift_python_types.FieldInfo(
            1,  # id
            _fbthrift_python_types.FieldQualifier.Optional, # qualifier
            "e",  # name
            "e",  # python name (from @python.Name annotation)
            lambda: _fbthrift_python_types.StructTypeInfo(CustomException),  # typeinfo
            None,  # default value
            None,  # adapter info
            False, # field type is primitive
        ),
    )


class _fbthrift_PrimitivesService_return_void_method_args(metaclass=_fbthrift_python_types.StructMeta):
    _fbthrift_SPEC = (
        _fbthrift_python_types.FieldInfo(
            1,  # id
            _fbthrift_python_types.FieldQualifier.Unqualified, # qualifier
            "id",  # name
            "id",  # python name (from @python.Name annotation)
            _fbthrift_python_types.typeinfo_i64,  # typeinfo
            None,  # default value
            None,  # adapter info
            True, # field type is primitive
        ),
    )


class _fbthrift_PrimitivesService_return_void_method_result(metaclass=_fbthrift_python_types.StructMeta):
    _fbthrift_SPEC = (
    )



_fbthrift_python_types.fill_specs(
    _fbthrift_PrimitivesService_init_args,
    _fbthrift_PrimitivesService_init_result,
    _fbthrift_PrimitivesService_method_that_throws_args,
    _fbthrift_PrimitivesService_method_that_throws_result,
    _fbthrift_PrimitivesService_return_void_method_args,
    _fbthrift_PrimitivesService_return_void_method_result,
)
