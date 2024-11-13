// @generated by Thrift for thrift/compiler/test/fixtures/params/src/module.thrift
// This file is probably not the place you want to edit!

//! Server definitions for `module`.

#![recursion_limit = "100000000"]
#![allow(non_camel_case_types, non_snake_case, non_upper_case_globals, unused_crate_dependencies, unused_imports, clippy::all)]


#[doc(inline)]
pub use :: as types;

pub mod errors {
    #[doc(inline)]
    pub use ::::services::nested_containers;
    #[doc(inline)]
    #[allow(ambiguous_glob_reexports)]
    pub use ::::services::nested_containers::*;
}

pub(crate) use crate as server;
pub(crate) use ::::services;



#[::async_trait::async_trait]
pub trait NestedContainers: ::std::marker::Send + ::std::marker::Sync + 'static {
    async fn mapList(
        &self,
        _foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::vec::Vec<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::MapListExn> {
        ::std::result::Result::Err(crate::services::nested_containers::MapListExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "NestedContainers",
                "mapList",
            ),
        ))
    }
    async fn mapSet(
        &self,
        _foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::MapSetExn> {
        ::std::result::Result::Err(crate::services::nested_containers::MapSetExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "NestedContainers",
                "mapSet",
            ),
        ))
    }
    async fn listMap(
        &self,
        _foo: ::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::ListMapExn> {
        ::std::result::Result::Err(crate::services::nested_containers::ListMapExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "NestedContainers",
                "listMap",
            ),
        ))
    }
    async fn listSet(
        &self,
        _foo: ::std::vec::Vec<::std::collections::BTreeSet<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::ListSetExn> {
        ::std::result::Result::Err(crate::services::nested_containers::ListSetExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "NestedContainers",
                "listSet",
            ),
        ))
    }
    async fn turtles(
        &self,
        _foo: ::std::vec::Vec<::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>>>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::TurtlesExn> {
        ::std::result::Result::Err(crate::services::nested_containers::TurtlesExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "NestedContainers",
                "turtles",
            ),
        ))
    }
}

#[::async_trait::async_trait]
impl<T> NestedContainers for ::std::boxed::Box<T>
where
    T: NestedContainers + Send + Sync + ?Sized,
{
    async fn mapList(
        &self,
        foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::vec::Vec<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::MapListExn> {
        (**self).mapList(
            foo,
        ).await
    }
    async fn mapSet(
        &self,
        foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::MapSetExn> {
        (**self).mapSet(
            foo,
        ).await
    }
    async fn listMap(
        &self,
        foo: ::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::ListMapExn> {
        (**self).listMap(
            foo,
        ).await
    }
    async fn listSet(
        &self,
        foo: ::std::vec::Vec<::std::collections::BTreeSet<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::ListSetExn> {
        (**self).listSet(
            foo,
        ).await
    }
    async fn turtles(
        &self,
        foo: ::std::vec::Vec<::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>>>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::TurtlesExn> {
        (**self).turtles(
            foo,
        ).await
    }
}

#[::async_trait::async_trait]
impl<T> NestedContainers for ::std::sync::Arc<T>
where
    T: NestedContainers + Send + Sync + ?Sized,
{
    async fn mapList(
        &self,
        foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::vec::Vec<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::MapListExn> {
        (**self).mapList(
            foo,
        ).await
    }
    async fn mapSet(
        &self,
        foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::MapSetExn> {
        (**self).mapSet(
            foo,
        ).await
    }
    async fn listMap(
        &self,
        foo: ::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::ListMapExn> {
        (**self).listMap(
            foo,
        ).await
    }
    async fn listSet(
        &self,
        foo: ::std::vec::Vec<::std::collections::BTreeSet<::std::primitive::i32>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::ListSetExn> {
        (**self).listSet(
            foo,
        ).await
    }
    async fn turtles(
        &self,
        foo: ::std::vec::Vec<::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>>>>,
    ) -> ::std::result::Result<(), crate::services::nested_containers::TurtlesExn> {
        (**self).turtles(
            foo,
        ).await
    }
}
/// Processor for NestedContainers's methods.
#[derive(Clone, Debug)]
pub struct NestedContainersProcessor<P, H, R, RS> {
    service: H,
    supa: ::fbthrift::NullServiceProcessor<P, R, RS>,
    _phantom: ::std::marker::PhantomData<(P, H, R, RS)>,
}


struct Args_NestedContainers_mapList {
    foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::vec::Vec<::std::primitive::i32>>,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_NestedContainers_mapList {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "NestedContainers.mapList"))]
    fn read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("foo", ::fbthrift::TType::Map, 1),
        ];
        let mut field_foo = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::Map, 1) => field_foo = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            foo: field_foo.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "NestedContainers.mapList", "foo"))?,
        })
    }
}


