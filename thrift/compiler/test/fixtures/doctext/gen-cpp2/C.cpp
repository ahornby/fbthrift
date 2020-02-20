/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "thrift/compiler/test/fixtures/doctext/gen-cpp2/C.h"
#include "thrift/compiler/test/fixtures/doctext/gen-cpp2/C.tcc"
#include "thrift/compiler/test/fixtures/doctext/gen-cpp2/module_metadata.h"
#include <thrift/lib/cpp2/gen/service_cpp.h>

namespace cpp2 {
std::unique_ptr<apache::thrift::AsyncProcessor> CSvIf::getProcessor() {
  return std::make_unique<CAsyncProcessor>(this);
}


void CSvIf::f() {
  apache::thrift::detail::si::throw_app_exn_unimplemented("f");
}

folly::SemiFuture<folly::Unit> CSvIf::semifuture_f() {
  return apache::thrift::detail::si::semifuture([&] { return f(); });
}

folly::Future<folly::Unit> CSvIf::future_f() {
  return apache::thrift::detail::si::future(semifuture_f(), getThreadManager());
}


void CSvIf::async_tm_f(std::unique_ptr<apache::thrift::HandlerCallback<void>> callback) {
  apache::thrift::detail::si::async_tm(this, std::move(callback), [&] { return future_f(); });
}

void CSvNull::f() {
  return;
}

const char* CAsyncProcessor::getServiceName() {
  return "C";
}

void CAsyncProcessor::getServiceMetadata(apache::thrift::metadata::ThriftServiceMetadataResponse& response) {
  ::apache::thrift::detail::md::ServiceMetadata<CSvIf>::gen(response.metadata, response.context);
}

void CAsyncProcessor::process(apache::thrift::ResponseChannelRequest::UniquePtr req, std::unique_ptr<folly::IOBuf> buf, apache::thrift::protocol::PROTOCOL_TYPES protType, apache::thrift::Cpp2RequestContext* context, folly::EventBase* eb, apache::thrift::concurrency::ThreadManager* tm) {
  apache::thrift::detail::ap::process(this, std::move(req), std::move(buf), protType, context, eb, tm);
}

bool CAsyncProcessor::isOnewayMethod(const folly::IOBuf* buf, const apache::thrift::transport::THeader* header) {
  return apache::thrift::detail::ap::is_oneway_method(buf, header, onewayMethods_);
}

std::shared_ptr<folly::RequestContext> CAsyncProcessor::getBaseContextForRequest() {
  return iface_->getBaseContextForRequest();
}

std::unordered_set<std::string> CAsyncProcessor::onewayMethods_ {};
const CAsyncProcessor::ProcessMap& CAsyncProcessor::getBinaryProtocolProcessMap() {
  return binaryProcessMap_;
}

const CAsyncProcessor::ProcessMap CAsyncProcessor::binaryProcessMap_ {
  {"f", &CAsyncProcessor::_processInThread_f<apache::thrift::BinaryProtocolReader, apache::thrift::BinaryProtocolWriter>},
};

const CAsyncProcessor::ProcessMap& CAsyncProcessor::getCompactProtocolProcessMap() {
  return compactProcessMap_;
}

const CAsyncProcessor::ProcessMap CAsyncProcessor::compactProcessMap_ {
  {"f", &CAsyncProcessor::_processInThread_f<apache::thrift::CompactProtocolReader, apache::thrift::CompactProtocolWriter>},
};

} // cpp2
