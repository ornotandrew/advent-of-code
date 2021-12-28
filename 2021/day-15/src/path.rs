use std::cmp::Ordering;

use crate::cave::Position;

#[derive(Copy, Clone)]
pub struct Path {
    pub pos: Position,
    pub cost: usize,
    pub end: Position,
}

impl Path {
    pub fn f(&self) -> usize {
        self.cost + (self.end.row - self.pos.row) + (self.end.col - self.pos.col)
    }
}

impl Eq for Path {}

impl Ord for Path {
    fn cmp(&self, other: &Self) -> Ordering {
        other.f().cmp(&self.f())
    }
}

impl PartialOrd for Path {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl PartialEq for Path {
    fn eq(&self, other: &Self) -> bool {
        self.cost == other.cost
    }
}
