mod paper;
use std::io::{self, BufRead};

use crate::paper::Paper;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let mut paper: Paper = input.parse().unwrap();
    for fold in paper.folds.clone() {
        paper = paper.fold(fold);
    }
    println!("{}", paper);
}
