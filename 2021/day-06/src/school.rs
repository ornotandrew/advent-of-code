pub struct School {
    groups: [usize; 9],
}

impl School {
    pub fn new(initial_fish: Vec<usize>) -> School {
        let mut groups: [usize; 9] = [0; 9];
        for t in initial_fish {
            groups[t] += 1
        }

        School { groups }
    }

    pub fn simulate_day(&mut self) {
        let mut next_day: [usize; 9] = [0; 9];
        for (i, group_count) in self.groups.iter().enumerate().rev() {
            if i == 0 {
                next_day[6] += group_count;
                next_day[8] += group_count;
                continue;
            }
            next_day[i - 1] = *group_count;
        }
        self.groups = next_day;
    }

    pub fn get_total_fish(&self) -> usize {
        self.groups.iter().sum()
    }
}
