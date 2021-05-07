/**
 * Autogenerated by Thrift for src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/basic/gen-cpp2/DbMixedStackArgumentsAsyncClient.h"

#include <thrift/lib/cpp2/gen/client_cpp.h>

namespace cpp2 {
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::string, ::std::string*>> DbMixedStackArguments_getDataByKey0_pargs;
typedef apache::thrift::ThriftPresult<true, apache::thrift::FieldData<0, ::apache::thrift::type_class::binary, ::std::string*>> DbMixedStackArguments_getDataByKey0_presult;
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::string, ::std::string*>> DbMixedStackArguments_getDataByKey1_pargs;
typedef apache::thrift::ThriftPresult<true, apache::thrift::FieldData<0, ::apache::thrift::type_class::binary, ::std::string*>> DbMixedStackArguments_getDataByKey1_presult;

template <typename Protocol_>
void DbMixedStackArgumentsAsyncClient::getDataByKey0T(Protocol_* prot, apache::thrift::RpcOptions rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::std::string& p_key) {

  std::shared_ptr<apache::thrift::transport::THeader> header(ctx, &ctx->header);
  DbMixedStackArguments_getDataByKey0_pargs args;
  args.get<0>().value = const_cast<::std::string*>(&p_key);
  auto sizer = [&](Protocol_* p) { return args.serializedSizeZC(p); };
  auto writer = [&](Protocol_* p) { args.write(p); };
  static constexpr std::string_view methodName = "getDataByKey0";
  apache::thrift::clientSendT<apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, Protocol_>(prot, std::move(rpcOptions), std::move(callback), contextStack, std::move(header), channel_.get(), apache::thrift::ManagedStringView::from_static(methodName), writer, sizer);
  ctx->reqContext.setRequestHeader(nullptr);
}

template <typename Protocol_>
void DbMixedStackArgumentsAsyncClient::getDataByKey1T(Protocol_* prot, apache::thrift::RpcOptions rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::std::string& p_key) {

  std::shared_ptr<apache::thrift::transport::THeader> header(ctx, &ctx->header);
  DbMixedStackArguments_getDataByKey1_pargs args;
  args.get<0>().value = const_cast<::std::string*>(&p_key);
  auto sizer = [&](Protocol_* p) { return args.serializedSizeZC(p); };
  auto writer = [&](Protocol_* p) { args.write(p); };
  static constexpr std::string_view methodName = "getDataByKey1";
  apache::thrift::clientSendT<apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, Protocol_>(prot, std::move(rpcOptions), std::move(callback), contextStack, std::move(header), channel_.get(), apache::thrift::ManagedStringView::from_static(methodName), writer, sizer);
  ctx->reqContext.setRequestHeader(nullptr);
}



void DbMixedStackArgumentsAsyncClient::getDataByKey0(std::unique_ptr<apache::thrift::RequestCallback> callback, const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  getDataByKey0(rpcOptions, std::move(callback), p_key);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey0(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, const ::std::string& p_key) {
  auto ctx = getDataByKey0Ctx(&rpcOptions);
  apache::thrift::RequestCallback::Context callbackContext;
  callbackContext.protocolId =
      apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* contextStack = ctx->ctx.get();
  if (callback) {
    callbackContext.ctx = std::move(ctx->ctx);
  }
  auto wrappedCallback = apache::thrift::toRequestClientCallbackPtr(std::move(callback), std::move(callbackContext));
  getDataByKey0Impl(rpcOptions, std::move(ctx), contextStack, std::move(wrappedCallback), p_key);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey0Impl(const apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::std::string& p_key) {
  switch (apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolWriter writer;
      getDataByKey0T(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_key);
      break;
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolWriter writer;
      getDataByKey0T(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_key);
      break;
    }
    default:
    {
      apache::thrift::detail::ac::throw_app_exn("Could not find Protocol");
    }
  }
}

std::shared_ptr<::apache::thrift::detail::ac::ClientRequestContext> DbMixedStackArgumentsAsyncClient::getDataByKey0Ctx(apache::thrift::RpcOptions* rpcOptions) {
  return std::make_shared<apache::thrift::detail::ac::ClientRequestContext>(
      channel_->getProtocolId(),
      rpcOptions ? rpcOptions->releaseWriteHeaders() : std::map<std::string, std::string>{},
      handlers_,
      getServiceName(),
      "DbMixedStackArguments.getDataByKey0");
}

void DbMixedStackArgumentsAsyncClient::sync_getDataByKey0(::std::string& _return, const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  sync_getDataByKey0(rpcOptions, _return, p_key);
}

void DbMixedStackArgumentsAsyncClient::sync_getDataByKey0(apache::thrift::RpcOptions& rpcOptions, ::std::string& _return, const ::std::string& p_key) {
  apache::thrift::ClientReceiveState returnState;
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto ctx = getDataByKey0Ctx(&rpcOptions);
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  getDataByKey0Impl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_key);
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
      recv_getDataByKey0(_return, returnState);
  });
}


folly::Future<::std::string> DbMixedStackArgumentsAsyncClient::future_getDataByKey0(const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  return future_getDataByKey0(rpcOptions, p_key);
}

folly::SemiFuture<::std::string> DbMixedStackArgumentsAsyncClient::semifuture_getDataByKey0(const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  return semifuture_getDataByKey0(rpcOptions, p_key);
}

folly::Future<::std::string> DbMixedStackArgumentsAsyncClient::future_getDataByKey0(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  folly::Promise<::std::string> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::FutureCallback<::std::string>>(std::move(promise), recv_wrapped_getDataByKey0, channel_);
  getDataByKey0(rpcOptions, std::move(callback), p_key);
  return future;
}

folly::SemiFuture<::std::string> DbMixedStackArgumentsAsyncClient::semifuture_getDataByKey0(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  auto callbackAndFuture = makeSemiFutureCallback(recv_wrapped_getDataByKey0, channel_);
  auto callback = std::move(callbackAndFuture.first);
  getDataByKey0(rpcOptions, std::move(callback), p_key);
  return std::move(callbackAndFuture.second);
}

folly::Future<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> DbMixedStackArgumentsAsyncClient::header_future_getDataByKey0(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  folly::Promise<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::HeaderFutureCallback<::std::string>>(std::move(promise), recv_wrapped_getDataByKey0, channel_);
  getDataByKey0(rpcOptions, std::move(callback), p_key);
  return future;
}

folly::SemiFuture<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> DbMixedStackArgumentsAsyncClient::header_semifuture_getDataByKey0(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  auto callbackAndFuture = makeHeaderSemiFutureCallback(recv_wrapped_getDataByKey0, channel_);
  auto callback = std::move(callbackAndFuture.first);
  getDataByKey0(rpcOptions, std::move(callback), p_key);
  return std::move(callbackAndFuture.second);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey0(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, const ::std::string& p_key) {
  getDataByKey0(std::make_unique<apache::thrift::FunctionReplyCallback>(std::move(callback)), p_key);
}

#if FOLLY_HAS_COROUTINES
#endif // FOLLY_HAS_COROUTINES
folly::exception_wrapper DbMixedStackArgumentsAsyncClient::recv_wrapped_getDataByKey0(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  if (state.isException()) {
    return std::move(state.exception());
  }
  if (!state.buf()) {
    return folly::make_exception_wrapper<apache::thrift::TApplicationException>("recv_ called without result");
  }

  using result = DbMixedStackArguments_getDataByKey0_presult;
  constexpr auto const fname = "getDataByKey0";
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

void DbMixedStackArgumentsAsyncClient::recv_getDataByKey0(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  auto ew = recv_wrapped_getDataByKey0(_return, state);
  if (ew) {
    ew.throw_exception();
  }
}

void DbMixedStackArgumentsAsyncClient::recv_instance_getDataByKey0(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_getDataByKey0(_return, state);
}

folly::exception_wrapper DbMixedStackArgumentsAsyncClient::recv_instance_wrapped_getDataByKey0(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_wrapped_getDataByKey0(_return, state);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey1(std::unique_ptr<apache::thrift::RequestCallback> callback, const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  getDataByKey1(rpcOptions, std::move(callback), p_key);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey1(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, const ::std::string& p_key) {
  auto ctx = getDataByKey1Ctx(&rpcOptions);
  apache::thrift::RequestCallback::Context callbackContext;
  callbackContext.protocolId =
      apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto* contextStack = ctx->ctx.get();
  if (callback) {
    callbackContext.ctx = std::move(ctx->ctx);
  }
  auto wrappedCallback = apache::thrift::toRequestClientCallbackPtr(std::move(callback), std::move(callbackContext));
  getDataByKey1Impl(rpcOptions, std::move(ctx), contextStack, std::move(wrappedCallback), p_key);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey1Impl(const apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::detail::ac::ClientRequestContext> ctx, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::std::string& p_key) {
  switch (apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId()) {
    case apache::thrift::protocol::T_BINARY_PROTOCOL:
    {
      apache::thrift::BinaryProtocolWriter writer;
      getDataByKey1T(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_key);
      break;
    }
    case apache::thrift::protocol::T_COMPACT_PROTOCOL:
    {
      apache::thrift::CompactProtocolWriter writer;
      getDataByKey1T(&writer, rpcOptions, std::move(ctx), contextStack, std::move(callback), p_key);
      break;
    }
    default:
    {
      apache::thrift::detail::ac::throw_app_exn("Could not find Protocol");
    }
  }
}

std::shared_ptr<::apache::thrift::detail::ac::ClientRequestContext> DbMixedStackArgumentsAsyncClient::getDataByKey1Ctx(apache::thrift::RpcOptions* rpcOptions) {
  return std::make_shared<apache::thrift::detail::ac::ClientRequestContext>(
      channel_->getProtocolId(),
      rpcOptions ? rpcOptions->releaseWriteHeaders() : std::map<std::string, std::string>{},
      handlers_,
      getServiceName(),
      "DbMixedStackArguments.getDataByKey1");
}

void DbMixedStackArgumentsAsyncClient::sync_getDataByKey1(::std::string& _return, const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  sync_getDataByKey1(rpcOptions, _return, p_key);
}

void DbMixedStackArgumentsAsyncClient::sync_getDataByKey1(apache::thrift::RpcOptions& rpcOptions, ::std::string& _return, const ::std::string& p_key) {
  apache::thrift::ClientReceiveState returnState;
  apache::thrift::ClientSyncCallback<false> callback(&returnState);
  auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
  auto evb = apache::thrift::GeneratedAsyncClient::getChannel()->getEventBase();
  auto ctx = getDataByKey1Ctx(&rpcOptions);
  auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(&callback);
  getDataByKey1Impl(rpcOptions, ctx, ctx->ctx.get(), std::move(wrappedCallback), p_key);
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
      recv_getDataByKey1(_return, returnState);
  });
}


folly::Future<::std::string> DbMixedStackArgumentsAsyncClient::future_getDataByKey1(const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  return future_getDataByKey1(rpcOptions, p_key);
}

folly::SemiFuture<::std::string> DbMixedStackArgumentsAsyncClient::semifuture_getDataByKey1(const ::std::string& p_key) {
  ::apache::thrift::RpcOptions rpcOptions;
  return semifuture_getDataByKey1(rpcOptions, p_key);
}

folly::Future<::std::string> DbMixedStackArgumentsAsyncClient::future_getDataByKey1(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  folly::Promise<::std::string> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::FutureCallback<::std::string>>(std::move(promise), recv_wrapped_getDataByKey1, channel_);
  getDataByKey1(rpcOptions, std::move(callback), p_key);
  return future;
}

folly::SemiFuture<::std::string> DbMixedStackArgumentsAsyncClient::semifuture_getDataByKey1(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  auto callbackAndFuture = makeSemiFutureCallback(recv_wrapped_getDataByKey1, channel_);
  auto callback = std::move(callbackAndFuture.first);
  getDataByKey1(rpcOptions, std::move(callback), p_key);
  return std::move(callbackAndFuture.second);
}

folly::Future<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> DbMixedStackArgumentsAsyncClient::header_future_getDataByKey1(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  folly::Promise<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> promise;
  auto future = promise.getFuture();
  auto callback = std::make_unique<apache::thrift::HeaderFutureCallback<::std::string>>(std::move(promise), recv_wrapped_getDataByKey1, channel_);
  getDataByKey1(rpcOptions, std::move(callback), p_key);
  return future;
}

folly::SemiFuture<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> DbMixedStackArgumentsAsyncClient::header_semifuture_getDataByKey1(apache::thrift::RpcOptions& rpcOptions, const ::std::string& p_key) {
  auto callbackAndFuture = makeHeaderSemiFutureCallback(recv_wrapped_getDataByKey1, channel_);
  auto callback = std::move(callbackAndFuture.first);
  getDataByKey1(rpcOptions, std::move(callback), p_key);
  return std::move(callbackAndFuture.second);
}

void DbMixedStackArgumentsAsyncClient::getDataByKey1(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, const ::std::string& p_key) {
  getDataByKey1(std::make_unique<apache::thrift::FunctionReplyCallback>(std::move(callback)), p_key);
}

#if FOLLY_HAS_COROUTINES
#endif // FOLLY_HAS_COROUTINES
folly::exception_wrapper DbMixedStackArgumentsAsyncClient::recv_wrapped_getDataByKey1(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  if (state.isException()) {
    return std::move(state.exception());
  }
  if (!state.buf()) {
    return folly::make_exception_wrapper<apache::thrift::TApplicationException>("recv_ called without result");
  }

  using result = DbMixedStackArguments_getDataByKey1_presult;
  constexpr auto const fname = "getDataByKey1";
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

void DbMixedStackArgumentsAsyncClient::recv_getDataByKey1(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  auto ew = recv_wrapped_getDataByKey1(_return, state);
  if (ew) {
    ew.throw_exception();
  }
}

void DbMixedStackArgumentsAsyncClient::recv_instance_getDataByKey1(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_getDataByKey1(_return, state);
}

folly::exception_wrapper DbMixedStackArgumentsAsyncClient::recv_instance_wrapped_getDataByKey1(::std::string& _return, ::apache::thrift::ClientReceiveState& state) {
  return recv_wrapped_getDataByKey1(_return, state);
}


} // cpp2
