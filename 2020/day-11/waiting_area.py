from itertools import product


def apply_rules(seat, num_adjacent, people_are_tolerant=False):
    if seat == ".":
        return "."
    elif num_adjacent == 0 and seat == "L":
        return "#"
    elif num_adjacent >= (5 if people_are_tolerant else 4) and seat == "#":
        return "L"
    else:
        return seat


class WaitingArea:
    def __init__(self, use_visible_rules=False):
        self.use_visible_rules = use_visible_rules
        with open("input.txt", "r") as f:
            self.rows = f.read().splitlines()

    def is_within_area(self, coords):
        r, c = coords
        return 0 <= r < len(self.rows) and 0 <= c < len(self.rows[0])

    def count_occupied_adjacent_seats(self):
        deltas = [-1, 0, 1]
        counts = [[0] * len(self.rows[0]) for row in self.rows]
        for r, row in enumerate(self.rows):
            for c, col in enumerate(row):
                if col != "#":
                    continue

                for d_r, d_c in product(deltas, repeat=2):
                    coords = (r + d_r, c + d_c)
                    if (d_r, d_c) != (0, 0) and self.is_within_area(coords):
                        counts[coords[0]][coords[1]] += 1
        return counts

    def count_occupied_visible_seats(self):
        deltas = [-1, 0, 1]
        counts = [[0] * len(self.rows[0]) for row in self.rows]

        def increment_in_direction(coords, deltas):
            r, c = coords
            d_r, d_c = deltas
            r_prime, c_prime = r + d_r, c + d_c
            while self.is_within_area((r_prime, c_prime)):
                counts[r_prime][c_prime] += 1
                if self.rows[r_prime][c_prime] != ".":
                    break
                r_prime, c_prime = r_prime + d_r, c_prime + d_c

        for r, row in enumerate(self.rows):
            for c, col in enumerate(row):
                if col != "#":
                    continue

                for d_r, d_c in product(deltas, repeat=2):
                    if (d_r, d_c) != (0, 0):
                        increment_in_direction((r, c), (d_r, d_c))

        return counts

    def __next__(self):
        counts = self.count_occupied_visible_seats() if self.use_visible_rules else self.count_occupied_adjacent_seats()
        new_rows = [
            "".join(
                [
                    apply_rules(col, counts[r][c], people_are_tolerant=self.use_visible_rules)
                    for c, col in enumerate(row)
                ]
            )
            for r, row in enumerate(self.rows)
        ]
        self.rows = new_rows
        return str(self)

    def __str__(self):
        return "\n".join(self.rows)
