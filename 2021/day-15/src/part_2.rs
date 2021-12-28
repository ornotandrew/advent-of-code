mod cave;
mod path;
use std::io::{self, BufRead};

use crate::cave::Cave;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let cave: Cave = input.parse().unwrap();
    let cave = cave.expand_cave();
    let path = cave.safest_route();

    println!("{}", path.cost);
}
