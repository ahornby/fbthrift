/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/terse_write/src/deprecated_terse_write.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */

#include "thrift/compiler/test/fixtures/terse_write/gen-cpp2/deprecated_terse_write_data.h"

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

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyStruct>::name = "MyStruct";
THRIFT_DATA_MEMBER const std::array<std::string_view, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyStruct>::fields_names = { {
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyStruct>::fields_ids = { {
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyStruct>::fields_types = { {
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyStruct>::storage_names = { {
}};
THRIFT_DATA_MEMBER const std::array<int, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyStruct>::isset_indexes = { {
}};

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyUnion>::name = "MyUnion";
THRIFT_DATA_MEMBER const std::array<std::string_view, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyUnion>::fields_names = { {
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyUnion>::fields_ids = { {
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyUnion>::fields_types = { {
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyUnion>::storage_names = { {
}};
THRIFT_DATA_MEMBER const std::array<int, 0> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::MyUnion>::isset_indexes = { {
}};

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::StructLevelTerseStruct>::name = "StructLevelTerseStruct";
THRIFT_DATA_MEMBER const std::array<std::string_view, 15> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::StructLevelTerseStruct>::fields_names = { {
  "bool_field"sv,
  "byte_field"sv,
  "short_field"sv,
  "int_field"sv,
  "long_field"sv,
  "float_field"sv,
  "double_field"sv,
  "string_field"sv,
  "binary_field"sv,
  "enum_field"sv,
  "list_field"sv,
  "set_field"sv,
  "map_field"sv,
  "struct_field"sv,
  "union_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 15> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::StructLevelTerseStruct>::fields_ids = { {
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
  9,
  10,
  11,
  12,
  13,
  14,
  15,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 15> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::StructLevelTerseStruct>::fields_types = { {
  TType::T_BOOL,
  TType::T_BYTE,
  TType::T_I16,
  TType::T_I32,
  TType::T_I64,
  TType::T_FLOAT,
  TType::T_DOUBLE,
  TType::T_STRING,
  TType::T_STRING,
  TType::T_I32,
  TType::T_LIST,
  TType::T_SET,
  TType::T_MAP,
  TType::T_STRUCT,
  TType::T_STRUCT,
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 15> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::StructLevelTerseStruct>::storage_names = { {
  "__fbthrift_field_bool_field"sv,
  "__fbthrift_field_byte_field"sv,
  "__fbthrift_field_short_field"sv,
  "__fbthrift_field_int_field"sv,
  "__fbthrift_field_long_field"sv,
  "__fbthrift_field_float_field"sv,
  "__fbthrift_field_double_field"sv,
  "__fbthrift_field_string_field"sv,
  "__fbthrift_field_binary_field"sv,
  "__fbthrift_field_enum_field"sv,
  "__fbthrift_field_list_field"sv,
  "__fbthrift_field_set_field"sv,
  "__fbthrift_field_map_field"sv,
  "__fbthrift_field_struct_field"sv,
  "__fbthrift_field_union_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int, 15> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::StructLevelTerseStruct>::isset_indexes = { {
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
}};

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::FieldLevelTerseStruct>::name = "FieldLevelTerseStruct";
THRIFT_DATA_MEMBER const std::array<std::string_view, 30> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::FieldLevelTerseStruct>::fields_names = { {
  "terse_bool_field"sv,
  "terse_byte_field"sv,
  "terse_short_field"sv,
  "terse_int_field"sv,
  "terse_long_field"sv,
  "terse_float_field"sv,
  "terse_double_field"sv,
  "terse_string_field"sv,
  "terse_binary_field"sv,
  "terse_enum_field"sv,
  "terse_list_field"sv,
  "terse_set_field"sv,
  "terse_map_field"sv,
  "terse_struct_field"sv,
  "bool_field"sv,
  "byte_field"sv,
  "short_field"sv,
  "int_field"sv,
  "long_field"sv,
  "float_field"sv,
  "double_field"sv,
  "string_field"sv,
  "binary_field"sv,
  "enum_field"sv,
  "list_field"sv,
  "set_field"sv,
  "map_field"sv,
  "struct_field"sv,
  "union_field"sv,
  "iobuf_ptr_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 30> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::FieldLevelTerseStruct>::fields_ids = { {
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
  9,
  10,
  11,
  12,
  13,
  14,
  15,
  16,
  17,
  18,
  19,
  20,
  21,
  22,
  23,
  24,
  25,
  26,
  27,
  28,
  29,
  30,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 30> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::FieldLevelTerseStruct>::fields_types = { {
  TType::T_BOOL,
  TType::T_BYTE,
  TType::T_I16,
  TType::T_I32,
  TType::T_I64,
  TType::T_FLOAT,
  TType::T_DOUBLE,
  TType::T_STRING,
  TType::T_STRING,
  TType::T_I32,
  TType::T_LIST,
  TType::T_SET,
  TType::T_MAP,
  TType::T_STRUCT,
  TType::T_BOOL,
  TType::T_BYTE,
  TType::T_I16,
  TType::T_I32,
  TType::T_I64,
  TType::T_FLOAT,
  TType::T_DOUBLE,
  TType::T_STRING,
  TType::T_STRING,
  TType::T_I32,
  TType::T_LIST,
  TType::T_SET,
  TType::T_MAP,
  TType::T_STRUCT,
  TType::T_STRUCT,
  TType::T_BYTE,
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 30> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::FieldLevelTerseStruct>::storage_names = { {
  "__fbthrift_field_terse_bool_field"sv,
  "__fbthrift_field_terse_byte_field"sv,
  "__fbthrift_field_terse_short_field"sv,
  "__fbthrift_field_terse_int_field"sv,
  "__fbthrift_field_terse_long_field"sv,
  "__fbthrift_field_terse_float_field"sv,
  "__fbthrift_field_terse_double_field"sv,
  "__fbthrift_field_terse_string_field"sv,
  "__fbthrift_field_terse_binary_field"sv,
  "__fbthrift_field_terse_enum_field"sv,
  "__fbthrift_field_terse_list_field"sv,
  "__fbthrift_field_terse_set_field"sv,
  "__fbthrift_field_terse_map_field"sv,
  "__fbthrift_field_terse_struct_field"sv,
  "__fbthrift_field_bool_field"sv,
  "__fbthrift_field_byte_field"sv,
  "__fbthrift_field_short_field"sv,
  "__fbthrift_field_int_field"sv,
  "__fbthrift_field_long_field"sv,
  "__fbthrift_field_float_field"sv,
  "__fbthrift_field_double_field"sv,
  "__fbthrift_field_string_field"sv,
  "__fbthrift_field_binary_field"sv,
  "__fbthrift_field_enum_field"sv,
  "__fbthrift_field_list_field"sv,
  "__fbthrift_field_set_field"sv,
  "__fbthrift_field_map_field"sv,
  "__fbthrift_field_struct_field"sv,
  "__fbthrift_field_union_field"sv,
  "__fbthrift_field_iobuf_ptr_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int, 30> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::FieldLevelTerseStruct>::isset_indexes = { {
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  -1,
  0,
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
  9,
  10,
  11,
  12,
  13,
  14,
  15,
}};

THRIFT_DATA_MEMBER const std::string_view TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::CppRefStructFields>::name = "CppRefStructFields";
THRIFT_DATA_MEMBER const std::array<std::string_view, 2> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::CppRefStructFields>::fields_names = { {
  "primitive_ref_field"sv,
  "struct_ref_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int16_t, 2> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::CppRefStructFields>::fields_ids = { {
  1,
  2,
}};
THRIFT_DATA_MEMBER const std::array<protocol::TType, 2> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::CppRefStructFields>::fields_types = { {
  TType::T_I32,
  TType::T_STRUCT,
}};
THRIFT_DATA_MEMBER const std::array<std::string_view, 2> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::CppRefStructFields>::storage_names = { {
  "__fbthrift_field_primitive_ref_field"sv,
  "__fbthrift_field_struct_ref_field"sv,
}};
THRIFT_DATA_MEMBER const std::array<int, 2> TStructDataStorage<::facebook::thrift::test::terse_write::deprecated::CppRefStructFields>::isset_indexes = { {
  -1,
  -1,
}};

} // namespace thrift
} // namespace apache
