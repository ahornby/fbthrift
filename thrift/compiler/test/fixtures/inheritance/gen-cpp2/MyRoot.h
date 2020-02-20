/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#pragma once

#include <thrift/lib/cpp2/gen/service_h.h>

#include "thrift/compiler/test/fixtures/inheritance/gen-cpp2/MyRootAsyncClient.h"
#include "thrift/compiler/test/fixtures/inheritance/gen-cpp2/module_types.h"

namespace folly {
  class IOBuf;
  class IOBufQueue;
}
namespace apache { namespace thrift {
  class Cpp2RequestContext;
  class BinaryProtocolReader;
  class CompactProtocolReader;
  namespace transport { class THeader; }
}}

namespace cpp2 {

class MyRootSvAsyncIf {
 public:
  virtual ~MyRootSvAsyncIf() {}
  virtual void async_tm_do_root(std::unique_ptr<apache::thrift::HandlerCallback<void>> callback) = 0;
  virtual folly::Future<folly::Unit> future_do_root() = 0;
  virtual folly::SemiFuture<folly::Unit> semifuture_do_root() = 0;
};

class MyRootAsyncProcessor;

class MyRootSvIf : public MyRootSvAsyncIf, public apache::thrift::ServerInterface {
 public:
  typedef MyRootAsyncProcessor ProcessorType;
  std::unique_ptr<apache::thrift::AsyncProcessor> getProcessor() override;
  virtual void do_root();
  folly::Future<folly::Unit> future_do_root() override;
  folly::SemiFuture<folly::Unit> semifuture_do_root() override;
  void async_tm_do_root(std::unique_ptr<apache::thrift::HandlerCallback<void>> callback) override;
};

class MyRootSvNull : public MyRootSvIf {
 public:
  void do_root() override;
};

class MyRootAsyncProcessor : public ::apache::thrift::GeneratedAsyncProcessor {
 public:
  const char* getServiceName() override;
  void getServiceMetadata(apache::thrift::metadata::ThriftServiceMetadataResponse& response) override;
  using BaseAsyncProcessor = void;
 protected:
  MyRootSvIf* iface_;
 public:
  void process(apache::thrift::ResponseChannelRequest::UniquePtr req, std::unique_ptr<folly::IOBuf> buf, apache::thrift::protocol::PROTOCOL_TYPES protType, apache::thrift::Cpp2RequestContext* context, folly::EventBase* eb, apache::thrift::concurrency::ThreadManager* tm) override;
 protected:
  bool isOnewayMethod(const folly::IOBuf* buf, const apache::thrift::transport::THeader* header) override;
  std::shared_ptr<folly::RequestContext> getBaseContextForRequest() override;
 private:
  static std::unordered_set<std::string> onewayMethods_;
 public:
  using ProcessFunc = GeneratedAsyncProcessor::ProcessFunc<MyRootAsyncProcessor>;
  using ProcessMap = GeneratedAsyncProcessor::ProcessMap<ProcessFunc>;
  static const MyRootAsyncProcessor::ProcessMap& getBinaryProtocolProcessMap();
  static const MyRootAsyncProcessor::ProcessMap& getCompactProtocolProcessMap();
 private:
  static const MyRootAsyncProcessor::ProcessMap binaryProcessMap_;
   static const MyRootAsyncProcessor::ProcessMap compactProcessMap_;
 private:
  template <typename ProtocolIn_, typename ProtocolOut_>
  void _processInThread_do_root(apache::thrift::ResponseChannelRequest::UniquePtr req, std::unique_ptr<folly::IOBuf> buf, apache::thrift::Cpp2RequestContext* ctx, folly::EventBase* eb, apache::thrift::concurrency::ThreadManager* tm);
  template <typename ProtocolIn_, typename ProtocolOut_>
  void process_do_root(apache::thrift::ResponseChannelRequest::UniquePtr req, std::unique_ptr<folly::IOBuf> buf,apache::thrift::Cpp2RequestContext* ctx,folly::EventBase* eb, apache::thrift::concurrency::ThreadManager* tm);
  template <class ProtocolIn_, class ProtocolOut_>
  static folly::IOBufQueue return_do_root(int32_t protoSeqId, apache::thrift::ContextStack* ctx);
  template <class ProtocolIn_, class ProtocolOut_>
  static void throw_wrapped_do_root(apache::thrift::ResponseChannelRequest::UniquePtr req,int32_t protoSeqId,apache::thrift::ContextStack* ctx,folly::exception_wrapper ew,apache::thrift::Cpp2RequestContext* reqCtx);
 public:
  MyRootAsyncProcessor(MyRootSvIf* iface) :
      iface_(iface) {}

  virtual ~MyRootAsyncProcessor() {}
};

} // cpp2
