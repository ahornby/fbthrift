#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

from abc import ABCMeta
import typing as _typing

import folly.iobuf as _fbthrift_iobuf

import apache.thrift.metadata.thrift_types as _fbthrift_metadata
from thrift.python.serializer import serialize_iobuf, deserialize, Protocol
from thrift.python.server import ServiceInterface, RpcKind, PythonUserException

import my.namespacing.test.module.module.thrift_types
import my.namespacing.test.module.module.thrift_metadata

class TestServiceInterface(
    ServiceInterface,
    metaclass=ABCMeta
):

    @staticmethod
    def service_name() -> bytes:
        return b"TestService"

    def getFunctionTable(self) -> _typing.Mapping[bytes, _typing.Callable[..., object]]:
        functionTable = {
            b"init": (RpcKind.SINGLE_REQUEST_SINGLE_RESPONSE, self._fbthrift__handler_init),
        }
        return {**super().getFunctionTable(), **functionTable}

    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.TestService"

    @staticmethod
    def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
        return my.namespacing.test.module.module.thrift_metadata.gen_metadata_service_TestService()

    @staticmethod
    def __get_metadata_service_response__() -> _fbthrift_metadata.ThriftServiceMetadataResponse:
        return my.namespacing.test.module.module.thrift_metadata._fbthrift_metadata_service_response_TestService()



    async def init(
            self,
            int1: int
        ) -> int:
        raise NotImplementedError("async def init is not implemented")

    async def _fbthrift__handler_init(self, args: _fbthrift_iobuf.IOBuf, protocol: Protocol) -> _fbthrift_iobuf.IOBuf:
        args_struct = deserialize(my.namespacing.test.module.module.thrift_types._fbthrift_TestService_init_args, args, protocol)
        value = await self.init(args_struct.int1,)
        return_struct = my.namespacing.test.module.module.thrift_types._fbthrift_TestService_init_result(success=value)
        return serialize_iobuf(return_struct, protocol)

