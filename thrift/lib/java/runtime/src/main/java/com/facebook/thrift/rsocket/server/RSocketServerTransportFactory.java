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

package com.facebook.thrift.rsocket.server;

import com.facebook.swift.service.ThriftServerConfig;
import com.facebook.thrift.server.RpcServerHandler;
import com.facebook.thrift.server.ServerTransportFactory;
import com.facebook.thrift.util.RpcServerUtils;
import java.net.InetSocketAddress;
import java.net.SocketAddress;
import reactor.core.publisher.Mono;

public class RSocketServerTransportFactory
    implements ServerTransportFactory<RSocketServerTransport> {
  private final ThriftServerConfig config;

  public RSocketServerTransportFactory(ThriftServerConfig config) {
    this.config = config;
  }

  @Override
  public Mono<? extends RSocketServerTransport> createServerTransport(
      SocketAddress bindAddress, RpcServerHandler rpcServerHandler) {
    return RSocketServerTransport.createInstance(parsePort(bindAddress), rpcServerHandler, config);
  }

  public Mono<? extends RSocketServerTransport> createServerTransport(
      RpcServerHandler rpcServerHandler) {
    return createServerTransport(
        new InetSocketAddress("localhost", parsePort(config.getPort())), rpcServerHandler);
  }

  private int parsePort(int port) {
    return port == 0 ? RpcServerUtils.findFreePort() : port;
  }

  private SocketAddress parsePort(SocketAddress socketAddress) {
    if (socketAddress instanceof InetSocketAddress
        && ((InetSocketAddress) socketAddress).getPort() == 0) {
      InetSocketAddress inetSocketAddress = (InetSocketAddress) socketAddress;
      return new InetSocketAddress(
          inetSocketAddress.getHostString(), parsePort(inetSocketAddress.getPort()));
    }
    return socketAddress;
  }
}
