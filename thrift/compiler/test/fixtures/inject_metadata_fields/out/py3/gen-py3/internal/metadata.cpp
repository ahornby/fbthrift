/**
 * Autogenerated by Thrift for thrift/annotation/internal.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/inject_metadata_fields/gen-py3/internal/metadata.h"

namespace facebook {
namespace thrift {
namespace annotation {
::apache::thrift::metadata::ThriftMetadata internal_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftServiceMetadataResponse response;
  ::apache::thrift::metadata::ThriftMetadata& metadata = *response.metadata_ref();
  ::apache::thrift::detail::md::StructMetadata<InjectMetadataFields>::gen(metadata);
  return metadata;
}
} // namespace facebook
} // namespace thrift
} // namespace annotation
