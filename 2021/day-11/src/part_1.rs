mod grid;
use std::io::{self, BufRead};

use grid::Grid;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let mut grid: Grid = input.parse().unwrap();
    for _ in 0..100 {
        grid = grid.step();
    }

    println!("{}", grid.total_flashes);
}
