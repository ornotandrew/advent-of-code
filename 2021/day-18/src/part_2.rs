mod homework;
use std::io::{self, BufRead};

use crate::homework::Homework;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let homework: Homework = input.parse().unwrap();

    println!("{}", homework.largest_magnitude());
}
