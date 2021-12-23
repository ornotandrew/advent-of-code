use std::{
    collections::{HashMap, HashSet},
    str::FromStr,
};

#[derive(Debug)]
pub struct Syntax {
    opening_chars: HashSet<char>,
    pairs: HashMap<char, char>,
    lines: Vec<String>,
}

impl Syntax {
    pub fn find_corrupted(&self) -> Vec<char> {
        let mut corrupted: Vec<char> = vec![];
        for line in self.lines.iter() {
            let mut stack: Vec<char> = vec![];
            for char in line.chars() {
                if self.opening_chars.contains(&char) {
                    stack.push(char);
                    continue;
                }
                if &stack.pop().unwrap() != self.pairs.get(&char).unwrap() {
                    corrupted.push(char);
                    break;
                }
            }
        }

        corrupted
    }

    pub fn corruption_score(corruptions: Vec<char>) -> u32 {
        let scores: HashMap<char, u32> = vec![(')', 3), (']', 57), ('}', 1197), ('>', 25137)]
            .into_iter()
            .collect();
        corruptions
            .iter()
            .fold(0, |acc, cur| acc + scores.get(&cur).unwrap())
    }

    pub fn find_incomplete(&self) -> Vec<String> {
        let mut incomplete: Vec<String> = vec![];
        'line: for line in self.lines.iter() {
            let mut stack: Vec<char> = vec![];
            for char in line.chars() {
                if self.opening_chars.contains(&char) {
                    stack.push(char);
                    continue;
                }
                if &stack.pop().unwrap() != self.pairs.get(&char).unwrap() {
                    // this line is corrupted - ignore it
                    continue 'line;
                }
            }
            incomplete.push(stack.iter().collect());
        }

        incomplete
            .iter()
            .map(|l| {
                l.chars()
                    .rev()
                    .map(|c| self.pairs.get(&c).unwrap())
                    .collect()
            })
            .collect()
    }

    pub fn incomplete_scores(incomplete: Vec<String>) -> Vec<u64> {
        let scores: HashMap<char, u32> = vec![(')', 1), (']', 2), ('}', 3), ('>', 4)]
            .into_iter()
            .collect();
        incomplete
            .iter()
            .map(|line| {
                line.chars().fold(0_u64, |acc, cur| {
                    (acc * 5) + *scores.get(&cur).unwrap() as u64
                })
            })
            .collect()
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Syntax {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Syntax {
            opening_chars: vec!['(', '[', '{', '<'].into_iter().collect(),
            pairs: vec![
                ('(', ')'),
                (')', '('),
                ('[', ']'),
                (']', '['),
                ('{', '}'),
                ('}', '{'),
                ('<', '>'),
                ('>', '<'),
            ]
            .into_iter()
            .collect(),
            lines: s.trim().split('\n').map(|l| l.to_string()).collect(),
        })
    }
}
