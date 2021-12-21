mod display;
mod segment;
use crate::display::Display;
use std::io::{self, BufRead};

const KNOWN_LENGTHS: [usize; 4] = [2, 3, 4, 7];

fn main() {
    let input: Vec<Display> = io::stdin()
        .lock()
        .lines()
        .map(|l| Display::new(l.unwrap()))
        .collect();

    let mut answer = 0;
    for display in &input {
        for output_value in &display.output_values {
            if KNOWN_LENGTHS.contains(&output_value.len()) {
                answer += 1;
            }
        }
    }

    println!("{}", answer);
}
