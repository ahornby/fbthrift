/**
 * Autogenerated by Thrift for c.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/transitive-deps/gen-py3/c/metadata.h"

namespace cpp2 {
::apache::thrift::metadata::ThriftMetadata c_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftServiceMetadataResponse response;
  ::apache::thrift::metadata::ThriftMetadata& metadata = *response.metadata_ref();
  ::apache::thrift::detail::md::StructMetadata<C>::gen(metadata);
  ::apache::thrift::detail::md::ExceptionMetadata<E>::gen(metadata);
  return metadata;
}
} // namespace cpp2
