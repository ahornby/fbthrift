#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from thrift.py3.server cimport ServiceInterface


cdef class MyServiceInterface(ServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_hasDataById
    cdef bint _for_cython_getDataById
    cdef bint _for_cython_putDataById
    cdef bint _for_cython_lobDataById
    pass

cdef class MyServiceFastInterface(ServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_hasDataById
    cdef bint _for_cython_getDataById
    cdef bint _for_cython_putDataById
    cdef bint _for_cython_lobDataById
    pass

cdef class DbMixedStackArgumentsInterface(ServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_getDataByKey0
    cdef bint _for_cython_getDataByKey1
    pass

