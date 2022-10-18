/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/mcpp2-compare/src/reflection.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */

#include "thrift/compiler/test/fixtures/mcpp2-compare/gen-cpp2/reflection_data.h"

#include <thrift/lib/cpp2/gen/module_data_cpp.h>

#if defined(__GNUC__) && defined(__linux__) && !FOLLY_MOBILE
// This attribute is applied to the static data members to ensure that they are
// not stripped from the compiled binary, in order to keep them available for
// use by debuggers at runtime.
//
// The attribute works by forcing all of the data members (both used and unused
// ones) into the same section. This stops the linker from stripping the unused
// data, as it works on a per-section basis and only removes sections if they
// are entirely unused.
#define THRIFT_DATA_SECTION [[gnu::section(".rodata.thrift.data")]]
#else
#define THRIFT_DATA_SECTION
#endif

namespace apache {
namespace thrift {

THRIFT_DATA_SECTION const std::array<folly::StringPiece, 1> TStructDataStorage<::cpp2::ReflectionStruct>::fields_names = {{
  "fieldA",
}};
THRIFT_DATA_SECTION const std::array<int16_t, 1> TStructDataStorage<::cpp2::ReflectionStruct>::fields_ids = {{
  1,
}};
THRIFT_DATA_SECTION const std::array<protocol::TType, 1> TStructDataStorage<::cpp2::ReflectionStruct>::fields_types = {{
  TType::T_I32,
}};
THRIFT_DATA_SECTION const std::array<folly::StringPiece, 1> TStructDataStorage<::cpp2::ReflectionStruct>::storage_names = {{
  "__fbthrift_field_fieldA",
}};
THRIFT_DATA_SECTION const std::array<int, 1> TStructDataStorage<::cpp2::ReflectionStruct>::isset_indexes = {{
  0,
}};

} // namespace thrift
} // namespace apache
