use std::{collections::HashMap, str::FromStr};

#[derive(Debug)]
pub struct Polymer {
    last_item: String,
    rules: HashMap<String, String>,
    pair_counts: HashMap<String, usize>,
}

impl Polymer {
    pub fn step(&mut self) {
        let mut new_counts: HashMap<String, usize> = HashMap::new();
        for (pair, count) in &self.pair_counts {
            match self.rules.get(pair) {
                Some(to_insert) => {
                    let left = (&pair[..1]).to_string() + to_insert;
                    *new_counts.entry(left).or_insert(0) += count;

                    let right = to_insert.to_string() + &(&pair[1..]).to_string();
                    *new_counts.entry(right).or_insert(0) += count;
                }
                None => {}
            };
        }
        self.pair_counts = new_counts;
    }

    pub fn get_difference(&self) -> usize {
        let mut counts: HashMap<&str, usize> = HashMap::new();
        for (pair, count) in &self.pair_counts {
            *counts.entry(&pair[..1]).or_insert(0) += count;
        }
        *counts.entry(&self.last_item).or_insert(0) += 1;

        let max = counts.values().max().unwrap();
        let min = counts.values().min().unwrap();

        max - min
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Polymer {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let lines: Vec<&str> = s.trim().split("\n").collect();

        let sequence = lines[0];
        let mut pair_counts: HashMap<String, usize> = HashMap::new();
        for pos in 0..sequence.len() - 1 {
            let pair = &sequence[pos..pos + 2];
            *pair_counts.entry(pair.to_string()).or_insert(0) += 1;
        }

        let mut rules: HashMap<String, String> = HashMap::new();
        for line in &lines[2..] {
            let (pair, between) = line.split_once(" -> ").unwrap();
            rules.insert(pair.to_string(), between.to_string());
        }

        Ok(Polymer {
            rules,
            pair_counts,
            last_item: sequence.chars().last().unwrap().to_string(),
        })
    }
}
