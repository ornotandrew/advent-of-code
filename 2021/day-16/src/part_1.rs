mod packet;
use std::io::{self, BufRead};

use crate::packet::Packet;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let packet: Packet = input.parse().unwrap();

    println!("{}", packet.sum_version_numbers());
}
