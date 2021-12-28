use std::{
    collections::{BinaryHeap, HashMap},
    fmt,
    str::FromStr,
};

use crate::path::Path;

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
pub struct Position {
    pub row: usize,
    pub col: usize,
}

type Grid = Vec<Vec<usize>>;

fn wrap(n: usize) -> usize {
    if n <= 9 {
        n
    } else {
        n % 9
    }
}

fn extend_downward(grid: &Grid, times: usize) -> Grid {
    let mut new_grid: Grid = vec![];
    for i in 0..times {
        for row in grid {
            let new_row: Vec<usize> = row.iter().map(|col| wrap(col + i)).collect();
            new_grid.push(new_row);
        }
    }

    new_grid
}

fn extend_rightward(grid: &Grid, times: usize) -> Grid {
    let mut new_grid: Grid = vec![];
    for row_num in 0..grid.len() {
        let mut new_row: Vec<usize> = vec![];
        for i in 0..times {
            let mut new_cols: Vec<usize> = grid[row_num].iter().map(|col| wrap(col + i)).collect();
            new_row.append(&mut new_cols);
        }
        new_grid.push(new_row);
    }

    new_grid
}

#[derive(Debug)]
pub struct Cave {
    grid: Grid,
}

enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Cave {
    pub fn expand_cave(&self) -> Cave {
        Cave {
            grid: extend_rightward(&extend_downward(&self.grid, 5), 5),
        }
    }

    fn next_paths(&self, path: &Path) -> Vec<Path> {
        let mut next_positions: Vec<Position> = vec![];
        for dir in [
            Direction::Up,
            Direction::Down,
            Direction::Left,
            Direction::Right,
        ] {
            match dir {
                Direction::Up => {
                    if path.pos.row == 0 {
                        continue;
                    }
                    next_positions.push(Position {
                        row: path.pos.row - 1,
                        col: path.pos.col,
                    });
                }
                Direction::Down => {
                    if path.pos.row == self.grid.len() - 1 {
                        continue;
                    }
                    next_positions.push(Position {
                        row: path.pos.row + 1,
                        col: path.pos.col,
                    });
                }
                Direction::Left => {
                    if path.pos.col == 0 {
                        continue;
                    }
                    next_positions.push(Position {
                        row: path.pos.row,
                        col: path.pos.col - 1,
                    });
                }
                Direction::Right => {
                    if path.pos.col == self.grid[path.pos.row].len() - 1 {
                        continue;
                    }
                    next_positions.push(Position {
                        row: path.pos.row,
                        col: path.pos.col + 1,
                    })
                }
            };
        }

        next_positions
            .into_iter()
            .map(|pos| Path {
                pos,
                end: path.end,
                cost: path.cost + self.grid[pos.row][pos.col],
            })
            .collect()
    }

    pub fn safest_route(&self) -> Path {
        let mut safest_cost_map: HashMap<Position, usize> = HashMap::new();
        let mut queue: BinaryHeap<Path> = BinaryHeap::new();

        let end = Position {
            row: self.grid.len() - 1,
            col: self.grid[0].len() - 1,
        };
        let start = Path {
            pos: Position { row: 0, col: 0 },
            cost: 0,
            end,
        };
        queue.push(start);

        loop {
            // the top-most path will always be the current cheapest route
            let path = queue.pop().unwrap(); // we assume there is always a route
            if path.pos == end {
                return path;
            }

            for next_path in self.next_paths(&path) {
                let current_lowest_cost = safest_cost_map.get(&next_path.pos);
                if current_lowest_cost.is_some() && *current_lowest_cost.unwrap() <= next_path.f() {
                    continue;
                };

                safest_cost_map.insert(next_path.pos, next_path.f());
                queue.push(next_path);
            }
        }
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Cave {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let grid: Grid = s
            .trim()
            .split("\n")
            .map(|line| {
                line.chars()
                    .map(|c| c.to_string().parse().unwrap())
                    .collect()
            })
            .collect();

        Ok(Cave { grid })
    }
}

impl fmt::Display for Cave {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        for row in &self.grid {
            for col in row {
                write!(f, "{}", col).unwrap();
            }
            write!(f, "\n").unwrap();
        }
        return Ok(());
    }
}
