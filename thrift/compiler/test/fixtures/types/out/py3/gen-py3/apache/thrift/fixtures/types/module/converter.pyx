
#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/types/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

cimport apache.thrift.fixtures.types.module.types as _fbthrift_ctypes


cdef shared_ptr[_fbthrift_cbindings.cempty_struct] empty_struct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.empty_struct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object empty_struct_from_cpp(const shared_ptr[_fbthrift_cbindings.cempty_struct]& c_struct):
    return _fbthrift_ctypes.empty_struct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cdecorated_struct] decorated_struct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.decorated_struct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object decorated_struct_from_cpp(const shared_ptr[_fbthrift_cbindings.cdecorated_struct]& c_struct):
    return _fbthrift_ctypes.decorated_struct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cContainerStruct] ContainerStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ContainerStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ContainerStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cContainerStruct]& c_struct):
    return _fbthrift_ctypes.ContainerStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cCppTypeStruct] CppTypeStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.CppTypeStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object CppTypeStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cCppTypeStruct]& c_struct):
    return _fbthrift_ctypes.CppTypeStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cVirtualStruct] VirtualStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.VirtualStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object VirtualStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cVirtualStruct]& c_struct):
    return _fbthrift_ctypes.VirtualStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMyStructWithForwardRefEnum] MyStructWithForwardRefEnum_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.MyStructWithForwardRefEnum?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object MyStructWithForwardRefEnum_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyStructWithForwardRefEnum]& c_struct):
    return _fbthrift_ctypes.MyStructWithForwardRefEnum._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cTrivialNumeric] TrivialNumeric_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.TrivialNumeric?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object TrivialNumeric_from_cpp(const shared_ptr[_fbthrift_cbindings.cTrivialNumeric]& c_struct):
    return _fbthrift_ctypes.TrivialNumeric._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cTrivialNestedWithDefault] TrivialNestedWithDefault_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.TrivialNestedWithDefault?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object TrivialNestedWithDefault_from_cpp(const shared_ptr[_fbthrift_cbindings.cTrivialNestedWithDefault]& c_struct):
    return _fbthrift_ctypes.TrivialNestedWithDefault._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cComplexString] ComplexString_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ComplexString?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ComplexString_from_cpp(const shared_ptr[_fbthrift_cbindings.cComplexString]& c_struct):
    return _fbthrift_ctypes.ComplexString._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cComplexNestedWithDefault] ComplexNestedWithDefault_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ComplexNestedWithDefault?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ComplexNestedWithDefault_from_cpp(const shared_ptr[_fbthrift_cbindings.cComplexNestedWithDefault]& c_struct):
    return _fbthrift_ctypes.ComplexNestedWithDefault._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMinPadding] MinPadding_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.MinPadding?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object MinPadding_from_cpp(const shared_ptr[_fbthrift_cbindings.cMinPadding]& c_struct):
    return _fbthrift_ctypes.MinPadding._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMinPaddingWithCustomType] MinPaddingWithCustomType_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.MinPaddingWithCustomType?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object MinPaddingWithCustomType_from_cpp(const shared_ptr[_fbthrift_cbindings.cMinPaddingWithCustomType]& c_struct):
    return _fbthrift_ctypes.MinPaddingWithCustomType._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMyStruct] MyStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.MyStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object MyStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyStruct]& c_struct):
    return _fbthrift_ctypes.MyStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMyDataItem] MyDataItem_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.MyDataItem?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object MyDataItem_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyDataItem]& c_struct):
    return _fbthrift_ctypes.MyDataItem._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cRenaming] Renaming_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.Renaming?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object Renaming_from_cpp(const shared_ptr[_fbthrift_cbindings.cRenaming]& c_struct):
    return _fbthrift_ctypes.Renaming._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cAnnotatedTypes] AnnotatedTypes_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.AnnotatedTypes?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object AnnotatedTypes_from_cpp(const shared_ptr[_fbthrift_cbindings.cAnnotatedTypes]& c_struct):
    return _fbthrift_ctypes.AnnotatedTypes._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cForwardUsageRoot] ForwardUsageRoot_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ForwardUsageRoot?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ForwardUsageRoot_from_cpp(const shared_ptr[_fbthrift_cbindings.cForwardUsageRoot]& c_struct):
    return _fbthrift_ctypes.ForwardUsageRoot._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cForwardUsageStruct] ForwardUsageStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ForwardUsageStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ForwardUsageStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cForwardUsageStruct]& c_struct):
    return _fbthrift_ctypes.ForwardUsageStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cForwardUsageByRef] ForwardUsageByRef_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.ForwardUsageByRef?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object ForwardUsageByRef_from_cpp(const shared_ptr[_fbthrift_cbindings.cForwardUsageByRef]& c_struct):
    return _fbthrift_ctypes.ForwardUsageByRef._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cIncompleteMap] IncompleteMap_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.IncompleteMap?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object IncompleteMap_from_cpp(const shared_ptr[_fbthrift_cbindings.cIncompleteMap]& c_struct):
    return _fbthrift_ctypes.IncompleteMap._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cIncompleteMapDep] IncompleteMapDep_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.IncompleteMapDep?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object IncompleteMapDep_from_cpp(const shared_ptr[_fbthrift_cbindings.cIncompleteMapDep]& c_struct):
    return _fbthrift_ctypes.IncompleteMapDep._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cCompleteMap] CompleteMap_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.CompleteMap?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object CompleteMap_from_cpp(const shared_ptr[_fbthrift_cbindings.cCompleteMap]& c_struct):
    return _fbthrift_ctypes.CompleteMap._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cCompleteMapDep] CompleteMapDep_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.CompleteMapDep?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object CompleteMapDep_from_cpp(const shared_ptr[_fbthrift_cbindings.cCompleteMapDep]& c_struct):
    return _fbthrift_ctypes.CompleteMapDep._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cIncompleteList] IncompleteList_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.IncompleteList?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object IncompleteList_from_cpp(const shared_ptr[_fbthrift_cbindings.cIncompleteList]& c_struct):
    return _fbthrift_ctypes.IncompleteList._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cIncompleteListDep] IncompleteListDep_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.IncompleteListDep?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object IncompleteListDep_from_cpp(const shared_ptr[_fbthrift_cbindings.cIncompleteListDep]& c_struct):
    return _fbthrift_ctypes.IncompleteListDep._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cCompleteList] CompleteList_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.CompleteList?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object CompleteList_from_cpp(const shared_ptr[_fbthrift_cbindings.cCompleteList]& c_struct):
    return _fbthrift_ctypes.CompleteList._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cCompleteListDep] CompleteListDep_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.CompleteListDep?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object CompleteListDep_from_cpp(const shared_ptr[_fbthrift_cbindings.cCompleteListDep]& c_struct):
    return _fbthrift_ctypes.CompleteListDep._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cAdaptedList] AdaptedList_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.AdaptedList?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object AdaptedList_from_cpp(const shared_ptr[_fbthrift_cbindings.cAdaptedList]& c_struct):
    return _fbthrift_ctypes.AdaptedList._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cDependentAdaptedList] DependentAdaptedList_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.DependentAdaptedList?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object DependentAdaptedList_from_cpp(const shared_ptr[_fbthrift_cbindings.cDependentAdaptedList]& c_struct):
    return _fbthrift_ctypes.DependentAdaptedList._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cAllocatorAware] AllocatorAware_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.AllocatorAware?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object AllocatorAware_from_cpp(const shared_ptr[_fbthrift_cbindings.cAllocatorAware]& c_struct):
    return _fbthrift_ctypes.AllocatorAware._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cAllocatorAware2] AllocatorAware2_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.AllocatorAware2?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object AllocatorAware2_from_cpp(const shared_ptr[_fbthrift_cbindings.cAllocatorAware2]& c_struct):
    return _fbthrift_ctypes.AllocatorAware2._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cTypedefStruct] TypedefStruct_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.TypedefStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object TypedefStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cTypedefStruct]& c_struct):
    return _fbthrift_ctypes.TypedefStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithDoubleUnderscores] StructWithDoubleUnderscores_convert_to_cpp(object inst) except*:
    return (<_fbthrift_ctypes.StructWithDoubleUnderscores?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE

cdef object StructWithDoubleUnderscores_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithDoubleUnderscores]& c_struct):
    return _fbthrift_ctypes.StructWithDoubleUnderscores._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

