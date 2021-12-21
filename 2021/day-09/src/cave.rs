use std::{collections::HashSet, num::ParseIntError, str::FromStr};

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
pub struct Position {
    pub row: usize,
    pub col: usize,
}

#[derive(Debug)]
pub struct Cave {
    heights: Vec<Vec<u32>>,
}

impl Cave {
    fn surrounding_positions(&self, pos: Position) -> Vec<Position> {
        let mut positions: Vec<Position> = vec![];
        if pos.row > 0 {
            positions.push(Position {
                row: pos.row - 1,
                col: pos.col,
            });
        }
        if pos.row < self.heights.len() - 1 {
            positions.push(Position {
                row: pos.row + 1,
                col: pos.col,
            });
        }
        if pos.col > 0 {
            positions.push(Position {
                row: pos.row,
                col: pos.col - 1,
            });
        }
        if pos.col < self.heights[pos.row].len() - 1 {
            positions.push(Position {
                row: pos.row,
                col: pos.col + 1,
            });
        }

        positions
    }

    pub fn find_low_points(&self) -> Vec<(Position, u32)> {
        let mut low_points: Vec<(Position, u32)> = vec![];
        for (row_num, row) in self.heights.iter().enumerate() {
            for (col_num, height) in row.iter().enumerate() {
                let surrounding_positions = self.surrounding_positions(Position {
                    row: row_num,
                    col: col_num,
                });
                if surrounding_positions
                    .iter()
                    .all(|h| &self.heights[h.row][h.col] > height)
                {
                    low_points.push((
                        Position {
                            row: row_num,
                            col: col_num,
                        },
                        *height,
                    ));
                }
            }
        }

        low_points
    }

    fn num_positions_in_basin(&self, pos: Position) -> usize {
        let mut num_positions = 1;

        let mut queue: Vec<Position> = vec![pos];
        let mut checked: HashSet<Position> = HashSet::new();
        checked.insert(pos);

        while queue.len() > 0 {
            let position = queue.pop().unwrap();
            let surrounding_positions = self.surrounding_positions(position);
            for p in surrounding_positions {
                if checked.contains(&p) {
                    continue;
                }
                if self.heights[p.row][p.col] != 9 {
                    checked.insert(p);
                    queue.push(p);
                    num_positions += 1;
                }
            }
        }
        num_positions
    }

    pub fn find_basin_sizes(&self) -> Vec<usize> {
        let mut basin_sizes: Vec<usize> = vec![];
        let low_points = self.find_low_points();
        for low_point in low_points {
            let s = self.num_positions_in_basin(low_point.0);
            basin_sizes.push(s);
        }

        basin_sizes
    }
}

impl FromStr for Cave {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let rows: Vec<&str> = s.trim().split('\n').collect();

        Ok(Cave {
            heights: rows
                .iter()
                .map(|row| row.chars().map(|c| c.to_digit(10).unwrap()).collect())
                .collect(),
        })
    }
}
