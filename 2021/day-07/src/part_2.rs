use std::{
    cmp::min,
    io::{self, BufRead},
};

fn main() {
    let mut positions: Vec<isize> = io::stdin()
        .lock()
        .lines()
        .next()
        .unwrap()
        .unwrap()
        .split(",")
        .map(|l| l.parse().unwrap())
        .collect();

    positions.sort_unstable();

    let mut min_cost: Option<isize> = None;
    for current_pos in *positions.first().unwrap()..(positions.last().unwrap() + 1) {
        let cost = positions.iter().fold(0, |acc, pos| {
            let diff = (pos - current_pos).abs();
            let distance = diff * (diff + 1) / 2;
            acc + distance
        });

        min_cost = match min_cost {
            None => Some(cost),
            Some(current_min) => Some(min(cost, current_min)),
        }
    }

    println!("{}", min_cost.unwrap());
}
