/**
 * Autogenerated by Thrift for src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/basic-stack-arguments/gen-cpp2/MyServiceFastAsyncClient.h"

#include <thrift/lib/cpp2/gen/client_cpp.h>

namespace cpp2 {
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::integral, ::std::int64_t*>> MyServiceFast_hasDataById_pargs;
typedef apache::thrift::ThriftPresult<true, apache::thrift::FieldData<0, ::apache::thrift::type_class::integral, bool*>> MyServiceFast_hasDataById_presult;
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::integral, ::std::int64_t*>> MyServiceFast_getDataById_pargs;
typedef apache::thrift::ThriftPresult<true, apache::thrift::FieldData<0, ::apache::thrift::type_class::string, ::std::string*>> MyServiceFast_getDataById_presult;
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::integral, ::std::int64_t*>, apache::thrift::FieldData<2, ::apache::thrift::type_class::string, ::std::string*>> MyServiceFast_putDataById_pargs;
typedef apache::thrift::ThriftPresult<true> MyServiceFast_putDataById_presult;
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::integral, ::std::int64_t*>, apache::thrift::FieldData<2, ::apache::thrift::type_class::string, ::std::string*>> MyServiceFast_lobDataById_pargs;

template <typename Protocol_>
void MyServiceFastAsyncClient::hasDataByIdT(Protocol_* prot, apache::thrift::RpcOptions rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id) {

  std::shared_ptr<apache::thrift::transport::THeader> header(ctx, &ctx->header);
  MyServiceFast_hasDataById_pargs args;
  args.get<0>().value = &p_id;
  auto sizer = [&](Protocol_* p) { return args.serializedSizeZC(p); };
  auto writer = [&](Protocol_* p) { args.write(p); };
  static constexpr std::string_view methodName = "hasDataById";
  apache::thrift::clientSendT<apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, Protocol_>(prot, std::move(rpcOptions), std::move(callback), contextStack, std::move(header), channel_.get(), apache::thrift::ManagedStringView::from_static(methodName), writer, sizer);
  ctx->reqContext.setRequestHeader(nullptr);
}

template <typename Protocol_>
void MyServiceFastAsyncClient::getDataByIdT(Protocol_* prot, apache::thrift::RpcOptions rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id) {

  std::shared_ptr<apache::thrift::transport::THeader> header(ctx, &ctx->header);
  MyServiceFast_getDataById_pargs args;
  args.get<0>().value = &p_id;
  auto sizer = [&](Protocol_* p) { return args.serializedSizeZC(p); };
  auto writer = [&](Protocol_* p) { args.write(p); };
  static constexpr std::string_view methodName = "getDataById";
  apache::thrift::clientSendT<apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, Protocol_>(prot, std::move(rpcOptions), std::move(callback), contextStack, std::move(header), channel_.get(), apache::thrift::ManagedStringView::from_static(methodName), writer, sizer);
  ctx->reqContext.setRequestHeader(nullptr);
}

template <typename Protocol_>
void MyServiceFastAsyncClient::putDataByIdT(Protocol_* prot, apache::thrift::RpcOptions rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, const ::std::string& p_data) {

  std::shared_ptr<apache::thrift::transport::THeader> header(ctx, &ctx->header);
  MyServiceFast_putDataById_pargs args;
  args.get<0>().value = &p_id;
  args.get<1>().value = const_cast<::std::string*>(&p_data);
  auto sizer = [&](Protocol_* p) { return args.serializedSizeZC(p); };
  auto writer = [&](Protocol_* p) { args.write(p); };
  static constexpr std::string_view methodName = "putDataById";
  apache::thrift::clientSendT<apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, Protocol_>(prot, std::move(rpcOptions), std::move(callback), contextStack, std::move(header), channel_.get(), apache::thrift::ManagedStringView::from_static(methodName), writer, sizer);
  ctx->reqContext.setRequestHeader(nullptr);
}

template <typename Protocol_>
void MyServiceFastAsyncClient::lobDataByIdT(Protocol_* prot, apache::thrift::RpcOptions rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, const ::std::string& p_data) {

  std::shared_ptr<apache::thrift::transport::THeader> header(ctx, &ctx->header);
  MyServiceFast_lobDataById_pargs args;
  args.get<0>().value = &p_id;
  args.get<1>().value = const_cast<::std::string*>(&p_data);
  auto sizer = [&](Protocol_* p) { return args.serializedSizeZC(p); };
  auto writer = [&](Protocol_* p) { args.write(p); };
  static constexpr std::string_view methodName = "lobDataById";
  apache::thrift::clientSendT<apache::thrift::RpcKind::SINGLE_REQUEST_NO_RESPONSE, Protocol_>(prot, std::move(rpcOptions), std::move(callback), contextStack, std::move(header), channel_.get(), apache::thrift::ManagedStringView::from_static(methodName), writer, sizer);
  ctx->reqContext.setRequestHeader(nullptr);
}



void MyServiceFastAsyncClient::hasDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  hasDataById(rpcOptions, std::move(callback), p_id);
}

void MyServiceFastAsyncClient::hasDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id) {
  auto ctx = hasDataByIdCtx(&rpcOptions);
  apache::thrift::RequestCallback::Context callbackContext;
  callbackContext.protocolId =
      apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* contextStack = ctx->ctx.get();
  if (callback) {
    callbackContext.ctx = std::move(ctx->ctx);
  }
  auto wrappedCallback = apache::thrift::toRequestClientCallbackPtr(std::move(callback), std::move(callbackContext));
  hasDataByIdImpl(rpcOptions, std::move(ctx), contextStack, std::move(wrappedCallback), p_id);
}

void MyServiceFastAsyncClient::hasDataByIdImpl(const apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id) {
  switch (apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolWriter writer;
      hasDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id);
      break;
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolWriter writer;
      hasDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id);
      break;
    }
    default:
    {
      apache::thrift::detail::ac::throw_app_exn("Could not find Protocol");
    }
  }
}

std::shared_ptr<::apache::thrift::detail::ac::ClientRequestContext> MyServiceFastAsyncClient::hasDataByIdCtx(apache::thrift::RpcOptions* rpcOptions) {
  return std::make_shared<apache::thrift::detail::ac::ClientRequestContext>(
      channel_->getProtocolId(),
      rpcOptions ? rpcOptions->releaseWriteHeaders() : std::map<std::string, std::string>{},
      handlers_,
      getServiceName(),
      "MyServiceFast.hasDataById");
}

bool MyServiceFastAsyncClient::sync_hasDataById(::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  return sync_hasDataById(rpcOptions, p_id);
}

bool MyServiceFastAsyncClient::sync_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  apache::thrift::ClientReceiveState returnState;
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto ctx = hasDataByIdCtx(&rpcOptions);
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  hasDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id);
  callback.waitUntilDone(evb);

  if (returnState.isException()) {
    returnState.exception().throw_exception();
  }
  returnState.resetProtocolId(protocolId);
  returnState.resetCtx(std::move(ctx->ctx));
  SCOPE_EXIT {
    if (returnState.header() && !returnState.header()->getHeaders().empty()) {
      rpcOptions.setReadHeaders(returnState.header()->releaseHeaders());
    }
  };
  return folly::fibers::runInMainContext([&] {
      return recv_hasDataById(returnState);
  });
}

folly::Try<apache::thrift::RpcResponseComplete<bool>>
MyServiceFastAsyncClient::sync_complete_hasDataById(
    apache::thrift::RpcOptions& rpcOptions,  ::std::int64_t p_id) {
  apache::thrift::ClientReceiveState returnState;
  auto ctx = hasDataByIdCtx(&rpcOptions);
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  const auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* const evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  hasDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id);

  callback.waitUntilDone(evb);
  returnState.resetProtocolId(protocolId);
  returnState.resetCtx(std::move(ctx->ctx));

  folly::Try<apache::thrift::RpcResponseComplete<bool>> tryResponse;
  if (!returnState.buf()) {
    assert(returnState.isException());
  	tryResponse.emplaceException(std::move(returnState.exception()));
  } else {
    tryResponse.emplace();
    tryResponse->responseContext.rpcSizeStats = returnState.getRpcSizeStats();
    if (auto* header = returnState.header()) {
      if (!header->getHeaders().empty()) {
  	    tryResponse->responseContext.headers = header->releaseHeaders();
      }
      if (auto load = header->getServerLoad()) {
        tryResponse->responseContext.serverLoad = *load;
      }
    }
    tryResponse->response = folly::makeTryWith([&] {
      return folly::fibers::runInMainContext([&] {
        return recv_hasDataById(returnState);
      });
    });
  }
  return tryResponse;
}

folly::Future<bool> MyServiceFastAsyncClient::future_hasDataById(::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  return future_hasDataById(rpcOptions, p_id);
}

folly::SemiFuture<bool> MyServiceFastAsyncClient::semifuture_hasDataById(::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  return semifuture_hasDataById(rpcOptions, p_id);
}

folly::Future<bool> MyServiceFastAsyncClient::future_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  folly::Promise<bool> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::FutureCallback<bool>>(std::move(promise), recv_wrapped_hasDataById, channel_);
  hasDataById(rpcOptions, std::move(callback), p_id);
  return future;
}

folly::SemiFuture<bool> MyServiceFastAsyncClient::semifuture_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  auto callbackAndFuture = makeSemiFutureCallback(recv_wrapped_hasDataById, channel_);
  auto callback = std::move(callbackAndFuture.first);
  hasDataById(rpcOptions, std::move(callback), p_id);
  return std::move(callbackAndFuture.second);
}

folly::Future<std::pair<bool, std::unique_ptr<apache::thrift::transport::THeader>>> MyServiceFastAsyncClient::header_future_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  folly::Promise<std::pair<bool, std::unique_ptr<apache::thrift::transport::THeader>>> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::HeaderFutureCallback<bool>>(std::move(promise), recv_wrapped_hasDataById, channel_);
  hasDataById(rpcOptions, std::move(callback), p_id);
  return future;
}

folly::SemiFuture<std::pair<bool, std::unique_ptr<apache::thrift::transport::THeader>>> MyServiceFastAsyncClient::header_semifuture_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  auto callbackAndFuture = makeHeaderSemiFutureCallback(recv_wrapped_hasDataById, channel_);
  auto callback = std::move(callbackAndFuture.first);
  hasDataById(rpcOptions, std::move(callback), p_id);
  return std::move(callbackAndFuture.second);
}

void MyServiceFastAsyncClient::hasDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id) {
  hasDataById(std::make_unique<apache::thrift::FunctionReplyCallback>(std::move(callback)), p_id);
}

#if FOLLY_HAS_COROUTINES
#endif // FOLLY_HAS_COROUTINES
folly::exception_wrapper MyServiceFastAsyncClient::recv_wrapped_hasDataById(bool& _return, ::apache::thrift::ClientReceiveState& state) {
  if (state.isException()) {
    return std::move(state.exception());
  }
  if (!state.buf()) {
    return folly::make_exception_wrapper<apache::thrift::TApplicationException>("recv_ called without result");
  }

  using result = MyServiceFast_hasDataById_presult;
  constexpr auto const fname = "hasDataById";
  switch (state.protocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolReader reader;
      return apache::thrift::detail::ac::recv_wrapped<result>(
          fname, &reader, state, _return);
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolReader reader;
      return apache::thrift::detail::ac::recv_wrapped<result>(
          fname, &reader, state, _return);
    }
    default:
    {
    }
  }
  return folly::make_exception_wrapper<apache::thrift::TApplicationException>("Could not find Protocol");
}

bool MyServiceFastAsyncClient::recv_hasDataById(::apache::thrift::ClientReceiveState& state) {
  bool _return;
  auto ew = recv_wrapped_hasDataById(_return, state);
  if (ew) {
    ew.throw_exception();
  }
  return _return;
}

bool MyServiceFastAsyncClient::recv_instance_hasDataById(::apache::thrift::ClientReceiveState& state) {
  return recv_hasDataById(state);
}

folly::exception_wrapper MyServiceFastAsyncClient::recv_instance_wrapped_hasDataById(bool& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_wrapped_hasDataById(_return, state);
}

void MyServiceFastAsyncClient::getDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  getDataById(rpcOptions, std::move(callback), p_id);
}

void MyServiceFastAsyncClient::getDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id) {
  auto ctx = getDataByIdCtx(&rpcOptions);
  apache::thrift::RequestCallback::Context callbackContext;
  callbackContext.protocolId =
      apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* contextStack = ctx->ctx.get();
  if (callback) {
    callbackContext.ctx = std::move(ctx->ctx);
  }
  auto wrappedCallback = apache::thrift::toRequestClientCallbackPtr(std::move(callback), std::move(callbackContext));
  getDataByIdImpl(rpcOptions, std::move(ctx), contextStack, std::move(wrappedCallback), p_id);
}

void MyServiceFastAsyncClient::getDataByIdImpl(const apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id) {
  switch (apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolWriter writer;
      getDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id);
      break;
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolWriter writer;
      getDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id);
      break;
    }
    default:
    {
      apache::thrift::detail::ac::throw_app_exn("Could not find Protocol");
    }
  }
}

std::shared_ptr<::apache::thrift::detail::ac::ClientRequestContext> MyServiceFastAsyncClient::getDataByIdCtx(apache::thrift::RpcOptions* rpcOptions) {
  return std::make_shared<apache::thrift::detail::ac::ClientRequestContext>(
      channel_->getProtocolId(),
      rpcOptions ? rpcOptions->releaseWriteHeaders() : std::map<std::string, std::string>{},
      handlers_,
      getServiceName(),
      "MyServiceFast.getDataById");
}

void MyServiceFastAsyncClient::sync_getDataById(::std::string& _return, ::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  sync_getDataById(rpcOptions, _return, p_id);
}

void MyServiceFastAsyncClient::sync_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::string& _return, ::std::int64_t p_id) {
  apache::thrift::ClientReceiveState returnState;
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto ctx = getDataByIdCtx(&rpcOptions);
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  getDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id);
  callback.waitUntilDone(evb);

  if (returnState.isException()) {
    returnState.exception().throw_exception();
  }
  returnState.resetProtocolId(protocolId);
  returnState.resetCtx(std::move(ctx->ctx));
  SCOPE_EXIT {
    if (returnState.header() && !returnState.header()->getHeaders().empty()) {
      rpcOptions.setReadHeaders(returnState.header()->releaseHeaders());
    }
  };
  return folly::fibers::runInMainContext([&] {
      recv_getDataById(_return, returnState);
  });
}

folly::Try<apache::thrift::RpcResponseComplete<::std::string>>
MyServiceFastAsyncClient::sync_complete_getDataById(
    apache::thrift::RpcOptions& rpcOptions,  ::std::int64_t p_id) {
  apache::thrift::ClientReceiveState returnState;
  auto ctx = getDataByIdCtx(&rpcOptions);
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  const auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* const evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  getDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id);

  callback.waitUntilDone(evb);
  returnState.resetProtocolId(protocolId);
  returnState.resetCtx(std::move(ctx->ctx));

  folly::Try<apache::thrift::RpcResponseComplete<::std::string>> tryResponse;
  if (!returnState.buf()) {
    assert(returnState.isException());
  	tryResponse.emplaceException(std::move(returnState.exception()));
  } else {
    tryResponse.emplace();
    tryResponse->responseContext.rpcSizeStats = returnState.getRpcSizeStats();
    if (auto* header = returnState.header()) {
      if (!header->getHeaders().empty()) {
  	    tryResponse->responseContext.headers = header->releaseHeaders();
      }
      if (auto load = header->getServerLoad()) {
        tryResponse->responseContext.serverLoad = *load;
      }
    }
    tryResponse->response = folly::makeTryWith([&] {
      return folly::fibers::runInMainContext([&] {
        ::std::string rv;
        recv_getDataById(rv, returnState);
        return rv;
      });
    });
  }
  return tryResponse;
}

folly::Future<::std::string> MyServiceFastAsyncClient::future_getDataById(::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  return future_getDataById(rpcOptions, p_id);
}

folly::SemiFuture<::std::string> MyServiceFastAsyncClient::semifuture_getDataById(::std::int64_t p_id) {
  ::apache::thrift::RpcOptions rpcOptions;
  return semifuture_getDataById(rpcOptions, p_id);
}

folly::Future<::std::string> MyServiceFastAsyncClient::future_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  folly::Promise<::std::string> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::FutureCallback<::std::string>>(std::move(promise), recv_wrapped_getDataById, channel_);
  getDataById(rpcOptions, std::move(callback), p_id);
  return future;
}

folly::SemiFuture<::std::string> MyServiceFastAsyncClient::semifuture_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  auto callbackAndFuture = makeSemiFutureCallback(recv_wrapped_getDataById, channel_);
  auto callback = std::move(callbackAndFuture.first);
  getDataById(rpcOptions, std::move(callback), p_id);
  return std::move(callbackAndFuture.second);
}

folly::Future<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> MyServiceFastAsyncClient::header_future_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  folly::Promise<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::HeaderFutureCallback<::std::string>>(std::move(promise), recv_wrapped_getDataById, channel_);
  getDataById(rpcOptions, std::move(callback), p_id);
  return future;
}

folly::SemiFuture<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> MyServiceFastAsyncClient::header_semifuture_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
  auto callbackAndFuture = makeHeaderSemiFutureCallback(recv_wrapped_getDataById, channel_);
  auto callback = std::move(callbackAndFuture.first);
  getDataById(rpcOptions, std::move(callback), p_id);
  return std::move(callbackAndFuture.second);
}

void MyServiceFastAsyncClient::getDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id) {
  getDataById(std::make_unique<apache::thrift::FunctionReplyCallback>(std::move(callback)), p_id);
}

#if FOLLY_HAS_COROUTINES
#endif // FOLLY_HAS_COROUTINES
folly::exception_wrapper MyServiceFastAsyncClient::recv_wrapped_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  if (state.isException()) {
    return std::move(state.exception());
  }
  if (!state.buf()) {
    return folly::make_exception_wrapper<apache::thrift::TApplicationException>("recv_ called without result");
  }

  using result = MyServiceFast_getDataById_presult;
  constexpr auto const fname = "getDataById";
  switch (state.protocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolReader reader;
      return apache::thrift::detail::ac::recv_wrapped<result>(
          fname, &reader, state, _return);
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolReader reader;
      return apache::thrift::detail::ac::recv_wrapped<result>(
          fname, &reader, state, _return);
    }
    default:
    {
    }
  }
  return folly::make_exception_wrapper<apache::thrift::TApplicationException>("Could not find Protocol");
}

void MyServiceFastAsyncClient::recv_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  auto ew = recv_wrapped_getDataById(_return, state);
  if (ew) {
    ew.throw_exception();
  }
}

void MyServiceFastAsyncClient::recv_instance_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_getDataById(_return, state);
}

folly::exception_wrapper MyServiceFastAsyncClient::recv_instance_wrapped_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_wrapped_getDataById(_return, state);
}

void MyServiceFastAsyncClient::putDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  putDataById(rpcOptions, std::move(callback), p_id, p_data);
}

void MyServiceFastAsyncClient::putDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data) {
  auto ctx = putDataByIdCtx(&rpcOptions);
  apache::thrift::RequestCallback::Context callbackContext;
  callbackContext.protocolId =
      apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* contextStack = ctx->ctx.get();
  if (callback) {
    callbackContext.ctx = std::move(ctx->ctx);
  }
  auto wrappedCallback = apache::thrift::toRequestClientCallbackPtr(std::move(callback), std::move(callbackContext));
  putDataByIdImpl(rpcOptions, std::move(ctx), contextStack, std::move(wrappedCallback), p_id, p_data);
}

void MyServiceFastAsyncClient::putDataByIdImpl(const apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, const ::std::string& p_data) {
  switch (apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolWriter writer;
      putDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id, p_data);
      break;
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolWriter writer;
      putDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id, p_data);
      break;
    }
    default:
    {
      apache::thrift::detail::ac::throw_app_exn("Could not find Protocol");
    }
  }
}

std::shared_ptr<::apache::thrift::detail::ac::ClientRequestContext> MyServiceFastAsyncClient::putDataByIdCtx(apache::thrift::RpcOptions* rpcOptions) {
  return std::make_shared<apache::thrift::detail::ac::ClientRequestContext>(
      channel_->getProtocolId(),
      rpcOptions ? rpcOptions->releaseWriteHeaders() : std::map<std::string, std::string>{},
      handlers_,
      getServiceName(),
      "MyServiceFast.putDataById");
}

void MyServiceFastAsyncClient::sync_putDataById(::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  sync_putDataById(rpcOptions, p_id, p_data);
}

void MyServiceFastAsyncClient::sync_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  apache::thrift::ClientReceiveState returnState;
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto ctx = putDataByIdCtx(&rpcOptions);
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  putDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id, p_data);
  callback.waitUntilDone(evb);

  if (returnState.isException()) {
    returnState.exception().throw_exception();
  }
  returnState.resetProtocolId(protocolId);
  returnState.resetCtx(std::move(ctx->ctx));
  SCOPE_EXIT {
    if (returnState.header() && !returnState.header()->getHeaders().empty()) {
      rpcOptions.setReadHeaders(returnState.header()->releaseHeaders());
    }
  };
  return folly::fibers::runInMainContext([&] {
      recv_putDataById(returnState);
  });
}

folly::Try<apache::thrift::RpcResponseComplete<void>>
MyServiceFastAsyncClient::sync_complete_putDataById(
    apache::thrift::RpcOptions& rpcOptions,  ::std::int64_t p_id, const ::std::string& p_data) {
  apache::thrift::ClientReceiveState returnState;
  auto ctx = putDataByIdCtx(&rpcOptions);
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  const auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* const evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  putDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id, p_data);

  callback.waitUntilDone(evb);
  returnState.resetProtocolId(protocolId);
  returnState.resetCtx(std::move(ctx->ctx));

  folly::Try<apache::thrift::RpcResponseComplete<void>> tryResponse;
  if (!returnState.buf()) {
    assert(returnState.isException());
  	tryResponse.emplaceException(std::move(returnState.exception()));
  } else {
    tryResponse.emplace();
    tryResponse->responseContext.rpcSizeStats = returnState.getRpcSizeStats();
    if (auto* header = returnState.header()) {
      if (!header->getHeaders().empty()) {
  	    tryResponse->responseContext.headers = header->releaseHeaders();
      }
      if (auto load = header->getServerLoad()) {
        tryResponse->responseContext.serverLoad = *load;
      }
    }
    tryResponse->response = folly::makeTryWith([&] {
      return folly::fibers::runInMainContext([&] {
        return recv_putDataById(returnState);
      });
    });
  }
  return tryResponse;
}

folly::Future<folly::Unit> MyServiceFastAsyncClient::future_putDataById(::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  return future_putDataById(rpcOptions, p_id, p_data);
}

folly::SemiFuture<folly::Unit> MyServiceFastAsyncClient::semifuture_putDataById(::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  return semifuture_putDataById(rpcOptions, p_id, p_data);
}

folly::Future<folly::Unit> MyServiceFastAsyncClient::future_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  folly::Promise<folly::Unit> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::FutureCallback<folly::Unit>>(std::move(promise), recv_wrapped_putDataById, channel_);
  putDataById(rpcOptions, std::move(callback), p_id, p_data);
  return future;
}

folly::SemiFuture<folly::Unit> MyServiceFastAsyncClient::semifuture_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  auto callbackAndFuture = makeSemiFutureCallback(recv_wrapped_putDataById, channel_);
  auto callback = std::move(callbackAndFuture.first);
  putDataById(rpcOptions, std::move(callback), p_id, p_data);
  return std::move(callbackAndFuture.second);
}

folly::Future<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> MyServiceFastAsyncClient::header_future_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  folly::Promise<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::HeaderFutureCallback<folly::Unit>>(std::move(promise), recv_wrapped_putDataById, channel_);
  putDataById(rpcOptions, std::move(callback), p_id, p_data);
  return future;
}

folly::SemiFuture<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> MyServiceFastAsyncClient::header_semifuture_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  auto callbackAndFuture = makeHeaderSemiFutureCallback(recv_wrapped_putDataById, channel_);
  auto callback = std::move(callbackAndFuture.first);
  putDataById(rpcOptions, std::move(callback), p_id, p_data);
  return std::move(callbackAndFuture.second);
}

void MyServiceFastAsyncClient::putDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id, const ::std::string& p_data) {
  putDataById(std::make_unique<apache::thrift::FunctionReplyCallback>(std::move(callback)), p_id, p_data);
}

#if FOLLY_HAS_COROUTINES
#endif // FOLLY_HAS_COROUTINES
folly::exception_wrapper MyServiceFastAsyncClient::recv_wrapped_putDataById(::apache::thrift::ClientReceiveState& state) {
  if (state.isException()) {
    return std::move(state.exception());
  }
  if (!state.buf()) {
    return folly::make_exception_wrapper<apache::thrift::TApplicationException>("recv_ called without result");
  }

  using result = MyServiceFast_putDataById_presult;
  constexpr auto const fname = "putDataById";
  switch (state.protocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolReader reader;
      return apache::thrift::detail::ac::recv_wrapped<result>(
          fname, &reader, state);
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolReader reader;
      return apache::thrift::detail::ac::recv_wrapped<result>(
          fname, &reader, state);
    }
    default:
    {
    }
  }
  return folly::make_exception_wrapper<apache::thrift::TApplicationException>("Could not find Protocol");
}

void MyServiceFastAsyncClient::recv_putDataById(::apache::thrift::ClientReceiveState& state) {
  auto ew = recv_wrapped_putDataById(state);
  if (ew) {
    ew.throw_exception();
  }
}

void MyServiceFastAsyncClient::recv_instance_putDataById(::apache::thrift::ClientReceiveState& state) {
  recv_putDataById(state);
}

folly::exception_wrapper MyServiceFastAsyncClient::recv_instance_wrapped_putDataById(::apache::thrift::ClientReceiveState& state) {
  return recv_wrapped_putDataById(state);
}

void MyServiceFastAsyncClient::lobDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  lobDataById(rpcOptions, std::move(callback), p_id, p_data);
}

void MyServiceFastAsyncClient::lobDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data) {
  auto ctx = lobDataByIdCtx(&rpcOptions);
  apache::thrift::RequestCallback::Context callbackContext;
  callbackContext.oneWay = true;
  callbackContext.protocolId =
      apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* contextStack = ctx->ctx.get();
  if (callback) {
    callbackContext.ctx = std::move(ctx->ctx);
  }
  auto wrappedCallback = apache::thrift::toRequestClientCallbackPtr(std::move(callback), std::move(callbackContext));
  lobDataByIdImpl(rpcOptions, std::move(ctx), contextStack, std::move(wrappedCallback), p_id, p_data);
}

void MyServiceFastAsyncClient::lobDataByIdImpl(const apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, const ::std::string& p_data) {
  switch (apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolWriter writer;
      lobDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id, p_data);
      break;
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolWriter writer;
      lobDataByIdT(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_id, p_data);
      break;
    }
    default:
    {
      apache::thrift::detail::ac::throw_app_exn("Could not find Protocol");
    }
  }
}

std::shared_ptr<::apache::thrift::detail::ac::ClientRequestContext> MyServiceFastAsyncClient::lobDataByIdCtx(apache::thrift::RpcOptions* rpcOptions) {
  return std::make_shared<apache::thrift::detail::ac::ClientRequestContext>(
      channel_->getProtocolId(),
      rpcOptions ? rpcOptions->releaseWriteHeaders() : std::map<std::string, std::string>{},
      handlers_,
      getServiceName(),
      "MyServiceFast.lobDataById");
}

void MyServiceFastAsyncClient::sync_lobDataById(::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  sync_lobDataById(rpcOptions, p_id, p_data);
}

void MyServiceFastAsyncClient::sync_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  apache::thrift::ClientReceiveState returnState;
  apache::thrift::ClientSyncCallback<true> callback(&returnState);
  auto evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto ctx = lobDataByIdCtx(&rpcOptions);
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  lobDataByIdImpl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_id, p_data);
  callback.waitUntilDone(evb);

  if (returnState.isException()) {
    returnState.exception().throw_exception();
  }
}


folly::Future<folly::Unit> MyServiceFastAsyncClient::future_lobDataById(::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  return future_lobDataById(rpcOptions, p_id, p_data);
}

folly::SemiFuture<folly::Unit> MyServiceFastAsyncClient::semifuture_lobDataById(::std::int64_t p_id, const ::std::string& p_data) {
  ::apache::thrift::RpcOptions rpcOptions;
  return semifuture_lobDataById(rpcOptions, p_id, p_data);
}

folly::Future<folly::Unit> MyServiceFastAsyncClient::future_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  folly::Promise<folly::Unit> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::OneWayFutureCallback>(std::move(promise), channel_);
  lobDataById(rpcOptions, std::move(callback), p_id, p_data);
  return future;
}

folly::SemiFuture<folly::Unit> MyServiceFastAsyncClient::semifuture_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
  auto callbackAndFuture = makeOneWaySemiFutureCallback(channel_);
  auto callback = std::move(callbackAndFuture.first);
  lobDataById(rpcOptions, std::move(callback), p_id, p_data);
  return std::move(callbackAndFuture.second);
}


void MyServiceFastAsyncClient::lobDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id, const ::std::string& p_data) {
  lobDataById(std::make_unique<apache::thrift::FunctionReplyCallback>(std::move(callback)), p_id, p_data);
}

#if FOLLY_HAS_COROUTINES
#endif // FOLLY_HAS_COROUTINES

} // cpp2
