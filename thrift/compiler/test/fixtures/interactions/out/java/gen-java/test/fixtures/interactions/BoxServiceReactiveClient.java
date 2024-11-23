/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.interactions;

import static com.facebook.swift.service.SwiftConstants.STICKY_HASH_KEY;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicLong;
import org.apache.thrift.protocol.*;
import org.apache.thrift.ClientPushMetadata;
import org.apache.thrift.InteractionCreate;
import org.apache.thrift.InteractionTerminate;
import com.facebook.thrift.client.ResponseWrapper;
import com.facebook.thrift.client.RpcOptions;
import com.facebook.thrift.util.Readers;

public class BoxServiceReactiveClient 
  implements BoxService.Reactive {
  private static final AtomicLong _interactionCounter = new AtomicLong(0);

  protected final org.apache.thrift.ProtocolId _protocolId;
  protected final reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient;
  protected final reactor.core.publisher.Mono<Map<String, String>> _headersMono;
  protected final reactor.core.publisher.Mono<Map<String, String>> _persistentHeadersMono;
  protected final Set<Long> _activeInteractions;

  private static final TField _getABoxSession_REQ_FIELD_DESC = new TField("req", TType.STRUCT, (short)1);
  private static final java.util.Map<Short, com.facebook.thrift.payload.Reader> _getABoxSession_EXCEPTION_READERS = java.util.Collections.emptyMap();
  private static final java.util.Map<Short, com.facebook.thrift.payload.Reader> _getABox_EXCEPTION_READERS_INT = java.util.Collections.emptyMap();

  static {
  }

  public BoxServiceReactiveClient(org.apache.thrift.ProtocolId _protocolId, reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient) {
    
    this._protocolId = _protocolId;
    this._rpcClient = _rpcClient;
    this._headersMono = reactor.core.publisher.Mono.empty();
    this._persistentHeadersMono = reactor.core.publisher.Mono.empty();
    this._activeInteractions = ConcurrentHashMap.newKeySet();
  }

  public BoxServiceReactiveClient(org.apache.thrift.ProtocolId _protocolId, reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient, Map<String, String> _headers, Map<String, String> _persistentHeaders) {
    this(_protocolId, _rpcClient, reactor.core.publisher.Mono.just(_headers != null ? _headers : java.util.Collections.emptyMap()), reactor.core.publisher.Mono.just(_persistentHeaders != null ? _persistentHeaders : java.util.Collections.emptyMap()), new AtomicLong(), ConcurrentHashMap.newKeySet());
  }

  public BoxServiceReactiveClient(org.apache.thrift.ProtocolId _protocolId, reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient, reactor.core.publisher.Mono<Map<String, String>> _headersMono, reactor.core.publisher.Mono<Map<String, String>> _persistentHeadersMono) {
    this(_protocolId, _rpcClient, _headersMono, _persistentHeadersMono, new AtomicLong(), ConcurrentHashMap.newKeySet());
  }

  public BoxServiceReactiveClient(org.apache.thrift.ProtocolId _protocolId, reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient, Map<String, String> _headers, Map<String, String> _persistentHeaders, AtomicLong interactionCounter, Set<Long> activeInteractions) {
    this(_protocolId,_rpcClient, reactor.core.publisher.Mono.just(_headers != null ? _headers : java.util.Collections.emptyMap()), reactor.core.publisher.Mono.just(_persistentHeaders != null ? _persistentHeaders : java.util.Collections.emptyMap()), interactionCounter, activeInteractions);
  }

  public BoxServiceReactiveClient(org.apache.thrift.ProtocolId _protocolId, reactor.core.publisher.Mono<? extends com.facebook.thrift.client.RpcClient> _rpcClient, reactor.core.publisher.Mono<Map<String, String>> _headersMono, reactor.core.publisher.Mono<Map<String, String>> _persistentHeadersMono, AtomicLong interactionCounter, Set<Long> activeInteractions) {
    
    this._protocolId = _protocolId;
    this._rpcClient = _rpcClient;
    this._headersMono = _headersMono;
    this._persistentHeadersMono = _persistentHeadersMono;
    this._activeInteractions = activeInteractions;
  }

  @java.lang.Override
  public void dispose() {}

  private com.facebook.thrift.payload.Writer _creategetABoxSessionWriter(final test.fixtures.interactions.ShouldBeBoxed req) {
    return oprot -> {
      try {
        {
          oprot.writeFieldBegin(_getABoxSession_REQ_FIELD_DESC);

          test.fixtures.interactions.ShouldBeBoxed _iter0 = req;

          _iter0.write0(oprot);
          oprot.writeFieldEnd();
        }


      } catch (Throwable _e) {
        com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
        throw reactor.core.Exceptions.propagate(_e);
      }
    };
  }

  private static final com.facebook.thrift.payload.Reader _getABoxSession_READER = Readers.wrap(test.fixtures.interactions.ShouldBeBoxed.asReader());

  @java.lang.Override
  public reactor.core.publisher.Mono<com.facebook.thrift.client.ResponseWrapper<test.fixtures.interactions.ShouldBeBoxed>> getABoxSessionWrapper(final test.fixtures.interactions.ShouldBeBoxed req,  final com.facebook.thrift.client.RpcOptions rpcOptions) {
    return _rpcClient
      .flatMap(_rpc -> getHeaders(rpcOptions).flatMap(headers -> {
        org.apache.thrift.RequestRpcMetadata _metadata = new org.apache.thrift.RequestRpcMetadata.Builder()
                .setName("getABoxSession")
                .setKind(org.apache.thrift.RpcKind.SINGLE_REQUEST_SINGLE_RESPONSE)
                .setOtherMetadata(headers)
                .setProtocol(_protocolId)
                .build();

            com.facebook.thrift.payload.ClientRequestPayload<test.fixtures.interactions.ShouldBeBoxed> _crp =
                com.facebook.thrift.payload.ClientRequestPayload.create(
                    "BoxService",
                    _creategetABoxSessionWriter(req),
                    _getABoxSession_READER,
                    _getABoxSession_EXCEPTION_READERS,
                    _metadata,
                    java.util.Collections.emptyMap());

            return _rpc
                .singleRequestSingleResponse(_crp, rpcOptions).transform(com.facebook.thrift.util.MonoPublishingTransformer.getInstance()).doOnNext(_p -> {if(_p.getException() != null) throw reactor.core.Exceptions.propagate(_p.getException());});
      }));
  }

  @java.lang.Override
  public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABoxSession(final test.fixtures.interactions.ShouldBeBoxed req,  final com.facebook.thrift.client.RpcOptions rpcOptions) {
    return getABoxSessionWrapper(req,  rpcOptions).map(_p -> _p.getData());
  }

  @java.lang.Override
  public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABoxSession(final test.fixtures.interactions.ShouldBeBoxed req) {
    return getABoxSession(req,  com.facebook.thrift.client.RpcOptions.EMPTY);
  }


  public class BoxedInteractionImpl implements BoxedInteraction {
    private final long interactionId;

    BoxedInteractionImpl(long interactionId) {
      this.interactionId = interactionId;
      com.facebook.thrift.client.ThriftClientStatsHolder.getThriftClientStats().interactionCreated("BoxedInteraction");
    }

    private final java.util.Map<Short, com.facebook.thrift.payload.Reader> _getABox_EXCEPTION_READERS = java.util.Collections.emptyMap();

    private com.facebook.thrift.payload.Writer _creategetABoxWriter() {
      return oprot -> {
        try {

        } catch (Throwable _e) {
          com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
          throw reactor.core.Exceptions.propagate(_e);
        }
      };
    }

    private final com.facebook.thrift.payload.Reader _getABox_READER = Readers.wrap(test.fixtures.interactions.ShouldBeBoxed.asReader());

    public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABox() {
      return getABoxWrapper( com.facebook.thrift.client.RpcOptions.EMPTY).map(_p -> _p.getData());
    }

    @java.lang.Override
    public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABox(RpcOptions rpcOptions)  {
      return getABoxWrapper( rpcOptions).map(_p -> _p.getData());
    }

    @java.lang.Override
    public reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.interactions.ShouldBeBoxed>> getABoxWrapper(RpcOptions rpcOptions)  {
      return _rpcClient
        .contextWrite(ctx -> ctx.put(STICKY_HASH_KEY, interactionId))
        .flatMap(_rpc -> getHeaders(rpcOptions).flatMap(headers -> {
          String interactionName = "BoxedInteraction.getABox";
          org.apache.thrift.RequestRpcMetadata.Builder _metadataBuilder = new org.apache.thrift.RequestRpcMetadata.Builder()
                  .setName(interactionName)
                  .setKind(org.apache.thrift.RpcKind.SINGLE_REQUEST_SINGLE_RESPONSE)
                  .setOtherMetadata(headers)
                  .setProtocol(_protocolId);

          if (_activeInteractions.contains(interactionId)) {
            _metadataBuilder.setInteractionId(interactionId);
          } else {
            _metadataBuilder.setInteractionCreate(
              new InteractionCreate.Builder()
                  .setInteractionId(interactionId)
                  .setInteractionName("BoxedInteraction")
                  .build());
            _metadataBuilder.setInteractionId(0L);
            _activeInteractions.add(interactionId);
          }

          org.apache.thrift.RequestRpcMetadata _metadata = _metadataBuilder.build();

          com.facebook.thrift.payload.ClientRequestPayload<test.fixtures.interactions.ShouldBeBoxed> _crp =
              com.facebook.thrift.payload.ClientRequestPayload.create(
                  "BoxedInteraction",
                  _creategetABoxWriter(),
                  _getABox_READER,
                  _getABox_EXCEPTION_READERS_INT,
                  _metadata,
                  java.util.Collections.emptyMap());

          return _rpc
              .singleRequestSingleResponse(_crp, rpcOptions).doOnNext(_p -> {if(_p.getException() != null) throw reactor.core.Exceptions.propagate(_p.getException());});
        }));
    }

    @java.lang.Override
    public void dispose() {
      com.facebook.thrift.client.ThriftClientStatsHolder.getThriftClientStats().interactionDisposed("BoxedInteraction");
      _activeInteractions.remove(interactionId);
      _rpcClient
        .contextWrite(ctx -> ctx.put(STICKY_HASH_KEY, interactionId))
        .flatMap(_rpc -> {
          InteractionTerminate term = new InteractionTerminate.Builder().setInteractionId(interactionId).build();
          ClientPushMetadata metadata = ClientPushMetadata.fromInteractionTerminate(term);
          return _rpc.metadataPush(metadata, com.facebook.thrift.client.RpcOptions.EMPTY);
        }).subscribe();
    }
  }

  public BoxedInteraction createBoxedInteraction() {
      return new BoxedInteractionImpl(_interactionCounter.incrementAndGet());
  }

  private reactor.core.publisher.Mono<Map<String, String>> getHeaders(com.facebook.thrift.client.RpcOptions rpcOptions) {
      Map<String, String> requestHeaders = new HashMap<>();
      if (rpcOptions.getRequestHeaders() != null && !rpcOptions.getRequestHeaders().isEmpty()) {
          requestHeaders.putAll(rpcOptions.getRequestHeaders());
      }

      return _headersMono.defaultIfEmpty(java.util.Collections.emptyMap()).zipWith(_persistentHeadersMono.defaultIfEmpty(java.util.Collections.emptyMap()), (headers, persistentHeaders) -> {
          Map<String, String> result = new HashMap<>();
          result.putAll(requestHeaders);
          result.putAll(headers);
          result.putAll(persistentHeaders);
          return result;
      });
  }
}
