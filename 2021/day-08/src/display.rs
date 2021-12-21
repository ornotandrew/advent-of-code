use std::collections::HashSet;

use crate::segment::Segment;

pub type Pattern = HashSet<Segment>;

#[derive(Debug)]
pub struct Display {
    pub signal_patterns: [Pattern; 10],
    pub output_values: [Pattern; 4],
}

impl Display {
    pub fn new(line: String) -> Display {
        let parts: Vec<Vec<Pattern>> = line
            .split(" | ")
            .map(|part| -> Vec<HashSet<Segment>> {
                let strings: Vec<String> = part.split(" ").map(|s| s.to_string()).collect();
                strings
                    .iter()
                    .map(|s| {
                        s.chars()
                            .map(|char| Segment::from_str(char))
                            .collect::<HashSet<Segment>>()
                    })
                    .collect()
            })
            .collect();

        Display {
            signal_patterns: parts[0].clone().try_into().unwrap(),
            output_values: parts[1].clone().try_into().unwrap(),
        }
    }

    fn get_pattern(
        &self,
        length: usize,
        containing: Option<Vec<&Pattern>>,
        not_containing: Option<Vec<&Pattern>>,
        contained_by: Option<Vec<&Pattern>>,
    ) -> Pattern {
        self.signal_patterns
            .iter()
            .filter(|p| p.len() == length)
            .filter(|p| match &containing {
                Some(containing) => containing
                    .iter()
                    .all(|c| p.intersection(c).cloned().collect::<Pattern>().len() == c.len()),
                None => true,
            })
            .filter(|p| match &not_containing {
                Some(not_containing) => not_containing
                    .iter()
                    .all(|nc| p.intersection(nc).cloned().collect::<Pattern>().len() != nc.len()),
                None => true,
            })
            .filter(|p| match &contained_by {
                Some(contained_by) => contained_by
                    .iter()
                    .all(|cb| p.intersection(cb).cloned().collect::<Pattern>().len() == p.len()),
                None => true,
            })
            .cloned()
            .next()
            .unwrap()
    }

    pub fn derive_numbers(&self) -> [Pattern; 10] {
        // First, find numbers that have a unique number of segments
        // ============================================================
        let one = self.get_pattern(2, None, None, None);
        let four = self.get_pattern(4, None, None, None);
        let seven = self.get_pattern(3, None, None, None);
        let eight = self.get_pattern(7, None, None, None);

        // Then, use this information to find the other numbers by checking
        // if they contain the smaller numbers. For example, the number "3"
        // contains the number "1" (segments C and F).
        // ============================================================
        let three = self.get_pattern(5, Some(vec![&one]), None, None);
        let nine = self.get_pattern(6, Some(vec![&four]), None, None);

        // The remaining numbers can be found using the same logic, and
        // using the numbers we've already found. For example, the
        // number "0" is the only number with 6 segments which
        // - contains "1"
        // - does not contain "9"
        // ============================================================
        let zero = self.get_pattern(6, Some(vec![&one]), Some(vec![&nine]), None);
        let six = self.get_pattern(6, None, Some(vec![&zero, &nine]), None);
        let five = self.get_pattern(5, None, Some(vec![&three]), Some(vec![&six]));
        let two = self.get_pattern(5, None, Some(vec![&three, &five]), None);

        [zero, one, two, three, four, five, six, seven, eight, nine]
    }
}
