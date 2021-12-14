mod bingo;
use bingo::Game;
use std::io::{self, BufRead};

fn main() {
    let input: Vec<String> = io::stdin().lock().lines().map(|l| l.unwrap()).collect();

    let numbers: Vec<usize> = (&input[0]).split(",").map(|n| n.parse().unwrap()).collect();

    let mut game = Game::from_lines(&input[2..]);
    let mut num_winners = 0;
    let mut winners: Vec<bool> = vec![];
    for _ in 0..game.boards.len() {
        winners.push(false);
    }

    let mut get_final_score = || -> Option<usize> {
        for i in 0..numbers.len() {
            game.draw_number(numbers[i]);
            // check if we have only one board remaining
            for j in 0..winners.len() {
                if winners[j] {
                    continue;
                }
                if game.boards[j].has_won() {
                    winners[j] = true;
                    num_winners += 1;
                    if num_winners == winners.len() {
                        return Some(game.boards[j].get_score() * numbers[i]);
                    }
                }
            }
        }
        return None;
    };

    let answer = get_final_score();
    match answer {
        Some(answer) => println!("{}", answer),
        None => println!("None"),
    }
}
