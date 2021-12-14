mod school;
use std::io::{self, BufRead};

use crate::school::School;

fn main() {
    let input: Vec<usize> = io::stdin()
        .lock()
        .lines()
        .next()
        .unwrap()
        .unwrap()
        .split(",")
        .map(|l| l.parse().unwrap())
        .collect();

    let mut school = School::new(input);
    for _ in 0..256 {
        school.simulate_day();
    }

    println!("{}", school.get_total_fish());
}
