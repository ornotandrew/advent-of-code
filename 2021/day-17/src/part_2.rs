mod trench;
use std::io::{self, BufRead};

use crate::trench::Trench;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let trench: Trench = input.parse().unwrap();

    println!("{}", trench.all_trajectories_for_target());
}
