/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/default_values_rectification_after/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */

#include "thrift/compiler/test/fixtures/default_values_rectification_after/gen-cpp2/module_data.h"

#include <thrift/lib/cpp2/gen/module_data_cpp.h>

FOLLY_CLANG_DISABLE_WARNING("-Wunused-macros")

#if defined(__GNUC__) && defined(__linux__) && !FOLLY_MOBILE
// These attributes are applied to the static data members to ensure that they
// are not stripped from the compiled binary, in order to keep them available
// for use by debuggers at runtime.
//
// The "used" attribute is required to ensure the compiler always emits unused
// data.
//
// The "section" attribute is required to stop the linker from stripping used
// data. It works by forcing all of the data members (both used and unused ones)
// into the same section. As the linker strips data on a per-section basis, it
// is then unable to remove unused data without also removing used data.
// This has a similar effect to the "retain" attribute, but works with older
// toolchains.
#define THRIFT_DATA_MEMBER [[gnu::used]] [[gnu::section(".rodata.thrift.data")]]
#else
#define THRIFT_DATA_MEMBER
#endif

namespace apache {
namespace thrift {

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::EmptyStruct>::name = "EmptyStruct";
THRIFT_DATA_MEMBER const std::array<std::string_view, 0> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::EmptyStruct>::fields_names = { {
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 0> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::EmptyStruct>::fields_ids = { {
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 0> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::EmptyStruct>::fields_types = { {
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 0> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::EmptyStruct>::storage_names = { {
}};
THRIFT_DATA_MEMBER const std::array<int, 0> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::EmptyStruct>::isset_indexes = { {
}};

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::TestStruct>::name = "TestStruct";
THRIFT_DATA_MEMBER const std::array<std::string_view, 8> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::TestStruct>::fields_names = { {
  "unqualified_int_field"sv,
  "unqualified_bool_field"sv,
  "unqualified_list_field"sv,
  "unqualified_struct_field"sv,
  "optional_int_field"sv,
  "optional_bool_field"sv,
  "optional_list_field"sv,
  "optional_struct_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 8> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::TestStruct>::fields_ids = { {
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 8> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::TestStruct>::fields_types = { {
  TType::T_I32,
  TType::T_BOOL,
  TType::T_LIST,
  TType::T_STRUCT,
  TType::T_I32,
  TType::T_BOOL,
  TType::T_LIST,
  TType::T_STRUCT,
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 8> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::TestStruct>::storage_names = { {
  "__fbthrift_field_unqualified_int_field"sv,
  "__fbthrift_field_unqualified_bool_field"sv,
  "__fbthrift_field_unqualified_list_field"sv,
  "__fbthrift_field_unqualified_struct_field"sv,
  "__fbthrift_field_optional_int_field"sv,
  "__fbthrift_field_optional_bool_field"sv,
  "__fbthrift_field_optional_list_field"sv,
  "__fbthrift_field_optional_struct_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int, 8> TStructDataStorage<::facebook::thrift::compiler::test::fixtures::default_values_rectification::TestStruct>::isset_indexes = { {
  0,
  1,
  2,
  3,
  4,
  5,
  6,
  7,
}};

} // namespace thrift
} // namespace apache
