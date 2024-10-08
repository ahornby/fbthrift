/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/basic/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#pragma once

#include <functional>
#include <folly/Range.h>

#include "thrift/compiler/test/fixtures/basic/gen-cpp2/module_data.h"
#include "thrift/compiler/test/fixtures/basic/gen-cpp2/module_types.h"
#include "thrift/compiler/test/fixtures/basic/gen-cpp2/module_metadata.h"
namespace thrift {
namespace py3 {



template<>
inline void reset_field<::test::fixtures::basic::MyStruct>(
    ::test::fixtures::basic::MyStruct& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.MyIntField_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().MyIntField_ref());
      return;
    case 1:
      obj.MyStringField_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().MyStringField_ref());
      return;
    case 2:
      obj.MyDataField_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().MyDataField_ref());
      return;
    case 3:
      obj.myEnum_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().myEnum_ref());
      return;
    case 4:
      obj.oneway_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().oneway_ref());
      return;
    case 5:
      obj.readonly_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().readonly_ref());
      return;
    case 6:
      obj.idempotent_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().idempotent_ref());
      return;
    case 7:
      obj.floatSet_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().floatSet_ref());
      return;
    case 8:
      obj.no_hack_codegen_field_ref().copy_from(default_inst<::test::fixtures::basic::MyStruct>().no_hack_codegen_field_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::basic::Containers>(
    ::test::fixtures::basic::Containers& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.I32List_ref().copy_from(default_inst<::test::fixtures::basic::Containers>().I32List_ref());
      return;
    case 1:
      obj.StringSet_ref().copy_from(default_inst<::test::fixtures::basic::Containers>().StringSet_ref());
      return;
    case 2:
      obj.StringToI64Map_ref().copy_from(default_inst<::test::fixtures::basic::Containers>().StringToI64Map_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::basic::MyDataItem>(
    ::test::fixtures::basic::MyDataItem& obj, uint16_t index) {
  switch (index) {
  }
}

template<>
inline void reset_field<::test::fixtures::basic::MyException>(
    ::test::fixtures::basic::MyException& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.MyIntField_ref().copy_from(default_inst<::test::fixtures::basic::MyException>().MyIntField_ref());
      return;
    case 1:
      obj.MyStringField_ref().copy_from(default_inst<::test::fixtures::basic::MyException>().MyStringField_ref());
      return;
    case 2:
      obj.myStruct_ref().copy_from(default_inst<::test::fixtures::basic::MyException>().myStruct_ref());
      return;
    case 3:
      obj.myUnion_ref().copy_from(default_inst<::test::fixtures::basic::MyException>().myUnion_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::basic::MyExceptionWithMessage>(
    ::test::fixtures::basic::MyExceptionWithMessage& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.MyIntField_ref().copy_from(default_inst<::test::fixtures::basic::MyExceptionWithMessage>().MyIntField_ref());
      return;
    case 1:
      obj.MyStringField_ref().copy_from(default_inst<::test::fixtures::basic::MyExceptionWithMessage>().MyStringField_ref());
      return;
    case 2:
      obj.myStruct_ref().copy_from(default_inst<::test::fixtures::basic::MyExceptionWithMessage>().myStruct_ref());
      return;
    case 3:
      obj.myUnion_ref().copy_from(default_inst<::test::fixtures::basic::MyExceptionWithMessage>().myUnion_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::basic::ReservedKeyword>(
    ::test::fixtures::basic::ReservedKeyword& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.reserved_field_ref().copy_from(default_inst<::test::fixtures::basic::ReservedKeyword>().reserved_field_ref());
      return;
  }
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::MyStruct>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::Containers>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::MyDataItem>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::MyUnion>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::MyException>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::MyExceptionWithMessage>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::ReservedKeyword>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::basic::UnionToBeRenamed>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}
} // namespace py3
} // namespace thrift
