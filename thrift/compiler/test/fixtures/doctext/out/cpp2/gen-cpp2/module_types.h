/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/doctext/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include <thrift/lib/cpp2/gen/module_types_h.h>



namespace apache {
namespace thrift {
namespace ident {
struct useless_field;
struct i;
struct s;
struct message;
} // namespace ident
namespace detail {
#ifndef APACHE_THRIFT_ACCESSOR_useless_field
#define APACHE_THRIFT_ACCESSOR_useless_field
APACHE_THRIFT_DEFINE_ACCESSOR(useless_field);
#endif
#ifndef APACHE_THRIFT_ACCESSOR_i
#define APACHE_THRIFT_ACCESSOR_i
APACHE_THRIFT_DEFINE_ACCESSOR(i);
#endif
#ifndef APACHE_THRIFT_ACCESSOR_s
#define APACHE_THRIFT_ACCESSOR_s
APACHE_THRIFT_DEFINE_ACCESSOR(s);
#endif
#ifndef APACHE_THRIFT_ACCESSOR_message
#define APACHE_THRIFT_ACCESSOR_message
APACHE_THRIFT_DEFINE_ACCESSOR(message);
#endif
} // namespace detail
} // namespace thrift
} // namespace apache

// BEGIN declare_enums
namespace cpp2 {

/** Glean {"file": "thrift/compiler/test/fixtures/doctext/src/module.thrift", "name": "B", "kind": "enum" } */
enum class B {
  HELLO = 0,
};



} // namespace cpp2

namespace std {
template<> struct hash<::cpp2::B> :
  ::apache::thrift::detail::enum_hash<::cpp2::B> {};
} // std

namespace apache { namespace thrift {


template <> struct TEnumDataStorage<::cpp2::B>;

template <> struct TEnumTraits<::cpp2::B> {
  using type = ::cpp2::B;

  static constexpr std::size_t const size = 1;
  static folly::Range<type const*> const values;
  static folly::Range<std::string_view const*> const names;
  static const std::string_view __fbthrift_module_name_internal_do_not_use;

  static bool findName(type value, std::string_view* out) noexcept;
  static bool findValue(std::string_view name, type* out) noexcept;

  template <class ...>
  FOLLY_ERASE static std::string_view typeName() noexcept {
    return "B";
  }

  template <class ...>
  FOLLY_ERASE static constexpr std::string_view moduleName() noexcept {
    return "module";
  }

  static char const* findName(type value) noexcept {
    std::string_view ret;
    (void)findName(value, &ret);
    return ret.data();
  }
  static constexpr type min() { return type::HELLO; }
  static constexpr type max() { return type::HELLO; }
};


}} // apache::thrift


// END declare_enums
// BEGIN forward_declare
namespace cpp2 {
class A;
class U;
class Bang;
} // namespace cpp2
// END forward_declare
namespace apache::thrift::detail::annotation {
} // namespace apache::thrift::detail::annotation

namespace apache::thrift::detail::qualifier {
} // namespace apache::thrift::detail::qualifier

// BEGIN hash_and_equal_to
// END hash_and_equal_to
namespace cpp2 {
using ::apache::thrift::detail::operator!=;
using ::apache::thrift::detail::operator>;
using ::apache::thrift::detail::operator<=;
using ::apache::thrift::detail::operator>=;

/** Glean {"file": "thrift/compiler/test/fixtures/doctext/src/module.thrift", "name": "lanyard", "kind": "typedef" } */
typedef ::std::string lanyard;
/** Glean {"file": "thrift/compiler/test/fixtures/doctext/src/module.thrift", "name": "number", "kind": "typedef" } */
typedef ::std::int32_t number;

/** Glean {"file": "thrift/compiler/test/fixtures/doctext/src/module.thrift", "name": "A", "kind": "struct" } */
class A final  {
 private:
  friend struct ::apache::thrift::detail::st::struct_private_access;
  template<class> friend struct ::apache::thrift::detail::invoke_reffer;

