#
# Autogenerated by Thrift for included.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#


import folly.iobuf as _fbthrift_iobuf

from thrift.py3.reflection import (
    NumberType as __NumberType,
    StructType as __StructType,
    Qualifier as __Qualifier,
    StructSpec as __StructSpec,
    ListSpec as __ListSpec,
    SetSpec as __SetSpec,
    MapSpec as __MapSpec,
    FieldSpec as __FieldSpec,
)


import apache.thrift.fixtures.types.included.types as _apache_thrift_fixtures_types_included_types



def get_reflection__std_unordered_map__Map__i32_string() -> __MapSpec:
    return __MapSpec._fbthrift_create(
        key=int,
        key_kind=__NumberType.I32,
        value=str,
        value_kind=__NumberType.NOT_A_NUMBER,
    )

def get_reflection__List__std_unordered_map__Map__i32_string() -> __ListSpec :
    return __ListSpec._fbthrift_create(
        value=_apache_thrift_fixtures_types_included_types.std_unordered_map__Map__i32_string,
        kind=__NumberType.NOT_A_NUMBER,
    )

