/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.interactions;

import java.util.*;

public class FactoriesReactiveBlockingWrapper 
  implements Factories {
  private final Factories.Reactive _delegate;

  public FactoriesReactiveBlockingWrapper(Factories.Reactive _delegate) {
    
    this._delegate = _delegate;
  }

  public FactoriesReactiveBlockingWrapper(org.apache.thrift.ProtocolId _protocolId, reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient, Map<String, String> _headers, Map<String, String> _persistentHeaders) {
    this(new FactoriesReactiveClient(_protocolId, _rpcClient, _headers, _persistentHeaders));
  }

  @java.lang.Override
  public void close() {
    _delegate.dispose();
  }

  @java.lang.Override
  public void foo() throws org.apache.thrift.TException {
      fooWrapper(com.facebook.thrift.client.RpcOptions.EMPTY);
  }

  @java.lang.Override
  public void foo(
        com.facebook.thrift.client.RpcOptions rpcOptions) throws org.apache.thrift.TException {
      fooWrapper(rpcOptions);
  }

  @java.lang.Override
  public com.facebook.thrift.client.ResponseWrapper<Void> fooWrapper(
    com.facebook.thrift.client.RpcOptions rpcOptions) throws org.apache.thrift.TException {
      try {
        return _delegate.fooWrapper(rpcOptions).block();
      } catch (Throwable t) {
        throw com.facebook.thrift.util.ExceptionUtil.wrap(t);
      }
  }
  @java.lang.Override
  public void interact( final int arg) throws org.apache.thrift.TException {
      interactWrapper(arg, com.facebook.thrift.client.RpcOptions.EMPTY);
  }

  @java.lang.Override
  public void interact(
        final int arg,
        com.facebook.thrift.client.RpcOptions rpcOptions) throws org.apache.thrift.TException {
      interactWrapper(arg,rpcOptions);
  }

  @java.lang.Override
  public com.facebook.thrift.client.ResponseWrapper<Void> interactWrapper(
    final int arg,
    com.facebook.thrift.client.RpcOptions rpcOptions) throws org.apache.thrift.TException {
      try {
        return _delegate.interactWrapper(arg, rpcOptions).block();
      } catch (Throwable t) {
        throw com.facebook.thrift.util.ExceptionUtil.wrap(t);
      }
  }
  @java.lang.Override
  public int interactFast() throws org.apache.thrift.TException {
      return interactFastWrapper(com.facebook.thrift.client.RpcOptions.EMPTY).getData();
  }

  @java.lang.Override
  public int interactFast(
        com.facebook.thrift.client.RpcOptions rpcOptions) throws org.apache.thrift.TException {
      return interactFastWrapper(rpcOptions).getData();
  }

  @java.lang.Override
  public com.facebook.thrift.client.ResponseWrapper<Integer> interactFastWrapper(
    com.facebook.thrift.client.RpcOptions rpcOptions) throws org.apache.thrift.TException {
      try {
        return _delegate.interactFastWrapper(rpcOptions).block();
      } catch (Throwable t) {
        throw com.facebook.thrift.util.ExceptionUtil.wrap(t);
      }
  }

  public MyInteraction createMyInteraction() {
      throw new UnsupportedOperationException("Interactions are not yet supported on ReactiveBlockingWrapper Interfaces!");
  }

  public MyInteractionFast createMyInteractionFast() {
      throw new UnsupportedOperationException("Interactions are not yet supported on ReactiveBlockingWrapper Interfaces!");
  }

  public SerialInteraction createSerialInteraction() {
      throw new UnsupportedOperationException("Interactions are not yet supported on ReactiveBlockingWrapper Interfaces!");
  }
}
