#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/inheritance/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from thrift.py3.server cimport ServiceInterface


cdef class MyRootInterface(ServiceInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_do_root
    pass

cdef class MyNodeInterface(MyRootInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_do_mid
    pass

cdef class MyLeafInterface(MyNodeInterface):
    # these are to avoid weird Cython multiple inheritance issue
    cdef bint _for_cython_do_leaf
    pass

