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

from cpython.ref cimport PyObject
from libcpp.memory cimport shared_ptr
from folly cimport cFollyTry
from thrift.python.client cimport ssl as thrift_ssl
from thrift.python.client.request_channel cimport (
    cRequestChannel_ptr,
    ClientType as cClientType,
    ChannelFactory,
)
from thrift.python.protocol cimport Protocol as cProtocol


cdef void requestchannel_callback(
    cFollyTry[cRequestChannel_ptr]&& result,
    PyObject* userData,
)

cpdef object get_proxy_factory()

cdef object get_client_with_channel_factory(
    clientKlass,
    shared_ptr[ChannelFactory] channel_factory,
    host = ?,
    port = ?,
    path = ?,
    double timeout = ?,
    cClientType client_type = ?,
    cProtocol protocol = ?,
    thrift_ssl.SSLContext ssl_context = ?,
    double ssl_timeout = ?,
)
