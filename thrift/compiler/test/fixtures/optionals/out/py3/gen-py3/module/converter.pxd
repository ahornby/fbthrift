#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/optionals/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libcpp.memory cimport shared_ptr

cimport module.cbindings as _fbthrift_cbindings


cdef shared_ptr[_fbthrift_cbindings.cColor] Color_convert_to_cpp(object inst) except*
cdef object Color_from_cpp(const shared_ptr[_fbthrift_cbindings.cColor]& c_struct)

cdef shared_ptr[_fbthrift_cbindings.cVehicle] Vehicle_convert_to_cpp(object inst) except*
cdef object Vehicle_from_cpp(const shared_ptr[_fbthrift_cbindings.cVehicle]& c_struct)

cdef shared_ptr[_fbthrift_cbindings.cPerson] Person_convert_to_cpp(object inst) except*
cdef object Person_from_cpp(const shared_ptr[_fbthrift_cbindings.cPerson]& c_struct)

