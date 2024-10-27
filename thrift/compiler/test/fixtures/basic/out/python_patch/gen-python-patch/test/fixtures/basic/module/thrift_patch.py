# @generated
from __future__ import annotations

import typing
import typing as _typing

from common.thrift.patch.detail.dynamic_patch import (
    BaseStructPatch,
    ListPatch,
    SetPatch,
    MapPatch,
    OptionalFieldPatch,
)

from common.thrift.patch.detail.py_bindings.DynamicPatch import (
    BoolPatch,
    BytePatch,
    I16Patch,
    I32Patch,
    I64Patch,
    FloatPatch,
    DoublePatch,
    StringPatch,
    BinaryPatch,
)

import thrift.python.types as _fbthrift_python_types
import folly.iobuf as _fbthrift_iobuf

import test.fixtures.basic.module.thrift_types



class MyStructPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.MyStruct]):
    pass

    @property
    def MyIntField(self) -> OptionalFieldPatch[
            int,
            I64Patch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_i64_patch(),
            self._patch,
            1,
            _fbthrift_python_types.typeinfo_i64)
    @property
    def MyStringField(self) -> OptionalFieldPatch[
            str,
            StringPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_string_patch(),
            self._patch,
            2,
            _fbthrift_python_types.typeinfo_string)
    @property
    def MyDataField(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyDataItem,
            MyDataItemPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyDataItemPatch(patch),
            self._patch,
            3,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyDataItem))
    @property
    def myEnum(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyEnum,
            test.fixtures.basic.module.thrift_types.MyEnum]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_enum_patch(),
            self._patch,
            4,
            _fbthrift_python_types.EnumTypeInfo(test.fixtures.basic.module.thrift_types.MyEnum))
    @property
    def oneway(self) -> OptionalFieldPatch[
            bool,
            BoolPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_bool_patch(),
            self._patch,
            5,
            _fbthrift_python_types.typeinfo_bool)
    @property
    def readonly(self) -> OptionalFieldPatch[
            bool,
            BoolPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_bool_patch(),
            self._patch,
            6,
            _fbthrift_python_types.typeinfo_bool)
    @property
    def idempotent(self) -> OptionalFieldPatch[
            bool,
            BoolPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_bool_patch(),
            self._patch,
            7,
            _fbthrift_python_types.typeinfo_bool)
    @property
    def floatSet(self) -> OptionalFieldPatch[
            _typing.AbstractSet[float],
            SetPatch[float]]:

        return OptionalFieldPatch(
            lambda patch, type_info: SetPatch(patch.as_set_patch(), type_info),
            self._patch,
            8,
            _fbthrift_python_types.SetTypeInfo(_fbthrift_python_types.typeinfo_float))
    @property
    def no_hack_codegen_field(self) -> OptionalFieldPatch[
            str,
            StringPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_string_patch(),
            self._patch,
            9,
            _fbthrift_python_types.typeinfo_string)

class ContainersPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.Containers]):
    pass

    @property
    def I32List(self) -> OptionalFieldPatch[
            _typing.Sequence[int],
            ListPatch[int]]:

        return OptionalFieldPatch(
            lambda patch, type_info: ListPatch(patch.as_list_patch(), type_info),
            self._patch,
            1,
            _fbthrift_python_types.ListTypeInfo(_fbthrift_python_types.typeinfo_i32))
    @property
    def StringSet(self) -> OptionalFieldPatch[
            _typing.AbstractSet[str],
            SetPatch[str]]:

        return OptionalFieldPatch(
            lambda patch, type_info: SetPatch(patch.as_set_patch(), type_info),
            self._patch,
            2,
            _fbthrift_python_types.SetTypeInfo(_fbthrift_python_types.typeinfo_string))
    @property
    def StringToI64Map(self) -> OptionalFieldPatch[
            _typing.Mapping[str, int],
            MapPatch[str, I64Patch]]:

        return OptionalFieldPatch(
            lambda patch, type_info: MapPatch(lambda patch, type_info: patch.as_i64_patch(), patch.as_map_patch(), type_info),
            self._patch,
            3,
            _fbthrift_python_types.MapTypeInfo(_fbthrift_python_types.typeinfo_string, _fbthrift_python_types.typeinfo_i64))

class MyDataItemPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.MyDataItem]):
    pass


class MyUnionPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.MyUnion]):
    pass

    @property
    def myEnum(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyEnum,
            test.fixtures.basic.module.thrift_types.MyEnum]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_enum_patch(),
            self._patch,
            1,
            _fbthrift_python_types.EnumTypeInfo(test.fixtures.basic.module.thrift_types.MyEnum))
    @property
    def myStruct(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyStruct,
            MyStructPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyStructPatch(patch),
            self._patch,
            2,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyStruct))
    @property
    def myDataItem(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyDataItem,
            MyDataItemPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyDataItemPatch(patch),
            self._patch,
            3,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyDataItem))
    @property
    def floatSet(self) -> OptionalFieldPatch[
            _typing.AbstractSet[float],
            SetPatch[float]]:

        return OptionalFieldPatch(
            lambda patch, type_info: SetPatch(patch.as_set_patch(), type_info),
            self._patch,
            4,
            _fbthrift_python_types.SetTypeInfo(_fbthrift_python_types.typeinfo_float))

class MyExceptionPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.MyException]):
    pass

    @property
    def MyIntField(self) -> OptionalFieldPatch[
            int,
            I64Patch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_i64_patch(),
            self._patch,
            1,
            _fbthrift_python_types.typeinfo_i64)
    @property
    def MyStringField(self) -> OptionalFieldPatch[
            str,
            StringPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_string_patch(),
            self._patch,
            2,
            _fbthrift_python_types.typeinfo_string)
    @property
    def myStruct(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyStruct,
            MyStructPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyStructPatch(patch),
            self._patch,
            3,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyStruct))
    @property
    def myUnion(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyUnion,
            MyUnionPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyUnionPatch(patch),
            self._patch,
            4,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyUnion))

class MyExceptionWithMessagePatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.MyExceptionWithMessage]):
    pass

    @property
    def MyIntField(self) -> OptionalFieldPatch[
            int,
            I64Patch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_i64_patch(),
            self._patch,
            1,
            _fbthrift_python_types.typeinfo_i64)
    @property
    def MyStringField(self) -> OptionalFieldPatch[
            str,
            StringPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_string_patch(),
            self._patch,
            2,
            _fbthrift_python_types.typeinfo_string)
    @property
    def myStruct(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyStruct,
            MyStructPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyStructPatch(patch),
            self._patch,
            3,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyStruct))
    @property
    def myUnion(self) -> OptionalFieldPatch[
            test.fixtures.basic.module.thrift_types.MyUnion,
            MyUnionPatch]:

        return OptionalFieldPatch(
            lambda patch, type_info: MyUnionPatch(patch),
            self._patch,
            4,
            _fbthrift_python_types.StructTypeInfo(test.fixtures.basic.module.thrift_types.MyUnion))

class ReservedKeywordPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.ReservedKeyword]):
    pass

    @property
    def reserved_field(self) -> OptionalFieldPatch[
            int,
            I32Patch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_i32_patch(),
            self._patch,
            1,
            _fbthrift_python_types.typeinfo_i32)

class UnionToBeRenamedPatch(BaseStructPatch[test.fixtures.basic.module.thrift_types.UnionToBeRenamed]):
    pass

    @property
    def reserved_field(self) -> OptionalFieldPatch[
            int,
            I32Patch]:

        return OptionalFieldPatch(
            lambda patch, type_info: patch.as_i32_patch(),
            self._patch,
            1,
            _fbthrift_python_types.typeinfo_i32)
