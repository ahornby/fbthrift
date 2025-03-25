// @generated by Thrift for thrift/compiler/test/fixtures/int_limits/src/module.thrift
// This file is probably not the place you want to edit!


#![recursion_limit = "100000000"]
#![allow(non_camel_case_types, non_snake_case, non_upper_case_globals, unused_crate_dependencies, clippy::redundant_closure, clippy::type_complexity)]

pub mod consts;
#[doc(inline)]
pub use self::consts::*;
#[allow(unused_imports)]
pub(crate) use crate as types;

#[derive(Clone, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct Limits {
    pub max_i64_field: ::std::primitive::i64,
    pub min_i64_field: ::std::primitive::i64,
    pub max_i32_field: ::std::primitive::i32,
    pub min_i32_field: ::std::primitive::i32,
    pub max_i16_field: ::std::primitive::i16,
    pub min_i16_field: ::std::primitive::i16,
    pub max_byte_field: ::std::primitive::i8,
    pub min_byte_field: ::std::primitive::i8,
    // This field forces `..Default::default()` when instantiating this
    // struct, to make code future-proof against new fields added later to
    // the definition in Thrift. If you don't want this, add the annotation
    // `@rust.Exhaustive` to the Thrift struct to eliminate this field.
    #[doc(hidden)]
    pub _dot_dot_Default_default: self::dot_dot::OtherFields,
}



#[allow(clippy::derivable_impls)]
impl ::std::default::Default for self::Limits {
    fn default() -> Self {
        Self {
            max_i64_field: 9223372036854775807,
            min_i64_field: -9223372036854775808,
            max_i32_field: 2147483647,
            min_i32_field: -2147483648,
            max_i16_field: 32767,
            min_i16_field: -32768,
            max_byte_field: 127,
            min_byte_field: -128,
            _dot_dot_Default_default: self::dot_dot::OtherFields(()),
        }
    }
}

impl ::std::fmt::Debug for self::Limits {
    fn fmt(&self, formatter: &mut ::std::fmt::Formatter) -> ::std::fmt::Result {
        formatter
            .debug_struct("Limits")
            .field("max_i64_field", &self.max_i64_field)
            .field("min_i64_field", &self.min_i64_field)
            .field("max_i32_field", &self.max_i32_field)
            .field("min_i32_field", &self.min_i32_field)
            .field("max_i16_field", &self.max_i16_field)
            .field("min_i16_field", &self.min_i16_field)
            .field("max_byte_field", &self.max_byte_field)
            .field("min_byte_field", &self.min_byte_field)
            .finish()
    }
}

unsafe impl ::std::marker::Send for self::Limits {}
unsafe impl ::std::marker::Sync for self::Limits {}
impl ::std::marker::Unpin for self::Limits {}
impl ::std::panic::RefUnwindSafe for self::Limits {}
impl ::std::panic::UnwindSafe for self::Limits {}

impl ::fbthrift::GetTType for self::Limits {
    const TTYPE: ::fbthrift::TType = ::fbthrift::TType::Struct;
}

impl ::fbthrift::GetTypeNameType for self::Limits {
    fn type_name_type() -> fbthrift::TypeNameType {
        ::fbthrift::TypeNameType::StructType
    }
}

impl<P> ::fbthrift::Serialize<P> for self::Limits
where
    P: ::fbthrift::ProtocolWriter,
{
    #[inline]
    fn write(&self, p: &mut P) {
        p.write_struct_begin("Limits");
        p.write_field_begin("max_i64_field", ::fbthrift::TType::I64, 1);
        ::fbthrift::Serialize::write(&self.max_i64_field, p);
        p.write_field_end();
        p.write_field_begin("min_i64_field", ::fbthrift::TType::I64, 2);
        ::fbthrift::Serialize::write(&self.min_i64_field, p);
        p.write_field_end();
        p.write_field_begin("max_i32_field", ::fbthrift::TType::I32, 3);
        ::fbthrift::Serialize::write(&self.max_i32_field, p);
        p.write_field_end();
        p.write_field_begin("min_i32_field", ::fbthrift::TType::I32, 4);
        ::fbthrift::Serialize::write(&self.min_i32_field, p);
        p.write_field_end();
        p.write_field_begin("max_i16_field", ::fbthrift::TType::I16, 5);
        ::fbthrift::Serialize::write(&self.max_i16_field, p);
        p.write_field_end();
        p.write_field_begin("min_i16_field", ::fbthrift::TType::I16, 6);
        ::fbthrift::Serialize::write(&self.min_i16_field, p);
        p.write_field_end();
        p.write_field_begin("max_byte_field", ::fbthrift::TType::Byte, 7);
        ::fbthrift::Serialize::write(&self.max_byte_field, p);
        p.write_field_end();
        p.write_field_begin("min_byte_field", ::fbthrift::TType::Byte, 8);
        ::fbthrift::Serialize::write(&self.min_byte_field, p);
        p.write_field_end();
        p.write_field_stop();
        p.write_struct_end();
    }
}