  //  used by a static_assert in the corresponding source
  static constexpr bool __fbthrift_cpp2_gen_json = false;
  static constexpr bool __fbthrift_cpp2_is_runtime_annotation = false;
  static std::string_view __fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord);
  static std::string_view __fbthrift_get_class_name();
  template <class ...>
  FOLLY_ERASE static constexpr std::string_view __fbthrift_get_module_name() noexcept {
    return "module";
  }
  using __fbthrift_reflection_ident_list = folly::tag_t<
    ::apache::thrift::ident::useless_field
  >;

  static constexpr std::int16_t __fbthrift_reflection_field_id_list[] = {0,1};
  using __fbthrift_reflection_type_tags = folly::tag_t<
    ::apache::thrift::type::i32_t
  >;

  static constexpr std::size_t __fbthrift_field_size_v = 1;

  template<class T>
  using __fbthrift_id = ::apache::thrift::type::field_id<__fbthrift_reflection_field_id_list[folly::to_underlying(T::value)]>;

  template<class T>
  using __fbthrift_type_tag = ::apache::thrift::detail::at<__fbthrift_reflection_type_tags, T::value>;

  template<class T>
  using __fbthrift_ident = ::apache::thrift::detail::at<__fbthrift_reflection_ident_list, T::value>;

  template<class T> using __fbthrift_ordinal = ::apache::thrift::type::ordinal_tag<
    ::apache::thrift::detail::getFieldOrdinal<T,
                                              __fbthrift_reflection_ident_list,
                                              __fbthrift_reflection_type_tags>(
      __fbthrift_reflection_field_id_list
    )
  >;
  void __fbthrift_clear();
  void __fbthrift_clear_terse_fields();
  bool __fbthrift_is_empty() const;

 public:
  using __fbthrift_cpp2_type = A;
  static constexpr bool __fbthrift_cpp2_is_union =
    false;
  static constexpr bool __fbthrift_cpp2_uses_op_encode =
    false;


 public:

  A() :
      __fbthrift_field_useless_field() {
  }
  // FragileConstructor for use in initialization lists only.
  [[deprecated("This constructor is deprecated")]]
  A(apache::thrift::FragileConstructor, ::std::int32_t useless_field__arg);

  A(A&&) = default;

  A(const A&) = default;


  A& operator=(A&&) = default;

  A& operator=(const A&) = default;
 private:
  ::std::int32_t __fbthrift_field_useless_field;
 private:
  apache::thrift::detail::isset_bitset<1, apache::thrift::detail::IssetBitsetOption::Unpacked> __isset;

 public:

