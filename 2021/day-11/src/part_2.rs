mod grid;
use std::io::{self, BufRead};

use grid::Grid;

fn main() {
    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let mut grid: Grid = input.parse().unwrap();
    let mut i = 0;
    loop {
        i += 1;
        grid = grid.step();
        if grid.all_flashed() {
            println!("{}", i);
            return;
        }
    }
}