impl<P> ::fbthrift::Deserialize<P> for self::Limits
where
    P: ::fbthrift::ProtocolReader,
{
    #[inline]
    fn read(p: &mut P) -> ::anyhow::Result<Self> {
        static FIELDS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("max_byte_field", ::fbthrift::TType::Byte, 7),
            ::fbthrift::Field::new("max_i16_field", ::fbthrift::TType::I16, 5),
            ::fbthrift::Field::new("max_i32_field", ::fbthrift::TType::I32, 3),
            ::fbthrift::Field::new("max_i64_field", ::fbthrift::TType::I64, 1),
            ::fbthrift::Field::new("min_byte_field", ::fbthrift::TType::Byte, 8),
            ::fbthrift::Field::new("min_i16_field", ::fbthrift::TType::I16, 6),
            ::fbthrift::Field::new("min_i32_field", ::fbthrift::TType::I32, 4),
            ::fbthrift::Field::new("min_i64_field", ::fbthrift::TType::I64, 2),
        ];
        let mut field_max_i64_field = ::std::option::Option::None;
        let mut field_min_i64_field = ::std::option::Option::None;
        let mut field_max_i32_field = ::std::option::Option::None;
        let mut field_min_i32_field = ::std::option::Option::None;
        let mut field_max_i16_field = ::std::option::Option::None;
        let mut field_min_i16_field = ::std::option::Option::None;
        let mut field_max_byte_field = ::std::option::Option::None;
        let mut field_min_byte_field = ::std::option::Option::None;
        let _ = ::anyhow::Context::context(p.read_struct_begin(|_| ()), "Expected a Limits")?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), FIELDS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::I64, 1) => field_max_i64_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::I64, 2) => field_min_i64_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::I32, 3) => field_max_i32_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::I32, 4) => field_min_i32_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::I16, 5) => field_max_i16_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::I16, 6) => field_min_i16_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::Byte, 7) => field_max_byte_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (::fbthrift::TType::Byte, 8) => field_min_byte_field = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            max_i64_field: field_max_i64_field.unwrap_or(9223372036854775807),
            min_i64_field: field_min_i64_field.unwrap_or(-9223372036854775808),
            max_i32_field: field_max_i32_field.unwrap_or(2147483647),
            min_i32_field: field_min_i32_field.unwrap_or(-2147483648),
            max_i16_field: field_max_i16_field.unwrap_or(32767),
            min_i16_field: field_min_i16_field.unwrap_or(-32768),
            max_byte_field: field_max_byte_field.unwrap_or(127),
            min_byte_field: field_min_byte_field.unwrap_or(-128),
            _dot_dot_Default_default: self::dot_dot::OtherFields(()),
        })
    }
}


impl ::fbthrift::metadata::ThriftAnnotations for Limits {
    fn get_structured_annotation<T: Sized + 'static>() -> ::std::option::Option<T> {
        #[allow(unused_variables)]
        let type_id = ::std::any::TypeId::of::<T>();

        ::std::option::Option::None
    }

    fn get_field_structured_annotation<T: Sized + 'static>(field_id: ::std::primitive::i16) -> ::std::option::Option<T> {
        #[allow(unused_variables)]
        let type_id = ::std::any::TypeId::of::<T>();

        #[allow(clippy::match_single_binding)]
        match field_id {
            1 => {
            },
            2 => {
            },
            3 => {
            },
            4 => {
            },
            5 => {
            },
            6 => {
            },
            7 => {
            },
            8 => {
            },
            _ => {}
        }

        ::std::option::Option::None
    }
}



mod dot_dot {
    #[derive(Copy, Clone, PartialEq, Eq, PartialOrd, Ord, Hash)]
    pub struct OtherFields(pub(crate) ());

    #[allow(dead_code)] // if serde isn't being used
    pub(super) fn default_for_serde_deserialize() -> OtherFields {
        OtherFields(())
    }
}

pub(crate) mod r#impl {
    use ::ref_cast::RefCast;

    #[derive(RefCast)]
    #[repr(transparent)]
    pub(crate) struct LocalImpl<T>(T);

    #[allow(unused)]
    pub(crate) fn write<T, P>(value: &T, p: &mut P)
    where
        LocalImpl<T>: ::fbthrift::Serialize<P>,
        P: ::fbthrift::ProtocolWriter,
    {
        ::fbthrift::Serialize::write(LocalImpl::ref_cast(value), p);
    }

    #[allow(unused)]
    pub(crate) fn read<T, P>(p: &mut P) -> ::anyhow::Result<T>
    where
        LocalImpl<T>: ::fbthrift::Deserialize<P>,
        P: ::fbthrift::ProtocolReader,
    {
        let value: LocalImpl<T> = ::fbthrift::Deserialize::read(p)?;
        ::std::result::Result::Ok(value.0)
    }
}


#[doc(hidden)]
#[deprecated]
#[allow(hidden_glob_reexports)]
pub mod __constructors {
}
