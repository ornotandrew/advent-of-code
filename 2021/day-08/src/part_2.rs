mod display;
mod segment;
use crate::display::Display;
use std::io::{self, BufRead};

fn main() {
    let input: Vec<Display> = io::stdin()
        .lock()
        .lines()
        .map(|l| Display::new(l.unwrap()))
        .collect();

    let mut answer = 0;

    for display in &input {
        let numbers = display.derive_numbers();
        for (i, output_value) in (&display.output_values).iter().rev().enumerate() {
            let number = numbers.iter().position(|n| n == output_value).unwrap();
            answer += number * 10_usize.pow(i.try_into().unwrap());
        }
    }

    println!("{}", answer);
}
