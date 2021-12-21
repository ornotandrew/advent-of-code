use std::io::{self, BufRead};

fn main() {
    let mut positions: Vec<isize> = io::stdin()
        .lock()
        .lines()
        .next()
        .unwrap()
        .unwrap()
        .split(",")
        .map(|l| l.parse().unwrap())
        .collect();

    positions.sort_unstable();

    let median = positions[positions.len() / 2];
    let answer = positions
        .iter()
        .fold(0, |acc, pos| acc + (pos - median).abs());

    println!("{}", answer);
}
