use std::io::{self, BufRead};

enum Direction {
    Forward,
    Down,
    Up,
    Unknown,
}

fn main() {
    let input: Vec<(Direction, i32)> = io::stdin()
        .lock()
        .lines()
        .map(|l| {
            let line = l.unwrap();
            let parts: Vec<&str> = line.split(" ").collect();
            let direction = match parts[0] {
                "forward" => Direction::Forward,
                "down" => Direction::Down,
                "up" => Direction::Up,
                _ => Direction::Unknown,
            };
            let units: i32 = parts[1].parse().unwrap();
            (direction, units)
        })
        .collect();

    let mut horizontal = 0;
    let mut depth = 0;
    for (direction, units) in input {
        match direction {
            Direction::Forward => horizontal += units,
            Direction::Down => depth += units,
            Direction::Up => depth -= units,
            Direction::Unknown => {}
        }
    }

    println!("{}", horizontal * depth);
}
