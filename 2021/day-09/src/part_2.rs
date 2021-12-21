mod cave;
use std::io::{self, BufRead};

use crate::cave::Cave;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let cave: Cave = input.parse().unwrap();
    let mut basin_sizes = cave.find_basin_sizes();
    basin_sizes.sort_by(|a, b| b.cmp(a));

    println!("{}", basin_sizes[0] * basin_sizes[1] * basin_sizes[2]);
}
