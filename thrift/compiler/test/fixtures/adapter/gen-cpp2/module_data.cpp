/**
 * Autogenerated by Thrift for src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/adapter/gen-cpp2/module_data.h"

#include <thrift/lib/cpp2/gen/module_data_cpp.h>

namespace apache {
namespace thrift {

const std::array<folly::StringPiece, 7> TStructDataStorage<::cpp2::Foo>::fields_names = {{
  "intField",
  "optionalIntField",
  "intFieldWithDefault",
  "setField",
  "optionalSetField",
  "mapField",
  "optionalMapField",
}};
const std::array<int16_t, 7> TStructDataStorage<::cpp2::Foo>::fields_ids = {{
  1,
  2,
  3,
  4,
  5,
  6,
  7,
}};
const std::array<protocol::TType, 7> TStructDataStorage<::cpp2::Foo>::fields_types = {{
  TType::T_I32,
  TType::T_I32,
  TType::T_I32,
  TType::T_SET,
  TType::T_SET,
  TType::T_MAP,
  TType::T_MAP,
}};

const std::array<folly::StringPiece, 4> TStructDataStorage<::cpp2::Bar>::fields_names = {{
  "structField",
  "optionalStructField",
  "structListField",
  "optionalStructListField",
}};
const std::array<int16_t, 4> TStructDataStorage<::cpp2::Bar>::fields_ids = {{
  1,
  2,
  3,
  4,
}};
const std::array<protocol::TType, 4> TStructDataStorage<::cpp2::Bar>::fields_types = {{
  TType::T_STRUCT,
  TType::T_STRUCT,
  TType::T_LIST,
  TType::T_LIST,
}};

} // namespace thrift
} // namespace apache
