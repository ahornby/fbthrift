/**
 * Autogenerated by Thrift for thrift/annotation/python.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "thrift/compiler/test/fixtures/exceptions/gen-py3/python/metadata.h"

namespace facebook {
namespace thrift {
namespace annotation {
namespace python {
::apache::thrift::metadata::ThriftMetadata python_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftServiceMetadataResponse response;
  ::apache::thrift::metadata::ThriftMetadata& metadata = *response.metadata_ref();
  ::apache::thrift::detail::md::StructMetadata<Py3Hidden>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<PyDeprecatedHidden>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<Flags>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<Name>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<Adapter>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<UseCAPI>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<Py3EnableCppAdapter>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<MigrationBlockingAllowInheritance>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<NoIntBaseClassDeprecated>::gen(metadata);
  return metadata;
}
} // namespace facebook
} // namespace thrift
} // namespace annotation
} // namespace python
