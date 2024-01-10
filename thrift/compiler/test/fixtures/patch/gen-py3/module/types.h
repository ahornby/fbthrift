/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#pragma once

#include <functional>
#include <folly/Range.h>

#include <thrift/lib/py3/enums.h>
#include "thrift/compiler/test/fixtures/patch/src/gen-cpp2/module_data.h"
#include "thrift/compiler/test/fixtures/patch/src/gen-cpp2/module_types.h"
#include "thrift/compiler/test/fixtures/patch/src/gen-cpp2/module_metadata.h"
namespace thrift {
namespace py3 {


template<>
inline const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::test::fixtures::patch::MyEnum>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}


template<>
inline const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::test::fixtures::patch::InnerUnion::Type>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}

template<>
inline const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::test::fixtures::patch::MyUnion::Type>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}


template<>
inline void reset_field<::test::fixtures::patch::MyData>(
    ::test::fixtures::patch::MyData& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.data1_ref().copy_from(default_inst<::test::fixtures::patch::MyData>().data1_ref());
      return;
    case 1:
      obj.data2_ref().copy_from(default_inst<::test::fixtures::patch::MyData>().data2_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::patch::MyDataWithCustomDefault>(
    ::test::fixtures::patch::MyDataWithCustomDefault& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.data1_ref().copy_from(default_inst<::test::fixtures::patch::MyDataWithCustomDefault>().data1_ref());
      return;
    case 1:
      obj.data2_ref().copy_from(default_inst<::test::fixtures::patch::MyDataWithCustomDefault>().data2_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::patch::MyStruct>(
    ::test::fixtures::patch::MyStruct& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.boolVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().boolVal_ref());
      return;
    case 1:
      obj.byteVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().byteVal_ref());
      return;
    case 2:
      obj.i16Val_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().i16Val_ref());
      return;
    case 3:
      obj.i32Val_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().i32Val_ref());
      return;
    case 4:
      obj.i64Val_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().i64Val_ref());
      return;
    case 5:
      obj.floatVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().floatVal_ref());
      return;
    case 6:
      obj.doubleVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().doubleVal_ref());
      return;
    case 7:
      obj.stringVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().stringVal_ref());
      return;
    case 8:
      obj.binaryVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().binaryVal_ref());
      return;
    case 9:
      obj.enumVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().enumVal_ref());
      return;
    case 10:
      obj.structVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().structVal_ref());
      return;
    case 11:
      obj.unionVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().unionVal_ref());
      return;
    case 12:
      obj.lateStructVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().lateStructVal_ref());
      return;
    case 13:
      obj.optBoolVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optBoolVal_ref());
      return;
    case 14:
      obj.optByteVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optByteVal_ref());
      return;
    case 15:
      obj.optI16Val_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optI16Val_ref());
      return;
    case 16:
      obj.optI32Val_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optI32Val_ref());
      return;
    case 17:
      obj.optI64Val_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optI64Val_ref());
      return;
    case 18:
      obj.optFloatVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optFloatVal_ref());
      return;
    case 19:
      obj.optDoubleVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optDoubleVal_ref());
      return;
    case 20:
      obj.optStringVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optStringVal_ref());
      return;
    case 21:
      obj.optBinaryVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optBinaryVal_ref());
      return;
    case 22:
      obj.optEnumVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optEnumVal_ref());
      return;
    case 23:
      obj.optStructVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optStructVal_ref());
      return;
    case 24:
      obj.optLateStructVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optLateStructVal_ref());
      return;
    case 25:
      obj.optListVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optListVal_ref());
      return;
    case 26:
      obj.optSetVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optSetVal_ref());
      return;
    case 27:
      obj.optMapVal_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().optMapVal_ref());
      return;
    case 28:
      obj.listMap_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().listMap_ref());
      return;
    case 29:
      obj.mapMap_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().mapMap_ref());
      return;
    case 30:
      obj.i32WithCustomDefault_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().i32WithCustomDefault_ref());
      return;
    case 31:
      obj.structWithCustomDefault_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().structWithCustomDefault_ref());
      return;
    case 32:
      obj.structWithFieldCustomDefault_ref().copy_from(default_inst<::test::fixtures::patch::MyStruct>().structWithFieldCustomDefault_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::patch::LateDefStruct>(
    ::test::fixtures::patch::LateDefStruct& obj, uint16_t index) {
  switch (index) {
  }
}

template<>
inline void reset_field<::test::fixtures::patch::Recursive>(
    ::test::fixtures::patch::Recursive& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.nodes_ref().copy_from(default_inst<::test::fixtures::patch::Recursive>().nodes_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::patch::Bar>(
    ::test::fixtures::patch::Bar& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.loop_ref().reset();
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::patch::Loop>(
    ::test::fixtures::patch::Loop& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.bar_ref().copy_from(default_inst<::test::fixtures::patch::Loop>().bar_ref());
      return;
  }
}

template<>
inline void reset_field<::test::fixtures::patch::RefFields>(
    ::test::fixtures::patch::RefFields& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.unique_ref().reset();
      return;
    case 1:
      obj.shared_const_ref().reset();
      return;
    case 2:
      obj.shared_mustable_ref().reset();
      return;
    case 3:
      obj.opt_unique_ref().reset();
      return;
    case 4:
      obj.opt_shared_const_ref().reset();
      return;
    case 5:
      obj.opt_shared_mustable_ref().reset();
      return;
    case 6:
      obj.opt_box_ref().copy_from(default_inst<::test::fixtures::patch::RefFields>().opt_box_ref());
      return;
  }
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::MyData>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::MyDataWithCustomDefault>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::InnerUnion>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::MyUnion>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::MyStruct>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::LateDefStruct>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::Recursive>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::Bar>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::Loop>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
inline const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::test::fixtures::patch::RefFields>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}
} // namespace py3
} // namespace thrift
