#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-annotations/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#


cdef extern from "thrift/compiler/test/fixtures/basic-annotations/gen-cpp2/module_handlers.h" namespace "::cpp2":
    cdef cppclass cMyServiceSvIf "::cpp2::MyServiceSvIf":
        pass

cdef extern from "thrift/compiler/test/fixtures/basic-annotations/gen-cpp2/module_handlers.h" namespace "::cpp2":
    cdef cppclass cMyServicePrioParentSvIf "::cpp2::MyServicePrioParentSvIf":
        pass

cdef extern from "thrift/compiler/test/fixtures/basic-annotations/gen-cpp2/module_handlers.h" namespace "::cpp2":
    cdef cppclass cMyServicePrioChildSvIf "::cpp2::MyServicePrioChildSvIf":
        pass

cdef extern from "thrift/compiler/test/fixtures/basic-annotations/gen-cpp2/module_handlers.h" namespace "::cpp2":
    cdef cppclass cGoodServiceSvIf "::cpp2::GoodServiceSvIf":
        pass

cdef extern from "thrift/compiler/test/fixtures/basic-annotations/gen-cpp2/module_handlers.h" namespace "::cpp2":
    cdef cppclass cFooBarBazServiceSvIf "::cpp2::FooBarBazServiceSvIf":
        pass
