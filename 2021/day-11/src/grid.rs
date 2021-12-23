use std::{collections::HashSet, fmt, str::FromStr};

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
pub struct Position {
    pub row: usize,
    pub col: usize,
}

#[derive(Debug)]
pub struct Grid {
    pub total_flashes: u64,
    rows: Vec<Vec<u32>>,
}

impl Grid {
    fn get_surrounding_positions(&self, pos: Position) -> Vec<Position> {
        let mut surrounding: Vec<Position> = vec![];
        for row_offset in -1..2 {
            for col_offset in -1..2 {
                let row = pos.row as i32 + row_offset;
                let col = pos.col as i32 + col_offset;
                if row < 0 || col < 0 || (row_offset == 0 && col_offset == 0) {
                    continue;
                }
                let row = row as usize;
                let col = col as usize;
                if row >= self.rows.len() || col >= self.rows[row].len() {
                    continue;
                }
                surrounding.push(Position { row, col });
            }
        }
        surrounding
    }

    pub fn step(&self) -> Self {
        let mut next = Grid {
            total_flashes: self.total_flashes,
            rows: self.rows.clone(),
        };

        // first, increase everything by one
        for row in next.rows.iter_mut() {
            for col in row.iter_mut() {
                *col += 1;
            }
        }

        // then, flash each cell with value > 9 until they stop flashing
        let mut flashed_positions: HashSet<Position> = HashSet::new();
        let mut to_increment: Vec<Position> = vec![];
        loop {
            for pos in to_increment {
                for surrounding_pos in self.get_surrounding_positions(pos) {
                    next.rows[surrounding_pos.row][surrounding_pos.col] += 1;
                }
            }
            to_increment = vec![];

            for row in 0..next.rows.len() {
                for col in 0..next.rows[row].len() {
                    let pos = Position { row, col };
                    if next.rows[row][col] > 9 {
                        flashed_positions.insert(pos);
                        next.total_flashes += 1;
                        next.rows[row][col] = 0;
                        to_increment.push(pos);
                    }
                }
            }
            if to_increment.len() == 0 {
                break;
            }
        }

        // set any cell that flashed back to zero
        for pos in flashed_positions {
            next.rows[pos.row][pos.col] = 0;
        }

        next
    }

    pub fn all_flashed(&self) -> bool {
        self.rows.iter().all(|row| row.iter().all(|col| *col == 0))
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Grid {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Grid {
            total_flashes: 0,
            rows: s
                .trim()
                .split('\n')
                .map(|l| l.chars().map(|c| c.to_string().parse().unwrap()).collect())
                .collect(),
        })
    }
}

impl fmt::Display for Grid {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        for row in &self.rows {
            write!(f, "\n").unwrap();
            for col in row {
                write!(f, "{:>3 }", col).unwrap()
            }
        }

        write!(f, "\n({})\n", self.total_flashes).unwrap();
        return Ok(());
    }
}
