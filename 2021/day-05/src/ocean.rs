use std::{cmp::Ordering, collections::HashMap, fmt};

pub type Vent = ((isize, isize), (isize, isize));
pub type Position = (isize, isize);

pub struct Ocean {
    pub vents: Vec<Vent>,
    pub position_counts: HashMap<Position, usize>,
}

impl Ocean {
    pub fn load_from_lines(lines: Vec<String>, include_diagonal_vents: bool) -> Ocean {
        let vents: Vec<Vent> = lines
            .iter()
            .map(|l| -> Vent {
                let parts: Vec<Vec<isize>> = l
                    .split(" -> ")
                    .take(2)
                    .map(|part| {
                        part.split(",")
                            .map(|n| n.parse().unwrap())
                            .take(2)
                            .collect::<Vec<isize>>()
                    })
                    .collect();
                ((parts[0][0], parts[0][1]), (parts[1][0], parts[1][1]))
            })
            .collect();

        let mut position_counts: HashMap<Position, usize> = HashMap::new();
        for vent in vents.iter() {
            for p in get_all_positions(&vent.0, &vent.1, include_diagonal_vents) {
                let existing_count = &position_counts.get(&p);
                match existing_count {
                    Some(&count) => position_counts.insert(p, count + 1),
                    None => position_counts.insert(p, 1),
                };
            }
        }

        Ocean {
            vents,
            position_counts,
        }
    }
}

impl fmt::Display for Ocean {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> fmt::Result {
        let (max_x, max_y) = self
            .position_counts
            .keys()
            .cloned()
            .reduce(|acc, (x, y)| {
                let (mut max_x, mut max_y) = acc;
                if x > max_x {
                    max_x = x;
                }
                if y > max_y {
                    max_y = y;
                }
                (max_x, max_y)
            })
            .unwrap();

        for y in 0..max_y + 1 {
            for x in 0..max_x + 1 {
                match self.position_counts.get(&(x, y)) {
                    Some(count) => write!(f, "{}", count).unwrap(),
                    None => write!(f, ".").unwrap(),
                };
            }
            write!(f, "\n").unwrap();
        }

        Ok(())
    }
}

// this assumes the lines are either horizontal or vertical
fn get_all_positions(
    start: &Position,
    end: &Position,
    include_diagonal_vents: bool,
) -> Vec<Position> {
    if !include_diagonal_vents && !(start.0 == end.0 || start.1 == end.1) {
        return vec![];
    }

    let mut positions: Vec<Position> = vec![];

    let get_step = |start: &isize, end: &isize| -> isize {
        match start.cmp(end) {
            Ordering::Less => 1,
            Ordering::Equal => 0,
            Ordering::Greater => -1,
        }
    };

    let x_step = get_step(&start.0, &end.0);
    let y_step = get_step(&start.1, &end.1);

    let mut current = start.clone();
    while current != *end {
        positions.push(current);
        current = (current.0 + x_step, current.1 + y_step);
    }
    positions.push(current);
    positions
}
