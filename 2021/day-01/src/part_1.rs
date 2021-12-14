use std::io::{self, BufRead};

fn main() {
    let mut answer = 0;
    let mut prev_num = 0;
    for (i, line) in io::stdin().lock().lines().enumerate() {
        let num: i32 = line.unwrap().parse().unwrap();
        if i > 0 && num > prev_num {
            answer += 1;
        }
        prev_num = num;
    }

    println!("{}", answer)
}