  bool operator==(const A&) const;
  bool operator<(const A&) const;

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&> useless_field_ref() const& {
    return {this->__fbthrift_field_useless_field, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&&> useless_field_ref() const&& {
    return {static_cast<const T&&>(this->__fbthrift_field_useless_field), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&> useless_field_ref() & {
    return {this->__fbthrift_field_useless_field, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&&> useless_field_ref() && {
    return {static_cast<T&&>(this->__fbthrift_field_useless_field), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&> useless_field() const& {
    return {this->__fbthrift_field_useless_field, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&&> useless_field() const&& {
    return {static_cast<const T&&>(this->__fbthrift_field_useless_field), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&> useless_field() & {
    return {this->__fbthrift_field_useless_field, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&&> useless_field() && {
    return {static_cast<T&&>(this->__fbthrift_field_useless_field), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "useless_field" } */
  [[deprecated("Use `FOO.useless_field().value();` instead of `FOO.get_useless_field();`")]]
  ::std::int32_t get_useless_field() const {
    return __fbthrift_field_useless_field;
  }

  /** Glean { "field": "useless_field" } */
  [[deprecated("Use `FOO.useless_field() = BAR;` instead of `FOO.set_useless_field(BAR);`")]]
  ::std::int32_t& set_useless_field(::std::int32_t useless_field_) {
    useless_field_ref() = useless_field_;
    return __fbthrift_field_useless_field;
  }

  template <class Protocol_>
  unsigned long read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t serializedSize(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t serializedSizeZC(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t write(Protocol_* prot_) const;

 private:
  template <class Protocol_>
  void readNoXfer(Protocol_* iprot);

  friend class ::apache::thrift::Cpp2Ops<A>;
  friend void swap(A& a, A& b);
};

template <class Protocol_>
unsigned long A::read(Protocol_* iprot) {
  auto _xferStart = iprot->getCursorPosition();
  readNoXfer(iprot);
  return iprot->getCursorPosition() - _xferStart;
}


/** Glean {"file": "thrift/compiler/test/fixtures/doctext/src/module.thrift", "name": "U", "kind": "union" } */
class U final  {
 private:
  friend struct ::apache::thrift::detail::st::struct_private_access;
  template<class> friend struct ::apache::thrift::detail::invoke_reffer;

  //  used by a static_assert in the corresponding source
  static constexpr bool __fbthrift_cpp2_gen_json = false;
  static constexpr bool __fbthrift_cpp2_is_runtime_annotation = false;
  static std::string_view __fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord);
  static std::string_view __fbthrift_get_class_name();
  template <class ...>
  FOLLY_ERASE static constexpr std::string_view __fbthrift_get_module_name() noexcept {
    return "module";
  }
  using __fbthrift_reflection_ident_list = folly::tag_t<
    ::apache::thrift::ident::i,
    ::apache::thrift::ident::s
  >;

  static constexpr std::int16_t __fbthrift_reflection_field_id_list[] = {0,1,2};
  using __fbthrift_reflection_type_tags = folly::tag_t<
    ::apache::thrift::type::i32_t,
    ::apache::thrift::type::string_t
  >;

  static constexpr std::size_t __fbthrift_field_size_v = 2;

  template<class T>
  using __fbthrift_id = ::apache::thrift::type::field_id<__fbthrift_reflection_field_id_list[folly::to_underlying(T::value)]>;

  template<class T>
  using __fbthrift_type_tag = ::apache::thrift::detail::at<__fbthrift_reflection_type_tags, T::value>;

  template<class T>
  using __fbthrift_ident = ::apache::thrift::detail::at<__fbthrift_reflection_ident_list, T::value>;

  template<class T> using __fbthrift_ordinal = ::apache::thrift::type::ordinal_tag<
    ::apache::thrift::detail::getFieldOrdinal<T,
                                              __fbthrift_reflection_ident_list,
                                              __fbthrift_reflection_type_tags>(
      __fbthrift_reflection_field_id_list
    )
  >;
  void __fbthrift_clear();
  void __fbthrift_destruct();
  bool __fbthrift_is_empty() const;

 public:
  using __fbthrift_cpp2_type = U;
  static constexpr bool __fbthrift_cpp2_is_union =
    true;
  static constexpr bool __fbthrift_cpp2_uses_op_encode =
    false;


 public:
  enum Type : int {
    __EMPTY__ = 0,
    i = 1,
    s = 2,
  } ;

  U()
      : type_(folly::to_underlying(Type::__EMPTY__)) {}

  U(U&& rhs) noexcept
      : type_(folly::to_underlying(Type::__EMPTY__)) {
    if (this == &rhs) { return; }
    switch (rhs.getType()) {
      case Type::__EMPTY__:
      {
        return;
      }
      case Type::i:
      {
        set_i(std::move(rhs.value_.i));
        break;
      }
      case Type::s:
      {
        set_s(std::move(rhs.value_.s));
        break;
      }
      default:
      {
        assert(false);
        break;
      }
    }
    apache::thrift::clear(rhs);
  }

  U(const U& rhs);

  U& operator=(U&& rhs) noexcept {
    if (this == &rhs) { return *this; }
    switch (rhs.getType()) {
      case Type::__EMPTY__:
      {
        __fbthrift_clear();
        return *this;
      }
      case Type::i:
      {
        set_i(std::move(rhs.value_.i));
        break;
      }
      case Type::s:
      {
        set_s(std::move(rhs.value_.s));
        break;
      }
      default:
      {
        assert(false);
        __fbthrift_clear();
      }
    }
    apache::thrift::clear(rhs);
    return *this;
  }

  U& operator=(const U& rhs);

  ~U();

  union storage_type {
    ::std::int32_t i;
    ::std::string s;

    storage_type() {}
    ~storage_type() {}
  } ;

  bool operator==(const U&) const;
  bool operator<(const U&) const;

  /** Glean { "field": "i" } */
  template <typename... A, std::enable_if_t<!sizeof...(A), int> = 0>
  ::std::int32_t& set_i(::std::int32_t t = ::std::int32_t()) {
    using T0 = ::std::int32_t;
    using T = folly::type_t<T0, A...>;
    __fbthrift_clear();
    type_ = folly::to_underlying(Type::i);
    ::new (std::addressof(value_.i)) T(t);
    return value_.i;
  }


  /** Glean { "field": "s" } */
  template <typename... A, std::enable_if_t<!sizeof...(A), int> = 0>
  ::std::string& set_s(::std::string const &t) {
    using T0 = ::std::string;
    using T = folly::type_t<T0, A...>;
    __fbthrift_clear();
    type_ = folly::to_underlying(Type::s);
    ::new (std::addressof(value_.s)) T(t);
    return value_.s;
  }

  /** Glean { "field": "s" } */
  template <typename... A, std::enable_if_t<!sizeof...(A), int> = 0>
  ::std::string& set_s(::std::string&& t) {
    using T0 = ::std::string;
    using T = folly::type_t<T0, A...>;
    __fbthrift_clear();
    type_ = folly::to_underlying(Type::s);
    ::new (std::addressof(value_.s)) T(std::move(t));
    return value_.s;
  }

  /** Glean { "field": "s" } */
  template<typename... T, typename = ::apache::thrift::safe_overload_t<::std::string, T...>> ::std::string& set_s(T&&... t) {
    __fbthrift_clear();
    type_ = folly::to_underlying(Type::s);
    ::new (std::addressof(value_.s)) ::std::string(std::forward<T>(t)...);
    return value_.s;
  }


  ::std::int32_t const& get_i() const {
    if (getType() != Type::i) {
      ::apache::thrift::detail::throw_on_bad_union_field_access();
    }
    return value_.i;
  }

  ::std::string const& get_s() const {
    if (getType() != Type::s) {
      ::apache::thrift::detail::throw_on_bad_union_field_access();
    }
    return value_.s;
  }

  ::std::int32_t& mutable_i() {
    assert(getType() == Type::i);
    return value_.i;
  }

  ::std::string& mutable_s() {
    assert(getType() == Type::s);
    return value_.s;
  }

  template <typename..., typename T = ::std::int32_t>
  T move_i() {
    assert(getType() == Type::i);
    return std::move(value_.i);
  }

  template <typename..., typename T = ::std::string>
  T move_s() {
    assert(getType() == Type::s);
    return std::move(value_.s);
  }

  /** Glean { "field": "i" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::union_field_ref<const T&> i_ref() const& {
    return {value_.i, type_, folly::to_underlying(Type::i), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }

  /** Glean { "field": "i" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::union_field_ref<const T&&> i_ref() const&& {
    return {std::move(value_.i), type_, folly::to_underlying(Type::i), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }

  /** Glean { "field": "i" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::union_field_ref<T&> i_ref() & {
    return {value_.i, type_, folly::to_underlying(Type::i), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }

  /** Glean { "field": "i" } */
  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::union_field_ref<T&&> i_ref() && {
    return {std::move(value_.i), type_, folly::to_underlying(Type::i), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }
  /** Glean { "field": "s" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::union_field_ref<const T&> s_ref() const& {
    return {value_.s, type_, folly::to_underlying(Type::s), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }

  /** Glean { "field": "s" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::union_field_ref<const T&&> s_ref() const&& {
    return {std::move(value_.s), type_, folly::to_underlying(Type::s), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }

  /** Glean { "field": "s" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::union_field_ref<T&> s_ref() & {
    return {value_.s, type_, folly::to_underlying(Type::s), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }

  /** Glean { "field": "s" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::union_field_ref<T&&> s_ref() && {
    return {std::move(value_.s), type_, folly::to_underlying(Type::s), this, ::apache::thrift::detail::union_field_ref_owner_vtable_for<decltype(*this)>};
  }
  Type getType() const { return static_cast<Type>(type_); }

  template <class Protocol_>
  unsigned long read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t serializedSize(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t serializedSizeZC(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t write(Protocol_* prot_) const;
 protected:
  storage_type value_;
  std::underlying_type_t<Type> type_;

 private:
  template <class Protocol_>
  void readNoXfer(Protocol_* iprot);

  friend class ::apache::thrift::Cpp2Ops<U>;
  friend void swap(U& a, U& b);
};

template <class Protocol_>
unsigned long U::read(Protocol_* iprot) {
  auto _xferStart = iprot->getCursorPosition();
  readNoXfer(iprot);
  return iprot->getCursorPosition() - _xferStart;
}


/** Glean {"file": "thrift/compiler/test/fixtures/doctext/src/module.thrift", "name": "Bang", "kind": "exception" } */
class FOLLY_EXPORT Bang : public virtual apache::thrift::TException {
 private:
  friend struct ::apache::thrift::detail::st::struct_private_access;
  template<class> friend struct ::apache::thrift::detail::invoke_reffer;

  //  used by a static_assert in the corresponding source
  static constexpr bool __fbthrift_cpp2_gen_json = false;
  static constexpr bool __fbthrift_cpp2_is_runtime_annotation = false;
  static std::string_view __fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord);
  static std::string_view __fbthrift_get_class_name();
  template <class ...>
  FOLLY_ERASE static constexpr std::string_view __fbthrift_get_module_name() noexcept {
    return "module";
  }
  using __fbthrift_reflection_ident_list = folly::tag_t<
    ::apache::thrift::ident::message
  >;

  static constexpr std::int16_t __fbthrift_reflection_field_id_list[] = {0,1};
  using __fbthrift_reflection_type_tags = folly::tag_t<
    ::apache::thrift::type::string_t
  >;

  static constexpr std::size_t __fbthrift_field_size_v = 1;

  template<class T>
  using __fbthrift_id = ::apache::thrift::type::field_id<__fbthrift_reflection_field_id_list[folly::to_underlying(T::value)]>;

  template<class T>
  using __fbthrift_type_tag = ::apache::thrift::detail::at<__fbthrift_reflection_type_tags, T::value>;

  template<class T>
  using __fbthrift_ident = ::apache::thrift::detail::at<__fbthrift_reflection_ident_list, T::value>;

  template<class T> using __fbthrift_ordinal = ::apache::thrift::type::ordinal_tag<
    ::apache::thrift::detail::getFieldOrdinal<T,
                                              __fbthrift_reflection_ident_list,
                                              __fbthrift_reflection_type_tags>(
      __fbthrift_reflection_field_id_list
    )
  >;
  void __fbthrift_clear();
  void __fbthrift_clear_terse_fields();
  bool __fbthrift_is_empty() const;
  static constexpr ::apache::thrift::ExceptionKind __fbthrift_cpp2_gen_exception_kind =
         ::apache::thrift::ExceptionKind::UNSPECIFIED;
  static constexpr ::apache::thrift::ExceptionSafety __fbthrift_cpp2_gen_exception_safety =
         ::apache::thrift::ExceptionSafety::UNSPECIFIED;
  static constexpr ::apache::thrift::ExceptionBlame __fbthrift_cpp2_gen_exception_blame =
         ::apache::thrift::ExceptionBlame::UNSPECIFIED;

 public:
  using __fbthrift_cpp2_type = Bang;
  static constexpr bool __fbthrift_cpp2_is_union =
    false;
  static constexpr bool __fbthrift_cpp2_uses_op_encode =
    false;


 public:

  Bang();

  // FragileConstructor for use in initialization lists only.
  [[deprecated("This constructor is deprecated")]]
  Bang(apache::thrift::FragileConstructor, ::std::string message__arg);

  Bang(Bang&&) noexcept;

  Bang(const Bang& src);


  Bang& operator=(Bang&&) noexcept;
  Bang& operator=(const Bang& src);

  ~Bang() override;

 private:
  ::std::string __fbthrift_field_message;
 private:
  apache::thrift::detail::isset_bitset<1, apache::thrift::detail::IssetBitsetOption::Unpacked> __isset;

 public:

  bool operator==(const Bang&) const;
  bool operator<(const Bang&) const;

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&> message_ref() const& {
    return {this->__fbthrift_field_message, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&&> message_ref() const&& {
    return {static_cast<const T&&>(this->__fbthrift_field_message), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<T&> message_ref() & {
    return {this->__fbthrift_field_message, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<T&&> message_ref() && {
    return {static_cast<T&&>(this->__fbthrift_field_message), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&> message() const& {
    return {this->__fbthrift_field_message, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&&> message() const&& {
    return {static_cast<const T&&>(this->__fbthrift_field_message), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<T&> message() & {
    return {this->__fbthrift_field_message, __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  template <typename..., typename T = ::std::string>
  FOLLY_ERASE ::apache::thrift::field_ref<T&&> message() && {
    return {static_cast<T&&>(this->__fbthrift_field_message), __isset.at(0), __isset.bit(0)};
  }

  /** Glean { "field": "message" } */
  [[deprecated("Use `FOO.message().value();` instead of `FOO.get_message();`")]]
  const ::std::string& get_message() const& {
    return __fbthrift_field_message;
  }

  /** Glean { "field": "message" } */
  [[deprecated("Use `FOO.message().value();` instead of `FOO.get_message();`")]]
  ::std::string get_message() && {
    return std::move(__fbthrift_field_message);
  }

  /** Glean { "field": "message" } */
  template <typename T_Bang_message_struct_setter = ::std::string>
  [[deprecated("Use `FOO.message() = BAR;` instead of `FOO.set_message(BAR);`")]]
  ::std::string& set_message(T_Bang_message_struct_setter&& message_) {
    message_ref() = std::forward<T_Bang_message_struct_setter>(message_);
    return __fbthrift_field_message;
  }

  template <class Protocol_>
  unsigned long read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t serializedSize(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t serializedSizeZC(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t write(Protocol_* prot_) const;

  const char* what() const noexcept override {
    return "::cpp2::Bang";
  }

 private:
  template <class Protocol_>
  void readNoXfer(Protocol_* iprot);

  friend class ::apache::thrift::Cpp2Ops<Bang>;
  friend void swap(Bang& a, Bang& b);
};

template <class Protocol_>
unsigned long Bang::read(Protocol_* iprot) {
  auto _xferStart = iprot->getCursorPosition();
  readNoXfer(iprot);
  return iprot->getCursorPosition() - _xferStart;
}


} // namespace cpp2

namespace apache { namespace thrift {

template <> struct TEnumDataStorage<::cpp2::U::Type>;

template <> struct TEnumTraits<::cpp2::U::Type> {
  using type = ::cpp2::U::Type;

  static constexpr std::size_t const size = 2;
  static folly::Range<type const*> const values;
  static folly::Range<std::string_view const*> const names;

  static bool findName(type value, std::string_view* out) noexcept;
  static bool findValue(std::string_view name, type* out) noexcept;

  static char const* findName(type value) noexcept {
    std::string_view ret;
    (void)findName(value, &ret);
    return ret.data();
  }
};
}} // apache::thrift
