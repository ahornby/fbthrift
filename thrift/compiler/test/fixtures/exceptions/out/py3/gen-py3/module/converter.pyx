
#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/exceptions/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

cimport module.types as _fbthrift_ctypes

cdef shared_ptr[_fbthrift_cbindings.cFiery] Fiery_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.Fiery?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object Fiery_from_cpp(const shared_ptr[_fbthrift_cbindings.cFiery]& c_struct):
    return _fbthrift_ctypes.Fiery._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cSerious] Serious_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.Serious?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object Serious_from_cpp(const shared_ptr[_fbthrift_cbindings.cSerious]& c_struct):
    return _fbthrift_ctypes.Serious._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cComplexFieldNames] ComplexFieldNames_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ComplexFieldNames?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ComplexFieldNames_from_cpp(const shared_ptr[_fbthrift_cbindings.cComplexFieldNames]& c_struct):
    return _fbthrift_ctypes.ComplexFieldNames._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cCustomFieldNames] CustomFieldNames_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.CustomFieldNames?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object CustomFieldNames_from_cpp(const shared_ptr[_fbthrift_cbindings.cCustomFieldNames]& c_struct):
    return _fbthrift_ctypes.CustomFieldNames._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cExceptionWithPrimitiveField] ExceptionWithPrimitiveField_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ExceptionWithPrimitiveField?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ExceptionWithPrimitiveField_from_cpp(const shared_ptr[_fbthrift_cbindings.cExceptionWithPrimitiveField]& c_struct):
    return _fbthrift_ctypes.ExceptionWithPrimitiveField._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cExceptionWithStructuredAnnotation] ExceptionWithStructuredAnnotation_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ExceptionWithStructuredAnnotation?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ExceptionWithStructuredAnnotation_from_cpp(const shared_ptr[_fbthrift_cbindings.cExceptionWithStructuredAnnotation]& c_struct):
    return _fbthrift_ctypes.ExceptionWithStructuredAnnotation._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cBanal] Banal_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.Banal?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object Banal_from_cpp(const shared_ptr[_fbthrift_cbindings.cBanal]& c_struct):
    return _fbthrift_ctypes.Banal._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)


