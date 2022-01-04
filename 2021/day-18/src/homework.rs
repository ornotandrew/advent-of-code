use std::str::FromStr;

pub type Num = (u32, u32); // (depth, num)
pub type Term = Vec<Num>;

#[derive(Debug)]
pub struct Homework {
    pub terms: Vec<Term>,
}

impl Homework {
    fn _reduce(t: &mut Term) {
        // first, explode
        for i in 0..t.len() - 1 {
            let depth = t[i].0;
            if depth == 5 {
                let (l, r) = (t[i].1, t[i + 1].1);
                t[i] = (4, 0);
                t.remove(i + 1);
                if i > 0 {
                    t[i - 1] = (t[i - 1].0, t[i - 1].1 + l);
                }
                if i + 1 < t.len() {
                    t[i + 1] = (t[i + 1].0, t[i + 1].1 + r);
                }
                return Homework::_reduce(t);
            }
        }
        // then, split
        for i in 0..t.len() {
            let (depth, num) = t[i];
            if num > 9 {
                t[i] = (depth + 1, num / 2);
                t.insert(i + 1, (depth + 1, (num + 1) / 2));
                return Homework::_reduce(t);
            }
        }
    }

    fn _add(a: &Term, b: &Term) -> Term {
        let mut t: Term = vec![];
        for (depth, num) in a {
            t.push((depth + 1, *num));
        }
        for (depth, num) in b {
            t.push((depth + 1, *num));
        }
        t
    }

    pub fn reduce(&mut self) -> Term {
        // first, add up all the terms
        let mut t: Term = self.terms[0].clone();
        for i in 1..self.terms.len() {
            t = Homework::_add(&t, &self.terms[i]);
            Homework::_reduce(&mut t);
        }

        t
    }

    fn _magnitude(i: &mut usize, depth: u32, term: &Term) -> u32 {
        let mut mag = 0;
        // left
        if term[*i].0 == depth {
            mag += 3 * term[*i].1;
            *i += 1;
        } else {
            mag += 3 * Homework::_magnitude(i, depth + 1, term);
        }
        // right
        if term[*i].0 == depth {
            mag += 2 * term[*i].1;
            *i += 1;
        } else {
            mag += 2 * Homework::_magnitude(i, depth + 1, term);
        }

        mag
    }

    pub fn magnitude(t: &Term) -> u32 {
        Homework::_magnitude(&mut 0, 1, &t)
    }

    pub fn largest_magnitude(&self) -> u32 {
        let mut max = 0;
        for i in 0..self.terms.len() {
            for j in i + 1..self.terms.len() {
                for (i_l, i_r) in vec![(i, j), (j, i)] {
                    let mut sum = Homework::_add(&self.terms[i_l], &self.terms[i_r]);
                    Homework::_reduce(&mut sum);
                    let mag = Homework::magnitude(&sum);
                    if mag > max {
                        max = mag;
                    }
                }
            }
        }

        max
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Homework {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Homework {
            terms: s
                .trim()
                .split("\n")
                .map(|line| {
                    let mut term: Term = vec![];
                    // we are only given digits 0-9 in the input
                    let mut depth = 0;
                    for c in line.chars() {
                        match c {
                            '[' => depth += 1,
                            ']' => depth -= 1,
                            '0'..='9' => term.push((depth, c.to_digit(10).unwrap())),
                            _ => {}
                        }
                    }
                    term
                })
                .collect(),
        })
    }
}

pub fn print_term(t: &Term) {
    let mut current_depth = 0;
    for (depth, num) in t {
        while current_depth < *depth {
            print!("[");
            current_depth += 1;
        }
        while current_depth > *depth {
            print!("]");
            current_depth -= 1;
        }
        print!("{},", num);
    }

    while current_depth > 0 {
        print!("]");
        current_depth -= 1;
    }
    print!("\n");
}
