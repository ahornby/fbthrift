#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/transitive-deps/src/a.thrift
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

import b.types as _b_types
import c.types as _c_types

import a.types as _a_types



def get_reflection__A() -> __StructSpec:
    spec: __StructSpec = __StructSpec._fbthrift_create(
        name="A",
        kind=__StructType.STRUCT,
        annotations={
        },
    )
    spec.add_field(
        __FieldSpec._fbthrift_create(
            id=1,
            name="b",
            py_name="b",
            type=_a_types.List__List__c_C,
            kind=__NumberType.NOT_A_NUMBER,
            qualifier=__Qualifier.UNQUALIFIED,
            default=None,
            annotations={
            },
        ),
    )
    spec.add_field(
        __FieldSpec._fbthrift_create(
            id=2,
            name="other",
            py_name="other",
            type=_a_types.List__c_C,
            kind=__NumberType.NOT_A_NUMBER,
            qualifier=__Qualifier.UNQUALIFIED,
            default=None,
            annotations={
            },
        ),
    )
    return spec
def get_reflection__List__c_C() -> __ListSpec :
    return __ListSpec._fbthrift_create(
        value=_c_types.C,
        kind=__NumberType.NOT_A_NUMBER,
    )

def get_reflection__List__List__c_C() -> __ListSpec :
    return __ListSpec._fbthrift_create(
        value=_a_types.List__c_C,
        kind=__NumberType.NOT_A_NUMBER,
    )

