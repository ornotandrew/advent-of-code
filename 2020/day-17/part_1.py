from life_3d import read_initial_grid, step

grid = read_initial_grid(3)

for i in range(6):
    grid = step(grid)

print(len(grid))
