# Copyright (c) Meta Platforms, Inc. and affiliates.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import unittest

from parameterized import parameterized_class

from schema_evolution_test.thrift_types import (
    MyStruct_V1,
    MyStruct_V2,
    MyUnion_V1,
    MyUnion_V2,
)
from thrift.python.protocol import Protocol

from thrift.python.serializer import deserialize, serialize
from thrift.python.types import StructOrUnion


@parameterized_class(
    ("protocol"),
    [
        (Protocol.BINARY,),
        (Protocol.COMPACT,),
    ],
)
class TestSchemaCompatibility(unittest.TestCase):
    def serialize_deserialize_round_trip(
        self, struct_or_union: StructOrUnion, protocol: Protocol
    ) -> None:
        serialized = serialize(struct_or_union, protocol=protocol)
        struct_or_union_v2 = deserialize(
            type(struct_or_union), serialized, protocol=protocol
        )
        self.assertEqual(struct_or_union, struct_or_union_v2)

    def test_union_compatibility_sanity_test(self) -> None:
        # Just a sanity check to ensure that each type can be serialized and
        # deserialized correctly
        self.serialize_deserialize_round_trip(MyUnion_V1(), self.protocol)
        self.serialize_deserialize_round_trip(MyUnion_V2(), self.protocol)
        self.serialize_deserialize_round_trip(MyStruct_V1(), self.protocol)
        self.serialize_deserialize_round_trip(MyStruct_V2(), self.protocol)

    def test_union_compatibility(self) -> None:
        u_v2 = MyUnion_V2(string_field="Hello world!")
        serialized = serialize(u_v2, protocol=self.protocol)

        # We should be able to deserialize the v2 union to v1 union, even if the
        # v2 union has a new field that is not present in the v1 union. In this
        # case, we expect the deserialized v1 union to be empty.
        u_v1 = deserialize(MyUnion_V1, serialized, protocol=self.protocol)
        self.assertEqual(MyUnion_V1(), u_v1)

    def test_struct_compatibility(self) -> None:
        s_v2 = MyStruct_V2(
            i32_field=11, union_field=MyUnion_V2(string_field="Hello world!")
        )
        serialized = serialize(s_v2, protocol=self.protocol)

        # We should be able to deserialize the v2 struct to v1 struct, even if
        # the v2 union has a new field that is not present in the v1 union. In
        # this case, we expect the deserialized v1 struct to have an empty v1
        # union.
        s_v1 = deserialize(MyStruct_V1, serialized, protocol=self.protocol)
        self.assertEqual(11, s_v1.i32_field)
        self.assertEqual(MyUnion_V1(), s_v1.union_field)
