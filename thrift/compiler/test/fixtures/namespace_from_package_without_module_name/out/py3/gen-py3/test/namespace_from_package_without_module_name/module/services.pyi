#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/namespace_from_package_without_module_name/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import folly.iobuf as _fbthrift_iobuf
import typing as _typing
from thrift.py3.server import RequestContext, ServiceInterface
from abc import abstractmethod, ABCMeta

import test.namespace_from_package_without_module_name.module.types as _test_namespace_from_package_without_module_name_module_types

_TestServiceInterfaceT = _typing.TypeVar('_TestServiceInterfaceT', bound='TestServiceInterface')


class TestServiceInterface(
    ServiceInterface,
    metaclass=ABCMeta,
):


    @abstractmethod
    async def init(
        self,
        int1: int
    ) -> int: ...
    pass


