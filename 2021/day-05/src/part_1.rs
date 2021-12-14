mod ocean;
use std::io::{self, BufRead};

use crate::ocean::Ocean;

fn main() {
    let input: Vec<String> = io::stdin().lock().lines().map(|l| l.unwrap()).collect();
    let ocean = Ocean::load_from_lines(input, false);

    let answer = ocean
        .position_counts
        .iter()
        .filter(|(_, count)| **count > 1)
        .count();

    println!("{}", answer);
}
