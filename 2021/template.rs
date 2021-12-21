use std::io::{self, BufRead};

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let mut answer = 0;

    println!("{}", answer);
}
