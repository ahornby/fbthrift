#
# Autogenerated by Thrift for foo.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libcpp.memory cimport shared_ptr

cimport foo.types as _fbthrift_ctypes


cdef shared_ptr[_fbthrift_ctypes.cFields] Fields_convert_to_cpp(object inst) except*
cdef object Fields_from_cpp(const shared_ptr[_fbthrift_ctypes.cFields]& c_struct)

