#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import typing as _typing

import apache.thrift.metadata.thrift_types as _fbthrift_metadata
import folly.iobuf as _fbthrift_iobuf
from thrift.python.client import (
    AsyncClient as _fbthrift_python_AsyncClient,
    SyncClient as _fbthrift_python_SyncClient,
    Client as _fbthrift_python_Client,
)
from thrift.python.client.omni_client import InteractionMethodPosition as _fbthrift_InteractionMethodPosition, FunctionQualifier as _fbthrift_FunctionQualifier
from thrift.python.common import RpcOptions
import thrift.python.mutable_exceptions as _fbthrift_python_mutable_exceptions
import thrift.python.mutable_types as _fbthrift_python_mutable_types
import thrift.python.exceptions as _fbthrift_python_exceptions
import thrift.python.types as _fbthrift_python_types
import apache.thrift.fixtures.types.module.thrift_mutable_types
import apache.thrift.fixtures.types.module.thrift_metadata
import apache.thrift.fixtures.types.included.thrift_mutable_types

class SomeService(_fbthrift_python_Client["SomeService.Async", "SomeService.Sync"]):
    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.SomeService"
    
    @staticmethod
    def __get_thrift_uri__() -> _typing.Optional[str]:
        return "apache.org/thrift/fixtures/types/SomeService"
    
    @staticmethod
    def __get_thrift_unstructured_annotations_DEPRECATED__() -> _typing.Mapping[str, str]:
        return {
        }
    
    @staticmethod
    def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
        return apache.thrift.fixtures.types.module.thrift_metadata.gen_metadata_service_SomeService()
    
    class Async(_fbthrift_python_AsyncClient):
        @staticmethod
        def __get_thrift_name__() -> str:
            return "module.SomeService"
    
        @staticmethod
        def __get_thrift_uri__() -> _typing.Optional[str]:
            return "apache.org/thrift/fixtures/types/SomeService"
    
        @staticmethod
        def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
            return apache.thrift.fixtures.types.module.thrift_metadata.gen_metadata_service_SomeService()
    
        async def bounce_map(
            self,
            m: _typing.Mapping[int, str],
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> _typing.Mapping[int, str]:
            _fbthrift_resp = await self._send_request(
                "SomeService",
                "bounce_map",
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_bounce_map_args(
                    m=m,),
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_bounce_map_result,
                qualifier = _fbthrift_FunctionQualifier.Unspecified,
                uri_or_name="apache.org/thrift/fixtures/types/SomeService",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            # shortcut to success path for non-void returns
            if _fbthrift_resp.success is not None:
                return _fbthrift_resp.success
            raise _fbthrift_python_exceptions.ApplicationError(
                _fbthrift_python_exceptions.ApplicationErrorType.MISSING_RESULT,
                "Empty Response",
            )
    
        async def binary_keyed_map(
            self,
            r: _typing.Sequence[int],
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> _typing.Mapping[bytes, int]:
            _fbthrift_resp = await self._send_request(
                "SomeService",
                "binary_keyed_map",
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_binary_keyed_map_args(
                    r=r,),
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_binary_keyed_map_result,
                qualifier = _fbthrift_FunctionQualifier.Unspecified,
                uri_or_name="apache.org/thrift/fixtures/types/SomeService",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            # shortcut to success path for non-void returns
            if _fbthrift_resp.success is not None:
                return _fbthrift_resp.success
            raise _fbthrift_python_exceptions.ApplicationError(
                _fbthrift_python_exceptions.ApplicationErrorType.MISSING_RESULT,
                "Empty Response",
            )
    
    class Sync(_fbthrift_python_SyncClient):
        @staticmethod
        def __get_thrift_name__() -> str:
            return "module.SomeService"
    
        @staticmethod
        def __get_thrift_uri__() -> _typing.Optional[str]:
            return "apache.org/thrift/fixtures/types/SomeService"
    
        @staticmethod
        def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
            return apache.thrift.fixtures.types.module.thrift_metadata.gen_metadata_service_SomeService()
    
        def bounce_map(
            self,
            m: _typing.Mapping[int, str],
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> _typing.Mapping[int, str]:
            _fbthrift_resp = self._send_request(
                "SomeService",
                "bounce_map",
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_bounce_map_args(
                    m=m,),
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_bounce_map_result,
                uri_or_name="apache.org/thrift/fixtures/types/SomeService",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            # shortcut to success path for non-void returns
            if _fbthrift_resp.success is not None:
                return _fbthrift_resp.success
            raise _fbthrift_python_exceptions.ApplicationError(
                _fbthrift_python_exceptions.ApplicationErrorType.MISSING_RESULT,
                "Empty Response",
            )
    
        def binary_keyed_map(
            self,
            r: _typing.Sequence[int],
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> _typing.Mapping[bytes, int]:
            _fbthrift_resp = self._send_request(
                "SomeService",
                "binary_keyed_map",
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_binary_keyed_map_args(
                    r=r,),
                apache.thrift.fixtures.types.module.thrift_mutable_types._fbthrift_SomeService_binary_keyed_map_result,
                uri_or_name="apache.org/thrift/fixtures/types/SomeService",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            # shortcut to success path for non-void returns
            if _fbthrift_resp.success is not None:
                return _fbthrift_resp.success
            raise _fbthrift_python_exceptions.ApplicationError(
                _fbthrift_python_exceptions.ApplicationErrorType.MISSING_RESULT,
                "Empty Response",
            )

