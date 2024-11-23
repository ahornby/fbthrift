#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/interactions/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import folly.iobuf as _fbthrift_iobuf
import thrift.py3.types
import thrift.py3.client
import thrift.py3.common
import typing as _typing
from types import TracebackType

import test.fixtures.interactions.module.types as _test_fixtures_interactions_module_types
import test.fixtures.another_interactions.shared.types as _test_fixtures_another_interactions_shared_types
import test.fixtures.another_interactions.shared.clients as _test_fixtures_another_interactions_shared_clients


_MyServiceT = _typing.TypeVar('_MyServiceT', bound='MyService')


class MyService(thrift.py3.client.Client):

    async def foo(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def interact(
        self,
        arg: int,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def interactFast(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def serialize(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ResponseAndClientBufferedStream__i32_i32: ...

    def createMyInteraction(self) -> MyService_MyInteraction: ...
    def async_createMyInteraction(self) -> MyService_MyInteraction: ...
    def createMyInteractionFast(self) -> MyService_MyInteractionFast: ...
    def async_createMyInteractionFast(self) -> MyService_MyInteractionFast: ...
    def createSerialInteraction(self) -> MyService_SerialInteraction: ...
    def async_createSerialInteraction(self) -> MyService_SerialInteraction: ...

_MyService_MyInteraction = _typing.TypeVar('_MyService_MyInteraction', bound='MyService_MyInteraction')


class MyService_MyInteraction(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def ping(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def truthify(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ClientBufferedStream__bool: ...


_MyService_MyInteractionFast = _typing.TypeVar('_MyService_MyInteractionFast', bound='MyService_MyInteractionFast')


class MyService_MyInteractionFast(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def ping(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def truthify(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ClientBufferedStream__bool: ...


_MyService_SerialInteraction = _typing.TypeVar('_MyService_SerialInteraction', bound='MyService_SerialInteraction')


class MyService_SerialInteraction(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...



_FactoriesT = _typing.TypeVar('_FactoriesT', bound='Factories')


class Factories(thrift.py3.client.Client):

    async def foo(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def interact(
        self,
        arg: int,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def interactFast(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def serialize(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ResponseAndClientBufferedStream__i32_i32: ...



_PerformT = _typing.TypeVar('_PerformT', bound='Perform')


class Perform(thrift.py3.client.Client):

    async def foo(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    def createMyInteraction(self) -> Perform_MyInteraction: ...
    def async_createMyInteraction(self) -> Perform_MyInteraction: ...
    def createMyInteractionFast(self) -> Perform_MyInteractionFast: ...
    def async_createMyInteractionFast(self) -> Perform_MyInteractionFast: ...
    def createSerialInteraction(self) -> Perform_SerialInteraction: ...
    def async_createSerialInteraction(self) -> Perform_SerialInteraction: ...

_Perform_MyInteraction = _typing.TypeVar('_Perform_MyInteraction', bound='Perform_MyInteraction')


class Perform_MyInteraction(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def ping(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def truthify(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ClientBufferedStream__bool: ...


_Perform_MyInteractionFast = _typing.TypeVar('_Perform_MyInteractionFast', bound='Perform_MyInteractionFast')


class Perform_MyInteractionFast(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def ping(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def truthify(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ClientBufferedStream__bool: ...


_Perform_SerialInteraction = _typing.TypeVar('_Perform_SerialInteraction', bound='Perform_SerialInteraction')


class Perform_SerialInteraction(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...



_InteractWithSharedT = _typing.TypeVar('_InteractWithSharedT', bound='InteractWithShared')


class InteractWithShared(thrift.py3.client.Client):

    async def do_some_similar_things(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_another_interactions_shared_types.DoSomethingResult: ...

    def createSharedInteraction(self) -> InteractWithShared_SharedInteraction: ...
    def async_createSharedInteraction(self) -> InteractWithShared_SharedInteraction: ...
    def createMyInteraction(self) -> InteractWithShared_MyInteraction: ...
    def async_createMyInteraction(self) -> InteractWithShared_MyInteraction: ...

_InteractWithShared_SharedInteraction = _typing.TypeVar('_InteractWithShared_SharedInteraction', bound='InteractWithShared_SharedInteraction')


class InteractWithShared_SharedInteraction(thrift.py3.client.Client):

    async def init(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def do_something(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_another_interactions_shared_types.DoSomethingResult: ...

    async def tear_down(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...


_InteractWithShared_MyInteraction = _typing.TypeVar('_InteractWithShared_MyInteraction', bound='InteractWithShared_MyInteraction')


class InteractWithShared_MyInteraction(thrift.py3.client.Client):

    async def frobnicate(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> int: ...

    async def ping(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> None: ...

    async def truthify(
        self,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ClientBufferedStream__bool: ...



_BoxServiceT = _typing.TypeVar('_BoxServiceT', bound='BoxService')


class BoxService(thrift.py3.client.Client):

    async def getABoxSession(
        self,
        req: _test_fixtures_interactions_module_types.ShouldBeBoxed,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_interactions_module_types.ShouldBeBoxed: ...

