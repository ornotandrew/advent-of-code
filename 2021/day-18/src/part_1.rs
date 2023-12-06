mod homework;
use std::io::{self, BufRead};

use homework::Term;

use crate::homework::Homework;

fn expect_reduction(input: &str, expected: &str) {
    let mut homework: Homework = input.parse().unwrap();
    let actual = homework.reduce();
    let parsed: Term = expected.parse::<Homework>().unwrap().terms[0].clone();
    assert_eq!(actual, parsed);
}

fn expect_magnitude(input: &str, mag: usize) {
    let homework: Homework = input.parse().unwrap();
    Homework::magnitude(&homework.terms[0]);
}

fn main() {
    // tests
    expect_reduction(
        "[[[[4,3],4],4],[7,[[8,4],9]]]\n[1,1]",
        "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
    );
    expect_reduction(
        "[1,1]\n[2,2]\n[3,3]\n[4,4]",
        "[[[[1,1],[2,2]],[3,3]],[4,4]]",
    );
    expect_reduction(
        "[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]",
        "[[[[3,0],[5,3]],[4,4]],[5,5]]",
    );
    expect_reduction(
        "[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]\n[6,6]",
        "[[[[5,0],[7,4]],[5,5]],[6,6]]",
    );
    expect_reduction(
        "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
        "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
    );
    expect_reduction(
        "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
        "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"
    );
    expect_reduction(
        "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
        "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"
    );
    expect_reduction(
        "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]\n[7,[5,[[3,8],[1,4]]]]",
        "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
    );
    expect_reduction(
        "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]\n[[2,[2,2]],[8,[8,1]]]",
        "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
    );
    expect_reduction(
        "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]\n[2,9]",
        "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
    );
    expect_reduction(
        "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]\n[1,[[[9,3],9],[[9,0],[0,7]]]]",
        "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
    );
    expect_reduction(
        "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]\n[[[5,[7,4]],7],1]",
        "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
    );
    expect_reduction(
        "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]\n[[[[4,2],2],6],[8,7]]",
        "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
    );
    expect_reduction(
        "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]\n[7,[5,[[3,8],[1,4]]]]\n[[2,[2,2]],[8,[8,1]]]\n[2,9]\n[1,[[[9,3],9],[[9,0],[0,7]]]]\n[[[5,[7,4]],7],1]\n[[[[4,2],2],6],[8,7]]",
        "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
    );

    expect_magnitude("[9,1]", 29);
    expect_magnitude("[[9,1],[1,9]]", 129);
    expect_magnitude("[[9,1],[1,9]]", 29);
    expect_magnitude("[[1,2],[[3,4],5]]", 143);
    expect_magnitude("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384);
    expect_magnitude("[[[[1,1],[2,2]],[3,3]],[4,4]]", 445);
    expect_magnitude("[[[[3,0],[5,3]],[4,4]],[5,5]]", 791);
    expect_magnitude("[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137);
    expect_magnitude(
        "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
        3488,
    );

    let input: String = io::stdin().lock().lines().fold(String::new(), |acc, cur| {
        format!("{}\n{}", acc, cur.unwrap())
    });

    let mut homework: Homework = input.parse().unwrap();
    println!("{}", Homework::magnitude(&homework.reduce()));
}