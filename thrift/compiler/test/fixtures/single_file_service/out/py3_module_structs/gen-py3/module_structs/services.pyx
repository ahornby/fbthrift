#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/single_file_service/src/module_structs.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

cimport cython
from typing import AsyncIterator
from cpython.version cimport PY_VERSION_HEX
from libc.stdint cimport (
    int8_t as cint8_t,
    int16_t as cint16_t,
    int32_t as cint32_t,
    int64_t as cint64_t,
)
from libcpp.memory cimport shared_ptr, make_shared, unique_ptr
from libcpp.string cimport string
from libcpp cimport bool as cbool
from cpython cimport bool as pbool
from libcpp.vector cimport vector
from libcpp.set cimport set as cset
from libcpp.map cimport map as cmap
from libcpp.utility cimport move as cmove
from libcpp.pair cimport pair
from cython.operator cimport dereference as deref
from cpython.ref cimport PyObject
from thrift.python.exceptions cimport (
    ApplicationError as __ApplicationError,
    cTApplicationException,
    cTApplicationExceptionType__UNKNOWN,
)
from thrift.py3.server cimport ServiceInterface
from thrift.python.server_impl.request_context cimport RequestContext, Cpp2RequestContext
from thrift.python.server_impl.request_context import RequestContext
from folly cimport (
  cFollyPromise,
  cFollyUnit,
  c_unit,
)
from thrift.python.common cimport (
    cThriftServiceMetadataResponse as __fbthrift_cThriftServiceMetadataResponse,
    ServiceMetadata,
    MetadataBox as __MetadataBox,
)

from thrift.py3.server cimport THRIFT_REQUEST_CONTEXT as __THRIFT_REQUEST_CONTEXT
from thrift.py3.types cimport make_unique

cimport folly.futures
from folly.executor cimport get_executor
cimport folly.iobuf as _fbthrift_iobuf
import folly.iobuf as _fbthrift_iobuf
from folly.iobuf cimport move as move_iobuf
from folly.memory cimport to_shared_ptr as __to_shared_ptr

cimport module_structs.types as _module_structs_types
cimport module_structs.cbindings as _module_structs_cbindings
import module_structs.types as _module_structs_types

cimport module_structs.services_interface as _fbthrift_services_interface

import asyncio
import functools
import sys
import traceback
import types as _py_types




@cython.auto_pickle(False)
cdef class Promise_cFollyUnit:
    cdef cFollyPromise[cFollyUnit]* cPromise

    def __cinit__(self):
        self.cPromise = new cFollyPromise[cFollyUnit](cFollyPromise[cFollyUnit].makeEmpty())

    def __dealloc__(self):
        del self.cPromise

    @staticmethod
    cdef _fbthrift_create(cFollyPromise[cFollyUnit] cPromise):
        cdef Promise_cFollyUnit inst = Promise_cFollyUnit.__new__(Promise_cFollyUnit)
        inst.cPromise[0] = cmove(cPromise)
        return inst