struct Args_NestedContainers_mapSet {
    foo: ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_NestedContainers_mapSet {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "NestedContainers.mapSet"))]
    fn read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("foo", ::fbthrift::TType::Map, 1),
        ];
        let mut field_foo = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::Map, 1) => field_foo = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            foo: field_foo.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "NestedContainers.mapSet", "foo"))?,
        })
    }
}


struct Args_NestedContainers_listMap {
    foo: ::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::primitive::i32>>,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_NestedContainers_listMap {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "NestedContainers.listMap"))]
    fn read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("foo", ::fbthrift::TType::List, 1),
        ];
        let mut field_foo = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::List, 1) => field_foo = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            foo: field_foo.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "NestedContainers.listMap", "foo"))?,
        })
    }
}


struct Args_NestedContainers_listSet {
    foo: ::std::vec::Vec<::std::collections::BTreeSet<::std::primitive::i32>>,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_NestedContainers_listSet {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "NestedContainers.listSet"))]
    fn read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("foo", ::fbthrift::TType::List, 1),
        ];
        let mut field_foo = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::List, 1) => field_foo = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            foo: field_foo.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "NestedContainers.listSet", "foo"))?,
        })
    }
}


struct Args_NestedContainers_turtles {
    foo: ::std::vec::Vec<::std::vec::Vec<::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeMap<::std::primitive::i32, ::std::collections::BTreeSet<::std::primitive::i32>>>>>,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_NestedContainers_turtles {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "NestedContainers.turtles"))]
    fn read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("foo", ::fbthrift::TType::List, 1),
        ];
        let mut field_foo = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::List, 1) => field_foo = ::std::option::Option::Some(::fbthrift::Deserialize::read(p)?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            foo: field_foo.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "NestedContainers.turtles", "foo"))?,
        })
    }
}

