mod cave;
use std::io::{self, BufRead};

use crate::cave::Cave;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let cave: Cave = input.parse().unwrap();

    println!("{}", cave.all_possible_paths(false).len());
}
