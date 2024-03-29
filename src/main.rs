extern crate clap;

use std::convert::Infallible;
use std::net::SocketAddr;

use hyper::{Server};
use hyper::service::{make_service_fn, service_fn};
use clap::{Arg, App};

mod fastq;
mod api;
mod tokenizer;

#[tokio::main]
async fn main()
{
    let matches = App::new("GenomDB")
        .version("0.1.0")
        .author("Billy Hadlow")
        .about("GenomDB")
        .arg(Arg::with_name("config")
             .short("c")
             .long("config")
             .help("Config file")
             .required(false)
             .takes_value(true))
        .arg(Arg::with_name("shard")
             .short("s")
             .long("shard")
             .help("The shard ID")
             .required(true)
             .takes_value(true))
        .arg(Arg::with_name("port")
             .short("p")
             .long("port")
             .help("Port to host the shard on")
             .required(true)
             .takes_value(true))
        .get_matches();

    let _shard = matches.value_of("shard").unwrap_or("No shard ID has been provided");
    let port = matches.value_of("port").unwrap_or("No port has been provided");
    let _config = matches.value_of("config").unwrap_or("No config file has been provided");
    let file = "test.fastq";

    let addr = SocketAddr::from(([127, 0, 0, 1], port.parse::<u16>().unwrap()));

    let make_svc = make_service_fn(|_conn| async
    {
        Ok::<_, Infallible>(service_fn(api::routes))
    });

    let server = Server::bind(&addr).serve(make_svc);

    if let Err(e) = server.await
    {
        eprintln!("server error: {}", e);
    }
}
