/**
 * Autogenerated by Thrift for transitive.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/includes/gen-py3/transitive/metadata.h"

namespace cpp2 {
::apache::thrift::metadata::ThriftMetadata transitive_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftServiceMetadataResponse response;
  ::apache::thrift::metadata::ThriftMetadata& metadata = *response.metadata_ref();
  ::apache::thrift::detail::md::StructMetadata<Foo>::gen(metadata);
  return metadata;
}
} // namespace cpp2
