/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/interactions/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include "thrift/compiler/test/fixtures/interactions/gen-cpp2/BoxService.h"

#include <thrift/lib/cpp2/gen/service_tcc.h>

namespace cpp2 {
typedef apache::thrift::ThriftPresult<false, apache::thrift::FieldData<1, ::apache::thrift::type_class::structure, ::cpp2::ShouldBeBoxed*>> BoxService_getABoxSession_pargs;
typedef apache::thrift::ThriftPresult<true, apache::thrift::FieldData<0, ::apache::thrift::type_class::structure, ::cpp2::ShouldBeBoxed*>> BoxService_getABoxSession_presult;
template <typename ProtocolIn_, typename ProtocolOut_>
void BoxServiceAsyncProcessor::setUpAndProcess_getABoxSession(apache::thrift::ResponseChannelRequest::UniquePtr req, apache::thrift::SerializedCompressedRequest&& serializedRequest, apache::thrift::Cpp2RequestContext* ctx, folly::EventBase* eb, [[maybe_unused]] apache::thrift::concurrency::ThreadManager* tm) {
  if (!setUpRequestProcessing(req, ctx, eb, tm, apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, iface_, "BoxedInteraction", true)) {
    return;
  }
  auto scope = iface_->getRequestExecutionScope(ctx, apache::thrift::concurrency::NORMAL);
  ctx->setRequestExecutionScope(std::move(scope));
  processInThread(std::move(req), std::move(serializedRequest), ctx, eb, tm, apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, &BoxServiceAsyncProcessor::executeRequest_getABoxSession<ProtocolIn_, ProtocolOut_>, this);
}

template <typename ProtocolIn_, typename ProtocolOut_>
void BoxServiceAsyncProcessor::executeRequest_getABoxSession(apache::thrift::ServerRequest&& serverRequest) {
  auto tile = serverRequest.requestContext()->releaseTile();
  // make sure getRequestContext is null
  // so async calls don't accidentally use it
  iface_->setRequestContext(nullptr);
  struct ArgsState {
    std::unique_ptr<::cpp2::ShouldBeBoxed> uarg_req = std::make_unique<::cpp2::ShouldBeBoxed>();
    BoxService_getABoxSession_pargs pargs() {
      BoxService_getABoxSession_pargs args;
      args.get<0>().value = uarg_req.get();
      return args;
    }

    auto asTupleOfRefs() & {
      return std::tie(
        std::as_const(*uarg_req)
      );
    }
  } args;

  auto ctxStack = apache::thrift::ContextStack::create(
    this->getEventHandlersSharedPtr(),
    this->getServiceName(),
    "BoxService.getABoxSession",
    serverRequest.requestContext());
  try {
    auto pargs = args.pargs();
    deserializeRequest<ProtocolIn_>(pargs, "getABoxSession", apache::thrift::detail::ServerRequestHelper::compressedRequest(std::move(serverRequest)).uncompress(), ctxStack.get());
  }
  catch (...) {
    folly::exception_wrapper ew(std::current_exception());
    apache::thrift::detail::ap::process_handle_exn_deserialization<ProtocolOut_>(
        ew
        , apache::thrift::detail::ServerRequestHelper::request(std::move(serverRequest))
        , serverRequest.requestContext()
        , apache::thrift::detail::ServerRequestHelper::eventBase(serverRequest)
        , "getABoxSession");
    return;
  }
  auto requestPileNotification = apache::thrift::detail::ServerRequestHelper::moveRequestPileNotification(serverRequest);
  auto concurrencyControllerNotification = apache::thrift::detail::ServerRequestHelper::moveConcurrencyControllerNotification(serverRequest);
  auto callback = apache::thrift::HandlerCallbackPtr<apache::thrift::TileAndResponse<apache::thrift::ServiceHandler<::cpp2::BoxService>::BoxedInteractionIf, std::unique_ptr<::cpp2::ShouldBeBoxed>>>::make(
    apache::thrift::detail::ServerRequestHelper::request(std::move(serverRequest))
    , std::move(ctxStack)
    , this->getServiceName()
    , "BoxService"
    , "getABoxSession"
    , return_getABoxSession<ProtocolIn_,ProtocolOut_>
    , throw_wrapped_getABoxSession<ProtocolIn_, ProtocolOut_>
    , serverRequest.requestContext()->getProtoSeqId()
    , apache::thrift::detail::ServerRequestHelper::eventBase(serverRequest)
    , apache::thrift::detail::ServerRequestHelper::executor(serverRequest)
    , serverRequest.requestContext()
    , requestPileNotification
    , concurrencyControllerNotification, std::move(serverRequest.requestData())
    , std::move(tile));
  const auto makeExecuteHandler = [&] {
    return [ifacePtr = iface_](auto&& cb, ArgsState args) mutable {
      (void)args;
      ifacePtr->async_tm_getABoxSession(std::move(cb), std::move(args.uarg_req));
    };
  };
#if FOLLY_HAS_COROUTINES
  if (apache::thrift::detail::shouldProcessServiceInterceptorsOnRequest(*callback)) {
    [](auto callback, auto executeHandler, ArgsState args) -> folly::coro::Task<void> {
      auto argRefs = args.asTupleOfRefs();
      co_await apache::thrift::detail::processServiceInterceptorsOnRequest(
          *callback,
          apache::thrift::detail::ServiceInterceptorOnRequestArguments(argRefs));
      executeHandler(std::move(callback), std::move(args));
    }(std::move(callback), makeExecuteHandler(), std::move(args))
              .scheduleOn(apache::thrift::detail::ServerRequestHelper::executor(serverRequest))
              .startInlineUnsafe();
  } else {
    makeExecuteHandler()(std::move(callback), std::move(args));
  }
#else
  makeExecuteHandler()(std::move(callback), std::move(args));
#endif // FOLLY_HAS_COROUTINES
}

template <class ProtocolIn_, class ProtocolOut_>
apache::thrift::SerializedResponse BoxServiceAsyncProcessor::return_getABoxSession(apache::thrift::ContextStack* ctx, ::cpp2::ShouldBeBoxed const& _return) {
  ProtocolOut_ prot;
  ::cpp2::BoxService_getABoxSession_presult result;
  result.get<0>().value = const_cast<::cpp2::ShouldBeBoxed*>(&_return);
  result.setIsSet(0, true);
  return serializeResponse("getABoxSession", &prot, ctx, result);
}

template <class ProtocolIn_, class ProtocolOut_>
void BoxServiceAsyncProcessor::throw_wrapped_getABoxSession(apache::thrift::ResponseChannelRequest::UniquePtr req,[[maybe_unused]] int32_t protoSeqId,apache::thrift::ContextStack* ctx,folly::exception_wrapper ew,apache::thrift::Cpp2RequestContext* reqCtx) {
  if (!ew) {
    return;
  }
  {
    apache::thrift::detail::ap::process_throw_wrapped_handler_error<ProtocolOut_>(
        ew, std::move(req), reqCtx, ctx, "getABoxSession");
    return;
  }
}


typedef apache::thrift::ThriftPresult<false> BoxService_BoxedInteraction_getABox_pargs;
typedef apache::thrift::ThriftPresult<true, apache::thrift::FieldData<0, ::apache::thrift::type_class::structure, ::cpp2::ShouldBeBoxed*>> BoxService_BoxedInteraction_getABox_presult;
template <typename ProtocolIn_, typename ProtocolOut_>
void BoxServiceAsyncProcessor::setUpAndProcess_BoxedInteraction_getABox(apache::thrift::ResponseChannelRequest::UniquePtr req, apache::thrift::SerializedCompressedRequest&& serializedRequest, apache::thrift::Cpp2RequestContext* ctx, folly::EventBase* eb, [[maybe_unused]] apache::thrift::concurrency::ThreadManager* tm) {
  if (!setUpRequestProcessing(req, ctx, eb, tm, apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, iface_, "BoxedInteraction")) {
    return;
  }
  auto scope = iface_->getRequestExecutionScope(ctx, apache::thrift::concurrency::NORMAL);
  ctx->setRequestExecutionScope(std::move(scope));
  processInThread(std::move(req), std::move(serializedRequest), ctx, eb, tm, apache::thrift::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE, &BoxServiceAsyncProcessor::executeRequest_BoxedInteraction_getABox<ProtocolIn_, ProtocolOut_>, this);
}

template <typename ProtocolIn_, typename ProtocolOut_>
void BoxServiceAsyncProcessor::executeRequest_BoxedInteraction_getABox(apache::thrift::ServerRequest&& serverRequest) {
  auto tile = serverRequest.requestContext()->releaseTile();
  // make sure getRequestContext is null
  // so async calls don't accidentally use it
  iface_->setRequestContext(nullptr);
  struct ArgsState {
    BoxService_BoxedInteraction_getABox_pargs pargs() {
      BoxService_BoxedInteraction_getABox_pargs args;
      return args;
    }

    auto asTupleOfRefs() & {
      return std::tie(
      );
    }
  } args;

  auto ctxStack = apache::thrift::ContextStack::create(
    this->getEventHandlersSharedPtr(),
    this->getServiceName(),
    "BoxService.BoxedInteraction.getABox",
    serverRequest.requestContext());
  auto& iface = static_cast<apache::thrift::ServiceHandler<BoxService>::BoxedInteractionIf&>(*tile);
  try {
    auto pargs = args.pargs();
    deserializeRequest<ProtocolIn_>(pargs, "BoxedInteraction.getABox", apache::thrift::detail::ServerRequestHelper::compressedRequest(std::move(serverRequest)).uncompress(), ctxStack.get());
  }
  catch (...) {
    folly::exception_wrapper ew(std::current_exception());
    apache::thrift::detail::ap::process_handle_exn_deserialization<ProtocolOut_>(
        ew
        , apache::thrift::detail::ServerRequestHelper::request(std::move(serverRequest))
        , serverRequest.requestContext()
        , apache::thrift::detail::ServerRequestHelper::eventBase(serverRequest)
        , "BoxedInteraction.getABox");
    return;
  }
  auto requestPileNotification = apache::thrift::detail::ServerRequestHelper::moveRequestPileNotification(serverRequest);
  auto concurrencyControllerNotification = apache::thrift::detail::ServerRequestHelper::moveConcurrencyControllerNotification(serverRequest);
  auto callback = apache::thrift::HandlerCallbackPtr<std::unique_ptr<::cpp2::ShouldBeBoxed>>::make(
    apache::thrift::detail::ServerRequestHelper::request(std::move(serverRequest))
    , std::move(ctxStack)
    , this->getServiceName()
    , "BoxService"
    , "BoxedInteraction.getABox"
    , return_BoxedInteraction_getABox<ProtocolIn_,ProtocolOut_>
    , throw_wrapped_BoxedInteraction_getABox<ProtocolIn_, ProtocolOut_>
    , serverRequest.requestContext()->getProtoSeqId()
    , apache::thrift::detail::ServerRequestHelper::eventBase(serverRequest)
    , apache::thrift::detail::ServerRequestHelper::executor(serverRequest)
    , serverRequest.requestContext()
    , requestPileNotification
    , concurrencyControllerNotification, std::move(serverRequest.requestData())
    , std::move(tile));
  const auto makeExecuteHandler = [&] {
    return [ifacePtr = &iface](auto&& cb, ArgsState args) mutable {
      (void)args;
      ifacePtr->async_tm_getABox(std::move(cb));
    };
  };
#if FOLLY_HAS_COROUTINES
  if (apache::thrift::detail::shouldProcessServiceInterceptorsOnRequest(*callback)) {
    [](auto callback, auto executeHandler, ArgsState args) -> folly::coro::Task<void> {
      auto argRefs = args.asTupleOfRefs();
      co_await apache::thrift::detail::processServiceInterceptorsOnRequest(
          *callback,
          apache::thrift::detail::ServiceInterceptorOnRequestArguments(argRefs));
      executeHandler(std::move(callback), std::move(args));
    }(std::move(callback), makeExecuteHandler(), std::move(args))
              .scheduleOn(apache::thrift::detail::ServerRequestHelper::executor(serverRequest))
              .startInlineUnsafe();
  } else {
    makeExecuteHandler()(std::move(callback), std::move(args));
  }
#else
  makeExecuteHandler()(std::move(callback), std::move(args));
#endif // FOLLY_HAS_COROUTINES
}

template <class ProtocolIn_, class ProtocolOut_>
apache::thrift::SerializedResponse BoxServiceAsyncProcessor::return_BoxedInteraction_getABox(apache::thrift::ContextStack* ctx, ::cpp2::ShouldBeBoxed const& _return) {
  ProtocolOut_ prot;
  ::cpp2::BoxService_BoxedInteraction_getABox_presult result;
  result.get<0>().value = const_cast<::cpp2::ShouldBeBoxed*>(&_return);
  result.setIsSet(0, true);
  return serializeResponse("BoxedInteraction.getABox", &prot, ctx, result);
}

template <class ProtocolIn_, class ProtocolOut_>
void BoxServiceAsyncProcessor::throw_wrapped_BoxedInteraction_getABox(apache::thrift::ResponseChannelRequest::UniquePtr req,[[maybe_unused]] int32_t protoSeqId,apache::thrift::ContextStack* ctx,folly::exception_wrapper ew,apache::thrift::Cpp2RequestContext* reqCtx) {
  if (!ew) {
    return;
  }
  {
    apache::thrift::detail::ap::process_throw_wrapped_handler_error<ProtocolOut_>(
        ew, std::move(req), reqCtx, ctx, "BoxedInteraction.getABox");
    return;
  }
}

} // namespace cpp2
