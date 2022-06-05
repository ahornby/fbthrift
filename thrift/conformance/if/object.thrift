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

// TODO(ytj): remove this file in favor of "thrift/lib/thrift/object.thrift"
include "thrift/lib/thrift/object.thrift"

namespace cpp2 apache.thrift.conformance
namespace php apache_thrift
namespace py thrift.conformance.object
namespace py.asyncio thrift_asyncio.conformance.object
namespace py3 thrift.conformance
namespace java.swift org.apache.thrift.conformance
namespace java2 org.apache.thrift.conformance
namespace go thrift.conformance.object

typedef object.ProtocolObject Object
typedef object.ProtocolValue Value
