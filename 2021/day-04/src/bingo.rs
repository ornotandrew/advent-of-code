use std::fmt;

pub struct Game {
    pub boards: Vec<Board>,
}

pub struct Board {
    pub size: usize,
    pub numbers: Vec<Vec<Option<usize>>>,
}

impl Board {
    pub fn rows(&self) -> Vec<Vec<Option<usize>>> {
        self.numbers.iter().cloned().collect()
    }

    pub fn cols(&self) -> Vec<Vec<Option<usize>>> {
        let mut cols: Vec<Vec<Option<usize>>> = vec![];
        for i in 0..self.size {
            let mut col: Vec<Option<usize>> = vec![];
            for row in &(self.numbers) {
                col.push(row[i]);
            }
            cols.push(col);
        }
        cols.iter().cloned().collect()
    }

    pub fn has_won(&self) -> bool {
        fn slice_is_empty(slice: &Vec<Option<usize>>) -> bool {
            !slice.iter().any(|n| match n {
                Some(_) => true,
                None => false,
            })
        }
        self.rows().iter().any(|row| slice_is_empty(row))
            || self.cols().iter().any(|row| slice_is_empty(row))
    }

    pub fn get_score(&self) -> usize {
        let mut acc: usize = 0;
        for row in &(self.numbers) {
            for maybe_col in row {
                match maybe_col {
                    Some(col) => acc += col,
                    None => {}
                }
            }
        }
        acc
    }
}

impl Game {
    pub fn from_lines(lines: &[String]) -> Game {
        let mut boards: Vec<Board> = vec![];
        let mut current_board = Board {
            size: 0,
            numbers: vec![],
        };

        for line in lines {
            if line == &"" {
                current_board.size = current_board.numbers.len();
                boards.push(current_board);
                current_board = Board {
                    size: 0,
                    numbers: vec![],
                };
                continue;
            }
            current_board.numbers.push(
                line.split(" ")
                    .filter(|x| x != &"")
                    .map(|n| Some(n.parse().unwrap()))
                    .collect(),
            );
        }
        current_board.size = current_board.numbers.len();
        boards.push(current_board);
        Game { boards }
    }

    pub fn draw_number(&mut self, n: usize) {
        for board in &mut (self.boards) {
            for row in &mut (board.numbers) {
                for maybe_col in row {
                    match maybe_col {
                        None => {}
                        Some(col) => {
                            if *col == n {
                                *maybe_col = None;
                            }
                        }
                    }
                }
            }
        }
    }
}

impl fmt::Display for Board {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "Board {{").unwrap();

        for row in &self.numbers {
            write!(f, "\n  ").unwrap();
            for col in row {
                match col {
                    None => write!(f, " x ").unwrap(),
                    Some(col) => write!(f, "{: >2} ", col).unwrap(),
                }
            }
        }

        write!(f, "\n}}").unwrap();
        return Ok(());
    }
}
