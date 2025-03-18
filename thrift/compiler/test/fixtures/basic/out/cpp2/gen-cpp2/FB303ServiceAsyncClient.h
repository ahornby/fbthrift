/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/basic/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include <thrift/lib/cpp2/gen/client_h.h>

#include "thrift/compiler/test/fixtures/basic/gen-cpp2/module_types.h"

namespace apache { namespace thrift {
  class Cpp2RequestContext;
  namespace detail { namespace ac { struct ClientRequestContext; }}
  namespace transport { class THeader; }
}}

namespace test::fixtures::basic {
class FB303Service;
} // namespace test::fixtures::basic
namespace apache::thrift {

template <>
class Client<::test::fixtures::basic::FB303Service> : public apache::thrift::GeneratedAsyncClient {
 public:
  using apache::thrift::GeneratedAsyncClient::GeneratedAsyncClient;

  char const* getServiceName() const noexcept override {
    return "FB303Service";
  }

  static const char* __fbthrift_thrift_uri() {
    return "test.dev/fixtures/basic/FB303Service";
  }


  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual void simple_rpc(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int32_t p_int_parameter);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual void simple_rpc(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int32_t p_int_parameter);
 protected:
  void fbthrift_serialize_and_send_simple_rpc(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int32_t p_int_parameter, bool stealRpcOptions = false);
 public:

  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual void sync_simple_rpc(::test::fixtures::basic::ReservedKeyword& _return, ::std::int32_t p_int_parameter);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual void sync_simple_rpc(apache::thrift::RpcOptions& rpcOptions, ::test::fixtures::basic::ReservedKeyword& _return, ::std::int32_t p_int_parameter);

  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual folly::Future<::test::fixtures::basic::ReservedKeyword> future_simple_rpc(::std::int32_t p_int_parameter);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual folly::SemiFuture<::test::fixtures::basic::ReservedKeyword> semifuture_simple_rpc(::std::int32_t p_int_parameter);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual folly::Future<::test::fixtures::basic::ReservedKeyword> future_simple_rpc(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_int_parameter);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual folly::SemiFuture<::test::fixtures::basic::ReservedKeyword> semifuture_simple_rpc(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_int_parameter);

#if FOLLY_HAS_COROUTINES
#if __clang__
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  template <int = 0>
  folly::coro::Task<::test::fixtures::basic::ReservedKeyword> co_simple_rpc(::std::int32_t p_int_parameter) {
    return co_simple_rpc<false>(nullptr, p_int_parameter);
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  template <int = 0>
  folly::coro::Task<::test::fixtures::basic::ReservedKeyword> co_simple_rpc(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_int_parameter) {
    return co_simple_rpc<true>(&rpcOptions, p_int_parameter);
  }
#else
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  folly::coro::Task<::test::fixtures::basic::ReservedKeyword> co_simple_rpc(::std::int32_t p_int_parameter) {
    co_return co_await folly::coro::detachOnCancel(semifuture_simple_rpc(p_int_parameter));
  }
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  folly::coro::Task<::test::fixtures::basic::ReservedKeyword> co_simple_rpc(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_int_parameter) {
    co_return co_await folly::coro::detachOnCancel(semifuture_simple_rpc(rpcOptions, p_int_parameter));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<::test::fixtures::basic::ReservedKeyword> co_simple_rpc(apache::thrift::RpcOptions* rpcOptions, ::std::int32_t p_int_parameter) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = simple_rpcCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if (ctx != nullptr) {
      auto argsAsRefs = std::tie(p_int_parameter);
      ctx->processClientInterceptorsOnRequest(apache::thrift::ClientInterceptorOnRequestArguments(argsAsRefs), header.get(), hasRpcOptions ? *rpcOptions : *defaultRpcOptions).throwUnlessValue();
    }
    if constexpr (hasRpcOptions) {
      fbthrift_serialize_and_send_simple_rpc(*rpcOptions, header, ctx.get(), std::move(wrappedCallback), p_int_parameter);
    } else {
      fbthrift_serialize_and_send_simple_rpc(*defaultRpcOptions, header, ctx.get(), std::move(wrappedCallback), p_int_parameter);
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
    ::test::fixtures::basic::ReservedKeyword _return;
    if (auto ew = recv_wrapped_simple_rpc(_return, returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
    co_return _return;
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual void simple_rpc(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int32_t p_int_parameter);


  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  static folly::exception_wrapper recv_wrapped_simple_rpc(::test::fixtures::basic::ReservedKeyword& _return, ::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  static void recv_simple_rpc(::test::fixtures::basic::ReservedKeyword& _return, ::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual void recv_instance_simple_rpc(::test::fixtures::basic::ReservedKeyword& _return, ::apache::thrift::ClientReceiveState& state);
  /** Glean {"file": "thrift/compiler/test/fixtures/basic/src/module.thrift", "service": "FB303Service", "function": "simple_rpc"} */
  virtual folly::exception_wrapper recv_instance_wrapped_simple_rpc(::test::fixtures::basic::ReservedKeyword& _return, ::apache::thrift::ClientReceiveState& state);
 private:
  apache::thrift::SerializedRequest fbthrift_serialize_simple_rpc(const RpcOptions& rpcOptions, apache::thrift::transport::THeader& header, apache::thrift::ContextStack* contextStack, ::std::int32_t p_int_parameter);
  template <typename RpcOptions>
  void fbthrift_send_simple_rpc(apache::thrift::SerializedRequest&& request, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::RequestClientCallback::Ptr callback, std::unique_ptr<folly::IOBuf> interceptorFrameworkMetadata);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> simple_rpcCtx(apache::thrift::RpcOptions* rpcOptions);
  template <typename CallbackType>
  folly::SemiFuture<::test::fixtures::basic::ReservedKeyword> fbthrift_semifuture_simple_rpc(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_int_parameter);
 public:
};

} // namespace apache::thrift

namespace test::fixtures::basic {
using FB303ServiceAsyncClient [[deprecated("Use apache::thrift::Client<FB303Service> instead")]] = ::apache::thrift::Client<FB303Service>;
} // namespace test::fixtures::basic
