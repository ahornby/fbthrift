
#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#


cdef extern from "thrift/compiler/test/fixtures/includes/gen-cpp2/includes_types.h":
    cdef cppclass cIncluded "::cpp2::Included":
        cIncluded()

cdef cIncluded Included_convert_to_cpp(object inst) except*
cdef object Included_from_cpp(const cIncluded& c_struct)

