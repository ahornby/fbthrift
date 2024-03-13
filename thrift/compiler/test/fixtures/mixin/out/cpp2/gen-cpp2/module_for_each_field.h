/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/mixin/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include "thrift/compiler/test/fixtures/mixin/gen-cpp2/module_metadata.h"
#include <thrift/lib/cpp2/visitation/for_each.h>

namespace apache {
namespace thrift {
namespace detail {

template <>
struct ForEachField<::cpp2::Mixin1> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field1_ref()...);
  }
};

template <>
struct ForEachField<::cpp2::Mixin2> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).m1_ref()...);
    f(1, static_cast<T&&>(t).field2_ref()...);
  }
};

template <>
struct ForEachField<::cpp2::Mixin3Base> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field3_ref()...);
  }
};

template <>
struct ForEachField<::cpp2::Foo> {
  template <typename F, typename... T>
  void operator()([[maybe_unused]] F&& f, [[maybe_unused]] T&&... t) const {
    f(0, static_cast<T&&>(t).field4_ref()...);
    f(1, static_cast<T&&>(t).m2_ref()...);
    f(2, static_cast<T&&>(t).m3_ref()...);
  }
};
} // namespace detail
} // namespace thrift
} // namespace apache