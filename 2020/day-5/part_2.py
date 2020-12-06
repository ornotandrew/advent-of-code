from itertools import product
from seats import get_seat_id

with open("input.txt", "r") as f:
    all_bps = f.read().splitlines()

cols_by_row = {}

for bp in all_bps:
    row, col = bp[:7], bp[7:]
    if not row in cols_by_row:
        cols_by_row[row] = []
    cols_by_row[row].append(col)

# we want to exclude the first and last rows
middle_rows = sorted(cols_by_row.keys())[1:-1]

my_row = next(row for row in middle_rows if len(cols_by_row[row]) < 8)

possible_cols = ["".join(chars) for chars in product("LR", repeat=3)]
my_col = next(col for col in possible_cols if col not in cols_by_row[my_row])

print(get_seat_id(my_row + my_col))
