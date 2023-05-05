/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/service-schema/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */

#include "thrift/compiler/test/fixtures/service-schema/gen-cpp2/module_constants.h"

#include <thrift/lib/cpp2/gen/module_constants_cpp.h>


namespace cpp2 {

::apache::thrift::type::Schema const& module_constants::schemaPrimitivesService() {
  static folly::Indestructible<::apache::thrift::type::Schema> const instance{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Schema>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::definitions>(std::initializer_list<::apache::thrift::type::Definition>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Definition>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::enumDef>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Enum>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("Result")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::values>(std::initializer_list<::apache::thrift::type::EnumValue>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::EnumValue>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("OK")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::value>(static_cast<::std::int32_t>(0))),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::EnumValue>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("SO_SO")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::value>(static_cast<::std::int32_t>(1))),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::EnumValue>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("GOOD")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::value>(static_cast<::std::int32_t>(2)))})))),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Definition>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::exceptionDef>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Exception>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("CustomException")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::fields>(std::initializer_list<::apache::thrift::type::Field>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Field>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("name")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::id>(static_cast<::apache::thrift::type::FieldId>(1)), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FieldQualifier::Fill), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::type, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{3}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::stringType>( ::apache::thrift::type::Void::Unused))))))}), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::safety>( ::apache::thrift::type::ErrorSafety::Unspecified), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::kind>( ::apache::thrift::type::ErrorKind::Unspecified), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::blame>( ::apache::thrift::type::ErrorBlame::Unspecified)))),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Definition>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::serviceDef>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Service>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("PrimitivesService")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::annotations>(std::initializer_list<::apache::thrift::type::ValueId>{static_cast<::apache::thrift::type::ValueId>(1)}))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::functions>(std::initializer_list<::apache::thrift::type::Function>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Function>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("init")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FunctionQualifier::Unspecified), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::returnTypes>(std::initializer_list<::apache::thrift::type::ReturnType>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::ReturnType>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::thriftType, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{1}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::i64Type>( ::apache::thrift::type::Void::Unused))))))}), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::paramlist>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Paramlist>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::fields>(std::initializer_list<::apache::thrift::type::Field>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Field>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("param0")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::id>(static_cast<::apache::thrift::type::FieldId>(1)), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FieldQualifier::Fill), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::type, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{3}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::i64Type>( ::apache::thrift::type::Void::Unused)))))),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Field>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("param1")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::id>(static_cast<::apache::thrift::type::FieldId>(2)), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FieldQualifier::Fill), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::type, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{3}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::i64Type>( ::apache::thrift::type::Void::Unused))))))})))),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Function>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("method_that_throws")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FunctionQualifier::Unspecified), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::returnTypes>(std::initializer_list<::apache::thrift::type::ReturnType>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::ReturnType>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::thriftType, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{1}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::enumType>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeUri>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))))))))}), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::paramlist>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Paramlist>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::fields>(std::initializer_list<::apache::thrift::type::Field>{}))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::exceptions>(std::initializer_list<::apache::thrift::type::Field>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Field>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("e")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::id>(static_cast<::apache::thrift::type::FieldId>(1)), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FieldQualifier::Fill), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::type, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{3}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::exceptionType>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeUri>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))))))))})),
  ::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Function>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("return_void_method")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FunctionQualifier::Unspecified), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::returnTypes>(std::initializer_list<::apache::thrift::type::ReturnType>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::ReturnType>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::thriftType, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{1}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>()))))}), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::paramlist>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Paramlist>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::fields>(std::initializer_list<::apache::thrift::type::Field>{::apache::thrift::detail::make_structured_constant<::apache::thrift::type::Field>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::attrs>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::DefinitionAttrs>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(apache::thrift::StringTraits<std::string>::fromStringLiteral("id")), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::uri>(apache::thrift::StringTraits<std::string>::fromStringLiteral("")))), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::id>(static_cast<::apache::thrift::type::FieldId>(1)), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::qualifier>( ::apache::thrift::type::FieldQualifier::Fill), ::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::type, ::apache::thrift::InlineAdapter<::apache::thrift::type::Type>, ::apache::thrift::FieldId{3}>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeStruct>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::name>(::apache::thrift::detail::make_structured_constant<::apache::thrift::type::TypeName>(::apache::thrift::detail::wrap_struct_argument<::apache::thrift::ident::i64Type>( ::apache::thrift::type::Void::Unused))))))}))))}))))}))};
  return *instance;
}

} // cpp2
