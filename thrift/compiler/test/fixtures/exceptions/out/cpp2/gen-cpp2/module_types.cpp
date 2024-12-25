/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/exceptions/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#include "thrift/compiler/test/fixtures/exceptions/gen-cpp2/module_types.h"
#include "thrift/compiler/test/fixtures/exceptions/gen-cpp2/module_types.tcc"

#include <thrift/lib/cpp2/gen/module_types_cpp.h>

#include "thrift/compiler/test/fixtures/exceptions/gen-cpp2/module_data.h"
[[maybe_unused]] static constexpr std::string_view kModuleName = "module";


namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::Fiery>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::Fiery>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view Fiery::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<Fiery>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view Fiery::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<Fiery>::name;
}

Fiery::Fiery(const Fiery&) = default;
Fiery& Fiery::operator=(const Fiery&) = default;
Fiery::Fiery() {
}

Fiery::Fiery(std::string __message) : Fiery() {
  __fbthrift_field_message = std::move(__message);
}


Fiery::~Fiery() {}

Fiery::Fiery([[maybe_unused]] Fiery&& other) noexcept :
    __fbthrift_field_message(std::move(other.__fbthrift_field_message)) {
}

Fiery& Fiery::operator=([[maybe_unused]] Fiery&& other) noexcept {
    this->__fbthrift_field_message = std::move(other.__fbthrift_field_message);
    return *this;
}


Fiery::Fiery(apache::thrift::FragileConstructor, ::std::string message__arg) :
    __fbthrift_field_message(std::move(message__arg)) { 
}


void Fiery::__fbthrift_clear() {
  // clear all fields
  this->__fbthrift_field_message = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
}

void Fiery::__fbthrift_clear_terse_fields() {
}

bool Fiery::__fbthrift_is_empty() const {
  return false;
}

