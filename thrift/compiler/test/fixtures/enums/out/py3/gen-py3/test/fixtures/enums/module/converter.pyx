
#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/enums/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

cimport test.fixtures.enums.module.types as _fbthrift_ctypes

cdef shared_ptr[_fbthrift_cbindings.cSomeStruct] SomeStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.SomeStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object SomeStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cSomeStruct]& c_struct):
    return _fbthrift_ctypes.SomeStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMyStruct] MyStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.MyStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object MyStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyStruct]& c_struct):
    return _fbthrift_ctypes.MyStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)


