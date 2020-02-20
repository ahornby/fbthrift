/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "thrift/compiler/test/fixtures/exceptions/gen-cpp2/Raiser.h"
#include "thrift/compiler/test/fixtures/exceptions/gen-cpp2/Raiser.tcc"
#include "thrift/compiler/test/fixtures/exceptions/gen-cpp2/module_metadata.h"
#include <thrift/lib/cpp2/gen/service_cpp.h>

namespace cpp2 {
std::unique_ptr<apache::thrift::AsyncProcessor> RaiserSvIf::getProcessor() {
  return std::make_unique<RaiserAsyncProcessor>(this);
}


void RaiserSvIf::doBland() {
  apache::thrift::detail::si::throw_app_exn_unimplemented("doBland");
}

folly::SemiFuture<folly::Unit> RaiserSvIf::semifuture_doBland() {
  return apache::thrift::detail::si::semifuture([&] { return doBland(); });
}

folly::Future<folly::Unit> RaiserSvIf::future_doBland() {
  return apache::thrift::detail::si::future(semifuture_doBland(), getThreadManager());
}


void RaiserSvIf::async_tm_doBland(std::unique_ptr<apache::thrift::HandlerCallback<void>> callback) {
  apache::thrift::detail::si::async_tm(this, std::move(callback), [&] { return future_doBland(); });
}

void RaiserSvIf::doRaise() {
  apache::thrift::detail::si::throw_app_exn_unimplemented("doRaise");
}

folly::SemiFuture<folly::Unit> RaiserSvIf::semifuture_doRaise() {
  return apache::thrift::detail::si::semifuture([&] { return doRaise(); });
}

folly::Future<folly::Unit> RaiserSvIf::future_doRaise() {
  return apache::thrift::detail::si::future(semifuture_doRaise(), getThreadManager());
}


void RaiserSvIf::async_tm_doRaise(std::unique_ptr<apache::thrift::HandlerCallback<void>> callback) {
  apache::thrift::detail::si::async_tm(this, std::move(callback), [&] { return future_doRaise(); });
}

void RaiserSvIf::get200(::std::string& /*_return*/) {
  apache::thrift::detail::si::throw_app_exn_unimplemented("get200");
}

folly::SemiFuture<std::unique_ptr<::std::string>> RaiserSvIf::semifuture_get200() {
  return apache::thrift::detail::si::semifuture_returning_uptr([&](::std::string& _return) { get200(_return); });
}

folly::Future<std::unique_ptr<::std::string>> RaiserSvIf::future_get200() {
  return apache::thrift::detail::si::future(semifuture_get200(), getThreadManager());
}


void RaiserSvIf::async_tm_get200(std::unique_ptr<apache::thrift::HandlerCallback<std::unique_ptr<::std::string>>> callback) {
  apache::thrift::detail::si::async_tm(this, std::move(callback), [&] { return future_get200(); });
}


void RaiserSvIf::get500(::std::string& /*_return*/) {
  apache::thrift::detail::si::throw_app_exn_unimplemented("get500");
}

folly::SemiFuture<std::unique_ptr<::std::string>> RaiserSvIf::semifuture_get500() {
  return apache::thrift::detail::si::semifuture_returning_uptr([&](::std::string& _return) { get500(_return); });
}

folly::Future<std::unique_ptr<::std::string>> RaiserSvIf::future_get500() {
  return apache::thrift::detail::si::future(semifuture_get500(), getThreadManager());
}


void RaiserSvIf::async_tm_get500(std::unique_ptr<apache::thrift::HandlerCallback<std::unique_ptr<::std::string>>> callback) {
  apache::thrift::detail::si::async_tm(this, std::move(callback), [&] { return future_get500(); });
}


void RaiserSvNull::doBland() {
  return;
}

void RaiserSvNull::doRaise() {
  return;
}

void RaiserSvNull::get200(::std::string& /*_return*/) {}

void RaiserSvNull::get500(::std::string& /*_return*/) {}

const char* RaiserAsyncProcessor::getServiceName() {
  return "Raiser";
}

void RaiserAsyncProcessor::getServiceMetadata(apache::thrift::metadata::ThriftServiceMetadataResponse& response) {
  ::apache::thrift::detail::md::ServiceMetadata<RaiserSvIf>::gen(response.metadata, response.context);
}

void RaiserAsyncProcessor::process(apache::thrift::ResponseChannelRequest::UniquePtr req, std::unique_ptr<folly::IOBuf> buf, apache::thrift::protocol::PROTOCOL_TYPES protType, apache::thrift::Cpp2RequestContext* context, folly::EventBase* eb, apache::thrift::concurrency::ThreadManager* tm) {
  apache::thrift::detail::ap::process(this, std::move(req), std::move(buf), protType, context, eb, tm);
}

bool RaiserAsyncProcessor::isOnewayMethod(const folly::IOBuf* buf, const apache::thrift::transport::THeader* header) {
  return apache::thrift::detail::ap::is_oneway_method(buf, header, onewayMethods_);
}

std::shared_ptr<folly::RequestContext> RaiserAsyncProcessor::getBaseContextForRequest() {
  return iface_->getBaseContextForRequest();
}

std::unordered_set<std::string> RaiserAsyncProcessor::onewayMethods_ {};
const RaiserAsyncProcessor::ProcessMap& RaiserAsyncProcessor::getBinaryProtocolProcessMap() {
  return binaryProcessMap_;
}

const RaiserAsyncProcessor::ProcessMap RaiserAsyncProcessor::binaryProcessMap_ {
  {"doBland", &RaiserAsyncProcessor::_processInThread_doBland<apache::thrift::BinaryProtocolReader, apache::thrift::BinaryProtocolWriter>},
  {"doRaise", &RaiserAsyncProcessor::_processInThread_doRaise<apache::thrift::BinaryProtocolReader, apache::thrift::BinaryProtocolWriter>},
  {"get200", &RaiserAsyncProcessor::_processInThread_get200<apache::thrift::BinaryProtocolReader, apache::thrift::BinaryProtocolWriter>},
  {"get500", &RaiserAsyncProcessor::_processInThread_get500<apache::thrift::BinaryProtocolReader, apache::thrift::BinaryProtocolWriter>},
};

const RaiserAsyncProcessor::ProcessMap& RaiserAsyncProcessor::getCompactProtocolProcessMap() {
  return compactProcessMap_;
}

const RaiserAsyncProcessor::ProcessMap RaiserAsyncProcessor::compactProcessMap_ {
  {"doBland", &RaiserAsyncProcessor::_processInThread_doBland<apache::thrift::CompactProtocolReader, apache::thrift::CompactProtocolWriter>},
  {"doRaise", &RaiserAsyncProcessor::_processInThread_doRaise<apache::thrift::CompactProtocolReader, apache::thrift::CompactProtocolWriter>},
  {"get200", &RaiserAsyncProcessor::_processInThread_get200<apache::thrift::CompactProtocolReader, apache::thrift::CompactProtocolWriter>},
  {"get500", &RaiserAsyncProcessor::_processInThread_get500<apache::thrift::CompactProtocolReader, apache::thrift::CompactProtocolWriter>},
};

} // cpp2
