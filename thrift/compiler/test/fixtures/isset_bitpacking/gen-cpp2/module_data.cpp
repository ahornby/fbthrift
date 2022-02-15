/**
 * Autogenerated by Thrift for src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */

#include "thrift/compiler/test/fixtures/isset_bitpacking/gen-cpp2/module_data.h"

#include <thrift/lib/cpp2/gen/module_data_cpp.h>

namespace apache {
namespace thrift {

const std::array<folly::StringPiece, 4> TStructDataStorage<::cpp2::Default>::fields_names = {{
  "field1",
  "field2",
  "field3",
  "field4",
}};
const std::array<int16_t, 4> TStructDataStorage<::cpp2::Default>::fields_ids = {{
  1,
  2,
  3,
  4,
}};
const std::array<protocol::TType, 4> TStructDataStorage<::cpp2::Default>::fields_types = {{
  TType::T_I32,
  TType::T_I32,
  TType::T_STRING,
  TType::T_DOUBLE,
}};

const std::array<folly::StringPiece, 4> TStructDataStorage<::cpp2::NonAtomic>::fields_names = {{
  "field1",
  "field2",
  "field3",
  "field4",
}};
const std::array<int16_t, 4> TStructDataStorage<::cpp2::NonAtomic>::fields_ids = {{
  1,
  2,
  3,
  4,
}};
const std::array<protocol::TType, 4> TStructDataStorage<::cpp2::NonAtomic>::fields_types = {{
  TType::T_I32,
  TType::T_I32,
  TType::T_STRING,
  TType::T_DOUBLE,
}};

const std::array<folly::StringPiece, 4> TStructDataStorage<::cpp2::Atomic>::fields_names = {{
  "field1",
  "field2",
  "field3",
  "field4",
}};
const std::array<int16_t, 4> TStructDataStorage<::cpp2::Atomic>::fields_ids = {{
  1,
  2,
  3,
  4,
}};
const std::array<protocol::TType, 4> TStructDataStorage<::cpp2::Atomic>::fields_types = {{
  TType::T_I32,
  TType::T_I32,
  TType::T_STRING,
  TType::T_DOUBLE,
}};

} // namespace thrift
} // namespace apache
