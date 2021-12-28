use std::{
    cmp::{max, min},
    str::FromStr,
};

#[derive(Debug)]
pub struct Range {
    min: isize,
    max: isize,
}

#[derive(Debug)]
pub struct Target {
    x: Range,
    y: Range,
}

#[derive(Debug)]
pub struct Trench {
    target: Target,
}

#[derive(Debug)]
pub struct Trajectory {
    vx: isize,
    vy: isize,
}

impl Trench {
    fn launch(&self, t: Trajectory) -> bool {
        let (mut x, mut y) = (0_isize, 0_isize);
        let mut i = 0;
        loop {
            x += max(t.vx - i, 0);
            y += t.vy - i;
            if x >= self.target.x.min
                && x <= self.target.x.max
                && y >= self.target.y.min
                && y <= self.target.y.max
            {
                return true;
            }
            if x > self.target.x.max || y < self.target.y.min {
                return false;
            }

            i += 1;
        }
    }

    pub fn highest_possible_y_value(&self) -> usize {
        let n = -self.target.y.min - 1;
        (n * (n + 1) / 2) as usize
    }

    pub fn all_trajectories_for_target(&self) -> usize {
        let mut count: usize = 0;
        let x_min = (self.target.x.min as f64).sqrt().floor() as isize;
        let x_max = self.target.x.max;
        let y_min = self.target.y.min;
        let y_max = self.target.y.min.abs();

        for vx in x_min..x_max + 1 {
            for vy in y_min..y_max {
                let t = Trajectory { vx, vy };
                if self.launch(t) {
                    count += 1;
                }
            }
        }

        count
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Trench {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (x, y) = s
            .trim()
            .split_once(": ")
            .unwrap()
            .1
            .split_once(", ")
            .unwrap();
        let (x_min, x_max) = &x[2..].split_once("..").unwrap();
        let (y_min, y_max) = &y[2..].split_once("..").unwrap();
        Ok(Trench {
            target: Target {
                x: Range {
                    min: x_min.parse().unwrap(),
                    max: x_max.parse().unwrap(),
                },
                y: Range {
                    min: y_min.parse().unwrap(),
                    max: y_max.parse().unwrap(),
                },
            },
        })
    }
}
