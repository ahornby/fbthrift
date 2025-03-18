/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include <thrift/lib/cpp2/gen/client_h.h>

#include "thrift/compiler/test/fixtures/basic-stack-arguments/gen-cpp2/module_types.h"

namespace apache { namespace thrift {
  class Cpp2RequestContext;
  namespace detail { namespace ac { struct ClientRequestContext; }}
  namespace transport { class THeader; }
}}

namespace cpp2 {
class MyServiceFast;
} // namespace cpp2
namespace apache::thrift {

template <>
class Client<::cpp2::MyServiceFast> : public apache::thrift::GeneratedAsyncClient {
 public:
  using apache::thrift::GeneratedAsyncClient::GeneratedAsyncClient;

  char const* getServiceName() const noexcept override {
    return "MyServiceFast";
  }


  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual void hasDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual void hasDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id);
 protected:
  void fbthrift_serialize_and_send_hasDataById(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, bool stealRpcOptions = false);
 public:

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual bool sync_hasDataById(::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual bool sync_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual folly::Future<bool> future_hasDataById(::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual folly::SemiFuture<bool> semifuture_hasDataById(::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual folly::Future<bool> future_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual folly::SemiFuture<bool> semifuture_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  FOLLY_NODISCARD [[deprecated("To be replaced by new API soon")]] virtual folly::Try<apache::thrift::RpcResponseComplete<bool>> sync_complete_hasDataById(
      apache::thrift::RpcOptions&& rpcOptions,  ::std::int64_t p_id);

#if FOLLY_HAS_COROUTINES
#if __clang__
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  template <int = 0>
  folly::coro::Task<bool> co_hasDataById(::std::int64_t p_id) {
    return co_hasDataById<false>(nullptr, p_id);
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  template <int = 0>
  folly::coro::Task<bool> co_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
    return co_hasDataById<true>(&rpcOptions, p_id);
  }
#else
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  folly::coro::Task<bool> co_hasDataById(::std::int64_t p_id) {
    co_return co_await folly::coro::detachOnCancel(semifuture_hasDataById(p_id));
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  folly::coro::Task<bool> co_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
    co_return co_await folly::coro::detachOnCancel(semifuture_hasDataById(rpcOptions, p_id));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<bool> co_hasDataById(apache::thrift::RpcOptions* rpcOptions, ::std::int64_t p_id) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = hasDataByIdCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if (ctx != nullptr) {
      auto argsAsRefs = std::tie(p_id);
      ctx->processClientInterceptorsOnRequest(apache::thrift::ClientInterceptorOnRequestArguments(argsAsRefs), header.get(), hasRpcOptions ? *rpcOptions : *defaultRpcOptions).throwUnlessValue();
    }
    if constexpr (hasRpcOptions) {
      fbthrift_serialize_and_send_hasDataById(*rpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id);
    } else {
      fbthrift_serialize_and_send_hasDataById(*defaultRpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (ctx != nullptr) {
      ctx->processClientInterceptorsOnResponse(returnState.header()).throwUnlessValue();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    bool _return;
    if (auto ew = recv_wrapped_hasDataById(_return, returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
    co_return _return;
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual void hasDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id);


  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  static folly::exception_wrapper recv_wrapped_hasDataById(bool& _return, ::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  static bool recv_hasDataById(::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual bool recv_instance_hasDataById(::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "hasDataById"} */
  virtual folly::exception_wrapper recv_instance_wrapped_hasDataById(bool& _return, ::apache::thrift::ClientReceiveState& state);
 private:
  apache::thrift::SerializedRequest fbthrift_serialize_hasDataById(const RpcOptions& rpcOptions, apache::thrift::transport::THeader& header, apache::thrift::ContextStack* contextStack, ::std::int64_t p_id);
  template <typename RpcOptions>
  void fbthrift_send_hasDataById(apache::thrift::SerializedRequest&& request, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::RequestClientCallback::Ptr callback, std::unique_ptr<folly::IOBuf> interceptorFrameworkMetadata);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> hasDataByIdCtx(apache::thrift::RpcOptions* rpcOptions);
  template <typename CallbackType>
  folly::SemiFuture<bool> fbthrift_semifuture_hasDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);
 public:
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual void getDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual void getDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id);
 protected:
  void fbthrift_serialize_and_send_getDataById(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, bool stealRpcOptions = false);
 public:

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual void sync_getDataById(::std::string& _return, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual void sync_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::string& _return, ::std::int64_t p_id);

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual folly::Future<::std::string> future_getDataById(::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual folly::SemiFuture<::std::string> semifuture_getDataById(::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual folly::Future<::std::string> future_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual folly::SemiFuture<::std::string> semifuture_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  FOLLY_NODISCARD [[deprecated("To be replaced by new API soon")]] virtual folly::Try<apache::thrift::RpcResponseComplete<::std::string>> sync_complete_getDataById(
      apache::thrift::RpcOptions&& rpcOptions,  ::std::int64_t p_id);

#if FOLLY_HAS_COROUTINES
#if __clang__
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  template <int = 0>
  folly::coro::Task<::std::string> co_getDataById(::std::int64_t p_id) {
    return co_getDataById<false>(nullptr, p_id);
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  template <int = 0>
  folly::coro::Task<::std::string> co_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
    return co_getDataById<true>(&rpcOptions, p_id);
  }
#else
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  folly::coro::Task<::std::string> co_getDataById(::std::int64_t p_id) {
    co_return co_await folly::coro::detachOnCancel(semifuture_getDataById(p_id));
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  folly::coro::Task<::std::string> co_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id) {
    co_return co_await folly::coro::detachOnCancel(semifuture_getDataById(rpcOptions, p_id));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<::std::string> co_getDataById(apache::thrift::RpcOptions* rpcOptions, ::std::int64_t p_id) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = getDataByIdCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if (ctx != nullptr) {
      auto argsAsRefs = std::tie(p_id);
      ctx->processClientInterceptorsOnRequest(apache::thrift::ClientInterceptorOnRequestArguments(argsAsRefs), header.get(), hasRpcOptions ? *rpcOptions : *defaultRpcOptions).throwUnlessValue();
    }
    if constexpr (hasRpcOptions) {
      fbthrift_serialize_and_send_getDataById(*rpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id);
    } else {
      fbthrift_serialize_and_send_getDataById(*defaultRpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (ctx != nullptr) {
      ctx->processClientInterceptorsOnResponse(returnState.header()).throwUnlessValue();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    ::std::string _return;
    if (auto ew = recv_wrapped_getDataById(_return, returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
    co_return _return;
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual void getDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id);


  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  static folly::exception_wrapper recv_wrapped_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  static void recv_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual void recv_instance_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "getDataById"} */
  virtual folly::exception_wrapper recv_instance_wrapped_getDataById(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
 private:
  apache::thrift::SerializedRequest fbthrift_serialize_getDataById(const RpcOptions& rpcOptions, apache::thrift::transport::THeader& header, apache::thrift::ContextStack* contextStack, ::std::int64_t p_id);
  template <typename RpcOptions>
  void fbthrift_send_getDataById(apache::thrift::SerializedRequest&& request, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::RequestClientCallback::Ptr callback, std::unique_ptr<folly::IOBuf> interceptorFrameworkMetadata);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> getDataByIdCtx(apache::thrift::RpcOptions* rpcOptions);
  template <typename CallbackType>
  folly::SemiFuture<::std::string> fbthrift_semifuture_getDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id);
 public:
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual void putDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual void putDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data);
 protected:
  void fbthrift_serialize_and_send_putDataById(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, const ::std::string& p_data, bool stealRpcOptions = false);
 public:

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual void sync_putDataById(::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual void sync_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual folly::Future<folly::Unit> future_putDataById(::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual folly::SemiFuture<folly::Unit> semifuture_putDataById(::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual folly::Future<folly::Unit> future_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual folly::SemiFuture<folly::Unit> semifuture_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  FOLLY_NODISCARD [[deprecated("To be replaced by new API soon")]] virtual folly::Try<apache::thrift::RpcResponseComplete<void>> sync_complete_putDataById(
      apache::thrift::RpcOptions&& rpcOptions,  ::std::int64_t p_id, const ::std::string& p_data);

#if FOLLY_HAS_COROUTINES
#if __clang__
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  template <int = 0>
  folly::coro::Task<void> co_putDataById(::std::int64_t p_id, const ::std::string& p_data) {
    return co_putDataById<false>(nullptr, p_id, p_data);
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  template <int = 0>
  folly::coro::Task<void> co_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
    return co_putDataById<true>(&rpcOptions, p_id, p_data);
  }
#else
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  folly::coro::Task<void> co_putDataById(::std::int64_t p_id, const ::std::string& p_data) {
    co_await folly::coro::detachOnCancel(semifuture_putDataById(p_id, p_data));
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  folly::coro::Task<void> co_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
    co_await folly::coro::detachOnCancel(semifuture_putDataById(rpcOptions, p_id, p_data));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<void> co_putDataById(apache::thrift::RpcOptions* rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = putDataByIdCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if (ctx != nullptr) {
      auto argsAsRefs = std::tie(p_id, p_data);
      ctx->processClientInterceptorsOnRequest(apache::thrift::ClientInterceptorOnRequestArguments(argsAsRefs), header.get(), hasRpcOptions ? *rpcOptions : *defaultRpcOptions).throwUnlessValue();
    }
    if constexpr (hasRpcOptions) {
      fbthrift_serialize_and_send_putDataById(*rpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id, p_data);
    } else {
      fbthrift_serialize_and_send_putDataById(*defaultRpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id, p_data);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (ctx != nullptr) {
      ctx->processClientInterceptorsOnResponse(returnState.header()).throwUnlessValue();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    if (auto ew = recv_wrapped_putDataById(returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual void putDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id, const ::std::string& p_data);


  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  static folly::exception_wrapper recv_wrapped_putDataById(::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  static void recv_putDataById(::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual void recv_instance_putDataById(::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "putDataById"} */
  virtual folly::exception_wrapper recv_instance_wrapped_putDataById(::apache::thrift::ClientReceiveState& state);
 private:
  apache::thrift::SerializedRequest fbthrift_serialize_putDataById(const RpcOptions& rpcOptions, apache::thrift::transport::THeader& header, apache::thrift::ContextStack* contextStack, ::std::int64_t p_id, const ::std::string& p_data);
  template <typename RpcOptions>
  void fbthrift_send_putDataById(apache::thrift::SerializedRequest&& request, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::RequestClientCallback::Ptr callback, std::unique_ptr<folly::IOBuf> interceptorFrameworkMetadata);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> putDataByIdCtx(apache::thrift::RpcOptions* rpcOptions);
  template <typename CallbackType>
  folly::SemiFuture<folly::Unit> fbthrift_semifuture_putDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);
 public:
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual void lobDataById(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual void lobDataById(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int64_t p_id, const ::std::string& p_data);
 protected:
  void fbthrift_serialize_and_send_lobDataById(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int64_t p_id, const ::std::string& p_data, bool stealRpcOptions = false);
 public:

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual void sync_lobDataById(::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual void sync_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual folly::Future<folly::Unit> future_lobDataById(::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual folly::SemiFuture<folly::Unit> semifuture_lobDataById(::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual folly::Future<folly::Unit> future_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual folly::SemiFuture<folly::Unit> semifuture_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);

#if FOLLY_HAS_COROUTINES
#if __clang__
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  template <int = 0>
  folly::coro::Task<void> co_lobDataById(::std::int64_t p_id, const ::std::string& p_data) {
    return co_lobDataById<false>(nullptr, p_id, p_data);
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  template <int = 0>
  folly::coro::Task<void> co_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
    return co_lobDataById<true>(&rpcOptions, p_id, p_data);
  }
#else
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  folly::coro::Task<void> co_lobDataById(::std::int64_t p_id, const ::std::string& p_data) {
    co_await folly::coro::detachOnCancel(semifuture_lobDataById(p_id, p_data));
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  folly::coro::Task<void> co_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
    co_await folly::coro::detachOnCancel(semifuture_lobDataById(rpcOptions, p_id, p_data));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<void> co_lobDataById(apache::thrift::RpcOptions* rpcOptions, ::std::int64_t p_id, const ::std::string& p_data) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<true> callback(&returnState, co_await folly::coro::co_current_executor);
    auto [ctx, header] = lobDataByIdCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<true>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if (ctx != nullptr) {
      auto argsAsRefs = std::tie(p_id, p_data);
      ctx->processClientInterceptorsOnRequest(apache::thrift::ClientInterceptorOnRequestArguments(argsAsRefs), header.get(), hasRpcOptions ? *rpcOptions : *defaultRpcOptions).throwUnlessValue();
    }
    if constexpr (hasRpcOptions) {
      fbthrift_serialize_and_send_lobDataById(*rpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id, p_data);
    } else {
      fbthrift_serialize_and_send_lobDataById(*defaultRpcOptions, header, ctx.get(), std::move(wrappedCallback), p_id, p_data);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  /** Glean {"file": "thrift/compiler/test/fixtures/basic-stack-arguments/src/module.thrift", "service": "MyServiceFast", "function": "lobDataById"} */
  virtual void lobDataById(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int64_t p_id, const ::std::string& p_data);


 private:
  apache::thrift::SerializedRequest fbthrift_serialize_lobDataById(const RpcOptions& rpcOptions, apache::thrift::transport::THeader& header, apache::thrift::ContextStack* contextStack, ::std::int64_t p_id, const ::std::string& p_data);
  template <typename RpcOptions>
  void fbthrift_send_lobDataById(apache::thrift::SerializedRequest&& request, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::RequestClientCallback::Ptr callback, std::unique_ptr<folly::IOBuf> interceptorFrameworkMetadata);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> lobDataByIdCtx(apache::thrift::RpcOptions* rpcOptions);
  template <typename CallbackType>
  folly::SemiFuture<folly::Unit> fbthrift_semifuture_lobDataById(apache::thrift::RpcOptions& rpcOptions, ::std::int64_t p_id, const ::std::string& p_data);
 public:
};

} // namespace apache::thrift

namespace cpp2 {
using MyServiceFastAsyncClient [[deprecated("Use apache::thrift::Client<MyServiceFast> instead")]] = ::apache::thrift::Client<MyServiceFast>;
} // namespace cpp2
