mod syntax;
use std::io::{self, BufRead};

use crate::syntax::Syntax;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let syntax: Syntax = input.parse().unwrap();
    let mut scores = Syntax::incomplete_scores(syntax.find_incomplete());
    scores.sort();

    println!("{}", scores[scores.len() / 2]);
}
