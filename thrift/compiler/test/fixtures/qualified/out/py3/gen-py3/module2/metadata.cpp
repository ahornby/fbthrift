/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/qualified/src/module2.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/qualified/gen-py3/module2/metadata.h"

namespace module2 {
::apache::thrift::metadata::ThriftMetadata module2_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftServiceMetadataResponse response;
  ::apache::thrift::metadata::ThriftMetadata& metadata = *response.metadata_ref();
  ::apache::thrift::detail::md::StructMetadata<Struct>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<BigStruct>::gen(metadata);
  return metadata;
}
} // namespace module2
