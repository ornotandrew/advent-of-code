use std::io::{self, BufRead};

fn main() {
    let input: Vec<i32> = io::stdin()
        .lock()
        .lines()
        .map(|l| l.unwrap().parse().unwrap())
        .collect();

    let mut answer = 0;

    println!("{}", answer);
}
