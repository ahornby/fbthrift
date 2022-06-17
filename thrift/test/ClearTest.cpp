/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <folly/portability/GTest.h>
#include <thrift/lib/cpp2/op/Clear.h>
#include <thrift/test/gen-cpp2/clear_types.h>

namespace apache::thrift::test {
namespace {

template <typename T>
void checkIsDefault(T& obj) {
  EXPECT_EQ(*obj.bool_field(), false);
  EXPECT_EQ(*obj.byte_field(), 0);
  EXPECT_EQ(*obj.short_field(), 0);
  EXPECT_EQ(*obj.int_field(), 0);
  EXPECT_EQ(*obj.long_field(), 0);
  EXPECT_EQ(*obj.float_field(), 0.0);
  EXPECT_EQ(*obj.double_field(), 0.0);
  EXPECT_EQ(*obj.string_field(), "");
  EXPECT_EQ(*obj.binary_field(), "");
  EXPECT_EQ(*obj.enum_field(), MyEnum::ME0);
  EXPECT_TRUE(obj.list_field()->empty());
  EXPECT_TRUE(obj.set_field()->empty());
  EXPECT_TRUE(obj.map_field()->empty());
  EXPECT_EQ(*obj.struct_field()->int_field(), 0);
}

TEST(ClearTest, StructWithNoDefaultStruct) {
  StructWithNoDefaultStruct obj;
  checkIsDefault(obj);
}

TEST(ClearTest, StructWithDefaultStruct) {
  StructWithDefaultStruct obj;
  apache::thrift::clear(obj);
  checkIsDefault(obj);
}

TEST(ClearTest, RefFields) {
  StructWithNoDefaultStruct obj;
  obj.ref_field_ref()->int_field() = 42;

  auto obj2 = obj;
  EXPECT_EQ(*obj.ref_field_ref()->int_field(), 42);
  EXPECT_EQ(*obj2.ref_field_ref()->int_field(), 42);

  apache::thrift::clear(obj);
  EXPECT_EQ(*obj.ref_field_ref()->int_field(), 0);
  EXPECT_EQ(*obj2.ref_field_ref()->int_field(), 42);
}

TEST(AdaptTest, ThriftClearTestStruct) {
  static_assert(!folly::is_detected_v<
                adapt_detail::ClearType,
                AdapterWithContext,
                AdaptedWithContext<int64_t, ThriftClearTestStruct, 1>>);

  auto obj = ThriftClearTestStruct();

  obj.data()->value = 42;
  obj.meta() = "foo";

  EXPECT_EQ(obj.data()->value, 42);
  EXPECT_EQ(obj.data()->fieldId, 1);
  EXPECT_EQ(*obj.data()->meta, "foo");

  apache::thrift::clear(obj);

  EXPECT_EQ(obj.data()->value, 0);
  EXPECT_EQ(obj.data()->fieldId, 1);
  EXPECT_EQ(*obj.data()->meta, "");
}

TEST(AdaptTest, AdapterClearTestStruct) {
  static_assert(folly::is_detected_v<
                adapt_detail::ClearType,
                AdapterWithContextAndClear,
                AdaptedWithContext<int64_t, AdapterClearTestStruct, 1>>);

  auto obj = AdapterClearTestStruct();

  obj.data()->value = 42;
  obj.meta() = "foo";

  EXPECT_EQ(obj.data()->value, 42);
  EXPECT_EQ(obj.data()->fieldId, 1);
  EXPECT_EQ(*obj.data()->meta, "foo");

  apache::thrift::clear(obj);

  EXPECT_EQ(obj.data()->value, 0);
  EXPECT_EQ(obj.data()->fieldId, 1);
  EXPECT_EQ(*obj.data()->meta, "");
}

TEST(AdaptTest, AdapterClearTestStructOpClear) {
  static_assert(folly::is_detected_v<
                adapt_detail::ClearType,
                AdapterWithContextAndClear,
                AdaptedWithContext<int64_t, AdapterClearTestStruct, 1>>);
  using namespace apache::thrift::type;
  using field_type_tag = field<
      adapted<AdapterWithContextAndClear, i64_t>,
      FieldContext<AdapterClearTestStruct, 1>>;

  auto obj = AdapterClearTestStruct();

  obj.data()->value = 42;
  obj.meta() = "foo";

  apache::thrift::op::clear<field_type_tag>(*obj.data(), obj);

  EXPECT_EQ(obj.data()->value, 0);
  EXPECT_EQ(obj.data()->fieldId, 1);
  EXPECT_EQ(*obj.data()->meta, "foo");
}

} // namespace
} // namespace apache::thrift::test
