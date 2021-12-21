mod cave;
use std::io::{self, BufRead};

use crate::cave::Cave;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let cave: Cave = input.parse().unwrap();
    let low_points = cave.find_low_points();

    println!("{}", low_points.iter().fold(0, |acc, cur| acc + cur.1 + 1));
}
