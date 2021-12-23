use std::fmt;
use std::{collections::HashSet, str::FromStr};

#[derive(PartialEq, Eq, Hash, Debug)]
pub struct Point {
    x: usize,
    y: usize,
}

#[derive(Debug, Copy, Clone)]
pub enum Fold {
    X(usize),
    Y(usize),
}

#[derive(Debug)]
pub struct Paper {
    pub points: HashSet<Point>,
    pub folds: Vec<Fold>,
}

impl Paper {
    pub fn fold(&self, fold: Fold) -> Paper {
        let mut next = Paper {
            points: HashSet::new(),
            folds: self.folds.clone(),
        };

        for point in &self.points {
            next.points.insert(match fold {
                Fold::X(x) => Point {
                    x: {
                        if point.x <= x {
                            point.x
                        } else {
                            x - (point.x - x)
                        }
                    },
                    y: point.y,
                },
                Fold::Y(y) => Point {
                    x: point.x,
                    y: {
                        if point.y <= y {
                            point.y
                        } else {
                            y - (point.y - y)
                        }
                    },
                },
            });
        }

        next
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Paper {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut points: HashSet<Point> = HashSet::new();
        let mut folds: Vec<Fold> = vec![];
        for line in s.trim().split("\n").collect::<Vec<&str>>() {
            if line == "" {
                continue;
            }
            if line.contains("fold") {
                let coord = line.split_once("fold along ").unwrap().1;
                let (axis, position) = coord.split_once("=").unwrap();
                let position: usize = position.parse().unwrap();
                folds.push(match axis {
                    "x" => Fold::X(position),
                    "y" => Fold::Y(position),
                    _ => unreachable!(),
                });
                continue;
            }

            let (x, y) = line.split_once(",").unwrap();
            points.insert(Point {
                x: x.parse().unwrap(),
                y: y.parse().unwrap(),
            });
        }

        Ok(Paper { points, folds })
    }
}

impl fmt::Display for Paper {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let max_x = self.points.iter().map(|p| p.x).max().unwrap();
        let max_y = self.points.iter().map(|p| p.y).max().unwrap();

        for y in 0..max_y + 1 {
            for x in 0..max_x + 1 {
                write!(
                    f,
                    "{}",
                    match self.points.contains(&Point { x, y }) {
                        true => "#",
                        false => ".",
                    }
                )
                .unwrap()
            }
            write!(f, "\n").unwrap();
        }

        return Ok(());
    }
}
