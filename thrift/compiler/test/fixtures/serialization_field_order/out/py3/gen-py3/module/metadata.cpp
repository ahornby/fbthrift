/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/serialization_field_order/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/serialization_field_order/gen-py3/module/metadata.h"

namespace cpp2 {
::apache::thrift::metadata::ThriftMetadata module_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftServiceMetadataResponse response;
  ::apache::thrift::metadata::ThriftMetadata& metadata = *response.metadata_ref();
  ::apache::thrift::detail::md::StructMetadata<Foo>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<Foo2>::gen(metadata);
  return metadata;
}
} // namespace cpp2
