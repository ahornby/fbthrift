/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-enum/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include <thrift/lib/cpp2/gen/module_constants_h.h>

#include "thrift/compiler/test/fixtures/basic-enum/gen-cpp2/module_types.h"

namespace test::fixtures::enumstrict {
/** Glean {"file": "thrift/compiler/test/fixtures/basic-enum/src/module.thrift"} */
namespace module_constants {

  /** Glean {"constant": "kOne"} */
  constexpr ::test::fixtures::enumstrict::MyEnum const kOne_ =  ::test::fixtures::enumstrict::MyEnum::ONE;
  /** Glean {"constant": "kOne"} */
  constexpr ::test::fixtures::enumstrict::MyEnum kOne() {
    return kOne_;
  }

  /** Glean {"constant": "enumNames"} */
  ::std::map<::test::fixtures::enumstrict::MyEnum, ::std::string> const& enumNames();

  FOLLY_EXPORT ::std::string_view _fbthrift_schema_e97c3188d96b91fe();
  FOLLY_EXPORT ::folly::Range<const ::std::string_view*> _fbthrift_schema_e97c3188d96b91fe_includes();

} // namespace module_constants
} // namespace test::fixtures::enumstrict
