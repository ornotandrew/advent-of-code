mod polymer;
use std::io::{self, BufRead};

use crate::polymer::Polymer;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let mut polymer: Polymer = input.parse().unwrap();
    for _ in 0..10 {
        polymer.step();
    }

    println!("{}", polymer.get_difference());
}
