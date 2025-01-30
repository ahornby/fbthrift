#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/py3/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
from module.thrift_types import *
import importlib
import thrift.python.types as _fbthrift_python_types
import module.thrift_types as _module_thrift_types
def get_types_reflection():
    return importlib.import_module(
        "module.types_reflection"
    )
class List__i16__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i16,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__i16(_fbthrift_python_types.List, metaclass=List__i16__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__i16._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__i16()

class List__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__i32(_fbthrift_python_types.List, metaclass=List__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__i32()

class List__i64__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i64,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__i64(_fbthrift_python_types.List, metaclass=List__i64__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__i64._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__i64()

class List__string__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__string(_fbthrift_python_types.List, metaclass=List__string__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__string._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__string()

class List__SimpleStruct__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.StructTypeInfo(_module_thrift_types.SimpleStruct),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__SimpleStruct(_fbthrift_python_types.List, metaclass=List__SimpleStruct__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__SimpleStruct._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__SimpleStruct()

class Set__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Set) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Set__i32(_fbthrift_python_types.Set, metaclass=Set__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Set__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Set__i32()

class Set__string__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Set) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Set__string(_fbthrift_python_types.Set, metaclass=Set__string__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Set__string._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Set__string()

class Map__string_string__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
            _fbthrift_python_types.typeinfo_string,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__string_string(_fbthrift_python_types.Map, metaclass=Map__string_string__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__string_string._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_string()

class Map__string_SimpleStruct__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
            _fbthrift_python_types.StructTypeInfo(_module_thrift_types.SimpleStruct),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__string_SimpleStruct(_fbthrift_python_types.Map, metaclass=Map__string_SimpleStruct__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__string_SimpleStruct._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_SimpleStruct()

class Map__string_i16__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
            _fbthrift_python_types.typeinfo_i16,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__string_i16(_fbthrift_python_types.Map, metaclass=Map__string_i16__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__string_i16._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_i16()

class List__List__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.ListTypeInfo(_fbthrift_python_types.typeinfo_i32),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__List__i32(_fbthrift_python_types.List, metaclass=List__List__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__List__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__List__i32()

class Map__string_i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__string_i32(_fbthrift_python_types.Map, metaclass=Map__string_i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__string_i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_i32()

class Map__string_Map__string_i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
            _fbthrift_python_types.MapTypeInfo(_fbthrift_python_types.typeinfo_string, _fbthrift_python_types.typeinfo_i32),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__string_Map__string_i32(_fbthrift_python_types.Map, metaclass=Map__string_Map__string_i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__string_Map__string_i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_Map__string_i32()

class List__Set__string__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.SetTypeInfo(_fbthrift_python_types.typeinfo_string),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__Set__string(_fbthrift_python_types.List, metaclass=List__Set__string__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__Set__string._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__Set__string()

class Map__string_List__SimpleStruct__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_string,
            _fbthrift_python_types.ListTypeInfo(_fbthrift_python_types.StructTypeInfo(_module_thrift_types.SimpleStruct)),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__string_List__SimpleStruct(_fbthrift_python_types.Map, metaclass=Map__string_List__SimpleStruct__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__string_List__SimpleStruct._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__string_List__SimpleStruct()

class List__List__string__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.ListTypeInfo(_fbthrift_python_types.typeinfo_string),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__List__string(_fbthrift_python_types.List, metaclass=List__List__string__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__List__string._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__List__string()

class List__Set__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.SetTypeInfo(_fbthrift_python_types.typeinfo_i32),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__Set__i32(_fbthrift_python_types.List, metaclass=List__Set__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__Set__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__Set__i32()

class List__Map__string_string__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.MapTypeInfo(_fbthrift_python_types.typeinfo_string, _fbthrift_python_types.typeinfo_string),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__Map__string_string(_fbthrift_python_types.List, metaclass=List__Map__string_string__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__Map__string_string._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__Map__string_string()

class List__binary__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_binary,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__binary(_fbthrift_python_types.List, metaclass=List__binary__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__binary._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__binary()

class Set__binary__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_binary,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Set) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Set__binary(_fbthrift_python_types.Set, metaclass=Set__binary__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Set__binary._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Set__binary()

class List__AnEnum__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.EnumTypeInfo(_module_thrift_types.AnEnum),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__AnEnum(_fbthrift_python_types.List, metaclass=List__AnEnum__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__AnEnum._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__AnEnum()

class _std_unordered_map__Map__i32_i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _std_unordered_map__Map__i32_i32(_fbthrift_python_types.Map, metaclass=_std_unordered_map__Map__i32_i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_std_unordered_map__Map__i32_i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___std_unordered_map__Map__i32_i32()

class _MyType__List__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _MyType__List__i32(_fbthrift_python_types.List, metaclass=_MyType__List__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_MyType__List__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___MyType__List__i32()

class _MyType__Set__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Set) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _MyType__Set__i32(_fbthrift_python_types.Set, metaclass=_MyType__Set__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_MyType__Set__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___MyType__Set__i32()

class _MyType__Map__i32_i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _MyType__Map__i32_i32(_fbthrift_python_types.Map, metaclass=_MyType__Map__i32_i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_MyType__Map__i32_i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___MyType__Map__i32_i32()

class _py3_simple_AdaptedList__List__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _py3_simple_AdaptedList__List__i32(_fbthrift_python_types.List, metaclass=_py3_simple_AdaptedList__List__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_py3_simple_AdaptedList__List__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___py3_simple_AdaptedList__List__i32()

class _py3_simple_AdaptedSet__Set__i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Set) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _py3_simple_AdaptedSet__Set__i32(_fbthrift_python_types.Set, metaclass=_py3_simple_AdaptedSet__Set__i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_py3_simple_AdaptedSet__Set__i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___py3_simple_AdaptedSet__Set__i32()

class _py3_simple_AdaptedMap__Map__i32_i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class _py3_simple_AdaptedMap__Map__i32_i32(_fbthrift_python_types.Map, metaclass=_py3_simple_AdaptedMap__Map__i32_i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *_py3_simple_AdaptedMap__Map__i32_i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection___py3_simple_AdaptedMap__Map__i32_i32()

class Map__i32_double__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.typeinfo_i32,
            _fbthrift_python_types.typeinfo_double,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__i32_double(_fbthrift_python_types.Map, metaclass=Map__i32_double__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__i32_double._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__i32_double()

class List__Map__i32_double__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.MapTypeInfo(_fbthrift_python_types.typeinfo_i32, _fbthrift_python_types.typeinfo_double),
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.List) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class List__Map__i32_double(_fbthrift_python_types.List, metaclass=List__Map__i32_double__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *List__Map__i32_double._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__List__Map__i32_double()

class Map__AnEnumRenamed_i32__Meta(type):
    def _fbthrift_type_info(cls):
        return (
            _fbthrift_python_types.EnumTypeInfo(_module_thrift_types.AnEnumRenamed),
            _fbthrift_python_types.typeinfo_i32,
        )

    def __instancecheck__(cls, instance):
        return (
            isinstance(instance, _fbthrift_python_types.Map) and
            instance._fbthrift_same_type(*cls._fbthrift_type_info())
        )

class Map__AnEnumRenamed_i32(_fbthrift_python_types.Map, metaclass=Map__AnEnumRenamed_i32__Meta):
    def __init__(self, *args, **kwargs):
        super().__init__(
            *Map__AnEnumRenamed_i32._fbthrift_type_info(),
            *args,
            **kwargs,
        )

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Map__AnEnumRenamed_i32()

