from slope import Slope

cases = [
    (1, 1),
    (1, 3),
    (1, 5),
    (1, 7),
    (2, 1),
]

slope = Slope()

answer = 1

for d_row, d_col in cases:
    row, col = 0, 0
    count = 0
    while row < len(slope.trees):
        if slope[row, col]:
            count += 1
        row += d_row
        col += d_col
    answer *= count

print(answer)
