/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#pragma once

#include <functional>
#include <folly/Range.h>

#include "thrift/compiler/test/fixtures/basic-stack-arguments/gen-cpp2/module_data.h"
#include "thrift/compiler/test/fixtures/basic-stack-arguments/gen-cpp2/module_types.h"
#include "thrift/compiler/test/fixtures/basic-stack-arguments/gen-cpp2/module_metadata.h"
namespace thrift {
namespace py3 {



template<>
inline void reset_field<::cpp2::MyStruct>(
    ::cpp2::MyStruct& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.MyIntField_ref().copy_from(default_inst<::cpp2::MyStruct>().MyIntField_ref());
      return;
    case 1:
      obj.MyStringField_ref().copy_from(default_inst<::cpp2::MyStruct>().MyStringField_ref());
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
} // namespace py3
} // namespace thrift
