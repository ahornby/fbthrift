// @generated by autocargo

use std::env;
use std::fs;
use std::path::Path;
use thrift_compiler::Config;
use thrift_compiler::GenContext;
const CRATEMAP: &str = "\
hack crate //thrift/annotation:hack-rust
scope scope //thrift/annotation:scope-rust
";
#[rustfmt::skip]
fn main() {
    println!("cargo:rerun-if-changed=thrift_build.rs");
    let out_dir = env::var_os("OUT_DIR").expect("OUT_DIR env not provided");
    let cratemap_path = Path::new(&out_dir).join("cratemap");
    fs::write(cratemap_path, CRATEMAP).expect("Failed to write cratemap");
    Config::from_env(GenContext::Types)
        .expect("Failed to instantiate thrift_compiler::Config")
        .base_path("../../../../..")
        .types_crate("hack__types")
        .clients_crate("hack__clients")
        .run(["../../../../annotation/hack.thrift"])
        .expect("Failed while running thrift compilation");
}
