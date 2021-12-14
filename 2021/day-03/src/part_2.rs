use std::io::{self, BufRead};

fn filter_remaining_numbers(
    numbers: Vec<u32>,
    col_from_right: u32,
    use_dominant_bit: bool,
) -> Vec<u32> {
    // first, calculate the dominant bit in this column (from the remaining numbers)
    let mut acc: i32 = 0;
    for i in &numbers {
        acc += if (i >> (col_from_right - 1) & 1) > 0 {
            1
        } else {
            -1
        };
    }
    let relevant_bit: u32 = if acc >= 0 {
        if use_dominant_bit {
            1
        } else {
            0
        }
    } else {
        if use_dominant_bit {
            0
        } else {
            1
        }
    };

    numbers
        .into_iter()
        .filter(|n| (n >> (col_from_right - 1) & 1) == relevant_bit)
        .collect::<Vec<u32>>()
}

fn get_o2_and_co2(numbers: Vec<u32>, num_bits: u32) -> (u32, u32) {
    let mut col_from_left = 0;
    let mut o2_numbers = numbers.clone();
    while o2_numbers.len() > 1 {
        o2_numbers = filter_remaining_numbers(o2_numbers, num_bits - col_from_left, true);
        col_from_left += 1;
    }

    col_from_left = 0;
    let mut co2_numbers = numbers.clone();
    while co2_numbers.len() > 1 {
        co2_numbers = filter_remaining_numbers(co2_numbers, num_bits - col_from_left, false);
        col_from_left += 1;
    }

    return (o2_numbers[0], co2_numbers[0]);
}

fn main() {
    let mut num_bits: u32 = 0;
    let input: Vec<u32> = io::stdin()
        .lock()
        .lines()
        .enumerate()
        .map(|(i, l)| {
            let line = l.unwrap();
            let len: u32 = line.len().try_into().unwrap();
            if i == 0 {
                num_bits = len;
            }
            u32::from_str_radix(line.as_str(), 2).unwrap()
        })
        .collect();

    let (o2, co2) = get_o2_and_co2(input, num_bits);
    println!("{}", o2 * co2);
}
