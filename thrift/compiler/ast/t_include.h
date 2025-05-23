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

#pragma once

#include <fmt/core.h>
#include <thrift/compiler/ast/t_node.h>

namespace apache::thrift::compiler {

class t_program;

/**
 * This node models an include statement in the Thrift program.
 */
class t_include : public t_node {
 public:
  t_include(
      t_program* program,
      std::string raw_path,
      std::optional<std::string> alias)
      : program_(program),
        raw_path_(std::move(raw_path)),
        alias_(std::move(alias)) {}

  t_program* get_program() const { return program_; }

  std::string_view raw_path() const { return raw_path_; }

  std::optional<std::string_view> alias() const { return alias_; }

  void set_str_range(source_range rng) { range_ = rng; }
  source_range str_range() const { return range_; }

 private:
  t_program* program_;
  std::string raw_path_;
  std::optional<std::string> alias_;
  source_range range_;
};

} // namespace apache::thrift::compiler
