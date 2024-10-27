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
import module.thrift_mutable_types
import module.thrift_metadata

class Raiser(_fbthrift_python_Client["Raiser.Async", "Raiser.Sync"]):
    @staticmethod
    def __get_thrift_name__() -> str:
        return "module.Raiser"
    
    @staticmethod
    def __get_thrift_uri__() -> _typing.Optional[str]:
        return None
    
    @staticmethod
    def __get_thrift_unstructured_annotations_DEPRECATED__() -> _typing.Mapping[str, str]:
        return {
        }
    
    @staticmethod
    def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
        return module.thrift_metadata.gen_metadata_service_Raiser()
    
    class Async(_fbthrift_python_AsyncClient):
        @staticmethod
        def __get_thrift_name__() -> str:
            return "module.Raiser"
    
        @staticmethod
        def __get_thrift_uri__() -> _typing.Optional[str]:
            return None
    
        @staticmethod
        def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
            return module.thrift_metadata.gen_metadata_service_Raiser()
    
        async def doBland(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> None:
            _fbthrift_resp = await self._send_request(
                "Raiser",
                "doBland",
                module.thrift_mutable_types._fbthrift_Raiser_doBland_args(),
                module.thrift_mutable_types._fbthrift_Raiser_doBland_result,
                qualifier = _fbthrift_FunctionQualifier.Unspecified,
                uri_or_name="Raiser",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
    
        async def doRaise(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> None:
            _fbthrift_resp = await self._send_request(
                "Raiser",
                "doRaise",
                module.thrift_mutable_types._fbthrift_Raiser_doRaise_args(),
                module.thrift_mutable_types._fbthrift_Raiser_doRaise_result,
                qualifier = _fbthrift_FunctionQualifier.Unspecified,
                uri_or_name="Raiser",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            if _fbthrift_resp.b is not None:
                raise _fbthrift_resp.b
            if _fbthrift_resp.f is not None:
                raise _fbthrift_resp.f
            if _fbthrift_resp.s is not None:
                raise _fbthrift_resp.s
    
        async def get200(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> str:
            _fbthrift_resp = await self._send_request(
                "Raiser",
                "get200",
                module.thrift_mutable_types._fbthrift_Raiser_get200_args(),
                module.thrift_mutable_types._fbthrift_Raiser_get200_result,
                qualifier = _fbthrift_FunctionQualifier.Unspecified,
                uri_or_name="Raiser",
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
    
        async def get500(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> str:
            _fbthrift_resp = await self._send_request(
                "Raiser",
                "get500",
                module.thrift_mutable_types._fbthrift_Raiser_get500_args(),
                module.thrift_mutable_types._fbthrift_Raiser_get500_result,
                qualifier = _fbthrift_FunctionQualifier.Unspecified,
                uri_or_name="Raiser",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            # shortcut to success path for non-void returns
            if _fbthrift_resp.success is not None:
                return _fbthrift_resp.success
            if _fbthrift_resp.f is not None:
                raise _fbthrift_resp.f
            if _fbthrift_resp.b is not None:
                raise _fbthrift_resp.b
            if _fbthrift_resp.s is not None:
                raise _fbthrift_resp.s
            raise _fbthrift_python_exceptions.ApplicationError(
                _fbthrift_python_exceptions.ApplicationErrorType.MISSING_RESULT,
                "Empty Response",
            )
    
    class Sync(_fbthrift_python_SyncClient):
        @staticmethod
        def __get_thrift_name__() -> str:
            return "module.Raiser"
    
        @staticmethod
        def __get_thrift_uri__() -> _typing.Optional[str]:
            return None
    
        @staticmethod
        def __get_metadata__() -> _fbthrift_metadata.ThriftMetadata:
            return module.thrift_metadata.gen_metadata_service_Raiser()
    
        def doBland(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> None:
            _fbthrift_resp = self._send_request(
                "Raiser",
                "doBland",
                module.thrift_mutable_types._fbthrift_Raiser_doBland_args(),
                module.thrift_mutable_types._fbthrift_Raiser_doBland_result,
                uri_or_name="Raiser",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
    
        def doRaise(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> None:
            _fbthrift_resp = self._send_request(
                "Raiser",
                "doRaise",
                module.thrift_mutable_types._fbthrift_Raiser_doRaise_args(),
                module.thrift_mutable_types._fbthrift_Raiser_doRaise_result,
                uri_or_name="Raiser",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            if _fbthrift_resp.b is not None:
                raise _fbthrift_resp.b
            if _fbthrift_resp.f is not None:
                raise _fbthrift_resp.f
            if _fbthrift_resp.s is not None:
                raise _fbthrift_resp.s
    
        def get200(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> str:
            _fbthrift_resp = self._send_request(
                "Raiser",
                "get200",
                module.thrift_mutable_types._fbthrift_Raiser_get200_args(),
                module.thrift_mutable_types._fbthrift_Raiser_get200_result,
                uri_or_name="Raiser",
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
    
        def get500(
            self,
            *,
            rpc_options: _typing.Optional[RpcOptions] = None,
        ) -> str:
            _fbthrift_resp = self._send_request(
                "Raiser",
                "get500",
                module.thrift_mutable_types._fbthrift_Raiser_get500_args(),
                module.thrift_mutable_types._fbthrift_Raiser_get500_result,
                uri_or_name="Raiser",
                rpc_options=rpc_options,
                is_mutable_types=True,
            )
            # shortcut to success path for non-void returns
            if _fbthrift_resp.success is not None:
                return _fbthrift_resp.success
            if _fbthrift_resp.f is not None:
                raise _fbthrift_resp.f
            if _fbthrift_resp.b is not None:
                raise _fbthrift_resp.b
            if _fbthrift_resp.s is not None:
                raise _fbthrift_resp.s
            raise _fbthrift_python_exceptions.ApplicationError(
                _fbthrift_python_exceptions.ApplicationErrorType.MISSING_RESULT,
                "Empty Response",
            )