bool Fiery::operator==([[maybe_unused]] const Fiery& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool Fiery::operator<([[maybe_unused]] const Fiery& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


void swap([[maybe_unused]] Fiery& a, [[maybe_unused]] Fiery& b) {
  using ::std::swap;
  swap(a.__fbthrift_field_message, b.__fbthrift_field_message);
}

template void Fiery::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t Fiery::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t Fiery::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t Fiery::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void Fiery::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t Fiery::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t Fiery::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t Fiery::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::Serious>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::Serious>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view Serious::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<Serious>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view Serious::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<Serious>::name;
}

Serious::Serious(const Serious&) = default;
Serious& Serious::operator=(const Serious&) = default;
Serious::Serious() {
}

Serious::Serious(std::string __message) : Serious() {
  __fbthrift_field_sonnet = std::move(__message);
}


Serious::~Serious() {}

Serious::Serious([[maybe_unused]] Serious&& other) noexcept :
    __fbthrift_field_sonnet(std::move(other.__fbthrift_field_sonnet)),
    __isset(other.__isset) {
}

Serious& Serious::operator=([[maybe_unused]] Serious&& other) noexcept {
    this->__fbthrift_field_sonnet = std::move(other.__fbthrift_field_sonnet);
    __isset = other.__isset;
    return *this;
}


Serious::Serious(apache::thrift::FragileConstructor, ::std::string sonnet__arg) :
    __fbthrift_field_sonnet(std::move(sonnet__arg)) { 
  __isset.set(folly::index_constant<0>(), true);
}


void Serious::__fbthrift_clear() {
  // clear all fields
  this->__fbthrift_field_sonnet = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  __isset = {};
}

void Serious::__fbthrift_clear_terse_fields() {
}

bool Serious::__fbthrift_is_empty() const {
  return !(this->__isset.get(0));
}

bool Serious::operator==([[maybe_unused]] const Serious& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool Serious::operator<([[maybe_unused]] const Serious& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


const ::std::string* Serious::get_sonnet() const& {
  return sonnet_ref().has_value() ? std::addressof(__fbthrift_field_sonnet) : nullptr;
}

::std::string* Serious::get_sonnet() & {
  return sonnet_ref().has_value() ? std::addressof(__fbthrift_field_sonnet) : nullptr;
}

void swap([[maybe_unused]] Serious& a, [[maybe_unused]] Serious& b) {
  using ::std::swap;
  swap(a.__fbthrift_field_sonnet, b.__fbthrift_field_sonnet);
  swap(a.__isset, b.__isset);
}

template void Serious::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t Serious::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t Serious::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t Serious::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void Serious::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t Serious::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t Serious::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t Serious::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::ComplexFieldNames>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::ComplexFieldNames>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view ComplexFieldNames::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<ComplexFieldNames>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view ComplexFieldNames::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<ComplexFieldNames>::name;
}

ComplexFieldNames::ComplexFieldNames(const ComplexFieldNames&) = default;
ComplexFieldNames& ComplexFieldNames::operator=(const ComplexFieldNames&) = default;
ComplexFieldNames::ComplexFieldNames() {
}

ComplexFieldNames::ComplexFieldNames(std::string __message) : ComplexFieldNames() {
  __fbthrift_field_internal_error_message = std::move(__message);
}


ComplexFieldNames::~ComplexFieldNames() {}

ComplexFieldNames::ComplexFieldNames([[maybe_unused]] ComplexFieldNames&& other) noexcept :
    __fbthrift_field_error_message(std::move(other.__fbthrift_field_error_message)),
    __fbthrift_field_internal_error_message(std::move(other.__fbthrift_field_internal_error_message)),
    __isset(other.__isset) {
}

ComplexFieldNames& ComplexFieldNames::operator=([[maybe_unused]] ComplexFieldNames&& other) noexcept {
    this->__fbthrift_field_error_message = std::move(other.__fbthrift_field_error_message);
    this->__fbthrift_field_internal_error_message = std::move(other.__fbthrift_field_internal_error_message);
    __isset = other.__isset;
    return *this;
}


ComplexFieldNames::ComplexFieldNames(apache::thrift::FragileConstructor, ::std::string error_message__arg, ::std::string internal_error_message__arg) :
    __fbthrift_field_error_message(std::move(error_message__arg)),
    __fbthrift_field_internal_error_message(std::move(internal_error_message__arg)) { 
  __isset.set(folly::index_constant<0>(), true);
  __isset.set(folly::index_constant<1>(), true);
}


void ComplexFieldNames::__fbthrift_clear() {
  // clear all fields
  this->__fbthrift_field_error_message = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  this->__fbthrift_field_internal_error_message = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  __isset = {};
}

void ComplexFieldNames::__fbthrift_clear_terse_fields() {
}

bool ComplexFieldNames::__fbthrift_is_empty() const {
  return false;
}

bool ComplexFieldNames::operator==([[maybe_unused]] const ComplexFieldNames& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool ComplexFieldNames::operator<([[maybe_unused]] const ComplexFieldNames& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


void swap([[maybe_unused]] ComplexFieldNames& a, [[maybe_unused]] ComplexFieldNames& b) {
  using ::std::swap;
  swap(a.__fbthrift_field_error_message, b.__fbthrift_field_error_message);
  swap(a.__fbthrift_field_internal_error_message, b.__fbthrift_field_internal_error_message);
  swap(a.__isset, b.__isset);
}

template void ComplexFieldNames::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t ComplexFieldNames::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t ComplexFieldNames::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t ComplexFieldNames::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void ComplexFieldNames::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t ComplexFieldNames::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t ComplexFieldNames::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t ComplexFieldNames::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::CustomFieldNames>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::CustomFieldNames>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view CustomFieldNames::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<CustomFieldNames>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view CustomFieldNames::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<CustomFieldNames>::name;
}

CustomFieldNames::CustomFieldNames(const CustomFieldNames&) = default;
CustomFieldNames& CustomFieldNames::operator=(const CustomFieldNames&) = default;
CustomFieldNames::CustomFieldNames() {
}

CustomFieldNames::CustomFieldNames(std::string __message) : CustomFieldNames() {
  __fbthrift_field_internal_error_message = std::move(__message);
}


CustomFieldNames::~CustomFieldNames() {}

CustomFieldNames::CustomFieldNames([[maybe_unused]] CustomFieldNames&& other) noexcept :
    __fbthrift_field_error_message(std::move(other.__fbthrift_field_error_message)),
    __fbthrift_field_internal_error_message(std::move(other.__fbthrift_field_internal_error_message)),
    __isset(other.__isset) {
}

CustomFieldNames& CustomFieldNames::operator=([[maybe_unused]] CustomFieldNames&& other) noexcept {
    this->__fbthrift_field_error_message = std::move(other.__fbthrift_field_error_message);
    this->__fbthrift_field_internal_error_message = std::move(other.__fbthrift_field_internal_error_message);
    __isset = other.__isset;
    return *this;
}


CustomFieldNames::CustomFieldNames(apache::thrift::FragileConstructor, ::std::string error_message__arg, ::std::string internal_error_message__arg) :
    __fbthrift_field_error_message(std::move(error_message__arg)),
    __fbthrift_field_internal_error_message(std::move(internal_error_message__arg)) { 
  __isset.set(folly::index_constant<0>(), true);
  __isset.set(folly::index_constant<1>(), true);
}


void CustomFieldNames::__fbthrift_clear() {
  // clear all fields
  this->__fbthrift_field_error_message = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  this->__fbthrift_field_internal_error_message = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  __isset = {};
}

void CustomFieldNames::__fbthrift_clear_terse_fields() {
}

bool CustomFieldNames::__fbthrift_is_empty() const {
  return false;
}

bool CustomFieldNames::operator==([[maybe_unused]] const CustomFieldNames& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool CustomFieldNames::operator<([[maybe_unused]] const CustomFieldNames& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


void swap([[maybe_unused]] CustomFieldNames& a, [[maybe_unused]] CustomFieldNames& b) {
  using ::std::swap;
  swap(a.__fbthrift_field_error_message, b.__fbthrift_field_error_message);
  swap(a.__fbthrift_field_internal_error_message, b.__fbthrift_field_internal_error_message);
  swap(a.__isset, b.__isset);
}

template void CustomFieldNames::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t CustomFieldNames::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t CustomFieldNames::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t CustomFieldNames::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void CustomFieldNames::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t CustomFieldNames::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t CustomFieldNames::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t CustomFieldNames::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::ExceptionWithPrimitiveField>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::ExceptionWithPrimitiveField>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view ExceptionWithPrimitiveField::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<ExceptionWithPrimitiveField>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view ExceptionWithPrimitiveField::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<ExceptionWithPrimitiveField>::name;
}

ExceptionWithPrimitiveField::ExceptionWithPrimitiveField(const ExceptionWithPrimitiveField&) = default;
ExceptionWithPrimitiveField& ExceptionWithPrimitiveField::operator=(const ExceptionWithPrimitiveField&) = default;
ExceptionWithPrimitiveField::ExceptionWithPrimitiveField() :
    __fbthrift_field_error_code() {
}

ExceptionWithPrimitiveField::ExceptionWithPrimitiveField(std::string __message) : ExceptionWithPrimitiveField() {
  __fbthrift_field_message = std::move(__message);
}


ExceptionWithPrimitiveField::~ExceptionWithPrimitiveField() {}

ExceptionWithPrimitiveField::ExceptionWithPrimitiveField([[maybe_unused]] ExceptionWithPrimitiveField&& other) noexcept :
    __fbthrift_field_message(std::move(other.__fbthrift_field_message)),
    __fbthrift_field_error_code(std::move(other.__fbthrift_field_error_code)),
    __isset(other.__isset) {
}

ExceptionWithPrimitiveField& ExceptionWithPrimitiveField::operator=([[maybe_unused]] ExceptionWithPrimitiveField&& other) noexcept {
    this->__fbthrift_field_message = std::move(other.__fbthrift_field_message);
    this->__fbthrift_field_error_code = std::move(other.__fbthrift_field_error_code);
    __isset = other.__isset;
    return *this;
}


ExceptionWithPrimitiveField::ExceptionWithPrimitiveField(apache::thrift::FragileConstructor, ::std::string message__arg, ::std::int32_t error_code__arg) :
    __fbthrift_field_message(std::move(message__arg)),
    __fbthrift_field_error_code(std::move(error_code__arg)) { 
  __isset.set(folly::index_constant<0>(), true);
  __isset.set(folly::index_constant<1>(), true);
}


void ExceptionWithPrimitiveField::__fbthrift_clear() {
  // clear all fields
  this->__fbthrift_field_message = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  this->__fbthrift_field_error_code = ::std::int32_t();
  __isset = {};
}

void ExceptionWithPrimitiveField::__fbthrift_clear_terse_fields() {
}

bool ExceptionWithPrimitiveField::__fbthrift_is_empty() const {
  return false;
}

bool ExceptionWithPrimitiveField::operator==([[maybe_unused]] const ExceptionWithPrimitiveField& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool ExceptionWithPrimitiveField::operator<([[maybe_unused]] const ExceptionWithPrimitiveField& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


::std::int32_t ExceptionWithPrimitiveField::get_error_code() const {
  return __fbthrift_field_error_code;
}

::std::int32_t& ExceptionWithPrimitiveField::set_error_code(::std::int32_t error_code_) {
  error_code_ref() = error_code_;
  return __fbthrift_field_error_code;
}

void swap([[maybe_unused]] ExceptionWithPrimitiveField& a, [[maybe_unused]] ExceptionWithPrimitiveField& b) {
  using ::std::swap;
  swap(a.__fbthrift_field_message, b.__fbthrift_field_message);
  swap(a.__fbthrift_field_error_code, b.__fbthrift_field_error_code);
  swap(a.__isset, b.__isset);
}

template void ExceptionWithPrimitiveField::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t ExceptionWithPrimitiveField::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t ExceptionWithPrimitiveField::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t ExceptionWithPrimitiveField::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void ExceptionWithPrimitiveField::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t ExceptionWithPrimitiveField::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t ExceptionWithPrimitiveField::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t ExceptionWithPrimitiveField::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::ExceptionWithStructuredAnnotation>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::ExceptionWithStructuredAnnotation>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view ExceptionWithStructuredAnnotation::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<ExceptionWithStructuredAnnotation>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view ExceptionWithStructuredAnnotation::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<ExceptionWithStructuredAnnotation>::name;
}

ExceptionWithStructuredAnnotation::ExceptionWithStructuredAnnotation(const ExceptionWithStructuredAnnotation&) = default;
ExceptionWithStructuredAnnotation& ExceptionWithStructuredAnnotation::operator=(const ExceptionWithStructuredAnnotation&) = default;
ExceptionWithStructuredAnnotation::ExceptionWithStructuredAnnotation() :
    __fbthrift_field_error_code() {
}

ExceptionWithStructuredAnnotation::ExceptionWithStructuredAnnotation(std::string __message) : ExceptionWithStructuredAnnotation() {
  __fbthrift_field_message_field = std::move(__message);
}


ExceptionWithStructuredAnnotation::~ExceptionWithStructuredAnnotation() {}

ExceptionWithStructuredAnnotation::ExceptionWithStructuredAnnotation([[maybe_unused]] ExceptionWithStructuredAnnotation&& other) noexcept :
    __fbthrift_field_message_field(std::move(other.__fbthrift_field_message_field)),
    __fbthrift_field_error_code(std::move(other.__fbthrift_field_error_code)),
    __isset(other.__isset) {
}

ExceptionWithStructuredAnnotation& ExceptionWithStructuredAnnotation::operator=([[maybe_unused]] ExceptionWithStructuredAnnotation&& other) noexcept {
    this->__fbthrift_field_message_field = std::move(other.__fbthrift_field_message_field);
    this->__fbthrift_field_error_code = std::move(other.__fbthrift_field_error_code);
    __isset = other.__isset;
    return *this;
}


ExceptionWithStructuredAnnotation::ExceptionWithStructuredAnnotation(apache::thrift::FragileConstructor, ::std::string message_field__arg, ::std::int32_t error_code__arg) :
    __fbthrift_field_message_field(std::move(message_field__arg)),
    __fbthrift_field_error_code(std::move(error_code__arg)) { 
  __isset.set(folly::index_constant<0>(), true);
  __isset.set(folly::index_constant<1>(), true);
}


void ExceptionWithStructuredAnnotation::__fbthrift_clear() {
  // clear all fields
  this->__fbthrift_field_message_field = apache::thrift::StringTraits<std::string>::fromStringLiteral("");
  this->__fbthrift_field_error_code = ::std::int32_t();
  __isset = {};
}

void ExceptionWithStructuredAnnotation::__fbthrift_clear_terse_fields() {
}

bool ExceptionWithStructuredAnnotation::__fbthrift_is_empty() const {
  return false;
}

bool ExceptionWithStructuredAnnotation::operator==([[maybe_unused]] const ExceptionWithStructuredAnnotation& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool ExceptionWithStructuredAnnotation::operator<([[maybe_unused]] const ExceptionWithStructuredAnnotation& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


::std::int32_t ExceptionWithStructuredAnnotation::get_error_code() const {
  return __fbthrift_field_error_code;
}

::std::int32_t& ExceptionWithStructuredAnnotation::set_error_code(::std::int32_t error_code_) {
  error_code_ref() = error_code_;
  return __fbthrift_field_error_code;
}

void swap([[maybe_unused]] ExceptionWithStructuredAnnotation& a, [[maybe_unused]] ExceptionWithStructuredAnnotation& b) {
  using ::std::swap;
  swap(a.__fbthrift_field_message_field, b.__fbthrift_field_message_field);
  swap(a.__fbthrift_field_error_code, b.__fbthrift_field_error_code);
  swap(a.__isset, b.__isset);
}

template void ExceptionWithStructuredAnnotation::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t ExceptionWithStructuredAnnotation::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t ExceptionWithStructuredAnnotation::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t ExceptionWithStructuredAnnotation::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void ExceptionWithStructuredAnnotation::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t ExceptionWithStructuredAnnotation::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t ExceptionWithStructuredAnnotation::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t ExceptionWithStructuredAnnotation::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace apache {
namespace thrift {
namespace detail {

void TccStructTraits<::cpp2::Banal>::translateFieldName(
    std::string_view _fname,
    int16_t& fid,
    apache::thrift::protocol::TType& _ftype) noexcept {
  using data = apache::thrift::TStructDataStorage<::cpp2::Banal>;
  static const st::translate_field_name_table table{
      data::fields_size,
      data::fields_names.data(),
      data::fields_ids.data(),
      data::fields_types.data()};
  st::translate_field_name(_fname, fid, _ftype, table);
}

} // namespace detail
} // namespace thrift
} // namespace apache

namespace cpp2 {

std::string_view Banal::__fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord) {
  if (ord == ::apache::thrift::FieldOrdinal{0}) { return {}; }
  return apache::thrift::TStructDataStorage<Banal>::fields_names[folly::to_underlying(ord) - 1];
}
std::string_view Banal::__fbthrift_get_class_name() {
  return apache::thrift::TStructDataStorage<Banal>::name;
}

Banal::Banal(const Banal&) = default;
Banal& Banal::operator=(const Banal&) = default;
Banal::Banal() {
}


Banal::~Banal() {}

Banal::Banal([[maybe_unused]] Banal&& other) noexcept{}
Banal& Banal::operator=([[maybe_unused]] Banal&& other) noexcept {
    return *this;
}


Banal::Banal(apache::thrift::FragileConstructor) {}


void Banal::__fbthrift_clear() {
  // clear all fields
}

void Banal::__fbthrift_clear_terse_fields() {
}

bool Banal::__fbthrift_is_empty() const {
  return true;
}

bool Banal::operator==([[maybe_unused]] const Banal& rhs) const {
  return ::apache::thrift::op::detail::StructEquality{}(*this, rhs);
}

bool Banal::operator<([[maybe_unused]] const Banal& rhs) const {
  return ::apache::thrift::op::detail::StructLessThan{}(*this, rhs);
}


void swap([[maybe_unused]] Banal& a, [[maybe_unused]] Banal& b) {
  using ::std::swap;
}

template void Banal::readNoXfer<>(apache::thrift::BinaryProtocolReader*);
template uint32_t Banal::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t Banal::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t Banal::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template void Banal::readNoXfer<>(apache::thrift::CompactProtocolReader*);
template uint32_t Banal::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t Banal::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t Banal::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;


} // namespace cpp2

namespace cpp2 { namespace {
[[maybe_unused]] FOLLY_ERASE void validateAdapters() {
}
}} // namespace cpp2
namespace apache::thrift::detail::annotation {
}