impl<P, H, R, RS> NestedContainersProcessor<P, H, R, RS>
where
    P: ::fbthrift::Protocol + ::std::marker::Send + ::std::marker::Sync + 'static,
    P::Frame: ::std::marker::Send + 'static,
    P::Deserializer: ::std::marker::Send,
    H: NestedContainers,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<P::Frame, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = <P as ::fbthrift::Protocol>::Frame>
        + ::std::marker::Send + ::std::marker::Sync,
    ::fbthrift::ProtocolDecoded<P>: ::std::clone::Clone,
    ::fbthrift::ProtocolEncodedFinal<P>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    pub fn new(service: H) -> Self {
        Self {
            service,
            supa: ::fbthrift::NullServiceProcessor::new(),
            _phantom: ::std::marker::PhantomData,
        }
    }

    pub fn into_inner(self) -> H {
        self.service
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "NestedContainers.mapList"))]
    async fn handle_mapList<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"NestedContainers";
        const METHOD_NAME: &::std::ffi::CStr = c"mapList";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"NestedContainers.mapList";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_NestedContainers_mapList = ::fbthrift::Deserialize::read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.mapList(
                _args.foo,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "NestedContainers.mapList", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "NestedContainers.mapList", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("NestedContainers.mapList", exn);
                ::tracing::error!(method = "NestedContainers.mapList", panic = ?aexn);
                ::std::result::Result::Err(crate::services::nested_containers::MapListExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::nested_containers::MapListExn>(
            "mapList",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "NestedContainers.mapSet"))]
    async fn handle_mapSet<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"NestedContainers";
        const METHOD_NAME: &::std::ffi::CStr = c"mapSet";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"NestedContainers.mapSet";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_NestedContainers_mapSet = ::fbthrift::Deserialize::read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.mapSet(
                _args.foo,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "NestedContainers.mapSet", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "NestedContainers.mapSet", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("NestedContainers.mapSet", exn);
                ::tracing::error!(method = "NestedContainers.mapSet", panic = ?aexn);
                ::std::result::Result::Err(crate::services::nested_containers::MapSetExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::nested_containers::MapSetExn>(
            "mapSet",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "NestedContainers.listMap"))]
    async fn handle_listMap<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"NestedContainers";
        const METHOD_NAME: &::std::ffi::CStr = c"listMap";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"NestedContainers.listMap";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_NestedContainers_listMap = ::fbthrift::Deserialize::read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.listMap(
                _args.foo,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "NestedContainers.listMap", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "NestedContainers.listMap", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("NestedContainers.listMap", exn);
                ::tracing::error!(method = "NestedContainers.listMap", panic = ?aexn);
                ::std::result::Result::Err(crate::services::nested_containers::ListMapExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::nested_containers::ListMapExn>(
            "listMap",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "NestedContainers.listSet"))]
    async fn handle_listSet<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"NestedContainers";
        const METHOD_NAME: &::std::ffi::CStr = c"listSet";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"NestedContainers.listSet";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_NestedContainers_listSet = ::fbthrift::Deserialize::read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.listSet(
                _args.foo,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "NestedContainers.listSet", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "NestedContainers.listSet", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("NestedContainers.listSet", exn);
                ::tracing::error!(method = "NestedContainers.listSet", panic = ?aexn);
                ::std::result::Result::Err(crate::services::nested_containers::ListSetExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::nested_containers::ListSetExn>(
            "listSet",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "NestedContainers.turtles"))]
    async fn handle_turtles<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"NestedContainers";
        const METHOD_NAME: &::std::ffi::CStr = c"turtles";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"NestedContainers.turtles";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_NestedContainers_turtles = ::fbthrift::Deserialize::read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.turtles(
                _args.foo,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "NestedContainers.turtles", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "NestedContainers.turtles", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("NestedContainers.turtles", exn);
                ::tracing::error!(method = "NestedContainers.turtles", panic = ?aexn);
                ::std::result::Result::Err(crate::services::nested_containers::TurtlesExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::nested_containers::TurtlesExn>(
            "turtles",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }
}

#[::async_trait::async_trait]
impl<P, H, R, RS> ::fbthrift::ServiceProcessor<P> for NestedContainersProcessor<P, H, R, RS>
where
    P: ::fbthrift::Protocol + ::std::marker::Send + ::std::marker::Sync + 'static,
    P::Deserializer: ::std::marker::Send,
    H: NestedContainers,
    P::Frame: ::std::marker::Send + 'static,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = <P as ::fbthrift::Protocol>::Frame>
        + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<P::Frame, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    ::fbthrift::ProtocolDecoded<P>: ::std::clone::Clone,
    ::fbthrift::ProtocolEncodedFinal<P>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    type RequestContext = R;
    type ReplyState = RS;

    #[inline]
    fn method_idx(&self, name: &[::std::primitive::u8]) -> ::std::result::Result<::std::primitive::usize, ::fbthrift::ApplicationException> {
        match name {
            b"mapList" => ::std::result::Result::Ok(0usize),
            b"mapSet" => ::std::result::Result::Ok(1usize),
            b"listMap" => ::std::result::Result::Ok(2usize),
            b"listSet" => ::std::result::Result::Ok(3usize),
            b"turtles" => ::std::result::Result::Ok(4usize),
            _ => ::std::result::Result::Err(::fbthrift::ApplicationException::unknown_method()),
        }
    }

    #[allow(clippy::match_single_binding)]
    async fn handle_method(
        &self,
        idx: ::std::primitive::usize,
        _p: &mut P::Deserializer,
        _req: ::fbthrift::ProtocolDecoded<P>,
        _req_ctxt: &R,
        _reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        match idx {
            0usize => {
                self.handle_mapList(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            1usize => {
                self.handle_mapSet(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            2usize => {
                self.handle_listMap(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            3usize => {
                self.handle_listSet(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            4usize => {
                self.handle_turtles(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            bad => panic!(
                "{}: unexpected method idx {}",
                "NestedContainersProcessor",
                bad
            ),
        }
    }

    #[allow(clippy::match_single_binding)]
    #[inline]
    fn create_interaction_idx(&self, name: &::std::primitive::str) -> ::anyhow::Result<::std::primitive::usize> {
        match name {
            _ => ::anyhow::bail!("Unknown interaction"),
        }
    }

    #[allow(clippy::match_single_binding)]
    fn handle_create_interaction(
        &self,
        idx: ::std::primitive::usize,
    ) -> ::anyhow::Result<
        ::std::sync::Arc<dyn ::fbthrift::ThriftService<P::Frame, Handler = (), RequestContext = Self::RequestContext, ReplyState = Self::ReplyState> + ::std::marker::Send + 'static>
    > {
        match idx {
            bad => panic!(
                "{}: unexpected method idx {}",
                "NestedContainersProcessor",
                bad
            ),
        }
    }

    async fn handle_on_termination(&self) {
    }
}

#[::async_trait::async_trait]
impl<P, H, R, RS> ::fbthrift::ThriftService<P::Frame> for NestedContainersProcessor<P, H, R, RS>
where
    P: ::fbthrift::Protocol + ::std::marker::Send + ::std::marker::Sync + 'static,
    P::Deserializer: ::std::marker::Send,
    P::Frame: ::std::marker::Send + 'static,
    H: NestedContainers,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = <P as ::fbthrift::Protocol>::Frame>
        + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<P::Frame, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    ::fbthrift::ProtocolDecoded<P>: ::std::clone::Clone,
    ::fbthrift::ProtocolEncodedFinal<P>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    type Handler = H;
    type RequestContext = R;
    type ReplyState = RS;

    #[tracing::instrument(level="trace", skip_all, fields(service = "NestedContainers"))]
    async fn call(
        &self,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
    ) -> ::anyhow::Result<()> {
        use ::fbthrift::{ProtocolReader as _, ServiceProcessor as _};
        let mut p = P::deserializer(req.clone());
        let (idx, mty, seqid) = p.read_message_begin(|name| self.method_idx(name))?;
        if mty != ::fbthrift::MessageType::Call {
            return ::std::result::Result::Err(::std::convert::From::from(::fbthrift::ApplicationException::new(
                ::fbthrift::ApplicationExceptionErrorCode::InvalidMessageType,
                format!("message type {:?} not handled", mty)
            )));
        }
        let idx = match idx {
            ::std::result::Result::Ok(idx) => idx,
            ::std::result::Result::Err(_) => {
                return self.supa.call(req, req_ctxt, reply_state).await;
            }
        };
        self.handle_method(idx, &mut p, req, req_ctxt, reply_state, seqid).await?;
        p.read_message_end()?;

        ::std::result::Result::Ok(())
    }

    fn create_interaction(
        &self,
        name: &::std::primitive::str,
    ) -> ::anyhow::Result<
        ::std::sync::Arc<dyn ::fbthrift::ThriftService<P::Frame, Handler = (), RequestContext = R, ReplyState = RS> + ::std::marker::Send + 'static>
    > {
        use ::fbthrift::{ServiceProcessor as _};
        let idx = self.create_interaction_idx(name);
        let idx = match idx {
            ::anyhow::Result::Ok(idx) => idx,
            ::anyhow::Result::Err(_) => {
                return self.supa.create_interaction(name);
            }
        };
        self.handle_create_interaction(idx)
    }

    fn get_method_metadata(&self) -> &'static [::fbthrift::processor::MethodMetadata] {
        &[
            // From module.NestedContainers:
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "mapList",
              starts_interaction: false,
              interaction_name: None,
            },
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "mapSet",
              starts_interaction: false,
              interaction_name: None,
            },
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "listMap",
              starts_interaction: false,
              interaction_name: None,
            },
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "listSet",
              starts_interaction: false,
              interaction_name: None,
            },
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "turtles",
              starts_interaction: false,
              interaction_name: None,
            },
        ]
    }

    async fn on_termination(&self) {
        use ::fbthrift::{ServiceProcessor as _};
        self.handle_on_termination().await
    }
}

/// Construct a new instance of a NestedContainers service.
///
/// This is called when a new instance of a Thrift service Processor
/// is needed for a particular Thrift protocol.
#[::tracing::instrument(level="debug", skip_all, fields(proto = ?proto))]
pub fn make_NestedContainers_server<F, H, R, RS>(
    proto: ::fbthrift::ProtocolID,
    handler: H,
) -> ::std::result::Result<::std::boxed::Box<dyn ::fbthrift::ThriftService<F, Handler = H, RequestContext = R, ReplyState = RS> + ::std::marker::Send + 'static>, ::fbthrift::ApplicationException>
where
    F: ::fbthrift::Framing + ::std::marker::Send + ::std::marker::Sync + 'static,
    H: NestedContainers,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = F> + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<F, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    ::fbthrift::FramingDecoded<F>: ::std::clone::Clone,
    ::fbthrift::FramingEncodedFinal<F>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    match proto {
        ::fbthrift::ProtocolID::BinaryProtocol => {
            ::std::result::Result::Ok(::std::boxed::Box::new(NestedContainersProcessor::<::fbthrift::BinaryProtocol<F>, H, R, RS>::new(handler)))
        }
        ::fbthrift::ProtocolID::CompactProtocol => {
            ::std::result::Result::Ok(::std::boxed::Box::new(NestedContainersProcessor::<::fbthrift::CompactProtocol<F>, H, R, RS>::new(handler)))
        }
        bad => {
            ::tracing::error!(method = "NestedContainers.", invalid_protocol = ?bad);
            ::std::result::Result::Err(::fbthrift::ApplicationException::invalid_protocol(bad))
        }
    }
}
