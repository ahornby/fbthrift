/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.interactions;

import java.util.*;
import org.apache.thrift.protocol.*;
import com.facebook.thrift.util.Readers;

public class FactoriesRpcServerHandler 
  implements com.facebook.thrift.server.RpcServerHandler {

  private final java.util.Map<String, com.facebook.thrift.server.RpcServerHandler> _methodMap;

  private final Factories.Reactive _delegate;

  private final java.util.List<com.facebook.thrift.payload.Reader> _fooReaders;
  private final java.util.List<com.facebook.thrift.payload.Reader> _interactReaders;
  private final java.util.List<com.facebook.thrift.payload.Reader> _interactFastReaders;
  private final java.util.List<com.facebook.thrift.payload.Reader> _serializeReaders;
  private final com.facebook.thrift.server.generated.SingleRequestStreamResponseDelegate<com.facebook.thrift.model.StreamResponse<Integer,Integer>> _delegate_serialize;
  private final com.facebook.thrift.server.generated.StreamWithFirstResponseHandler<Integer,Integer> _handler_serialize;

  private final java.util.List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers;

  public FactoriesRpcServerHandler(Factories _delegate,
                                    java.util.List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
    this(new FactoriesBlockingReactiveWrapper(_delegate), _eventHandlers);
  }

  public FactoriesRpcServerHandler(Factories.Async _delegate,
                                    java.util.List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
    this(new FactoriesAsyncReactiveWrapper(_delegate), _eventHandlers);
  }

  public FactoriesRpcServerHandler(Factories.Reactive _delegate,
                                    java.util.List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
    
    this._methodMap = new java.util.HashMap<>();
    this._delegate = _delegate;
    this._eventHandlers = _eventHandlers;

    _methodMap.put("foo", this);
    _fooReaders = _create_foo_request_readers();
    _methodMap.put("interact", this);
    _interactReaders = _create_interact_request_readers();
    _methodMap.put("interactFast", this);
    _interactFastReaders = _create_interactFast_request_readers();

    _methodMap.put("serialize", this);
    _serializeReaders = _create_serialize_request_readers();
    _delegate_serialize = (java.util.List<Object> _list) -> {
        java.util.Iterator _iterator = _list.iterator();

        return _delegate.serialize();
    };

    _handler_serialize =
    new com.facebook.thrift.server.generated.StreamWithFirstResponseHandler<>(
      _delegate_serialize,
      _first_ResponseWriterFactory_serialize,
      _ResponseWriterFactory_serialize,
      _serializeReaders,
      "serialize",
      new Class[]{  },
      new Integer[]{  },
      new Class[]{  },
      new Integer[]{  }
    );


  }

  private final static java.util.List<com.facebook.thrift.payload.Reader> _create_serialize_request_readers() {
    java.util.List<com.facebook.thrift.payload.Reader> _readerList = new java.util.ArrayList<>();


    return _readerList;
  }

  private final static com.facebook.thrift.server.generated.ResponseWriterFactory _ResponseWriterFactory_serialize = (java.lang.Object _o, com.facebook.swift.service.ContextChain _chain, com.facebook.thrift.payload.ServerRequestPayload _requestPayload) -> {
    return oprot -> {
      try {
        oprot.writeStructBegin(com.facebook.thrift.util.RpcPayloadUtil.TSTRUCT);

        
        int _iter0 = (int)_o;
        oprot.writeFieldBegin(com.facebook.thrift.util.RpcPayloadUtil.I32_FIELD);
oprot.writeI32(_iter0);
        oprot.writeFieldEnd();

        oprot.writeFieldStop();
        oprot.writeStructEnd();

        _chain.postWrite(_o);
      } catch (Throwable _e) {
        com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
        throw reactor.core.Exceptions.propagate(_e);
      }
    };
  };

  private final static com.facebook.thrift.server.generated.ResponseWriterFactory _first_ResponseWriterFactory_serialize = (java.lang.Object _o, com.facebook.swift.service.ContextChain _chain, com.facebook.thrift.payload.ServerRequestPayload _requestPayload) -> {
    return oprot -> {
      try {
        oprot.writeStructBegin(com.facebook.thrift.util.RpcPayloadUtil.TSTRUCT);

        
        int _iter0 = (int)_o;
        oprot.writeFieldBegin(com.facebook.thrift.util.RpcPayloadUtil.I32_FIELD);
oprot.writeI32(_iter0);
        oprot.writeFieldEnd();

        oprot.writeFieldStop();
        oprot.writeStructEnd();

        _chain.postWrite(_o);
      } catch (Throwable _e) {
        com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
        throw reactor.core.Exceptions.propagate(_e);
      }
    };
  };



  private static java.util.List<com.facebook.thrift.payload.Reader> _create_foo_request_readers() {
    java.util.List<com.facebook.thrift.payload.Reader> _readerList = new java.util.ArrayList<>();


    return _readerList;
  }

  private static com.facebook.thrift.payload.Writer _create_foo_response_writer(
      final java.lang.Object _r,
      final com.facebook.swift.service.ContextChain _chain,
      final int _seqId) {
      return oprot -> {
      try {
        oprot.writeStructBegin(com.facebook.thrift.util.RpcPayloadUtil.TSTRUCT);

        

        oprot.writeFieldStop();
        oprot.writeStructEnd();

        _chain.postWrite(_r);
      } catch (Throwable _e) {
        com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
        throw reactor.core.Exceptions.propagate(_e);
      }
    };
  }


  private static reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload>
    _dofoo(
    Factories.Reactive _delegate,
    com.facebook.thrift.payload.ServerRequestPayload _payload,
    java.util.List<com.facebook.thrift.payload.Reader> _readers,
    com.facebook.swift.service.ContextChain _chain) {
          _chain.preRead();
          java.util.List<java.lang.Object>_data = _payload.getData(_readers);
          java.util.Iterator<java.lang.Object> _iterator = _data.iterator();


          _chain.postRead(_data);

          reactor.core.publisher.Mono<Void> _delegateResponse =
            _delegate
              .foo()
              .doFirst(() -> com.facebook.nifty.core.RequestContexts.setCurrentContext(_payload.getRequestContext()));

          reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload> _internalResponse =
            _delegateResponse.map(_response -> {
              _chain.preWrite(_response);
              com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload =
                com.facebook.thrift.util.RpcPayloadUtil.createServerResponsePayload(
                  _payload,
                  _create_foo_response_writer(_response, _chain, _payload.getMessageSeqId()));

                return _serverResponsePayload;
            })
            .switchIfEmpty(
              reactor.core.publisher.Mono.fromSupplier(
                () -> {
                  _chain.preWrite(null);
                  return com.facebook.thrift.util.RpcPayloadUtil.createServerResponsePayload(
                    _payload,
                    _create_foo_response_writer(null, _chain, _payload.getMessageSeqId()));
                }
              )
            )
            .<com.facebook.thrift.payload.ServerResponsePayload>onErrorResume(_t -> {
                _chain.preWriteException(_t);
                // exception is not of user declared type
                String _errorMessage = String.format("Internal error processing foo: %s", _t.getMessage() == null ? "<null>" : _t.getMessage());
                org.apache.thrift.TApplicationException _tApplicationException =
                    new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.INTERNAL_ERROR, _errorMessage);
                _tApplicationException.initCause(_t);
                com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload =
                    com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(),  _chain);

                return reactor.core.publisher.Mono.just(_serverResponsePayload);
            })
            .doFinally(__ -> {
              _chain.done();
            });

          return _internalResponse;
  }
  private static java.util.List<com.facebook.thrift.payload.Reader> _create_interact_request_readers() {
    java.util.List<com.facebook.thrift.payload.Reader> _readerList = new java.util.ArrayList<>();

    
    _readerList.add(Readers.i32Reader());

    return _readerList;
  }

  private static com.facebook.thrift.payload.Writer _create_interact_response_writer(
      final java.lang.Object _r,
      final com.facebook.swift.service.ContextChain _chain,
      final int _seqId) {
      return oprot -> {
      try {
        oprot.writeStructBegin(com.facebook.thrift.util.RpcPayloadUtil.TSTRUCT);

        

        oprot.writeFieldStop();
        oprot.writeStructEnd();

        _chain.postWrite(_r);
      } catch (Throwable _e) {
        com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
        throw reactor.core.Exceptions.propagate(_e);
      }
    };
  }


  private static reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload>
    _dointeract(
    Factories.Reactive _delegate,
    com.facebook.thrift.payload.ServerRequestPayload _payload,
    java.util.List<com.facebook.thrift.payload.Reader> _readers,
    com.facebook.swift.service.ContextChain _chain) {
          _chain.preRead();
          java.util.List<java.lang.Object>_data = _payload.getData(_readers);
          java.util.Iterator<java.lang.Object> _iterator = _data.iterator();

          int arg = (int) _iterator.next();

          _chain.postRead(_data);

          reactor.core.publisher.Mono<Void> _delegateResponse =
            _delegate
              .interact(arg)
              .doFirst(() -> com.facebook.nifty.core.RequestContexts.setCurrentContext(_payload.getRequestContext()));

          reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload> _internalResponse =
            _delegateResponse.map(_response -> {
              _chain.preWrite(_response);
              com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload =
                com.facebook.thrift.util.RpcPayloadUtil.createServerResponsePayload(
                  _payload,
                  _create_interact_response_writer(_response, _chain, _payload.getMessageSeqId()));

                return _serverResponsePayload;
            })
            .switchIfEmpty(
              reactor.core.publisher.Mono.fromSupplier(
                () -> {
                  _chain.preWrite(null);
                  return com.facebook.thrift.util.RpcPayloadUtil.createServerResponsePayload(
                    _payload,
                    _create_interact_response_writer(null, _chain, _payload.getMessageSeqId()));
                }
              )
            )
            .<com.facebook.thrift.payload.ServerResponsePayload>onErrorResume(_t -> {
                _chain.preWriteException(_t);
                // exception is not of user declared type
                String _errorMessage = String.format("Internal error processing interact: %s", _t.getMessage() == null ? "<null>" : _t.getMessage());
                org.apache.thrift.TApplicationException _tApplicationException =
                    new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.INTERNAL_ERROR, _errorMessage);
                _tApplicationException.initCause(_t);
                com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload =
                    com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(),  _chain);

                return reactor.core.publisher.Mono.just(_serverResponsePayload);
            })
            .doFinally(__ -> {
              _chain.done();
            });

          return _internalResponse;
  }
  private static java.util.List<com.facebook.thrift.payload.Reader> _create_interactFast_request_readers() {
    java.util.List<com.facebook.thrift.payload.Reader> _readerList = new java.util.ArrayList<>();


    return _readerList;
  }

  private static com.facebook.thrift.payload.Writer _create_interactFast_response_writer(
      final java.lang.Object _r,
      final com.facebook.swift.service.ContextChain _chain,
      final int _seqId) {
      return oprot -> {
      try {
        oprot.writeStructBegin(com.facebook.thrift.util.RpcPayloadUtil.TSTRUCT);

        
        int _iter0 = (int)_r;
        oprot.writeFieldBegin(com.facebook.thrift.util.RpcPayloadUtil.I32_FIELD);
oprot.writeI32(_iter0);
        oprot.writeFieldEnd();

        oprot.writeFieldStop();
        oprot.writeStructEnd();

        _chain.postWrite(_r);
      } catch (Throwable _e) {
        com.facebook.thrift.util.NettyUtil.releaseIfByteBufTProtocol(oprot);
        throw reactor.core.Exceptions.propagate(_e);
      }
    };
  }


  private static reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload>
    _dointeractFast(
    Factories.Reactive _delegate,
    com.facebook.thrift.payload.ServerRequestPayload _payload,
    java.util.List<com.facebook.thrift.payload.Reader> _readers,
    com.facebook.swift.service.ContextChain _chain) {
          _chain.preRead();
          java.util.List<java.lang.Object>_data = _payload.getData(_readers);
          java.util.Iterator<java.lang.Object> _iterator = _data.iterator();


          _chain.postRead(_data);

          reactor.core.publisher.Mono<Integer> _delegateResponse =
            _delegate
              .interactFast()
              .doFirst(() -> com.facebook.nifty.core.RequestContexts.setCurrentContext(_payload.getRequestContext()));

          reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload> _internalResponse =
            _delegateResponse.map(_response -> {
              _chain.preWrite(_response);
              com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload =
                com.facebook.thrift.util.RpcPayloadUtil.createServerResponsePayload(
                  _payload,
                  _create_interactFast_response_writer(_response, _chain, _payload.getMessageSeqId()));

                return _serverResponsePayload;
            })
            .switchIfEmpty(
              reactor.core.publisher.Mono.fromSupplier(
                () -> {
                  org.apache.thrift.TApplicationException _tApplicationException =
                    new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.MISSING_RESULT, "method interactFast returned null");
                  return com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), _chain);
                }
              )
            )
            .<com.facebook.thrift.payload.ServerResponsePayload>onErrorResume(_t -> {
                _chain.preWriteException(_t);
                // exception is not of user declared type
                String _errorMessage = String.format("Internal error processing interactFast: %s", _t.getMessage() == null ? "<null>" : _t.getMessage());
                org.apache.thrift.TApplicationException _tApplicationException =
                    new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.INTERNAL_ERROR, _errorMessage);
                _tApplicationException.initCause(_t);
                com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload =
                    com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(),  _chain);

                return reactor.core.publisher.Mono.just(_serverResponsePayload);
            })
            .doFinally(__ -> {
              _chain.done();
            });

          return _internalResponse;
  }


  @java.lang.Override
  public reactor.core.publisher.Flux<com.facebook.thrift.payload.ServerResponsePayload> singleRequestStreamingResponse(com.facebook.thrift.payload.ServerRequestPayload _payload) {
    final String _name = _payload.getRequestRpcMetadata().getName();
    reactor.core.publisher.Flux<com.facebook.thrift.payload.ServerResponsePayload> _retVal = reactor.core.publisher.Flux.defer(() -> {
    com.facebook.swift.service.ContextChain _chain;
    try {
      String qualifiedMethodName = "Factories." + _name;
      _chain = new com.facebook.swift.service.ContextChain(_eventHandlers, qualifiedMethodName, _payload.getRequestContext());
    } catch (Throwable _t) {
      org.apache.thrift.TApplicationException _tApplicationException = new org.apache.thrift.TApplicationException(_t.getMessage());
      com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload = com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), null);
      return reactor.core.publisher.Flux.just(_serverResponsePayload);
    }

    reactor.core.publisher.Flux<com.facebook.thrift.payload.ServerResponsePayload> _result;
    try {
      switch(_name) {
        case "serialize":
          _result = _handler_serialize.handleStream(_payload, _chain);
          break;
        default: {
            _chain.preRead();
            org.apache.thrift.TApplicationException _tApplicationException = new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.UNKNOWN_METHOD, "no method found with name " + _name);
            com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload = com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), _chain);
            return reactor.core.publisher.Flux.just(_serverResponsePayload);
        }
      }
    } catch (org.apache.thrift.TApplicationException _tApplicationException) {
      com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload = com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), _chain);
      return reactor.core.publisher.Flux.just(_serverResponsePayload);
    } catch (Throwable _t) {
      _result = reactor.core.publisher.Flux.error(_t);
    }
    return _result;
    });
    if (com.facebook.thrift.util.resources.RpcResources.isForceExecutionOffEventLoop()) {
      _retVal = _retVal.subscribeOn(com.facebook.thrift.util.resources.RpcResources.getOffLoopScheduler());
    }
    return _retVal;
  }

  @java.lang.Override
  public reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload> singleRequestSingleResponse(com.facebook.thrift.payload.ServerRequestPayload _payload) {
    final String _name = _payload.getRequestRpcMetadata().getName();
    reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload> _retVal = reactor.core.publisher.Mono.defer(() -> {
    com.facebook.swift.service.ContextChain _chain;
    try {
      String qualifiedMethodName = "Factories." + _name;
      _chain = new com.facebook.swift.service.ContextChain(_eventHandlers, qualifiedMethodName, _payload.getRequestContext());
    } catch (Throwable _t) {
      org.apache.thrift.TApplicationException _tApplicationException = new org.apache.thrift.TApplicationException(_t.getMessage());
      com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload = com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), null);
      return reactor.core.publisher.Mono.just(_serverResponsePayload);
    }

    reactor.core.publisher.Mono<com.facebook.thrift.payload.ServerResponsePayload> _result;
    try {
      switch (_name) {
        case "foo":
          _result = _dofoo(_delegate, _payload, _fooReaders, _chain);
        break;
        case "interact":
          _result = _dointeract(_delegate, _payload, _interactReaders, _chain);
        break;
        case "interactFast":
          _result = _dointeractFast(_delegate, _payload, _interactFastReaders, _chain);
        break;
        default: {
            _chain.preRead();
            org.apache.thrift.TApplicationException _tApplicationException = new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.UNKNOWN_METHOD, "no method found with name " + _name);
            com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload = com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), _chain);
            return reactor.core.publisher.Mono.just(_serverResponsePayload);
        }
      }
    } catch (org.apache.thrift.TApplicationException _tApplicationException) {
      com.facebook.thrift.payload.ServerResponsePayload _serverResponsePayload = com.facebook.thrift.util.RpcPayloadUtil.fromTApplicationException(_tApplicationException, _payload.getRequestRpcMetadata(), _chain);
      return reactor.core.publisher.Mono.just(_serverResponsePayload);
    } catch (Throwable _t) {
      _result = reactor.core.publisher.Mono.error(_t);
    }

    return _result;
    });

    if (com.facebook.thrift.util.resources.RpcResources.isForceExecutionOffEventLoop()) {
      _retVal = _retVal.subscribeOn(com.facebook.thrift.util.resources.RpcResources.getOffLoopScheduler());
    }
    return _retVal;
  }

  @java.lang.Override
  public reactor.core.publisher.Mono<Void> singleRequestNoResponse(com.facebook.thrift.payload.ServerRequestPayload _payload) {
    final String _name = _payload.getRequestRpcMetadata().getName();

    reactor.core.publisher.Mono<Void> _retVal = reactor.core.publisher.Mono.defer(() -> {
    com.facebook.swift.service.ContextChain _chain;
    try {
      String qualifiedMethodName = "Factories." + _name;
      _chain = new com.facebook.swift.service.ContextChain(_eventHandlers, qualifiedMethodName, _payload.getRequestContext());
    } catch (Throwable _t) {
      return reactor.core.publisher.Mono.error(_t);
    }

    reactor.core.publisher.Mono<Void> _result;
    try {
      switch (_name) {
        default: {
          _result = reactor.core.publisher.Mono.error(new org.apache.thrift.TApplicationException(org.apache.thrift.TApplicationException.UNKNOWN_METHOD, "no method found with name " + _name));
        }
      }
    } catch (Throwable _t) {
      _result = reactor.core.publisher.Mono.error(_t);
    }

    return _result;
    });

    if (com.facebook.thrift.util.resources.RpcResources.isForceExecutionOffEventLoop()) {
      _retVal = _retVal.subscribeOn(com.facebook.thrift.util.resources.RpcResources.getOffLoopScheduler());
    }
    return _retVal;
  }

  public java.util.Map<String, com.facebook.thrift.server.RpcServerHandler> getMethodMap() {
      return _methodMap;
  }

}
