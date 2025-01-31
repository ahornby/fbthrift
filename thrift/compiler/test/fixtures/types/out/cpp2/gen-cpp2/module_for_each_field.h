/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/types/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include "thrift/compiler/test/fixtures/types/gen-cpp2/module_metadata.h"
#include <thrift/lib/cpp2/visitation/for_each.h>

namespace apache {
namespace thrift {
namespace detail {

template <>
struct ForEachField<::apache::thrift::fixtures::types::empty_struct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::decorated_struct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::ContainerStruct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).fieldA_ref()...);
    f(1, static_cast<T&&>(t).fieldB_ref()...);
    f(2, static_cast<T&&>(t).fieldC_ref()...);
    f(3, static_cast<T&&>(t).fieldD_ref()...);
    f(4, static_cast<T&&>(t).fieldE_ref()...);
    f(5, static_cast<T&&>(t).fieldF_ref()...);
    f(6, static_cast<T&&>(t).fieldG_ref()...);
    f(7, static_cast<T&&>(t).fieldH_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::CppTypeStruct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).fieldA_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::VirtualStruct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).MyIntField_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::MyStructWithForwardRefEnum> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).a_ref()...);
    f(1, static_cast<T&&>(t).b_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::TrivialNumeric> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).a_ref()...);
    f(1, static_cast<T&&>(t).b_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::TrivialNestedWithDefault> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).z_ref()...);
    f(1, static_cast<T&&>(t).n_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::ComplexString> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).a_ref()...);
    f(1, static_cast<T&&>(t).b_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::ComplexNestedWithDefault> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).z_ref()...);
    f(1, static_cast<T&&>(t).n_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::MinPadding> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).small_ref()...);
    f(1, static_cast<T&&>(t).big_ref()...);
    f(2, static_cast<T&&>(t).medium_ref()...);
    f(3, static_cast<T&&>(t).biggish_ref()...);
    f(4, static_cast<T&&>(t).tiny_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::MinPaddingWithCustomType> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).small_ref()...);
    f(1, static_cast<T&&>(t).biggish_ref()...);
    f(2, static_cast<T&&>(t).medium_ref()...);
    f(3, static_cast<T&&>(t).big_ref()...);
    f(4, static_cast<T&&>(t).tiny_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::MyStruct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).MyIntField_ref()...);
    f(1, static_cast<T&&>(t).MyStringField_ref()...);
    f(2, static_cast<T&&>(t).majorVer_ref()...);
    f(3, static_cast<T&&>(t).data_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::MyDataItem> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::Renamed> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).bar_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::AnnotatedTypes> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).binary_field_ref()...);
    f(1, static_cast<T&&>(t).list_field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::ForwardUsageRoot> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).ForwardUsageStruct_ref()...);
    f(1, static_cast<T&&>(t).ForwardUsageByRef_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::ForwardUsageStruct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).foo_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::ForwardUsageByRef> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).foo_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::IncompleteMap> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::IncompleteMapDep> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::CompleteMap> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::CompleteMapDep> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::IncompleteList> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::IncompleteListDep> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::CompleteList> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::CompleteListDep> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::AdaptedList> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::detail::AdaptedListDep> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::DependentAdaptedList> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::detail::DependentAdaptedListDep> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::AllocatorAware> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).aa_list_ref()...);
    f(1, static_cast<T&&>(t).aa_set_ref()...);
    f(2, static_cast<T&&>(t).aa_map_ref()...);
    f(3, static_cast<T&&>(t).aa_string_ref()...);
    f(4, static_cast<T&&>(t).not_a_container_ref()...);
    f(5, static_cast<T&&>(t).aa_unique_ref()...);
    f(6, static_cast<T&&>(t).aa_shared_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::AllocatorAware2> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).not_a_container_ref()...);
    f(1, static_cast<T&&>(t).box_field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::TypedefStruct> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).i32_field_ref()...);
    f(1, static_cast<T&&>(t).IntTypedef_field_ref()...);
    f(2, static_cast<T&&>(t).UintTypedef_field_ref()...);
  }
};

template <>
struct ForEachField<::apache::thrift::fixtures::types::StructWithDoubleUnderscores> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).__field_ref()...);
  }
};
} // namespace detail
} // namespace thrift
} // namespace apache
