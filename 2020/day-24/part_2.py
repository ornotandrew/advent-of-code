from floor import small_lines, step, follow_directions
from part_1 import black_tiles

grid = black_tiles

for i in range(100):
    grid = step(grid)

print(len(grid))
