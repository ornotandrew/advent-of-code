mod bingo;
use bingo::Game;
use std::io::{self, BufRead};

fn main() {
    let input: Vec<String> = io::stdin().lock().lines().map(|l| l.unwrap()).collect();

    let numbers: Vec<usize> = (&input[0]).split(",").map(|n| n.parse().unwrap()).collect();

    let mut game = Game::from_lines(&input[2..]);
    let mut score: Option<usize> = None;
    let mut i = 0;
    while score == None || i >= numbers.len() {
        game.draw_number(numbers[i]);
        for board in &game.boards {
            if board.has_won() {
                score = Some(board.get_score());
            }
        }
        i += 1;
    }

    match score {
        Some(score) => println!("{}", score * numbers[i - 1]),
        None => {}
    }
}
