/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/includes/src/matching_struct_names.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#pragma once

#include <functional>
#include <folly/Range.h>

#include "thrift/compiler/test/fixtures/includes/gen-cpp2/matching_struct_names_data.h"
#include "thrift/compiler/test/fixtures/includes/gen-cpp2/matching_struct_names_types.h"
#include "thrift/compiler/test/fixtures/includes/gen-cpp2/matching_struct_names_metadata.h"
namespace thrift {
namespace py3 {



template<>
inline void reset_field<::cpp2::MyStruct>(
    ::cpp2::MyStruct& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.field_ref().copy_from(default_inst<::cpp2::MyStruct>().field_ref());
      return;
  }
}

template<>
inline void reset_field<::cpp2::Combo>(
    ::cpp2::Combo& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.listOfOurMyStructLists_ref().copy_from(default_inst<::cpp2::Combo>().listOfOurMyStructLists_ref());
      return;
    case 1:
      obj.theirMyStructList_ref().copy_from(default_inst<::cpp2::Combo>().theirMyStructList_ref());
      return;
    case 2:
      obj.ourMyStructList_ref().copy_from(default_inst<::cpp2::Combo>().ourMyStructList_ref());
      return;
    case 3:
      obj.listOfTheirMyStructList_ref().copy_from(default_inst<::cpp2::Combo>().listOfTheirMyStructList_ref());
      return;
  }
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::cpp2::MyStruct>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::cpp2::Combo>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}
} // namespace py3
} // namespace thrift
