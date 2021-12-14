use std::io::{self, BufRead};

fn main() {
    let mut num_bits: u32 = 0;
    let input: Vec<u32> = io::stdin()
        .lock()
        .lines()
        .map(|l| {
            let line = l.unwrap();
            let len: u32 = line.len().try_into().unwrap();
            if len > num_bits {
                num_bits = len;
            }
            u32::from_str_radix(line.as_str(), 2).unwrap()
        })
        .collect();

    let mut gamma = 0;
    for bit in 0..num_bits {
        let mut acc = 0;
        for i in &input {
            acc += if 1 << bit & i > 0 { 1 } else { -1 }
        }

        if acc > 0 {
            gamma += 1 << bit;
        }
    }

    let mask = u32::pow(2, num_bits) - 1;
    let answer = gamma * (gamma ^ mask);

    println!("{}", answer);
}
